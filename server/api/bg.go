package api

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bitbucket.bri.co.id/scm/addons/addons-bg-service/server/db"
	account_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/account"
	filelistener_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/filelistener"
	system_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/system"
	task_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/task"
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

	for _, v := range data {
		currency, err := v.ToPB(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		result.Data = append(result.Data, &currency)
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

	me, err := s.manager.GetMeFromJWT(ctx, "")
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

	taskFilter := &task_pb.Task{
		Type: "BG Issuing",
	}

	filter := []string{
		"data.publishing.thirdPartyID:" + strconv.FormatUint(req.ThirdPartyID, 10),
	}

	dataReq := &task_pb.ListTaskRequest{
		Task:   taskFilter,
		Filter: strings.Join(filter, ","),
		In:     me.CompanyIDs,
	}

	dataList, err := taskClient.GetListTask(ctx, dataReq, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	names := []string{}

	for _, v := range dataList.Data {

		taskData := pb.IssuingData{}
		json.Unmarshal([]byte(v.Data), &taskData)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		if !contains(names, taskData.Applicant.Name) {
			names = append(names, taskData.Applicant.Name)
		}

	}

	for _, v := range names {

		result.Data = append(result.Data, &pb.ApplicantName{
			Name: v,
		})

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

	me, err := s.manager.GetMeFromJWT(ctx, "")
	if err != nil {
		return nil, err
	}

	apiReq := &ApiInquiryBenficiaryRequest{
		ThirdPartyID: req.ThirdPartyID,
	}

	res, err := ApiInquiryBeneficiary(ctx, apiReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	data := []*pb.BeneficiaryName{}

	if res.ResponseCode == "00" {
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

	if req.Type == 0 {

		result.Data = data

	} else {

		mappedBeneficiaryIDs := []string{}

		mappingFilter := []string{
			"company_id:" + strconv.FormatUint(me.CompanyID, 10),
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

		for _, v := range mappingORMs {
			if v.BeneficiaryID != 10101010 {
				if !contains(mappedBeneficiaryIDs, strconv.FormatUint(v.BeneficiaryID, 10)) {
					mappedBeneficiaryIDs = append(mappedBeneficiaryIDs, strconv.FormatUint(v.BeneficiaryID, 10))
				}
			}
		}

		for _, v := range data {
			if contains(mappedBeneficiaryIDs, strconv.FormatUint(v.BeneficiaryId, 10)) {
				result.Data = append(result.Data, v)
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

	me, err := s.manager.GetMeFromJWT(ctx, "")
	if err != nil {
		return nil, err
	}

	if me.UserType == "ba" {

		apiReq := &ApiInquiryThirdPartyByStatusRequest{
			Status: "Active",
		}

		res, err := ApiInquiryThirdPartyByStatus(ctx, apiReq)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		if res.ResponseCode == "00" {
			for _, v := range res.ResponseData {
				result.Data = append(result.Data, &pb.ThirdParty{
					Id:   v.ThirdPartyID,
					Name: v.FullName,
				})
			}
		}

	} else {

		filter := &db.ListFilter{}

		filterMapped := []string{
			"company_id:" + strconv.FormatUint(me.CompanyID, 10),
		}
		if req.Type == *pb.ThirdPartyType_NeedMapping.Enum() {
			filterMapped = append(filterMapped, "is_mapped:false")
		} else if req.Type == *pb.ThirdPartyType_IsMapped.Enum() {
			filterMapped = append(filterMapped, "is_mapped:true")
		}

		filter.Filter = strings.Join(filterMapped, ",")

		thirdPartyNameList, err := s.provider.GetMapping(ctx, filter)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		logrus.Print(thirdPartyNameList)

		ids := []string{}

		for _, v := range thirdPartyNameList {
			if !contains(ids, strconv.FormatUint(v.ThirdPartyID, 10)) {
				ids = append(ids, strconv.FormatUint(v.ThirdPartyID, 10))
			}
		}

		for _, v := range ids {

			id, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
			}

			name := ""

			apiReq := &ApiInquiryThirdPartyByIDRequest{
				ThirdPartyID: id,
			}

			res, err := ApiInquiryThirdPartyByID(ctx, apiReq)
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

	res, err := ApiInquiryLimitIndividual(ctx, apiReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	if res.ResponseCode != "00" {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", string(*res.ResponseMessage))
	}

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

		res, err := ApiDownload(ctx, apiReq)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		if res.ResponseCode != "00" {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", res.ResponseMessage)
		}

		for _, v := range res.ResponseData {
			result.Data = append(result.Data, v.Url)
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

	me, err := s.manager.GetMeFromJWT(ctx, "")
	if err != nil {
		return nil, err
	}

	filterData := []string{
		"company_id:" + strconv.FormatUint(me.CompanyID, 10),
		"is_mapped:true",
	}

	if req.Transaction.GetThirdPartyID() > 0 {
		filterData = append(filterData, "third_party_id:"+strconv.FormatUint(req.Transaction.ThirdPartyID, 10))
	}

	filter := &db.ListFilter{
		Filter: strings.Join(filterData, ","),
	}

	logrus.Println("---------------------------")
	logrus.Println(filter.Filter)
	logrus.Println("---------------------------")

	mappingORM, err := s.provider.GetMapping(ctx, filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	beneficiaryIDs := []string{}
	for _, v := range mappingORM {
		if v.BeneficiaryID == 10101010 {
			res, err := s.GetBeneficiaryName(ctx, &pb.GetBeneficiaryNameRequest{ThirdPartyID: v.ThirdPartyID})
			if err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				}
			}

			for _, d := range res.Data {
				if !contains(beneficiaryIDs, strconv.FormatUint(d.BeneficiaryId, 10)) {
					beneficiaryIDs = append(beneficiaryIDs, strconv.FormatUint(d.BeneficiaryId, 10))
				}
			}
		} else {
			if !contains(beneficiaryIDs, strconv.FormatUint(v.BeneficiaryID, 10)) {
				beneficiaryIDs = append(beneficiaryIDs, strconv.FormatUint(v.BeneficiaryID, 10))
			}
		}
	}

	httpReqParamsOpt := ApiListTransactionRequest{
		Page:  uint64(req.Page),
		Limit: uint64(req.Limit),
	}

	logrus.Println("---------------------------")
	logrus.Println(strings.Join(beneficiaryIDs, ","))
	logrus.Println("---------------------------")

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

	res, err := ApiListTransaction(ctx, apiReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	if res.ResponseCode != "00" {

		logrus.Error("Failed To Transfer Data : ", res.ResponseMessage)

	} else {

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

		res, err := ApiListTransaction(ctx, apiReq)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		if res.ResponseCode != "00" {
			logrus.Error("Failed To Transfer Data : ", res.ResponseMessage)
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", res.ResponseMessage)
		}

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

	return result, nil
}

func (s *Server) CreateTransaction(ctx context.Context, req *pb.CreateTransactionRequest) (*pb.CreateTransactionResponse, error) {
	result := &pb.CreateTransactionResponse{
		Error:   false,
		Code:    200,
		Message: "Data",
	}

	me, err := s.manager.GetMeFromJWT(ctx, "")
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

		logrus.Println("----------------------")
		logrus.Println("Save BG Mapping")
		logrus.Println("----------------------")

		taskData := req.MappingData
		taskDataBak := req.MappingDataBackup

		ids := []string{}

		for _, v := range taskDataBak {

			needDelete := true

			for _, check := range taskData {

				if check.ThirdPartyID == v.ThirdPartyID {

					needDelete = false
					logrus.Println("Break at: " + strconv.FormatUint(check.ThirdPartyID, 10))
					break

				}

			}

			logrus.Println(needDelete)

			if needDelete {

				logrus.Println("----------------------")
				logrus.Println("Get Mapping Digital Task Data")
				logrus.Println(v)
				logrus.Println("----------------------")

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

				logrus.Println("----------------------")
				logrus.Println("Mapping Digital Task Response:")
				logrus.Println(taskMappingDigitalRes.Data)
				logrus.Println("----------------------")

				for _, taskMappingDigitalResData := range taskMappingDigitalRes.Data {

					taskMappingDigitalData := []*pb.MappingDigitalData{}
					json.Unmarshal([]byte(taskMappingDigitalResData.GetData()), &taskMappingDigitalData)
					if err != nil {
						return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
					}

					logrus.Println("----------------------")
					logrus.Println("To Delete Mapping Digital Task ID: " + strconv.FormatUint(taskMappingDigitalResData.TaskID, 10))
					logrus.Println("----------------------")

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

						for _, mappingORM := range mappingORMs {
							if mappingORM.Id > 0 {
								if !contains(ids, strconv.FormatUint(mappingORM.Id, 10)) {
									ids = append(ids, strconv.FormatUint(mappingORM.Id, 10))
								}
							}
						}
					}

				}

				logrus.Println("----------------------")
				logrus.Println("To Delete Mapping Digital Data:")
				logrus.Println(ids)
				logrus.Println("----------------------")

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

		logrus.Println("----------------------")
		logrus.Println("To Delete Mapping Data:")
		logrus.Println(ids)
		logrus.Println("----------------------")

		err = s.provider.DeleteMapping(ctx, ids)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		for _, v := range taskData {

			data := &pb.MappingORM{
				CompanyID:     v.CompanyID,
				ThirdPartyID:  v.ThirdPartyID,
				BeneficiaryID: 10101010,
				IsMapped:      false,
				CreatedByID:   me.UserID,
				UpdatedByID:   me.UserID,
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

	case "BG Mapping Digital":

		taskData := req.MappingDigitalData
		taskDataBak := req.MappingDigitalDataBackup

		ids := []string{}

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

			for _, mappingORM := range mappingORMs {
				if mappingORM.Id > 0 {
					if !contains(ids, strconv.FormatUint(mappingORM.Id, 10)) {
						ids = append(ids, strconv.FormatUint(mappingORM.Id, 10))
					}
				}
			}

		}

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

			for _, mappingORM := range mappingORMs {
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

		for _, v := range taskData {

			data := &pb.MappingORM{
				CompanyID:     v.CompanyID,
				ThirdPartyID:  v.ThirdPartyID,
				BeneficiaryID: v.BeneficiaryId,
				IsMapped:      true,
				CreatedByID:   me.UserID,
				UpdatedByID:   me.UserID,
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

	return result, nil
}

func (s *Server) DeleteTransaction(ctx context.Context, req *pb.DeleteTransactionRequest) (*pb.DeleteTransactionResponse, error) {
	result := &pb.DeleteTransactionResponse{
		Error:   false,
		Code:    200,
		Message: "Data",
	}

	me, err := s.manager.GetMeFromJWT(ctx, "")
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

		logrus.Println("----------------------")
		logrus.Println("Save BG Mapping")
		logrus.Println("----------------------")

		taskData := req.MappingData
		// taskDataBak := req.MappingDataBackup

		ids := []string{}

		for _, v := range taskData {

			logrus.Println("----------------------")
			logrus.Println("Get Mapping Digital Task Data")
			logrus.Println(v)
			logrus.Println("----------------------")

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

			logrus.Println("----------------------")
			logrus.Println("Mapping Digital Task Response:")
			logrus.Println(taskMappingDigitalRes.Data)
			logrus.Println("----------------------")

			for _, taskMappingDigitalResData := range taskMappingDigitalRes.Data {

				taskMappingDigitalData := []*pb.MappingDigitalData{}
				json.Unmarshal([]byte(taskMappingDigitalResData.GetData()), &taskMappingDigitalData)
				if err != nil {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				}

				logrus.Println("----------------------")
				logrus.Println("To Delete Mapping Digital Task ID: " + strconv.FormatUint(taskMappingDigitalResData.TaskID, 10))
				logrus.Println("----------------------")

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

					for _, mappingORM := range mappingORMs {
						if mappingORM.Id > 0 {
							if !contains(ids, strconv.FormatUint(mappingORM.Id, 10)) {
								ids = append(ids, strconv.FormatUint(mappingORM.Id, 10))
							}
						}
					}
				}

			}

			logrus.Println("----------------------")
			logrus.Println("To Delete Mapping Digital Data:")
			logrus.Println(ids)
			logrus.Println("----------------------")

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

		logrus.Println("----------------------")
		logrus.Println("To Delete Mapping Data:")
		logrus.Println(ids)
		logrus.Println("----------------------")

		err = s.provider.DeleteMapping(ctx, ids)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

	case "BG Mapping Digital":

		taskData := req.MappingDigitalData
		// taskDataBak := req.MappingDigitalDataBackup

		ids := []string{}

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

			for _, mappingORM := range mappingORMs {
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

		for _, v := range taskData {

			filter := []string{
				"company_id:" + strconv.FormatUint(v.CompanyID, 10),
			}

			taskMappingRes, err := taskClient.GetListTask(ctx, &task_pb.ListTaskRequest{Filter: strings.Join(filter, ","), Task: &task_pb.Task{Type: "BG Mapping"}, Page: 1, Limit: 1}, grpc.Header(&header), grpc.Trailer(&trailer))
			if err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				}
			}

			for _, d := range taskMappingRes.Data {

				taskMappingData := []*pb.MappingData{}
				json.Unmarshal([]byte(d.GetData()), &taskMappingData)
				if err != nil {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				}

				for _, dd := range taskMappingData {

					if v.ThirdPartyID == dd.ThirdPartyID {

						data := &pb.MappingORM{
							CompanyID:     v.CompanyID,
							ThirdPartyID:  v.ThirdPartyID,
							BeneficiaryID: 10101010,
							IsMapped:      dd.IsAllowAllBeneficiary,
							CreatedByID:   me.UserID,
							UpdatedByID:   me.UserID,
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

	return result, nil
}

func (s *Server) CreateIssuing(ctx context.Context, req *pb.CreateIssuingRequest) (*pb.CreateIssuingResponse, error) {
	result := &pb.CreateIssuingResponse{
		Error:   false,
		Code:    200,
		Message: "Data",
	}

	// me, err := s.manager.GetMeFromJWT(ctx, "")
	// if err != nil {
	// 	return nil, err
	// }

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		ctx = metadata.NewOutgoingContext(context.Background(), md)
	}
	var header, trailer metadata.MD

	systemConn, err := grpc.Dial(getEnv("SYSTEM_SERVICE", ":9101"), opts...)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed connect to System Service: %v", err)
	}
	defer systemConn.Close()

	systemClient := system_pb.NewApiServiceClient(systemConn)

	fileConn, err := grpc.Dial(getEnv("FILELISTENER_SERVICE", ":9201"), opts...)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed connect to System Service: %v", err)
	}
	defer fileConn.Close()

	fileClient := filelistener_pb.NewFileProcessorServiceClient(fileConn)

	isIndividu := uint64(req.Data.Applicant.GetApplicantType().Number())
	dateEstablished := ""

	if isIndividu == 0 {
		dateEstablished = req.Data.Applicant.GetDateEstablished()
		if dateEstablished == "" {
			return nil, status.Errorf(codes.InvalidArgument, "Internal Error: %v", "Empty value on dateEstablished when isIndividu is true")
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

	openingBranchORMs, err := systemClient.ListMdBranch(ctx, &system_pb.ListMdBranchRequest{
		Data: &system_pb.MdBranch{
			Id: req.Data.Publishing.GetOpeningBranchId(),
		},
	}, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "Opening Branch not found")
		} else {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}
	}

	if len(openingBranchORMs.Data) == 0 {
		return nil, status.Errorf(codes.NotFound, "Opening Branch not found")
	}

	openingBranch := openingBranchORMs.Data[0]

	publishingBranchORMs, err := systemClient.ListMdBranch(ctx, &system_pb.ListMdBranchRequest{
		Data: &system_pb.MdBranch{
			Id: req.Data.Publishing.GetPublishingBranchId(),
		},
	}, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "Publishing Branch not found")
		} else {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}
	}

	if len(publishingBranchORMs.Data) == 0 {
		return nil, status.Errorf(codes.NotFound, "Opening Branch not found")
	}

	publishingBranch := publishingBranchORMs.Data[0]

	if req.Data.Publishing.BgType == pb.BgType_GovernmentPaymentGuarantee {
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
		// nonCashAccountNo = req.Data.Project.GetNonCashAccountNo()
		// nonCashAccountAmount = req.Data.Project.GetNonCashAccountAmount()

		if cashAccountNo == "" || cashAccountAmount <= 0.0 {
			return nil, status.Errorf(codes.InvalidArgument, "Bad Request: %v", "Empty value on required field(s) when hold account is selected")
		}

	case pb.ContractGuaranteeType_NonCashLoan: // Non Cash Loan

		counterGuaranteeTypeString = map[string]string{"0": "customer limit"}

		// cashAccountNo = req.Data.Project.GetCashAccountNo()
		// cashAccountAmount = req.Data.Project.GetCashAccountAmount()
		nonCashAccountNo = req.Data.Project.GetNonCashAccountNo()
		nonCashAccountAmount = req.Data.Project.GetNonCashAccountAmount()

		inquiryLimit, err := ApiInquiryLimitIndividual(ctx, &ApiInquiryLimitIndividualRequest{Cif: req.Data.Account.Cif})
		if err != nil {
			logrus.Println("Error Limit Individual: ", err.Error())
			return nil, status.Errorf(codes.NotFound, "You are not allowed for Non Cash Loan facility")
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

		inquiryLimit, err := ApiInquiryLimitIndividual(ctx, &ApiInquiryLimitIndividualRequest{Cif: req.Data.Account.Cif})
		if err != nil {
			logrus.Println("Error Limit Individual: ", err.Error())
			return nil, status.Errorf(codes.NotFound, "You are not allowed for Combination facility")
		}

		customerLimitId = strconv.FormatUint(inquiryLimit.ResponseData[0].CustomerLimitId, 10)
		customerLimitAmount = float64(inquiryLimit.ResponseData[0].AvailableAmount)

		isEndOfYearBg = "1"

		if nonCashAccountNo == "" || nonCashAccountAmount <= 0.0 || cashAccountNo == "" || cashAccountAmount <= 0.0 {
			return nil, status.Errorf(codes.InvalidArgument, "Bad Request: %v", "Empty value on required field(s) when combination account is selected")
		}

	default:

		return nil, status.Errorf(codes.InvalidArgument, "Bad Request: %v", "Invalid Contract Guarantee Type")

	}

	openingBranchPadded := fmt.Sprintf("%05d", openingBranch.Id)
	publishingBranchPadded := fmt.Sprintf("%05d", publishingBranch.Id)

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
		// LegalDocument:          req.Data.Document.GetBusinessLegal(),
		// ContractDocument:       req.Data.Document.GetBg(),
		// Sp3Document:            req.Data.Document.GetSp(),
		// OthersDocument:         req.Data.Document.GetOther(),
	}

	businessLegalFile, err := fileClient.FileDownloadHandler(ctx, &filelistener_pb.FileDownloadHandlerRequest{
		ObjectName:         req.Data.Document.GetBusinessLegal(),
		ContentDisposition: "inline",
	}, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	legalDocument, err := ApiUploadEncode(ctx, &ApiUploadEncodeRequest{
		ChannelId: getEnv("BG_CHANNEL_ID", "2"),
		Document:  base64.RawStdEncoding.EncodeToString(businessLegalFile.Data),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	httpReqData.LegalDocument = legalDocument.ResponseData.Filename

	bgFile, err := fileClient.FileDownloadHandler(ctx, &filelistener_pb.FileDownloadHandlerRequest{
		ObjectName:         req.Data.Document.GetBg(),
		ContentDisposition: "inline",
	}, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	contractDocument, err := ApiUploadEncode(ctx, &ApiUploadEncodeRequest{
		ChannelId: getEnv("BG_CHANNEL_ID", "2"),
		Document:  base64.RawStdEncoding.EncodeToString(bgFile.Data),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	httpReqData.ContractDocument = contractDocument.ResponseData.Filename

	spFile, err := fileClient.FileDownloadHandler(ctx, &filelistener_pb.FileDownloadHandlerRequest{
		ObjectName:         req.Data.Document.GetSp(),
		ContentDisposition: "inline",
	}, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	sp3Document, err := ApiUploadEncode(ctx, &ApiUploadEncodeRequest{
		ChannelId: getEnv("BG_CHANNEL_ID", "2"),
		Document:  base64.RawStdEncoding.EncodeToString(spFile.Data),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	httpReqData.Sp3Document = sp3Document.ResponseData.Filename

	otherFile, err := fileClient.FileDownloadHandler(ctx, &filelistener_pb.FileDownloadHandlerRequest{
		ObjectName:         req.Data.Document.GetOther(),
		ContentDisposition: "inline",
	}, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	otherDocument, err := ApiUploadEncode(ctx, &ApiUploadEncodeRequest{
		ChannelId: getEnv("BG_CHANNEL_ID", "2"),
		Document:  base64.RawStdEncoding.EncodeToString(otherFile.Data),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	httpReqData.OthersDocument = otherDocument.ResponseData.Filename

	createIssuingRes, err := ApiCreateIssuing(ctx, &httpReqData)
	if err != nil {
		return nil, err
	}

	httpReqParamsOpt := ApiBgTrackingRequest{
		RegistrationNo: createIssuingRes.Data.RegistrationNo,
	}

	apiReq := &httpReqParamsOpt

	checkIssuingRes, err := ApiCheckIssuingStatus(ctx, apiReq)
	if err != nil {
		return nil, err
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

	taskRes, err := taskClient.GetTaskByID(ctx, &task_pb.GetTaskByIDReq{ID: req.TaskID, Type: "BG Issuing"}, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	taskData := pb.IssuingData{}
	json.Unmarshal([]byte(taskRes.Data.GetData()), &taskData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	logrus.Print(taskData)

	httpReqParamsOpt := ApiBgTrackingRequest{
		RegistrationNo: taskData.RegistrationNo,
	}

	apiReq := &httpReqParamsOpt

	res, err := ApiCheckIssuingStatus(ctx, apiReq)
	if err != nil {
		return nil, err
	}

	// channelId, err := strconv.ParseUint(getEnv("BG_CHANNEL_ID", "2"), 10, 64)
	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	// }

	// httpTransactionReqParamsOpt := ApiListTransactionRequest{
	// 	ReferenceNo: res.Data.ReferenceNo,
	// 	ChannelId:   channelId,
	// }

	// apiTransactionReq := &httpTransactionReqParamsOpt

	// transactionRes, err := ApiListTransaction(ctx, apiTransactionReq)

	// if len(transactionRes.ResponseData) <= 0 {
	// 	return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	// }

	taskData.ReferenceNo = res.Data.ReferenceNo

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

	_, err = taskClient.UpdateTaskData(ctx, taskReq, grpc.Header(&header), grpc.Trailer(&trailer))
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

	decodedBytes, err := base64.StdEncoding.DecodeString(req.GetData())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Bad Request: File is corrupted")
	}

	contentType := http.DetectContentType(decodedBytes)

	if contentType != "application/pdf" {
		return nil, status.Errorf(codes.InvalidArgument, "Bad Request: Invalid filetype")
	}

	httpReqParamsOpt := ApiUploadEncodeRequest{
		Document: req.GetData(),
	}

	apiReq := &httpReqParamsOpt

	res, err := ApiUploadEncode(ctx, apiReq)
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

func (s *Server) checkAccountNoIsValid(ctx context.Context, accountConn *grpc.ClientConn, accountNo string) (bool, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		ctx = metadata.NewOutgoingContext(context.Background(), md)
	}
	var header, trailer metadata.MD

	accountClient := account_pb.NewApiServiceClient(accountConn)
	account, err := accountClient.ValidateAccount(ctx, &account_pb.ValidateAccountRequest{
		AccountNo: accountNo,
	}, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		logrus.Println("Error rpc", err)
		return false, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	if account.Error {
		if account.Code == 404 {
			logrus.Println("Error not found 2")
			return false, status.Errorf(codes.NotFound, "Bad Request: Account not found")
		}
		logrus.Println("Error internal")
		return false, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	return true, nil
}
