package api

import (
	"context"
	"os"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"bitbucket.bri.co.id/scm/addons/addons-bg-service/server/db"
	manager "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/jwt"
	svc "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs"
	pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/pb"
	"github.com/sirupsen/logrus"
)

const apiServicePath string = "/bg.service.v1.ApiService/"

var (
	log                        *logrus.Logger
	transferTransactionDataTag string
)

// Server represents the server implementation of the SW API.
type Server struct {
	provider *db.GormProvider
	manager  *manager.JWTManager
	svcConn  *svc.ServiceConnection
	logger   *logrus.Logger

	pb.ApiServiceServer
}

type DataPublish struct {
	DataType string
	Data     string
}

func New(
	jwt_secret string,
	jwt_duration string,
	db01 *gorm.DB,
	svcConn *svc.ServiceConnection,
	logger *logrus.Logger,
) *Server {
	log = logger
	secret := jwt_secret
	tokenDuration, err := time.ParseDuration(jwt_duration)
	if err != nil {
		logrus.Panic(err)
	}

	return &Server{
		provider:         db.NewProvider(db01, log),
		manager:          manager.NewJWTManager(secret, tokenDuration, svcConn),
		svcConn:          svcConn,
		logger:           logger,
		ApiServiceServer: nil,
	}
}

func (s *Server) GetManager() *manager.JWTManager {
	return s.manager
}

func (s *Server) NotImplementedError() error {
	st := status.New(codes.Unimplemented, "Not implemented yet")
	return st.Err()
}

func (s *Server) UnauthorizedError() error {
	st := status.New(codes.Unauthenticated, "Unauthorized")
	return st.Err()
}

func (s *Server) HealthCheck(ctx context.Context, _ *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{Message: "API Running !"}, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
