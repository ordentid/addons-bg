package api

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"bitbucket.bri.co.id/scm/addons/addons-bg-service/server/db"
	manager "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/jwt"
	company_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/company"
	notification_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/notification"
	system_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/system"
	task_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/task"
	workflow_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/workflow"
	pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (s *Server) GetCurrency(ctx context.Context, req *pb.GetCurrencyRequest) (*pb.GetCurrencyResponse, error) {

	result := &pb.GetCurrencyResponse{
		Error:   false,
		Code:    200,
		Message: "List Data",
		Data:    []*pb.Currency{},
	}

	data, err := s.provider.GetCurrency(ctx, &db.ListFilter{})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	if len(data) > 0 {
		for _, v := range data {
			currency, err := v.ToPB(ctx)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
			}

			result.Data = append(result.Data, &currency)
		}
	}

	return result, nil

}

func (s *Server) GetApplicantName(ctx context.Context, req *pb.GetApplicantNameRequest) (*pb.GetApplicantNameResponse, error) {

	result := &pb.GetApplicantNameResponse{
		Error:   false,
		Code:    200,
		Message: "List Data",
		Data:    []*pb.ApplicantName{},
	}

	var newCtx context.Context

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		newCtx = metadata.NewOutgoingContext(context.Background(), md)
	}

	currentUser, userMD, err := s.manager.GetMeFromMD(ctx)
	if err != nil {
		return nil, err
	}
	if currentUser == nil {
		return nil, s.UnauthorizedError()
	}
	var trailer metadata.MD

	taskClient := s.svcConn.TaskServiceClient()

	taskFilter := &task_pb.Task{
		Type:      "BG Issuing",
		CompanyID: currentUser.CompanyID,
	}

	filter := []string{
		"data.publishing.thirdPartyID:" + strconv.FormatUint(req.ThirdPartyID, 10),
	}

	dataReq := &task_pb.ListTaskRequest{
		Task:   taskFilter,
		Filter: strings.Join(filter, ","),
	}

	dataList, err := taskClient.GetListTask(newCtx, dataReq, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	names := []string{}

	if len(dataList.Data) > 0 {

		for _, v := range dataList.Data {

			taskData := pb.IssuingData{}
			err = json.Unmarshal([]byte(v.Data), &taskData)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
			}

			if !contains(names, taskData.Applicant.Name) {
				names = append(names, taskData.Applicant.Name)
			}

		}

	}

	if len(names) > 0 {

		for _, v := range names {

			result.Data = append(result.Data, &pb.ApplicantName{Name: v})

		}

	}

	return result, nil

}

func (s *Server) GetBeneficiaryName(ctx context.Context, req *pb.GetBeneficiaryNameRequest) (*pb.GetBeneficiaryNameResponse, error) {

	result := &pb.GetBeneficiaryNameResponse{
		Error:   false,
		Code:    200,
		Message: "List Data",
		Data:    []*pb.BeneficiaryName{},
	}

	if req.ThirdPartyID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Bad Request: Third Party ID is required")
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		ctx = metadata.NewOutgoingContext(context.Background(), md)
	}

	currentUser, _, err := s.manager.GetMeFromMD(ctx)
	if err != nil {
		return nil, err
	}
	if currentUser == nil {
		return nil, s.UnauthorizedError()
	}

	apiReq := &ApiInquiryBenficiaryRequest{
		ThirdPartyID: req.ThirdPartyID,
	}

	res, err := s.ApiInquiryBeneficiary(ctx, apiReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	data := []*pb.BeneficiaryName{}

	if res.ResponseCode == "00" {

		if len(res.ResponseData) > 0 {

			for _, v := range res.ResponseData {
				data = append(data, &pb.BeneficiaryName{
					BeneficiaryId: v.BeneficiaryID,
					ThirdPartyId:  v.ThirdPartyID,
					Cif:           v.Cif,
					Fullname:      v.FullName,
					Status:        v.Status,
				})
			}

		}

	}

	if req.Type == 0 {

		result.Data = data

	} else {

		mappedBeneficiaryIDs := []string{}

		mappingFilter := []string{
			"company_id:" + strconv.FormatUint(currentUser.CompanyID, 10),
			"third_party_id:" + strconv.FormatUint(req.ThirdPartyID, 10),
			"is_mapped:true",
		}

		filter := &db.ListFilter{
			Filter: strings.Join(mappingFilter, ","),
		}

		mappingORMs, err := s.provider.GetMapping(ctx, filter)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		if len(mappingORMs) > 0 {

			for _, v := range mappingORMs {

				if v.BeneficiaryID != 10101010 {

					if !contains(mappedBeneficiaryIDs, strconv.FormatUint(v.BeneficiaryID, 10)) {
						mappedBeneficiaryIDs = append(mappedBeneficiaryIDs, strconv.FormatUint(v.BeneficiaryID, 10))
					}

				} else {

					for _, d := range data {

						if !contains(mappedBeneficiaryIDs, strconv.FormatUint(d.BeneficiaryId, 10)) {
							mappedBeneficiaryIDs = append(mappedBeneficiaryIDs, strconv.FormatUint(d.BeneficiaryId, 10))
						}

					}

				}

			}

		}

		if len(data) > 0 {

			for _, v := range data {

				if contains(mappedBeneficiaryIDs, strconv.FormatUint(v.BeneficiaryId, 10)) {
					result.Data = append(result.Data, v)
				}

			}

		}

	}

	return result, nil

}

func (s *Server) GetThirdParty(ctx context.Context, req *pb.GetThirdPartyRequest) (*pb.GetThirdPartyResponse, error) {

	result := &pb.GetThirdPartyResponse{
		Error:   false,
		Code:    200,
		Message: "List Data",
		Data:    []*pb.ThirdParty{},
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		ctx = metadata.NewOutgoingContext(context.Background(), md)
	}

	currentUser, _, err := s.manager.GetMeFromMD(ctx)
	if err != nil {
		return nil, err
	}
	if currentUser == nil {
		return nil, s.UnauthorizedError()
	}

	logrus.Println("==========> User Type:", currentUser.UserType)

	if currentUser.UserType == "ba" {

		apiReq := &ApiInquiryThirdPartyByStatusRequest{
			Status: "Active",
		}

		res, err := s.ApiInquiryThirdPartyByStatus(ctx, apiReq)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		if res.ResponseCode == "00" {

			if len(res.ResponseData) > 0 {

				for _, v := range res.ResponseData {
					result.Data = append(result.Data, &pb.ThirdParty{
						Id:   v.ThirdPartyID,
						Name: v.FullName,
					})
				}

			}

		}

	} else {

		if req.Type != pb.ThirdPartyType_All {

			filter := &db.ListFilter{}

			filterMapped := []string{
				"company_id:" + strconv.FormatUint(currentUser.CompanyID, 10),
			}
			if req.Type == pb.ThirdPartyType_NeedMapping {
				filterMapped = append(filterMapped, "is_mapped:false")
			} else if req.Type == pb.ThirdPartyType_IsMapped {
				filterMapped = append(filterMapped, "is_mapped:true")
			}

			filter.Filter = strings.Join(filterMapped, ",")
			logrus.Println("==========> Mapping Filter:", filter.Filter)

			thirdPartyNameList, err := s.provider.GetMapping(ctx, filter)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
			}

			logrus.Println("==========> ThirdParty List:", thirdPartyNameList)

			ids := []string{}

			if len(thirdPartyNameList) > 0 {
				for _, v := range thirdPartyNameList {
					if !contains(ids, strconv.FormatUint(v.ThirdPartyID, 10)) {
						ids = append(ids, strconv.FormatUint(v.ThirdPartyID, 10))
					}
				}
			}

			if len(ids) > 0 {

				for _, v := range ids {

					id, err := strconv.ParseUint(v, 10, 64)
					if err != nil {
						return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
					}

					name := ""

					apiReq := &ApiInquiryThirdPartyByIDRequest{
						ThirdPartyID: id,
					}

					res, err := s.ApiInquiryThirdPartyByID(ctx, apiReq)
					if err != nil {
						return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
					}

					if res.ResponseCode == "00" {
						name = res.ResponseData.FullName
					}

					result.Data = append(result.Data, &pb.ThirdParty{
						Id:   id,
						Name: name,
					})

				}

			}

		} else {

			apiReq := &ApiInquiryThirdPartyByStatusRequest{
				Status: "Active",
			}

			res, err := s.ApiInquiryThirdPartyByStatus(ctx, apiReq)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
			}

			if res.ResponseCode == "00" {

				if len(res.ResponseData) > 0 {

					for _, v := range res.ResponseData {
						result.Data = append(result.Data, &pb.ThirdParty{
							Id:   v.ThirdPartyID,
							Name: v.FullName,
						})
					}

				}

			}

		}

	}

	return result, nil

}

func (s *Server) GetCustomerLimit(ctx context.Context, req *pb.GetCustomerLimitRequest) (*pb.GetCustomerLimitResponse, error) {

	result := &pb.GetCustomerLimitResponse{
		Error:   false,
		Code:    200,
		Message: "List Data",
		Data:    []*pb.CustomerLimit{},
	}

	apiReq := &ApiInquiryLimitIndividualRequest{}

	res, err := s.ApiInquiryLimitIndividual(ctx, apiReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	if res.ResponseCode != "00" {
		return nil, status.Errorf(codes.Internal, string(*res.ResponseMessage))
	}

	if len(res.ResponseData) > 0 {

		for _, v := range res.ResponseData {
			result.Data = append(result.Data, &pb.CustomerLimit{
				CustomerLimitId:   v.CustomerLimitId,
				Code:              v.Code,
				Fullname:          v.Fullname,
				Cif:               v.Cif,
				PtkNo:             v.PtkNo,
				Currency:          v.Currency,
				Plafond:           v.Plafond,
				ReservationAmount: v.ReservationAmount,
				OutstandingAmount: v.OutstandingAmount,
				AvailableAmount:   v.AvailableAmount,
				ExpiryDate:        v.ExpiryDate,
				PnRm:              v.PnRm,
				NameRm:            v.NameRm,
				CreatedDate:       v.CreatedDate,
				ModifiedDate:      v.ModifiedDate,
				Status:            v.Status,
			})
		}

	}

	return result, nil

}

func (s *Server) GetTransactionAttachment(ctx context.Context, req *pb.GetTransactionAttachmentRequest) (*pb.GetTransactionAttachmentResponse, error) {

	result := &pb.GetTransactionAttachmentResponse{
		Error:   false,
		Code:    200,
		Message: "Data",
	}

	if req.ReferenceNo != "" {

		apiReq := &ApiDownloadRequest{
			ReferenceNo: req.ReferenceNo,
		}

		res, err := s.ApiDownload(ctx, apiReq)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		if res.ResponseCode != "00" {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", res.ResponseMessage)
		}

		if len(res.ResponseData) > 0 {

			for _, v := range res.ResponseData {
				result.Data = append(result.Data, v.Url)
			}

		}

	}

	return result, nil

}

func (s *Server) GetTransaction(ctx context.Context, req *pb.GetTransactionRequest) (*pb.GetTransactionResponse, error) {

	result := &pb.GetTransactionResponse{
		Error:   false,
		Code:    200,
		Message: "List Data",
		Data:    []*pb.Transaction{},
	}

	result.Pagination = &pb.PaginationResponse{
		Limit:      req.Limit,
		Page:       req.Page,
		TotalRows:  0,
		TotalPages: 0,
	}

	currentUser, err := s.manager.GetMeFromJWT(ctx, "")
	if err != nil {
		return nil, err
	}

	filterData := []string{
		"company_id:" + currentUser.CompanyID,
		"is_mapped:true",
	}

	if req.Transaction.GetThirdPartyID() > 0 {
		filterData = append(filterData, "third_party_id:"+strconv.FormatUint(req.Transaction.ThirdPartyID, 10))
	}

	filter := &db.ListFilter{
		Filter: strings.Join(filterData, ","),
	}

	mappingORMs, err := s.provider.GetMapping(ctx, filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	beneficiaryIDs := []string{}

	if len(mappingORMs) > 0 {

		for _, v := range mappingORMs {

			if v.BeneficiaryID == 10101010 {

				res, err := s.GetBeneficiaryName(ctx, &pb.GetBeneficiaryNameRequest{ThirdPartyID: v.ThirdPartyID})
				if err != nil {
					if !errors.Is(err, gorm.ErrRecordNotFound) {
						return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
					}
				}

				if len(res.Data) > 0 {

					for _, d := range res.Data {
						if !contains(beneficiaryIDs, strconv.FormatUint(d.BeneficiaryId, 10)) {
							beneficiaryIDs = append(beneficiaryIDs, strconv.FormatUint(d.BeneficiaryId, 10))
						}
					}

				}

			} else {

				if !contains(beneficiaryIDs, strconv.FormatUint(v.BeneficiaryID, 10)) {
					beneficiaryIDs = append(beneficiaryIDs, strconv.FormatUint(v.BeneficiaryID, 10))
				}

			}

		}

	}

	httpReqParamsOpt := ApiListTransactionRequest{
		Page:  uint64(req.Page),
		Limit: uint64(req.Limit),
	}

	if len(beneficiaryIDs) > 0 {
		httpReqParamsOpt.BeneficiaryId = strings.Join(beneficiaryIDs, ",")
	} else {
		httpReqParamsOpt.BeneficiaryId = "10101010"
	}

	if req.Transaction != nil {
		if req.Transaction.StartDate != "" && req.Transaction.EndDate != "" {
			httpReqParamsOpt.StartDate = req.Transaction.StartDate
			httpReqParamsOpt.EndDate = req.Transaction.EndDate
		} else {
			return nil, status.Errorf(codes.InvalidArgument, "Start Date and End Date is Required")
		}

		if req.Transaction.BeneficiaryID > 0 {
			httpReqParamsOpt.BeneficiaryId = strconv.FormatUint(req.Transaction.BeneficiaryID, 10)
		}

		if req.Transaction.ThirdPartyID > 0 {
			httpReqParamsOpt.ThirdPartyId = req.Transaction.ThirdPartyID
		}

		if req.Transaction.ClaimPeriod > 0 {
			httpReqParamsOpt.ClaimPeriod = strconv.FormatUint(uint64(req.Transaction.ClaimPeriod), 10)
		}

		if req.Transaction.Status != "" {
			httpReqParamsOpt.Status = req.Transaction.Status
		}

		if req.Transaction.ReferenceNo != "" {
			httpReqParamsOpt.ReferenceNo = req.Transaction.ReferenceNo
		}

		if req.Transaction.ChannelID > 0 {
			httpReqParamsOpt.ChannelId = req.Transaction.ChannelID
		}

		if req.Transaction.ApplicantName != "" {
			httpReqParamsOpt.ApplicantName = req.Transaction.ApplicantName
		}
	}

	apiReq := &httpReqParamsOpt

	res, err := s.ApiListTransaction(ctx, apiReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	if res.ResponseCode != "00" {

		logrus.Error("Failed To Transfer Data : ", res.ResponseMessage)

	} else {

		if len(res.ResponseData) > 0 {

			for _, d := range res.ResponseData {

				transactionPB := &pb.Transaction{
					TransactionID:     d.TransactionId,
					ThirdPartyID:      d.ThirdPartyId,
					ThirdPartyName:    d.ThirdPartyName,
					ReferenceNo:       d.ReferenceNo,
					RegistrationNo:    d.RegistrationNo,
					ApplicantName:     d.ApplicantName,
					BeneficiaryID:     d.BeneficiaryId,
					BeneficiaryName:   d.BeneficiaryName,
					IssueDate:         d.IssueDate,
					EffectiveDate:     d.EffectiveDate,
					ExpiryDate:        d.ExpiryDate,
					ClaimPeriod:       d.ClaimPeriod,
					ClosingDate:       d.ClosingDate,
					Currency:          d.Currency,
					Amount:            d.Amount,
					CreatedDate:       d.CreatedDate,
					ModifiedDate:      d.ModifiedDate,
					Remark:            d.Remark,
					Status:            d.Status,
					ChannelID:         d.ChannelId,
					ChannelName:       d.ChannelName,
					TransactionTypeID: pb.BgType(d.TransactionTypeId),
				}

				result.Data = append(result.Data, transactionPB)

			}

		}

		result.Pagination = &pb.PaginationResponse{
			Limit:      int32(res.Pagination.Limit),
			Page:       int32(res.Pagination.Page),
			TotalRows:  int64(res.Pagination.TotalRecord),
			TotalPages: int32(res.Pagination.TotalPage),
		}

	}

	return result, nil

}

func (s *Server) GetTransactionDetail(ctx context.Context, req *pb.GetTransactionDetailRequest) (*pb.GetTransactionDetailResponse, error) {

	result := &pb.GetTransactionDetailResponse{
		Error:   false,
		Code:    200,
		Message: "Data",
	}

	if req.ReferenceNo != "" {

		httpReqParamsOpt := ApiListTransactionRequest{
			ReferenceNo: req.ReferenceNo,
			Page:        1,
			Limit:       1,
		}

		apiReq := &httpReqParamsOpt

		res, err := s.ApiListTransaction(ctx, apiReq)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		if res.ResponseCode != "00" {
			logrus.Error("Failed To Transfer Data : ", res.ResponseMessage)
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", res.ResponseMessage)
		}

		if len(res.ResponseData) > 0 {

			d := res.ResponseData[0]

			result.Data = &pb.Transaction{
				TransactionID:     d.TransactionId,
				ThirdPartyID:      d.ThirdPartyId,
				ThirdPartyName:    d.ThirdPartyName,
				ReferenceNo:       d.ReferenceNo,
				RegistrationNo:    d.RegistrationNo,
				ApplicantName:     d.ApplicantName,
				BeneficiaryID:     d.BeneficiaryId,
				BeneficiaryName:   d.BeneficiaryName,
				IssueDate:         d.IssueDate,
				EffectiveDate:     d.EffectiveDate,
				ExpiryDate:        d.ExpiryDate,
				ClaimPeriod:       d.ClaimPeriod,
				ClosingDate:       d.ClosingDate,
				Currency:          d.Currency,
				Amount:            d.Amount,
				CreatedDate:       d.CreatedDate,
				ModifiedDate:      d.ModifiedDate,
				Remark:            d.Remark,
				Status:            d.Status,
				ChannelID:         d.ChannelId,
				ChannelName:       d.ChannelName,
				TransactionTypeID: pb.BgType(d.TransactionTypeId),
			}

		}

	}

	return result, nil

}

func (s *Server) CreateTransaction(ctx context.Context, req *pb.CreateTransactionRequest) (*pb.CreateTransactionResponse, error) {

	result := &pb.CreateTransactionResponse{
		Error:   false,
		Code:    200,
		Message: "Data",
	}

	currentUser, err := s.manager.GetMeFromJWT(ctx, "")
	if err != nil {
		return nil, err
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		ctx = metadata.NewOutgoingContext(context.Background(), md)
	}
	var header, trailer metadata.MD

	taskConn, err := grpc.Dial(getEnv("TASK_SERVICE", ":9090"), opts...)
	if err != nil {
		logrus.Errorln("Failed connect to Task Service: %v", err)
		return nil, status.Errorf(codes.Internal, "Error Internal")
	}
	taskConn.Connect()
	defer taskConn.Close()

	taskClient := task_pb.NewTaskServiceClient(taskConn)

	switch req.Type {
	case "BG Mapping":

		taskData := req.MappingData
		taskDataBak := req.MappingDataBackup

		ids := []string{}

		if len(taskDataBak) > 0 {

			for _, v := range taskDataBak {

				needDelete := true

				if len(taskData) > 0 {

					for _, check := range taskData {

						if check.ThirdPartyID == v.ThirdPartyID {

							needDelete = false
							logrus.Println("Break at: " + strconv.FormatUint(check.ThirdPartyID, 10))
							break

						}

					}

				}

				if needDelete {

					filter := []string{
						"company_id:" + strconv.FormatUint(v.CompanyID, 10),
						"data.0.thirdPartyID:" + strconv.FormatUint(v.ThirdPartyID, 10),
					}

					taskMappingDigitalRes, err := taskClient.GetListTask(ctx, &task_pb.ListTaskRequest{Filter: strings.Join(filter, ","), Task: &task_pb.Task{Type: "BG Mapping Digital"}, Page: 1, Limit: 1}, grpc.Header(&header), grpc.Trailer(&trailer))
					if err != nil {
						if !errors.Is(err, gorm.ErrRecordNotFound) {
							return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
						}
					}

					if len(taskMappingDigitalRes.Data) > 0 {

						for _, taskMappingDigitalResData := range taskMappingDigitalRes.Data {

							taskMappingDigitalData := []*pb.MappingDigitalData{}
							err = json.Unmarshal([]byte(taskMappingDigitalResData.GetData()), &taskMappingDigitalData)
							if err != nil {
								return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
							}

							_, err := taskClient.SetTask(ctx, &task_pb.SetTaskRequest{TaskID: taskMappingDigitalResData.TaskID, Action: "delete", Comment: "delete"}, grpc.Header(&header), grpc.Trailer(&trailer))
							if err != nil {
								return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
							}

							if len(taskMappingDigitalData) > 0 {

								mappingFilter := []string{
									"company_id:" + strconv.FormatUint(taskMappingDigitalData[0].CompanyID, 10),
									"third_party_id:" + strconv.FormatUint(taskMappingDigitalData[0].ThirdPartyID, 10),
								}

								mappingListFilter := &db.ListFilter{
									Filter: strings.Join(mappingFilter, ","),
								}

								mappingORMs, err := s.provider.GetMapping(ctx, mappingListFilter)
								if err != nil {
									if !errors.Is(err, gorm.ErrRecordNotFound) {
										return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
									}
								}

								if len(mappingORMs) > 0 {

									for _, mappingORM := range mappingORMs {
										if mappingORM.Id > 0 {
											if !contains(ids, strconv.FormatUint(mappingORM.Id, 10)) {
												ids = append(ids, strconv.FormatUint(mappingORM.Id, 10))
											}
										}
									}

								}

							}

						}

					}

				}

				mappingORM, err := s.provider.GetMappingDetail(ctx, &pb.MappingORM{ThirdPartyID: v.ThirdPartyID, BeneficiaryID: 10101010, CompanyID: v.CompanyID})
				if err != nil {
					if !errors.Is(err, gorm.ErrRecordNotFound) {
						return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
					}
				}

				if mappingORM.Id > 0 {
					if !contains(ids, strconv.FormatUint(mappingORM.Id, 10)) {
						ids = append(ids, strconv.FormatUint(mappingORM.Id, 10))
					}
				}

			}

		}

		err = s.provider.DeleteMapping(ctx, ids)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		if len(taskData) > 0 {

			for _, v := range taskData {

				userID, err := strconv.ParseUint(currentUser.UserID, 10, 64)
				if err != nil {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				}

				data := &pb.MappingORM{
					CompanyID:     v.CompanyID,
					ThirdPartyID:  v.ThirdPartyID,
					BeneficiaryID: 10101010,
					IsMapped:      false,
					CreatedByID:   userID,
					UpdatedByID:   userID,
				}

				if v.IsAllowAllBeneficiary {
					data.IsMapped = true
				}

				mappingORM, err := s.provider.GetMappingDetail(ctx, &pb.MappingORM{ThirdPartyID: v.ThirdPartyID, BeneficiaryID: 10101010, CompanyID: v.CompanyID})
				if err != nil {
					if !errors.Is(err, gorm.ErrRecordNotFound) {
						return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
					}
				}

				if mappingORM.Id > 0 {
					data.Id = mappingORM.Id
				}

				mappingORM, err = s.provider.UpdateOrCreateMapping(ctx, data)
				if err != nil {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				}

				mappingPB, err := mappingORM.ToPB(ctx)
				if err != nil {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				}

				result.Data = append(result.Data, &mappingPB)

			}

		}

	case "BG Mapping Digital":

		taskData := req.MappingDigitalData
		taskDataBak := req.MappingDigitalDataBackup

		ids := []string{}

		if len(taskDataBak) > 0 {

			for _, v := range taskDataBak {

				mappingFilter := []string{
					"company_id:" + strconv.FormatUint(v.CompanyID, 10),
					"third_party_id:" + strconv.FormatUint(v.ThirdPartyID, 10),
				}

				mappingListFilter := &db.ListFilter{
					Filter: strings.Join(mappingFilter, ","),
				}

				mappingORMs, err := s.provider.GetMapping(ctx, mappingListFilter)
				if err != nil {
					if !errors.Is(err, gorm.ErrRecordNotFound) {
						return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
					}
				}

				if len(mappingORMs) > 0 {

					for _, mappingORM := range mappingORMs {
						if mappingORM.Id > 0 {
							if !contains(ids, strconv.FormatUint(mappingORM.Id, 10)) {
								ids = append(ids, strconv.FormatUint(mappingORM.Id, 10))
							}
						}
					}

				}

			}

		}

		if len(taskData) > 0 {

			for _, v := range taskData {

				mappingFilter := []string{
					"company_id:" + strconv.FormatUint(v.CompanyID, 10),
					"third_party_id:" + strconv.FormatUint(v.ThirdPartyID, 10),
					"beneficiary_id:10101010",
				}

				mappingListFilter := &db.ListFilter{
					Filter: strings.Join(mappingFilter, ","),
				}

				mappingORMs, err := s.provider.GetMapping(ctx, mappingListFilter)
				if err != nil {
					if !errors.Is(err, gorm.ErrRecordNotFound) {
						return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
					}
				}

				if len(mappingORMs) > 0 {

					for _, mappingORM := range mappingORMs {
						if mappingORM.Id > 0 {
							if !contains(ids, strconv.FormatUint(mappingORM.Id, 10)) {
								ids = append(ids, strconv.FormatUint(mappingORM.Id, 10))
							}
						}
					}

				}

			}

		}

		err = s.provider.DeleteMapping(ctx, ids)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		if len(taskData) > 0 {

			for _, v := range taskData {

				userID, err := strconv.ParseUint(currentUser.UserID, 10, 64)
				if err != nil {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				}

				data := &pb.MappingORM{
					CompanyID:     v.CompanyID,
					ThirdPartyID:  v.ThirdPartyID,
					BeneficiaryID: v.BeneficiaryId,
					IsMapped:      true,
					CreatedByID:   userID,
					UpdatedByID:   userID,
				}

				mappingORM, err := s.provider.GetMappingDetail(ctx, &pb.MappingORM{ThirdPartyID: v.ThirdPartyID, BeneficiaryID: v.BeneficiaryId, CompanyID: v.CompanyID})
				if err != nil {
					if !errors.Is(err, gorm.ErrRecordNotFound) {
						return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
					}
				}

				if mappingORM.Id > 0 {
					data.Id = mappingORM.Id
				}

				mappingORM, err = s.provider.UpdateOrCreateMapping(ctx, data)
				if err != nil {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				}

				mappingPB, err := mappingORM.ToPB(ctx)
				if err != nil {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				}

				result.Data = append(result.Data, &mappingPB)

			}

		}

	}

	return result, nil

}

func (s *Server) DeleteTransaction(ctx context.Context, req *pb.DeleteTransactionRequest) (*pb.DeleteTransactionResponse, error) {

	result := &pb.DeleteTransactionResponse{
		Error:   false,
		Code:    200,
		Message: "Data",
	}

	var newCtx context.Context

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		newCtx = metadata.NewOutgoingContext(context.Background(), md)
	}

	currentUser, userMD, err := s.manager.GetMeFromMD(ctx)
	if err != nil {
		return nil, err
	}
	if currentUser == nil {
		return nil, s.UnauthorizedError()
	}
	var trailer metadata.MD

	taskClient := s.svcConn.TaskServiceClient()

	switch req.Type {
	case "BG Mapping":

		taskData := req.MappingData
		// taskDataBak := req.MappingDataBackup

		ids := []string{}

		for _, v := range taskData {

			filter := []string{
				"company_id:" + strconv.FormatUint(v.CompanyID, 10),
				"data.0.thirdPartyID:" + strconv.FormatUint(v.ThirdPartyID, 10),
			}

			taskMappingDigitalRes, err := taskClient.GetListTask(newCtx, &task_pb.ListTaskRequest{Filter: strings.Join(filter, ","), Task: &task_pb.Task{Type: "BG Mapping Digital"}, Page: 1, Limit: 1}, grpc.Header(&userMD), grpc.Trailer(&trailer))
			if err != nil {
				logrus.Println("[api][DeleteTransaction] Failed when execute GetListTask:", err)
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, err
				}
			}

			if len(taskMappingDigitalRes.Data) > 0 {

				for _, taskMappingDigitalResData := range taskMappingDigitalRes.Data {

					taskMappingDigitalData := []*pb.MappingDigitalData{}
					err = json.Unmarshal([]byte(taskMappingDigitalResData.GetData()), &taskMappingDigitalData)
					if err != nil {
						return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
					}

					_, err := taskClient.SetTask(newCtx, &task_pb.SetTaskRequest{TaskID: taskMappingDigitalResData.TaskID, Action: "delete", Comment: "delete"}, grpc.Header(&userMD), grpc.Trailer(&trailer))
					if err != nil {
						logrus.Println("[api][DeleteTransaction] Failed when execute SetTask:", err)
						return nil, err
					}

					if len(taskMappingDigitalData) > 0 {

						mappingFilter := []string{
							"company_id:" + strconv.FormatUint(taskMappingDigitalData[0].CompanyID, 10),
							"third_party_id:" + strconv.FormatUint(taskMappingDigitalData[0].ThirdPartyID, 10),
						}

						mappingListFilter := &db.ListFilter{
							Filter: strings.Join(mappingFilter, ","),
						}

						mappingORMs, err := s.provider.GetMapping(ctx, mappingListFilter)
						if err != nil {
							if !errors.Is(err, gorm.ErrRecordNotFound) {
								return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
							}
						}

						if len(mappingORMs) > 0 {

							for _, mappingORM := range mappingORMs {
								if mappingORM.Id > 0 {
									if !contains(ids, strconv.FormatUint(mappingORM.Id, 10)) {
										ids = append(ids, strconv.FormatUint(mappingORM.Id, 10))
									}
								}
							}

						}

					}

				}

			}

			mappingORM, err := s.provider.GetMappingDetail(ctx, &pb.MappingORM{ThirdPartyID: v.ThirdPartyID, BeneficiaryID: 10101010, CompanyID: v.CompanyID})
			if err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				}
			}

			if mappingORM.Id > 0 {
				if !contains(ids, strconv.FormatUint(mappingORM.Id, 10)) {
					ids = append(ids, strconv.FormatUint(mappingORM.Id, 10))
				}
			}

		}

		err = s.provider.DeleteMapping(ctx, ids)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

	case "BG Mapping Digital":

		taskData := req.MappingDigitalData
		// taskDataBak := req.MappingDigitalDataBackup

		ids := []string{}

		if len(taskData) > 0 {

			for _, v := range taskData {

				mappingFilter := []string{
					"company_id:" + strconv.FormatUint(v.CompanyID, 10),
					"third_party_id:" + strconv.FormatUint(v.ThirdPartyID, 10),
				}

				mappingListFilter := &db.ListFilter{
					Filter: strings.Join(mappingFilter, ","),
				}

				mappingORMs, err := s.provider.GetMapping(ctx, mappingListFilter)
				if err != nil {
					if !errors.Is(err, gorm.ErrRecordNotFound) {
						return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
					}
				}

				if len(mappingORMs) > 0 {

					for _, mappingORM := range mappingORMs {
						if mappingORM.Id > 0 {
							if !contains(ids, strconv.FormatUint(mappingORM.Id, 10)) {
								ids = append(ids, strconv.FormatUint(mappingORM.Id, 10))
							}
						}
					}

				}

			}

		}

		err = s.provider.DeleteMapping(ctx, ids)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		if len(taskData) > 0 {

			for _, v := range taskData {

				filter := []string{
					"company_id:" + strconv.FormatUint(v.CompanyID, 10),
				}

				taskMappingRes, err := taskClient.GetListTask(newCtx, &task_pb.ListTaskRequest{Filter: strings.Join(filter, ","), Task: &task_pb.Task{Type: "BG Mapping"}, Page: 1, Limit: 1}, grpc.Header(&userMD), grpc.Trailer(&trailer))
				if err != nil {
					logrus.Println("[api][DeleteTransaction] Failed when execute GetListTask:", err)
					if !errors.Is(err, gorm.ErrRecordNotFound) {
						return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
					}
				}

				if len(taskMappingRes.Data) > 0 {

					for _, d := range taskMappingRes.Data {

						taskMappingData := []*pb.MappingData{}
						err = json.Unmarshal([]byte(d.GetData()), &taskMappingData)
						if err != nil {
							return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
						}

						if len(taskMappingData) > 0 {

							for _, dd := range taskMappingData {

								if v.ThirdPartyID == dd.ThirdPartyID {

									userID := currentUser.UserID

									data := &pb.MappingORM{
										CompanyID:     v.CompanyID,
										ThirdPartyID:  v.ThirdPartyID,
										BeneficiaryID: 10101010,
										IsMapped:      dd.IsAllowAllBeneficiary,
										CreatedByID:   userID,
										UpdatedByID:   userID,
									}

									mappingORM, err := s.provider.GetMappingDetail(ctx, &pb.MappingORM{ThirdPartyID: v.ThirdPartyID, BeneficiaryID: 10101010, CompanyID: v.CompanyID})
									if err != nil {
										if !errors.Is(err, gorm.ErrRecordNotFound) {
											return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
										}
									}

									if mappingORM.Id > 0 {
										data.Id = mappingORM.Id
									}

									_, err = s.provider.UpdateOrCreateMapping(ctx, data)
									if err != nil {
										return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
									}

								}

							}

						}

					}

				}

			}

		}

	}

	return result, nil

}

func (s *Server) CreateIssuing(ctx context.Context, req *pb.CreateIssuingRequest) (*pb.CreateIssuingResponse, error) {

	result := &pb.CreateIssuingResponse{
		Error:   false,
		Code:    200,
		Message: "Data",
	}

	var newCtx context.Context

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		newCtx = metadata.NewOutgoingContext(context.Background(), md)
	}

	currentUser, userMD, err := s.manager.GetMeFromMD(ctx)
	if err != nil {
		return nil, err
	}
	if currentUser == nil {
		return nil, s.UnauthorizedError()
	}
	var trailer metadata.MD

	systemClient := s.svcConn.SystemServiceClient()

	isIndividu := uint64(req.Data.Applicant.GetApplicantType().Number())
	dateEstablished := ""

	if isIndividu == 0 {
		dateEstablished = req.Data.Applicant.GetDateEstablished()
		if dateEstablished == "" {
			return nil, status.Errorf(codes.InvalidArgument, "Bad Request: %v", "Empty value on dateEstablished when isIndividu is true")
		}
	}

	var gender string

	if req.Data.Applicant.GetGender().Number() == 0 {
		gender = "Laki-laki"
	} else {
		gender = "Perempuan"
	}

	contractGuaranteeType := req.Data.Project.GetContractGuaranteeType()

	var counterGuaranteeTypeString map[string]string
	insuranceLimitId := ""
	sp3No := ""
	nonCashAccountNo := ""
	nonCashAccountAmount := 0.0
	cashAccountNo := ""
	cashAccountAmount := 0.0
	customerLimitId := ""
	customerLimitAmount := 0.0
	isEndOfYearBg := "0"

	openingBranchORMs, err := systemClient.ListMdBranch(newCtx, &system_pb.ListMdBranchRequest{
		Data: &system_pb.MdBranch{
			Id: req.Data.Publishing.GetOpeningBranchId(),
		},
	}, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "Opening Branch not found")
		} else {
			return nil, err
		}
	}

	if len(openingBranchORMs.Data) == 0 {
		return nil, status.Errorf(codes.NotFound, "Opening Branch not found")
	}

	openingBranch := openingBranchORMs.Data[0]

	publishingBranchORMs, err := systemClient.ListMdBranch(newCtx, &system_pb.ListMdBranchRequest{
		Data: &system_pb.MdBranch{
			Id: req.Data.Publishing.GetPublishingBranchId(),
		},
	}, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "Publishing Branch not found")
		} else {
			return nil, err
		}
	}

	if len(publishingBranchORMs.Data) == 0 {
		return nil, status.Errorf(codes.NotFound, "Publishing Branch not found")
	}

	publishingBranch := publishingBranchORMs.Data[0]

	if req.Data.Publishing.BgType == pb.BgType_GovernmentPaymentGuarantee {
		isEndOfYearBg = "1"
		if req.Data.Project.GetNrkNumber() == "" {
			return nil, status.Errorf(codes.InvalidArgument, "Bad Request: %v", "Empty value on required NRK Number field when Government Payment Guarantee is selected")
		}
	}

	switch contractGuaranteeType {
	case pb.ContractGuaranteeType_Insurance: // Insurance

		counterGuaranteeTypeString = map[string]string{"0": "insurance limit"}

		insuranceLimitId = req.Data.Project.GetInsuranceLimitId()
		sp3No = req.Data.Project.GetSp3No()

		if insuranceLimitId == "" || sp3No == "" {
			return nil, status.Errorf(codes.InvalidArgument, "Bad Request: %v", "Empty value on required field(s) when insurance limit is selected")
		}

	case pb.ContractGuaranteeType_Cash: // Tunai / Cash

		counterGuaranteeTypeString = map[string]string{"0": "hold account"}

		cashAccountNo = req.Data.Project.GetCashAccountNo()
		cashAccountAmount = req.Data.Project.GetCashAccountAmount()

		if cashAccountNo == "" || cashAccountAmount <= 0.0 {
			return nil, status.Errorf(codes.InvalidArgument, "Bad Request: %v", "Empty value on required field(s) when hold account is selected")
		}

	case pb.ContractGuaranteeType_NonCashLoan: // Non Cash Loan

		counterGuaranteeTypeString = map[string]string{"0": "customer limit"}

		nonCashAccountNo = req.Data.Project.GetNonCashAccountNo()
		nonCashAccountAmount = req.Data.Project.GetNonCashAccountAmount()

		if req.Data.Account.Cif == "" {
			return nil, status.Errorf(codes.InvalidArgument, "You are not allowed for Non Cash Loan facility")
		}

		inquiryLimit, err := s.ApiInquiryLimitIndividual(ctx, &ApiInquiryLimitIndividualRequest{Cif: req.Data.Account.Cif})
		if err != nil {
			logrus.Println("Error Limit Individual: ", err.Error())
			return nil, status.Errorf(codes.InvalidArgument, "You are not allowed for Non Cash Loan facility")
		}

		customerLimitId = strconv.FormatUint(inquiryLimit.ResponseData[0].CustomerLimitId, 10)
		customerLimitAmount = float64(inquiryLimit.ResponseData[0].AvailableAmount)

		if nonCashAccountNo == "" || nonCashAccountAmount <= 0.0 {
			return nil, status.Errorf(codes.InvalidArgument, "Bad Request: %v", "Empty value on required field(s) when customer limit is selected")
		}

	case pb.ContractGuaranteeType_Combination: // Kombinasi

		counterGuaranteeTypeString = map[string]string{"0": "customer limit", "1": "hold account"}

		cashAccountNo = req.Data.Project.GetCashAccountNo()
		cashAccountAmount = req.Data.Project.GetCashAccountAmount()
		nonCashAccountNo = req.Data.Project.GetNonCashAccountNo()
		nonCashAccountAmount = req.Data.Project.GetNonCashAccountAmount()

		if req.Data.Account.Cif == "" {
			return nil, status.Errorf(codes.InvalidArgument, "You are not allowed for Non Cash Loan facility")
		}

		inquiryLimit, err := s.ApiInquiryLimitIndividual(ctx, &ApiInquiryLimitIndividualRequest{Cif: req.Data.Account.Cif})
		if err != nil {
			logrus.Println("Error Limit Individual: ", err.Error())
			return nil, status.Errorf(codes.InvalidArgument, "You are not allowed for Combination facility")
		}

		customerLimitId = strconv.FormatUint(inquiryLimit.ResponseData[0].CustomerLimitId, 10)
		customerLimitAmount = float64(inquiryLimit.ResponseData[0].AvailableAmount)

		if nonCashAccountNo == "" || nonCashAccountAmount <= 0.0 || cashAccountNo == "" || cashAccountAmount <= 0.0 {
			return nil, status.Errorf(codes.InvalidArgument, "Bad Request: %v", "Empty value on required field(s) when combination account is selected")
		}

	default:

		return nil, status.Errorf(codes.InvalidArgument, "Bad Request: %v", "Invalid Contract Guarantee Type")

	}

	openingBranchID, err := strconv.ParseInt(openingBranch.Id, 10, 64)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	publishingBranchID, err := strconv.ParseInt(publishingBranch.Id, 10, 64)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	openingBranchPadded := fmt.Sprintf("%05d", openingBranchID)
	publishingBranchPadded := fmt.Sprintf("%05d", publishingBranchID)

	httpReqData := ApiBgIssuingRequest{
		AccountNo:              req.Data.Account.GetAccountNumber(),
		ApplicantName:          req.Data.Applicant.GetName(),
		ApplicantAddress:       req.Data.Applicant.GetAddress(),
		IsIndividu:             isIndividu,
		NIK:                    req.Data.Applicant.GetNik(),
		BirthDate:              req.Data.Applicant.GetBirthDate(),
		Gender:                 gender,
		NPWPNo:                 req.Data.Applicant.GetNpwpNo(),
		DateEstablished:        dateEstablished,
		CompanyType:            uint64(req.Data.Applicant.GetCompanyType().Number()),
		IsPlafond:              0,
		TransactionType:        uint64(req.Data.Publishing.GetBgType().Number()),
		IsEndOfYearBg:          isEndOfYearBg,
		NRK:                    req.Data.Project.GetNrkNumber(),
		ProjectName:            req.Data.Project.GetName(),
		ThirdPartyId:           req.Data.Publishing.GetThirdPartyID(),
		BeneficiaryName:        req.Data.Applicant.GetBeneficiaryName(),
		ProjectAmount:          req.Data.Project.GetProjectAmount(),
		ContractNo:             req.Data.Project.GetContractNumber(),
		ContractDate:           req.Data.Project.GetProjectDate(),
		Currency:               req.Data.Project.GetBgCurrency(),
		Amount:                 req.Data.Project.GetBgAmount(),
		EffectiveDate:          req.Data.Publishing.GetEffectiveDate(),
		MaturityDate:           req.Data.Publishing.GetExpiryDate(),
		ClaimPeriod:            req.Data.Publishing.GetClaimPeriod(),
		IssuingBranch:          openingBranchPadded,
		PublishingBranch:       publishingBranchPadded,
		ContraGuarantee:        counterGuaranteeTypeString,
		InsuranceLimitId:       insuranceLimitId,
		SP3No:                  sp3No,
		HoldAccountNo:          cashAccountNo,
		HoldAccountAmount:      cashAccountAmount,
		ConsumerLimitId:        customerLimitId,
		ConsumerLimitAmount:    customerLimitAmount,
		ApplicantContactPerson: req.Data.Applicant.GetContactPerson(),
		ApplicantPhoneNumber:   req.Data.Applicant.GetPhoneNumber(),
		ApplicantEmail:         req.Data.Applicant.GetEmail(),
		ChannelId:              getEnv("BG_CHANNEL_ID", "2"),
		ApplicantCustomerId:    "0",
		BeneficiaryCustomerId:  "0",
		LegalDocument:          req.Data.Document.GetFileBusinessLegal(),
		ContractDocument:       req.Data.Document.GetFileTender(),
		Sp3Document:            req.Data.Document.GetFileSp(),
		OthersDocument:         req.Data.Document.GetFileOther(),
	}

	createIssuingRes, err := s.ApiCreateIssuing(ctx, &httpReqData)
	if err != nil {
		logrus.Println("Failed to create issuing: ", err)
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	httpReqParamsOpt := ApiBgTrackingRequest{
		RegistrationNo: createIssuingRes.Data.RegistrationNo,
	}

	apiReq := &httpReqParamsOpt

	checkIssuingRes, err := s.ApiCheckIssuingStatus(ctx, apiReq)
	if err != nil {
		logrus.Println("Failed to check issuing: ", err)
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	result.Data = &pb.IssuingPortal{
		RegistrationNo:  checkIssuingRes.Data.RegistrationNo,
		ReferenceNo:     checkIssuingRes.Data.ReferenceNo,
		WarkatUrl:       checkIssuingRes.Data.WarkatUrl,
		WarkatUrlPublic: checkIssuingRes.Data.WarkatUrlPublic,
		Status:          checkIssuingRes.Data.Status,
		ModifiedDate:    checkIssuingRes.Data.ModifiedDate,
	}

	return result, nil

}

func (s *Server) CheckIssuingStatus(ctx context.Context, req *pb.CheckIssuingRequest) (*pb.CheckIssuingResponse, error) {

	result := &pb.CheckIssuingResponse{
		Error:   false,
		Code:    200,
		Message: "Data",
	}

	var newCtx context.Context

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		newCtx = metadata.NewOutgoingContext(context.Background(), md)
	}

	currentUser, userMD, err := s.manager.GetMeFromMD(ctx)
	if err != nil {
		return nil, err
	}
	if currentUser == nil {
		return nil, s.UnauthorizedError()
	}
	var trailer metadata.MD

	taskClient := s.svcConn.TaskServiceClient()

	taskRes, err := taskClient.GetTaskByID(newCtx, &task_pb.GetTaskByIDReq{ID: req.TaskID, Type: "BG Issuing"}, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	var taskData *pb.IssuingData
	json.Unmarshal([]byte(taskRes.Data.GetData()), &taskData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	logrus.Print(taskData)

	httpReqParamsOpt := ApiBgTrackingRequest{
		RegistrationNo: taskData.RegistrationNo,
	}

	apiReq := &httpReqParamsOpt

	res, err := s.ApiCheckIssuingStatus(ctx, apiReq)
	if err != nil {
		return nil, err
	}

	taskData.ReferenceNo = res.Data.ReferenceNo

	if contains([]string{"Rejected", "Cancelled"}, res.Data.ReferenceNo) {
		taskData.ReferenceNo = "-"
	}

	data, err := json.Marshal(taskData)
	if err != nil {
		logrus.Error("Failed To Marshal : ", taskData)
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	taskReq := &task_pb.UpdateTaskDataReq{
		Type:   "BG Issuing",
		TaskID: req.TaskID,
		Data:   string(data),
	}

	_, err = taskClient.UpdateTaskData(newCtx, taskReq, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		logrus.Error("Failed To Transfer Data : ", "FAK")
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	result.Data = &pb.IssuingPortal{
		RegistrationNo:  res.Data.RegistrationNo,
		ReferenceNo:     res.Data.ReferenceNo,
		WarkatUrl:       res.Data.WarkatUrl,
		WarkatUrlPublic: res.Data.WarkatUrlPublic,
		Status:          res.Data.Status,
		ModifiedDate:    res.Data.ModifiedDate,
	}

	return result, nil

}

func (s *Server) FileUpload(ctx context.Context, req *pb.FileUploadRequest) (*pb.FileUploadResponse, error) {

	result := &pb.FileUploadResponse{
		Error:   false,
		Code:    200,
		Message: "Success",
	}

	_, err := base64.StdEncoding.DecodeString(req.GetData())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Bad Request: File is corrupted")
	}

	// contentType := http.DetectContentType(decodedBytes)

	// if contentType != "application/pdf" {
	// 	return nil, status.Errorf(codes.InvalidArgument, "Bad Request: Invalid filetype")
	// }

	httpReqParamsOpt := ApiUploadEncodeRequest{
		Document: req.GetData(),
	}

	apiReq := &httpReqParamsOpt

	res, err := s.ApiUploadEncode(ctx, apiReq)
	if err != nil {
		return nil, err
	}

	resultData := &pb.FileUploadData{
		FileName:        res.ResponseData.Filename,
		DocumentPath:    res.ResponseData.DocumentPath,
		UploadDate:      res.ResponseData.UploadDate,
		UploadedFileUrl: res.ResponseData.UploadFileUrl,
	}

	result.Data = resultData

	return result, nil

}

func (s *Server) CheckIndividualLimit(ctx context.Context, req *pb.CheckIndividualLimitRequest) (*pb.CheckIndividualLimitResponse, error) {

	result := &pb.CheckIndividualLimitResponse{
		Error:    false,
		Code:     200,
		Message:  "Success",
		HasLimit: false,
	}

	inquiryLimit, err := s.ApiInquiryLimitIndividual(ctx, &ApiInquiryLimitIndividualRequest{Cif: req.Cif})
	if err != nil {
		logrus.Println("Error Limit Individual: ", err.Error())
	}

	if inquiryLimit.ResponseCode == "00" {
		result.HasLimit = true
	}

	return result, nil

}

func (s *Server) NotificationRequestBuilder(ctx context.Context, nextStep string, task *task_pb.Task, action string, username string, emails []string) (*notification_pb.SendNotificationWorkflowRequest, error) {

	var newCtx context.Context

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		newCtx = metadata.NewOutgoingContext(context.Background(), md)
	}

	taskClient := s.svcConn.TaskServiceClient()
	companyClient := s.svcConn.CompanyServiceClient()

	taskRes, err := taskClient.GetTaskByID(newCtx, &task_pb.GetTaskByIDReq{Type: task.GetType(), ID: task.GetTaskID()})
	if err != nil {
		logrus.Errorln("[api][func: NotificationRequestBuilder] Unable to Get Task by ID:", err.Error())
		return nil, err
	}

	if !taskRes.GetFound() {
		logrus.Errorln("[api][func: NotificationRequestBuilder] Task not found")
		return nil, status.Errorf(codes.NotFound, "Task not found")
	}

	task = taskRes.GetData()

	loc, _ := time.LoadLocation("Asia/Jakarta")

	var workflowDoc *workflow_pb.ValidateWorkflowData
	err = json.Unmarshal([]byte(taskRes.GetData().GetWorkflowDoc()), &workflowDoc)
	if err != nil {
		logrus.Errorln("[api][func: NotificationRequestBuilder] Unable to Unmarshal Workflow Data:", err)
		return nil, err
	}

	userID := []uint64{}

	eventName := ""
	switch action {
	case "send approval":
		eventName = "Created new transaction and sent for approval"
	case "send other approval":
		eventName = "Created new transaction and sent for approval"
	case "complete":
		eventName = "Transaction request gets final approval and sent for processing"
		userID = []uint64{task.GetCreatedByID()}
	case "approve":
		eventName = "Transaction request gets approval"
		userID = []uint64{task.GetCreatedByID()}
	case "error":
		eventName = "Transaction error"
		userID = []uint64{task.GetCreatedByID()}
	case "timeout":
		eventName = "Transaction suspended/timeout"
		userID = []uint64{task.GetCreatedByID()}
	case "success":
		eventName = "Transaction success"
		userID = []uint64{task.GetCreatedByID()}
	case "waiting":
		eventName = "Transaction waiting"
		userID = []uint64{task.GetCreatedByID()}
	case "reject":
		eventName = "Transaction request gets rejected"
		userID = []uint64{task.GetCreatedByID()}
	case "rework":
		eventName = "Transaction request sent for rework"
		userID = []uint64{task.GetCreatedByID()}
	default:
		return nil, nil
	}

	company, err := companyClient.DetailCompany(newCtx, &company_pb.CompanyParams{
		CompanyID: taskRes.GetData().GetCompanyID(),
	})
	if err != nil {
		logrus.Errorln("[api][func: NotificationRequestBuilder] Unable to Detail Company:", err)
		return nil, err
	}

	// Set status info
	statusInfo := ""
	if nextStep != "" {
		statusInfo = fmt.Sprintf(" on %v", strings.Title(nextStep))
	}

	notificationData := &pb.NotificationData{
		USERNAME_MAKER:    task.CreatedByName,
		USERNAME_APPROVER: task.LastApprovedByName,
		CREATED_DATETIME:  task.CreatedAt.AsTime().In(loc).Format("2006-01-02 15:04:05"),
		CREATED_DATE:      task.CreatedAt.AsTime().In(loc).Format("2006-01-02"),
		CREATED_TIME:      task.CreatedAt.AsTime().In(loc).Format("15:04:05"),
		EVENT_DATETIME:    task.UpdatedAt.AsTime().In(loc).Format("2006-01-02 15:04:05"),
		EVENT_DATE:        task.UpdatedAt.AsTime().In(loc).Format("2006-01-02"),
		EVENT_TIME:        task.UpdatedAt.AsTime().In(loc).Format("15:04:05"),
		TASK_ID:           strconv.FormatUint(task.GetTaskID(), 10),
		USERNAME_REJECTOR: task.LastRejectedByName,
		COMPANY_NAME:      company.GetCompanyName(),
		USERNAME_CHECKER:  username,
		USERNAME_RELEASER: username,
		MODULE:            taskRes.GetData().GetType(),
		STATUS_ACTION:     action,
		STATUS_INFO:       statusInfo,
		STATUS_SEND:       "Needs Approval",
		REASON:            task.Reasons,
		COMMENT:           task.Comment,
	}

	notificationDataByte, err := json.Marshal(notificationData)
	if err != nil {
		logrus.Errorln("[api][func: NotificationRequestBuilder] Unable to Marshal Notification Data:", err)
		return nil, err
	}

	requestData := &notification_pb.SendNotificationWorkflowRequest{
		ModuleID:    0,
		EventID:     0,
		ModuleName:  taskRes.GetData().GetType(),
		EventName:   eventName,
		Data:        string(notificationDataByte),
		RoleIDs:     workflowDoc.GetWorkflow().GetCurrentRoleIDs(),
		Step:        workflowDoc.GetWorkflow().GetCurrentStep(),
		CompanyID:   task.GetCompanyID(),
		UserID:      userID,
		CustomEmail: emails,
	}

	requestDataByte, err := json.Marshal(requestData)
	if err != nil {
		logrus.Errorln("[api][func: NotificationRequestBuilder] Unable to Marshal Send Notification Workflow Request Data:", err)
		return nil, err
	}

	logrus.Println("[api][func: NotificationRequestBuilder] Send Notification Workflow Request Data:", string(requestDataByte))

	return requestData, nil

}

func (s *Server) FilterBuilder(ctx context.Context, currentUser manager.UserData) (status []string, filter []string, err error) {

	// - Maker: 1. Draft, 2. Returned, 3. Pending, 4. Request for Delete, 5. Approved, 6. Rejected
	// - Signer: 1. Pending, 2. Request for Delete, 3. Approved, 4. Rejected

	if contains(currentUser.Authorities, "maker") {

		status = []string{"2", "3", "1", "6", "4", "5"}
		filter = []string{"status:<>0", "status:<>7"}

	} else if contains(currentUser.Authorities, "checker") || contains(currentUser.Authorities, "signer") || contains(currentUser.Authorities, "releaser") {

		if contains(currentUser.Authorities, "maker") {

			status = []string{"2", "3", "1", "6", "4", "5"}
			filter = []string{"status:<>0", "status:<>7"}

		} else {

			status = []string{"1", "6", "4", "5"}
			filter = []string{"status:<>0", "status:<>2", "status:<>3", "status:<>7"}

		}

	} else {

		return nil, nil, errors.New("permission denied")

	}

	return status, filter, nil

}
