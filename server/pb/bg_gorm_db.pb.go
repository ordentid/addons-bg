// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
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

type TaskStatus int32

const (
	TaskStatus_AllStatus     TaskStatus = 0
	TaskStatus_Pending       TaskStatus = 1
	TaskStatus_Draft         TaskStatus = 2
	TaskStatus_Returned      TaskStatus = 3
	TaskStatus_Approved      TaskStatus = 4
	TaskStatus_Rejected      TaskStatus = 5
	TaskStatus_DeleteRequest TaskStatus = 6
	TaskStatus_Deleted       TaskStatus = 7
)

// Enum value maps for TaskStatus.
var (
	TaskStatus_name = map[int32]string{
		0: "AllStatus",
		1: "Pending",
		2: "Draft",
		3: "Returned",
		4: "Approved",
		5: "Rejected",
		6: "DeleteRequest",
		7: "Deleted",
	}
	TaskStatus_value = map[string]int32{
		"AllStatus":     0,
		"Pending":       1,
		"Draft":         2,
		"Returned":      3,
		"Approved":      4,
		"Rejected":      5,
		"DeleteRequest": 6,
		"Deleted":       7,
	}
)

func (x TaskStatus) Enum() *TaskStatus {
	p := new(TaskStatus)
	*p = x
	return p
}

func (x TaskStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_bg_gorm_db_proto_enumTypes[2].Descriptor()
}

func (TaskStatus) Type() protoreflect.EnumType {
	return &file_bg_gorm_db_proto_enumTypes[2]
}

func (x TaskStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskStatus.Descriptor instead.
func (TaskStatus) EnumDescriptor() ([]byte, []int) {
	return file_bg_gorm_db_proto_rawDescGZIP(), []int{2}
}

type TaskStep int32

const (
	TaskStep_AllStep  TaskStep = 0
	TaskStep_Maker    TaskStep = 1
	TaskStep_Checker  TaskStep = 2
	TaskStep_Signer   TaskStep = 3
	TaskStep_Releaser TaskStep = 4
)

// Enum value maps for TaskStep.
var (
	TaskStep_name = map[int32]string{
		0: "AllStep",
		1: "Maker",
		2: "Checker",
		3: "Signer",
		4: "Releaser",
	}
	TaskStep_value = map[string]int32{
		"AllStep":  0,
		"Maker":    1,
		"Checker":  2,
		"Signer":   3,
		"Releaser": 4,
	}
)

func (x TaskStep) Enum() *TaskStep {
	p := new(TaskStep)
	*p = x
	return p
}

func (x TaskStep) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskStep) Descriptor() protoreflect.EnumDescriptor {
	return file_bg_gorm_db_proto_enumTypes[3].Descriptor()
}

func (TaskStep) Type() protoreflect.EnumType {
	return &file_bg_gorm_db_proto_enumTypes[3]
}

func (x TaskStep) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskStep.Descriptor instead.
func (TaskStep) EnumDescriptor() ([]byte, []int) {
	return file_bg_gorm_db_proto_rawDescGZIP(), []int{3}
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

type BgTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskID             uint64                 `protobuf:"varint,1,opt,name=taskID,proto3" json:"taskID,omitempty"`
	TransactionID      string                 `protobuf:"bytes,2,opt,name=transactionID,proto3" json:"transactionID,omitempty"`
	Status             TaskStatus             `protobuf:"varint,3,opt,name=status,proto3,enum=bg.service.v1.TaskStatus" json:"status,omitempty"`
	Step               TaskStep               `protobuf:"varint,4,opt,name=step,proto3,enum=bg.service.v1.TaskStep" json:"step,omitempty"`
	CreatedByID        uint64                 `protobuf:"varint,5,opt,name=createdByID,proto3" json:"createdByID,omitempty"`
	LastApprovedByID   uint64                 `protobuf:"varint,6,opt,name=lastApprovedByID,proto3" json:"lastApprovedByID,omitempty"`
	LastRejectedByID   uint64                 `protobuf:"varint,7,opt,name=lastRejectedByID,proto3" json:"lastRejectedByID,omitempty"`
	Data               string                 `protobuf:"bytes,8,opt,name=data,proto3" json:"data,omitempty"`
	Comment            string                 `protobuf:"bytes,9,opt,name=comment,proto3" json:"comment,omitempty"`
	Reasons            string                 `protobuf:"bytes,10,opt,name=reasons,proto3" json:"reasons,omitempty"`
	LastApprovedByName string                 `protobuf:"bytes,11,opt,name=lastApprovedByName,proto3" json:"lastApprovedByName,omitempty"`
	LastRejectedByName string                 `protobuf:"bytes,12,opt,name=lastRejectedByName,proto3" json:"lastRejectedByName,omitempty"`
	UpdatedByID        uint64                 `protobuf:"varint,13,opt,name=updatedByID,proto3" json:"updatedByID,omitempty"`
	UpdatedByName      string                 `protobuf:"bytes,14,opt,name=updatedByName,proto3" json:"updatedByName,omitempty"`
	CreatedByName      string                 `protobuf:"bytes,15,opt,name=createdByName,proto3" json:"createdByName,omitempty"`
	DataBak            string                 `protobuf:"bytes,16,opt,name=dataBak,proto3" json:"dataBak,omitempty"`
	WorkflowDoc        string                 `protobuf:"bytes,17,opt,name=workflowDoc,proto3" json:"workflowDoc,omitempty"`
	CompanyID          uint64                 `protobuf:"varint,18,opt,name=companyID,proto3" json:"companyID,omitempty"`
	HoldingID          uint64                 `protobuf:"varint,19,opt,name=holdingID,proto3" json:"holdingID,omitempty"`
	CreatedAt          *timestamppb.Timestamp `protobuf:"bytes,21,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt          *timestamppb.Timestamp `protobuf:"bytes,22,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt          *timestamppb.Timestamp `protobuf:"bytes,23,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *BgTask) Reset() {
	*x = BgTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bg_gorm_db_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BgTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BgTask) ProtoMessage() {}

func (x *BgTask) ProtoReflect() protoreflect.Message {
	mi := &file_bg_gorm_db_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BgTask.ProtoReflect.Descriptor instead.
func (*BgTask) Descriptor() ([]byte, []int) {
	return file_bg_gorm_db_proto_rawDescGZIP(), []int{2}
}

func (x *BgTask) GetTaskID() uint64 {
	if x != nil {
		return x.TaskID
	}
	return 0
}

func (x *BgTask) GetTransactionID() string {
	if x != nil {
		return x.TransactionID
	}
	return ""
}

func (x *BgTask) GetStatus() TaskStatus {
	if x != nil {
		return x.Status
	}
	return TaskStatus_AllStatus
}

func (x *BgTask) GetStep() TaskStep {
	if x != nil {
		return x.Step
	}
	return TaskStep_AllStep
}

func (x *BgTask) GetCreatedByID() uint64 {
	if x != nil {
		return x.CreatedByID
	}
	return 0
}

func (x *BgTask) GetLastApprovedByID() uint64 {
	if x != nil {
		return x.LastApprovedByID
	}
	return 0
}

func (x *BgTask) GetLastRejectedByID() uint64 {
	if x != nil {
		return x.LastRejectedByID
	}
	return 0
}

func (x *BgTask) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *BgTask) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *BgTask) GetReasons() string {
	if x != nil {
		return x.Reasons
	}
	return ""
}

func (x *BgTask) GetLastApprovedByName() string {
	if x != nil {
		return x.LastApprovedByName
	}
	return ""
}

func (x *BgTask) GetLastRejectedByName() string {
	if x != nil {
		return x.LastRejectedByName
	}
	return ""
}

func (x *BgTask) GetUpdatedByID() uint64 {
	if x != nil {
		return x.UpdatedByID
	}
	return 0
}

func (x *BgTask) GetUpdatedByName() string {
	if x != nil {
		return x.UpdatedByName
	}
	return ""
}

func (x *BgTask) GetCreatedByName() string {
	if x != nil {
		return x.CreatedByName
	}
	return ""
}

func (x *BgTask) GetDataBak() string {
	if x != nil {
		return x.DataBak
	}
	return ""
}

func (x *BgTask) GetWorkflowDoc() string {
	if x != nil {
		return x.WorkflowDoc
	}
	return ""
}

func (x *BgTask) GetCompanyID() uint64 {
	if x != nil {
		return x.CompanyID
	}
	return 0
}

func (x *BgTask) GetHoldingID() uint64 {
	if x != nil {
		return x.HoldingID
	}
	return 0
}

func (x *BgTask) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *BgTask) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *BgTask) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
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
	0x01, 0x1a, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x22, 0x94, 0x08,
	0x0a, 0x06, 0x42, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x22, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0a, 0xba, 0xb9, 0x19, 0x06, 0x0a, 0x04,
	0x28, 0x01, 0x40, 0x01, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x12, 0x31, 0x0a, 0x0d,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0b, 0xba, 0xb9, 0x19, 0x07, 0x0a, 0x05, 0x3a, 0x01, 0x31, 0x40, 0x01,
	0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12,
	0x3e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x19, 0x2e, 0x62, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0b, 0xba, 0xb9, 0x19, 0x07,
	0x0a, 0x05, 0x3a, 0x01, 0x31, 0x40, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x38, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e,
	0x62, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x53, 0x74, 0x65, 0x70, 0x42, 0x0b, 0xba, 0xb9, 0x19, 0x07, 0x0a, 0x05, 0x3a, 0x01,
	0x31, 0x40, 0x01, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x12, 0x2a, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x42, 0x08,
	0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x10, 0x6c, 0x61, 0x73, 0x74, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x42, 0x79, 0x49,
	0x44, 0x12, 0x2a, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x6c, 0x61, 0x73,
	0x74, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x21, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xba, 0xb9, 0x19,
	0x09, 0x0a, 0x07, 0x12, 0x05, 0x6a, 0x73, 0x6f, 0x6e, 0x62, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x26, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0c, 0xba, 0xb9, 0x19, 0x08, 0x0a, 0x06, 0x12, 0x04, 0x74, 0x65, 0x78, 0x74, 0x52,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0xba, 0xb9, 0x19, 0x10, 0x0a,
	0x0e, 0x12, 0x0c, 0x76, 0x61, 0x72, 0x63, 0x68, 0x61, 0x72, 0x28, 0x32, 0x35, 0x35, 0x29, 0x52,
	0x07, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x33, 0x0a, 0x12, 0x6c, 0x61, 0x73, 0x74,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x41,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a,
	0x12, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x42, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x12,
	0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x42, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49,
	0x44, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x49, 0x44, 0x12, 0x29, 0x0a, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01,
	0x52, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x29, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0d, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x64, 0x61,
	0x74, 0x61, 0x42, 0x61, 0x6b, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xe0, 0x41, 0x03,
	0xba, 0xb9, 0x19, 0x09, 0x0a, 0x07, 0x12, 0x05, 0x6a, 0x73, 0x6f, 0x6e, 0x62, 0x52, 0x07, 0x64,
	0x61, 0x74, 0x61, 0x42, 0x61, 0x6b, 0x12, 0x33, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x44, 0x6f, 0x63, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0xba, 0xb9, 0x19,
	0x0d, 0x0a, 0x0b, 0x12, 0x05, 0x6a, 0x73, 0x6f, 0x6e, 0x62, 0x3a, 0x02, 0x7b, 0x7d, 0x52, 0x0b,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x44, 0x6f, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x18, 0x12, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x6f, 0x6c,
	0x64, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x18, 0x13, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x68, 0x6f,
	0x6c, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x16,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x3a, 0x0d, 0xba, 0xb9, 0x19, 0x09, 0x08, 0x01, 0x1a, 0x05, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x2a, 0x55, 0x0a, 0x08, 0x42, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0d, 0x0a, 0x09, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x43,
	0x6c, 0x61, 0x69, 0x6d, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d,
	0x43, 0x6c, 0x6f, 0x73, 0x69, 0x6e, 0x67, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x10, 0x03, 0x12,
	0x0a, 0x0a, 0x06, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x10, 0x04, 0x2a, 0xbc, 0x01, 0x0a, 0x06,
	0x42, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x75, 0x6c, 0x6c, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x42, 0x69, 0x64, 0x42, 0x6f, 0x6e, 0x64, 0x10, 0x01, 0x12, 0x12, 0x0a,
	0x0e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x10,
	0x02, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65,
	0x42, 0x6f, 0x6e, 0x64, 0x10, 0x03, 0x12, 0x1e, 0x0a, 0x1a, 0x47, 0x6f, 0x76, 0x65, 0x72, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x75, 0x61, 0x72, 0x61,
	0x6e, 0x74, 0x65, 0x65, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x6f, 0x6e, 0x64, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x50,
	0x72, 0x6f, 0x63, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x6f, 0x6e, 0x64, 0x10, 0x06,
	0x12, 0x17, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x69, 0x73, 0x6b, 0x42, 0x6f, 0x6e, 0x64, 0x10, 0x07, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x73, 0x42, 0x6f, 0x6e, 0x64, 0x10, 0x08, 0x2a, 0x7d, 0x0a, 0x0a, 0x54, 0x61,
	0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x6c, 0x6c, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x72, 0x61, 0x66, 0x74, 0x10, 0x02, 0x12,
	0x0c, 0x0a, 0x08, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x10, 0x03, 0x12, 0x0c, 0x0a,
	0x08, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x52,
	0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0x06, 0x12, 0x0b, 0x0a, 0x07,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x07, 0x2a, 0x49, 0x0a, 0x08, 0x54, 0x61, 0x73,
	0x6b, 0x53, 0x74, 0x65, 0x70, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x65, 0x70,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x61, 0x6b, 0x65, 0x72, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x69,
	0x67, 0x6e, 0x65, 0x72, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x72, 0x10, 0x04, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
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

var file_bg_gorm_db_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_bg_gorm_db_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_bg_gorm_db_proto_goTypes = []interface{}{
	(BgStatus)(0),                 // 0: bg.service.v1.BgStatus
	(BgType)(0),                   // 1: bg.service.v1.BgType
	(TaskStatus)(0),               // 2: bg.service.v1.TaskStatus
	(TaskStep)(0),                 // 3: bg.service.v1.TaskStep
	(*Mapping)(nil),               // 4: bg.service.v1.Mapping
	(*Currency)(nil),              // 5: bg.service.v1.Currency
	(*BgTask)(nil),                // 6: bg.service.v1.BgTask
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_bg_gorm_db_proto_depIdxs = []int32{
	7, // 0: bg.service.v1.Mapping.createdAt:type_name -> google.protobuf.Timestamp
	7, // 1: bg.service.v1.Mapping.updatedAt:type_name -> google.protobuf.Timestamp
	7, // 2: bg.service.v1.Currency.createdAt:type_name -> google.protobuf.Timestamp
	7, // 3: bg.service.v1.Currency.updatedAt:type_name -> google.protobuf.Timestamp
	2, // 4: bg.service.v1.BgTask.status:type_name -> bg.service.v1.TaskStatus
	3, // 5: bg.service.v1.BgTask.step:type_name -> bg.service.v1.TaskStep
	7, // 6: bg.service.v1.BgTask.createdAt:type_name -> google.protobuf.Timestamp
	7, // 7: bg.service.v1.BgTask.updatedAt:type_name -> google.protobuf.Timestamp
	7, // 8: bg.service.v1.BgTask.deletedAt:type_name -> google.protobuf.Timestamp
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
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
		file_bg_gorm_db_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BgTask); i {
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
			NumEnums:      4,
			NumMessages:   3,
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
