package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"bitbucket.bri.co.id/scm/addons/addons-bg-service/server/db"
	company_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/company"
	pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/pb"
	"github.com/google/go-querystring/query"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type ApiListTransactionRequest struct {
	StartDate             string `url:"startDate"`
	EndDate               string `url:"endDate"`
	Branch                string `url:"branch"`
	ApplicantName         string `url:"applicationName"`
	ClaimPeriod           string `url:"claimPeriod"`
	Status                string `url:"status"`
	ReferenceNo           string `url:"referenceNo"`
	EventPeriod           string `url:"eventPeriod"`
	BeneficiaryName       string `url:"beneficiaryName"`
	ThirdPartyId          string `url:"thirdPartyId"`
	ChannelId             string `url:"channel_id"`
	ApplicationCustomerId string `url:"applicant_customer_id"`
	BeneficiaryCustomerId string `url:"beneficiary_customer_id"`
	Page                  string `url:"page"`
	Limit                 string `url:"limit"`
}

type ApiListTransactionResponse struct {
	ResponseCode    uint64                `json:"responseCode,string"`
	ResponseMessage string                `json:"responseMessage"`
	Pagination      ApiPaginationResponse `json:"pagination"`
	ResponseData    []*ApiTransaction     `json:"responseData"`
}

type ApiPaginationResponse struct {
	Page        uint64 `json:"page,string"`
	Limit       uint64 `json:"limit,string"`
	TotalRecord uint64 `json:"totalRecord,string"`
	TotalPage   uint32 `json:"totalPage"`
}

type ApiTransaction struct {
	TransactionId     uint64  `json:"transactionId,string"`
	ThirdPartyId      uint64  `json:"thirdPartyId,string"`
	ReferenceNo       string  `json:"referenceNo"`
	ApplicantName     string  `json:"applicantName"`
	BeneficiaryName   string  `json:"beneficiaryName"`
	IssueDate         string  `json:"issueDate"`
	EffectiveDate     string  `json:"effectiveDate"`
	ExpiryDate        string  `json:"expiryDate"`
	ClaimPeriod       uint64  `json:"claimPeriod,string"`
	ClosingDate       string  `json:"closingDate"`
	Currency          string  `json:"currency"`
	Amount            float64 `json:"amount,string"`
	CreatedDate       string  `json:"createdDate"`
	ModifiedDate      string  `json:"modifiedDate"`
	Remark            string  `json:"remark"`
	Status            string  `json:"status"`
	ChannelId         uint64  `json:"channelId,string"`
	ChannelName       string  `json:"channelName"`
	TransactionTypeId uint64  `json:"transactionTypeId,string"`
	DocumentPath      string  `json:"documentPath"`
}

func (s *Server) GetTransaction(ctx context.Context, req *pb.GetTransactionRequest) (*pb.GetTransactionResponse, error) {
	result := &pb.GetTransactionResponse{
		Error:   false,
		Code:    200,
		Message: "List Data",
		Data:    []*pb.Transaction{},
	}

	var data pb.TransactionORM
	var err error
	if req.Transaction != nil {
		data, err = req.Transaction.ToORM(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}
	}

	result.Pagination = setPagination(req.Page, req.Limit)
	sort := &pb.Sort{
		Column:    req.GetSort(),
		Direction: req.GetDir().Enum().String(),
	}

	filter := &db.ListFilter{
		Data:   &data,
		Filter: req.Filter,
		Query:  req.Query,
	}

	listORM, err := s.provider.GetTransaction(ctx, filter, result.Pagination, sort)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	list := []*pb.Transaction{}
	for _, v := range listORM {
		transaction, err := v.ToPB(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}
		list = append(list, &transaction)
	}

	result.Data = list

	return result, nil
}

func (s *Server) GetTransactionDetail(ctx context.Context, req *pb.GetTransactionDetailRequest) (*pb.GetTransactionDetailResponse, error) {
	result := &pb.GetTransactionDetailResponse{
		Error:   false,
		Code:    200,
		Message: "Data",
	}

	if req.TransactionID > 0 {
		orm, err := s.provider.GetTransactionDetail(ctx, &pb.TransactionORM{TransactionID: req.TransactionID})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		data, err := orm.ToPB(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		result.Data = &data
	}

	return result, nil
}

func (s *Server) CreateTransaction(ctx context.Context, req *pb.CreateTransactionRequest) (*pb.CreateTransactionResponse, error) {
	result := &pb.CreateTransactionResponse{
		Error:   false,
		Code:    200,
		Message: "Data",
		Data:    []*pb.Transaction{},
	}

	me, err := s.manager.GetMeFromJWT(ctx, "")
	if err != nil {
		return nil, err
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		ctx = metadata.NewOutgoingContext(context.Background(), md)
	}
	// var header, trailer metadata.MD

	companyConn, err := grpc.Dial(getEnv("COMPANY_SERVICE", ":9092"), opts...)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed connect to Company Service: %v", err)
	}
	defer companyConn.Close()

	companyClient := company_pb.NewApiServiceClient(companyConn)

	company, err := companyClient.ListCompanyData(ctx, &company_pb.ListCompanyDataReq{CompanyID: req.CompanyID})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}
	if !(len(company.GetData()) > 0) {
		return nil, status.Errorf(codes.NotFound, "Company not found.")
	}

	// taskConn, err := grpc.Dial(getEnv("TASK_SERVICE", ":9090"), opts...)
	// if err != nil {
	// 	logrus.Errorln("Failed connect to Task Service: %v", err)
	// 	return nil, status.Errorf(codes.Internal, "Error Internal")
	// }
	// defer taskConn.Close()

	// taskClient := task_pb.NewTaskServiceClient(taskConn)

	for _, v := range req.ThirdParty {
		httpReqParamsOpt := ApiListTransactionRequest{
			ThirdPartyId: strconv.FormatUint(v.ThirdPartyID, 10),
			Page:         "1",
			Limit:        "10",
		}

		httpReqParams, err := query.Values(httpReqParamsOpt)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		proxyURL, err := url.Parse("http://localhost:5002")
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
		httpReq, err := http.NewRequest("GET", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0/listTransaction?"+httpReqParams.Encode(), nil)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		httpReq.Header.Add("Authorization", "Basic YnJpY2FtczpCcmljYW1zNGRkMG5z")

		httpRes, err := client.Do(httpReq)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}
		defer httpRes.Body.Close()

		var httpResData ApiListTransactionResponse
		httpResBody, err := ioutil.ReadAll(httpRes.Body)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		err = json.Unmarshal(httpResBody, &httpResData)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		if httpResData.ResponseCode != 00 {
			logrus.Error("Failed To Transfer Data : ", httpResData.ResponseMessage)
		} else {
			for _, d := range httpResData.ResponseData {
				data := &pb.TransactionORM{
					Amount:             d.Amount,
					ApplicantName:      d.ApplicantName,
					BeneficiaryName:    d.BeneficiaryName,
					ChannelID:          d.ChannelId,
					ChannelName:        d.ChannelName,
					ClaimPeriod:        d.ClaimPeriod,
					ClosingDate:        d.ClosingDate,
					CompanyID:          req.CompanyID,
					CreatedByID:        me.UserID,
					Currency:           d.Currency,
					DocumentPath:       d.DocumentPath,
					EffectiveDate:      d.EffectiveDate,
					ExpiryDate:         d.ExpiryDate,
					IsAllowBeneficiary: v.IsAllowBeneficiary,
					IssueDate:          d.IssueDate,
					ReferenceNo:        d.ReferenceNo,
					Remark:             d.Remark,
					Status:             "Pending",
					ThirdPartyID:       d.ThirdPartyId,
					TransactionID:      d.TransactionId,
					TransactionStatus:  d.Status,
					TransactionTypeID:  d.TransactionTypeId,
					UpdatedByID:        me.UserID,
				}

				transactionORM, err := s.provider.UpdateOrCreateTransaction(ctx, data)
				if err != nil {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				}

				transaction, err := transactionORM.ToPB(ctx)
				if err != nil {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				}

				result.Data = append(result.Data, &transaction)
			}
		}
	}

	return result, nil
}

func (s *Server) UpdateTransaction(ctx context.Context, req *pb.UpdateTransactionRequest) (*pb.UpdateTransactionResponse, error) {
	result := &pb.UpdateTransactionResponse{
		Error:   false,
		Code:    200,
		Message: "Data",
	}

	return result, nil
}
