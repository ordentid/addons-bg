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

	return result, nil
}

func (s *Server) GetBeneficiaryName(ctx context.Context, req *pb.GetBeneficiaryNameRequest) (*pb.GetBeneficiaryNameResponse, error) {
	result := &pb.GetBeneficiaryNameResponse{
		Error:   false,
		Code:    200,
		Message: "List Data",
		Data:    []*pb.BeneficiaryName{},
	}

	httpReqData := ApiInquiryBenficiaryRequest{
		ThirdPartyID: req.ThirdPartyID,
	}

	httpReqParam, err := query.Values(httpReqData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	client := &http.Client{}
	if getEnv("ENV", "PRODUCTION") != "PRODUCTION" {
		proxyURL, err := url.Parse("http://localhost:5100")
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
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

	filterData := "company_id:" + strconv.FormatUint(me.CompanyID, 10)

	if req.Transaction.ThirdPartyID > 0 {
		filterData = filterData + ",third_party_id:" + strconv.FormatUint(req.Transaction.ThirdPartyID, 10)
	}

	filter := &db.ListFilter{
		Filter: filterData,
	}

	logrus.Println(filter.Data)

	mappingORM, err := s.provider.GetMapping(ctx, filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	beneficiaryIDs := []string{}
	for _, v := range mappingORM {
		if !contains(beneficiaryIDs, strconv.FormatUint(v.BeneficiaryID, 10)) {
			beneficiaryIDs = append(beneficiaryIDs, strconv.FormatUint(v.BeneficiaryID, 10))
		}
	}

	httpReqParamsOpt := ApiListTransactionRequest{
		Page:  uint64(req.Page),
		Limit: uint64(req.Limit),
	}

	httpReqParamsOpt.BeneficiaryId = strings.Join(beneficiaryIDs, ",")

	logrus.Println(httpReqParamsOpt.BeneficiaryId)

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

	client := &http.Client{}
	if getEnv("ENV", "PRODUCTION") != "PRODUCTION" {
		proxyURL, err := url.Parse("http://localhost:5100")
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
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

	taskRes, err := taskClient.GetTaskByID(ctx, &task_pb.GetTaskByIDReq{ID: req.TaskID}, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	task := &pb.Task{
		TaskID:             taskRes.Data.GetTaskID(),
		Type:               taskRes.Data.GetType(),
		Status:             taskRes.Data.GetStatus().String(),
		Step:               taskRes.Data.GetStep().String(),
		FeatureID:          taskRes.Data.GetFeatureID(),
		LastApprovedByID:   taskRes.Data.GetLastApprovedByID(),
		LastRejectedByID:   taskRes.Data.GetLastRejectedByID(),
		LastApprovedByName: taskRes.Data.GetLastApprovedByName(),
		LastRejectedByName: taskRes.Data.GetLastRejectedByName(),
		CreatedByName:      taskRes.Data.GetCreatedByName(),
		UpdatedByName:      taskRes.Data.GetUpdatedByName(),
		Reasons:            taskRes.Data.GetReasons(),
		Comment:            taskRes.Data.GetComment(),
		CompanyID:          taskRes.Data.GetCompanyID(),
		HoldingID:          taskRes.Data.GetHoldingID(),
		CreatedAt:          taskRes.Data.GetCreatedAt(),
		UpdatedAt:          taskRes.Data.GetUpdatedAt(),
	}

	switch task.Type {
	case "BG Mapping":
		taskData := []*pb.MappingData{}
		json.Unmarshal([]byte(taskRes.Data.GetData()), &taskData)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		for _, v := range taskData {
			httpReqParamsOpt := ApiInquiryBenficiaryRequest{
				ThirdPartyID: v.ThirdPartyID,
			}

			httpReqParams, err := query.Values(httpReqParamsOpt)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
			}

			logrus.Println(httpReqParams.Encode())

			httpReq, err := http.NewRequest("GET", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0/inquiryBeneficiary?"+httpReqParams.Encode(), nil)
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
				data := &pb.MappingORM{
					CompanyID:     v.CompanyID,
					ThirdPartyID:  v.ThirdPartyID,
					BeneficiaryID: 9999,
					IsMapped:      false,
					CreatedByID:   me.UserID,
					UpdatedByID:   me.UserID,
				}

				if v.IsAllowAllBeneficiary {
					data.IsMapped = true
				}

				mappingORM, err := s.provider.GetMappingDetail(ctx, &pb.MappingORM{ThirdPartyID: v.ThirdPartyID, BeneficiaryID: 9999, CompanyID: v.CompanyID})
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
			} else {
				for _, d := range httpResData.ResponseData {
					data := &pb.MappingORM{
						CompanyID:     v.CompanyID,
						ThirdPartyID:  v.ThirdPartyID,
						BeneficiaryID: d.BeneficiaryID,
						IsMapped:      false,
						CreatedByID:   me.UserID,
						UpdatedByID:   me.UserID,
					}

					if v.IsAllowAllBeneficiary {
						data.IsMapped = true
					}

					mappingORM, err := s.provider.GetMappingDetail(ctx, &pb.MappingORM{ThirdPartyID: v.ThirdPartyID, BeneficiaryID: d.BeneficiaryID, CompanyID: v.CompanyID})
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

		}
	case "BG Mapping Digital":
		taskData := []*pb.MappingDigitalData{}
		json.Unmarshal([]byte(taskRes.Data.GetData()), &taskData)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		for _, v := range taskData {
			httpReqParamsOpt := ApiInquiryBenficiaryRequest{
				ThirdPartyID: v.ThirdPartyID,
			}

			httpReqParams, err := query.Values(httpReqParamsOpt)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
			}

			logrus.Println(httpReqParams.Encode())

			httpReq, err := http.NewRequest("GET", "http://api.close.dev.bri.co.id:5557/gateway/apiPortalBG/1.0/inquiryBeneficiary?"+httpReqParams.Encode(), nil)
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
				data := &pb.MappingORM{
					CompanyID:     v.CompanyID,
					ThirdPartyID:  v.ThirdPartyID,
					BeneficiaryID: 9999,
					IsMapped:      true,
					CreatedByID:   me.UserID,
					UpdatedByID:   me.UserID,
				}

				mappingORM, err := s.provider.GetMappingDetail(ctx, &pb.MappingORM{ThirdPartyID: v.ThirdPartyID, BeneficiaryID: 9999, CompanyID: v.CompanyID})
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
			} else {
				for _, d := range httpResData.ResponseData {

					data := &pb.MappingORM{
						CompanyID:     v.CompanyID,
						ThirdPartyID:  v.ThirdPartyID,
						BeneficiaryID: d.BeneficiaryID,
						IsMapped:      true,
						CreatedByID:   me.UserID,
						UpdatedByID:   me.UserID,
					}

					mappingORM, err := s.provider.GetMappingDetail(ctx, &pb.MappingORM{ThirdPartyID: v.ThirdPartyID, BeneficiaryID: d.BeneficiaryID, CompanyID: v.CompanyID})
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
