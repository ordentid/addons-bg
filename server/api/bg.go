package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"bitbucket.bri.co.id/scm/addons/addons-bg-service/server/db"
	pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/pb"
	"github.com/google/go-querystring/query"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ApiListTransactionRequest struct {
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
	ResponseCode    string                `json:"responseCode"`
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

func (s *Server) GetThirdPartyID(ctx context.Context, req *pb.GetThirdPartyIDRequest) (*pb.GetThirdPartyIDResponse, error) {
	result := &pb.GetThirdPartyIDResponse{
		Error:   false,
		Code:    200,
		Message: "List Data",
		Data:    []*pb.ThirdParty{},
	}

	httpReqParamsOpt := ApiListTransactionRequest{
		Page:  "1",
		Limit: "100",
	}

	httpReqParams, err := query.Values(httpReqParamsOpt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	// proxyURL, err := url.Parse("http://localhost:5002")
	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	// }

	// client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
	client := &http.Client{}

	logrus.Println(httpReqParams.Encode())

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

	logrus.Println(string(httpResBody))

	err = json.Unmarshal(httpResBody, &httpResData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	if httpResData.ResponseCode != "00" {
		logrus.Error("Failed To Transfer Data : ", httpResData.ResponseMessage)
	} else {
		for _, d := range httpResData.ResponseData {
			if d.ThirdPartyId > 0 {
				thirdPartyORM, _ := s.provider.GetThirdPartyDetail(ctx, &pb.ThirdPartyORM{Id: d.ThirdPartyId})
				var thirdPartyData *pb.ThirdPartyORM
				if thirdPartyORM == nil {
					thirdPartyData = &pb.ThirdPartyORM{Name: "-"}
				} else {
					thirdPartyData = &pb.ThirdPartyORM{Id: thirdPartyORM.Id, Name: "-"}
				}
				_, err := s.provider.UpdateOrCreateThirdParty(ctx, thirdPartyData)
				if err != nil {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				}
			}
		}
	}

	thirdPartyORMList, err := s.provider.GetThirdParty(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	for _, v := range thirdPartyORMList {
		thirdParty, err := v.ToPB(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}
		result.Data = append(result.Data, &thirdParty)
	}

	return result, nil
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
	}

	me, err := s.manager.GetMeFromJWT(ctx, "")
	if err != nil {
		return nil, err
	}

	data := &pb.TransactionORM{
		Amount:             req.Data.Amount,
		ApplicantName:      req.Data.ApplicantName,
		BeneficiaryName:    req.Data.BeneficiaryName,
		ChannelID:          req.Data.ChannelID,
		ChannelName:        req.Data.ChannelName,
		ClaimPeriod:        req.Data.ClaimPeriod,
		ClosingDate:        req.Data.ClosingDate,
		CompanyID:          req.Data.CompanyID,
		CreatedByID:        me.UserID,
		Currency:           req.Data.Currency,
		DocumentPath:       req.Data.DocumentPath,
		EffectiveDate:      req.Data.EffectiveDate,
		ExpiryDate:         req.Data.ExpiryDate,
		IsAllowBeneficiary: req.Data.IsAllowBeneficiary,
		IssueDate:          req.Data.IssueDate,
		ReferenceNo:        req.Data.ReferenceNo,
		Remark:             req.Data.Remark,
		Status:             "Pending",
		ThirdPartyID:       req.Data.ThirdPartyID,
		TransactionID:      req.Data.TransactionID,
		TransactionStatus:  req.Data.Status,
		TransactionTypeID:  req.Data.TransactionTypeID,
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

	result.Data = &transaction

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
