// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: api.proto

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

// TransactionServiceClient is the client API for TransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionServiceClient interface {
	ListTransaction(ctx context.Context, in *ListTransactionRequest, opts ...grpc.CallOption) (*ListTransactionResponse, error)
	GetTransactionLimit(ctx context.Context, in *TransactionLimitRequest, opts ...grpc.CallOption) (*TransactionLimitResponse, error)
	LimitSummary(ctx context.Context, in *LimitSummaryRequest, opts ...grpc.CallOption) (*LimitSummaryResponse, error)
	GetModuleLimit(ctx context.Context, in *ModuleLimitRequest, opts ...grpc.CallOption) (*ModuleLimitResponse, error)
	DetailTransaction(ctx context.Context, in *TransactionKeyMessage, opts ...grpc.CallOption) (*DetailTransactionResponse, error)
	DeleteTransaction(ctx context.Context, in *TransactionKeyMessage, opts ...grpc.CallOption) (*CommonResponse, error)
	CreateTransaction(ctx context.Context, in *TransactionMessage, opts ...grpc.CallOption) (*CreateTransactionResponse, error)
	UpdateTransaction(ctx context.Context, in *EditTransactionReq, opts ...grpc.CallOption) (*DetailTransactionResponse, error)
	CreateHostToHostConfig(ctx context.Context, in *HostToHostConfigMessage, opts ...grpc.CallOption) (*DetailHostToHostConfigResp, error)
	EditHostToHostConfig(ctx context.Context, in *EditHostToHostConfigReq, opts ...grpc.CallOption) (*DetailHostToHostConfigResp, error)
	DetailHostToHostConfig(ctx context.Context, in *HostToHostConfigKeyMessage, opts ...grpc.CallOption) (*DetailHostToHostConfigResp, error)
	DeleteHostToHostConfig(ctx context.Context, in *HostToHostConfigKeyMessage, opts ...grpc.CallOption) (*DetailHostToHostConfigResp, error)
	ListTransactionScheduler(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListTransactionSchedulerRes, error)
	ExecScheduler(ctx context.Context, in *ExecSchedulerReq, opts ...grpc.CallOption) (*CommonResponse, error)
	ExecUnsettledScheduler(ctx context.Context, in *ExecUnsettledSchedulerReq, opts ...grpc.CallOption) (*CommonResponse, error)
	BRIGateHardTokenValidation(ctx context.Context, in *BRIGateHardTokenValidationRequest, opts ...grpc.CallOption) (*BRIGateHardTokenValidationResponse, error)
	BRIGateGetPairRate(ctx context.Context, in *BRIGateGetPairRateRequest, opts ...grpc.CallOption) (*BRIGateGetPairRateResponse, error)
}

type transactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionServiceClient(cc grpc.ClientConnInterface) TransactionServiceClient {
	return &transactionServiceClient{cc}
}

func (c *transactionServiceClient) ListTransaction(ctx context.Context, in *ListTransactionRequest, opts ...grpc.CallOption) (*ListTransactionResponse, error) {
	out := new(ListTransactionResponse)
	err := c.cc.Invoke(ctx, "/transaction.service.v1.TransactionService/ListTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) GetTransactionLimit(ctx context.Context, in *TransactionLimitRequest, opts ...grpc.CallOption) (*TransactionLimitResponse, error) {
	out := new(TransactionLimitResponse)
	err := c.cc.Invoke(ctx, "/transaction.service.v1.TransactionService/GetTransactionLimit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) LimitSummary(ctx context.Context, in *LimitSummaryRequest, opts ...grpc.CallOption) (*LimitSummaryResponse, error) {
	out := new(LimitSummaryResponse)
	err := c.cc.Invoke(ctx, "/transaction.service.v1.TransactionService/LimitSummary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) GetModuleLimit(ctx context.Context, in *ModuleLimitRequest, opts ...grpc.CallOption) (*ModuleLimitResponse, error) {
	out := new(ModuleLimitResponse)
	err := c.cc.Invoke(ctx, "/transaction.service.v1.TransactionService/GetModuleLimit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) DetailTransaction(ctx context.Context, in *TransactionKeyMessage, opts ...grpc.CallOption) (*DetailTransactionResponse, error) {
	out := new(DetailTransactionResponse)
	err := c.cc.Invoke(ctx, "/transaction.service.v1.TransactionService/DetailTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) DeleteTransaction(ctx context.Context, in *TransactionKeyMessage, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/transaction.service.v1.TransactionService/DeleteTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) CreateTransaction(ctx context.Context, in *TransactionMessage, opts ...grpc.CallOption) (*CreateTransactionResponse, error) {
	out := new(CreateTransactionResponse)
	err := c.cc.Invoke(ctx, "/transaction.service.v1.TransactionService/CreateTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) UpdateTransaction(ctx context.Context, in *EditTransactionReq, opts ...grpc.CallOption) (*DetailTransactionResponse, error) {
	out := new(DetailTransactionResponse)
	err := c.cc.Invoke(ctx, "/transaction.service.v1.TransactionService/UpdateTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) CreateHostToHostConfig(ctx context.Context, in *HostToHostConfigMessage, opts ...grpc.CallOption) (*DetailHostToHostConfigResp, error) {
	out := new(DetailHostToHostConfigResp)
	err := c.cc.Invoke(ctx, "/transaction.service.v1.TransactionService/CreateHostToHostConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) EditHostToHostConfig(ctx context.Context, in *EditHostToHostConfigReq, opts ...grpc.CallOption) (*DetailHostToHostConfigResp, error) {
	out := new(DetailHostToHostConfigResp)
	err := c.cc.Invoke(ctx, "/transaction.service.v1.TransactionService/EditHostToHostConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) DetailHostToHostConfig(ctx context.Context, in *HostToHostConfigKeyMessage, opts ...grpc.CallOption) (*DetailHostToHostConfigResp, error) {
	out := new(DetailHostToHostConfigResp)
	err := c.cc.Invoke(ctx, "/transaction.service.v1.TransactionService/DetailHostToHostConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) DeleteHostToHostConfig(ctx context.Context, in *HostToHostConfigKeyMessage, opts ...grpc.CallOption) (*DetailHostToHostConfigResp, error) {
	out := new(DetailHostToHostConfigResp)
	err := c.cc.Invoke(ctx, "/transaction.service.v1.TransactionService/DeleteHostToHostConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) ListTransactionScheduler(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListTransactionSchedulerRes, error) {
	out := new(ListTransactionSchedulerRes)
	err := c.cc.Invoke(ctx, "/transaction.service.v1.TransactionService/ListTransactionScheduler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) ExecScheduler(ctx context.Context, in *ExecSchedulerReq, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/transaction.service.v1.TransactionService/ExecScheduler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) ExecUnsettledScheduler(ctx context.Context, in *ExecUnsettledSchedulerReq, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/transaction.service.v1.TransactionService/ExecUnsettledScheduler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) BRIGateHardTokenValidation(ctx context.Context, in *BRIGateHardTokenValidationRequest, opts ...grpc.CallOption) (*BRIGateHardTokenValidationResponse, error) {
	out := new(BRIGateHardTokenValidationResponse)
	err := c.cc.Invoke(ctx, "/transaction.service.v1.TransactionService/BRIGateHardTokenValidation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) BRIGateGetPairRate(ctx context.Context, in *BRIGateGetPairRateRequest, opts ...grpc.CallOption) (*BRIGateGetPairRateResponse, error) {
	out := new(BRIGateGetPairRateResponse)
	err := c.cc.Invoke(ctx, "/transaction.service.v1.TransactionService/BRIGateGetPairRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServiceServer is the server API for TransactionService service.
// All implementations must embed UnimplementedTransactionServiceServer
// for forward compatibility
type TransactionServiceServer interface {
	ListTransaction(context.Context, *ListTransactionRequest) (*ListTransactionResponse, error)
	GetTransactionLimit(context.Context, *TransactionLimitRequest) (*TransactionLimitResponse, error)
	LimitSummary(context.Context, *LimitSummaryRequest) (*LimitSummaryResponse, error)
	GetModuleLimit(context.Context, *ModuleLimitRequest) (*ModuleLimitResponse, error)
	DetailTransaction(context.Context, *TransactionKeyMessage) (*DetailTransactionResponse, error)
	DeleteTransaction(context.Context, *TransactionKeyMessage) (*CommonResponse, error)
	CreateTransaction(context.Context, *TransactionMessage) (*CreateTransactionResponse, error)
	UpdateTransaction(context.Context, *EditTransactionReq) (*DetailTransactionResponse, error)
	CreateHostToHostConfig(context.Context, *HostToHostConfigMessage) (*DetailHostToHostConfigResp, error)
	EditHostToHostConfig(context.Context, *EditHostToHostConfigReq) (*DetailHostToHostConfigResp, error)
	DetailHostToHostConfig(context.Context, *HostToHostConfigKeyMessage) (*DetailHostToHostConfigResp, error)
	DeleteHostToHostConfig(context.Context, *HostToHostConfigKeyMessage) (*DetailHostToHostConfigResp, error)
	ListTransactionScheduler(context.Context, *Empty) (*ListTransactionSchedulerRes, error)
	ExecScheduler(context.Context, *ExecSchedulerReq) (*CommonResponse, error)
	ExecUnsettledScheduler(context.Context, *ExecUnsettledSchedulerReq) (*CommonResponse, error)
	BRIGateHardTokenValidation(context.Context, *BRIGateHardTokenValidationRequest) (*BRIGateHardTokenValidationResponse, error)
	BRIGateGetPairRate(context.Context, *BRIGateGetPairRateRequest) (*BRIGateGetPairRateResponse, error)
	mustEmbedUnimplementedTransactionServiceServer()
}

// UnimplementedTransactionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransactionServiceServer struct {
}

func (UnimplementedTransactionServiceServer) ListTransaction(context.Context, *ListTransactionRequest) (*ListTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTransaction not implemented")
}
func (UnimplementedTransactionServiceServer) GetTransactionLimit(context.Context, *TransactionLimitRequest) (*TransactionLimitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionLimit not implemented")
}
func (UnimplementedTransactionServiceServer) LimitSummary(context.Context, *LimitSummaryRequest) (*LimitSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LimitSummary not implemented")
}
func (UnimplementedTransactionServiceServer) GetModuleLimit(context.Context, *ModuleLimitRequest) (*ModuleLimitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModuleLimit not implemented")
}
func (UnimplementedTransactionServiceServer) DetailTransaction(context.Context, *TransactionKeyMessage) (*DetailTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetailTransaction not implemented")
}
func (UnimplementedTransactionServiceServer) DeleteTransaction(context.Context, *TransactionKeyMessage) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTransaction not implemented")
}
func (UnimplementedTransactionServiceServer) CreateTransaction(context.Context, *TransactionMessage) (*CreateTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}
func (UnimplementedTransactionServiceServer) UpdateTransaction(context.Context, *EditTransactionReq) (*DetailTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTransaction not implemented")
}
func (UnimplementedTransactionServiceServer) CreateHostToHostConfig(context.Context, *HostToHostConfigMessage) (*DetailHostToHostConfigResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHostToHostConfig not implemented")
}
func (UnimplementedTransactionServiceServer) EditHostToHostConfig(context.Context, *EditHostToHostConfigReq) (*DetailHostToHostConfigResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditHostToHostConfig not implemented")
}
func (UnimplementedTransactionServiceServer) DetailHostToHostConfig(context.Context, *HostToHostConfigKeyMessage) (*DetailHostToHostConfigResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetailHostToHostConfig not implemented")
}
func (UnimplementedTransactionServiceServer) DeleteHostToHostConfig(context.Context, *HostToHostConfigKeyMessage) (*DetailHostToHostConfigResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHostToHostConfig not implemented")
}
func (UnimplementedTransactionServiceServer) ListTransactionScheduler(context.Context, *Empty) (*ListTransactionSchedulerRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTransactionScheduler not implemented")
}
func (UnimplementedTransactionServiceServer) ExecScheduler(context.Context, *ExecSchedulerReq) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecScheduler not implemented")
}
func (UnimplementedTransactionServiceServer) ExecUnsettledScheduler(context.Context, *ExecUnsettledSchedulerReq) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecUnsettledScheduler not implemented")
}
func (UnimplementedTransactionServiceServer) BRIGateHardTokenValidation(context.Context, *BRIGateHardTokenValidationRequest) (*BRIGateHardTokenValidationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BRIGateHardTokenValidation not implemented")
}
func (UnimplementedTransactionServiceServer) BRIGateGetPairRate(context.Context, *BRIGateGetPairRateRequest) (*BRIGateGetPairRateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BRIGateGetPairRate not implemented")
}
func (UnimplementedTransactionServiceServer) mustEmbedUnimplementedTransactionServiceServer() {}

// UnsafeTransactionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionServiceServer will
// result in compilation errors.
type UnsafeTransactionServiceServer interface {
	mustEmbedUnimplementedTransactionServiceServer()
}

func RegisterTransactionServiceServer(s grpc.ServiceRegistrar, srv TransactionServiceServer) {
	s.RegisterService(&TransactionService_ServiceDesc, srv)
}

func _TransactionService_ListTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).ListTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.service.v1.TransactionService/ListTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).ListTransaction(ctx, req.(*ListTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_GetTransactionLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GetTransactionLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.service.v1.TransactionService/GetTransactionLimit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GetTransactionLimit(ctx, req.(*TransactionLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_LimitSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LimitSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).LimitSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.service.v1.TransactionService/LimitSummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).LimitSummary(ctx, req.(*LimitSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_GetModuleLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModuleLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GetModuleLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.service.v1.TransactionService/GetModuleLimit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GetModuleLimit(ctx, req.(*ModuleLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_DetailTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionKeyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).DetailTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.service.v1.TransactionService/DetailTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).DetailTransaction(ctx, req.(*TransactionKeyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_DeleteTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionKeyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).DeleteTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.service.v1.TransactionService/DeleteTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).DeleteTransaction(ctx, req.(*TransactionKeyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.service.v1.TransactionService/CreateTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).CreateTransaction(ctx, req.(*TransactionMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_UpdateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditTransactionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).UpdateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.service.v1.TransactionService/UpdateTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).UpdateTransaction(ctx, req.(*EditTransactionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_CreateHostToHostConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostToHostConfigMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).CreateHostToHostConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.service.v1.TransactionService/CreateHostToHostConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).CreateHostToHostConfig(ctx, req.(*HostToHostConfigMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_EditHostToHostConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditHostToHostConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).EditHostToHostConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.service.v1.TransactionService/EditHostToHostConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).EditHostToHostConfig(ctx, req.(*EditHostToHostConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_DetailHostToHostConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostToHostConfigKeyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).DetailHostToHostConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.service.v1.TransactionService/DetailHostToHostConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).DetailHostToHostConfig(ctx, req.(*HostToHostConfigKeyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_DeleteHostToHostConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostToHostConfigKeyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).DeleteHostToHostConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.service.v1.TransactionService/DeleteHostToHostConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).DeleteHostToHostConfig(ctx, req.(*HostToHostConfigKeyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_ListTransactionScheduler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).ListTransactionScheduler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.service.v1.TransactionService/ListTransactionScheduler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).ListTransactionScheduler(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_ExecScheduler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecSchedulerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).ExecScheduler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.service.v1.TransactionService/ExecScheduler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).ExecScheduler(ctx, req.(*ExecSchedulerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_ExecUnsettledScheduler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecUnsettledSchedulerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).ExecUnsettledScheduler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.service.v1.TransactionService/ExecUnsettledScheduler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).ExecUnsettledScheduler(ctx, req.(*ExecUnsettledSchedulerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_BRIGateHardTokenValidation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BRIGateHardTokenValidationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).BRIGateHardTokenValidation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.service.v1.TransactionService/BRIGateHardTokenValidation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).BRIGateHardTokenValidation(ctx, req.(*BRIGateHardTokenValidationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_BRIGateGetPairRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BRIGateGetPairRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).BRIGateGetPairRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.service.v1.TransactionService/BRIGateGetPairRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).BRIGateGetPairRate(ctx, req.(*BRIGateGetPairRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionService_ServiceDesc is the grpc.ServiceDesc for TransactionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transaction.service.v1.TransactionService",
	HandlerType: (*TransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTransaction",
			Handler:    _TransactionService_ListTransaction_Handler,
		},
		{
			MethodName: "GetTransactionLimit",
			Handler:    _TransactionService_GetTransactionLimit_Handler,
		},
		{
			MethodName: "LimitSummary",
			Handler:    _TransactionService_LimitSummary_Handler,
		},
		{
			MethodName: "GetModuleLimit",
			Handler:    _TransactionService_GetModuleLimit_Handler,
		},
		{
			MethodName: "DetailTransaction",
			Handler:    _TransactionService_DetailTransaction_Handler,
		},
		{
			MethodName: "DeleteTransaction",
			Handler:    _TransactionService_DeleteTransaction_Handler,
		},
		{
			MethodName: "CreateTransaction",
			Handler:    _TransactionService_CreateTransaction_Handler,
		},
		{
			MethodName: "UpdateTransaction",
			Handler:    _TransactionService_UpdateTransaction_Handler,
		},
		{
			MethodName: "CreateHostToHostConfig",
			Handler:    _TransactionService_CreateHostToHostConfig_Handler,
		},
		{
			MethodName: "EditHostToHostConfig",
			Handler:    _TransactionService_EditHostToHostConfig_Handler,
		},
		{
			MethodName: "DetailHostToHostConfig",
			Handler:    _TransactionService_DetailHostToHostConfig_Handler,
		},
		{
			MethodName: "DeleteHostToHostConfig",
			Handler:    _TransactionService_DeleteHostToHostConfig_Handler,
		},
		{
			MethodName: "ListTransactionScheduler",
			Handler:    _TransactionService_ListTransactionScheduler_Handler,
		},
		{
			MethodName: "ExecScheduler",
			Handler:    _TransactionService_ExecScheduler_Handler,
		},
		{
			MethodName: "ExecUnsettledScheduler",
			Handler:    _TransactionService_ExecUnsettledScheduler_Handler,
		},
		{
			MethodName: "BRIGateHardTokenValidation",
			Handler:    _TransactionService_BRIGateHardTokenValidation_Handler,
		},
		{
			MethodName: "BRIGateGetPairRate",
			Handler:    _TransactionService_BRIGateGetPairRate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
