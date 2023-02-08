package api

import (
	"context"
	"encoding/json"
	"strings"

	company_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/company"
	task_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/task"
	"bitbucket.bri.co.id/scm/addons/addons-bg-service/server/pb"
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
	if currentUser == nil || currentUser.UserType != "ba" {
		return nil, s.UnauthorizedError()
	}
	var trailer metadata.MD

	taskClient := s.svcConn.TaskServiceClient()
	companyClient := s.svcConn.CompanyServiceClient()

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

		case "checker":
			statuses = []string{"1", "6", "4", "5"}
			if len(req.Filter) > 0 {
				req.Filter = req.Filter + ","
			}
			req.Filter = req.Filter + "status:<>0,status:<>2,status:<>3,status:<>7"

		case "signer":
			statuses = []string{"1", "6", "4", "5"}
			if len(req.Filter) > 0 {
				req.Filter = req.Filter + ","
			}
			req.Filter = req.Filter + "status:<>0,status:<>2,status:<>3,status:<>7"

		case "releaser":
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
	if currentUser == nil || currentUser.UserType != "ba" {
		return nil, s.UnauthorizedError()
	}
	var trailer metadata.MD

	taskClient := s.svcConn.TaskServiceClient()
	companyClient := s.svcConn.CompanyServiceClient()

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

	var newCtx context.Context

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		newCtx = metadata.NewOutgoingContext(context.Background(), md)
	}

	currentUser, userMD, err := s.manager.GetMeFromMD(ctx)
	if err != nil {
		return nil, err
	}
	if currentUser == nil || currentUser.UserType != "ba" {
		return nil, s.UnauthorizedError()
	}
	var trailer metadata.MD

	taskClient := s.svcConn.TaskServiceClient()
	companyClient := s.svcConn.CompanyServiceClient()

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

		httpResData, err := s.ApiInquiryThirdPartyByID(ctx, &httpReqData)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		log.Println(httpResData.ResponseCode)

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
		log.Error("Failed To Marshal : ", taskData)
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
		log.Error("Failed To Transfer Data : ", err)
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
