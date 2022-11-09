// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: gorm_db.proto

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

type Currency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,51,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,52,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *Currency) Reset() {
	*x = Currency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gorm_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Currency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Currency) ProtoMessage() {}

func (x *Currency) ProtoReflect() protoreflect.Message {
	mi := &file_gorm_db_proto_msgTypes[0]
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
	return file_gorm_db_proto_rawDescGZIP(), []int{0}
}

func (x *Currency) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Currency) GetName() string {
	if x != nil {
		return x.Name
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

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId  uint64                 `protobuf:"varint,1,opt,name=transactionId,proto3" json:"transactionId,omitempty"`
	Status         string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"` // "status": 1 | 2 | 3 | 4,
	IsSchedule     bool                   `protobuf:"varint,3,opt,name=isSchedule,proto3" json:"isSchedule,omitempty"`
	ScheduleDate   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=schedule_date,json=scheduleDate,proto3" json:"schedule_date,omitempty"`
	ModuleId       uint64                 `protobuf:"varint,5,opt,name=moduleId,proto3" json:"moduleId,omitempty"`
	Feature        string                 `protobuf:"bytes,6,opt,name=feature,proto3" json:"feature,omitempty"`
	FeatureId      uint64                 `protobuf:"varint,7,opt,name=featureId,proto3" json:"featureId,omitempty"`
	ExternalId     string                 `protobuf:"bytes,8,opt,name=externalId,proto3" json:"externalId,omitempty"`
	JurnalSeq      string                 `protobuf:"bytes,9,opt,name=jurnalSeq,proto3" json:"jurnalSeq,omitempty"`
	ResponseDetail string                 `protobuf:"bytes,10,opt,name=responseDetail,proto3" json:"responseDetail,omitempty"`
	IsRetry        bool                   `protobuf:"varint,11,opt,name=isRetry,proto3" json:"isRetry,omitempty"`
	RetryNumber    uint32                 `protobuf:"varint,12,opt,name=retryNumber,proto3" json:"retryNumber,omitempty"`
	DebitAmount    string                 `protobuf:"bytes,13,opt,name=debitAmount,proto3" json:"debitAmount,omitempty"`
	CreditAmount   string                 `protobuf:"bytes,14,opt,name=creditAmount,proto3" json:"creditAmount,omitempty"`
	DebitCurrency  string                 `protobuf:"bytes,15,opt,name=debitCurrency,proto3" json:"debitCurrency,omitempty"`
	CreditCurrency string                 `protobuf:"bytes,16,opt,name=creditCurrency,proto3" json:"creditCurrency,omitempty"`
	DebitAccount   string                 `protobuf:"bytes,17,opt,name=debitAccount,proto3" json:"debitAccount,omitempty"`
	CreditAccount  string                 `protobuf:"bytes,18,opt,name=creditAccount,proto3" json:"creditAccount,omitempty"`
	CompanyId      uint64                 `protobuf:"varint,19,opt,name=companyId,proto3" json:"companyId,omitempty"`
	HoldingId      uint64                 `protobuf:"varint,20,opt,name=holdingId,proto3" json:"holdingId,omitempty"`
	DeletedAt      *timestamppb.Timestamp `protobuf:"bytes,21,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,22,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt      *timestamppb.Timestamp `protobuf:"bytes,23,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gorm_db_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_gorm_db_proto_msgTypes[1]
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
	return file_gorm_db_proto_rawDescGZIP(), []int{1}
}

func (x *Transaction) GetTransactionId() uint64 {
	if x != nil {
		return x.TransactionId
	}
	return 0
}

func (x *Transaction) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Transaction) GetIsSchedule() bool {
	if x != nil {
		return x.IsSchedule
	}
	return false
}

func (x *Transaction) GetScheduleDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ScheduleDate
	}
	return nil
}

func (x *Transaction) GetModuleId() uint64 {
	if x != nil {
		return x.ModuleId
	}
	return 0
}

func (x *Transaction) GetFeature() string {
	if x != nil {
		return x.Feature
	}
	return ""
}

func (x *Transaction) GetFeatureId() uint64 {
	if x != nil {
		return x.FeatureId
	}
	return 0
}

func (x *Transaction) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *Transaction) GetJurnalSeq() string {
	if x != nil {
		return x.JurnalSeq
	}
	return ""
}

func (x *Transaction) GetResponseDetail() string {
	if x != nil {
		return x.ResponseDetail
	}
	return ""
}

func (x *Transaction) GetIsRetry() bool {
	if x != nil {
		return x.IsRetry
	}
	return false
}

func (x *Transaction) GetRetryNumber() uint32 {
	if x != nil {
		return x.RetryNumber
	}
	return 0
}

func (x *Transaction) GetDebitAmount() string {
	if x != nil {
		return x.DebitAmount
	}
	return ""
}

func (x *Transaction) GetCreditAmount() string {
	if x != nil {
		return x.CreditAmount
	}
	return ""
}

func (x *Transaction) GetDebitCurrency() string {
	if x != nil {
		return x.DebitCurrency
	}
	return ""
}

func (x *Transaction) GetCreditCurrency() string {
	if x != nil {
		return x.CreditCurrency
	}
	return ""
}

func (x *Transaction) GetDebitAccount() string {
	if x != nil {
		return x.DebitAccount
	}
	return ""
}

func (x *Transaction) GetCreditAccount() string {
	if x != nil {
		return x.CreditAccount
	}
	return ""
}

func (x *Transaction) GetCompanyId() uint64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *Transaction) GetHoldingId() uint64 {
	if x != nil {
		return x.HoldingId
	}
	return 0
}

func (x *Transaction) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
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

type Scheduler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchedulerId  uint64                 `protobuf:"varint,1,opt,name=schedulerId,proto3" json:"schedulerId,omitempty"`
	Transaction  *Transaction           `protobuf:"bytes,2,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Status       uint32                 `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"` // 0 = status pending, 1 = status success, 2 status error/failed
	ScheduleDate *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=schedule_date,json=scheduleDate,proto3" json:"schedule_date,omitempty"`
	IsCutOff     bool                   `protobuf:"varint,5,opt,name=is_cut_off,json=isCutOff,proto3" json:"is_cut_off,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *Scheduler) Reset() {
	*x = Scheduler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gorm_db_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scheduler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scheduler) ProtoMessage() {}

func (x *Scheduler) ProtoReflect() protoreflect.Message {
	mi := &file_gorm_db_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scheduler.ProtoReflect.Descriptor instead.
func (*Scheduler) Descriptor() ([]byte, []int) {
	return file_gorm_db_proto_rawDescGZIP(), []int{2}
}

func (x *Scheduler) GetSchedulerId() uint64 {
	if x != nil {
		return x.SchedulerId
	}
	return 0
}

func (x *Scheduler) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *Scheduler) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Scheduler) GetScheduleDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ScheduleDate
	}
	return nil
}

func (x *Scheduler) GetIsCutOff() bool {
	if x != nil {
		return x.IsCutOff
	}
	return false
}

func (x *Scheduler) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type HostToHostConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigId             uint64 `protobuf:"varint,1,opt,name=configId,proto3" json:"configId,omitempty"`
	DirectoryUrl         string `protobuf:"bytes,2,opt,name=directoryUrl,proto3" json:"directoryUrl,omitempty"`
	DirectoryUsername    string `protobuf:"bytes,3,opt,name=directoryUsername,proto3" json:"directoryUsername,omitempty"`
	DirectoryPassword    string `protobuf:"bytes,4,opt,name=directoryPassword,proto3" json:"directoryPassword,omitempty"`
	HostToHostUserId     uint64 `protobuf:"varint,5,opt,name=hostToHostUserId,proto3" json:"hostToHostUserId,omitempty"`
	Encryption           string `protobuf:"bytes,6,opt,name=encryption,proto3" json:"encryption,omitempty"`
	EncryptionPrivateKey string `protobuf:"bytes,7,opt,name=encryptionPrivateKey,proto3" json:"encryptionPrivateKey,omitempty"`
	EncryptionPublicKey  string `protobuf:"bytes,8,opt,name=encryptionPublicKey,proto3" json:"encryptionPublicKey,omitempty"`
}

func (x *HostToHostConfig) Reset() {
	*x = HostToHostConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gorm_db_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostToHostConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostToHostConfig) ProtoMessage() {}

func (x *HostToHostConfig) ProtoReflect() protoreflect.Message {
	mi := &file_gorm_db_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostToHostConfig.ProtoReflect.Descriptor instead.
func (*HostToHostConfig) Descriptor() ([]byte, []int) {
	return file_gorm_db_proto_rawDescGZIP(), []int{3}
}

func (x *HostToHostConfig) GetConfigId() uint64 {
	if x != nil {
		return x.ConfigId
	}
	return 0
}

func (x *HostToHostConfig) GetDirectoryUrl() string {
	if x != nil {
		return x.DirectoryUrl
	}
	return ""
}

func (x *HostToHostConfig) GetDirectoryUsername() string {
	if x != nil {
		return x.DirectoryUsername
	}
	return ""
}

func (x *HostToHostConfig) GetDirectoryPassword() string {
	if x != nil {
		return x.DirectoryPassword
	}
	return ""
}

func (x *HostToHostConfig) GetHostToHostUserId() uint64 {
	if x != nil {
		return x.HostToHostUserId
	}
	return 0
}

func (x *HostToHostConfig) GetEncryption() string {
	if x != nil {
		return x.Encryption
	}
	return ""
}

func (x *HostToHostConfig) GetEncryptionPrivateKey() string {
	if x != nil {
		return x.EncryptionPrivateKey
	}
	return ""
}

func (x *HostToHostConfig) GetEncryptionPublicKey() string {
	if x != nil {
		return x.EncryptionPublicKey
	}
	return ""
}

var File_gorm_db_proto protoreflect.FileDescriptor

var file_gorm_db_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x6f, 0x72, 0x6d, 0x5f, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x16, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x72, 0x6d, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x01,
	0x0a, 0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0a, 0xba, 0xb9, 0x19, 0x06, 0x0a, 0x04, 0x28, 0x01,
	0x40, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x33, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0b, 0xe0, 0x41, 0x03, 0xba, 0xb9,
	0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x45, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x34,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x0b, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x3a, 0x10, 0xba, 0xb9, 0x19, 0x0c, 0x08, 0x01,
	0x1a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0xfc, 0x07, 0x0a, 0x0b, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x0d, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x0d, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x06, 0x0a, 0x04, 0x28, 0x01, 0x40, 0x01,
	0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x15, 0xba, 0xb9, 0x19, 0x11, 0x0a, 0x0f, 0x12, 0x0a, 0x76, 0x61, 0x72, 0x63, 0x68, 0x61, 0x72,
	0x28, 0x35, 0x29, 0x3a, 0x01, 0x31, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x69, 0x73, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x3f,
	0x0a, 0x0d, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x07, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0xba, 0xb9,
	0x19, 0x10, 0x0a, 0x0e, 0x12, 0x0c, 0x76, 0x61, 0x72, 0x63, 0x68, 0x61, 0x72, 0x28, 0x32, 0x35,
	0x35, 0x29, 0x52, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x0a, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xba,
	0xb9, 0x19, 0x08, 0x0a, 0x06, 0x12, 0x04, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0a, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x09, 0x6a, 0x75, 0x72, 0x6e, 0x61,
	0x6c, 0x53, 0x65, 0x71, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xba, 0xb9, 0x19, 0x08,
	0x0a, 0x06, 0x12, 0x04, 0x74, 0x65, 0x78, 0x74, 0x52, 0x09, 0x6a, 0x75, 0x72, 0x6e, 0x61, 0x6c,
	0x53, 0x65, 0x71, 0x12, 0x34, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xba, 0xb9, 0x19,
	0x08, 0x0a, 0x06, 0x12, 0x04, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x52,
	0x65, 0x74, 0x72, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x52, 0x65,
	0x74, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x74, 0x72, 0x79, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x65, 0x74, 0x72, 0x79, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x62, 0x69, 0x74, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x62, 0x69,
	0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x64,
	0x65, 0x62, 0x69, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x62, 0x69, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x62,
	0x69, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x64, 0x65, 0x62, 0x69, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a,
	0x0d, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x12,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64,
	0x18, 0x13, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x6f, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x68, 0x6f, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12,
	0x38, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x43, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x09, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19,
	0x02, 0x0a, 0x00, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x43,
	0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x17, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x09, 0xe0,
	0x41, 0x03, 0xba, 0xb9, 0x19, 0x02, 0x0a, 0x00, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x3a, 0x14, 0xba, 0xb9, 0x19, 0x10, 0x08, 0x01, 0x1a, 0x0c, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x96, 0x03, 0x0a, 0x09, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0d, 0xe0, 0x41,
	0x03, 0xba, 0xb9, 0x19, 0x06, 0x0a, 0x04, 0x28, 0x01, 0x40, 0x01, 0x52, 0x0b, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x12, 0x60, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x19, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x12, 0x22, 0x10, 0x0a, 0x0e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x52, 0x0b, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0xba, 0xb9, 0x19, 0x05,
	0x0a, 0x03, 0x3a, 0x01, 0x30, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4d, 0x0a,
	0x0d, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x0c, 0xba, 0xb9, 0x19, 0x08, 0x0a, 0x06, 0x3a, 0x04, 0x4e, 0x55, 0x4c, 0x4c, 0x52, 0x0c,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x0a,
	0x69, 0x73, 0x5f, 0x63, 0x75, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x42, 0x0d, 0xba, 0xb9, 0x19, 0x09, 0x0a, 0x07, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52,
	0x08, 0x69, 0x73, 0x43, 0x75, 0x74, 0x4f, 0x66, 0x66, 0x12, 0x43, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x09, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19,
	0x02, 0x0a, 0x00, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x3a, 0x12,
	0xba, 0xb9, 0x19, 0x0e, 0x08, 0x01, 0x1a, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x73, 0x22, 0xd5, 0x03, 0x0a, 0x10, 0x48, 0x6f, 0x73, 0x74, 0x54, 0x6f, 0x48, 0x6f, 0x73,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x29, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0d, 0xe0, 0x41, 0x03, 0xba, 0xb9,
	0x19, 0x06, 0x0a, 0x04, 0x28, 0x01, 0x40, 0x01, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x49, 0x64, 0x12, 0x2c, 0x0a, 0x0c, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x55,
	0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02,
	0x40, 0x01, 0x52, 0x0c, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x55, 0x72, 0x6c,
	0x12, 0x36, 0x0a, 0x11, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0xb9, 0x19,
	0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x11, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x11, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x11, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x34, 0x0a, 0x10, 0x68, 0x6f, 0x73, 0x74, 0x54, 0x6f, 0x48, 0x6f, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04,
	0x0a, 0x02, 0x40, 0x01, 0x52, 0x10, 0x68, 0x6f, 0x73, 0x74, 0x54, 0x6f, 0x48, 0x6f, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x14, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xba, 0xb9, 0x19, 0x0a, 0x0a, 0x08, 0x12, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x40, 0x01, 0x52, 0x14, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x40, 0x0a, 0x13, 0x65, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xba, 0xb9, 0x19, 0x0a, 0x0a, 0x08, 0x12,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x40, 0x01, 0x52, 0x13, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x3a, 0x1c, 0xba, 0xb9,
	0x19, 0x18, 0x08, 0x01, 0x1a, 0x14, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x6f, 0x5f, 0x68, 0x6f,
	0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gorm_db_proto_rawDescOnce sync.Once
	file_gorm_db_proto_rawDescData = file_gorm_db_proto_rawDesc
)

func file_gorm_db_proto_rawDescGZIP() []byte {
	file_gorm_db_proto_rawDescOnce.Do(func() {
		file_gorm_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_gorm_db_proto_rawDescData)
	})
	return file_gorm_db_proto_rawDescData
}

var file_gorm_db_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_gorm_db_proto_goTypes = []interface{}{
	(*Currency)(nil),              // 0: transaction.service.v1.Currency
	(*Transaction)(nil),           // 1: transaction.service.v1.Transaction
	(*Scheduler)(nil),             // 2: transaction.service.v1.Scheduler
	(*HostToHostConfig)(nil),      // 3: transaction.service.v1.HostToHostConfig
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_gorm_db_proto_depIdxs = []int32{
	4, // 0: transaction.service.v1.Currency.createdAt:type_name -> google.protobuf.Timestamp
	4, // 1: transaction.service.v1.Currency.updatedAt:type_name -> google.protobuf.Timestamp
	4, // 2: transaction.service.v1.Transaction.schedule_date:type_name -> google.protobuf.Timestamp
	4, // 3: transaction.service.v1.Transaction.deletedAt:type_name -> google.protobuf.Timestamp
	4, // 4: transaction.service.v1.Transaction.createdAt:type_name -> google.protobuf.Timestamp
	4, // 5: transaction.service.v1.Transaction.updatedAt:type_name -> google.protobuf.Timestamp
	1, // 6: transaction.service.v1.Scheduler.transaction:type_name -> transaction.service.v1.Transaction
	4, // 7: transaction.service.v1.Scheduler.schedule_date:type_name -> google.protobuf.Timestamp
	4, // 8: transaction.service.v1.Scheduler.createdAt:type_name -> google.protobuf.Timestamp
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_gorm_db_proto_init() }
func file_gorm_db_proto_init() {
	if File_gorm_db_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gorm_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_gorm_db_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_gorm_db_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scheduler); i {
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
		file_gorm_db_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostToHostConfig); i {
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
			RawDescriptor: file_gorm_db_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gorm_db_proto_goTypes,
		DependencyIndexes: file_gorm_db_proto_depIdxs,
		MessageInfos:      file_gorm_db_proto_msgTypes,
	}.Build()
	File_gorm_db_proto = out.File
	file_gorm_db_proto_rawDesc = nil
	file_gorm_db_proto_goTypes = nil
	file_gorm_db_proto_depIdxs = nil
}