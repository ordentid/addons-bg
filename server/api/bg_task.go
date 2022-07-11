package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	company_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/company"
	task_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/task"
	"bitbucket.bri.co.id/scm/addons/addons-bg-service/server/pb"
	"github.com/google/go-querystring/query"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (s *Server) GetTransactionTask(ctx context.Context, req *pb.GetTransactionTaskRequest) (*pb.GetTransactionTaskResponse, error) {
	result := &pb.GetTransactionTaskResponse{
		Error:   false,
		Code:    200,
		Message: "List Data",
		Data:    []*pb.TransactionTask{},
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

	companyConn, err := grpc.Dial(getEnv("COMPANY_SERVICE", ":9092"), opts...)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed connect to Company Service: %v", err)
	}
	defer companyConn.Close()

	companyClient := company_pb.NewApiServiceClient(companyConn)

	statuses := []string{}
	// - Maker: 1. Draft, 2. Returned, 3. Pending, 4. Request for Delete, 5. Approved, 6. Rejected
	// - Signer: 1. Pending, 2. Request for Delete, 3. Approved, 4. Rejected
	if len(me.Authorities) > 0 {
		switch strings.ToLower(me.Authorities[0]) {
		case "maker":
			statuses = []string{"2", "3", "1", "6", "4", "5"}
			if len(req.Filter) > 0 {
				req.Filter = req.Filter + ","
			}
			req.Filter = req.Filter + "status:<>0,status:<>7"

		case "signer":
			statuses = []string{"1", "6", "4", "5"}
			if len(req.Filter) > 0 {
				req.Filter = req.Filter + ","
			}
			req.Filter = req.Filter + "status:<>0,status:<>2,status:<>3,status:<>7"

		default:
			return nil, status.Errorf(codes.PermissionDenied, "Authority Denied")
		}
	}

	customOrder := ""
	if req.Sort == "status" {
		direction := ">"
		if req.Dir.String() == "DESC" {
			direction = "<"
		}
		customOrder = "status|" + direction + "|" + strings.Join(statuses, ",")
		req.Sort = ""
		req.Dir = 0
	} else if req.Sort == "" {
		customOrder = "status|>|" + strings.Join(statuses, ",")
		req.Dir = 0
	}

	filter := &task_pb.Task{
		Type: "BG Mapping",
	}

	if req.Status.Number() > 0 {
		filter.Status = task_pb.Statuses(req.Status.Number())
	}
	if req.Step.Number() > 0 {
		filter.Step = task_pb.Steps(req.Step.Number())
	}

	dataReq := &task_pb.ListTaskRequest{
		Task:        filter,
		Limit:       req.GetLimit(),
		Page:        req.GetPage(),
		Sort:        req.GetSort(),
		Dir:         task_pb.ListTaskRequestDirection(req.GetDir()),
		Filter:      req.GetFilter(),
		Query:       req.GetQuery(),
		CustomOrder: customOrder,
		In:          me.TaskFilter,
	}

	dataList, err := taskClient.GetListTask(ctx, dataReq, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	for _, v := range dataList.Data {
		task := &pb.Task{
			LastApprovedByID:   v.GetLastApprovedByID(),
			LastApprovedByName: v.GetLastApprovedByName(),
			LastRejectedByID:   v.GetLastRejectedByID(),
			LastRejectedByName: v.GetLastRejectedByName(),
			FeatureID:          v.GetFeatureID(),
			TaskID:             v.GetTaskID(),
			Status:             v.GetStatus().Enum().String(),
			Type:               v.GetType(),
			Step:               v.GetStep().String(),
			CreatedAt:          v.GetCreatedAt(),
			CreatedByName:      v.GetCreatedByName(),
			UpdatedAt:          v.GetUpdatedAt(),
		}

		taskData := []*pb.TransactionTaskData{}
		json.Unmarshal([]byte(v.Data), &taskData)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		var company *pb.Company
		var transactionTaskData []*pb.TransactionTaskData

		for i, v := range taskData {
			if i == 0 && len(v.Transaction) > 0 {
				companyRes, err := companyClient.ListCompanyData(ctx, &company_pb.ListCompanyDataReq{CompanyID: v.Transaction[0].GetCompanyID()}, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				}

				if len(companyRes.GetData()) > 0 {
					company = &pb.Company{
						CompanyID:          companyRes.Data[0].GetCompanyID(),
						HoldingID:          companyRes.Data[0].GetHoldingID(),
						GroupName:          companyRes.Data[0].GetGroupName(),
						CompanyName:        companyRes.Data[0].GetCompanyName(),
						HoldingCompanyName: companyRes.Data[0].GetHoldingCompanyName(),
						CreatedAt:          companyRes.Data[0].GetCreatedAt(),
						UpdatedAt:          companyRes.Data[0].GetUpdatedAt(),
					}
				} else {
					return nil, status.Errorf(codes.NotFound, "Company not found.")
				}
			}

			thirdPartyORM, err := s.provider.GetThirdPartyDetail(ctx, &pb.ThirdPartyORM{ThirdPartyID: v.ThirdParty.GetThirdPartyID()})
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
			}

			thirdParty, err := thirdPartyORM.ToPB(ctx)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
			}

			transactionTaskData = append(transactionTaskData, &pb.TransactionTaskData{
				ThirdParty:  &thirdParty,
				Transaction: v.Transaction,
			})
		}

		result.Data = append(result.Data, &pb.TransactionTask{
			Task:    task,
			Company: company,
			Data:    transactionTaskData,
		})
	}

	result.Pagination = &pb.PaginationResponse{
		Limit:      dataList.GetPagination().GetLimit(),
		Page:       dataList.GetPagination().GetPage(),
		TotalRows:  dataList.GetPagination().GetTotalRows(),
		TotalPages: dataList.GetPagination().GetTotalPages(),
	}

	return result, nil
}

func (s *Server) GetTransactionTaskDetail(ctx context.Context, req *pb.GetTransactionTaskDetailRequest) (*pb.GetTransactionTaskDetailResponse, error) {
	result := &pb.GetTransactionTaskDetailResponse{
		Error:   false,
		Code:    200,
		Message: "Success",
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

	taskConn, err := grpc.Dial(getEnv("TASK_SERVICE", ":9090"), opts...)
	if err != nil {
		logrus.Errorln("Failed connect to Task Service: %v", err)
		return nil, status.Errorf(codes.Internal, "Error Internal")
	}
	taskConn.Connect()
	defer taskConn.Close()

	taskClient := task_pb.NewTaskServiceClient(taskConn)

	companyConn, err := grpc.Dial(getEnv("COMPANY_SERVICE", ":9092"), opts...)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed connect to Company Service: %v", err)
	}
	defer companyConn.Close()

	companyClient := company_pb.NewApiServiceClient(companyConn)

	taskRes, err := taskClient.GetTaskByID(ctx, &task_pb.GetTaskByIDReq{ID: req.TaskID, Type: "BG Mapping"}, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	task := &pb.Task{
		LastApprovedByID:   taskRes.Data.GetLastApprovedByID(),
		LastApprovedByName: taskRes.Data.GetLastApprovedByName(),
		LastRejectedByID:   taskRes.Data.GetLastRejectedByID(),
		LastRejectedByName: taskRes.Data.GetLastRejectedByName(),
		FeatureID:          taskRes.Data.GetFeatureID(),
		TaskID:             taskRes.Data.GetTaskID(),
		Status:             taskRes.Data.GetStatus().Enum().String(),
		Type:               taskRes.Data.GetType(),
		Step:               taskRes.Data.GetStep().String(),
		CreatedAt:          taskRes.Data.GetCreatedAt(),
		CreatedByName:      taskRes.Data.GetCreatedByName(),
		UpdatedAt:          taskRes.Data.GetUpdatedAt(),
	}

	taskData := []*pb.TransactionTaskData{}
	json.Unmarshal([]byte(taskRes.Data.GetData()), &taskData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	var company *pb.Company
	var transactionTaskData []*pb.TransactionTaskData

	for i, v := range taskData {
		if i == 0 && len(v.Transaction) > 0 {
			companyRes, err := companyClient.ListCompanyData(ctx, &company_pb.ListCompanyDataReq{CompanyID: v.Transaction[0].GetCompanyID()}, grpc.Header(&header), grpc.Trailer(&trailer))
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
			}

			if len(companyRes.GetData()) > 0 {
				company = &pb.Company{
					CompanyID:          companyRes.Data[0].GetCompanyID(),
					HoldingID:          companyRes.Data[0].GetHoldingID(),
					GroupName:          companyRes.Data[0].GetGroupName(),
					CompanyName:        companyRes.Data[0].GetCompanyName(),
					HoldingCompanyName: companyRes.Data[0].GetHoldingCompanyName(),
					CreatedAt:          companyRes.Data[0].GetCreatedAt(),
					UpdatedAt:          companyRes.Data[0].GetUpdatedAt(),
				}
			} else {
				return nil, status.Errorf(codes.NotFound, "Company not found.")
			}
		}

		thirdPartyORM, err := s.provider.GetThirdPartyDetail(ctx, &pb.ThirdPartyORM{ThirdPartyID: v.ThirdParty.GetThirdPartyID()})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		thirdParty, err := thirdPartyORM.ToPB(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		transactionTaskData = append(transactionTaskData, &pb.TransactionTaskData{
			ThirdParty:  &thirdParty,
			Transaction: v.Transaction,
		})
	}

	result.Data = &pb.TransactionTask{
		Task:    task,
		Company: company,
		Data:    transactionTaskData,
	}

	return result, nil
}

func (s *Server) CreateTransactionTask(ctx context.Context, req *pb.CreateTransactionTaskRequest) (*pb.CreateTransactionTaskResponse, error) {
	result := &pb.CreateTransactionTaskResponse{
		Error:   false,
		Code:    200,
		Message: "Data",
	}

	me, err := s.manager.GetMeFromJWT(ctx, "")
	if err != nil {
		return nil, err
	}

	// proxyURL, err := url.Parse("http://localhost:5002")
	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	// }

	// client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
	client := &http.Client{}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		ctx = metadata.NewOutgoingContext(context.Background(), md)
	}
	var header, trailer metadata.MD

	companyConn, err := grpc.Dial(getEnv("COMPANY_SERVICE", ":9092"), opts...)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed connect to Company Service: %v", err)
	}
	defer companyConn.Close()

	companyClient := company_pb.NewApiServiceClient(companyConn)

	company, err := companyClient.ListCompanyData(ctx, &company_pb.ListCompanyDataReq{CompanyID: req.CompanyID}, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}
	if !(len(company.GetData()) > 0) {
		return nil, status.Errorf(codes.NotFound, "Company not found.")
	}

	taskConn, err := grpc.Dial(getEnv("TASK_SERVICE", ":9090"), opts...)
	if err != nil {
		logrus.Errorln("Failed connect to Task Service: %v", err)
		return nil, status.Errorf(codes.Internal, "Error Internal")
	}
	defer taskConn.Close()

	taskClient := task_pb.NewTaskServiceClient(taskConn)

	taskData := []*pb.TransactionTaskData{}
	for _, v := range req.ThirdParty {
		httpReqParamsOpt := ApiListTransactionRequest{
			ThirdPartyId: strconv.FormatUint(v.ThirdPartyID, 10),
			Page:         1,
			Limit:        100,
		}

		httpReqParams, err := query.Values(httpReqParamsOpt)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		logrus.Println(httpReqParams.Encode())

		httpReq, err := http.NewRequest("GET", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0/listTransaction?"+httpReqParams.Encode(), nil)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		httpReq.Header.Add("Authorization", "Basic YnJpY2FtczpCcmljYW1zNGRkMG5z")

		httpRes, err := client.Do(httpReq)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}
		defer httpRes.Body.Close()

		var httpResData ApiListTransactionResponse
		httpResBody, err := ioutil.ReadAll(httpRes.Body)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		err = json.Unmarshal(httpResBody, &httpResData)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		if httpResData.ResponseCode != "00" {
			logrus.Error("Failed To Transfer Data : ", httpResData.ResponseMessage)
		} else {
			transactionDataList := []*pb.Transaction{}
			for _, d := range httpResData.ResponseData {
				transactionData := &pb.Transaction{
					Amount:             d.Amount,
					ApplicantName:      d.ApplicantName,
					BeneficiaryName:    d.BeneficiaryName,
					ChannelID:          d.ChannelId,
					ChannelName:        d.ChannelName,
					ClaimPeriod:        d.ClaimPeriod,
					ClosingDate:        d.ClosingDate,
					CompanyID:          req.CompanyID,
					CreatedByID:        me.UserID,
					Currency:           d.Currency,
					DocumentPath:       d.DocumentPath,
					EffectiveDate:      d.EffectiveDate,
					ExpiryDate:         d.ExpiryDate,
					IsAllowBeneficiary: v.IsAllowBeneficiary,
					IssueDate:          d.IssueDate,
					ReferenceNo:        d.ReferenceNo,
					RegistrationNo:     d.RegistrationNo,
					Remark:             d.Remark,
					Status:             "Pending",
					ThirdPartyID:       d.ThirdPartyId,
					TransactionID:      d.TransactionId,
					TransactionStatus:  d.Status,
					TransactionTypeID:  d.TransactionTypeId,
					UpdatedByID:        me.UserID,
				}
				transactionDataList = append(transactionDataList, transactionData)
			}
			taskData = append(taskData, &pb.TransactionTaskData{
				ThirdParty: &pb.ThirdParty{
					Id:   v.GetThirdPartyID(),
					Name: fmt.Sprintf("THIRD PARTY %s", v.GetThirdPartyID()),
				},
				Transaction: transactionDataList,
			})
		}
	}

	data, err := json.Marshal(taskData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	taskReq := &task_pb.SaveTaskRequest{
		TaskID: req.TaskID,
		Task: &task_pb.Task{
			Type:        "BG Mapping",
			Data:        string(data),
			CreatedByID: me.UserID,
		},
	}

	if req.IsDraft {
		taskReq.IsDraft = true
	}

	taskRes, err := taskClient.SaveTaskWithData(ctx, taskReq, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	result.Data = &pb.Task{
		TaskID:             taskRes.Data.TaskID,
		Type:               taskRes.Data.Type,
		Status:             taskRes.Data.Status.String(),
		Step:               taskRes.Data.Step.String(),
		FeatureID:          taskRes.Data.FeatureID,
		LastApprovedByID:   taskRes.Data.LastApprovedByID,
		LastRejectedByID:   taskRes.Data.LastRejectedByID,
		LastApprovedByName: taskRes.Data.LastApprovedByName,
		LastRejectedByName: taskRes.Data.LastRejectedByName,
		CreatedByName:      taskRes.Data.CreatedByName,
		UpdatedByName:      taskRes.Data.UpdatedByName,
		Reasons:            taskRes.Data.Reasons,
		Comment:            taskRes.Data.Comment,
		CreatedAt:          taskRes.Data.CreatedAt,
		UpdatedAt:          taskRes.Data.UpdatedAt,
	}

	return result, nil
}
