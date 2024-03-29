// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: api.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ErrorBodyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   bool   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Code    uint32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ErrorBodyResponse) Reset() {
	*x = ErrorBodyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorBodyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorBodyResponse) ProtoMessage() {}

func (x *ErrorBodyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorBodyResponse.ProtoReflect.Descriptor instead.
func (*ErrorBodyResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorBodyResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *ErrorBodyResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ErrorBodyResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0d, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x57, 0x0a, 0x11, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x32, 0xc3, 0x1a, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xbf, 0x01, 0x0a, 0x0f, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4b,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x92, 0x41, 0x30, 0x0a, 0x04, 0x52, 0x65, 0x61,
	0x64, 0x12, 0x0d, 0x47, 0x65, 0x74, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x6c, 0x69, 0x73, 0x74,
	0x1a, 0x19, 0x67, 0x65, 0x74, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x12, 0xde, 0x01, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x2f, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x92, 0x41, 0x3e, 0x0a, 0x04,
	0x52, 0x65, 0x61, 0x64, 0x12, 0x1a, 0x47, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x1a, 0x1a, 0x47, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0xcc, 0x01, 0x0a,
	0x0c, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x2b, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x61, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20,
	0x12, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x92, 0x41, 0x38, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x11, 0x47, 0x65, 0x74, 0x20, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x20, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x1a, 0x1d, 0x47, 0x65,
	0x74, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x20, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0xbd, 0x01, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x2a,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x52, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12,
	0x1d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x92, 0x41,
	0x2a, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x10, 0x47, 0x65, 0x74, 0x20, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x20, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x1a, 0x10, 0x47, 0x65, 0x74, 0x20, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x20, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0xd6, 0x01, 0x0a, 0x11,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2d, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x31, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x5f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x7d, 0x92, 0x41, 0x34,
	0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x0f, 0x47, 0x65, 0x74, 0x20, 0x44, 0x61, 0x74, 0x61,
	0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a, 0x1b, 0x47, 0x65, 0x74, 0x20, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x12, 0xcc, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4b,
	0x65, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x26, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x60, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x2a, 0x20, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x7d, 0x92, 0x41, 0x35, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x17, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x64,
	0x61, 0x74, 0x61, 0x12, 0xc6, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x31, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x52, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15,
	0x22, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x34, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x6e, 0x65, 0x77,
	0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xc2, 0x01, 0x0a,
	0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2a, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x64, 0x69, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x31,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x4e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x1a, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x92, 0x41,
	0x30, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x12, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x7d, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x54,
	0x6f, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2f, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x54, 0x6f, 0x48, 0x6f, 0x73, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x32, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x48, 0x6f, 0x73, 0x74,
	0x54, 0x6f, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x7b, 0x0a, 0x14, 0x45, 0x64, 0x69, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x54, 0x6f, 0x48, 0x6f,
	0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2f, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x54, 0x6f, 0x48, 0x6f, 0x73, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x32, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x54, 0x6f, 0x48,
	0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x80, 0x01,
	0x0a, 0x16, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x54, 0x6f, 0x48, 0x6f,
	0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x32, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x54, 0x6f, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x4b, 0x65, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x32, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x48, 0x6f, 0x73, 0x74,
	0x54, 0x6f, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x80, 0x01, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x54,
	0x6f, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x32, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x54, 0x6f, 0x48, 0x6f, 0x73, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4b, 0x65, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x32, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x48,
	0x6f, 0x73, 0x74, 0x54, 0x6f, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x12, 0xe8, 0x01, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x12, 0x1d, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x33, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x22, 0x78, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x92, 0x41, 0x53, 0x0a, 0x04, 0x52, 0x65, 0x61,
	0x64, 0x12, 0x26, 0x47, 0x65, 0x74, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x6c, 0x69, 0x73, 0x74,
	0x20, 0x6f, 0x66, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x1a, 0x23, 0x67, 0x65, 0x74, 0x20, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x12, 0xd1,
	0x01, 0x0a, 0x0d, 0x45, 0x78, 0x65, 0x63, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x12, 0x28, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x26, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x6e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x22, 0x1a, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x46, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x20, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x20, 0x69, 0x6d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74,
	0x65, 0x6c, 0x79, 0x1a, 0x1d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x20, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x20, 0x69, 0x6d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65,
	0x6c, 0x79, 0x12, 0xa7, 0x02, 0x0a, 0x16, 0x45, 0x78, 0x65, 0x63, 0x55, 0x6e, 0x73, 0x65, 0x74,
	0x74, 0x6c, 0x65, 0x64, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x12, 0x31, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x55, 0x6e, 0x73, 0x65, 0x74,
	0x74, 0x6c, 0x65, 0x64, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x1a, 0x26, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xb1, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x2c, 0x12, 0x2a, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x75, 0x6e, 0x73, 0x65, 0x74, 0x74,
	0x6c, 0x65, 0x64, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x92, 0x41, 0x7c,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x27, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x20, 0x55, 0x6e, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x64, 0x20, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x20, 0x69, 0x6d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x6c,
	0x79, 0x1a, 0x49, 0x55, 0x6e, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x64, 0x20, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x20, 0x68, 0x61, 0x73, 0x20, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x20, 0x3d, 0x20, 0x30, 0x20, 0x28, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x29, 0x20,
	0x6f, 0x72, 0x20, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x20, 0x3d, 0x20, 0x32, 0x20, 0x28, 0x66,
	0x61, 0x69, 0x6c, 0x65, 0x64, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x29, 0x12, 0x85, 0x02, 0x0a,
	0x1a, 0x42, 0x52, 0x49, 0x47, 0x61, 0x74, 0x65, 0x48, 0x61, 0x72, 0x64, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x52, 0x49, 0x47, 0x61, 0x74, 0x65, 0x48, 0x61, 0x72, 0x64,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x52, 0x49, 0x47, 0x61, 0x74, 0x65, 0x48, 0x61, 0x72, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x70, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x22, 0x26, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x68, 0x61, 0x72,
	0x64, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x3c, 0x0a, 0x03, 0x41, 0x6c, 0x6c, 0x0a, 0x08, 0x45,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x0a, 0x07, 0x42, 0x52, 0x49, 0x47, 0x61, 0x74, 0x65,
	0x12, 0x1d, 0x42, 0x52, 0x49, 0x47, 0x41, 0x54, 0x45, 0x20, 0x48, 0x41, 0x52, 0x44, 0x20, 0x54,
	0x4f, 0x4b, 0x45, 0x4e, 0x20, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x1a,
	0x01, 0x2d, 0x62, 0x00, 0x12, 0xd1, 0x01, 0x0a, 0x12, 0x42, 0x52, 0x49, 0x47, 0x61, 0x74, 0x65,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x69, 0x72, 0x52, 0x61, 0x74, 0x65, 0x12, 0x31, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x52, 0x49, 0x47, 0x61, 0x74, 0x65, 0x47, 0x65, 0x74, 0x50,
	0x61, 0x69, 0x72, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x52, 0x49, 0x47, 0x61, 0x74, 0x65, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x69, 0x72, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x54, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x61, 0x74,
	0x65, 0x92, 0x41, 0x34, 0x0a, 0x03, 0x41, 0x6c, 0x6c, 0x0a, 0x08, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x0a, 0x07, 0x42, 0x52, 0x49, 0x47, 0x61, 0x74, 0x65, 0x12, 0x15, 0x42, 0x52,
	0x49, 0x47, 0x41, 0x54, 0x45, 0x20, 0x47, 0x45, 0x54, 0x20, 0x50, 0x41, 0x49, 0x52, 0x20, 0x52,
	0x41, 0x54, 0x45, 0x1a, 0x01, 0x2d, 0x62, 0x00, 0x42, 0xb6, 0x01, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x92, 0x41, 0xac, 0x01, 0x12, 0x1a, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x32, 0x03, 0x30, 0x2e,
	0x31, 0x2a, 0x01, 0x01, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x59, 0x0a, 0x57, 0x0a, 0x06, 0x62, 0x65,
	0x61, 0x72, 0x65, 0x72, 0x12, 0x4d, 0x08, 0x02, 0x12, 0x38, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2c, 0x20,
	0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x42, 0x65, 0x61, 0x72,
	0x65, 0x72, 0x3a, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x3c, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x3e, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x20, 0x02, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12,
	0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_proto_goTypes = []interface{}{
	(*ErrorBodyResponse)(nil),                  // 0: transaction.service.v1.errorBodyResponse
	(*Empty)(nil),                              // 1: transaction.service.v1.Empty
	(*ListTransactionRequest)(nil),             // 2: transaction.service.v1.ListTransactionRequest
	(*TransactionLimitRequest)(nil),            // 3: transaction.service.v1.TransactionLimitRequest
	(*LimitSummaryRequest)(nil),                // 4: transaction.service.v1.LimitSummaryRequest
	(*ModuleLimitRequest)(nil),                 // 5: transaction.service.v1.ModuleLimitRequest
	(*TransactionKeyMessage)(nil),              // 6: transaction.service.v1.TransactionKeyMessage
	(*TransactionMessage)(nil),                 // 7: transaction.service.v1.TransactionMessage
	(*EditTransactionReq)(nil),                 // 8: transaction.service.v1.EditTransactionReq
	(*HostToHostConfigMessage)(nil),            // 9: transaction.service.v1.HostToHostConfigMessage
	(*EditHostToHostConfigReq)(nil),            // 10: transaction.service.v1.EditHostToHostConfigReq
	(*HostToHostConfigKeyMessage)(nil),         // 11: transaction.service.v1.HostToHostConfigKeyMessage
	(*ExecSchedulerReq)(nil),                   // 12: transaction.service.v1.ExecSchedulerReq
	(*ExecUnsettledSchedulerReq)(nil),          // 13: transaction.service.v1.ExecUnsettledSchedulerReq
	(*BRIGateHardTokenValidationRequest)(nil),  // 14: transaction.service.v1.BRIGateHardTokenValidationRequest
	(*BRIGateGetPairRateRequest)(nil),          // 15: transaction.service.v1.BRIGateGetPairRateRequest
	(*ListTransactionResponse)(nil),            // 16: transaction.service.v1.ListTransactionResponse
	(*TransactionLimitResponse)(nil),           // 17: transaction.service.v1.TransactionLimitResponse
	(*LimitSummaryResponse)(nil),               // 18: transaction.service.v1.LimitSummaryResponse
	(*ModuleLimitResponse)(nil),                // 19: transaction.service.v1.ModuleLimitResponse
	(*DetailTransactionResponse)(nil),          // 20: transaction.service.v1.DetailTransactionResponse
	(*CommonResponse)(nil),                     // 21: transaction.service.v1.CommonResponse
	(*CreateTransactionResponse)(nil),          // 22: transaction.service.v1.CreateTransactionResponse
	(*DetailHostToHostConfigResp)(nil),         // 23: transaction.service.v1.DetailHostToHostConfigResp
	(*ListTransactionSchedulerRes)(nil),        // 24: transaction.service.v1.ListTransactionSchedulerRes
	(*BRIGateHardTokenValidationResponse)(nil), // 25: transaction.service.v1.BRIGateHardTokenValidationResponse
	(*BRIGateGetPairRateResponse)(nil),         // 26: transaction.service.v1.BRIGateGetPairRateResponse
}
var file_api_proto_depIdxs = []int32{
	2,  // 0: transaction.service.v1.TransactionService.ListTransaction:input_type -> transaction.service.v1.ListTransactionRequest
	3,  // 1: transaction.service.v1.TransactionService.GetTransactionLimit:input_type -> transaction.service.v1.TransactionLimitRequest
	4,  // 2: transaction.service.v1.TransactionService.LimitSummary:input_type -> transaction.service.v1.LimitSummaryRequest
	5,  // 3: transaction.service.v1.TransactionService.GetModuleLimit:input_type -> transaction.service.v1.ModuleLimitRequest
	6,  // 4: transaction.service.v1.TransactionService.DetailTransaction:input_type -> transaction.service.v1.TransactionKeyMessage
	6,  // 5: transaction.service.v1.TransactionService.DeleteTransaction:input_type -> transaction.service.v1.TransactionKeyMessage
	7,  // 6: transaction.service.v1.TransactionService.CreateTransaction:input_type -> transaction.service.v1.TransactionMessage
	8,  // 7: transaction.service.v1.TransactionService.UpdateTransaction:input_type -> transaction.service.v1.EditTransactionReq
	9,  // 8: transaction.service.v1.TransactionService.CreateHostToHostConfig:input_type -> transaction.service.v1.HostToHostConfigMessage
	10, // 9: transaction.service.v1.TransactionService.EditHostToHostConfig:input_type -> transaction.service.v1.EditHostToHostConfigReq
	11, // 10: transaction.service.v1.TransactionService.DetailHostToHostConfig:input_type -> transaction.service.v1.HostToHostConfigKeyMessage
	11, // 11: transaction.service.v1.TransactionService.DeleteHostToHostConfig:input_type -> transaction.service.v1.HostToHostConfigKeyMessage
	1,  // 12: transaction.service.v1.TransactionService.ListTransactionScheduler:input_type -> transaction.service.v1.Empty
	12, // 13: transaction.service.v1.TransactionService.ExecScheduler:input_type -> transaction.service.v1.ExecSchedulerReq
	13, // 14: transaction.service.v1.TransactionService.ExecUnsettledScheduler:input_type -> transaction.service.v1.ExecUnsettledSchedulerReq
	14, // 15: transaction.service.v1.TransactionService.BRIGateHardTokenValidation:input_type -> transaction.service.v1.BRIGateHardTokenValidationRequest
	15, // 16: transaction.service.v1.TransactionService.BRIGateGetPairRate:input_type -> transaction.service.v1.BRIGateGetPairRateRequest
	16, // 17: transaction.service.v1.TransactionService.ListTransaction:output_type -> transaction.service.v1.ListTransactionResponse
	17, // 18: transaction.service.v1.TransactionService.GetTransactionLimit:output_type -> transaction.service.v1.TransactionLimitResponse
	18, // 19: transaction.service.v1.TransactionService.LimitSummary:output_type -> transaction.service.v1.LimitSummaryResponse
	19, // 20: transaction.service.v1.TransactionService.GetModuleLimit:output_type -> transaction.service.v1.ModuleLimitResponse
	20, // 21: transaction.service.v1.TransactionService.DetailTransaction:output_type -> transaction.service.v1.DetailTransactionResponse
	21, // 22: transaction.service.v1.TransactionService.DeleteTransaction:output_type -> transaction.service.v1.CommonResponse
	22, // 23: transaction.service.v1.TransactionService.CreateTransaction:output_type -> transaction.service.v1.CreateTransactionResponse
	20, // 24: transaction.service.v1.TransactionService.UpdateTransaction:output_type -> transaction.service.v1.DetailTransactionResponse
	23, // 25: transaction.service.v1.TransactionService.CreateHostToHostConfig:output_type -> transaction.service.v1.DetailHostToHostConfigResp
	23, // 26: transaction.service.v1.TransactionService.EditHostToHostConfig:output_type -> transaction.service.v1.DetailHostToHostConfigResp
	23, // 27: transaction.service.v1.TransactionService.DetailHostToHostConfig:output_type -> transaction.service.v1.DetailHostToHostConfigResp
	23, // 28: transaction.service.v1.TransactionService.DeleteHostToHostConfig:output_type -> transaction.service.v1.DetailHostToHostConfigResp
	24, // 29: transaction.service.v1.TransactionService.ListTransactionScheduler:output_type -> transaction.service.v1.ListTransactionSchedulerRes
	21, // 30: transaction.service.v1.TransactionService.ExecScheduler:output_type -> transaction.service.v1.CommonResponse
	21, // 31: transaction.service.v1.TransactionService.ExecUnsettledScheduler:output_type -> transaction.service.v1.CommonResponse
	25, // 32: transaction.service.v1.TransactionService.BRIGateHardTokenValidation:output_type -> transaction.service.v1.BRIGateHardTokenValidationResponse
	26, // 33: transaction.service.v1.TransactionService.BRIGateGetPairRate:output_type -> transaction.service.v1.BRIGateGetPairRateResponse
	17, // [17:34] is the sub-list for method output_type
	0,  // [0:17] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	file_payload_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorBodyResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
