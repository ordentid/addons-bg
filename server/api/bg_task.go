package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	company_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/company"
	task_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/task"
	workflow_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/workflow"
	"bitbucket.bri.co.id/scm/addons/addons-bg-service/server/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (s *Server) GetTaskMapping(ctx context.Context, req *pb.GetTaskMappingRequest) (*pb.GetTaskMappingResponse, error) {

	result := &pb.GetTaskMappingResponse{
		Error:   false,
		Code:    200,
		Message: "List Data",
		Data:    []*pb.TaskMappingData{},
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		ctx = metadata.NewOutgoingContext(context.Background(), md)
	}

	currentUser, userMD, err := s.manager.GetMeFromMD(ctx)
	if err != nil {
		return nil, err
	}
	if currentUser == nil {
		return nil, s.unauthorizedError()
	}
	var trailer metadata.MD

	taskClient := s.scvConn.TaskServiceClient()
	companyClient := s.scvConn.CompanyServiceClient()

	statuses := []string{}
	// - Maker: 1. Draft, 2. Returned, 3. Pending, 4. Request for Delete, 5. Approved, 6. Rejected
	// - Signer: 1. Pending, 2. Request for Delete, 3. Approved, 4. Rejected
	if len(currentUser.Authorities) > 0 {
		switch strings.ToLower(currentUser.Authorities[0]) {
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
	}

	dataList, err := taskClient.GetListTask(ctx, dataReq, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	for _, v := range dataList.Data {
		task := &pb.Task{
			TaskID:             v.GetTaskID(),
			Type:               v.GetType(),
			Status:             v.GetStatus().String(),
			Step:               v.GetStep().String(),
			FeatureID:          v.GetFeatureID(),
			LastApprovedByID:   v.GetLastApprovedByID(),
			LastRejectedByID:   v.GetLastRejectedByID(),
			LastApprovedByName: v.GetLastApprovedByName(),
			LastRejectedByName: v.GetLastRejectedByName(),
			CreatedByName:      v.GetCreatedByName(),
			UpdatedByName:      v.GetUpdatedByName(),
			Reasons:            v.GetReasons(),
			Comment:            v.GetComment(),
			CompanyID:          v.GetCompanyID(),
			HoldingID:          v.GetHoldingID(),
			CreatedAt:          v.GetCreatedAt(),
			UpdatedAt:          v.GetUpdatedAt(),
		}

		taskData := []*pb.MappingData{}
		json.Unmarshal([]byte(v.Data), &taskData)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		var company *pb.Company

		companyRes, err := companyClient.ListCompanyDataV2(ctx, &company_pb.ListCompanyDataReq{CompanyID: v.GetCompanyID()}, grpc.Header(&userMD), grpc.Trailer(&trailer))
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

		result.Data = append(result.Data, &pb.TaskMappingData{
			Task:    task,
			Company: company,
			Data:    taskData,
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

func (s *Server) GetTaskMappingDetail(ctx context.Context, req *pb.GetTaskMappingDetailRequest) (*pb.GetTaskMappingDetailResponse, error) {

	result := &pb.GetTaskMappingDetailResponse{
		Error:   false,
		Code:    200,
		Message: "Success",
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		ctx = metadata.NewOutgoingContext(context.Background(), md)
	}

	currentUser, userMD, err := s.manager.GetMeFromMD(ctx)
	if err != nil {
		return nil, err
	}
	if currentUser == nil {
		return nil, s.unauthorizedError()
	}
	var trailer metadata.MD

	taskClient := s.scvConn.TaskServiceClient()
	companyClient := s.scvConn.CompanyServiceClient()

	taskRes, err := taskClient.GetTaskByID(ctx, &task_pb.GetTaskByIDReq{ID: req.TaskID, Type: "BG Mapping"}, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	task := &pb.Task{
		TaskID:             taskRes.Data.GetTaskID(),
		Type:               taskRes.Data.GetType(),
		Status:             taskRes.Data.GetStatus().String(),
		Step:               taskRes.Data.GetStep().String(),
		FeatureID:          taskRes.Data.GetFeatureID(),
		LastApprovedByID:   taskRes.Data.GetLastApprovedByID(),
		LastRejectedByID:   taskRes.Data.GetLastRejectedByID(),
		LastApprovedByName: taskRes.Data.GetLastApprovedByName(),
		LastRejectedByName: taskRes.Data.GetLastRejectedByName(),
		CreatedByName:      taskRes.Data.GetCreatedByName(),
		UpdatedByName:      taskRes.Data.GetUpdatedByName(),
		Reasons:            taskRes.Data.GetReasons(),
		Comment:            taskRes.Data.GetComment(),
		CompanyID:          taskRes.Data.GetCompanyID(),
		HoldingID:          taskRes.Data.GetHoldingID(),
		CreatedAt:          taskRes.Data.GetCreatedAt(),
		UpdatedAt:          taskRes.Data.GetUpdatedAt(),
	}

	taskData := []*pb.MappingData{}
	json.Unmarshal([]byte(taskRes.Data.GetData()), &taskData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	var company *pb.Company

	companyRes, err := companyClient.ListCompanyDataV2(ctx, &company_pb.ListCompanyDataReq{CompanyID: taskRes.Data.GetCompanyID()}, grpc.Header(&userMD), grpc.Trailer(&trailer))
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

	result.Data = &pb.TaskMappingData{
		Task:    task,
		Company: company,
		Data:    taskData,
	}

	return result, nil

}

func (s *Server) CreateTaskMapping(ctx context.Context, req *pb.CreateTaskMappingRequest) (*pb.CreateTaskMappingResponse, error) {

	result := &pb.CreateTaskMappingResponse{
		Error:   false,
		Code:    200,
		Message: "Success",
	}

	client := &http.Client{}
	if getEnv("ENV", "PRODUCTION") != "PRODUCTION" {
		proxyURL, err := url.Parse("http://localhost:5100")
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
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
		return nil, s.unauthorizedError()
	}
	var trailer metadata.MD

	taskClient := s.scvConn.TaskServiceClient()
	companyClient := s.scvConn.CompanyServiceClient()

	company, err := companyClient.ListCompanyDataV2(newCtx, &company_pb.ListCompanyDataReq{CompanyID: req.CompanyID}, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}
	if !(len(company.GetData()) > 0) {
		return nil, status.Errorf(codes.NotFound, "Company not found.")
	}

	taskData := []*pb.MappingData{}
	for _, v := range req.Data {
		name := ""

		httpReqData := ApiInquiryThirdPartyByIDRequest{
			ThirdPartyID: v.ThirdPartyID,
		}

		httpReqPayload, err := json.Marshal(httpReqData)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		httpReq, err := http.NewRequest("POST", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0/inquiryThirdParty", bytes.NewBuffer(httpReqPayload))
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		httpReq.Header.Add("Content-Type", "application/json")
		httpReq.Header.Add("Authorization", "Basic YnJpY2FtczpCcmljYW1zNGRkMG5z")

		httpRes, err := client.Do(httpReq)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}
		defer httpRes.Body.Close()

		var httpResData ApiInquiryThirdPartyByIDResponse
		err = json.NewDecoder(httpRes.Body).Decode(&httpResData)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		logrus.Println(httpResData.ResponseCode)

		if httpResData.ResponseCode == "00" {
			name = httpResData.ResponseData.FullName
		}

		taskData = append(taskData, &pb.MappingData{
			ThirdPartyID:          v.GetThirdPartyID(),
			ThirdPartyName:        name,
			CompanyID:             company.Data[0].CompanyID,
			CompanyName:           company.Data[0].CompanyName,
			IsAllowAllBeneficiary: v.GetIsAllowAllBeneficiary(),
		})
	}

	data, err := json.Marshal(taskData)
	if err != nil {
		logrus.Error("Failed To Marshal : ", taskData)
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	taskReq := &task_pb.SaveTaskRequest{
		TaskID: req.TaskID,
		Task: &task_pb.Task{
			Type:        "BG Mapping",
			Data:        string(data),
			CreatedByID: currentUser.UserID,
			CompanyID:   req.GetCompanyID(),
		},
	}

	if req.IsDraft {
		taskReq.IsDraft = true
	}

	taskRes, err := taskClient.SaveTaskWithData(newCtx, taskReq, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		logrus.Error("Failed To Transfer Data : ", "FAK")
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	result.Data = &pb.Task{
		TaskID:             taskRes.Data.GetTaskID(),
		Type:               taskRes.Data.GetType(),
		Status:             taskRes.Data.GetStatus().String(),
		Step:               taskRes.Data.GetStep().String(),
		FeatureID:          taskRes.Data.GetFeatureID(),
		LastApprovedByID:   taskRes.Data.GetLastApprovedByID(),
		LastRejectedByID:   taskRes.Data.GetLastRejectedByID(),
		LastApprovedByName: taskRes.Data.GetLastApprovedByName(),
		LastRejectedByName: taskRes.Data.GetLastRejectedByName(),
		CreatedByName:      taskRes.Data.GetCreatedByName(),
		UpdatedByName:      taskRes.Data.GetUpdatedByName(),
		Reasons:            taskRes.Data.GetReasons(),
		Comment:            taskRes.Data.GetComment(),
		CompanyID:          taskRes.Data.GetCompanyID(),
		HoldingID:          taskRes.Data.GetHoldingID(),
		CreatedAt:          taskRes.Data.GetCreatedAt(),
		UpdatedAt:          taskRes.Data.GetUpdatedAt(),
	}

	return result, nil

}

func (s *Server) GetTaskMappingDigital(ctx context.Context, req *pb.GetTaskMappingDigitalRequest) (*pb.GetTaskMappingDigitalResponse, error) {

	result := &pb.GetTaskMappingDigitalResponse{
		Error:   false,
		Code:    200,
		Message: "List Data",
		Data:    []*pb.TaskMappingDigitalData{},
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
		return nil, s.unauthorizedError()
	}
	var trailer metadata.MD

	taskClient := s.scvConn.TaskServiceClient()
	companyClient := s.scvConn.CompanyServiceClient()

	statuses := []string{}
	// - Maker: 1. Draft, 2. Returned, 3. Pending, 4. Request for Delete, 5. Approved, 6. Rejected
	// - Signer: 1. Pending, 2. Request for Delete, 3. Approved, 4. Rejected
	if len(currentUser.Authorities) > 0 {
		switch strings.ToLower(currentUser.Authorities[0]) {
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
		Type:      "BG Mapping Digital",
		CompanyID: currentUser.CompanyID,
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
	}

	dataList, err := taskClient.GetListTask(newCtx, dataReq, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	for _, v := range dataList.Data {
		task := &pb.Task{
			TaskID:             v.GetTaskID(),
			Type:               v.GetType(),
			Status:             v.GetStatus().String(),
			Step:               v.GetStep().String(),
			FeatureID:          v.GetFeatureID(),
			LastApprovedByID:   v.GetLastApprovedByID(),
			LastRejectedByID:   v.GetLastRejectedByID(),
			LastApprovedByName: v.GetLastApprovedByName(),
			LastRejectedByName: v.GetLastRejectedByName(),
			CreatedByName:      v.GetCreatedByName(),
			UpdatedByName:      v.GetUpdatedByName(),
			Reasons:            v.GetReasons(),
			Comment:            v.GetComment(),
			CompanyID:          v.GetCompanyID(),
			HoldingID:          v.GetHoldingID(),
			CreatedAt:          v.GetCreatedAt(),
			UpdatedAt:          v.GetUpdatedAt(),
		}

		taskData := []*pb.MappingDigitalData{}
		json.Unmarshal([]byte(v.Data), &taskData)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		var company *pb.Company

		companyRes, err := companyClient.ListCompanyDataV2(newCtx, &company_pb.ListCompanyDataReq{CompanyID: v.GetCompanyID()}, grpc.Header(&userMD), grpc.Trailer(&trailer))
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

		result.Data = append(result.Data, &pb.TaskMappingDigitalData{
			Task:    task,
			Company: company,
			Data:    taskData,
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

func (s *Server) GetTaskMappingDigitalDetail(ctx context.Context, req *pb.GetTaskMappingDigitalDetailRequest) (*pb.GetTaskMappingDigitalDetailResponse, error) {

	result := &pb.GetTaskMappingDigitalDetailResponse{
		Error:   false,
		Code:    200,
		Message: "Success",
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
		return nil, s.unauthorizedError()
	}
	var trailer metadata.MD

	taskClient := s.scvConn.TaskServiceClient()
	companyClient := s.scvConn.CompanyServiceClient()

	taskRes, err := taskClient.GetTaskByID(newCtx, &task_pb.GetTaskByIDReq{ID: req.TaskID, Type: "BG Mapping Digital"}, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	task := &pb.Task{
		TaskID:             taskRes.Data.GetTaskID(),
		Type:               taskRes.Data.GetType(),
		Status:             taskRes.Data.GetStatus().String(),
		Step:               taskRes.Data.GetStep().String(),
		FeatureID:          taskRes.Data.GetFeatureID(),
		LastApprovedByID:   taskRes.Data.GetLastApprovedByID(),
		LastRejectedByID:   taskRes.Data.GetLastRejectedByID(),
		LastApprovedByName: taskRes.Data.GetLastApprovedByName(),
		LastRejectedByName: taskRes.Data.GetLastRejectedByName(),
		CreatedByName:      taskRes.Data.GetCreatedByName(),
		UpdatedByName:      taskRes.Data.GetUpdatedByName(),
		Reasons:            taskRes.Data.GetReasons(),
		Comment:            taskRes.Data.GetComment(),
		CompanyID:          taskRes.Data.GetCompanyID(),
		HoldingID:          taskRes.Data.GetHoldingID(),
		CreatedAt:          taskRes.Data.GetCreatedAt(),
		UpdatedAt:          taskRes.Data.GetUpdatedAt(),
	}

	taskData := []*pb.MappingDigitalData{}
	json.Unmarshal([]byte(taskRes.Data.GetData()), &taskData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	var company *pb.Company

	companyRes, err := companyClient.ListCompanyDataV2(newCtx, &company_pb.ListCompanyDataReq{CompanyID: task.GetCompanyID()}, grpc.Header(&userMD), grpc.Trailer(&trailer))
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

	result.Data = &pb.TaskMappingDigitalData{
		Task:    task,
		Company: company,
		Data:    taskData,
	}

	return result, nil

}

func (s *Server) CreateTaskMappingDigital(ctx context.Context, req *pb.CreateTaskMappingDigitalRequest) (*pb.CreateTaskMappingDigitalResponse, error) {

	result := &pb.CreateTaskMappingDigitalResponse{
		Error:   false,
		Code:    200,
		Message: "Success",
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
		return nil, s.unauthorizedError()
	}
	var trailer metadata.MD

	taskClient := s.scvConn.TaskServiceClient()
	companyClient := s.scvConn.CompanyServiceClient()

	company, err := companyClient.ListCompanyDataV2(newCtx, &company_pb.ListCompanyDataReq{CompanyID: currentUser.CompanyID}, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}
	if !(len(company.GetData()) > 0) {
		return nil, status.Errorf(codes.NotFound, "Company not found.")
	}

	client := &http.Client{}
	if getEnv("ENV", "PRODUCTION") != "PRODUCTION" {
		proxyURL, err := url.Parse("http://localhost:5100")
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
	}

	name := ""

	httpReqData := ApiInquiryThirdPartyByIDRequest{
		ThirdPartyID: req.ThirdPartyID,
	}

	httpReqPayload, err := json.Marshal(httpReqData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	httpReq, err := http.NewRequest("POST", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0/inquiryThirdParty", bytes.NewBuffer(httpReqPayload))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", "Basic YnJpY2FtczpCcmljYW1zNGRkMG5z")

	httpRes, err := client.Do(httpReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}
	defer httpRes.Body.Close()

	var httpResData ApiInquiryThirdPartyByIDResponse
	err = json.NewDecoder(httpRes.Body).Decode(&httpResData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	logrus.Println(httpResData.ResponseCode)

	if httpResData.ResponseCode == "00" {
		name = httpResData.ResponseData.FullName
	}

	taskData := []*pb.MappingDigitalData{}
	for _, v := range req.Beneficiary {
		taskData = append(taskData, &pb.MappingDigitalData{
			ThirdPartyID:    req.GetThirdPartyID(),
			ThirdPartyName:  name,
			CompanyID:       company.Data[0].CompanyID,
			CompanyName:     company.Data[0].CompanyName,
			BeneficiaryId:   v.BeneficiaryId,
			BeneficiaryName: v.BeneficiaryName,
		})
	}

	logrus.Println("--------------------")
	logrus.Println(taskData)
	logrus.Println("--------------------")

	data, err := json.Marshal(taskData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	logrus.Println("--------------------")
	logrus.Println(string(data))
	logrus.Println("--------------------")

	taskReq := &task_pb.SaveTaskRequest{
		TaskID: req.TaskID,
		Task: &task_pb.Task{
			Type:        "BG Mapping Digital",
			Data:        string(data),
			CreatedByID: currentUser.UserID,
			CompanyID:   currentUser.CompanyID,
		},
	}

	if req.IsDraft {
		taskReq.IsDraft = true
	}

	taskRes, err := taskClient.SaveTaskWithData(newCtx, taskReq, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	result.Data = &pb.Task{
		TaskID:             taskRes.Data.GetTaskID(),
		Type:               taskRes.Data.GetType(),
		Status:             taskRes.Data.GetStatus().String(),
		Step:               taskRes.Data.GetStep().String(),
		FeatureID:          taskRes.Data.GetFeatureID(),
		LastApprovedByID:   taskRes.Data.GetLastApprovedByID(),
		LastRejectedByID:   taskRes.Data.GetLastRejectedByID(),
		LastApprovedByName: taskRes.Data.GetLastApprovedByName(),
		LastRejectedByName: taskRes.Data.GetLastRejectedByName(),
		CreatedByName:      taskRes.Data.GetCreatedByName(),
		UpdatedByName:      taskRes.Data.GetUpdatedByName(),
		Reasons:            taskRes.Data.GetReasons(),
		Comment:            taskRes.Data.GetComment(),
		CompanyID:          taskRes.Data.GetCompanyID(),
		HoldingID:          taskRes.Data.GetHoldingID(),
		CreatedAt:          taskRes.Data.GetCreatedAt(),
		UpdatedAt:          taskRes.Data.GetUpdatedAt(),
	}

	return result, nil

}

func (s *Server) GetTaskIssuing(ctx context.Context, req *pb.GetTaskIssuingRequest) (*pb.GetTaskIssuingResponse, error) {
	result := &pb.GetTaskIssuingResponse{
		Error:   false,
		Code:    200,
		Message: "List Data",
		Data:    []*pb.TaskIssuingData{},
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
		return nil, s.unauthorizedError()
	}
	var trailer metadata.MD

	taskClient := s.scvConn.TaskServiceClient()
	companyClient := s.scvConn.CompanyServiceClient()

	statuses := []string{}
	// - Maker: 1. Draft, 2. Returned, 3. Pending, 4. Request for Delete, 5. Approved, 6. Rejected
	// - Signer: 1. Pending, 2. Request for Delete, 3. Approved, 4. Rejected
	if len(currentUser.Authorities) > 0 {
		switch strings.ToLower(currentUser.Authorities[0]) {
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
		Type:      "BG Issuing",
		CompanyID: currentUser.CompanyID,
	}

	logrus.Println(filter)

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
	}

	dataList, err := taskClient.GetListTask(newCtx, dataReq, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	for _, v := range dataList.Data {
		task := &pb.Task{
			TaskID:             v.GetTaskID(),
			Type:               v.GetType(),
			Status:             v.GetStatus().String(),
			Step:               v.GetStep().String(),
			FeatureID:          v.GetFeatureID(),
			LastApprovedByID:   v.GetLastApprovedByID(),
			LastRejectedByID:   v.GetLastRejectedByID(),
			LastApprovedByName: v.GetLastApprovedByName(),
			LastRejectedByName: v.GetLastRejectedByName(),
			CreatedByName:      v.GetCreatedByName(),
			UpdatedByName:      v.GetUpdatedByName(),
			Reasons:            v.GetReasons(),
			Comment:            v.GetComment(),
			CompanyID:          v.GetCompanyID(),
			HoldingID:          v.GetHoldingID(),
			CreatedAt:          v.GetCreatedAt(),
			UpdatedAt:          v.GetUpdatedAt(),
		}

		taskData := pb.IssuingData{}
		json.Unmarshal([]byte(v.Data), &taskData)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		var company *pb.Company

		companyRes, err := companyClient.ListCompanyDataV2(newCtx, &company_pb.ListCompanyDataReq{CompanyID: v.CompanyID}, grpc.Header(&userMD), grpc.Trailer(&trailer))
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		workflow := pb.ValidateWorkflowData{}
		err = json.Unmarshal([]byte(v.GetWorkflowDoc()), &workflow)
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

		result.Data = append(result.Data, &pb.TaskIssuingData{
			Task:     task,
			Company:  company,
			Data:     &taskData,
			Workflow: &workflow,
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

func (s *Server) GetTaskIssuingDetail(ctx context.Context, req *pb.GetTaskIssuingDetailRequest) (*pb.GetTaskIssuingDetailResponse, error) {

	result := &pb.GetTaskIssuingDetailResponse{
		Error:   false,
		Code:    200,
		Message: "Success",
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
		return nil, s.unauthorizedError()
	}
	var trailer metadata.MD

	taskClient := s.scvConn.TaskServiceClient()
	companyClient := s.scvConn.CompanyServiceClient()

	taskRes, err := taskClient.GetTaskByID(newCtx, &task_pb.GetTaskByIDReq{ID: req.TaskID, Type: "BG Issuing"}, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	workflow := pb.ValidateWorkflowData{}
	err = json.Unmarshal([]byte(taskRes.Data.GetWorkflowDoc()), &workflow)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	task := &pb.Task{
		TaskID:             taskRes.Data.GetTaskID(),
		Type:               taskRes.Data.GetType(),
		Status:             taskRes.Data.GetStatus().String(),
		Step:               taskRes.Data.GetStep().String(),
		FeatureID:          taskRes.Data.GetFeatureID(),
		LastApprovedByID:   taskRes.Data.GetLastApprovedByID(),
		LastRejectedByID:   taskRes.Data.GetLastRejectedByID(),
		LastApprovedByName: taskRes.Data.GetLastApprovedByName(),
		LastRejectedByName: taskRes.Data.GetLastRejectedByName(),
		CreatedByName:      taskRes.Data.GetCreatedByName(),
		UpdatedByName:      taskRes.Data.GetUpdatedByName(),
		Reasons:            taskRes.Data.GetReasons(),
		Comment:            taskRes.Data.GetComment(),
		CompanyID:          taskRes.Data.GetCompanyID(),
		HoldingID:          taskRes.Data.GetHoldingID(),
		CreatedAt:          taskRes.Data.GetCreatedAt(),
		UpdatedAt:          taskRes.Data.GetUpdatedAt(),
	}

	taskData := pb.IssuingData{}
	json.Unmarshal([]byte(taskRes.Data.GetData()), &taskData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	var company *pb.Company

	companyRes, err := companyClient.ListCompanyDataV2(newCtx, &company_pb.ListCompanyDataReq{CompanyID: task.GetCompanyID()}, grpc.Header(&userMD), grpc.Trailer(&trailer))
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

	result.Data = &pb.TaskIssuingData{
		Task:     task,
		Company:  company,
		Data:     &taskData,
		Workflow: &workflow,
	}

	return result, nil

}

func (s *Server) CreateTaskIssuing(ctx context.Context, req *pb.CreateTaskIssuingRequest) (*pb.CreateTaskIssuingResponse, error) {

	result := &pb.CreateTaskIssuingResponse{
		Error:   false,
		Code:    200,
		Message: "Success",
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
		return nil, s.unauthorizedError()
	}
	var trailer metadata.MD

	taskClient := s.scvConn.TaskServiceClient()
	companyClient := s.scvConn.CompanyServiceClient()

	company, err := companyClient.ListCompanyDataV2(newCtx, &company_pb.ListCompanyDataReq{CompanyID: currentUser.CompanyID}, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}
	if !(len(company.GetData()) > 0) {
		return nil, status.Errorf(codes.NotFound, "Company not found.")
	}

	// client := &http.Client{}
	// if getEnv("ENV", "PRODUCTION") != "PRODUCTION" {
	// 	proxyURL, err := url.Parse("http://localhost:5100")
	// 	if err != nil {
	// 		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	// 	}

	// 	client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
	// }

	// httpReqParamsOpt := ApiListTransactionRequest{
	// 	ThirdPartyId: req.Data.Publishing.ThirdPartyID,
	// 	Page:         1,
	// 	Limit:        1,
	// }

	// httpReqParams, err := query.Values(httpReqParamsOpt)
	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	// }

	// logrus.Println(httpReqParams.Encode())

	// httpReq, err := http.NewRequest("GET", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0/listTransaction?"+httpReqParams.Encode(), nil)
	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	// }

	// httpReq.Header.Add("Authorization", "Basic YnJpY2FtczpCcmljYW1zNGRkMG5z")

	// httpRes, err := client.Do(httpReq)
	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	// }
	// defer httpRes.Body.Close()

	// var httpResData ApiListTransactionResponse
	// httpResBody, err := ioutil.ReadAll(httpRes.Body)
	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	// }

	// err = json.Unmarshal(httpResBody, &httpResData)
	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	// }

	// if httpResData.ResponseCode != "00" {
	// 	logrus.Error("Failed To Transfer Data : ", httpResData.ResponseMessage)
	// 	return nil, status.Errorf(codes.Internal, "Internal Error: %v", httpResData.ResponseMessage)
	// }

	taskData, err := json.Marshal(req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	taskReq := &task_pb.SaveTaskRequest{
		TaskID: req.TaskID,
		Task: &task_pb.Task{
			Type:        "BG Issuing",
			Data:        string(taskData),
			CreatedByID: currentUser.UserID,
			CompanyID:   currentUser.CompanyID,
		},
		TransactionAmount: req.Data.Project.BgAmount,
	}

	if req.IsDraft {
		taskReq.IsDraft = true
	}

	taskRes, err := taskClient.SaveTaskWithData(ctx, taskReq, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	result.Data = &pb.Task{
		TaskID:             taskRes.Data.GetTaskID(),
		Type:               taskRes.Data.GetType(),
		Status:             taskRes.Data.GetStatus().String(),
		Step:               taskRes.Data.GetStep().String(),
		FeatureID:          taskRes.Data.GetFeatureID(),
		LastApprovedByID:   taskRes.Data.GetLastApprovedByID(),
		LastRejectedByID:   taskRes.Data.GetLastRejectedByID(),
		LastApprovedByName: taskRes.Data.GetLastApprovedByName(),
		LastRejectedByName: taskRes.Data.GetLastRejectedByName(),
		CreatedByName:      taskRes.Data.GetCreatedByName(),
		UpdatedByName:      taskRes.Data.GetUpdatedByName(),
		Reasons:            taskRes.Data.GetReasons(),
		Comment:            taskRes.Data.GetComment(),
		CompanyID:          taskRes.Data.GetCompanyID(),
		HoldingID:          taskRes.Data.GetHoldingID(),
		CreatedAt:          taskRes.Data.GetCreatedAt(),
		UpdatedAt:          taskRes.Data.GetUpdatedAt(),
	}

	return result, nil

}

func (s *Server) TaskAction(ctx context.Context, req *pb.TaskActionRequest) (*pb.TaskActionResponse, error) {

	if req.GetAction() == "" || req.GetTaskID() < 1 {
		return nil, status.Error(codes.InvalidArgument, "Invalid Argument")
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
		return nil, s.unauthorizedError()
	}
	var trailer metadata.MD

	taskClient := s.scvConn.TaskServiceClient()
	workflowClient := s.scvConn.WorkflowServiceClient()

	// systemConn, err := grpc.Dial(getEnv("SYSTEM_SERVICE", ":9101"), opts...)
	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, "Failed connect to System Service: %v", err)
	// }
	// defer systemConn.Close()

	// systemClient := system_pb.NewApiServiceClient(systemConn)

	task, err := taskClient.GetTaskByID(ctx, &task_pb.GetTaskByIDReq{
		Type: "BG Issuing",
		ID:   req.GetTaskID(),
	})
	logrus.Println("task ===> ", task)
	if err != nil {
		logrus.Errorln("[api][func: TaskAction] error get task by ID: ", err)
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	if task.GetData() == nil {
		return nil, status.Error(codes.NotFound, "Task not found")
	}

	taskData := task.GetData()
	logrus.Println("taskData ===> ", taskData)
	if taskData.GetWorkflowDoc() == "{}" || taskData.GetWorkflowDoc() == "" {
		logrus.Errorln("[api][func: TaskAction] error workflow empty")
		return nil, status.Error(codes.InvalidArgument, "Bad Request: Workflow is empty")
	}

	var workflow *workflow_pb.ValidateWorkflowData
	err = json.Unmarshal([]byte(taskData.WorkflowDoc), &workflow)
	if err != nil {
		logrus.Errorln("[api][func: TaskAction] error unmarshal workflow: ", err)
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}
	if workflow == nil {
		logrus.Errorln("[api][func: TaskAction] error workflow is nil")
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	var workflowAction workflow_pb.ValidateWorkflowRequest_Action

	switch action := strings.ToLower(req.GetAction()); action {
	case "approve":
		workflowAction = workflow_pb.ValidateWorkflowRequest_APPROVE
		task.Data.LastApprovedByID = currentUser.UserID
		task.Data.LastApprovedByName = currentUser.Username

	case "reject":
		workflowAction = workflow_pb.ValidateWorkflowRequest_REJECT
		task.Data.Comment = req.GetComment()
		task.Data.Reasons = req.GetReasons()
		task.Data.LastRejectedByID = currentUser.UserID
		task.Data.LastRejectedByName = currentUser.Username

	case "rework":
		workflowAction = workflow_pb.ValidateWorkflowRequest_REQUEST_CHANGE
		task.Data.Comment = req.GetComment()
		task.Data.Reasons = req.GetReasons()
		task.Data.LastRejectedByID = currentUser.UserID
		task.Data.LastRejectedByName = currentUser.Username

	case "delete":
		workflowAction = workflow_pb.ValidateWorkflowRequest_REQUEST_DELETE
		task.Data.Comment = req.GetComment()
		task.Data.Reasons = req.GetReasons()

	default:
		return nil, status.Error(codes.InvalidArgument, "Bad Request: Invalid Action")
	}

	send, _ := metadata.FromOutgoingContext(ctx)
	ctx = metadata.NewOutgoingContext(ctx, metadata.Join(send, md))

	logrus.Println("[api][func: TaskAction] action: ", workflowAction)

	validateWorkflow, err := workflowClient.ValidateWorkflow(newCtx, &workflow_pb.ValidateWorkflowRequest{
		CurrentWorkflow: workflow.Workflow,
		Action:          workflowAction,
	}, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		return nil, err
	}
	if !validateWorkflow.IsValid {
		return nil, status.Error(codes.OK, validateWorkflow.Message)
	}

	logrus.Println("workflow validate ===> ", validateWorkflow)

	currentStep := validateWorkflow.GetData().GetWorkflow().GetCurrentStep()

	logrus.Println("current step ===> ", currentStep)

	var dataToSave []byte

	logrus.Println("Task Status ===> ", task.GetData().GetStatus())

	if currentStep == "complete" && task.GetData().GetStatus() != task_pb.Statuses_Approved {

		logrus.Println("[api][func: TaskAction] exec BG Issuing Portal Request")

		var issuingData *pb.IssuingData
		err = json.Unmarshal([]byte(taskData.Data), &issuingData)
		if err != nil {
			logrus.Errorln("[api][func: TaskAction] error unmarshal issuing: ", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

		createIssuing, err := s.CreateIssuing(ctx, &pb.CreateIssuingRequest{
			TaskID: task.Data.TaskID,
			Data:   issuingData,
		})
		if err != nil {
			logrus.Errorln("[api][func: TaskAction] failed to transfer data: ", err)
			return nil, err
		}

		issuingData.RegistrationNo = createIssuing.Data.RegistrationNo
		issuingData.ReferenceNo = createIssuing.Data.ReferenceNo

		dataToSave, err = json.Marshal(issuingData)
		if err != nil {
			logrus.Errorln("[api][func: TaskAction] error marshal saved data: ", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

	}

	if len(dataToSave) > 0 {
		taskData.Data = string(dataToSave)
	} else {
		taskData.Data = task.GetData().GetData()
	}

	currentWorkflow, err := json.Marshal(validateWorkflow.GetData())
	if err != nil {
		logrus.Errorln("[api][func: TaskAction] error marshal current workflow: ", err)
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}
	taskData.WorkflowDoc = string(currentWorkflow)

	saveReq := &task_pb.SaveTaskRequest{
		TaskID:            taskData.TaskID,
		Task:              taskData,
		IsDraft:           false,
		TransactionAmount: 0,
	}

	logrus.Println("saveReq ===> ", saveReq)

	savedTask, err := taskClient.SaveTaskWithWorkflow(newCtx, saveReq, grpc.Header(&userMD), grpc.Trailer(&trailer))
	logrus.Println("savedTask ===> ", savedTask)
	if err != nil {
		logrus.Errorln("[api][func: TaskAction] error save task: ", err)
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	res := &pb.TaskActionResponse{
		Error:   false,
		Code:    200,
		Message: "",
		Data: &pb.Task{
			TaskID:             savedTask.Data.TaskID,
			Type:               savedTask.Data.Type,
			Status:             savedTask.Data.Status.String(),
			Step:               savedTask.Data.Step.String(),
			FeatureID:          savedTask.Data.FeatureID,
			LastApprovedByID:   savedTask.Data.LastApprovedByID,
			LastRejectedByID:   savedTask.Data.LastRejectedByID,
			LastApprovedByName: savedTask.Data.LastApprovedByName,
			LastRejectedByName: savedTask.Data.LastRejectedByName,
			CreatedByName:      savedTask.Data.CreatedByName,
			UpdatedByName:      savedTask.Data.UpdatedByName,
			Reasons:            savedTask.Data.Reasons,
			Comment:            savedTask.Data.Comment,
			CreatedAt:          savedTask.Data.CreatedAt,
			UpdatedAt:          savedTask.Data.UpdatedAt,
		},
	}

	return res, nil

}
