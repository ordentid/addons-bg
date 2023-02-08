package api

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	account_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/account"
	company_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/company"
	menu_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/menu"
	system_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/system"
	task_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/task"
	transaction_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/transaction"
	workflow_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/workflow"
	"bitbucket.bri.co.id/scm/addons/addons-bg-service/server/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

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
		return nil, s.UnauthorizedError()
	}
	var trailer metadata.MD

	taskClient := s.svcConn.TaskServiceClient()
	companyClient := s.svcConn.CompanyServiceClient()

	statuses, filters, err := s.FilterBuilder(ctx, *currentUser)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Authority Denied")
	}

	req.Filter = strings.Join([]string{strings.Join(filters, ","), req.GetFilter()}, ",")

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
		CompanyID: currentUser.CompanyID,
	}

	if req.Status.Number() > 0 {
		filter.Status = task_pb.Statuses(req.Status.Number())
	}
	if req.Step.Number() > 0 {
		filter.Step = task_pb.Steps(req.Step.Number())
	}

	listTaskReq := &task_pb.ListTaskRequest{
		Task:         filter,
		Limit:        req.GetLimit(),
		Page:         req.GetPage(),
		Sort:         req.GetSort(),
		Dir:          task_pb.ListTaskRequestDirection(req.GetDir()),
		Filter:       req.GetFilter(),
		Query:        req.GetQuery(),
		FilterOr:     req.GetFilterOr(),
		CustomOrder:  customOrder,
		RoleIDFilter: currentUser.RoleIDs,
		UserIDFilter: currentUser.UserID,
		Services:     "BG Issuing",
	}

	listTaskRes, err := taskClient.GetListTask(newCtx, listTaskReq, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		log.Errorln("[api][func: GetTaskInternalTransfer] Unable to Get List Task:", err.Error())
		return nil, err
	}

	for _, v := range listTaskRes.Data {

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
		err = json.Unmarshal([]byte(v.Data), &taskData)
		if err != nil {
			log.Errorln("[api][func: GetTaskInternalTransfer] Unable to Unmarshal Task Data:", err.Error())
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

		company := &pb.Company{}

		companyRes, err := companyClient.ListCompanyDataV2(newCtx, &company_pb.ListCompanyDataReq{CompanyID: v.CompanyID}, grpc.Header(&userMD), grpc.Trailer(&trailer))
		if err != nil {
			log.Errorln("[api][func: GetTaskInternalTransfer] Unable to Get List Company:", err.Error())
			return nil, err
		}

		workflow := pb.ValidateWorkflowData{}
		err = json.Unmarshal([]byte(v.GetWorkflowDoc()), &workflow)
		if err != nil {
			log.Errorln("[api][func: GetTaskInternalTransfer] Unable to Unmarshal Workflow Data:", err.Error())
			return nil, status.Errorf(codes.Internal, "Internal Error")
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

		}

		result.Data = append(result.Data, &pb.TaskIssuingData{
			Task:     task,
			Company:  company,
			Data:     &taskData,
			Workflow: &workflow,
		})

	}

	result.Pagination = &pb.PaginationResponse{
		Limit:      listTaskRes.GetPagination().GetLimit(),
		Page:       listTaskRes.GetPagination().GetPage(),
		TotalRows:  listTaskRes.GetPagination().GetTotalRows(),
		TotalPages: listTaskRes.GetPagination().GetTotalPages(),
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
		return nil, s.UnauthorizedError()
	}

	var trailer metadata.MD

	taskClient := s.svcConn.TaskServiceClient()
	companyClient := s.svcConn.CompanyServiceClient()

	taskListRes, err := s.GetTaskIssuing(ctx, &pb.GetTaskIssuingRequest{
		Filter: fmt.Sprintf("task_id:%d", req.TaskID),
	})
	if err != nil {
		log.Errorln("[api][func: GetTaskIssuingDetail] Unable to Get Task Issuing:", err.Error())
		return nil, err
	}

	if len(taskListRes.GetData()) < 1 {
		return nil, status.Errorf(codes.NotFound, "Task Not Found")
	}

	taskRes, err := taskClient.GetTaskByID(newCtx, &task_pb.GetTaskByIDReq{ID: req.TaskID, Type: "BG Issuing"}, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		return nil, err
	}

	if taskRes.Data.CompanyID != currentUser.CompanyID {
		return nil, s.UnauthorizedError()
	}

	workflow := pb.ValidateWorkflowData{}
	err = json.Unmarshal([]byte(taskRes.Data.GetWorkflowDoc()), &workflow)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error")
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
	err = json.Unmarshal([]byte(taskRes.Data.GetData()), &taskData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	var company *pb.Company

	companyRes, err := companyClient.ListCompanyDataV2(newCtx, &company_pb.ListCompanyDataReq{CompanyID: task.GetCompanyID()}, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		return nil, err
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

	if err := req.Validate(); err != nil {
		log.Errorln("[api][func: CreateTaskIssuing] Bad Request:", err)
		return nil, status.Error(codes.InvalidArgument, "Invalid Argument")
	}

	if req.Data.Applicant.ApplicantType == 1 {
		if len(req.Data.Applicant.Nik) != 16 {
			return nil, status.Error(codes.InvalidArgument, "Bad Request: NIK is less than 16 characters")
		}
	}

	if req.Data.Publishing.LawArticle != "Pasal 1832" {
		log.Errorln("[api][func: CreateTaskIssuing] Bad Request: Law Article is not 'Pasal 1832'")
		return nil, status.Errorf(codes.InvalidArgument, "Bad Request: Law Article is not 'Pasal 1832'")
	}

	if req.Data.Publishing.PublishingType == pb.PublishingType_SingleBranch {
		if req.Data.Publishing.PublishingBranchId != req.Data.Publishing.OpeningBranchId {
			log.Errorln("[api][func: CreateTaskIssuing] Bad Request: Publishing and Opening Branch should match if Publishing Type is Single Branch")
			return nil, status.Errorf(codes.InvalidArgument, "Bad Request: Publishing and Opening Branch should match if Publishing Type is Single Branch")
		}
	}

	currentUser, userMD, err := s.manager.GetMeFromMD(ctx)
	if err != nil {
		return nil, err
	}
	if currentUser == nil || currentUser.UserType != "cu" {
		return nil, s.UnauthorizedError()
	}
	var trailer metadata.MD

	taskClient := s.svcConn.TaskServiceClient()
	companyClient := s.svcConn.CompanyServiceClient()
	systemClient := s.svcConn.SystemServiceClient()
	transactionClient := s.svcConn.TransactionServiceClient()
	workflowClient := s.svcConn.WorkflowServiceClient()
	menuClient := s.svcConn.MenuServiceClient()
	accountClient := s.svcConn.AccountServiceClient()

	// check user have access to BG Issuing on menu license
	menuMe, err := menuClient.GetMyMenu(newCtx, &menu_pb.GetMyMenuReq{}, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		return nil, err
	}
	isOn := false
	for _, menu1 := range menuMe.Data {
		if len(menu1.Menus) > 0 {
			for _, menu2 := range menu1.Menus {
				if menu2.ProductName == "BG Issuing" {
					isOn = true
				}
			}
		}
	}
	if !isOn {
		return nil, status.Error(codes.PermissionDenied, "Permission Denied")
	}

	// get OTP Validation
	if !req.IsDraft {
		if currentUser.IdToken != "" {
			if req.PassCode == "" {
				return nil, status.Error(codes.InvalidArgument, "Invalid Argument")
			}
			tokenValidRes, err := transactionClient.BRIGateHardTokenValidation(newCtx, &transaction_pb.BRIGateHardTokenValidationRequest{
				UserName: currentUser.IdToken,
				PassCode: req.PassCode,
			})
			if err != nil {
				log.Errorln("[api][func: CreateTaskIssuing] Failed when validate OTP", err.Error())
				return nil, err
			}
			if tokenValidRes.Data.ResponseCode != "00" {
				log.Errorln("[api][func: CreateTaskIssuing] Failed when validate OTP", tokenValidRes.Data.ResponseMessage)
				return nil, status.Error(codes.Aborted, "Hard Token Validation Fail")
			}
		}
	}

	company, err := companyClient.ListCompanyDataV2(newCtx, &company_pb.ListCompanyDataReq{CompanyID: currentUser.CompanyID}, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		log.Errorln("[api][func: CreateTaskIssuing] Failed when execute ListCompanyDataV2:", err.Error())
		return nil, err
	}

	if !(len(company.GetData()) > 0) {
		log.Errorln("[api][func: CreateTaskIssuing] Company not found")
		return nil, status.Errorf(codes.NotFound, "Company not found")
	}

	data := req.Data

	openingBranchORMs, err := systemClient.ListMdBranch(newCtx, &system_pb.ListMdBranchRequest{
		Data: &system_pb.MdBranch{
			Id: req.Data.Publishing.GetOpeningBranchId(),
		},
	}, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		log.Errorln("[api][func: CreateTaskIssuing] Failed when execute ListMdBranch:", err.Error())
		return nil, err
	}

	if len(openingBranchORMs.Data) == 0 {
		log.Errorln("[api][func: CreateTaskIssuing] Opening Branch not found")
		return nil, status.Errorf(codes.NotFound, "Opening Branch not found")
	}

	openingBranch := openingBranchORMs.Data[0]

	publishingBranchORMs, err := systemClient.ListMdBranch(newCtx, &system_pb.ListMdBranchRequest{
		Data: &system_pb.MdBranch{
			Id: req.Data.Publishing.GetPublishingBranchId(),
		},
	}, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		log.Errorln("[api][func: CreateTaskIssuing] Failed when execute ListMdBranch:", err.Error())
		return nil, err
	}

	if len(publishingBranchORMs.Data) == 0 {
		return nil, status.Errorf(codes.NotFound, "Publishing Branch not found")
	}

	publishingBranch := publishingBranchORMs.Data[0]

	data.Publishing.OpeningBranchName = openingBranch.GetDescription()
	data.Publishing.PublishingBranchName = publishingBranch.GetDescription()

	accountRes, err := accountClient.ListAccount(newCtx, &account_pb.ListAccountRequest{
		Account: &account_pb.Account{
			AccountNumber: data.GetAccount().GetAccountNumber(),
		},
	})
	if err != nil {
		log.Errorln("[api][func: CreateTaskIssuing] Failed when execute ListAccount:", err.Error())
		return nil, err
	}

	if len(accountRes.GetData()) < 1 {
		return nil, status.Errorf(codes.NotFound, "Account not found")
	}

	taskData, err := json.Marshal(data)
	if err != nil {
		log.Errorln("[api][func: CreateTaskIssuing] Unable to Marshal Data:", err.Error())
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	taskReq := &task_pb.SaveTaskRequest{
		TaskID: req.TaskID,
		Task: &task_pb.Task{
			Type:        "BG Issuing",
			Data:        string(taskData),
			CreatedByID: currentUser.UserID,
			CompanyID:   currentUser.CompanyID,
		},
		TransactionAmount:   req.Data.Project.BgAmount,
		TransactionCurrency: req.Data.Project.BgCurrency,
		CompanyID:           currentUser.CompanyID,
		HoldingID:           currentUser.CompanyID,
		SelectedAccountID:   accountRes.GetData()[0].GetAccountID(),
	}

	if req.IsDraft {
		taskReq.IsDraft = true
	}

	taskRes, err := taskClient.SaveTaskWithData(newCtx, taskReq, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		log.Errorln("[api][func: CreateTaskIssuing] Failed when execute SaveTaskWithData:", err.Error())
		return nil, err
	}

	updateTaskData := req.Data
	updateTaskData.TransactionID = "BGI" + strconv.FormatUint(taskRes.Data.TaskID, 10)

	updateData, err := json.Marshal(updateTaskData)
	if err != nil {
		log.Errorln("[api][func: CreateTaskIssuing] Unable to Marshal BG Issuing Data:", err)
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	log.Println("[api][func: CreateTaskIssuing] Task Data:", string(updateData))

	_, err = taskClient.UpdateTaskData(newCtx, &task_pb.UpdateTaskDataReq{
		Type:   "BG Issuing",
		TaskID: taskRes.Data.TaskID,
		Data:   string(updateData),
	}, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		log.Errorln("[api][func: CreateTaskIssuing] Unable to Update Task Data:", err)
		return nil, err
	}

	if taskRes.Data.Status != task_pb.Statuses_Draft {

		go func(ctx context.Context) error {

			log.Println("[api][func: CreateTaskIssuing] Auto Approve Task If Company Workflow is STP: START")

			var newCtx context.Context

			md, ok := metadata.FromIncomingContext(ctx)
			if ok {
				newCtx = metadata.NewOutgoingContext(context.Background(), md)
			}

			currentUser, userMD, err := s.manager.GetMeFromMD(ctx)
			if err != nil {
				return err
			}
			if currentUser == nil {
				return s.UnauthorizedError()
			}
			var trailer metadata.MD

			var workflow *workflow_pb.ValidateWorkflowData
			err = json.Unmarshal([]byte(taskRes.Data.WorkflowDoc), &workflow)
			if err != nil {
				log.Errorln("[api][func: CreateTaskIssuing] Unable to Unmarshal Data:", err)
				return status.Errorf(codes.Internal, "Internal Error")
			}
			nextStep := workflow.Workflow.CurrentStep

			companyWorkflow, err := workflowClient.GetCompanyWorkflow(newCtx, &workflow_pb.GetCompanyWorkflowRequest{
				CompanyID: currentUser.CompanyID,
			})
			if err != nil {
				log.Println("[api][func: CreateTaskIssuing] Failed when execute GetCompanyWorkflow function:", err.Error())
				return err
			}

			log.Println("[api][func: CreateTaskIssuing] Workflow STP is:", companyWorkflow.Data.IsTransactionSTP)

			if companyWorkflow.Data.IsTransactionSTP {

				_, err := taskClient.SetTask(newCtx, &task_pb.SetTaskRequest{
					TaskID:  taskRes.GetData().GetTaskID(),
					Action:  "approve",
					Comment: "",
					Reasons: "",
				}, grpc.Header(&userMD), grpc.Trailer(&trailer))
				if err != nil {
					log.Println("[api][func: CreateTaskIssuing] Failed when execute SetTask function:", err.Error())
					return err
				}

				createIssuing, err := s.CreateIssuing(ctx, &pb.CreateIssuingRequest{
					TaskID: taskRes.Data.TaskID,
					Data:   data,
				})
				if err != nil {
					log.Println("[api][func: CreateTaskIssuing] Failed when execute CreateIssuing function:", err.Error())
					return err
				}

				data.RegistrationNo = createIssuing.Data.RegistrationNo
				data.ReferenceNo = createIssuing.Data.ReferenceNo

				taskData, err = json.Marshal(data)
				if err != nil {
					log.Errorln("[api][func: CreateTaskIssuing] Unable to Marshal Data:", err)
					return status.Errorf(codes.Internal, "Internal Error")
				}

				_, err = taskClient.UpdateTaskData(newCtx, &task_pb.UpdateTaskDataReq{
					Type:   "BG Issuing",
					TaskID: taskRes.Data.TaskID,
					Data:   string(taskData),
				}, grpc.Header(&userMD), grpc.Trailer(&trailer))
				if err != nil {
					log.Println("[api][func: CreateTaskIssuing] Failed when execute UpdateTaskData function:", err.Error())
					return err
				}

				nextStep = " "

			}

			taskByIDRes, err := taskClient.GetTaskByID(newCtx, &task_pb.GetTaskByIDReq{
				Type: "BG Issuing",
				ID:   taskRes.Data.TaskID,
			})
			if err != nil {
				log.Errorln("[api][func: CreateTaskIssuing] Unable to Get Task By ID:", err)
				return status.Errorf(codes.Internal, "Internal Error")
			}

			_, err = s.provider.CreateBgTask(ctx, &pb.BgTaskORM{
				TaskID:             taskByIDRes.GetData().GetTaskID(),
				TransactionID:      fmt.Sprintf("BGI%v", taskByIDRes.GetData().GetTaskID()),
				Status:             int32(taskByIDRes.GetData().GetStatus()),
				Step:               int32(taskByIDRes.GetData().GetStep()),
				CreatedByID:        taskByIDRes.GetData().GetCreatedByID(),
				LastApprovedByID:   taskByIDRes.GetData().GetLastApprovedByID(),
				Data:               taskByIDRes.GetData().GetData(),
				Comment:            taskByIDRes.GetData().GetComment(),
				Reasons:            taskByIDRes.GetData().GetReasons(),
				LastApprovedByName: taskByIDRes.GetData().GetLastApprovedByName(),
				CreatedByName:      taskByIDRes.GetData().GetCreatedByName(),
				DataBak:            taskByIDRes.GetData().GetDataBak(),
				WorkflowDoc:        taskByIDRes.GetData().GetWorkflowDoc(),
				CompanyID:          taskByIDRes.GetData().GetCompanyID(),
				HoldingID:          taskByIDRes.GetData().GetHoldingID(),
			})
			if err != nil {
				log.Errorln("[api][func: CreateTaskIssuing] Failed when execute CreateBgTask:", err.Error())
				return status.Errorf(codes.Internal, "Internal Error")
			}

			log.Println("[api][func: CreateTaskIssuing] Auto Approve Task If Company Workflow is STP: END")

			log.Println("[api][func: CreateTaskIssuing] Send for Approval Notification: START")

			notificationClient := s.svcConn.NotificationServiceClient()

			sendNotificationPayload, err := s.NotificationRequestBuilder(ctx, nextStep, taskRes.GetData(), "send approval", currentUser.Username, []string{})
			if err != nil {
				log.Errorln("[api][func: CreateTaskIssuing] Failed when execute NotificationRequestbuilder:", err.Error())
				return status.Errorf(codes.Internal, "Internal Error")
			}

			_, err = notificationClient.SendNotificationWorkflow(newCtx, sendNotificationPayload)
			if err != nil {
				log.Errorln("[api][func: CreateTaskIssuing] Unable to Send Notification:", err.Error())
			}

			log.Println("[api][func: CreateTaskIssuing] Send for Approval Notification: DONE")

			return nil

		}(ctx)

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

func (s *Server) TaskIssuingAction(ctx context.Context, req *pb.TaskIssuingActionRequest) (*pb.TaskIssuingActionResponse, error) {

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
		return nil, s.UnauthorizedError()
	}
	var trailer metadata.MD

	taskClient := s.svcConn.TaskServiceClient()
	workflowClient := s.svcConn.WorkflowServiceClient()

	task, err := taskClient.GetTaskByID(newCtx, &task_pb.GetTaskByIDReq{Type: "BG Issuing", ID: req.GetTaskID()}, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		log.Errorln("[api][func: SetTaskInternalTransfer] Failed when execute GetTaskByID function:", err.Error())
		return nil, err
	}

	if task.GetData() == nil {
		return nil, status.Error(codes.NotFound, "Task not found")
	}

	if task.GetData().GetStatus() == task_pb.Statuses_Draft {

		if req.GetAction() == "delete" {

			action, err := taskClient.SetTask(newCtx, &task_pb.SetTaskRequest{
				TaskID:  req.GetTaskID(),
				Action:  req.GetAction(),
				Comment: req.GetComment(),
				Reasons: req.GetReasons(),
			}, grpc.Header(&userMD), grpc.Trailer(&trailer))
			if err != nil {
				log.Errorln("[api][func: SetTaskInternalTransfer] Unable to Set Task:", err.Error())
				return nil, err
			}

			res := &pb.TaskIssuingActionResponse{
				Error:   false,
				Code:    200,
				Message: "Task Deleted",
				Data: &pb.Task{
					TaskID:             action.Data.TaskID,
					Type:               action.Data.Type,
					Status:             action.Data.Status.String(),
					Step:               action.Data.Step.String(),
					FeatureID:          action.Data.FeatureID,
					LastApprovedByID:   action.Data.LastApprovedByID,
					LastRejectedByID:   action.Data.LastRejectedByID,
					LastApprovedByName: action.Data.LastApprovedByName,
					LastRejectedByName: action.Data.LastRejectedByName,
					CreatedByName:      action.Data.CreatedByName,
					UpdatedByName:      action.Data.UpdatedByName,
					Reasons:            action.Data.Reasons,
					Comment:            action.Data.Comment,
					CreatedAt:          action.Data.CreatedAt,
					UpdatedAt:          action.Data.UpdatedAt,
				},
			}

			return res, nil

		}

	}

	taskData := task.GetData()
	if taskData.GetWorkflowDoc() == "{}" || taskData.GetWorkflowDoc() == "" {
		log.Errorln("[api][func: SetTaskInternalTransfer] Failed: Workflow is empty")
		return nil, status.Error(codes.InvalidArgument, "Bad Request: Workflow is empty")
	}

	var workflow *workflow_pb.ValidateWorkflowData
	err = json.Unmarshal([]byte(taskData.WorkflowDoc), &workflow)
	if err != nil {
		log.Errorln("[api][func: SetTaskInternalTransfer] Unable to Unmarshal Data:", err)
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	if workflow == nil {
		log.Errorln("[api][func: SetTaskInternalTransfer] Failed: Workflow is nil")
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	var workflowAction workflow_pb.ValidateWorkflowRequest_Action

	switch action := strings.ToLower(req.GetAction()); action {
	case "approve":
		workflowAction = workflow_pb.ValidateWorkflowRequest_APPROVE
		task.Data.LastApprovedByID = currentUser.UserID
		task.Data.LastApprovedByName = currentUser.Username
		task.Data.LastRejectedByID = 0
		task.Data.LastRejectedByName = ""

	case "reject":
		workflowAction = workflow_pb.ValidateWorkflowRequest_REJECT
		if req.GetComment() != "" {
			if re.MatchString(req.GetComment()) {
				log.Errorln("[api][func: SetTaskInternalTransfer] Invalid Reject Comment Characters:", req.GetComment())
				return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument")
			} else {
				task.Data.Comment = req.GetComment()
			}
		} else {
			task.Data.Comment = "-"
		}
		task.Data.Reasons = req.GetReasons()
		task.Data.LastApprovedByID = 0
		task.Data.LastApprovedByName = ""
		task.Data.LastRejectedByID = currentUser.UserID
		task.Data.LastRejectedByName = currentUser.Username

	case "rework":
		workflowAction = workflow_pb.ValidateWorkflowRequest_REQUEST_CHANGE
		if req.GetComment() != "" {
			if re.MatchString(req.GetComment()) {
				log.Errorln("[api][func: SetTaskInternalTransfer] Invalid Rework Comment Characters:", req.GetComment())
				return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument")
			} else {
				task.Data.Comment = req.GetComment()
			}
		} else {
			task.Data.Comment = "-"
		}
		task.Data.Reasons = req.GetReasons()
		task.Data.LastApprovedByID = 0
		task.Data.LastApprovedByName = ""
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

	log.Println("[api][func: TaskAction] Workflow Action:", workflowAction)

	validateWorkflow, err := workflowClient.ValidateWorkflow(newCtx, &workflow_pb.ValidateWorkflowRequest{
		CurrentWorkflow: workflow.Workflow,
		Action:          workflowAction,
	}, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		log.Errorln("[api][func: TaskAction] Failed when execute ValidateWorkflow function:", err.Error())
		return nil, err
	}

	if !validateWorkflow.IsValid {
		return nil, status.Error(codes.OK, validateWorkflow.Message)
	}

	workflowResByte, err := json.Marshal(validateWorkflow.Data)
	if err != nil {
		log.Errorln("[api][func: TaskAction] Unable to Marshal Data:", err)
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	log.Println("[api][func: TaskAction] Validate Workflow:", string(workflowResByte))

	currentStep := validateWorkflow.GetData().GetWorkflow().GetCurrentStep()

	log.Println("[api][func: TaskAction] Current Step:", currentStep)

	var dataToSave []byte

	var issuingData *pb.IssuingData
	err = json.Unmarshal([]byte(taskData.Data), &issuingData)
	if err != nil {
		log.Errorln("[api][func: TaskAction] Unable to Unmarshal Data:", err)
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	log.Errorln("[api][func: TaskAction] Task Status:", task.GetData().GetStatus())

	if currentStep == "complete" && task.GetData().GetStatus() != task_pb.Statuses_Approved {

		log.Println("[api][func: TaskAction] Workflow Complete, Execute CreateIssuing...")

		createIssuing, err := s.CreateIssuing(ctx, &pb.CreateIssuingRequest{
			TaskID: task.Data.TaskID,
			Data:   issuingData,
		})
		if err != nil {
			log.Errorln("[api][func: TaskAction] Failed when execute CreateIssuing function:", err.Error())
			return nil, err
		}

		issuingData.RegistrationNo = createIssuing.Data.RegistrationNo
		issuingData.ReferenceNo = createIssuing.Data.ReferenceNo

		dataToSave, err = json.Marshal(issuingData)
		if err != nil {
			log.Errorln("[api][func: TaskAction] Unable to Marshal Data:", err)
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
		log.Errorln("[api][func: TaskAction] Unable to Marshal Data:", err)
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}
	taskData.WorkflowDoc = string(currentWorkflow)

	saveReq := &task_pb.SaveTaskRequest{
		TaskID:            taskData.TaskID,
		Task:              taskData,
		IsDraft:           false,
		TransactionAmount: 0,
	}

	savedTask, err := taskClient.SaveTaskWithWorkflow(newCtx, saveReq, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		log.Errorln("[api][func: TaskAction] Failed when execute SaveTaskWithWorkflow function:", err)
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	taskByIDRes, err := taskClient.GetTaskByID(newCtx, &task_pb.GetTaskByIDReq{
		Type: "BG Issuing",
		ID:   req.GetTaskID(),
	})
	if err != nil {
		log.Errorln("[api][func: TaskAction] Unable to Get Task By ID:", err)
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	_, err = s.provider.UpdateBgTask(ctx, taskByIDRes.GetData().GetTaskID(), &pb.BgTaskORM{
		TransactionID:      fmt.Sprintf("BGI%v", taskByIDRes.GetData().GetTaskID()),
		Status:             int32(taskByIDRes.GetData().GetStatus()),
		Step:               int32(taskByIDRes.GetData().GetStep()),
		CreatedByID:        taskByIDRes.GetData().GetCreatedByID(),
		LastApprovedByID:   taskByIDRes.GetData().GetLastApprovedByID(),
		LastRejectedByID:   taskByIDRes.GetData().GetLastRejectedByID(),
		Data:               taskByIDRes.GetData().GetData(),
		Comment:            taskByIDRes.GetData().GetComment(),
		Reasons:            taskByIDRes.GetData().GetReasons(),
		LastApprovedByName: taskByIDRes.GetData().GetLastApprovedByName(),
		LastRejectedByName: taskByIDRes.GetData().GetLastRejectedByName(),
		CreatedByName:      taskByIDRes.GetData().GetCreatedByName(),
		UpdatedByName:      taskByIDRes.GetData().GetUpdatedByName(),
		DataBak:            taskByIDRes.GetData().GetDataBak(),
		WorkflowDoc:        taskByIDRes.GetData().GetWorkflowDoc(),
		CompanyID:          taskByIDRes.GetData().GetCompanyID(),
		HoldingID:          taskByIDRes.GetData().GetHoldingID(),
	})
	if err != nil {
		log.Errorln("[api][func: CreateTaskIssuing] Failed when execute CreateBgTask:", err.Error())
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	go func() error {
		log.Println("[api][func: TaskAction] Task Action Notification: START")

		notificationClient := s.svcConn.NotificationServiceClient()

		// replace notificationAction from req.action to "complete" if currentStep is "complete"
		notificationAction := strings.ToLower(req.GetAction())
		if currentStep == "complete" {
			notificationAction = "complete"
		}

		sendNotificationPayload, err := s.NotificationRequestBuilder(ctx, currentStep, savedTask.GetData(), notificationAction, currentUser.Username, []string{})
		if err != nil {
			log.Errorln("[api][func: TaskAction] Failed when execute NotificationRequestbuilder:", err.Error())
			return status.Errorf(codes.Internal, "Internal Error")
		}

		_, err = notificationClient.SendNotificationWorkflow(newCtx, sendNotificationPayload)
		if err != nil {
			log.Errorln("[api][func: TaskAction] Unable to Send Notification", err.Error())
		}

		log.Println("[api][func: TaskAction] Task Action Notification: DONE")

		return nil

	}()

	res := &pb.TaskIssuingActionResponse{
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
