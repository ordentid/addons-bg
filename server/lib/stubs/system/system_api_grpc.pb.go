// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: system_api.proto

package pb

import (
	context "context"
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
	HealthCheck(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	CreateSystem(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	CreateSystemParam(ctx context.Context, in *CreateSystemRequest, opts ...grpc.CallOption) (*CreateTaskSystemResponse, error)
	GetMyTasksID(ctx context.Context, in *GetMyTasksIDRequest, opts ...grpc.CallOption) (*GetMyTasksIDResponse, error)
	GetMyTasks(ctx context.Context, in *SystemFilterRequest, opts ...grpc.CallOption) (*ListSystemResponse, error)
	ListSystemData(ctx context.Context, in *ListSystemDataRequest, opts ...grpc.CallOption) (*ListSystemDataResponse, error)
	SystemDetail(ctx context.Context, in *SystemDetailRequest, opts ...grpc.CallOption) (*SystemDetailResponse, error)
	SystemDetailByKey(ctx context.Context, in *SystemDetailByKeyRequest, opts ...grpc.CallOption) (*SystemDetailByKeyResponse, error)
	ListMdBiccc(ctx context.Context, in *ListMdBicccRequest, opts ...grpc.CallOption) (*ListMdBicccResponse, error)
	ListMdBenefType(ctx context.Context, in *ListMdBenefTypeRequest, opts ...grpc.CallOption) (*ListMdBenefTypeResponse, error)
	ListMdCountry(ctx context.Context, in *ListMdCountryRequest, opts ...grpc.CallOption) (*ListMdCountryResponse, error)
	ListMdEmploymentIndustry(ctx context.Context, in *ListMdEmploymentIndustryRequest, opts ...grpc.CallOption) (*ListMdEmploymentIndustryResponse, error)
	ListMdEmploymentPosition(ctx context.Context, in *ListMdEmploymentPositionRequest, opts ...grpc.CallOption) (*ListMdEmploymentPositionResponse, error)
	ListMdEmploymentStatus(ctx context.Context, in *ListMdEmploymentStatusRequest, opts ...grpc.CallOption) (*ListMdEmploymentStatusResponse, error)
	ListMdGender(ctx context.Context, in *ListMdGenderRequest, opts ...grpc.CallOption) (*ListMdGenderResponse, error)
	ListMdIdType(ctx context.Context, in *ListMdIdTypeRequest, opts ...grpc.CallOption) (*ListMdIdTypeResponse, error)
	ListMdKota(ctx context.Context, in *ListMdKotaRequest, opts ...grpc.CallOption) (*ListMdKotaResponse, error)
	ListMdProvince(ctx context.Context, in *ListMdProvinceRequest, opts ...grpc.CallOption) (*ListMdProvinceResponse, error)
	ListMdPurpose(ctx context.Context, in *ListMdPurposeRequest, opts ...grpc.CallOption) (*ListMdPurposeResponse, error)
	ListMdPurposeIntended(ctx context.Context, in *ListMdPurposeIntendedRequest, opts ...grpc.CallOption) (*ListMdPurposeIntendedResponse, error)
	ListMdRelationToSender(ctx context.Context, in *ListMdRelationToSenderRequest, opts ...grpc.CallOption) (*ListMdRelationToSenderResponse, error)
	ListMdBankKliring(ctx context.Context, in *ListMdBankKliringRequest, opts ...grpc.CallOption) (*ListMdBankKliringResponse, error)
	ListMdBranch(ctx context.Context, in *ListMdBranchRequest, opts ...grpc.CallOption) (*ListMdBranchResponse, error)
}

type apiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiServiceClient(cc grpc.ClientConnInterface) ApiServiceClient {
	return &apiServiceClient{cc}
}

func (c *apiServiceClient) HealthCheck(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/system.service.v1.ApiService/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateSystem(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/system.service.v1.ApiService/CreateSystem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateSystemParam(ctx context.Context, in *CreateSystemRequest, opts ...grpc.CallOption) (*CreateTaskSystemResponse, error) {
	out := new(CreateTaskSystemResponse)
	err := c.cc.Invoke(ctx, "/system.service.v1.ApiService/CreateSystemParam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetMyTasksID(ctx context.Context, in *GetMyTasksIDRequest, opts ...grpc.CallOption) (*GetMyTasksIDResponse, error) {
	out := new(GetMyTasksIDResponse)
	err := c.cc.Invoke(ctx, "/system.service.v1.ApiService/GetMyTasksID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetMyTasks(ctx context.Context, in *SystemFilterRequest, opts ...grpc.CallOption) (*ListSystemResponse, error) {
	out := new(ListSystemResponse)
	err := c.cc.Invoke(ctx, "/system.service.v1.ApiService/GetMyTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListSystemData(ctx context.Context, in *ListSystemDataRequest, opts ...grpc.CallOption) (*ListSystemDataResponse, error) {
	out := new(ListSystemDataResponse)
	err := c.cc.Invoke(ctx, "/system.service.v1.ApiService/ListSystemData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) SystemDetail(ctx context.Context, in *SystemDetailRequest, opts ...grpc.CallOption) (*SystemDetailResponse, error) {
	out := new(SystemDetailResponse)
	err := c.cc.Invoke(ctx, "/system.service.v1.ApiService/SystemDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) SystemDetailByKey(ctx context.Context, in *SystemDetailByKeyRequest, opts ...grpc.CallOption) (*SystemDetailByKeyResponse, error) {
	out := new(SystemDetailByKeyResponse)
	err := c.cc.Invoke(ctx, "/system.service.v1.ApiService/SystemDetailByKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListMdBiccc(ctx context.Context, in *ListMdBicccRequest, opts ...grpc.CallOption) (*ListMdBicccResponse, error) {
	out := new(ListMdBicccResponse)
	err := c.cc.Invoke(ctx, "/system.service.v1.ApiService/ListMdBiccc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListMdBenefType(ctx context.Context, in *ListMdBenefTypeRequest, opts ...grpc.CallOption) (*ListMdBenefTypeResponse, error) {
	out := new(ListMdBenefTypeResponse)
	err := c.cc.Invoke(ctx, "/system.service.v1.ApiService/ListMdBenefType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListMdCountry(ctx context.Context, in *ListMdCountryRequest, opts ...grpc.CallOption) (*ListMdCountryResponse, error) {
	out := new(ListMdCountryResponse)
	err := c.cc.Invoke(ctx, "/system.service.v1.ApiService/ListMdCountry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListMdEmploymentIndustry(ctx context.Context, in *ListMdEmploymentIndustryRequest, opts ...grpc.CallOption) (*ListMdEmploymentIndustryResponse, error) {
	out := new(ListMdEmploymentIndustryResponse)
	err := c.cc.Invoke(ctx, "/system.service.v1.ApiService/ListMdEmploymentIndustry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListMdEmploymentPosition(ctx context.Context, in *ListMdEmploymentPositionRequest, opts ...grpc.CallOption) (*ListMdEmploymentPositionResponse, error) {
	out := new(ListMdEmploymentPositionResponse)
	err := c.cc.Invoke(ctx, "/system.service.v1.ApiService/ListMdEmploymentPosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListMdEmploymentStatus(ctx context.Context, in *ListMdEmploymentStatusRequest, opts ...grpc.CallOption) (*ListMdEmploymentStatusResponse, error) {
	out := new(ListMdEmploymentStatusResponse)
	err := c.cc.Invoke(ctx, "/system.service.v1.ApiService/ListMdEmploymentStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListMdGender(ctx context.Context, in *ListMdGenderRequest, opts ...grpc.CallOption) (*ListMdGenderResponse, error) {
	out := new(ListMdGenderResponse)
	err := c.cc.Invoke(ctx, "/system.service.v1.ApiService/ListMdGender", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListMdIdType(ctx context.Context, in *ListMdIdTypeRequest, opts ...grpc.CallOption) (*ListMdIdTypeResponse, error) {
	out := new(ListMdIdTypeResponse)
	err := c.cc.Invoke(ctx, "/system.service.v1.ApiService/ListMdIdType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListMdKota(ctx context.Context, in *ListMdKotaRequest, opts ...grpc.CallOption) (*ListMdKotaResponse, error) {
	out := new(ListMdKotaResponse)
	err := c.cc.Invoke(ctx, "/system.service.v1.ApiService/ListMdKota", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListMdProvince(ctx context.Context, in *ListMdProvinceRequest, opts ...grpc.CallOption) (*ListMdProvinceResponse, error) {
	out := new(ListMdProvinceResponse)
	err := c.cc.Invoke(ctx, "/system.service.v1.ApiService/ListMdProvince", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListMdPurpose(ctx context.Context, in *ListMdPurposeRequest, opts ...grpc.CallOption) (*ListMdPurposeResponse, error) {
	out := new(ListMdPurposeResponse)
	err := c.cc.Invoke(ctx, "/system.service.v1.ApiService/ListMdPurpose", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListMdPurposeIntended(ctx context.Context, in *ListMdPurposeIntendedRequest, opts ...grpc.CallOption) (*ListMdPurposeIntendedResponse, error) {
	out := new(ListMdPurposeIntendedResponse)
	err := c.cc.Invoke(ctx, "/system.service.v1.ApiService/ListMdPurposeIntended", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListMdRelationToSender(ctx context.Context, in *ListMdRelationToSenderRequest, opts ...grpc.CallOption) (*ListMdRelationToSenderResponse, error) {
	out := new(ListMdRelationToSenderResponse)
	err := c.cc.Invoke(ctx, "/system.service.v1.ApiService/ListMdRelationToSender", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListMdBankKliring(ctx context.Context, in *ListMdBankKliringRequest, opts ...grpc.CallOption) (*ListMdBankKliringResponse, error) {
	out := new(ListMdBankKliringResponse)
	err := c.cc.Invoke(ctx, "/system.service.v1.ApiService/ListMdBankKliring", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListMdBranch(ctx context.Context, in *ListMdBranchRequest, opts ...grpc.CallOption) (*ListMdBranchResponse, error) {
	out := new(ListMdBranchResponse)
	err := c.cc.Invoke(ctx, "/system.service.v1.ApiService/ListMdBranch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServiceServer is the server API for ApiService service.
// All implementations must embed UnimplementedApiServiceServer
// for forward compatibility
type ApiServiceServer interface {
	HealthCheck(context.Context, *Empty) (*HealthCheckResponse, error)
	CreateSystem(context.Context, *CreateRequest) (*CreateResponse, error)
	CreateSystemParam(context.Context, *CreateSystemRequest) (*CreateTaskSystemResponse, error)
	GetMyTasksID(context.Context, *GetMyTasksIDRequest) (*GetMyTasksIDResponse, error)
	GetMyTasks(context.Context, *SystemFilterRequest) (*ListSystemResponse, error)
	ListSystemData(context.Context, *ListSystemDataRequest) (*ListSystemDataResponse, error)
	SystemDetail(context.Context, *SystemDetailRequest) (*SystemDetailResponse, error)
	SystemDetailByKey(context.Context, *SystemDetailByKeyRequest) (*SystemDetailByKeyResponse, error)
	ListMdBiccc(context.Context, *ListMdBicccRequest) (*ListMdBicccResponse, error)
	ListMdBenefType(context.Context, *ListMdBenefTypeRequest) (*ListMdBenefTypeResponse, error)
	ListMdCountry(context.Context, *ListMdCountryRequest) (*ListMdCountryResponse, error)
	ListMdEmploymentIndustry(context.Context, *ListMdEmploymentIndustryRequest) (*ListMdEmploymentIndustryResponse, error)
	ListMdEmploymentPosition(context.Context, *ListMdEmploymentPositionRequest) (*ListMdEmploymentPositionResponse, error)
	ListMdEmploymentStatus(context.Context, *ListMdEmploymentStatusRequest) (*ListMdEmploymentStatusResponse, error)
	ListMdGender(context.Context, *ListMdGenderRequest) (*ListMdGenderResponse, error)
	ListMdIdType(context.Context, *ListMdIdTypeRequest) (*ListMdIdTypeResponse, error)
	ListMdKota(context.Context, *ListMdKotaRequest) (*ListMdKotaResponse, error)
	ListMdProvince(context.Context, *ListMdProvinceRequest) (*ListMdProvinceResponse, error)
	ListMdPurpose(context.Context, *ListMdPurposeRequest) (*ListMdPurposeResponse, error)
	ListMdPurposeIntended(context.Context, *ListMdPurposeIntendedRequest) (*ListMdPurposeIntendedResponse, error)
	ListMdRelationToSender(context.Context, *ListMdRelationToSenderRequest) (*ListMdRelationToSenderResponse, error)
	ListMdBankKliring(context.Context, *ListMdBankKliringRequest) (*ListMdBankKliringResponse, error)
	ListMdBranch(context.Context, *ListMdBranchRequest) (*ListMdBranchResponse, error)
	mustEmbedUnimplementedApiServiceServer()
}

// UnimplementedApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApiServiceServer struct {
}

func (UnimplementedApiServiceServer) HealthCheck(context.Context, *Empty) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedApiServiceServer) CreateSystem(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSystem not implemented")
}
func (UnimplementedApiServiceServer) CreateSystemParam(context.Context, *CreateSystemRequest) (*CreateTaskSystemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSystemParam not implemented")
}
func (UnimplementedApiServiceServer) GetMyTasksID(context.Context, *GetMyTasksIDRequest) (*GetMyTasksIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyTasksID not implemented")
}
func (UnimplementedApiServiceServer) GetMyTasks(context.Context, *SystemFilterRequest) (*ListSystemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyTasks not implemented")
}
func (UnimplementedApiServiceServer) ListSystemData(context.Context, *ListSystemDataRequest) (*ListSystemDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSystemData not implemented")
}
func (UnimplementedApiServiceServer) SystemDetail(context.Context, *SystemDetailRequest) (*SystemDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemDetail not implemented")
}
func (UnimplementedApiServiceServer) SystemDetailByKey(context.Context, *SystemDetailByKeyRequest) (*SystemDetailByKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemDetailByKey not implemented")
}
func (UnimplementedApiServiceServer) ListMdBiccc(context.Context, *ListMdBicccRequest) (*ListMdBicccResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMdBiccc not implemented")
}
func (UnimplementedApiServiceServer) ListMdBenefType(context.Context, *ListMdBenefTypeRequest) (*ListMdBenefTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMdBenefType not implemented")
}
func (UnimplementedApiServiceServer) ListMdCountry(context.Context, *ListMdCountryRequest) (*ListMdCountryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMdCountry not implemented")
}
func (UnimplementedApiServiceServer) ListMdEmploymentIndustry(context.Context, *ListMdEmploymentIndustryRequest) (*ListMdEmploymentIndustryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMdEmploymentIndustry not implemented")
}
func (UnimplementedApiServiceServer) ListMdEmploymentPosition(context.Context, *ListMdEmploymentPositionRequest) (*ListMdEmploymentPositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMdEmploymentPosition not implemented")
}
func (UnimplementedApiServiceServer) ListMdEmploymentStatus(context.Context, *ListMdEmploymentStatusRequest) (*ListMdEmploymentStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMdEmploymentStatus not implemented")
}
func (UnimplementedApiServiceServer) ListMdGender(context.Context, *ListMdGenderRequest) (*ListMdGenderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMdGender not implemented")
}
func (UnimplementedApiServiceServer) ListMdIdType(context.Context, *ListMdIdTypeRequest) (*ListMdIdTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMdIdType not implemented")
}
func (UnimplementedApiServiceServer) ListMdKota(context.Context, *ListMdKotaRequest) (*ListMdKotaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMdKota not implemented")
}
func (UnimplementedApiServiceServer) ListMdProvince(context.Context, *ListMdProvinceRequest) (*ListMdProvinceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMdProvince not implemented")
}
func (UnimplementedApiServiceServer) ListMdPurpose(context.Context, *ListMdPurposeRequest) (*ListMdPurposeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMdPurpose not implemented")
}
func (UnimplementedApiServiceServer) ListMdPurposeIntended(context.Context, *ListMdPurposeIntendedRequest) (*ListMdPurposeIntendedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMdPurposeIntended not implemented")
}
func (UnimplementedApiServiceServer) ListMdRelationToSender(context.Context, *ListMdRelationToSenderRequest) (*ListMdRelationToSenderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMdRelationToSender not implemented")
}
func (UnimplementedApiServiceServer) ListMdBankKliring(context.Context, *ListMdBankKliringRequest) (*ListMdBankKliringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMdBankKliring not implemented")
}
func (UnimplementedApiServiceServer) ListMdBranch(context.Context, *ListMdBranchRequest) (*ListMdBranchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMdBranch not implemented")
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
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.service.v1.ApiService/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).HealthCheck(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateSystem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateSystem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.service.v1.ApiService/CreateSystem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateSystem(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateSystemParam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateSystemParam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.service.v1.ApiService/CreateSystemParam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateSystemParam(ctx, req.(*CreateSystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetMyTasksID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyTasksIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetMyTasksID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.service.v1.ApiService/GetMyTasksID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetMyTasksID(ctx, req.(*GetMyTasksIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetMyTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetMyTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.service.v1.ApiService/GetMyTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetMyTasks(ctx, req.(*SystemFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListSystemData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSystemDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListSystemData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.service.v1.ApiService/ListSystemData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListSystemData(ctx, req.(*ListSystemDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_SystemDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).SystemDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.service.v1.ApiService/SystemDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).SystemDetail(ctx, req.(*SystemDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_SystemDetailByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemDetailByKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).SystemDetailByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.service.v1.ApiService/SystemDetailByKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).SystemDetailByKey(ctx, req.(*SystemDetailByKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListMdBiccc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMdBicccRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListMdBiccc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.service.v1.ApiService/ListMdBiccc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListMdBiccc(ctx, req.(*ListMdBicccRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListMdBenefType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMdBenefTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListMdBenefType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.service.v1.ApiService/ListMdBenefType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListMdBenefType(ctx, req.(*ListMdBenefTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListMdCountry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMdCountryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListMdCountry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.service.v1.ApiService/ListMdCountry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListMdCountry(ctx, req.(*ListMdCountryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListMdEmploymentIndustry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMdEmploymentIndustryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListMdEmploymentIndustry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.service.v1.ApiService/ListMdEmploymentIndustry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListMdEmploymentIndustry(ctx, req.(*ListMdEmploymentIndustryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListMdEmploymentPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMdEmploymentPositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListMdEmploymentPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.service.v1.ApiService/ListMdEmploymentPosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListMdEmploymentPosition(ctx, req.(*ListMdEmploymentPositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListMdEmploymentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMdEmploymentStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListMdEmploymentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.service.v1.ApiService/ListMdEmploymentStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListMdEmploymentStatus(ctx, req.(*ListMdEmploymentStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListMdGender_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMdGenderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListMdGender(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.service.v1.ApiService/ListMdGender",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListMdGender(ctx, req.(*ListMdGenderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListMdIdType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMdIdTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListMdIdType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.service.v1.ApiService/ListMdIdType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListMdIdType(ctx, req.(*ListMdIdTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListMdKota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMdKotaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListMdKota(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.service.v1.ApiService/ListMdKota",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListMdKota(ctx, req.(*ListMdKotaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListMdProvince_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMdProvinceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListMdProvince(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.service.v1.ApiService/ListMdProvince",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListMdProvince(ctx, req.(*ListMdProvinceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListMdPurpose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMdPurposeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListMdPurpose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.service.v1.ApiService/ListMdPurpose",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListMdPurpose(ctx, req.(*ListMdPurposeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListMdPurposeIntended_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMdPurposeIntendedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListMdPurposeIntended(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.service.v1.ApiService/ListMdPurposeIntended",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListMdPurposeIntended(ctx, req.(*ListMdPurposeIntendedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListMdRelationToSender_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMdRelationToSenderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListMdRelationToSender(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.service.v1.ApiService/ListMdRelationToSender",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListMdRelationToSender(ctx, req.(*ListMdRelationToSenderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListMdBankKliring_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMdBankKliringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListMdBankKliring(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.service.v1.ApiService/ListMdBankKliring",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListMdBankKliring(ctx, req.(*ListMdBankKliringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListMdBranch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMdBranchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListMdBranch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.service.v1.ApiService/ListMdBranch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListMdBranch(ctx, req.(*ListMdBranchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiService_ServiceDesc is the grpc.ServiceDesc for ApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "system.service.v1.ApiService",
	HandlerType: (*ApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _ApiService_HealthCheck_Handler,
		},
		{
			MethodName: "CreateSystem",
			Handler:    _ApiService_CreateSystem_Handler,
		},
		{
			MethodName: "CreateSystemParam",
			Handler:    _ApiService_CreateSystemParam_Handler,
		},
		{
			MethodName: "GetMyTasksID",
			Handler:    _ApiService_GetMyTasksID_Handler,
		},
		{
			MethodName: "GetMyTasks",
			Handler:    _ApiService_GetMyTasks_Handler,
		},
		{
			MethodName: "ListSystemData",
			Handler:    _ApiService_ListSystemData_Handler,
		},
		{
			MethodName: "SystemDetail",
			Handler:    _ApiService_SystemDetail_Handler,
		},
		{
			MethodName: "SystemDetailByKey",
			Handler:    _ApiService_SystemDetailByKey_Handler,
		},
		{
			MethodName: "ListMdBiccc",
			Handler:    _ApiService_ListMdBiccc_Handler,
		},
		{
			MethodName: "ListMdBenefType",
			Handler:    _ApiService_ListMdBenefType_Handler,
		},
		{
			MethodName: "ListMdCountry",
			Handler:    _ApiService_ListMdCountry_Handler,
		},
		{
			MethodName: "ListMdEmploymentIndustry",
			Handler:    _ApiService_ListMdEmploymentIndustry_Handler,
		},
		{
			MethodName: "ListMdEmploymentPosition",
			Handler:    _ApiService_ListMdEmploymentPosition_Handler,
		},
		{
			MethodName: "ListMdEmploymentStatus",
			Handler:    _ApiService_ListMdEmploymentStatus_Handler,
		},
		{
			MethodName: "ListMdGender",
			Handler:    _ApiService_ListMdGender_Handler,
		},
		{
			MethodName: "ListMdIdType",
			Handler:    _ApiService_ListMdIdType_Handler,
		},
		{
			MethodName: "ListMdKota",
			Handler:    _ApiService_ListMdKota_Handler,
		},
		{
			MethodName: "ListMdProvince",
			Handler:    _ApiService_ListMdProvince_Handler,
		},
		{
			MethodName: "ListMdPurpose",
			Handler:    _ApiService_ListMdPurpose_Handler,
		},
		{
			MethodName: "ListMdPurposeIntended",
			Handler:    _ApiService_ListMdPurposeIntended_Handler,
		},
		{
			MethodName: "ListMdRelationToSender",
			Handler:    _ApiService_ListMdRelationToSender_Handler,
		},
		{
			MethodName: "ListMdBankKliring",
			Handler:    _ApiService_ListMdBankKliring_Handler,
		},
		{
			MethodName: "ListMdBranch",
			Handler:    _ApiService_ListMdBranch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "system_api.proto",
}
