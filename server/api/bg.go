package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"bitbucket.bri.co.id/scm/addons/addons-bg-service/server/db"
	pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/pb"
	"github.com/google/go-querystring/query"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ApiPaginationResponse struct {
	Page        uint64 `json:"page,string"`
	Limit       uint64 `json:"limit,string"`
	TotalRecord uint64 `json:"totalRecord,string"`
	TotalPage   uint32 `json:"totalPage"`
}

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
	Page                  uint64 `url:"page,string"`
	Limit                 uint64 `url:"limit,string"`
}

type ApiListTransactionResponse struct {
	ResponseCode    string                `json:"responseCode"`
	ResponseMessage string                `json:"responseMessage"`
	Pagination      ApiPaginationResponse `json:"pagination"`
	ResponseData    []*ApiTransaction     `json:"responseData"`
}

type ApiTransaction struct {
	TransactionId     uint64  `json:"transactionId,string"`
	ThirdPartyId      uint64  `json:"thirdPartyId,string"`
	ReferenceNo       string  `json:"referenceNo"`
	RegistrationNo    string  `json:"registrationNo"`
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

type ApiInquiryThirdPartyByIDRequest struct {
	ThirdPartyID uint64 `json:"thirdPartyId,string"`
}

type ApiInquiryThirdPartyByIDResponse struct {
	ResponseCode    uint64                `json:"responseCode,string"`
	ResponseMessage uint64                `json:"responseMessage,string"`
	ResponseData    *ApiInquiryThirdParty `json:"responseData"`
}

type ApiInquiryThirdParty struct {
	ThirdPartyID uint64 `json:"thirdPartyId,string"`
	Cif          string `json:"cif"`
	FullName     string `json:"fullName"`
	Status       string `json:"status"`
}

func (s *Server) GetThirdParty(ctx context.Context, req *pb.GetThirdPartyRequest) (*pb.GetThirdPartyResponse, error) {
	result := &pb.GetThirdPartyResponse{
		Error:   false,
		Code:    200,
		Message: "List Data",
		Data:    []*pb.ThirdParty{},
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

func (s *Server) GenerateThirdParty(ctx context.Context, req *pb.GenerateThirdPartyRequest) (*pb.GenerateThirdPartyResponse, error) {
	result := &pb.GenerateThirdPartyResponse{
		Error:   false,
		Code:    200,
		Message: "Success",
	}

	httpReqParamsOpt := ApiListTransactionRequest{
		Page:  req.Page,
		Limit: req.Limit,
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
	err = json.NewDecoder(httpRes.Body).Decode(&httpResData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	if httpResData.ResponseCode != "00" {
		logrus.Error("Failed To Transfer Data : ", httpResData.ResponseMessage)
	} else {
		idList := []string{}
		for _, d := range httpResData.ResponseData {
			if d.ThirdPartyId > 0 {
				if !contains(idList, strconv.FormatUint(d.ThirdPartyId, 10)) {
					logrus.Println("Third Party ID:", d.ThirdPartyId)
					idList = append(idList, strconv.FormatUint(d.ThirdPartyId, 10))
				}
			}
		}

		logrus.Println(idList)
		for _, d := range idList {
			id, err := strconv.ParseUint(d, 10, 64)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
			}

			_, err = s.provider.GetThirdPartyDetail(ctx, &pb.ThirdPartyORM{ThirdPartyID: id})
			if err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				} else {
					logrus.Println("THIRD PARTY " + d)

					// httpReqBodyData := &ApiInquiryThirdPartyByIDRequest{
					// 	ThirdPartyID: id,
					// }

					// httpReqBodyByte, err := json.Marshal(httpReqBodyData)
					// if err != nil {
					// 	return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
					// }

					// httpReq, err := http.NewRequest("POST", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0/inquiryThirdParty", strings.NewReader(string(httpReqBodyByte)))
					// if err != nil {
					// 	return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
					// }

					// httpReq.Header.Add("Authorization", "Basic YnJpY2FtczpCcmljYW1zNGRkMG5z")

					// httpRes, err := client.Do(httpReq)
					// if err != nil {
					// 	return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
					// }
					// defer httpRes.Body.Close()

					// var httpResData ApiInquiryThirdPartyByIDResponse
					// err = json.NewDecoder(httpRes.Body).Decode(&httpResData)
					// if err != nil {
					// 	return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
					// }

					_, err := s.provider.UpdateOrCreateThirdParty(ctx, &pb.ThirdPartyORM{ThirdPartyID: id, Name: "THIRD PARTY " + d})
					if err != nil {
						return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
					}
				}
			}
		}

		result.Pagination = &pb.ApiPaginationResponse{
			Page:        httpResData.Pagination.Page,
			Limit:       httpResData.Pagination.Limit,
			TotalRecord: httpResData.Pagination.TotalRecord,
			TotalPage:   httpResData.Pagination.TotalPage,
		}
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
		RegistrationNo:     req.Data.RegistrationNo,
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
