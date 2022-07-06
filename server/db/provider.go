package db

import (
	"context"

	pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/pb"
	"gorm.io/gorm"
)

type Provider interface {
	GetList(ctx context.Context, data interface{}, v *ListFilter, pagination *pb.PaginationResponse, sort *pb.Sort) (interface{}, error)
	GetFirst(ctx context.Context, filter interface{}, data interface{}) (interface{}, error)
}

type GormProvider struct {
	db_main *gorm.DB
}

func NewProvider(db *gorm.DB) *GormProvider {
	return &GormProvider{db_main: db}
}
