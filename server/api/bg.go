package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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
	ThirdPartyName    string  `json:"thirdPartyName"`
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
	ResponseCode    string                `json:"responseCode"`
	ResponseMessage string                `json:"responseMessage"`
	ResponseData    *ApiInquiryThirdParty `json:"responseData"`
}

type ApiInquiryThirdParty struct {
	ThirdPartyID uint64 `json:"thirdPartyId,string"`
	Cif          string `json:"cif"`
	FullName     string `json:"fullName"`
	Status       string `json:"status"`
}

type ApiDownloadRequest struct {
	ReferenceNo string `json:"referenceNo"`
}

type ApiDownloadResponse struct {
	ResponseCode    string      `json:"responseCode"`
	ResponseMessage string      `json:"responseMessage"`
	ResponseData    []UrlObject `json:"responseData"`
}

type UrlObject struct {
	Url string `json:"url"`
}

func (s *Server) GetApplicantName(ctx context.Context, req *pb.GetApplicantNameRequest) (*pb.GetApplicantNameResponse, error) {
	result := &pb.GetApplicantNameResponse{
		Error:   false,
		Code:    200,
		Message: "List Data",
		Data:    []*pb.GetApplicantNameData{},
	}

	applicantNameList, err := s.provider.GetApplicantName(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	result.Data = applicantNameList

	return result, nil
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

	httpReqData := ApiListTransactionRequest{
		Page:  req.Page,
		Limit: req.Limit,
	}

	httpReqParam, err := query.Values(httpReqData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	client := &http.Client{}
	if getEnv("ENV", "PRODUCTION") != "PRODUCTION" {
		proxyURL, err := url.Parse("http://localhost:5002")
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
	}

	httpReq, err := http.NewRequest("GET", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0/listTransaction?"+httpReqParam.Encode(), nil)
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
					httpReqParamsOpt := ApiListTransactionRequest{
						ThirdPartyId: d,
						Page:         1,
						Limit:        1,
					}

					httpReqParams, err := query.Values(httpReqParamsOpt)
					if err != nil {
						return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
					}

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

					err = json.Unmarshal(httpResBody, &httpResData)
					if err != nil {
						return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
					}

					if httpResData.ResponseCode != "00" {
						logrus.Error("Failed To Transfer Data : ", httpResData.ResponseMessage)
						return nil, status.Errorf(codes.Internal, "Internal Error: %v", httpResData.ResponseMessage)
					}

					data := &pb.ThirdPartyORM{ThirdPartyID: id, Name: fmt.Sprintf("THIRD PARTY %s", d)}
					if httpResData.ResponseCode == "00" {
						data.Name = httpResData.ResponseData[0].ThirdPartyName
					}

					_, err = s.provider.UpdateOrCreateThirdParty(ctx, data)
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

	client := &http.Client{}
	if getEnv("ENV", "PRODUCTION") != "PRODUCTION" {
		proxyURL, err := url.Parse("http://localhost:5002")
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
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

		httpReqData := ApiDownloadRequest{
			ReferenceNo: transaction.ReferenceNo,
		}

		httpReqPayload, err := json.Marshal(httpReqData)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		httpReq, err := http.NewRequest("POST", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0/downloadDigitalDocument", bytes.NewBuffer(httpReqPayload))
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		httpReq.Header.Add("Content-Type", "application/json")
		httpReq.Header.Add("Authorization", "Basic YnJpY2FtczpCcmljYW1zNGRkMG5z")

		httpRes, err := client.Do(httpReq)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}
		defer httpRes.Body.Close()

		var httpResData ApiDownloadResponse
		err = json.NewDecoder(httpRes.Body).Decode(&httpResData)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		logrus.Println(httpResData.ResponseCode)

		transaction.DocumentPath = ""
		if httpResData.ResponseCode == "00" {
			if len(httpResData.ResponseData) > 0 {
				transaction.DocumentPath = httpResData.ResponseData[0].Url
			}
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

		client := &http.Client{}
		if getEnv("ENV", "PRODUCTION") != "PRODUCTION" {
			proxyURL, err := url.Parse("http://localhost:5002")
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
			}

			client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
		}

		httpReqData := ApiDownloadRequest{
			ReferenceNo: data.ReferenceNo,
		}

		httpReqPayload, err := json.Marshal(httpReqData)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		httpReq, err := http.NewRequest("POST", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0/downloadDigitalDocument", bytes.NewBuffer(httpReqPayload))
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		httpReq.Header.Add("Content-Type", "application/json")
		httpReq.Header.Add("Authorization", "Basic YnJpY2FtczpCcmljYW1zNGRkMG5z")

		httpRes, err := client.Do(httpReq)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}
		defer httpRes.Body.Close()

		var httpResData ApiDownloadResponse
		err = json.NewDecoder(httpRes.Body).Decode(&httpResData)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		data.DocumentPath = ""
		if httpResData.ResponseCode == "00" {
			if len(httpResData.ResponseData) > 0 {
				data.DocumentPath = httpResData.ResponseData[0].Url
			}
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

	for _, v := range req.Data {
		// transactionORM, err := s.provider.GetTransactionDetail(ctx, &pb.TransactionORM{ReferenceNo: v.GetReferenceNo()})
		// if err != nil {
		// 	return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		// }

		data := &pb.TransactionORM{
			Amount:          v.GetAmount(),
			ApplicantName:   v.GetApplicantName(),
			BeneficiaryName: v.GetBeneficiaryName(),
			BgStatus:        int32(v.GetBgStatus()),
			BgType:          int32(v.GetBgType()),
			ChannelID:       v.GetChannelID(),
			ChannelName:     v.GetChannelName(),
			ClaimPeriod:     v.GetClaimPeriod(),
			ClosingDate:     v.GetClosingDate(),
			CompanyID:       v.GetCompanyID(),
			CreatedByID:     me.UserID,
			Currency:        v.GetCurrency(),
			DocumentPath:    v.GetDocumentPath(),
			EffectiveDate:   v.GetEffectiveDate(),
			ExpiryDate:      v.GetExpiryDate(),
			IssueDate:       v.GetIssueDate(),
			ReferenceNo:     v.GetReferenceNo(),
			RegistrationNo:  v.GetRegistrationNo(),
			Remark:          v.GetRemark(),
			Status:          int32(v.GetStatus()),
			ThirdPartyID:    v.GetThirdPartyID(),
			ThirdPartyName:  v.GetThirdPartyName(),
			TransactionID:   v.GetTransactionID(),
			UpdatedByID:     me.UserID,
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
