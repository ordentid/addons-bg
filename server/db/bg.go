package db

import (
	"context"
	"errors"

	pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type QueryBuilder struct {
	Filter        string
	FilterOr      string
	CollectiveAnd string
	In            string
	// Distinct      string
	// CustomOrder   string
	Sort *pb.Sort
}

func (p *GormProvider) GetCurrency(ctx context.Context, v *ListFilter) (data []*pb.CurrencyORM, err error) {
	query := p.db_main.Model(&pb.CurrencyORM{})
	if v.Data != nil {
		query = query.Where(v.Data)
	}

	query = query.Scopes(FilterScoope(v.Filter), QueryScoop(v.Query))

	if err := query.Find(&data).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Errorln(err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}
	}
	return data, nil
}

func (p *GormProvider) GetMapping(ctx context.Context, v *ListFilter) (data []*pb.MappingORM, err error) {
	query := p.db_main.Model(&pb.MappingORM{})
	if v.Data != nil {
		query = query.Where(v.Data)
	}

	query = query.Scopes(FilterScoope(v.Filter), QueryScoop(v.Query))

	if err := query.Find(&data).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Errorln(err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}
	}
	return data, nil
}

func (p *GormProvider) GetMappingDetail(ctx context.Context, v *pb.MappingORM) (data *pb.MappingORM, err error) {
	query := p.db_main

	query = query.Model(&pb.MappingORM{}).Where(v)

	if err := query.Find(&data).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Errorln(err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}
	}
	return data, nil
}

func (p *GormProvider) DeleteMapping(ctx context.Context, ids []string) error {
	if len(ids) > 0 {
		logrus.Println("----------------------")
		logrus.Println("Deleted Mapping Data:")
		logrus.Println(ids)
		logrus.Println("----------------------")
		if err := p.db_main.Where("\"id\" IN (?)", ids).Delete(&pb.MappingORM{}).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return status.Error(codes.NotFound, "ID Not Found")
			} else {
				return status.Error(codes.Internal, "Internal Error : "+err.Error())
			}
		}
	}
	return nil
}

func (p *GormProvider) UpdateOrCreateMapping(ctx context.Context, data *pb.MappingORM) (*pb.MappingORM, error) {
	if data.Id > 0 {
		model := &pb.MappingORM{
			Id: data.Id,
		}
		if err := p.db_main.Model(&model).Updates(&data).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, status.Error(codes.NotFound, "ID Not Found")
			} else {
				return nil, status.Error(codes.Internal, "Internal Error : "+err.Error())
			}
		}

		return model, nil
	} else {
		if err := p.db_main.Create(&data).Error; err != nil {
			return nil, status.Error(codes.Internal, "Internal Error : "+err.Error())
		}

		return data, nil
	}
}
