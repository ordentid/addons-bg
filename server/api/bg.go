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

	applicantNameList, err := s.provider.GetApplicantName(ctx, req.ThirdPartyID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	result.Data = applicantNameList

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
		proxyURL, err := url.Parse("http://localhost:5002")
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

	if me.UserType == "ba" {
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
	} else {
		thirdPartyNameList, err := s.provider.GetThirdPartyByCompany(ctx, me.CompanyID)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		for _, v := range thirdPartyNameList {
			result.Data = append(result.Data, &pb.ThirdParty{
				Id:           v.Id,
				Name:         v.Name,
				ThirdPartyID: v.Id,
			})
		}
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

	client := &http.Client{}
	if getEnv("ENV", "PRODUCTION") != "PRODUCTION" {
		proxyURL, err := url.Parse("http://localhost:5002")
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

	taskRes, err := taskClient.GetTaskByID(ctx, &task_pb.GetTaskByIDReq{ID: req.TaskID, Type: "BG Mapping"}, grpc.Header(&header), grpc.Trailer(&trailer))
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
			httpReqParamsOpt := ApiListTransactionRequest{
				ThirdPartyId: strconv.FormatUint(v.ThirdPartyID, 10),
				Page:         1,
				Limit:        100,
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

			for _, d := range httpResData.ResponseData {
				transactionStatus := pb.TransactionStatus_value["MappingDigital"]
				if v.IsAllowAllBeneficiary {
					transactionStatus = pb.TransactionStatus_value["PendingForIssuing"]
				}
				t := &pb.TransactionORM{
					Amount:          d.Amount,
					ApplicantName:   d.ApplicantName,
					BeneficiaryName: d.BeneficiaryName,
					BgStatus:        pb.BgStatus_value[strings.ReplaceAll(d.Status, " ", "")],
					BgType:          int32(d.TransactionTypeId),
					ChannelID:       d.ChannelId,
					ChannelName:     d.ChannelName,
					ClaimPeriod:     d.ClaimPeriod,
					ClosingDate:     d.ClosingDate,
					CompanyID:       v.CompanyID,
					CreatedByID:     me.UserID,
					Currency:        d.Currency,
					DocumentPath:    d.DocumentPath,
					EffectiveDate:   d.EffectiveDate,
					ExpiryDate:      d.ExpiryDate,
					IssueDate:       d.IssueDate,
					ReferenceNo:     d.ReferenceNo,
					RegistrationNo:  d.RegistrationNo,
					Remark:          d.Remark,
					Status:          transactionStatus,
					ThirdPartyID:    v.ThirdPartyID,
					ThirdPartyName:  v.ThirdPartyName,
					TransactionID:   d.TransactionId,
					UpdatedByID:     me.UserID,
				}

				transactionORM, err := s.provider.GetTransactionDetail(ctx, &pb.TransactionORM{ReferenceNo: d.ReferenceNo})
				if err != nil {
					if !errors.Is(err, gorm.ErrRecordNotFound) {
						return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
					}
				}

				if transactionORM.Id > 0 {
					t.Id = transactionORM.Id
				}

				transactionORM, err = s.provider.UpdateOrCreateTransaction(ctx, t)
				if err != nil {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				}

				transactionPB, err := transactionORM.ToPB(ctx)
				if err != nil {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				}

				result.Data = append(result.Data, &transactionPB)

			}
		}
	case "BG Mapping Digital":
		taskData := []*pb.MappingDigitalData{}
		json.Unmarshal([]byte(taskRes.Data.GetData()), &taskData)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}

		for _, v := range taskData {
			httpReqParamsOpt := ApiListTransactionRequest{
				ThirdPartyId:    strconv.FormatUint(v.ThirdPartyID, 10),
				BeneficiaryName: v.BeneficiaryName,
				Page:            1,
				Limit:           100,
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

			for _, d := range httpResData.ResponseData {
				t := &pb.TransactionORM{
					Amount:          d.Amount,
					ApplicantName:   d.ApplicantName,
					BeneficiaryName: d.BeneficiaryName,
					BgStatus:        pb.BgStatus_value[strings.ReplaceAll(d.Status, " ", "")],
					BgType:          int32(d.TransactionTypeId),
					ChannelID:       d.ChannelId,
					ChannelName:     d.ChannelName,
					ClaimPeriod:     d.ClaimPeriod,
					ClosingDate:     d.ClosingDate,
					CompanyID:       v.CompanyID,
					CreatedByID:     me.UserID,
					Currency:        d.Currency,
					DocumentPath:    d.DocumentPath,
					EffectiveDate:   d.EffectiveDate,
					ExpiryDate:      d.ExpiryDate,
					IssueDate:       d.IssueDate,
					ReferenceNo:     d.ReferenceNo,
					RegistrationNo:  d.RegistrationNo,
					Remark:          d.Remark,
					Status:          pb.TransactionStatus_value["PendingForIssuing"],
					ThirdPartyID:    v.ThirdPartyID,
					ThirdPartyName:  v.ThirdPartyName,
					TransactionID:   d.TransactionId,
					UpdatedByID:     me.UserID,
				}

				transactionORM, err := s.provider.GetTransactionDetail(ctx, &pb.TransactionORM{ReferenceNo: d.ReferenceNo})
				if err != nil {
					if !errors.Is(err, gorm.ErrRecordNotFound) {
						return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
					}
				}

				if transactionORM.Id > 0 {
					t.Id = transactionORM.Id
				}

				transactionORM, err = s.provider.UpdateOrCreateTransaction(ctx, t)
				if err != nil {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				}

				transactionPB, err := transactionORM.ToPB(ctx)
				if err != nil {
					return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
				}

				result.Data = append(result.Data, &transactionPB)

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
