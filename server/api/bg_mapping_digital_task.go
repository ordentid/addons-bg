package api

import (
	"context"
	"encoding/json"
	"strings"

	company_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/company"
	task_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/task"
	"bitbucket.bri.co.id/scm/addons/addons-bg-service/server/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

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
		return nil, s.UnauthorizedError()
	}
	var trailer metadata.MD

	taskClient := s.svcConn.TaskServiceClient()
	companyClient := s.svcConn.CompanyServiceClient()

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
		return nil, s.UnauthorizedError()
	}
	var trailer metadata.MD

	taskClient := s.svcConn.TaskServiceClient()
	companyClient := s.svcConn.CompanyServiceClient()

	taskRes, err := taskClient.GetTaskByID(newCtx, &task_pb.GetTaskByIDReq{ID: req.TaskID, Type: "BG Mapping Digital"}, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	if taskRes.Data.CompanyID != currentUser.CompanyID {
		return nil, s.UnauthorizedError()
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
		return nil, s.UnauthorizedError()
	}
	var trailer metadata.MD

	taskClient := s.svcConn.TaskServiceClient()
	companyClient := s.svcConn.CompanyServiceClient()

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
