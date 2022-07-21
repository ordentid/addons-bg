package db

import (
	"context"
	"errors"

	pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (p *GormProvider) GetApplicantName(ctx context.Context, thirdPartyID uint64) ([]*pb.ApplicantName, error) {
	data := []*pb.ApplicantName{}
	query := p.db_main

	query = query.Model(&pb.TransactionORM{}).Where(&pb.TransactionORM{Status: pb.TransactionStatus_value["MappingDigital"], ThirdPartyID: thirdPartyID})
	query = query.Select(`"applicant_name" as name, count("id") as total`)
	query = query.Group(`"applicant_name"`)

	if err := query.Find(&data).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Errorln(err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}
	}
	return data, nil
}

func (p *GormProvider) GetThirdPartyByCompany(ctx context.Context, companyID uint64) ([]*pb.ThirdPartyName, error) {
	data := []*pb.ThirdPartyName{}
	query := p.db_main

	query = query.Model(&pb.TransactionORM{}).Where(&pb.TransactionORM{Status: pb.TransactionStatus_value["MappingDigital"], CompanyID: companyID})
	query = query.Select(`"third_party_id" as id, "third_party_name" as name, count("id") as total`)
	query = query.Group(`"third_party_id", "third_party_name"`)

	if err := query.Find(&data).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Errorln(err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}
	}
	return data, nil
}

func (p *GormProvider) GetThirdParty(ctx context.Context) ([]*pb.ThirdPartyORM, error) {
	data := []*pb.ThirdPartyORM{}
	query := p.db_main

	if err := query.Find(&data).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Errorln(err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}
	}
	return data, nil
}

func (p *GormProvider) GetThirdPartyDetail(ctx context.Context, data *pb.ThirdPartyORM) (*pb.ThirdPartyORM, error) {
	query := p.db_main
	var err error
	if err = query.First(&data, &data).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Errorln(err)
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}
	}
	return data, err
}

func (p *GormProvider) UpdateOrCreateThirdParty(ctx context.Context, data *pb.ThirdPartyORM) (*pb.ThirdPartyORM, error) {
	if data.Id > 0 {
		model := &pb.ThirdPartyORM{
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

func (p *GormProvider) GetTransaction(ctx context.Context, v *ListFilter, pagination *pb.PaginationResponse, sort *pb.Sort) ([]*pb.TransactionORM, error) {
	data := []*pb.TransactionORM{}
	query := p.db_main
	if v.Data != nil {
		query = query.Preload(clause.Associations).Where(v.Data)
	}

	query = query.Scopes(FilterScoope(v.Filter), QueryScoop(v.Query))
	query = query.Scopes(Paginate(data, pagination, query), Sort(sort))
	query = query.Where("status > 0")

	if err := query.Find(&data).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Errorln(err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}
	}
	return data, nil
}

func (p *GormProvider) GetTransactionDetail(ctx context.Context, data *pb.TransactionORM) (*pb.TransactionORM, error) {
	query := p.db_main
	var err error
	if err = query.First(&data, &data).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Errorln(err)
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}
	}
	return data, err
}

func (p *GormProvider) UpdateOrCreateTransaction(ctx context.Context, data *pb.TransactionORM) (*pb.TransactionORM, error) {
	if data.Id > 0 {
		model := &pb.TransactionORM{
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
