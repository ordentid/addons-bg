package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"bitbucket.bri.co.id/scm/addons/addons-bg-service/server/db"
	task_pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/lib/stubs/task"
	pb "bitbucket.bri.co.id/scm/addons/addons-bg-service/server/pb"
	"github.com/google/go-querystring/query"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
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
	StartDate             string `url:"startDate"`
	EndDate               string `url:"endDate"`
	Branch                string `url:"branch"`
	ApplicantName         string `url:"applicationName"`
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

type ApiInquiryThirdPartyByStatusRequest struct {
	Status string `json:"status"`
}

type ApiInquiryThirdPartyByStatusResponse struct {
	ResponseCode    string                  `json:"responseCode"`
	ResponseMessage string                  `json:"responseMessage"`
	ResponseData    []*ApiInquiryThirdParty `json:"responseData"`
}

type ApiDownloadRequest struct {
	ReferenceNo string `json:"referenceNo"`
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

type ApiBeneficiary struct {
	BeneficiaryID uint64 `json:"beneficiaryId,string"`
	ThirdPartyID  uint64 `json:"thirdPartyId,string"`
	Cif           string `json:"cif"`
	FullName      string `json:"fullName"`
	CreatedDate   string `json:"createdDate"`
	ModifiedDate  string `json:"modifiedDate"`
	Status        string `json:"status"`
}

type ApiBgIssuingRequest struct {
	AccountNo              string `json:"account_no"`
	ApplicantName          string `json:"applicant_name"`
	ApplicantAddress       string `json:"applicant_address"`
	IsIndividu             string `json:"is_individu"`
	NIK                    string `json:"nik"`
	BirthDate              string `json:"birth_date"`
	Gender                 string `json:"gender"`
	NPWPNo                 string `json:"npwp_no"`
	DateEstablished        string `json:"tanggal_berdiri"`
	CompanyType            string `json:"company_type"`
	IsPlafond              string `json:"is_plafond"`
	TransactionType        string `json:"transaction_type"`
	IsEndOfYearBg          string `json:"is_bg_akhir_tahun"`
	NRK                    string `json:"nrk"`
	ProjectName            string `json:"project_name"`
	ThirdPartyId           string `json:"third_party_id"`
	BeneficiaryName        string `json:"beneficiary_name"`
	ProjectAmount          string `json:"project_amount"`
	ContractNo             string `json:"contract_no"`
	ContractDate           string `json:"contract_date"`
	Currency               string `json:"currency"`
	Amount                 string `json:"amount"`
	EffectiveDate          string `json:"effective_date"`
	MaturityDate           string `json:"maturity_date"`
	ClaimPeriod            string `json:"claim_periode"`
	IssuingBranch          string `json:"issuing_branch"`
	BranchPrinter          string `json:"pencetak_branch"`
	ContraGuarantee        string `json:"contra_guarantee"`
	InsuranceLimitId       string `json:"insurance_limit_id"`
	SP3No                  string `json:"sp3_no"`
	HoldAccountNo          string `json:"hold_account_no"`
	HoldAccountAmount      string `json:"hold_account_amount"`
	ConsumerLimitId        string `json:"consumer_limit_id"`
	ConsumerLimitAmount    string `json:"consumer_limit_amount"`
	ApplicantContactPerson string `json:"applicant_contact_person"`
	ApplicantPhoneNumber   string `json:"applicant_phone_number"`
	ApplicantEmail         string `json:"applicant_email"`
	ChannelId              string `json:"channel_id"`
	ApplicantCustomerId    string `json:"applicant_customer_id"`
	BeneficiaryCustomerId  string `json:"beneficiary_customer_id"`
	LegalDocument          string `json:"document_legalitas"`
	ContractDocument       string `json:"document_contract"`
	Sp3Document            string `json:"document_sp3"`
	OthersDocument         string `json:"document_others"`
}

type ApiBgIssuingResponse struct {
	ResponseCode    string           `json:"responseCode"`
	ResponseMessage string           `json:"responseMessage"`
	Data            ApiBgIssuingData `json:"responseData"`
}

type ApiBgIssuingData struct {
	RegistrationNo string `json:"registration_no"`
}

type UrlObject struct {
	Url string `json:"url"`
}

func (s *Server) GetApplicantName(ctx context.Context, req *pb.GetApplicantNameRequest) (*pb.GetApplicantNameResponse, error) {
	result := &pb.GetApplicantNameResponse{
		Error:   false,
		Code:    200,
		Message: "List Data",
		Data:    []*pb.ApplicantName{},
	}

	me, err := s.manager.GetMeFromJWT(ctx, "")
	if err != nil {
		return nil, err
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		ctx = metadata.NewOutgoingContext(context.Background(), md)
	}
	var header, trailer metadata.MD

	taskConn, err := grpc.Dial(getEnv("TASK_SERVICE", ":9090"), opts...)
	if err != nil {
		logrus.Errorln("Failed connect to Task Service: %v", err)
		return nil, status.Errorf(codes.Internal, "Error Internal")
	}
	taskConn.Connect()
	defer taskConn.Close()

	taskClient := task_pb.NewTaskServiceClient(taskConn)

	taskFilter := &task_pb.Task{
		Type: "BG Issuing",
	}

	filter := []string{
		"data.publishing.thirdPartyID:" + strconv.FormatUint(req.ThirdPartyID, 10),
	}

	dataReq := &task_pb.ListTaskRequest{
		Task:   taskFilter,
		Filter: strings.Join(filter, ","),
		In:     me.CompanyIDs,
	}

	dataList, err := taskClient.GetListTask(ctx, dataReq, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	names := []string{}

	for _, v := range dataList.Data {

		taskData := pb.IssuingData{}
		json.Unmarshal([]byte(v.Data), &taskData)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		if !contains(names, taskData.Applicant.Name) {
			names = append(names, taskData.Applicant.Name)
		}

	}

	for _, v := range names {

		result.Data = append(result.Data, &pb.ApplicantName{
			Name: v,
		})

	}

	return result, nil
}

func (s *Server) GetBeneficiaryName(ctx context.Context, req *pb.GetBeneficiaryNameRequest) (*pb.GetBeneficiaryNameResponse, error) {
	result := &pb.GetBeneficiaryNameResponse{
		Error:   false,
		Code:    200,
		Message: "List Data",
		Data:    []*pb.BeneficiaryName{},
	}

	client := &http.Client{}
	if getEnv("ENV", "PRODUCTION") != "PRODUCTION" {
		proxyURL, err := url.Parse("http://localhost:5100")
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
	}

	httpReqData := ApiInquiryBenficiaryRequest{
		ThirdPartyID: req.ThirdPartyID,
	}

	httpReqParam, err := query.Values(httpReqData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	httpReq, err := http.NewRequest("GET", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0/inquiryBeneficiary?"+httpReqParam.Encode(), nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	httpReq.Header.Add("Authorization", "Basic YnJpY2FtczpCcmljYW1zNGRkMG5z")

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

	if httpResData.ResponseCode == "00" {
		for _, v := range httpResData.ResponseData {
			result.Data = append(result.Data, &pb.BeneficiaryName{
				BeneficiaryId: v.BeneficiaryID,
				ThirdPartyId:  v.ThirdPartyID,
				Cif:           v.Cif,
				Fullname:      v.FullName,
				Status:        v.Status,
			})
		}
	}

	return result, nil
}

func (s *Server) GetThirdParty(ctx context.Context, req *pb.GetThirdPartyRequest) (*pb.GetThirdPartyResponse, error) {
	result := &pb.GetThirdPartyResponse{
		Error:   false,
		Code:    200,
		Message: "List Data",
		Data:    []*pb.ThirdParty{},
	}

	me, err := s.manager.GetMeFromJWT(ctx, "")
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	if getEnv("ENV", "PRODUCTION") != "PRODUCTION" {
		proxyURL, err := url.Parse("http://localhost:5100")
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
	}

	if me.UserType == "ba" {
		httpReqData := ApiInquiryThirdPartyByStatusRequest{
			Status: "Active",
		}

		httpReqPayload, err := json.Marshal(httpReqData)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		httpReq, err := http.NewRequest("POST", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0/inquiryThirdParty", bytes.NewBuffer(httpReqPayload))
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

		var httpResData ApiInquiryThirdPartyByStatusResponse
		err = json.NewDecoder(httpRes.Body).Decode(&httpResData)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		logrus.Println(httpResData.ResponseCode)

		if httpResData.ResponseCode == "00" {
			for _, v := range httpResData.ResponseData {
				result.Data = append(result.Data, &pb.ThirdParty{
					Id:   v.ThirdPartyID,
					Name: v.FullName,
				})
			}
		}
	} else {
		logrus.Println(req.Type)

		filter := &db.ListFilter{}

		filter.Filter = "company_id:" + strconv.FormatUint(me.CompanyID, 10)

		filterMapped := ""
		if req.Type == *pb.ThirdPartyType_NeedMapping.Enum() {
			filterMapped = ",is_mapped:false"
		} else if req.Type == *pb.ThirdPartyType_IsMapped.Enum() {
			filterMapped = ",is_mapped:true"
		}

		filter.Filter = filter.Filter + filterMapped

		logrus.Println("------------------------")
		logrus.Println(me.CompanyID)
		logrus.Println("------------------------")
		logrus.Println(filter.Filter)
		logrus.Println("------------------------")

		thirdPartyNameList, err := s.provider.GetMapping(ctx, filter)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		logrus.Print(thirdPartyNameList)

		ids := []string{}

		for _, v := range thirdPartyNameList {
			if !contains(ids, strconv.FormatUint(v.ThirdPartyID, 10)) {
				ids = append(ids, strconv.FormatUint(v.ThirdPartyID, 10))
			}
		}

		for _, v := range ids {
			id, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
			}

			name := ""

			httpReqData := ApiInquiryThirdPartyByIDRequest{
				ThirdPartyID: id,
			}

			httpReqPayload, err := json.Marshal(httpReqData)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
			}

			httpReq, err := http.NewRequest("POST", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0/inquiryThirdParty", bytes.NewBuffer(httpReqPayload))
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

			var httpResData ApiInquiryThirdPartyByIDResponse
			err = json.NewDecoder(httpRes.Body).Decode(&httpResData)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
			}

			logrus.Println(httpResData.ResponseCode)

			if httpResData.ResponseCode == "00" {
				name = httpResData.ResponseData.FullName
			}

			result.Data = append(result.Data, &pb.ThirdParty{
				Id:   id,
				Name: name,
			})
		}
	}

	return result, nil
}

func (s *Server) GetTransactionAttachment(ctx context.Context, req *pb.GetTransactionAttachmentRequest) (*pb.GetTransactionAttachmentResponse, error) {
	result := &pb.GetTransactionAttachmentResponse{
		Error:   false,
		Code:    200,
		Message: "Data",
	}

	if req.ReferenceNo != "" {
		client := &http.Client{}
		if getEnv("ENV", "PRODUCTION") != "PRODUCTION" {
			proxyURL, err := url.Parse("http://localhost:5100")
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
			}

			client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
		}

		httpReqData := ApiDownloadRequest{
			ReferenceNo: req.ReferenceNo,
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

		if httpResData.ResponseCode != "00" {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", httpResData.ResponseMessage)
		} else {
			for _, v := range httpResData.ResponseData {
				result.Data = append(result.Data, v.Url)
			}
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

	result.Pagination = &pb.PaginationResponse{
		Limit:      req.Limit,
		Page:       req.Page,
		TotalRows:  0,
		TotalPages: 0,
	}

	me, err := s.manager.GetMeFromJWT(ctx, "")
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	if getEnv("ENV", "PRODUCTION") != "PRODUCTION" {
		proxyURL, err := url.Parse("http://localhost:5100")
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
	}

	filterData := []string{
		"company_id:" + strconv.FormatUint(me.CompanyID, 10),
		"is_mapped:true",
	}

	if req.Transaction.ThirdPartyID > 0 {
		filterData = append(filterData, "third_party_id:"+strconv.FormatUint(req.Transaction.ThirdPartyID, 10))
	}

	filter := &db.ListFilter{
		Filter: strings.Join(filterData, ","),
	}

	logrus.Println("---------------------------")
	logrus.Println(filter.Filter)
	logrus.Println("---------------------------")

	mappingORM, err := s.provider.GetMapping(ctx, filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	beneficiaryIDs := []string{}
	for _, v := range mappingORM {
		if v.BeneficiaryID == 10101010 {
			res, err := s.GetBeneficiaryName(ctx, &pb.GetBeneficiaryNameRequest{ThirdPartyID: v.ThirdPartyID})
			if err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				}
			}

			for _, d := range res.Data {
				if !contains(beneficiaryIDs, strconv.FormatUint(d.BeneficiaryId, 10)) {
					beneficiaryIDs = append(beneficiaryIDs, strconv.FormatUint(d.BeneficiaryId, 10))
				}
			}
		} else {
			if !contains(beneficiaryIDs, strconv.FormatUint(v.BeneficiaryID, 10)) {
				beneficiaryIDs = append(beneficiaryIDs, strconv.FormatUint(v.BeneficiaryID, 10))
			}
		}
	}

	httpReqParamsOpt := ApiListTransactionRequest{
		Page:  uint64(req.Page),
		Limit: uint64(req.Limit),
	}

	logrus.Println("---------------------------")
	logrus.Println(strings.Join(beneficiaryIDs, ","))
	logrus.Println("---------------------------")

	httpReqParamsOpt.BeneficiaryId = strings.Join(beneficiaryIDs, ",")

	if req.Transaction != nil {
		if req.Transaction.StartDate != "" && req.Transaction.EndDate != "" {
			httpReqParamsOpt.StartDate = req.Transaction.StartDate
			httpReqParamsOpt.EndDate = req.Transaction.EndDate
		} else {
			return nil, status.Errorf(codes.InvalidArgument, "Start Date and End Date is Required")
		}

		if req.Transaction.BeneficiaryID > 0 {
			httpReqParamsOpt.BeneficiaryId = strconv.FormatUint(req.Transaction.BeneficiaryID, 10)
		}

		if req.Transaction.ThirdPartyID > 0 {
			httpReqParamsOpt.ThirdPartyId = req.Transaction.ThirdPartyID
		}

		if req.Transaction.ClaimPeriod > 0 {
			httpReqParamsOpt.ClaimPeriod = strconv.FormatUint(uint64(req.Transaction.ClaimPeriod), 10)
		}

		if req.Transaction.Status != "" {
			httpReqParamsOpt.Status = req.Transaction.Status
		}

		if req.Transaction.ReferenceNo != "" {
			httpReqParamsOpt.ReferenceNo = req.Transaction.ReferenceNo
		}
	}

	httpReqParams, err := query.Values(httpReqParamsOpt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	logrus.Println("---------------------------")
	logrus.Println(httpReqParams.Encode())
	logrus.Println("---------------------------")

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
		// return nil, status.Errorf(codes.Internal, "Internal Error: %v", httpResData.ResponseMessage)
	} else {
		for _, d := range httpResData.ResponseData {
			transactionPB := &pb.Transaction{
				TransactionID:     d.TransactionId,
				ThirdPartyID:      d.ThirdPartyId,
				ThirdPartyName:    d.ThirdPartyName,
				ReferenceNo:       d.ReferenceNo,
				RegistrationNo:    d.RegistrationNo,
				ApplicantName:     d.ApplicantName,
				BeneficiaryID:     d.BeneficiaryId,
				BeneficiaryName:   d.BeneficiaryName,
				IssueDate:         d.IssueDate,
				EffectiveDate:     d.EffectiveDate,
				ExpiryDate:        d.ExpiryDate,
				ClaimPeriod:       d.ClaimPeriod,
				ClosingDate:       d.ClosingDate,
				Currency:          d.Currency,
				Amount:            d.Amount,
				CreatedDate:       d.CreatedDate,
				ModifiedDate:      d.ModifiedDate,
				Remark:            d.Remark,
				Status:            d.Status,
				ChannelID:         d.ChannelId,
				ChannelName:       d.ChannelName,
				TransactionTypeID: pb.BgType(d.TransactionTypeId),
			}

			result.Data = append(result.Data, transactionPB)
		}

		result.Pagination = &pb.PaginationResponse{
			Limit:      int32(httpResData.Pagination.Limit),
			Page:       int32(httpResData.Pagination.Page),
			TotalRows:  int64(httpResData.Pagination.TotalRecord),
			TotalPages: int32(httpResData.Pagination.TotalPage),
		}
	}

	return result, nil
}

func (s *Server) GetTransactionDetail(ctx context.Context, req *pb.GetTransactionDetailRequest) (*pb.GetTransactionDetailResponse, error) {
	result := &pb.GetTransactionDetailResponse{
		Error:   false,
		Code:    200,
		Message: "Data",
	}

	if req.ReferenceNo != "" {
		client := &http.Client{}
		if getEnv("ENV", "PRODUCTION") != "PRODUCTION" {
			proxyURL, err := url.Parse("http://localhost:5100")
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
			}

			client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
		}

		httpReqParamsOpt := ApiListTransactionRequest{
			ReferenceNo: req.ReferenceNo,
			Page:        1,
			Limit:       1,
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

		d := httpResData.ResponseData[0]

		result.Data = &pb.Transaction{
			TransactionID:     d.TransactionId,
			ThirdPartyID:      d.ThirdPartyId,
			ThirdPartyName:    d.ThirdPartyName,
			ReferenceNo:       d.ReferenceNo,
			RegistrationNo:    d.RegistrationNo,
			ApplicantName:     d.ApplicantName,
			BeneficiaryID:     d.BeneficiaryId,
			BeneficiaryName:   d.BeneficiaryName,
			IssueDate:         d.IssueDate,
			EffectiveDate:     d.EffectiveDate,
			ExpiryDate:        d.ExpiryDate,
			ClaimPeriod:       d.ClaimPeriod,
			ClosingDate:       d.ClosingDate,
			Currency:          d.Currency,
			Amount:            d.Amount,
			CreatedDate:       d.CreatedDate,
			ModifiedDate:      d.ModifiedDate,
			Remark:            d.Remark,
			Status:            d.Status,
			ChannelID:         d.ChannelId,
			ChannelName:       d.ChannelName,
			TransactionTypeID: pb.BgType(d.TransactionTypeId),
		}
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

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		ctx = metadata.NewOutgoingContext(context.Background(), md)
	}
	var header, trailer metadata.MD

	taskConn, err := grpc.Dial(getEnv("TASK_SERVICE", ":9090"), opts...)
	if err != nil {
		logrus.Errorln("Failed connect to Task Service: %v", err)
		return nil, status.Errorf(codes.Internal, "Error Internal")
	}
	taskConn.Connect()
	defer taskConn.Close()

	taskClient := task_pb.NewTaskServiceClient(taskConn)

	switch req.Type {
	case "BG Mapping":

		logrus.Println("----------------------")
		logrus.Println("Save BG Mapping")
		logrus.Println("----------------------")

		taskData := req.MappingData
		taskDataBak := req.MappingDataBackup

		ids := []string{}

		for _, v := range taskDataBak {

			needDelete := true

			for _, check := range taskData {

				if check.ThirdPartyID == v.ThirdPartyID {

					needDelete = false
					logrus.Println("Break at: " + strconv.FormatUint(check.ThirdPartyID, 10))
					break

				}

			}

			logrus.Println(needDelete)

			if needDelete {

				logrus.Println("----------------------")
				logrus.Println("Get Mapping Digital Task Data")
				logrus.Println(v)
				logrus.Println("----------------------")

				filter := []string{
					"company_id:" + strconv.FormatUint(v.CompanyID, 10),
					"data.0.thirdPartyID:" + strconv.FormatUint(v.ThirdPartyID, 10),
				}

				taskMappingDigitalRes, err := taskClient.GetListTask(ctx, &task_pb.ListTaskRequest{Filter: strings.Join(filter, ","), Task: &task_pb.Task{Type: "BG Mapping Digital"}, Page: 1, Limit: 1}, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					if !errors.Is(err, gorm.ErrRecordNotFound) {
						return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
					}
				}

				logrus.Println("----------------------")
				logrus.Println("Mapping Digital Task Response:")
				logrus.Println(taskMappingDigitalRes.Data)
				logrus.Println("----------------------")

				for _, taskMappingDigitalResData := range taskMappingDigitalRes.Data {

					taskMappingDigitalData := []*pb.MappingDigitalData{}
					json.Unmarshal([]byte(taskMappingDigitalResData.GetData()), &taskMappingDigitalData)
					if err != nil {
						return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
					}

					logrus.Println("----------------------")
					logrus.Println("To Delete Mapping Digital Task ID: " + strconv.FormatUint(taskMappingDigitalResData.TaskID, 10))
					logrus.Println("----------------------")

					_, err := taskClient.SetTask(ctx, &task_pb.SetTaskRequest{TaskID: taskMappingDigitalResData.TaskID, Action: "delete", Comment: "delete"}, grpc.Header(&header), grpc.Trailer(&trailer))
					if err != nil {
						return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
					}

					if len(taskMappingDigitalData) > 0 {
						mappingFilter := []string{
							"company_id:" + strconv.FormatUint(taskMappingDigitalData[0].CompanyID, 10),
							"third_party_id:" + strconv.FormatUint(taskMappingDigitalData[0].ThirdPartyID, 10),
						}

						mappingListFilter := &db.ListFilter{
							Filter: strings.Join(mappingFilter, ","),
						}

						mappingORMs, err := s.provider.GetMapping(ctx, mappingListFilter)
						if err != nil {
							if !errors.Is(err, gorm.ErrRecordNotFound) {
								return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
							}
						}

						for _, mappingORM := range mappingORMs {
							if mappingORM.Id > 0 {
								if !contains(ids, strconv.FormatUint(mappingORM.Id, 10)) {
									ids = append(ids, strconv.FormatUint(mappingORM.Id, 10))
								}
							}
						}
					}

				}

				logrus.Println("----------------------")
				logrus.Println("To Delete Mapping Digital Data:")
				logrus.Println(ids)
				logrus.Println("----------------------")

			}

			mappingORM, err := s.provider.GetMappingDetail(ctx, &pb.MappingORM{ThirdPartyID: v.ThirdPartyID, BeneficiaryID: 10101010, CompanyID: v.CompanyID})
			if err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				}
			}

			if mappingORM.Id > 0 {
				if !contains(ids, strconv.FormatUint(mappingORM.Id, 10)) {
					ids = append(ids, strconv.FormatUint(mappingORM.Id, 10))
				}
			}

		}

		logrus.Println("----------------------")
		logrus.Println("To Delete Mapping Data:")
		logrus.Println(ids)
		logrus.Println("----------------------")

		err = s.provider.DeleteMapping(ctx, ids)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		for _, v := range taskData {

			data := &pb.MappingORM{
				CompanyID:     v.CompanyID,
				ThirdPartyID:  v.ThirdPartyID,
				BeneficiaryID: 10101010,
				IsMapped:      false,
				CreatedByID:   me.UserID,
				UpdatedByID:   me.UserID,
			}

			if v.IsAllowAllBeneficiary {
				data.IsMapped = true
			}

			mappingORM, err := s.provider.GetMappingDetail(ctx, &pb.MappingORM{ThirdPartyID: v.ThirdPartyID, BeneficiaryID: 10101010, CompanyID: v.CompanyID})
			if err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				}
			}

			if mappingORM.Id > 0 {
				data.Id = mappingORM.Id
			}

			mappingORM, err = s.provider.UpdateOrCreateMapping(ctx, data)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
			}

			mappingPB, err := mappingORM.ToPB(ctx)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
			}

			result.Data = append(result.Data, &mappingPB)

		}

	case "BG Mapping Digital":

		taskData := req.MappingDigitalData
		taskDataBak := req.MappingDigitalDataBackup

		ids := []string{}

		for _, v := range taskDataBak {

			mappingFilter := []string{
				"company_id:" + strconv.FormatUint(v.CompanyID, 10),
				"third_party_id:" + strconv.FormatUint(v.ThirdPartyID, 10),
			}

			mappingListFilter := &db.ListFilter{
				Filter: strings.Join(mappingFilter, ","),
			}

			mappingORMs, err := s.provider.GetMapping(ctx, mappingListFilter)
			if err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				}
			}

			for _, mappingORM := range mappingORMs {
				if mappingORM.Id > 0 {
					if !contains(ids, strconv.FormatUint(mappingORM.Id, 10)) {
						ids = append(ids, strconv.FormatUint(mappingORM.Id, 10))
					}
				}
			}

		}

		err = s.provider.DeleteMapping(ctx, ids)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		for _, v := range taskData {

			data := &pb.MappingORM{
				CompanyID:     v.CompanyID,
				ThirdPartyID:  v.ThirdPartyID,
				BeneficiaryID: v.BeneficiaryId,
				IsMapped:      true,
				CreatedByID:   me.UserID,
				UpdatedByID:   me.UserID,
			}

			mappingORM, err := s.provider.GetMappingDetail(ctx, &pb.MappingORM{ThirdPartyID: v.ThirdPartyID, BeneficiaryID: v.BeneficiaryId, CompanyID: v.CompanyID})
			if err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				}
			}

			if mappingORM.Id > 0 {
				data.Id = mappingORM.Id
			}

			mappingORM, err = s.provider.UpdateOrCreateMapping(ctx, data)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
			}

			mappingPB, err := mappingORM.ToPB(ctx)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
			}

			result.Data = append(result.Data, &mappingPB)

		}

	}

	return result, nil
}

func (s *Server) DeleteTransaction(ctx context.Context, req *pb.DeleteTransactionRequest) (*pb.DeleteTransactionResponse, error) {
	result := &pb.DeleteTransactionResponse{
		Error:   false,
		Code:    200,
		Message: "Data",
	}

	// me, err := s.manager.GetMeFromJWT(ctx, "")
	// if err != nil {
	// 	return nil, err
	// }

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		ctx = metadata.NewOutgoingContext(context.Background(), md)
	}
	var header, trailer metadata.MD

	taskConn, err := grpc.Dial(getEnv("TASK_SERVICE", ":9090"), opts...)
	if err != nil {
		logrus.Errorln("Failed connect to Task Service: %v", err)
		return nil, status.Errorf(codes.Internal, "Error Internal")
	}
	taskConn.Connect()
	defer taskConn.Close()

	taskClient := task_pb.NewTaskServiceClient(taskConn)

	switch req.Type {
	case "BG Mapping":

		logrus.Println("----------------------")
		logrus.Println("Save BG Mapping")
		logrus.Println("----------------------")

		taskData := req.MappingData
		// taskDataBak := req.MappingDataBackup

		ids := []string{}

		for _, v := range taskData {

			logrus.Println("----------------------")
			logrus.Println("Get Mapping Digital Task Data")
			logrus.Println(v)
			logrus.Println("----------------------")

			filter := []string{
				"company_id:" + strconv.FormatUint(v.CompanyID, 10),
				"data.0.thirdPartyID:" + strconv.FormatUint(v.ThirdPartyID, 10),
			}

			taskMappingDigitalRes, err := taskClient.GetListTask(ctx, &task_pb.ListTaskRequest{Filter: strings.Join(filter, ","), Task: &task_pb.Task{Type: "BG Mapping Digital"}, Page: 1, Limit: 1}, grpc.Header(&header), grpc.Trailer(&trailer))
			if err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				}
			}

			logrus.Println("----------------------")
			logrus.Println("Mapping Digital Task Response:")
			logrus.Println(taskMappingDigitalRes.Data)
			logrus.Println("----------------------")

			for _, taskMappingDigitalResData := range taskMappingDigitalRes.Data {

				taskMappingDigitalData := []*pb.MappingDigitalData{}
				json.Unmarshal([]byte(taskMappingDigitalResData.GetData()), &taskMappingDigitalData)
				if err != nil {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				}

				logrus.Println("----------------------")
				logrus.Println("To Delete Mapping Digital Task ID: " + strconv.FormatUint(taskMappingDigitalResData.TaskID, 10))
				logrus.Println("----------------------")

				_, err := taskClient.SetTask(ctx, &task_pb.SetTaskRequest{TaskID: taskMappingDigitalResData.TaskID, Action: "delete", Comment: "delete"}, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				}

				if len(taskMappingDigitalData) > 0 {
					mappingFilter := []string{
						"company_id:" + strconv.FormatUint(taskMappingDigitalData[0].CompanyID, 10),
						"third_party_id:" + strconv.FormatUint(taskMappingDigitalData[0].ThirdPartyID, 10),
					}

					mappingListFilter := &db.ListFilter{
						Filter: strings.Join(mappingFilter, ","),
					}

					mappingORMs, err := s.provider.GetMapping(ctx, mappingListFilter)
					if err != nil {
						if !errors.Is(err, gorm.ErrRecordNotFound) {
							return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
						}
					}

					for _, mappingORM := range mappingORMs {
						if mappingORM.Id > 0 {
							if !contains(ids, strconv.FormatUint(mappingORM.Id, 10)) {
								ids = append(ids, strconv.FormatUint(mappingORM.Id, 10))
							}
						}
					}
				}

			}

			logrus.Println("----------------------")
			logrus.Println("To Delete Mapping Digital Data:")
			logrus.Println(ids)
			logrus.Println("----------------------")

			mappingORM, err := s.provider.GetMappingDetail(ctx, &pb.MappingORM{ThirdPartyID: v.ThirdPartyID, BeneficiaryID: 10101010, CompanyID: v.CompanyID})
			if err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				}
			}

			if mappingORM.Id > 0 {
				if !contains(ids, strconv.FormatUint(mappingORM.Id, 10)) {
					ids = append(ids, strconv.FormatUint(mappingORM.Id, 10))
				}
			}

		}

		logrus.Println("----------------------")
		logrus.Println("To Delete Mapping Data:")
		logrus.Println(ids)
		logrus.Println("----------------------")

		err = s.provider.DeleteMapping(ctx, ids)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

	case "BG Mapping Digital":

		taskData := req.MappingDigitalData
		// taskDataBak := req.MappingDigitalDataBackup

		ids := []string{}

		for _, v := range taskData {

			mappingFilter := []string{
				"company_id:" + strconv.FormatUint(v.CompanyID, 10),
				"third_party_id:" + strconv.FormatUint(v.ThirdPartyID, 10),
			}

			mappingListFilter := &db.ListFilter{
				Filter: strings.Join(mappingFilter, ","),
			}

			mappingORMs, err := s.provider.GetMapping(ctx, mappingListFilter)
			if err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				}
			}

			byID := uint64(0)

			for _, mappingORM := range mappingORMs {
				if mappingORM.Id > 0 {
					byID = mappingORM.CreatedByID
					if !contains(ids, strconv.FormatUint(mappingORM.Id, 10)) {
						ids = append(ids, strconv.FormatUint(mappingORM.Id, 10))
					}
				}
			}

			data := &pb.MappingORM{
				CompanyID:     v.CompanyID,
				ThirdPartyID:  v.ThirdPartyID,
				BeneficiaryID: 10101010,
				IsMapped:      false,
				CreatedByID:   byID,
				UpdatedByID:   byID,
			}

			mappingORM, err := s.provider.GetMappingDetail(ctx, &pb.MappingORM{ThirdPartyID: v.ThirdPartyID, BeneficiaryID: 10101010, CompanyID: v.CompanyID})
			if err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				}
			}

			if mappingORM.Id > 0 {
				data.Id = mappingORM.Id
			}

			_, err = s.provider.UpdateOrCreateMapping(ctx, data)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
			}

		}

		err = s.provider.DeleteMapping(ctx, ids)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

	}

	return result, nil
}

func (s *Server) CreateIssuing(ctx context.Context, req *pb.CreateIssuingRequest) (*pb.CreateIssuingResponse, error) {
	result := &pb.CreateIssuingResponse{
		Error:   false,
		Code:    200,
		Message: "Data",
	}

	httpReqData := ApiBgIssuingRequest{
		AccountNo:              req.Data.Account.GetAccountNumber(),
		ApplicantName:          req.Data.Applicant.GetName(),
		ApplicantAddress:       req.Data.Applicant.GetAddress(),
		IsIndividu:             string(req.Data.Applicant.GetApplicantType().Number()),
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
		ThirdPartyId:           strconv.FormatUint(req.Data.Publishing.GetThirdPartyID(), 10),
		BeneficiaryName:        req.Data.Applicant.GetBeneficiaryName(),
		ProjectAmount:          strconv.FormatFloat(req.Data.Project.GetProjectAmount(), 'f', 10, 64),
		ContractNo:             req.Data.Project.GetContractNumber(),
		ContractDate:           req.Data.Project.GetProjectDate(),
		Currency:               req.Data.Project.GetBgCurrency(),
		Amount:                 strconv.FormatFloat(req.Data.Project.GetBgAmount(), 'f', 10, 64),
		EffectiveDate:          req.Data.Publishing.GetEffectiveDate(),
		MaturityDate:           req.Data.Publishing.GetExpiryDate(),
		ClaimPeriod:            strconv.FormatUint(req.Data.Publishing.GetClaimPeriod(), 10),
		IssuingBranch:          req.Data.Publishing.GetOpeningBranch(),
		BranchPrinter:          "Test",
		ContraGuarantee:        "Test",
		InsuranceLimitId:       "Test",
		SP3No:                  "Test",
		HoldAccountNo:          "Test",
		HoldAccountAmount:      "0",
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

	httpReq, err := http.NewRequest("POST", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0/applyBG", bytes.NewBuffer(httpReqPayload))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", "Basic YnJpY2FtczpCcmljYW1zNGRkMG5z")

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
