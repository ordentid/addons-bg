package api

import (
	"context"
	"encoding/json"
	"errors"
	"regexp"
	"strings"

	company_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/company"
	system_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/system"
	task_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/task"
	workflow_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/workflow"
	"bitbucket.bri.co.id/scm/addons/addons-bg-service/server/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
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
		return nil, s.unauthorizedError()
	}
	var trailer metadata.MD

	taskClient := s.scvConn.TaskServiceClient()
	companyClient := s.scvConn.CompanyServiceClient()

	logrus.Println("======> Current User: ", currentUser)
	logrus.Println("======> Authorities: ", currentUser.Authorities)

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

		httpResData, err := s.ApiInquiryThirdPartyByID(ctx, &httpReqData)
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
	if currentUser == nil || currentUser.UserType != "ca" {
		return nil, s.unauthorizedError()
	}
	var trailer metadata.MD

	taskClient := s.scvConn.TaskServiceClient()
	companyClient := s.scvConn.CompanyServiceClient()

	logrus.Println("======> Current User: ", currentUser)
	logrus.Println("======> Authorities: ", currentUser.Authorities)

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
	if currentUser == nil || currentUser.UserType != "ca" {
		return nil, s.unauthorizedError()
	}
	var trailer metadata.MD

	taskClient := s.scvConn.TaskServiceClient()
	companyClient := s.scvConn.CompanyServiceClient()

	taskRes, err := taskClient.GetTaskByID(newCtx, &task_pb.GetTaskByIDReq{ID: req.TaskID, Type: "BG Mapping Digital"}, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	if taskRes.Data.CompanyID != currentUser.CompanyID {
		return nil, s.unauthorizedError()
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
	if currentUser == nil || currentUser.UserType != "ca" {
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

	name := ""

	// check third party ID and beneficiary name is valid
	BenefName, err := s.GetBeneficiaryName(ctx, &pb.GetBeneficiaryNameRequest{
		ThirdPartyID: req.ThirdPartyID,
		Type:         pb.BeneficiaryType_AllBeneficiary,
	})
	if err != nil {
		logrus.Errorln("Failed to get beneficiary name")
		return nil, err
	}

	pass := false
	for _, v := range BenefName.Data {
		for _, vv := range req.Beneficiary {
			if vv.BeneficiaryId == v.BeneficiaryId {
				pass = true
				break
			}
		}
	}
	if !pass {
		logrus.Errorln("Invalid request data")
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument")
	}

	httpReqData := ApiInquiryThirdPartyByIDRequest{
		ThirdPartyID: req.ThirdPartyID,
	}

	httpResData, err := s.ApiInquiryThirdPartyByID(ctx, &httpReqData)
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
	if currentUser == nil || currentUser.UserType != "cu" {
		return nil, s.unauthorizedError()
	}
	var trailer metadata.MD

	taskClient := s.scvConn.TaskServiceClient()
	companyClient := s.scvConn.CompanyServiceClient()

	logrus.Println("======> Current User: ", currentUser)
	logrus.Println("======> Authorities: ", currentUser.Authorities)

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
	if currentUser == nil || currentUser.UserType != "cu" {
		return nil, s.unauthorizedError()
	}
	var trailer metadata.MD

	taskClient := s.scvConn.TaskServiceClient()
	companyClient := s.scvConn.CompanyServiceClient()

	taskRes, err := taskClient.GetTaskByID(newCtx, &task_pb.GetTaskByIDReq{ID: req.TaskID, Type: "BG Issuing"}, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	if taskRes.Data.CompanyID != currentUser.CompanyID {
		return nil, s.unauthorizedError()
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

	if req.Data.Applicant.Nik == "" || req.Data.Applicant.Email == "" ||
		req.Data.Applicant.PhoneNumber == "" || req.Data.Applicant.NpwpNo == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument")
	}

	if req.Data.Publishing.LawArticle != "Pasal 1832" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument")
	}
	if req.Data.Publishing.PublishingType.String() == "SingleBranch" {
		if req.Data.Publishing.PublishingBranchId != req.Data.Publishing.OpeningBranchId {
			return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument")
		}
	}

	currentUser, userMD, err := s.manager.GetMeFromMD(ctx)
	if err != nil {
		return nil, err
	}
	if currentUser == nil || currentUser.UserType != "cu" {
		return nil, s.unauthorizedError()
	}
	var trailer metadata.MD

	taskClient := s.scvConn.TaskServiceClient()
	companyClient := s.scvConn.CompanyServiceClient()
	systemClient := s.scvConn.SystemServiceClient()

	company, err := companyClient.ListCompanyDataV2(newCtx, &company_pb.ListCompanyDataReq{CompanyID: currentUser.CompanyID}, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}
	if !(len(company.GetData()) > 0) {
		return nil, status.Errorf(codes.NotFound, "Company not found.")
	}

	data := req.Data

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

	data.Publishing.OpeningBranchName = openingBranch.GetDescription()
	data.Publishing.PublishingBranchName = publishingBranch.GetDescription()

	taskData, err := json.Marshal(data)
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

func (s *Server) TaskAction(ctx context.Context, req *pb.TaskActionRequest) (*pb.TaskActionResponse, error) {

	if req.GetAction() == "" || req.GetTaskID() < 1 {
		return nil, status.Error(codes.InvalidArgument, "Invalid Argument")
	}

	re := regexp.MustCompile(`(<[a-z,\/]+.*?>)`)

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

	task, err := taskClient.GetTaskByID(newCtx, &task_pb.GetTaskByIDReq{
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
		if req.GetComment() != "" {
			if re.MatchString(req.GetComment()) {
				logrus.Errorf(`Error ---> Invalid Reject Comment Characters: %s`, req.GetComment())
				return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument")
			} else {
				task.Data.Comment = req.GetComment()
			}
		} else {
			task.Data.Comment = "-"
		}
		// task.Data.Comment = req.GetComment()
		task.Data.Reasons = req.GetReasons()
		task.Data.LastRejectedByID = currentUser.UserID
		task.Data.LastRejectedByName = currentUser.Username

	case "rework":
		workflowAction = workflow_pb.ValidateWorkflowRequest_REQUEST_CHANGE
		if req.GetComment() != "" {
			if re.MatchString(req.GetComment()) {
				logrus.Errorf(`Error ---> Invalid Rework Comment Characters: %s`, req.GetComment())
				return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument")
			} else {
				task.Data.Comment = req.GetComment()
			}
		} else {
			task.Data.Comment = "-"
		}
		// task.Data.Comment = req.GetComment()
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
