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
	pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/pb"
	"github.com/sirupsen/logrus"
)

const apiServicePath string = "/announcement.service.v1.ApiService/"

// Server represents the server implementation of the SW API.
type Server struct {
	provider *db.GormProvider
	manager  *manager.JWTManager

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
) *Server {
	secret := jwt_secret
	tokenDuration, err := time.ParseDuration(jwt_duration)
	if err != nil {
		logrus.Panic(err)
	}

	return &Server{
		provider:         db.NewProvider(db01),
		manager:          manager.NewJWTManager(secret, tokenDuration),
		ApiServiceServer: nil,
	}
}

func (s *Server) GetManager() *manager.JWTManager {
	return s.manager
}

func (s *Server) notImplementedError() error {
	st := status.New(codes.Unimplemented, "Not implemented yet")
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

func setPagination(page int32, limit int32) *pb.PaginationResponse {
	res := &pb.PaginationResponse{
		Limit: 10,
		Page:  1,
	}

	if limit == 0 && page == 0 {
		res.Limit = -1
		res.Page = -1
		return res
	} else {
		res.Limit = limit
		res.Page = page
	}

	if res.Page == 0 {
		res.Page = 1
	}

	switch {
	case res.Limit > 100:
		res.Limit = 100
	case res.Limit <= 0:
		res.Limit = 10
	}

	return res
}
