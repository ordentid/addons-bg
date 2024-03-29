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
	menuPB "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/menu"
	systemPB "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/system"
	taskPB "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/task"
	transactionPB "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/transaction"
	workflowPB "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/workflow"
	userPB "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/user"
)

type ServiceConnection struct {
	TaskService        *grpc.ClientConn
	AuthService        *grpc.ClientConn
	CompanyService     *grpc.ClientConn
	WorkflowService    *grpc.ClientConn
	SystemService      *grpc.ClientConn
	TransactionService *grpc.ClientConn
	AccountService     *grpc.ClientConn
	MenuService        *grpc.ClientConn
	UserService        *grpc.ClientConn
}

func InitServicesConn(
	certFile string,
	taskAddres string,
	authAddress string,
	companyAddress string,
	workflowAddress string,
	systemAddress string,
	transactionAddress string,
	accountAddress string,
	menuAddress string,
	userAddress string,
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

	// Company Service
	services.CompanyService, err = initGrpcClientConn(companyAddress, "Company Service", opts...)
	if err != nil {
		logrus.Fatalf("%v", err)
		os.Exit(1)
		return nil
	}

	// Workflow Service
	services.WorkflowService, err = initGrpcClientConn(workflowAddress, "Workflow Service", opts...)
	if err != nil {
		logrus.Fatalf("%v", err)
		os.Exit(1)
		return nil
	}

	// System Service
	services.SystemService, err = initGrpcClientConn(systemAddress, "System Service", opts...)
	if err != nil {
		logrus.Fatalf("%v", err)
		os.Exit(1)
		return nil
	}

	// Transaction Service
	services.TransactionService, err = initGrpcClientConn(transactionAddress, "Transaction Service", opts...)
	if err != nil {
		logrus.Fatalf("%v", err)
		os.Exit(1)
		return nil
	}

	// Account Service
	services.AccountService, err = initGrpcClientConn(accountAddress, "Account Service", opts...)
	if err != nil {
		logrus.Fatalf("%v", err)
		os.Exit(1)
		return nil
	}

	// Menu Service
	services.MenuService, err = initGrpcClientConn(menuAddress, "Menu Service", opts...)
	if err != nil {
		logrus.Fatalf("%v", err)
		os.Exit(1)
		return nil
	}

	// User Service
	services.UserService, err = initGrpcClientConn(userAddress, "User Service", opts...)
	if err != nil {
		logrus.Fatalf("%v", err)
		os.Exit(1)
		return nil
	}

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

func (s *ServiceConnection) WorkflowServiceClient() workflowPB.ApiServiceClient {
	return workflowPB.NewApiServiceClient(s.WorkflowService)
}

func (s *ServiceConnection) SystemServiceClient() systemPB.ApiServiceClient {
	return systemPB.NewApiServiceClient(s.SystemService)
}

func (s *ServiceConnection) TransactionServiceClient() transactionPB.TransactionServiceClient {
	return transactionPB.NewTransactionServiceClient(s.TransactionService)
}

func (s *ServiceConnection) AccountServiceClient() accountPB.ApiServiceClient {
	return accountPB.NewApiServiceClient(s.AccountService)
}

func (s *ServiceConnection) MenuServiceClient() menuPB.ApiServiceClient {
	return menuPB.NewApiServiceClient(s.MenuService)
}

func (s *ServiceConnection) UserServiceClient() userPB.ApiServiceClient {
	return userPB.NewApiServiceClient(s.UserService)
}

func (s *ServiceConnection) CloseAllServicesConn() {
	s.TaskService.Close()
	s.AuthService.Close()
	s.CompanyService.Close()
	s.WorkflowService.Close()
	s.SystemService.Close()
	s.TransactionService.Close()
	s.AccountService.Close()
	s.MenuService.Close()
	s.UserService.Close()
}

func (s *ServiceConnection) CloseTaskServiceConn() error {
	return s.TaskService.Close()
}

func (s *ServiceConnection) CloseAuthServiceConn() error {
	return s.AuthService.Close()
}

func (s *ServiceConnection) CloseCompanyServiceConn() error {
	return s.CompanyService.Close()
}

func (s *ServiceConnection) CloseWorkflowServiceConn() error {
	return s.WorkflowService.Close()
}

func (s *ServiceConnection) CloseSystemServiceConn() error {
	return s.SystemService.Close()
}

func (s *ServiceConnection) CloseTransactionServiceConn() error {
	return s.TransactionService.Close()
}

func (s *ServiceConnection) CloseAccountServiceConn() error {
	return s.AccountService.Close()
}

func (s *ServiceConnection) CloseMenuServiceConn() error {
	return s.MenuService.Close()
}

func (s *ServiceConnection) CloseUserServiceConn() error {
	return s.UserService.Close()
}
