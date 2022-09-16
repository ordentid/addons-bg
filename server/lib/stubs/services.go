package services

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	accountPB "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/account"
	authPB "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/auth"
	companyPB "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/company"
	taskPB "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/task"
	workflowPB "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/workflow"
)

type ServiceConnection struct {
	TaskService        *grpc.ClientConn
	AuthService        *grpc.ClientConn
	UserService        *grpc.ClientConn
	CompanyService     *grpc.ClientConn
	AccountService     *grpc.ClientConn
	SystemService      *grpc.ClientConn
	WorkflowService    *grpc.ClientConn
	TransactionService *grpc.ClientConn
}

func InitServicesConn(
	certFile string,
	taskAddres string,
	authAddress string,
	userAddress string,
	companyAddress string,
	workfloAddress string,
) *ServiceConnection {
	var err error
	var creds credentials.TransportCredentials
	if certFile != "" {
		creds, err = credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			logrus.Panic(err)
		}
	} else {
		creds = insecure.NewCredentials()
	}
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(creds))

	services := &ServiceConnection{}

	// Task Service
	services.TaskService, err = initGrpcClientConn(taskAddres, "Task Service", opts...)
	if err != nil {
		logrus.Fatalf("%v", err)
		os.Exit(1)
		return nil
	}

	// Auth Service
	services.AuthService, err = initGrpcClientConn(authAddress, "Auth Service", opts...)
	if err != nil {
		logrus.Fatalf("%v", err)
		os.Exit(1)
		return nil
	}

	// User Service
	// services.UserService, err = initGrpcClientConn(userAddress, "User Service", opts...)
	// if err != nil {
	// 	logrus.Fatalf("%v", err)
	// 	os.Exit(1)
	// 	return nil
	// }

	// Company Service
	services.CompanyService, err = initGrpcClientConn(companyAddress, "Company Service", opts...)
	if err != nil {
		logrus.Fatalf("%v", err)
		os.Exit(1)
		return nil
	}

	// Account Service
	// services.AccountService, err = initGrpcClientConn(accountAddress, "Account Service", opts...)
	// if err != nil {
	// 	logrus.Fatalf("%v", err)
	// 	os.Exit(1)
	// 	return nil
	// }

	// System Service
	// services.SystemService, err = initGrpcClientConn(systemAddress, "System Service", opts...)
	// if err != nil {
	// 	logrus.Fatalf("%v", err)
	// 	os.Exit(1)
	// 	return nil
	// }

	// Workflow Service
	services.WorkflowService, err = initGrpcClientConn(workfloAddress, "Workflow Service", opts...)
	if err != nil {
		logrus.Fatalf("%v", err)
		os.Exit(1)
		return nil
	}

	// Transaction Service
	// services.TransactionService, err = initGrpcClientConn(transactionAddress, "Transaction Service", opts...)
	// if err != nil {
	// 	logrus.Fatalf("%v", err)
	// 	os.Exit(1)
	// 	return nil
	// }

	return services
}

func initGrpcClientConn(address string, name string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	if address == "" {
		return nil, fmt.Errorf("%s address is empty", name)
	}

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed connect to %s: %v", name, err)
	}

	logrus.Println(fmt.Sprintf("[service - connection] %s State: %s", name, conn.GetState().String()))
	logrus.Println(fmt.Sprintf("[service - connection] %s Connected, on %s", name, address))

	return conn, nil
}

func (s *ServiceConnection) TaskServiceClient() taskPB.TaskServiceClient {
	return taskPB.NewTaskServiceClient(s.TaskService)
}

func (s *ServiceConnection) AuthServiceClient() authPB.ApiServiceClient {
	return authPB.NewApiServiceClient(s.AuthService)
}

func (s *ServiceConnection) CompanyServiceClient() companyPB.ApiServiceClient {
	return companyPB.NewApiServiceClient(s.CompanyService)
}

func (s *ServiceConnection) AccountServiceClient() accountPB.ApiServiceClient {
	return accountPB.NewApiServiceClient(s.AccountService)
}

func (s *ServiceConnection) WorkflowServiceClient() workflowPB.ApiServiceClient {
	return workflowPB.NewApiServiceClient(s.WorkflowService)
}

func (s *ServiceConnection) CloseAllServicesConn() {
	s.TaskService.Close()
	s.AuthService.Close()
	s.UserService.Close()
	s.CompanyService.Close()
	s.AccountService.Close()
	s.SystemService.Close()
	s.WorkflowService.Close()
	s.TransactionService.Close()
}

func (s *ServiceConnection) CloseAuthServiceConn() error {
	return s.AuthService.Close()
}

func (s *ServiceConnection) CloseUserServiceConn() error {
	return s.UserService.Close()
}

func (s *ServiceConnection) CloseTaskServiceConn() error {
	return s.TaskService.Close()
}

func (s *ServiceConnection) CloseCompanyServiceConn() error {
	return s.CompanyService.Close()
}

func (s *ServiceConnection) CloseAccountServiceConn() error {
	return s.AccountService.Close()
}

func (s *ServiceConnection) CloseSystemServiceConn() error {
	return s.SystemService.Close()
}

func (s *ServiceConnection) CloseWorkflowServiceConn() error {
	return s.WorkflowService.Close()
}

func (s *ServiceConnection) CloseTransactionServiceConn() error {
	return s.TransactionService.Close()
}
