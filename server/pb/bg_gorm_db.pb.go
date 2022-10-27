// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.16.1
// source: bg_gorm_db.proto

package pb

import (
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BgStatus int32

const (
	BgStatus_Cancelled     BgStatus = 0
	BgStatus_Active        BgStatus = 1
	BgStatus_ClaimPeriod   BgStatus = 2
	BgStatus_ClosingPeriod BgStatus = 3
	BgStatus_Closed        BgStatus = 4
)

// Enum value maps for BgStatus.
var (
	BgStatus_name = map[int32]string{
		0: "Cancelled",
		1: "Active",
		2: "ClaimPeriod",
		3: "ClosingPeriod",
		4: "Closed",
	}
	BgStatus_value = map[string]int32{
		"Cancelled":     0,
		"Active":        1,
		"ClaimPeriod":   2,
		"ClosingPeriod": 3,
		"Closed":        4,
	}
)

func (x BgStatus) Enum() *BgStatus {
	p := new(BgStatus)
	*p = x
	return p
}

func (x BgStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BgStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_bg_gorm_db_proto_enumTypes[0].Descriptor()
}

func (BgStatus) Type() protoreflect.EnumType {
	return &file_bg_gorm_db_proto_enumTypes[0]
}

func (x BgStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BgStatus.Descriptor instead.
func (BgStatus) EnumDescriptor() ([]byte, []int) {
	return file_bg_gorm_db_proto_rawDescGZIP(), []int{0}
}

type BgType int32

const (
	BgType_Null                       BgType = 0
	BgType_BidBond                    BgType = 1
	BgType_AdvancePayment             BgType = 2
	BgType_PerformanceBond            BgType = 3
	BgType_GovernmentPaymentGuarantee BgType = 4
	BgType_MaintenanceBond            BgType = 5
	BgType_ProcurementBond            BgType = 6
	BgType_TransactionRiskBond        BgType = 7
	BgType_CustomsBond                BgType = 8
)

// Enum value maps for BgType.
var (
	BgType_name = map[int32]string{
		0: "Null",
		1: "BidBond",
		2: "AdvancePayment",
		3: "PerformanceBond",
		4: "GovernmentPaymentGuarantee",
		5: "MaintenanceBond",
		6: "ProcurementBond",
		7: "TransactionRiskBond",
		8: "CustomsBond",
	}
	BgType_value = map[string]int32{
		"Null":                       0,
		"BidBond":                    1,
		"AdvancePayment":             2,
		"PerformanceBond":            3,
		"GovernmentPaymentGuarantee": 4,
		"MaintenanceBond":            5,
		"ProcurementBond":            6,
		"TransactionRiskBond":        7,
		"CustomsBond":                8,
	}
)

func (x BgType) Enum() *BgType {
	p := new(BgType)
	*p = x
	return p
}

func (x BgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BgType) Descriptor() protoreflect.EnumDescriptor {
	return file_bg_gorm_db_proto_enumTypes[1].Descriptor()
}

func (BgType) Type() protoreflect.EnumType {
	return &file_bg_gorm_db_proto_enumTypes[1]
}

func (x BgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BgType.Descriptor instead.
func (BgType) EnumDescriptor() ([]byte, []int) {
	return file_bg_gorm_db_proto_rawDescGZIP(), []int{1}
}

type Mapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CompanyID     uint64                 `protobuf:"varint,2,opt,name=companyID,proto3" json:"companyID,omitempty"`
	ThirdPartyID  uint64                 `protobuf:"varint,3,opt,name=thirdPartyID,proto3" json:"thirdPartyID,omitempty"`
	BeneficiaryID uint64                 `protobuf:"varint,4,opt,name=beneficiaryID,proto3" json:"beneficiaryID,omitempty"`
	IsMapped      bool                   `protobuf:"varint,5,opt,name=isMapped,proto3" json:"isMapped,omitempty"`
	CreatedByID   uint64                 `protobuf:"varint,41,opt,name=createdByID,proto3" json:"createdByID,omitempty"`
	UpdatedByID   uint64                 `protobuf:"varint,42,opt,name=updatedByID,proto3" json:"updatedByID,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,51,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,52,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *Mapping) Reset() {
	*x = Mapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bg_gorm_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mapping) ProtoMessage() {}

func (x *Mapping) ProtoReflect() protoreflect.Message {
	mi := &file_bg_gorm_db_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mapping.ProtoReflect.Descriptor instead.
func (*Mapping) Descriptor() ([]byte, []int) {
	return file_bg_gorm_db_proto_rawDescGZIP(), []int{0}
}

func (x *Mapping) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Mapping) GetCompanyID() uint64 {
	if x != nil {
		return x.CompanyID
	}
	return 0
}

func (x *Mapping) GetThirdPartyID() uint64 {
	if x != nil {
		return x.ThirdPartyID
	}
	return 0
}

func (x *Mapping) GetBeneficiaryID() uint64 {
	if x != nil {
		return x.BeneficiaryID
	}
	return 0
}

func (x *Mapping) GetIsMapped() bool {
	if x != nil {
		return x.IsMapped
	}
	return false
}

func (x *Mapping) GetCreatedByID() uint64 {
	if x != nil {
		return x.CreatedByID
	}
	return 0
}

func (x *Mapping) GetUpdatedByID() uint64 {
	if x != nil {
		return x.UpdatedByID
	}
	return 0
}

func (x *Mapping) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Mapping) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Currency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code      string                 `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	NameEN    string                 `protobuf:"bytes,3,opt,name=nameEN,proto3" json:"nameEN,omitempty"`
	NameID    string                 `protobuf:"bytes,4,opt,name=nameID,proto3" json:"nameID,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,51,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,52,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *Currency) Reset() {
	*x = Currency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bg_gorm_db_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Currency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Currency) ProtoMessage() {}

func (x *Currency) ProtoReflect() protoreflect.Message {
	mi := &file_bg_gorm_db_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Currency.ProtoReflect.Descriptor instead.
func (*Currency) Descriptor() ([]byte, []int) {
	return file_bg_gorm_db_proto_rawDescGZIP(), []int{1}
}

func (x *Currency) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Currency) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Currency) GetNameEN() string {
	if x != nil {
		return x.NameEN
	}
	return ""
}

func (x *Currency) GetNameID() string {
	if x != nil {
		return x.NameID
	}
	return ""
}

func (x *Currency) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Currency) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_bg_gorm_db_proto protoreflect.FileDescriptor

var file_bg_gorm_db_proto_rawDesc = []byte{
	0x0a, 0x10, 0x62, 0x67, 0x5f, 0x67, 0x6f, 0x72, 0x6d, 0x5f, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x62, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x67, 0x6f, 0x72, 0x6d, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x72,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x03, 0x0a, 0x07, 0x4d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x0a, 0xba, 0xb9, 0x19, 0x06, 0x0a, 0x04, 0x28, 0x01, 0x40, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x26, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x09, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x0c, 0x74, 0x68, 0x69, 0x72, 0x64,
	0x50, 0x61, 0x72, 0x74, 0x79, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x08, 0xba,
	0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x0c, 0x74, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61,
	0x72, 0x74, 0x79, 0x49, 0x44, 0x12, 0x2e, 0x0a, 0x0d, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63,
	0x69, 0x61, 0x72, 0x79, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x42, 0x08, 0xba, 0xb9,
	0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x0d, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69,
	0x61, 0x72, 0x79, 0x49, 0x44, 0x12, 0x2b, 0x0a, 0x08, 0x69, 0x73, 0x4d, 0x61, 0x70, 0x70, 0x65,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x42, 0x0f, 0xba, 0xb9, 0x19, 0x0b, 0x0a, 0x09, 0x3a,
	0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x40, 0x01, 0x52, 0x08, 0x69, 0x73, 0x4d, 0x61, 0x70, 0x70,
	0x65, 0x64, 0x12, 0x2d, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49,
	0x44, 0x18, 0x29, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0b, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x04,
	0x0a, 0x02, 0x40, 0x01, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49,
	0x44, 0x12, 0x2d, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44,
	0x18, 0x2a, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0b, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x04, 0x0a,
	0x02, 0x40, 0x01, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44,
	0x12, 0x45, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x33, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x0b, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x45, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x34, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0b, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x04, 0x0a,
	0x02, 0x40, 0x01, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x3a, 0x10,
	0xba, 0xb9, 0x19, 0x0c, 0x08, 0x01, 0x1a, 0x08, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73,
	0x22, 0xaa, 0x02, 0x0a, 0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1a, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0a, 0xba, 0xb9, 0x19, 0x06, 0x0a,
	0x04, 0x28, 0x01, 0x40, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40,
	0x01, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x06, 0x6e, 0x61, 0x6d, 0x65, 0x45,
	0x4e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40,
	0x01, 0x52, 0x06, 0x6e, 0x61, 0x6d, 0x65, 0x45, 0x4e, 0x12, 0x20, 0x0a, 0x06, 0x6e, 0x61, 0x6d,
	0x65, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a,
	0x02, 0x40, 0x01, 0x52, 0x06, 0x6e, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x12, 0x45, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x33, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0b, 0xe0, 0x41, 0x03, 0xba,
	0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x45, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x34, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x42, 0x0b, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x3a, 0x12, 0xba, 0xb9, 0x19, 0x0e, 0x08,
	0x01, 0x1a, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2a, 0x55, 0x0a,
	0x08, 0x42, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x50, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x6c, 0x6f, 0x73, 0x69, 0x6e, 0x67,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x6c, 0x6f, 0x73,
	0x65, 0x64, 0x10, 0x04, 0x2a, 0xbc, 0x01, 0x0a, 0x06, 0x42, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x75, 0x6c, 0x6c, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x69, 0x64,
	0x42, 0x6f, 0x6e, 0x64, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63,
	0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x65,
	0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x6f, 0x6e, 0x64, 0x10, 0x03, 0x12,
	0x1e, 0x0a, 0x1a, 0x47, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x47, 0x75, 0x61, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x10, 0x04, 0x12,
	0x13, 0x0a, 0x0f, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x6f,
	0x6e, 0x64, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x63, 0x75, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x42, 0x6f, 0x6e, 0x64, 0x10, 0x06, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x69, 0x73, 0x6b, 0x42, 0x6f, 0x6e, 0x64,
	0x10, 0x07, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x73, 0x42, 0x6f, 0x6e,
	0x64, 0x10, 0x08, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_bg_gorm_db_proto_rawDescOnce sync.Once
	file_bg_gorm_db_proto_rawDescData = file_bg_gorm_db_proto_rawDesc
)

func file_bg_gorm_db_proto_rawDescGZIP() []byte {
	file_bg_gorm_db_proto_rawDescOnce.Do(func() {
		file_bg_gorm_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_bg_gorm_db_proto_rawDescData)
	})
	return file_bg_gorm_db_proto_rawDescData
}

var file_bg_gorm_db_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_bg_gorm_db_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bg_gorm_db_proto_goTypes = []interface{}{
	(BgStatus)(0),                 // 0: bg.service.v1.BgStatus
	(BgType)(0),                   // 1: bg.service.v1.BgType
	(*Mapping)(nil),               // 2: bg.service.v1.Mapping
	(*Currency)(nil),              // 3: bg.service.v1.Currency
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_bg_gorm_db_proto_depIdxs = []int32{
	4, // 0: bg.service.v1.Mapping.createdAt:type_name -> google.protobuf.Timestamp
	4, // 1: bg.service.v1.Mapping.updatedAt:type_name -> google.protobuf.Timestamp
	4, // 2: bg.service.v1.Currency.createdAt:type_name -> google.protobuf.Timestamp
	4, // 3: bg.service.v1.Currency.updatedAt:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_bg_gorm_db_proto_init() }
func file_bg_gorm_db_proto_init() {
	if File_bg_gorm_db_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bg_gorm_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mapping); i {
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
		file_bg_gorm_db_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Currency); i {
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
			RawDescriptor: file_bg_gorm_db_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bg_gorm_db_proto_goTypes,
		DependencyIndexes: file_bg_gorm_db_proto_depIdxs,
		EnumInfos:         file_bg_gorm_db_proto_enumTypes,
		MessageInfos:      file_bg_gorm_db_proto_msgTypes,
	}.Build()
	File_bg_gorm_db_proto = out.File
	file_bg_gorm_db_proto_rawDesc = nil
	file_bg_gorm_db_proto_goTypes = nil
	file_bg_gorm_db_proto_depIdxs = nil
}
