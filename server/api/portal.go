package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/pb"
	"github.com/google/go-querystring/query"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ApiTransaction struct {
	TransactionId     uint64  `json:"transactionId,string"`
	ThirdPartyId      uint64  `json:"thirdPartyId,string"`
	ThirdPartyName    string  `json:"thirdPartyName"`
	ReferenceNo       string  `json:"referenceNo"`
	RegistrationNo    string  `json:"registrationNo"`
	ApplicantName     string  `json:"applicantName"`
	BeneficiaryId     uint64  `json:"beneficiaryId,string"`
	BeneficiaryName   string  `json:"beneficiaryName"`
	IssueDate         string  `json:"issueDate"`
	EffectiveDate     string  `json:"effectiveDate"`
	ExpiryDate        string  `json:"expiryDate"`
	ClaimPeriod       uint32  `json:"claimPeriod,string"`
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

type ApiInquiryThirdParty struct {
	ThirdPartyID uint64 `json:"thirdPartyId,string"`
	Cif          string `json:"cif"`
	FullName     string `json:"fullName"`
	Status       string `json:"status"`
}

type ApiBeneficiary struct {
	BeneficiaryID uint64 `json:"beneficiaryId,string"`
	ThirdPartyID  uint64 `json:"thirdPartyId,string"`
	Cif           string `json:"cif"`
	FullName      string `json:"fullname"`
	CreatedDate   string `json:"createdDate"`
	ModifiedDate  string `json:"modifiedDate"`
	Status        string `json:"status"`
}

type UrlObject struct {
	Url string `json:"url"`
}

type ApiPaginationResponse struct {
	Page        uint64 `json:"page,string"`
	Limit       uint64 `json:"limit,string"`
	TotalRecord uint64 `json:"totalRecord,string"`
	TotalPage   uint32 `json:"totalPage"`
}

type ApiListTransactionRequest struct {
	StartDate             string `url:"startDate"`
	EndDate               string `url:"endDate"`
	Branch                string `url:"branch"`
	ApplicantName         string `url:"applicantName"`
	ClaimPeriod           string `url:"claimPeriod"`
	Status                string `url:"status"`
	ReferenceNo           string `url:"referenceNo"`
	EventPeriod           string `url:"eventPeriod"`
	BeneficiaryId         string `url:"beneficiaryId,string"`
	BeneficiaryName       string `url:"beneficiaryName"`
	ThirdPartyId          uint64 `url:"thirdPartyId,string"`
	ThirdPartyName        string `url:"thirdPartyName"`
	ChannelId             uint64 `url:"channelId"`
	ChannelName           string `url:"channelName"`
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
	ResponseCode    string                `json:"responseCode"`
	ResponseMessage string                `json:"responseMessage"`
	ResponseData    *ApiInquiryThirdParty `json:"responseData"`
}

type ApiInquiryThirdPartyByStatusRequest struct {
	Status string `url:"status"`
}

type ApiInquiryThirdPartyByStatusResponse struct {
	ResponseCode    string                  `json:"responseCode"`
	ResponseMessage string                  `json:"responseMessage"`
	ResponseData    []*ApiInquiryThirdParty `json:"responseData"`
}

type ApiDownloadRequest struct {
	ReferenceNo string `json:"reference_no"`
}

type ApiDownloadResponse struct {
	ResponseCode    string      `json:"responseCode"`
	ResponseMessage string      `json:"responseMessage"`
	ResponseData    []UrlObject `json:"responseData"`
}

type ApiInquiryBenficiaryRequest struct {
	Cif          string `url:"cif"`
	Fullname     string `url:"fullname"`
	ThirdPartyID uint64 `url:"thirdPartyId"`
}

type ApiInquiryBenficiaryResponse struct {
	ResponseCode    string            `json:"responseCode"`
	ResponseMessage string            `json:"responseMessage"`
	ResponseData    []*ApiBeneficiary `json:"responseData"`
}

type ApiBgIssuingRequest struct {
	AccountNo              string  `json:"account_no"`
	ApplicantName          string  `json:"applicant_name"`
	ApplicantAddress       string  `json:"applicant_address"`
	IsIndividu             uint64  `json:"is_individu,string"`
	NIK                    string  `json:"nik"`
	BirthDate              string  `json:"birth_date"`
	Gender                 string  `json:"gender"`
	NPWPNo                 string  `json:"npwp_no"`
	DateEstablished        string  `json:"tanggal_berdiri"`
	CompanyType            string  `json:"company_type"`
	IsPlafond              string  `json:"is_plafond"`
	TransactionType        string  `json:"transaction_type"`
	IsEndOfYearBg          string  `json:"is_bg_akhir_tahun"`
	NRK                    string  `json:"nrk"`
	ProjectName            string  `json:"project_name"`
	ThirdPartyId           uint64  `json:"third_party_id,string"`
	BeneficiaryName        string  `json:"beneficiary_name"`
	ProjectAmount          float64 `json:"project_amount,string"`
	ContractNo             string  `json:"contract_no"`
	ContractDate           string  `json:"contract_date"`
	Currency               string  `json:"currency"`
	Amount                 float64 `json:"amount,string"`
	EffectiveDate          string  `json:"effective_date"`
	MaturityDate           string  `json:"maturity_date"`
	ClaimPeriod            uint64  `json:"claim_periode,string"`
	IssuingBranch          string  `json:"issuing_branch"`
	BranchPrinter          string  `json:"pencetak_branch"`
	ContraGuarantee        string  `json:"contra_guarantee"`
	InsuranceLimitId       string  `json:"insurance_limit_id"`
	SP3No                  string  `json:"sp3_no"`
	HoldAccountNo          string  `json:"hold_account_no"`
	HoldAccountAmount      float64 `json:"hold_account_amount,string"`
	ConsumerLimitId        string  `json:"consumer_limit_id"`
	ConsumerLimitAmount    string  `json:"consumer_limit_amount"`
	ApplicantContactPerson string  `json:"applicant_contact_person"`
	ApplicantPhoneNumber   string  `json:"applicant_phone_number"`
	ApplicantEmail         string  `json:"applicant_email"`
	ChannelId              string  `json:"channel_id"`
	ApplicantCustomerId    string  `json:"applicant_customer_id"`
	BeneficiaryCustomerId  string  `json:"beneficiary_customer_id"`
	LegalDocument          string  `json:"document_legalitas"`
	ContractDocument       string  `json:"document_contract"`
	Sp3Document            string  `json:"document_sp3"`
	OthersDocument         string  `json:"document_others"`
}

type ApiBgIssuingResponse struct {
	ResponseCode    string           `json:"code"`
	ResponseMessage string           `json:"message"`
	Data            ApiBgIssuingData `json:"data"`
}

type ApiBgIssuingData struct {
	RegistrationNo string `json:"registration_no"`
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

	httpReq, err := http.NewRequest("GET", getEnv("PORTAL_BG_URL", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0")+"/inquiryBeneficiary?"+httpReqParam.Encode(), nil)
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

	httpReqPayload, err := json.Marshal(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	httpReq, err := http.NewRequest("POST", getEnv("PORTAL_BG_URL", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0")+"/inquiryThirdParty", bytes.NewBuffer(httpReqPayload))
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

	httpReqPayload, err := json.Marshal(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	httpReq, err := http.NewRequest("POST", getEnv("PORTAL_BG_URL", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0")+"/inquiryThirdParty", bytes.NewBuffer(httpReqPayload))
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

	httpReqPayload, err := json.Marshal(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	httpReq, err := http.NewRequest("POST", getEnv("PORTAL_BG_URL", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0")+"/downloadDigitalDocument", bytes.NewBuffer(httpReqPayload))
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

	httpReq, err := http.NewRequest("GET", getEnv("PORTAL_BG_URL", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0")+"/listTransaction?"+httpReqParams.Encode(), nil)
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

func (s *Server) ApiCreateIssuing(ctx context.Context, req *pb.ApiCreateIssuingRequest) (*pb.ApiCreateIssuingResponse, error) {
	result := &pb.ApiCreateIssuingResponse{
		Error:   false,
		Code:    200,
		Message: "Data",
	}

	httpReqData := ApiBgIssuingRequest{
		AccountNo:              req.Data.Account.GetAccountNumber(),
		ApplicantName:          req.Data.Applicant.GetName(),
		ApplicantAddress:       req.Data.Applicant.GetAddress(),
		IsIndividu:             uint64(req.Data.Applicant.GetApplicantType().Number()),
		NIK:                    "Test",
		BirthDate:              req.Data.Applicant.GetBirthDate(),
		Gender:                 req.Data.Applicant.GetGender().String(),
		NPWPNo:                 "Test",
		DateEstablished:        req.Data.Applicant.GetDateEstablished(),
		CompanyType:            req.Data.Applicant.GetCompanyType().String(),
		IsPlafond:              "0",
		TransactionType:        req.Data.Publishing.GetBgType().String(),
		IsEndOfYearBg:          "0",
		NRK:                    req.Data.Project.GetNrkNumber(),
		ProjectName:            req.Data.Project.GetName(),
		ThirdPartyId:           req.Data.Publishing.GetThirdPartyID(),
		BeneficiaryName:        req.Data.Applicant.GetBeneficiaryName(),
		ProjectAmount:          req.Data.Project.GetProjectAmount(),
		ContractNo:             req.Data.Project.GetContractNumber(),
		ContractDate:           req.Data.Project.GetProjectDate(),
		Currency:               req.Data.Project.GetBgCurrency(),
		Amount:                 req.Data.Project.GetBgAmount(),
		EffectiveDate:          req.Data.Publishing.GetEffectiveDate(),
		MaturityDate:           req.Data.Publishing.GetExpiryDate(),
		ClaimPeriod:            req.Data.Publishing.GetClaimPeriod(),
		IssuingBranch:          req.Data.Publishing.GetOpeningBranch(),
		BranchPrinter:          "Test",
		ContraGuarantee:        "Test",
		InsuranceLimitId:       "Test",
		SP3No:                  "Test",
		HoldAccountNo:          "Test",
		HoldAccountAmount:      0.0,
		ConsumerLimitId:        "Test",
		ConsumerLimitAmount:    "0",
		ApplicantContactPerson: req.Data.Applicant.GetContactPerson(),
		ApplicantPhoneNumber:   "Test",
		ApplicantEmail:         "Test",
		ChannelId:              "Test",
		ApplicantCustomerId:    "Test",
		BeneficiaryCustomerId:  "Test",
		LegalDocument:          req.Data.Document.GetBusinessLegal(),
		ContractDocument:       req.Data.Document.GetBg(),
		Sp3Document:            req.Data.Document.GetSp(),
		OthersDocument:         req.Data.Document.GetOther(),
	}

	httpReqPayload, err := json.Marshal(httpReqData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	httpReq, err := http.NewRequest("POST", getEnv("PORTAL_BG_URL", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0")+"/applyBG", bytes.NewBuffer(httpReqPayload))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", "Basic "+getEnv("PORTAL_BG_API_KEY", ""))

	client := &http.Client{}
	httpRes, err := client.Do(httpReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}
	defer httpRes.Body.Close()

	var httpResData ApiBgIssuingResponse
	err = json.NewDecoder(httpRes.Body).Decode(&httpResData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	logrus.Println(httpResData.ResponseCode)

	if httpResData.ResponseCode == "00" {
		result.Data = &pb.IssuingPortal{
			RegistrationNo: httpResData.Data.RegistrationNo,
		}
	} else {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	return result, nil
}
