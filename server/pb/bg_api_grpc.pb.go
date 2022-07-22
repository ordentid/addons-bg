// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: bg_api.proto

package pb

import (
	context "context"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ApiServiceClient is the client API for ApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiServiceClient interface {
	HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	GetBeneficiaryName(ctx context.Context, in *GetBeneficiaryNameRequest, opts ...grpc.CallOption) (*GetBeneficiaryNameResponse, error)
	GetApplicantName(ctx context.Context, in *GetApplicantNameRequest, opts ...grpc.CallOption) (*GetApplicantNameResponse, error)
	GetThirdParty(ctx context.Context, in *GetThirdPartyRequest, opts ...grpc.CallOption) (*GetThirdPartyResponse, error)
	GenerateThirdParty(ctx context.Context, in *GenerateThirdPartyRequest, opts ...grpc.CallOption) (*GenerateThirdPartyResponse, error)
	GetTaskMappingFilterCompany(ctx context.Context, in *GetTaskMappingFilterCompanyRequest, opts ...grpc.CallOption) (*GetTaskMappingFilterCompanyResponse, error)
	GetTaskMappingFile(ctx context.Context, in *GetTaskMappingFileRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	GetTaskMapping(ctx context.Context, in *GetTaskMappingRequest, opts ...grpc.CallOption) (*GetTaskMappingResponse, error)
	GetTaskMappingDetail(ctx context.Context, in *GetTaskMappingDetailRequest, opts ...grpc.CallOption) (*GetTaskMappingDetailResponse, error)
	CreateTaskMapping(ctx context.Context, in *CreateTaskMappingRequest, opts ...grpc.CallOption) (*CreateTaskMappingResponse, error)
	GetTaskMappingDigitalFile(ctx context.Context, in *GetTaskMappingDigitalFileRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	GetTaskMappingDigital(ctx context.Context, in *GetTaskMappingDigitalRequest, opts ...grpc.CallOption) (*GetTaskMappingDigitalResponse, error)
	GetTaskMappingDigitalDetail(ctx context.Context, in *GetTaskMappingDigitalDetailRequest, opts ...grpc.CallOption) (*GetTaskMappingDigitalDetailResponse, error)
	CreateTaskMappingDigital(ctx context.Context, in *CreateTaskMappingDigitalRequest, opts ...grpc.CallOption) (*CreateTaskMappingDigitalResponse, error)
	GetTransactionFile(ctx context.Context, in *GetTransactionFileRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*GetTransactionResponse, error)
	GetTransactionDetail(ctx context.Context, in *GetTransactionDetailRequest, opts ...grpc.CallOption) (*GetTransactionDetailResponse, error)
	CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error)
	UpdateTransaction(ctx context.Context, in *UpdateTransactionRequest, opts ...grpc.CallOption) (*UpdateTransactionResponse, error)
	GetTaskIssuing(ctx context.Context, in *GetTaskIssuingRequest, opts ...grpc.CallOption) (*GetTaskIssuingResponse, error)
	GetTaskIssuingDetail(ctx context.Context, in *GetTaskIssuingDetailRequest, opts ...grpc.CallOption) (*GetTaskIssuingDetailResponse, error)
	CreateTaskIssuing(ctx context.Context, in *CreateTaskIssuingRequest, opts ...grpc.CallOption) (*CreateTaskIssuingResponse, error)
}

type apiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiServiceClient(cc grpc.ClientConnInterface) ApiServiceClient {
	return &apiServiceClient{cc}
}

func (c *apiServiceClient) HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/bg.service.v1.ApiService/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetBeneficiaryName(ctx context.Context, in *GetBeneficiaryNameRequest, opts ...grpc.CallOption) (*GetBeneficiaryNameResponse, error) {
	out := new(GetBeneficiaryNameResponse)
	err := c.cc.Invoke(ctx, "/bg.service.v1.ApiService/GetBeneficiaryName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetApplicantName(ctx context.Context, in *GetApplicantNameRequest, opts ...grpc.CallOption) (*GetApplicantNameResponse, error) {
	out := new(GetApplicantNameResponse)
	err := c.cc.Invoke(ctx, "/bg.service.v1.ApiService/GetApplicantName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetThirdParty(ctx context.Context, in *GetThirdPartyRequest, opts ...grpc.CallOption) (*GetThirdPartyResponse, error) {
	out := new(GetThirdPartyResponse)
	err := c.cc.Invoke(ctx, "/bg.service.v1.ApiService/GetThirdParty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GenerateThirdParty(ctx context.Context, in *GenerateThirdPartyRequest, opts ...grpc.CallOption) (*GenerateThirdPartyResponse, error) {
	out := new(GenerateThirdPartyResponse)
	err := c.cc.Invoke(ctx, "/bg.service.v1.ApiService/GenerateThirdParty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetTaskMappingFilterCompany(ctx context.Context, in *GetTaskMappingFilterCompanyRequest, opts ...grpc.CallOption) (*GetTaskMappingFilterCompanyResponse, error) {
	out := new(GetTaskMappingFilterCompanyResponse)
	err := c.cc.Invoke(ctx, "/bg.service.v1.ApiService/GetTaskMappingFilterCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetTaskMappingFile(ctx context.Context, in *GetTaskMappingFileRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/bg.service.v1.ApiService/GetTaskMappingFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetTaskMapping(ctx context.Context, in *GetTaskMappingRequest, opts ...grpc.CallOption) (*GetTaskMappingResponse, error) {
	out := new(GetTaskMappingResponse)
	err := c.cc.Invoke(ctx, "/bg.service.v1.ApiService/GetTaskMapping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetTaskMappingDetail(ctx context.Context, in *GetTaskMappingDetailRequest, opts ...grpc.CallOption) (*GetTaskMappingDetailResponse, error) {
	out := new(GetTaskMappingDetailResponse)
	err := c.cc.Invoke(ctx, "/bg.service.v1.ApiService/GetTaskMappingDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateTaskMapping(ctx context.Context, in *CreateTaskMappingRequest, opts ...grpc.CallOption) (*CreateTaskMappingResponse, error) {
	out := new(CreateTaskMappingResponse)
	err := c.cc.Invoke(ctx, "/bg.service.v1.ApiService/CreateTaskMapping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetTaskMappingDigitalFile(ctx context.Context, in *GetTaskMappingDigitalFileRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/bg.service.v1.ApiService/GetTaskMappingDigitalFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetTaskMappingDigital(ctx context.Context, in *GetTaskMappingDigitalRequest, opts ...grpc.CallOption) (*GetTaskMappingDigitalResponse, error) {
	out := new(GetTaskMappingDigitalResponse)
	err := c.cc.Invoke(ctx, "/bg.service.v1.ApiService/GetTaskMappingDigital", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetTaskMappingDigitalDetail(ctx context.Context, in *GetTaskMappingDigitalDetailRequest, opts ...grpc.CallOption) (*GetTaskMappingDigitalDetailResponse, error) {
	out := new(GetTaskMappingDigitalDetailResponse)
	err := c.cc.Invoke(ctx, "/bg.service.v1.ApiService/GetTaskMappingDigitalDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateTaskMappingDigital(ctx context.Context, in *CreateTaskMappingDigitalRequest, opts ...grpc.CallOption) (*CreateTaskMappingDigitalResponse, error) {
	out := new(CreateTaskMappingDigitalResponse)
	err := c.cc.Invoke(ctx, "/bg.service.v1.ApiService/CreateTaskMappingDigital", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetTransactionFile(ctx context.Context, in *GetTransactionFileRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/bg.service.v1.ApiService/GetTransactionFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*GetTransactionResponse, error) {
	out := new(GetTransactionResponse)
	err := c.cc.Invoke(ctx, "/bg.service.v1.ApiService/GetTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetTransactionDetail(ctx context.Context, in *GetTransactionDetailRequest, opts ...grpc.CallOption) (*GetTransactionDetailResponse, error) {
	out := new(GetTransactionDetailResponse)
	err := c.cc.Invoke(ctx, "/bg.service.v1.ApiService/GetTransactionDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error) {
	out := new(CreateTransactionResponse)
	err := c.cc.Invoke(ctx, "/bg.service.v1.ApiService/CreateTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) UpdateTransaction(ctx context.Context, in *UpdateTransactionRequest, opts ...grpc.CallOption) (*UpdateTransactionResponse, error) {
	out := new(UpdateTransactionResponse)
	err := c.cc.Invoke(ctx, "/bg.service.v1.ApiService/UpdateTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetTaskIssuing(ctx context.Context, in *GetTaskIssuingRequest, opts ...grpc.CallOption) (*GetTaskIssuingResponse, error) {
	out := new(GetTaskIssuingResponse)
	err := c.cc.Invoke(ctx, "/bg.service.v1.ApiService/GetTaskIssuing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetTaskIssuingDetail(ctx context.Context, in *GetTaskIssuingDetailRequest, opts ...grpc.CallOption) (*GetTaskIssuingDetailResponse, error) {
	out := new(GetTaskIssuingDetailResponse)
	err := c.cc.Invoke(ctx, "/bg.service.v1.ApiService/GetTaskIssuingDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateTaskIssuing(ctx context.Context, in *CreateTaskIssuingRequest, opts ...grpc.CallOption) (*CreateTaskIssuingResponse, error) {
	out := new(CreateTaskIssuingResponse)
	err := c.cc.Invoke(ctx, "/bg.service.v1.ApiService/CreateTaskIssuing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServiceServer is the server API for ApiService service.
// All implementations must embed UnimplementedApiServiceServer
// for forward compatibility
type ApiServiceServer interface {
	HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	GetBeneficiaryName(context.Context, *GetBeneficiaryNameRequest) (*GetBeneficiaryNameResponse, error)
	GetApplicantName(context.Context, *GetApplicantNameRequest) (*GetApplicantNameResponse, error)
	GetThirdParty(context.Context, *GetThirdPartyRequest) (*GetThirdPartyResponse, error)
	GenerateThirdParty(context.Context, *GenerateThirdPartyRequest) (*GenerateThirdPartyResponse, error)
	GetTaskMappingFilterCompany(context.Context, *GetTaskMappingFilterCompanyRequest) (*GetTaskMappingFilterCompanyResponse, error)
	GetTaskMappingFile(context.Context, *GetTaskMappingFileRequest) (*httpbody.HttpBody, error)
	GetTaskMapping(context.Context, *GetTaskMappingRequest) (*GetTaskMappingResponse, error)
	GetTaskMappingDetail(context.Context, *GetTaskMappingDetailRequest) (*GetTaskMappingDetailResponse, error)
	CreateTaskMapping(context.Context, *CreateTaskMappingRequest) (*CreateTaskMappingResponse, error)
	GetTaskMappingDigitalFile(context.Context, *GetTaskMappingDigitalFileRequest) (*httpbody.HttpBody, error)
	GetTaskMappingDigital(context.Context, *GetTaskMappingDigitalRequest) (*GetTaskMappingDigitalResponse, error)
	GetTaskMappingDigitalDetail(context.Context, *GetTaskMappingDigitalDetailRequest) (*GetTaskMappingDigitalDetailResponse, error)
	CreateTaskMappingDigital(context.Context, *CreateTaskMappingDigitalRequest) (*CreateTaskMappingDigitalResponse, error)
	GetTransactionFile(context.Context, *GetTransactionFileRequest) (*httpbody.HttpBody, error)
	GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionResponse, error)
	GetTransactionDetail(context.Context, *GetTransactionDetailRequest) (*GetTransactionDetailResponse, error)
	CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error)
	UpdateTransaction(context.Context, *UpdateTransactionRequest) (*UpdateTransactionResponse, error)
	GetTaskIssuing(context.Context, *GetTaskIssuingRequest) (*GetTaskIssuingResponse, error)
	GetTaskIssuingDetail(context.Context, *GetTaskIssuingDetailRequest) (*GetTaskIssuingDetailResponse, error)
	CreateTaskIssuing(context.Context, *CreateTaskIssuingRequest) (*CreateTaskIssuingResponse, error)
	mustEmbedUnimplementedApiServiceServer()
}

// UnimplementedApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApiServiceServer struct {
}

func (UnimplementedApiServiceServer) HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedApiServiceServer) GetBeneficiaryName(context.Context, *GetBeneficiaryNameRequest) (*GetBeneficiaryNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBeneficiaryName not implemented")
}
func (UnimplementedApiServiceServer) GetApplicantName(context.Context, *GetApplicantNameRequest) (*GetApplicantNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApplicantName not implemented")
}
func (UnimplementedApiServiceServer) GetThirdParty(context.Context, *GetThirdPartyRequest) (*GetThirdPartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetThirdParty not implemented")
}
func (UnimplementedApiServiceServer) GenerateThirdParty(context.Context, *GenerateThirdPartyRequest) (*GenerateThirdPartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateThirdParty not implemented")
}
func (UnimplementedApiServiceServer) GetTaskMappingFilterCompany(context.Context, *GetTaskMappingFilterCompanyRequest) (*GetTaskMappingFilterCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskMappingFilterCompany not implemented")
}
func (UnimplementedApiServiceServer) GetTaskMappingFile(context.Context, *GetTaskMappingFileRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskMappingFile not implemented")
}
func (UnimplementedApiServiceServer) GetTaskMapping(context.Context, *GetTaskMappingRequest) (*GetTaskMappingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskMapping not implemented")
}
func (UnimplementedApiServiceServer) GetTaskMappingDetail(context.Context, *GetTaskMappingDetailRequest) (*GetTaskMappingDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskMappingDetail not implemented")
}
func (UnimplementedApiServiceServer) CreateTaskMapping(context.Context, *CreateTaskMappingRequest) (*CreateTaskMappingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTaskMapping not implemented")
}
func (UnimplementedApiServiceServer) GetTaskMappingDigitalFile(context.Context, *GetTaskMappingDigitalFileRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskMappingDigitalFile not implemented")
}
func (UnimplementedApiServiceServer) GetTaskMappingDigital(context.Context, *GetTaskMappingDigitalRequest) (*GetTaskMappingDigitalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskMappingDigital not implemented")
}
func (UnimplementedApiServiceServer) GetTaskMappingDigitalDetail(context.Context, *GetTaskMappingDigitalDetailRequest) (*GetTaskMappingDigitalDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskMappingDigitalDetail not implemented")
}
func (UnimplementedApiServiceServer) CreateTaskMappingDigital(context.Context, *CreateTaskMappingDigitalRequest) (*CreateTaskMappingDigitalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTaskMappingDigital not implemented")
}
func (UnimplementedApiServiceServer) GetTransactionFile(context.Context, *GetTransactionFileRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionFile not implemented")
}
func (UnimplementedApiServiceServer) GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (UnimplementedApiServiceServer) GetTransactionDetail(context.Context, *GetTransactionDetailRequest) (*GetTransactionDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionDetail not implemented")
}
func (UnimplementedApiServiceServer) CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}
func (UnimplementedApiServiceServer) UpdateTransaction(context.Context, *UpdateTransactionRequest) (*UpdateTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTransaction not implemented")
}
func (UnimplementedApiServiceServer) GetTaskIssuing(context.Context, *GetTaskIssuingRequest) (*GetTaskIssuingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskIssuing not implemented")
}
func (UnimplementedApiServiceServer) GetTaskIssuingDetail(context.Context, *GetTaskIssuingDetailRequest) (*GetTaskIssuingDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskIssuingDetail not implemented")
}
func (UnimplementedApiServiceServer) CreateTaskIssuing(context.Context, *CreateTaskIssuingRequest) (*CreateTaskIssuingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTaskIssuing not implemented")
}
func (UnimplementedApiServiceServer) mustEmbedUnimplementedApiServiceServer() {}

// UnsafeApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServiceServer will
// result in compilation errors.
type UnsafeApiServiceServer interface {
	mustEmbedUnimplementedApiServiceServer()
}

func RegisterApiServiceServer(s grpc.ServiceRegistrar, srv ApiServiceServer) {
	s.RegisterService(&ApiService_ServiceDesc, srv)
}

func _ApiService_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bg.service.v1.ApiService/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).HealthCheck(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetBeneficiaryName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBeneficiaryNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetBeneficiaryName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bg.service.v1.ApiService/GetBeneficiaryName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetBeneficiaryName(ctx, req.(*GetBeneficiaryNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetApplicantName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApplicantNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetApplicantName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bg.service.v1.ApiService/GetApplicantName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetApplicantName(ctx, req.(*GetApplicantNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetThirdParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetThirdPartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetThirdParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bg.service.v1.ApiService/GetThirdParty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetThirdParty(ctx, req.(*GetThirdPartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GenerateThirdParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateThirdPartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GenerateThirdParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bg.service.v1.ApiService/GenerateThirdParty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GenerateThirdParty(ctx, req.(*GenerateThirdPartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetTaskMappingFilterCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskMappingFilterCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetTaskMappingFilterCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bg.service.v1.ApiService/GetTaskMappingFilterCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetTaskMappingFilterCompany(ctx, req.(*GetTaskMappingFilterCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetTaskMappingFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskMappingFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetTaskMappingFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bg.service.v1.ApiService/GetTaskMappingFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetTaskMappingFile(ctx, req.(*GetTaskMappingFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetTaskMapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetTaskMapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bg.service.v1.ApiService/GetTaskMapping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetTaskMapping(ctx, req.(*GetTaskMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetTaskMappingDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskMappingDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetTaskMappingDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bg.service.v1.ApiService/GetTaskMappingDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetTaskMappingDetail(ctx, req.(*GetTaskMappingDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateTaskMapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateTaskMapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bg.service.v1.ApiService/CreateTaskMapping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateTaskMapping(ctx, req.(*CreateTaskMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetTaskMappingDigitalFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskMappingDigitalFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetTaskMappingDigitalFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bg.service.v1.ApiService/GetTaskMappingDigitalFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetTaskMappingDigitalFile(ctx, req.(*GetTaskMappingDigitalFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetTaskMappingDigital_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskMappingDigitalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetTaskMappingDigital(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bg.service.v1.ApiService/GetTaskMappingDigital",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetTaskMappingDigital(ctx, req.(*GetTaskMappingDigitalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetTaskMappingDigitalDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskMappingDigitalDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetTaskMappingDigitalDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bg.service.v1.ApiService/GetTaskMappingDigitalDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetTaskMappingDigitalDetail(ctx, req.(*GetTaskMappingDigitalDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateTaskMappingDigital_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskMappingDigitalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateTaskMappingDigital(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bg.service.v1.ApiService/CreateTaskMappingDigital",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateTaskMappingDigital(ctx, req.(*CreateTaskMappingDigitalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetTransactionFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetTransactionFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bg.service.v1.ApiService/GetTransactionFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetTransactionFile(ctx, req.(*GetTransactionFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bg.service.v1.ApiService/GetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetTransaction(ctx, req.(*GetTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetTransactionDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetTransactionDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bg.service.v1.ApiService/GetTransactionDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetTransactionDetail(ctx, req.(*GetTransactionDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bg.service.v1.ApiService/CreateTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateTransaction(ctx, req.(*CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_UpdateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).UpdateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bg.service.v1.ApiService/UpdateTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).UpdateTransaction(ctx, req.(*UpdateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetTaskIssuing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskIssuingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetTaskIssuing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bg.service.v1.ApiService/GetTaskIssuing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetTaskIssuing(ctx, req.(*GetTaskIssuingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetTaskIssuingDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskIssuingDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetTaskIssuingDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bg.service.v1.ApiService/GetTaskIssuingDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetTaskIssuingDetail(ctx, req.(*GetTaskIssuingDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateTaskIssuing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskIssuingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateTaskIssuing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bg.service.v1.ApiService/CreateTaskIssuing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateTaskIssuing(ctx, req.(*CreateTaskIssuingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiService_ServiceDesc is the grpc.ServiceDesc for ApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bg.service.v1.ApiService",
	HandlerType: (*ApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _ApiService_HealthCheck_Handler,
		},
		{
			MethodName: "GetBeneficiaryName",
			Handler:    _ApiService_GetBeneficiaryName_Handler,
		},
		{
			MethodName: "GetApplicantName",
			Handler:    _ApiService_GetApplicantName_Handler,
		},
		{
			MethodName: "GetThirdParty",
			Handler:    _ApiService_GetThirdParty_Handler,
		},
		{
			MethodName: "GenerateThirdParty",
			Handler:    _ApiService_GenerateThirdParty_Handler,
		},
		{
			MethodName: "GetTaskMappingFilterCompany",
			Handler:    _ApiService_GetTaskMappingFilterCompany_Handler,
		},
		{
			MethodName: "GetTaskMappingFile",
			Handler:    _ApiService_GetTaskMappingFile_Handler,
		},
		{
			MethodName: "GetTaskMapping",
			Handler:    _ApiService_GetTaskMapping_Handler,
		},
		{
			MethodName: "GetTaskMappingDetail",
			Handler:    _ApiService_GetTaskMappingDetail_Handler,
		},
		{
			MethodName: "CreateTaskMapping",
			Handler:    _ApiService_CreateTaskMapping_Handler,
		},
		{
			MethodName: "GetTaskMappingDigitalFile",
			Handler:    _ApiService_GetTaskMappingDigitalFile_Handler,
		},
		{
			MethodName: "GetTaskMappingDigital",
			Handler:    _ApiService_GetTaskMappingDigital_Handler,
		},
		{
			MethodName: "GetTaskMappingDigitalDetail",
			Handler:    _ApiService_GetTaskMappingDigitalDetail_Handler,
		},
		{
			MethodName: "CreateTaskMappingDigital",
			Handler:    _ApiService_CreateTaskMappingDigital_Handler,
		},
		{
			MethodName: "GetTransactionFile",
			Handler:    _ApiService_GetTransactionFile_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _ApiService_GetTransaction_Handler,
		},
		{
			MethodName: "GetTransactionDetail",
			Handler:    _ApiService_GetTransactionDetail_Handler,
		},
		{
			MethodName: "CreateTransaction",
			Handler:    _ApiService_CreateTransaction_Handler,
		},
		{
			MethodName: "UpdateTransaction",
			Handler:    _ApiService_UpdateTransaction_Handler,
		},
		{
			MethodName: "GetTaskIssuing",
			Handler:    _ApiService_GetTaskIssuing_Handler,
		},
		{
			MethodName: "GetTaskIssuingDetail",
			Handler:    _ApiService_GetTaskIssuingDetail_Handler,
		},
		{
			MethodName: "CreateTaskIssuing",
			Handler:    _ApiService_CreateTaskIssuing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bg_api.proto",
}
