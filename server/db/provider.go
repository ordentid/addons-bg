package db

import (
	"context"
	"os"

	pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/pb"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var log *logrus.Logger
var transferTransactionDataTag string

type Provider interface {
	GetList(ctx context.Context, data interface{}, v *ListFilter, pagination *pb.PaginationResponse, sort *pb.Sort) (interface{}, error)
	GetFirst(ctx context.Context, filter interface{}, data interface{}) (interface{}, error)
}

type GormProvider struct {
	db_main *gorm.DB
	logger  *logrus.Logger
}

func NewProvider(db *gorm.DB, logger *logrus.Logger) *GormProvider {
	log = logger
	transferTransactionDataTag = getEnv("DATA_TAG1", "addons.bg")
	return &GormProvider{db_main: db}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
