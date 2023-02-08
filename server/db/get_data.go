package db

import (
	"context"
	"errors"

	pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ListFilter struct {
	Data   interface{}
	Filter string
	Query  string
}

func (p *GormProvider) GetList(ctx context.Context, data interface{}, v *ListFilter, pagination *pb.PaginationResponse, sort *pb.Sort) (interface{}, error) {
	query := p.db_main
	if v.Data != nil {
		query = query.Preload(clause.Associations).Where(v.Data)
	}

	query = query.Scopes(FilterScoope(v.Filter), QueryScoop(v.Query))
	query = query.Scopes(Paginate(data, pagination, query), Sort(sort))

	if err := query.Find(&data).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorln(err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}
	}
	return data, nil
}

func (p *GormProvider) GetFirst(ctx context.Context, data interface{}) (interface{}, error) {
	query := p.db_main
	var err error
	if err = query.First(&data).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorln(err)
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}
	}
	return data, err
}
