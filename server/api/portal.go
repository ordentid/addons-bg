package api

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
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
	ChannelId             uint64 `url:"channel_id"`
	ChannelName           string `url:"channel_name"`
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
	RegistrationNo string `json:"registrationNo"`
}

type ApiDownloadResponse struct {
	ResponseCode    string                  `json:"responseCode"`
	ResponseMessage string                  `json:"responseMessage"`
	ResponseData    ApiDownloadResponseData `json:"responseData"`
}

type ApiDownloadResponseData struct {
	RegistrationNo  string `json:"registrationNo"`
	ReferenceNo     string `json:"referenceNo"`
	WarkatUrl       string `json:"warkatUrl"`
	WarkatUrlPublic string `json:"warkatUrlPublic"`
	Status          string `json:"status"`
	ModifiedDate    string `json:"modifiedDate"`
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
	AccountNo              string            `json:"accountNumber"`
	ApplicantName          string            `json:"applicantName"`
	ApplicantAddress       string            `json:"applicantAddress"`
	IsIndividu             uint64            `json:"isIndividu,string"`
	NIK                    string            `json:"nik"`
	BirthDate              string            `json:"birthDate"`
	Gender                 string            `json:"gender"`
	NPWPNo                 string            `json:"npwp"`
	DateEstablished        string            `json:"tanggalBerdiri"`
	CompanyType            uint64            `json:"companyType,string"`
	IsPlafond              uint64            `json:"isPlafond,string"`
	TransactionType        uint64            `json:"transactionType"`
	TransactionTypeDesc    string            `json:"transactionTypeDesc"`
	IsEndOfYearBg          string            `json:"isBGAkhirTahun"`
	NRK                    string            `json:"nrk"`
	ProjectName            string            `json:"projectName"`
	ThirdPartyId           uint64            `json:"thirdPartyId,string"`
	BeneficiaryName        string            `json:"beneficiaryName"`
	ProjectAmount          float64           `json:"projectAmount,string"`
	ContractNo             string            `json:"contractNo"`
	ContractDate           string            `json:"contractDate"`
	Currency               string            `json:"currency"`
	Amount                 float64           `json:"amount,string"`
	EffectiveDate          string            `json:"effectiveDate"`
	MaturityDate           string            `json:"maturityDate"`
	ClaimPeriod            uint64            `json:"claimPeriod,string"`
	IssuingBranch          string            `json:"issuingBranch"`
	PublishingBranch       string            `json:"pencetakBranch"`
	ContraGuarantee        map[string]string `json:"contraGuarantee"`
	InsuranceLimitId       string            `json:"insuranceLimitId"`
	SP3No                  string            `json:"sp3No"`
	HoldAccountNo          string            `json:"holdAccountNo"`
	HoldAccountAmount      float64           `json:"holdAccountAmount,string"`
	ConsumerLimitId        string            `json:"customerLimitId"`
	ConsumerLimitAmount    float64           `json:"customerLimitAmount,string"`
	ApplicantContactPerson string            `json:"applicantContactPerson"`
	ApplicantPhoneNumber   string            `json:"applicantPhoneNumber"`
	ApplicantEmail         string            `json:"applicantEmail"`
	ChannelId              string            `json:"channelId"`
	ApplicantCustomerId    string            `json:"applicantCustomerId"`
	BeneficiaryCustomerId  string            `json:"beneficiaryCustomerId"`
	LegalDocument          string            `json:"documentLegalitas"`
	ContractDocument       string            `json:"documentContract"`
	Sp3Document            string            `json:"documentSp3"`
	OthersDocument         string            `json:"documentOthers"`
}

type ApiBgIssuingResponse struct {
	ResponseCode    string           `json:"responseCode"`
	ResponseMessage *json.RawMessage `json:"responseMessage"`
	Data            ApiBgIssuingData `json:"responseData"`
}

type ApiBgIssuingData struct {
	RegistrationNo string `json:"registrationNo"`
}

type ApiBgTrackingData struct {
	RegistrationNo  string `json:"registrationNo"`
	ReferenceNo     string `json:"referenceNo"`
	WarkatUrl       string `json:"warkatUrl"`
	WarkatUrlPublic string `json:"warkatUrlPublic"`
	Status          string `json:"status"`
	ModifiedDate    string `json:"modifiedDate"`
}

type ApiBgTrackingRequest struct {
	RegistrationNo string `json:"registrationNo"`
}

type ApiBgTrackingResponse struct {
	ResponseCode    string             `json:"responseCode"`
	ResponseMessage *json.RawMessage   `json:"responseMessage"`
	Data            *ApiBgTrackingData `json:"responseData"`
}

type ApiInquiryLimitIndividualRequest struct {
	Cif string `url:"cif"`
}

type ApiInquiryLimitIndividualResponse struct {
	ResponseCode    string                           `json:"responseCode"`
	ResponseMessage *json.RawMessage                 `json:"responseMessage"`
	ResponseData    []*ApiInquiryLimitIndividualData `json:"responseData"`
}

type ApiInquiryLimitIndividualData struct {
	CustomerLimitId   uint64  `json:"customerLimitId,string"`
	Code              string  `json:"code"`
	Fullname          string  `json:"fullname"`
	Cif               string  `json:"cif"`
	PtkNo             string  `json:"ptkNo"`
	Currency          string  `json:"currency"`
	Plafond           float64 `json:"plafond,string"`
	ReservationAmount int64   `json:"reservationAmount"`
	OutstandingAmount int64   `json:"outstandingAmount"`
	AvailableAmount   int64   `json:"availableAmount"`
	ExpiryDate        string  `json:"expiryDate"`
	PnRm              string  `json:"pnRm"`
	NameRm            string  `json:"nameRm"`
	CreatedDate       string  `json:"createdDate"`
	ModifiedDate      string  `json:"modifiedDate"`
	Status            string  `json:"status"`
}

type UploadEncodeData struct {
	Filename      string `json:"fileName"`
	DocumentPath  string `json:"documentPath"`
	UploadDate    string `json:"uploadDate"`
	UploadFileUrl string `json:"uploadedFileUrl"`
}

type ApiUploadEncodeRequest struct {
	ChannelId string `json:"channelId"`
	Document  string `json:"document"`
}

type ApiUploadEncodeResponse struct {
	ResponseCode    string           `json:"responseCode"`
	ResponseMessage *json.RawMessage `json:"responseMessage"`
	ResponseData    UploadEncodeData `json:"responseData"`
}

func (s *Server) ApiInquiryBeneficiary(ctx context.Context, req *ApiInquiryBenficiaryRequest) (*ApiInquiryBenficiaryResponse, error) {

	var transport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 15 * time.Second,
		}).Dial,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		TLSHandshakeTimeout: 15 * time.Second,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   15 * time.Second,
	}

	httpReqParam, err := query.Values(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("GET", getEnv("PORTAL_BG_URL", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0")+"/inquiryBeneficiary?"+httpReqParam.Encode(), nil)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Add("Authorization", "Basic "+getEnv("PORTAL_BG_API_KEY", ""))

	httpRes, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRes.Body.Close()

	var httpResData ApiInquiryBenficiaryResponse
	err = json.NewDecoder(httpRes.Body).Decode(&httpResData)
	if err != nil {
		return nil, err
	}

	httpResPayload, err := json.Marshal(httpResData)
	if err != nil {
		return nil, err
	}

	log.Println("Response:", string(httpResPayload))

	return &httpResData, nil

}

func (s *Server) ApiInquiryThirdPartyByStatus(ctx context.Context, req *ApiInquiryThirdPartyByStatusRequest) (*ApiInquiryThirdPartyByStatusResponse, error) {

	var transport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 15 * time.Second,
		}).Dial,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		TLSHandshakeTimeout: 15 * time.Second,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   15 * time.Second,
	}

	httpReqPayload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", getEnv("PORTAL_BG_URL", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0")+"/inquiryThirdParty", bytes.NewBuffer(httpReqPayload))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", "Basic "+getEnv("PORTAL_BG_API_KEY", ""))

	httpRes, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRes.Body.Close()

	var httpResData ApiInquiryThirdPartyByStatusResponse
	err = json.NewDecoder(httpRes.Body).Decode(&httpResData)
	if err != nil {
		return nil, err
	}

	httpResPayload, err := json.Marshal(httpResData)
	if err != nil {
		return nil, err
	}

	log.Println("Response:", string(httpResPayload))

	return &httpResData, nil

}

func (s *Server) ApiInquiryThirdPartyByID(ctx context.Context, req *ApiInquiryThirdPartyByIDRequest) (*ApiInquiryThirdPartyByIDResponse, error) {

	var transport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 15 * time.Second,
		}).Dial,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		TLSHandshakeTimeout: 15 * time.Second,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   15 * time.Second,
	}

	httpReqPayload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", getEnv("PORTAL_BG_URL", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0")+"/inquiryThirdParty", bytes.NewBuffer(httpReqPayload))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", "Basic "+getEnv("PORTAL_BG_API_KEY", ""))

	httpRes, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRes.Body.Close()

	var httpResData ApiInquiryThirdPartyByIDResponse
	err = json.NewDecoder(httpRes.Body).Decode(&httpResData)
	if err != nil {
		return nil, err
	}

	httpResPayload, err := json.Marshal(httpResData)
	if err != nil {
		return nil, err
	}

	log.Println("Response:", string(httpResPayload))

	return &httpResData, nil

}

func (s *Server) ApiDownload(ctx context.Context, req *ApiDownloadRequest) (*ApiDownloadResponse, error) {

	var transport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 15 * time.Second,
		}).Dial,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		TLSHandshakeTimeout: 15 * time.Second,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   15 * time.Second,
	}

	httpReqPayload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", getEnv("PORTAL_BG_URL", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0")+"/tracking", bytes.NewBuffer(httpReqPayload))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", "Basic "+getEnv("PORTAL_BG_API_KEY", ""))

	httpRes, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRes.Body.Close()

	var httpResData ApiDownloadResponse
	err = json.NewDecoder(httpRes.Body).Decode(&httpResData)
	if err != nil {
		return nil, err
	}

	httpResPayload, err := json.Marshal(httpResData)
	if err != nil {
		return nil, err
	}

	log.Println("Response:", string(httpResPayload))

	return &httpResData, nil

}

func (s *Server) ApiListTransaction(ctx context.Context, req *ApiListTransactionRequest) (*ApiListTransactionResponse, error) {

	var transport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 15 * time.Second,
		}).Dial,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		TLSHandshakeTimeout: 15 * time.Second,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   15 * time.Second,
	}

	httpReqParams, err := query.Values(req)
	if err != nil {
		return nil, err
	}

	httpReqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	log.Printf("[api][func: ApiListTransaction] REQUEST: %s", string(httpReqBytes))

	httpReq, err := http.NewRequest("GET", getEnv("PORTAL_BG_URL", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0")+"/listTransaction?"+httpReqParams.Encode(), nil)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Add("Authorization", "Basic "+getEnv("PORTAL_BG_API_KEY", ""))

	httpRes, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRes.Body.Close()

	var httpResData ApiListTransactionResponse
	err = json.NewDecoder(httpRes.Body).Decode(&httpResData)
	if err != nil {
		return nil, err
	}

	httpResDataBytes, err := json.Marshal(httpResData)
	if err != nil {
		return nil, err
	}

	log.Printf("[api][func: ApiListTransaction] RESPONSE: %s", string(httpResDataBytes))

	return &httpResData, nil

}

func (s *Server) ApiCreateIssuing(ctx context.Context, req *ApiBgIssuingRequest) (*ApiBgIssuingResponse, error) {

	log.WithField("request", req).Infoln("[api][func ApiCreateIssuing] request coming...")

	var transport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 15 * time.Second,
		}).Dial,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		TLSHandshakeTimeout: 15 * time.Second,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   15 * time.Second,
	}

	httpReqPayload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	log.Indo("[api][func: ApiCreateIssuing] REQUEST Marshal :", string(httpReqPayload))

	httpReq, err := http.NewRequest("POST", getEnv("PORTAL_BG_URL", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0")+"/applyBG", bytes.NewBuffer(httpReqPayload))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", "Basic "+getEnv("PORTAL_BG_API_KEY", ""))

	httpRes, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRes.Body.Close()

	var httpResData ApiBgIssuingResponse
	err = json.NewDecoder(httpRes.Body).Decode(&httpResData)
	if err != nil {
		return nil, err
	}

	httpResPayload, err := json.Marshal(httpResData)
	if err != nil {
		return nil, err
	}

	log.Println("Response:", string(httpResPayload))

	if httpResData.ResponseCode != "00" {
		return nil, errors.New(string(*httpResData.ResponseMessage))
	}

	result := &httpResData

	return result, nil

}

func (s *Server) ApiCheckIssuingStatus(ctx context.Context, req *ApiBgTrackingRequest) (*ApiBgTrackingResponse, error) {

	var transport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 15 * time.Second,
		}).Dial,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		TLSHandshakeTimeout: 15 * time.Second,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   15 * time.Second,
	}

	httpReqData := ApiBgIssuingData{
		RegistrationNo: req.RegistrationNo,
	}

	httpReqPayload, err := json.Marshal(httpReqData)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", getEnv("PORTAL_BG_URL", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0")+"/tracking", bytes.NewBuffer(httpReqPayload))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", "Basic "+getEnv("PORTAL_BG_API_KEY", ""))

	httpRes, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRes.Body.Close()

	var httpResData ApiBgTrackingResponse
	err = json.NewDecoder(httpRes.Body).Decode(&httpResData)
	if err != nil {
		return nil, err
	}

	httpResPayload, err := json.Marshal(httpResData)
	if err != nil {
		return nil, err
	}

	log.Println("Response:", string(httpResPayload))

	if httpResData.ResponseCode != "00" {
		return nil, errors.New(string(*httpResData.ResponseMessage))
	}

	return &httpResData, nil

}

func (s *Server) ApiInquiryLimitIndividual(ctx context.Context, req *ApiInquiryLimitIndividualRequest) (*ApiInquiryLimitIndividualResponse, error) {

	var transport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 15 * time.Second,
		}).Dial,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		TLSHandshakeTimeout: 15 * time.Second,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   15 * time.Second,
	}

	httpReqParam, err := query.Values(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("GET", getEnv("PORTAL_BG_URL", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0")+"/inquiryLimitIndividu?"+httpReqParam.Encode(), nil)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Add("Authorization", "Basic "+getEnv("PORTAL_BG_API_KEY", ""))

	httpRes, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRes.Body.Close()

	var httpResData ApiInquiryLimitIndividualResponse
	err = json.NewDecoder(httpRes.Body).Decode(&httpResData)
	if err != nil {
		return nil, err
	}

	httpResPayload, err := json.Marshal(httpResData)
	if err != nil {
		return nil, err
	}

	log.Println("Response:", string(httpResPayload))

	// if httpResData.ResponseCode != "00" {
	// 	return nil, errors.New(string(*httpResData.ResponseMessage))
	// }

	return &httpResData, nil

}

func (s *Server) ApiUploadEncode(ctx context.Context, req *ApiUploadEncodeRequest) (*ApiUploadEncodeResponse, error) {

	var transport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 15 * time.Second,
		}).Dial,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		TLSHandshakeTimeout: 15 * time.Second,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   15 * time.Second,
	}

	req.ChannelId = getEnv("BG_CHANNEL_ID", "2")

	httpReqPayload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", getEnv("PORTAL_BG_URL", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0")+"/uploadEncode", bytes.NewBuffer(httpReqPayload))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", "Basic "+getEnv("PORTAL_BG_API_KEY", ""))

	httpRes, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRes.Body.Close()

	var httpResData ApiUploadEncodeResponse
	err = json.NewDecoder(httpRes.Body).Decode(&httpResData)
	if err != nil {
		return nil, err
	}

	httpResPayload, err := json.Marshal(httpResData)
	if err != nil {
		return nil, err
	}

	log.Println("Response:", string(httpResPayload))

	if httpResData.ResponseCode != "00" {
		return nil, errors.New(string(*httpResData.ResponseMessage))
	}

	return &httpResData, nil

}
