// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: bg_gorm_db.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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
	BgStatus_Active   BgStatus = 0
	BgStatus_InActive BgStatus = 1
	BgStatus_Claimed  BgStatus = 2
)

// Enum value maps for BgStatus.
var (
	BgStatus_name = map[int32]string{
		0: "Active",
		1: "InActive",
		2: "Claimed",
	}
	BgStatus_value = map[string]int32{
		"Active":   0,
		"InActive": 1,
		"Claimed":  2,
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
	BgType_BidBond                    BgType = 0
	BgType_AdvancePayment             BgType = 1
	BgType_PerformanceBond            BgType = 2
	BgType_GovernmentPaymentGuarantee BgType = 3
	BgType_MaintenanceBond            BgType = 4
	BgType_ProcurementBond            BgType = 5
	BgType_TransactionRiskBond        BgType = 6
	BgType_CustomsBond                BgType = 7
)

// Enum value maps for BgType.
var (
	BgType_name = map[int32]string{
		0: "BidBond",
		1: "AdvancePayment",
		2: "PerformanceBond",
		3: "GovernmentPaymentGuarantee",
		4: "MaintenanceBond",
		5: "ProcurementBond",
		6: "TransactionRiskBond",
		7: "CustomsBond",
	}
	BgType_value = map[string]int32{
		"BidBond":                    0,
		"AdvancePayment":             1,
		"PerformanceBond":            2,
		"GovernmentPaymentGuarantee": 3,
		"MaintenanceBond":            4,
		"ProcurementBond":            5,
		"TransactionRiskBond":        6,
		"CustomsBond":                7,
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

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TransactionID      uint64                 `protobuf:"varint,2,opt,name=transactionID,proto3" json:"transactionID,omitempty"`
	ThirdPartyID       uint64                 `protobuf:"varint,3,opt,name=thirdPartyID,proto3" json:"thirdPartyID,omitempty"`
	ReferenceNo        string                 `protobuf:"bytes,4,opt,name=referenceNo,proto3" json:"referenceNo,omitempty"`
	ApplicantName      string                 `protobuf:"bytes,5,opt,name=applicantName,proto3" json:"applicantName,omitempty"`
	BeneficiaryName    string                 `protobuf:"bytes,6,opt,name=beneficiaryName,proto3" json:"beneficiaryName,omitempty"`
	IssueDate          string                 `protobuf:"bytes,7,opt,name=issueDate,proto3" json:"issueDate,omitempty"`
	EffectiveDate      string                 `protobuf:"bytes,8,opt,name=effectiveDate,proto3" json:"effectiveDate,omitempty"`
	ExpiryDate         string                 `protobuf:"bytes,9,opt,name=expiryDate,proto3" json:"expiryDate,omitempty"`
	ClaimPeriod        uint64                 `protobuf:"varint,10,opt,name=claimPeriod,proto3" json:"claimPeriod,omitempty"`
	ClosingDate        string                 `protobuf:"bytes,11,opt,name=closingDate,proto3" json:"closingDate,omitempty"`
	Currency           string                 `protobuf:"bytes,12,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount             float64                `protobuf:"fixed64,13,opt,name=amount,proto3" json:"amount,omitempty"`
	Remark             string                 `protobuf:"bytes,14,opt,name=remark,proto3" json:"remark,omitempty"`
	TransactionStatus  string                 `protobuf:"bytes,15,opt,name=transactionStatus,proto3" json:"transactionStatus,omitempty"`
	ChannelID          uint64                 `protobuf:"varint,16,opt,name=channelID,proto3" json:"channelID,omitempty"`
	ChannelName        string                 `protobuf:"bytes,17,opt,name=channelName,proto3" json:"channelName,omitempty"`
	TransactionTypeID  uint64                 `protobuf:"varint,18,opt,name=transactionTypeID,proto3" json:"transactionTypeID,omitempty"`
	DocumentPath       string                 `protobuf:"bytes,19,opt,name=documentPath,proto3" json:"documentPath,omitempty"`
	CompanyID          uint64                 `protobuf:"varint,21,opt,name=companyID,proto3" json:"companyID,omitempty"`
	Status             string                 `protobuf:"bytes,20,opt,name=status,proto3" json:"status,omitempty"`
	IsAllowBeneficiary bool                   `protobuf:"varint,22,opt,name=isAllowBeneficiary,proto3" json:"isAllowBeneficiary,omitempty"`
	CreatedByID        uint64                 `protobuf:"varint,41,opt,name=createdByID,proto3" json:"createdByID,omitempty"`
	UpdatedByID        uint64                 `protobuf:"varint,42,opt,name=updatedByID,proto3" json:"updatedByID,omitempty"`
	CreatedAt          *timestamppb.Timestamp `protobuf:"bytes,51,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt          *timestamppb.Timestamp `protobuf:"bytes,52,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bg_gorm_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_bg_gorm_db_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Transaction) GetTransactionID() uint64 {
	if x != nil {
		return x.TransactionID
	}
	return 0
}

func (x *Transaction) GetThirdPartyID() uint64 {
	if x != nil {
		return x.ThirdPartyID
	}
	return 0
}

func (x *Transaction) GetReferenceNo() string {
	if x != nil {
		return x.ReferenceNo
	}
	return ""
}

func (x *Transaction) GetApplicantName() string {
	if x != nil {
		return x.ApplicantName
	}
	return ""
}

func (x *Transaction) GetBeneficiaryName() string {
	if x != nil {
		return x.BeneficiaryName
	}
	return ""
}

func (x *Transaction) GetIssueDate() string {
	if x != nil {
		return x.IssueDate
	}
	return ""
}

func (x *Transaction) GetEffectiveDate() string {
	if x != nil {
		return x.EffectiveDate
	}
	return ""
}

func (x *Transaction) GetExpiryDate() string {
	if x != nil {
		return x.ExpiryDate
	}
	return ""
}

func (x *Transaction) GetClaimPeriod() uint64 {
	if x != nil {
		return x.ClaimPeriod
	}
	return 0
}

func (x *Transaction) GetClosingDate() string {
	if x != nil {
		return x.ClosingDate
	}
	return ""
}

func (x *Transaction) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Transaction) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Transaction) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *Transaction) GetTransactionStatus() string {
	if x != nil {
		return x.TransactionStatus
	}
	return ""
}

func (x *Transaction) GetChannelID() uint64 {
	if x != nil {
		return x.ChannelID
	}
	return 0
}

func (x *Transaction) GetChannelName() string {
	if x != nil {
		return x.ChannelName
	}
	return ""
}

func (x *Transaction) GetTransactionTypeID() uint64 {
	if x != nil {
		return x.TransactionTypeID
	}
	return 0
}

func (x *Transaction) GetDocumentPath() string {
	if x != nil {
		return x.DocumentPath
	}
	return ""
}

func (x *Transaction) GetCompanyID() uint64 {
	if x != nil {
		return x.CompanyID
	}
	return 0
}

func (x *Transaction) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Transaction) GetIsAllowBeneficiary() bool {
	if x != nil {
		return x.IsAllowBeneficiary
	}
	return false
}

func (x *Transaction) GetCreatedByID() uint64 {
	if x != nil {
		return x.CreatedByID
	}
	return 0
}

func (x *Transaction) GetUpdatedByID() uint64 {
	if x != nil {
		return x.UpdatedByID
	}
	return 0
}

func (x *Transaction) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Transaction) GetUpdatedAt() *timestamppb.Timestamp {
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
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x0a, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x0a, 0xba, 0xb9, 0x19, 0x06, 0x0a, 0x04, 0x28, 0x01, 0x40, 0x01, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04,
	0x0a, 0x02, 0x40, 0x01, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74,
	0x79, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x74, 0x68, 0x69, 0x72, 0x64,
	0x50, 0x61, 0x72, 0x74, 0x79, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x0b, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x4e, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0xb9,
	0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x0b, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x4e, 0x6f, 0x12, 0x2e, 0x0a, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04,
	0x0a, 0x02, 0x40, 0x01, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x6e, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x0f, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61,
	0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0xb9,
	0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x0f, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69,
	0x61, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x5d, 0x0a, 0x09, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x44, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3f, 0x92, 0x41, 0x34, 0x4a,
	0x0c, 0x22, 0x32, 0x30, 0x32, 0x32, 0x2d, 0x30, 0x37, 0x2d, 0x30, 0x36, 0x22, 0x8a, 0x01, 0x1c,
	0x5e, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x34, 0x7d, 0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x7b,
	0x32, 0x7d, 0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x32, 0x7d, 0x24, 0xa2, 0x02, 0x04, 0x64,
	0x61, 0x74, 0x65, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x09, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x65, 0x0a, 0x0d, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3f, 0x92,
	0x41, 0x34, 0x4a, 0x0c, 0x22, 0x32, 0x30, 0x32, 0x32, 0x2d, 0x30, 0x37, 0x2d, 0x30, 0x36, 0x22,
	0x8a, 0x01, 0x1c, 0x5e, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x34, 0x7d, 0x2d, 0x5b, 0x30, 0x2d,
	0x39, 0x5d, 0x7b, 0x32, 0x7d, 0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x32, 0x7d, 0x24, 0xa2,
	0x02, 0x04, 0x64, 0x61, 0x74, 0x65, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x0d,
	0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x5f, 0x0a,
	0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x3f, 0x92, 0x41, 0x34, 0x4a, 0x0c, 0x22, 0x32, 0x30, 0x32, 0x32, 0x2d, 0x30, 0x37,
	0x2d, 0x30, 0x36, 0x22, 0x8a, 0x01, 0x1c, 0x5e, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x34, 0x7d,
	0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x32, 0x7d, 0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x7b,
	0x32, 0x7d, 0x24, 0xa2, 0x02, 0x04, 0x64, 0x61, 0x74, 0x65, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02,
	0x40, 0x01, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2a,
	0x0a, 0x0b, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x0b, 0x63,
	0x6c, 0x61, 0x69, 0x6d, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x61, 0x0a, 0x0b, 0x63, 0x6c,
	0x6f, 0x73, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x3f, 0x92, 0x41, 0x34, 0x4a, 0x0c, 0x22, 0x32, 0x30, 0x32, 0x32, 0x2d, 0x30, 0x37, 0x2d, 0x30,
	0x36, 0x22, 0x8a, 0x01, 0x1c, 0x5e, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x34, 0x7d, 0x2d, 0x5b,
	0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x32, 0x7d, 0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x32, 0x7d,
	0x24, 0xa2, 0x02, 0x04, 0x64, 0x61, 0x74, 0x65, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01,
	0x52, 0x0b, 0x63, 0x6c, 0x6f, 0x73, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a,
	0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x12, 0x20, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x01, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52,
	0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x36, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x11, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x44, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x44, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x2c, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x49, 0x44, 0x18, 0x12, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x22, 0x0a,
	0x0c, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x74, 0x68, 0x18, 0x13, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x26, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a,
	0x02, 0x40, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x38, 0x0a, 0x12, 0x69,
	0x73, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72,
	0x79, 0x18, 0x16, 0x20, 0x01, 0x28, 0x08, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40,
	0x01, 0x52, 0x12, 0x69, 0x73, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69,
	0x63, 0x69, 0x61, 0x72, 0x79, 0x12, 0x2d, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x49, 0x44, 0x18, 0x29, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0b, 0xe0, 0x41, 0x03, 0xba,
	0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x49, 0x44, 0x12, 0x2d, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x49, 0x44, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0b, 0xe0, 0x41, 0x03, 0xba, 0xb9,
	0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x49, 0x44, 0x12, 0x45, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x33, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x0b, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x45, 0x0a, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x34, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0b, 0xe0, 0x41, 0x03, 0xba, 0xb9,
	0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x3a, 0x14, 0xba, 0xb9, 0x19, 0x10, 0x08, 0x01, 0x1a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2a, 0x31, 0x0a, 0x08, 0x42, 0x67, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x49, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x65, 0x64, 0x10, 0x02, 0x2a, 0xb2, 0x01, 0x0a, 0x06, 0x42,
	0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x69, 0x64, 0x42, 0x6f, 0x6e, 0x64,
	0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x6f, 0x6e, 0x64, 0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x47,
	0x6f, 0x76, 0x65, 0x72, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x47, 0x75, 0x61, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x4d,
	0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x6f, 0x6e, 0x64, 0x10, 0x04,
	0x12, 0x13, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x63, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42,
	0x6f, 0x6e, 0x64, 0x10, 0x05, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x69, 0x73, 0x6b, 0x42, 0x6f, 0x6e, 0x64, 0x10, 0x06, 0x12, 0x0f,
	0x0a, 0x0b, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x73, 0x42, 0x6f, 0x6e, 0x64, 0x10, 0x07, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
var file_bg_gorm_db_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_bg_gorm_db_proto_goTypes = []interface{}{
	(BgStatus)(0),                 // 0: bg.service.v1.BgStatus
	(BgType)(0),                   // 1: bg.service.v1.BgType
	(*Transaction)(nil),           // 2: bg.service.v1.Transaction
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_bg_gorm_db_proto_depIdxs = []int32{
	3, // 0: bg.service.v1.Transaction.createdAt:type_name -> google.protobuf.Timestamp
	3, // 1: bg.service.v1.Transaction.updatedAt:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_bg_gorm_db_proto_init() }
func file_bg_gorm_db_proto_init() {
	if File_bg_gorm_db_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bg_gorm_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
			NumMessages:   1,
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
