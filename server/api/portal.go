package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ApiTransaction struct {
	TransactionId     uint64  `json:"transaction_id,string"`
	ThirdPartyId      uint64  `json:"third_party_id,string"`
	ThirdPartyName    string  `json:"third_party_name"`
	ReferenceNo       string  `json:"reference_no"`
	RegistrationNo    string  `json:"registration_no"`
	ApplicantName     string  `json:"applicant_name"`
	BeneficiaryId     uint64  `json:"beneficiary_id,string"`
	BeneficiaryName   string  `json:"beneficiary_name"`
	IssueDate         string  `json:"issue_date"`
	EffectiveDate     string  `json:"effective_date"`
	ExpiryDate        string  `json:"expiry_date"`
	ClaimPeriod       uint32  `json:"claim_period,string"`
	ClosingDate       string  `json:"closing_date"`
	Currency          string  `json:"currency"`
	Amount            float64 `json:"amount,string"`
	CreatedDate       string  `json:"created_date"`
	ModifiedDate      string  `json:"modified_date"`
	Remark            string  `json:"remark"`
	Status            string  `json:"status"`
	ChannelId         uint64  `json:"channel_id,string"`
	ChannelName       string  `json:"channel_name"`
	TransactionTypeId uint64  `json:"transaction_type_id,string"`
	DocumentPath      string  `json:"document_path"`
}

type ApiInquiryThirdParty struct {
	ThirdPartyID uint64 `json:"third_party_id,string"`
	Cif          string `json:"cif"`
	FullName     string `json:"fullname"`
	Status       string `json:"status"`
}

type ApiBeneficiary struct {
	BeneficiaryID uint64 `json:"beneficiary_id,string"`
	ThirdPartyID  uint64 `json:"third_party_id,string"`
	Cif           string `json:"cif"`
	FullName      string `json:"fullname"`
	CreatedDate   string `json:"created_date"`
	ModifiedDate  string `json:"modified_date"`
	Status        string `json:"status"`
}

type UrlObject struct {
	Url string `json:"url"`
}

type ApiPaginationResponse struct {
	Page        uint64 `json:"page,string"`
	Limit       uint64 `json:"limit,string"`
	TotalRecord uint64 `json:"total_record,string"`
	TotalPage   uint32 `json:"total_page"`
}

type ApiListTransactionRequest struct {
	StartDate             string `url:"start_date"`
	EndDate               string `url:"end_date"`
	Branch                string `url:"branch"`
	ApplicantName         string `url:"applicant_name"`
	ClaimPeriod           string `url:"claim_period"`
	Status                string `url:"status"`
	ReferenceNo           string `url:"reference_no"`
	EventPeriod           string `url:"event_period"`
	BeneficiaryId         string `url:"beneficiary_id,string"`
	BeneficiaryName       string `url:"beneficiary_name"`
	ThirdPartyId          uint64 `url:"third_party_id,string"`
	ThirdPartyName        string `url:"third_party_name"`
	ChannelId             uint64 `url:"channel_id"`
	ChannelName           string `url:"channel_name"`
	ApplicationCustomerId string `url:"applicant_customer_id"`
	BeneficiaryCustomerId string `url:"beneficiary_customer_id"`
	Page                  uint64 `url:"page,string"`
	Limit                 uint64 `url:"limit,string"`
}

type ApiListTransactionResponse struct {
	ResponseCode    string                `json:"code"`
	ResponseMessage string                `json:"message"`
	Pagination      ApiPaginationResponse `json:"pagination"`
	ResponseData    []*ApiTransaction     `json:"data"`
}

type ApiInquiryThirdPartyByIDRequest struct {
	ThirdPartyID uint64 `json:"thirdPartyId,string"`
}

type ApiInquiryThirdPartyByIDResponse struct {
	ResponseCode    string                `json:"code"`
	ResponseMessage string                `json:"message"`
	ResponseData    *ApiInquiryThirdParty `json:"data"`
}

type ApiInquiryThirdPartyByStatusRequest struct {
	Status string `url:"status"`
}

type ApiInquiryThirdPartyByStatusResponse struct {
	ResponseCode    string                  `json:"code"`
	ResponseMessage string                  `json:"message"`
	ResponseData    []*ApiInquiryThirdParty `json:"data"`
}

type ApiDownloadRequest struct {
	ReferenceNo string `json:"reference_no"`
}

type ApiDownloadResponse struct {
	ResponseCode    string      `json:"code"`
	ResponseMessage string      `json:"message"`
	ResponseData    []UrlObject `json:"data"`
}

type ApiInquiryBenficiaryRequest struct {
	Cif          string `url:"cif"`
	Fullname     string `url:"fullname"`
	ThirdPartyID uint64 `url:"third_party_id"`
}

type ApiInquiryBenficiaryResponse struct {
	ResponseCode    string            `json:"code"`
	ResponseMessage string            `json:"message"`
	ResponseData    []*ApiBeneficiary `json:"data"`
}

func GetHttpClient(ctx context.Context) (*http.Client, error) {
	client := &http.Client{}
	if getEnv("ENV", "PRODUCTION") != "PRODUCTION" {
		proxyURL, err := url.Parse("http://localhost:5100")
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
	}
	return client, nil
}

func ApiInquiryBeneficiary(ctx context.Context, req *ApiInquiryBenficiaryRequest) (*ApiInquiryBenficiaryResponse, error) {
	client, err := GetHttpClient(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	httpReqParam, err := query.Values(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	httpReq, err := http.NewRequest("GET", getEnv("PORTAL_BG_URL", "https://tfapi.dev.bri.co.id/portalbg-api")+"/partnership/beneficiary?"+httpReqParam.Encode(), nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	httpReq.Header.Add("Authorization", "Basic "+getEnv("PORTAL_BG_API_KEY", ""))

	httpRes, err := client.Do(httpReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}
	defer httpRes.Body.Close()

	var httpResData ApiInquiryBenficiaryResponse
	err = json.NewDecoder(httpRes.Body).Decode(&httpResData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	return &httpResData, nil
}

func ApiInquiryThirdPartyByStatus(ctx context.Context, req *ApiInquiryThirdPartyByStatusRequest) (*ApiInquiryThirdPartyByStatusResponse, error) {
	client, err := GetHttpClient(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	httpReqParam, err := query.Values(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	httpReq, err := http.NewRequest("GET", getEnv("PORTAL_BG_URL", "https://tfapi.dev.bri.co.id/portalbg-api")+"/partnership/inquiry"+httpReqParam.Encode(), nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", "Basic "+getEnv("PORTAL_BG_API_KEY", ""))

	httpRes, err := client.Do(httpReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}
	defer httpRes.Body.Close()

	var httpResData ApiInquiryThirdPartyByStatusResponse
	err = json.NewDecoder(httpRes.Body).Decode(&httpResData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	return &httpResData, nil
}

func ApiInquiryThirdPartyByID(ctx context.Context, req *ApiInquiryThirdPartyByIDRequest) (*ApiInquiryThirdPartyByIDResponse, error) {
	client, err := GetHttpClient(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	httpReq, err := http.NewRequest("GET", getEnv("PORTAL_BG_URL", "https://tfapi.dev.bri.co.id/portalbg-api")+"/partnership/inquiry/"+strconv.FormatUint(req.ThirdPartyID, 10), nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", "Basic "+getEnv("PORTAL_BG_API_KEY", ""))

	httpRes, err := client.Do(httpReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}
	defer httpRes.Body.Close()

	var httpResData ApiInquiryThirdPartyByIDResponse
	err = json.NewDecoder(httpRes.Body).Decode(&httpResData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	return &httpResData, nil
}

func ApiDownload(ctx context.Context, req *ApiDownloadRequest) (*ApiDownloadResponse, error) {
	client, err := GetHttpClient(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	httpReqPayload := strings.NewReader("reference_no=" + req.ReferenceNo)

	httpReq, err := http.NewRequest("POST", getEnv("PORTAL_BG_URL", "https://tfapi.dev.bri.co.id/portalbg-api")+"/digitaldocument/byreference", httpReqPayload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	httpReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	httpReq.Header.Add("Authorization", "Basic "+getEnv("PORTAL_BG_API_KEY", ""))

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

	return &httpResData, nil
}

func ApiListTransaction(ctx context.Context, req *ApiListTransactionRequest) (*ApiListTransactionResponse, error) {
	client, err := GetHttpClient(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	httpReqParams, err := query.Values(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	logrus.Println("---------------------------")
	logrus.Println(httpReqParams.Encode())
	logrus.Println("---------------------------")

	httpReq, err := http.NewRequest("GET", getEnv("PORTAL_BG_URL", "https://tfapi.dev.bri.co.id/portalbg-api")+"/channel/monitoring/transaction?"+httpReqParams.Encode(), nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	httpReq.Header.Add("Authorization", "Basic "+getEnv("PORTAL_BG_API_KEY", ""))

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

	return &httpResData, nil
}
