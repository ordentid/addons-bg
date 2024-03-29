// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.3
// source: account_gorm_db.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	_ "github.com/mwitkow/go-proto-validators"
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

// Example User
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username  string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password  string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Role      string                 `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_gorm_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_account_gorm_db_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_account_gorm_db_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *User) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *User) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID        uint64                 `protobuf:"varint,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	CompanyID        uint64                 `protobuf:"varint,2,opt,name=companyID,proto3" json:"companyID,omitempty"`
	AccountNumber    string                 `protobuf:"bytes,3,opt,name=accountNumber,proto3" json:"accountNumber,omitempty"`
	AccountAlias     string                 `protobuf:"bytes,4,opt,name=accountAlias,proto3" json:"accountAlias,omitempty"`
	AccountName      string                 `protobuf:"bytes,5,opt,name=accountName,proto3" json:"accountName,omitempty"`
	AccountType      string                 `protobuf:"bytes,6,opt,name=accountType,proto3" json:"accountType,omitempty"`
	AccountStatus    string                 `protobuf:"bytes,7,opt,name=accountStatus,proto3" json:"accountStatus,omitempty"`
	AccountCurrency  string                 `protobuf:"bytes,8,opt,name=accountCurrency,proto3" json:"accountCurrency,omitempty"`
	AccessLevel      string                 `protobuf:"bytes,9,opt,name=accessLevel,proto3" json:"accessLevel,omitempty"`
	IsOwnedByCompany string                 `protobuf:"bytes,10,opt,name=isOwnedByCompany,proto3" json:"isOwnedByCompany,omitempty"`
	CreatedByID      uint64                 `protobuf:"varint,11,opt,name=createdByID,proto3" json:"createdByID,omitempty"`
	UpdatedByID      uint64                 `protobuf:"varint,12,opt,name=updatedByID,proto3" json:"updatedByID,omitempty"`
	DeletedByID      uint64                 `protobuf:"varint,13,opt,name=deletedByID,proto3" json:"deletedByID,omitempty"`
	RoleID           uint64                 `protobuf:"varint,14,opt,name=roleID,proto3" json:"roleID,omitempty"`
	Disabled         bool                   `protobuf:"varint,15,opt,name=disabled,proto3" json:"disabled,omitempty"`
	Cif              string                 `protobuf:"bytes,16,opt,name=cif,proto3" json:"cif,omitempty"`
	ProductCode      string                 `protobuf:"bytes,17,opt,name=productCode,proto3" json:"productCode,omitempty"`
	StatusCode       string                 `protobuf:"bytes,18,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	CreatedAt        *timestamppb.Timestamp `protobuf:"bytes,51,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt        *timestamppb.Timestamp `protobuf:"bytes,52,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt        *timestamppb.Timestamp `protobuf:"bytes,53,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_gorm_db_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_account_gorm_db_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_account_gorm_db_proto_rawDescGZIP(), []int{1}
}

func (x *Account) GetAccountID() uint64 {
	if x != nil {
		return x.AccountID
	}
	return 0
}

func (x *Account) GetCompanyID() uint64 {
	if x != nil {
		return x.CompanyID
	}
	return 0
}

func (x *Account) GetAccountNumber() string {
	if x != nil {
		return x.AccountNumber
	}
	return ""
}

func (x *Account) GetAccountAlias() string {
	if x != nil {
		return x.AccountAlias
	}
	return ""
}

func (x *Account) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *Account) GetAccountType() string {
	if x != nil {
		return x.AccountType
	}
	return ""
}

func (x *Account) GetAccountStatus() string {
	if x != nil {
		return x.AccountStatus
	}
	return ""
}

func (x *Account) GetAccountCurrency() string {
	if x != nil {
		return x.AccountCurrency
	}
	return ""
}

func (x *Account) GetAccessLevel() string {
	if x != nil {
		return x.AccessLevel
	}
	return ""
}

func (x *Account) GetIsOwnedByCompany() string {
	if x != nil {
		return x.IsOwnedByCompany
	}
	return ""
}

func (x *Account) GetCreatedByID() uint64 {
	if x != nil {
		return x.CreatedByID
	}
	return 0
}

func (x *Account) GetUpdatedByID() uint64 {
	if x != nil {
		return x.UpdatedByID
	}
	return 0
}

func (x *Account) GetDeletedByID() uint64 {
	if x != nil {
		return x.DeletedByID
	}
	return 0
}

func (x *Account) GetRoleID() uint64 {
	if x != nil {
		return x.RoleID
	}
	return 0
}

func (x *Account) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *Account) GetCif() string {
	if x != nil {
		return x.Cif
	}
	return ""
}

func (x *Account) GetProductCode() string {
	if x != nil {
		return x.ProductCode
	}
	return ""
}

func (x *Account) GetStatusCode() string {
	if x != nil {
		return x.StatusCode
	}
	return ""
}

func (x *Account) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Account) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Account) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type AccountRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID   uint64                 `protobuf:"varint,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	RoleID      uint64                 `protobuf:"varint,2,opt,name=roleID,proto3" json:"roleID,omitempty"`
	CreatedByID uint64                 `protobuf:"varint,3,opt,name=createdByID,proto3" json:"createdByID,omitempty"`
	UpdatedByID uint64                 `protobuf:"varint,4,opt,name=updatedByID,proto3" json:"updatedByID,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,51,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,52,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *AccountRole) Reset() {
	*x = AccountRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_gorm_db_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountRole) ProtoMessage() {}

func (x *AccountRole) ProtoReflect() protoreflect.Message {
	mi := &file_account_gorm_db_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountRole.ProtoReflect.Descriptor instead.
func (*AccountRole) Descriptor() ([]byte, []int) {
	return file_account_gorm_db_proto_rawDescGZIP(), []int{2}
}

func (x *AccountRole) GetAccountID() uint64 {
	if x != nil {
		return x.AccountID
	}
	return 0
}

func (x *AccountRole) GetRoleID() uint64 {
	if x != nil {
		return x.RoleID
	}
	return 0
}

func (x *AccountRole) GetCreatedByID() uint64 {
	if x != nil {
		return x.CreatedByID
	}
	return 0
}

func (x *AccountRole) GetUpdatedByID() uint64 {
	if x != nil {
		return x.UpdatedByID
	}
	return 0
}

func (x *AccountRole) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *AccountRole) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_account_gorm_db_proto protoreflect.FileDescriptor

var file_account_gorm_db_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x67, 0x6f, 0x72, 0x6d, 0x5f, 0x64,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x72, 0x6d, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2b, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe,
	0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x0d, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x06, 0x0a, 0x04, 0x28, 0x01,
	0x40, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13, 0xe0, 0x41, 0x02, 0xe2, 0xdf, 0x1f,
	0x02, 0x20, 0x01, 0xba, 0xb9, 0x19, 0x06, 0x0a, 0x04, 0x30, 0x01, 0x40, 0x01, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0xe0, 0x41, 0x02, 0xe2, 0xdf,
	0x1f, 0x02, 0x20, 0x01, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0xba, 0xb9, 0x19, 0x0a,
	0x0a, 0x08, 0x3a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x40, 0x01, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x12, 0x3e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x3e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x3e, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x3a, 0x0d, 0xba, 0xb9, 0x19, 0x09, 0x08, 0x01, 0x1a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22,
	0xcd, 0x0a, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x1a,
	0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x13, 0x0a, 0x11, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x44, 0x28, 0x01, 0x40, 0x01, 0x48, 0x01, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x37, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x19, 0xba, 0xb9, 0x19, 0x0f, 0x0a, 0x0d,
	0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x40, 0x01, 0xe2, 0xdf, 0x1f,
	0x02, 0x20, 0x01, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x12, 0x43,
	0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1d, 0xba, 0xb9, 0x19, 0x13, 0x0a, 0x11, 0x0a, 0x0d, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x40, 0x01, 0xe2, 0xdf,
	0x1f, 0x02, 0x20, 0x01, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x4b, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c,
	0x69, 0x61, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x27, 0xba, 0xb9, 0x19, 0x1d, 0x0a,
	0x1b, 0x0a, 0x0c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x12,
	0x05, 0x6a, 0x73, 0x6f, 0x6e, 0x62, 0x3a, 0x02, 0x5b, 0x5d, 0x40, 0x01, 0xe2, 0xdf, 0x1f, 0x02,
	0x20, 0x01, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x69, 0x61, 0x73,
	0x12, 0x3d, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xba, 0xb9, 0x19, 0x11, 0x0a, 0x0f, 0x0a, 0x0b, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x40, 0x01, 0xe2, 0xdf, 0x1f, 0x02,
	0x20, 0x01, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x3d, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xba, 0xb9, 0x19, 0x11, 0x0a, 0x0f, 0x0a, 0x0b, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x40, 0x01, 0xe2, 0xdf, 0x1f, 0x02, 0x20,
	0x01, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x43,
	0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1d, 0xba, 0xb9, 0x19, 0x13, 0x0a, 0x11, 0x0a, 0x0d, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x40, 0x01, 0xe2, 0xdf,
	0x1f, 0x02, 0x20, 0x01, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x49, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1f, 0xba, 0xb9,
	0x19, 0x15, 0x0a, 0x13, 0x0a, 0x0f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x40, 0x01, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x0f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x3d,
	0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x1b, 0xba, 0xb9, 0x19, 0x11, 0x0a, 0x0f, 0x0a, 0x0b, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x40, 0x01, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x4c, 0x0a,
	0x10, 0x69, 0x73, 0x4f, 0x77, 0x6e, 0x65, 0x64, 0x42, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x20, 0xba, 0xb9, 0x19, 0x16, 0x0a, 0x14, 0x0a,
	0x10, 0x49, 0x73, 0x4f, 0x77, 0x6e, 0x65, 0x64, 0x42, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x40, 0x01, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x10, 0x69, 0x73, 0x4f, 0x77, 0x6e,
	0x65, 0x64, 0x42, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x3d, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x1b, 0xba, 0xb9, 0x19, 0x11, 0x0a, 0x0f, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x49, 0x44, 0x40, 0x01, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x3d, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x1b, 0xba, 0xb9, 0x19, 0x11, 0x0a, 0x0f, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x49, 0x44, 0x40, 0x01, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x3d, 0x0a, 0x0b, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x42, 0x1b,
	0xba, 0xb9, 0x19, 0x11, 0x0a, 0x0f, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x49, 0x44, 0x40, 0x01, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65,
	0x49, 0x44, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0e, 0xba, 0xb9, 0x19, 0x0a, 0x0a, 0x08,
	0x0a, 0x06, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x44,
	0x12, 0x2c, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x08, 0x42, 0x10, 0xba, 0xb9, 0x19, 0x0c, 0x0a, 0x0a, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1d,
	0x0a, 0x03, 0x63, 0x69, 0x66, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xba, 0xb9, 0x19,
	0x07, 0x0a, 0x05, 0x0a, 0x03, 0x43, 0x69, 0x66, 0x52, 0x03, 0x63, 0x69, 0x66, 0x12, 0x35, 0x0a,
	0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x13, 0xba, 0xb9, 0x19, 0x0f, 0x0a, 0x0d, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x42, 0x12, 0xba, 0xb9, 0x19, 0x0e, 0x0a, 0x0c,
	0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x4e, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x33, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x14, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x0d,
	0x0a, 0x0b, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x4e, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x34, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x14, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x0d,
	0x0a, 0x0b, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x4e, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x35, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x14, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x0d,
	0x0a, 0x0b, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x52, 0x09, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x3a, 0x15, 0xba, 0xb9, 0x19, 0x11, 0x08, 0x01,
	0x1a, 0x0d, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22,
	0x98, 0x03, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x12,
	0x3a, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x1c, 0xba, 0xb9, 0x19, 0x11, 0x0a, 0x0f, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x44, 0x28, 0x01, 0x40, 0x01, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x03,
	0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x06, 0x72,
	0x6f, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x12, 0xba, 0xb9, 0x19,
	0x0e, 0x0a, 0x0c, 0x0a, 0x06, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x28, 0x01, 0x40, 0x01, 0x52,
	0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x3f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x1d, 0xe0, 0x41,
	0x03, 0xba, 0xb9, 0x19, 0x0f, 0x0a, 0x0d, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x49, 0x44, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x03, 0x52, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x3f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x42, 0x1d, 0xba,
	0xb9, 0x19, 0x0f, 0x0a, 0x0d, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79,
	0x49, 0x44, 0xe0, 0x41, 0x03, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x03, 0x52, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x43, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x33, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x09, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19,
	0x02, 0x0a, 0x00, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x43,
	0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x34, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x09, 0xe0,
	0x41, 0x03, 0xba, 0xb9, 0x19, 0x02, 0x0a, 0x00, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x3a, 0x15, 0xba, 0xb9, 0x19, 0x11, 0x08, 0x01, 0x1a, 0x0d, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_account_gorm_db_proto_rawDescOnce sync.Once
	file_account_gorm_db_proto_rawDescData = file_account_gorm_db_proto_rawDesc
)

func file_account_gorm_db_proto_rawDescGZIP() []byte {
	file_account_gorm_db_proto_rawDescOnce.Do(func() {
		file_account_gorm_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_gorm_db_proto_rawDescData)
	})
	return file_account_gorm_db_proto_rawDescData
}

var file_account_gorm_db_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_account_gorm_db_proto_goTypes = []interface{}{
	(*User)(nil),                  // 0: account.service.v1.User
	(*Account)(nil),               // 1: account.service.v1.Account
	(*AccountRole)(nil),           // 2: account.service.v1.AccountRole
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_account_gorm_db_proto_depIdxs = []int32{
	3, // 0: account.service.v1.User.created_at:type_name -> google.protobuf.Timestamp
	3, // 1: account.service.v1.User.updated_at:type_name -> google.protobuf.Timestamp
	3, // 2: account.service.v1.User.deleted_at:type_name -> google.protobuf.Timestamp
	3, // 3: account.service.v1.Account.createdAt:type_name -> google.protobuf.Timestamp
	3, // 4: account.service.v1.Account.updatedAt:type_name -> google.protobuf.Timestamp
	3, // 5: account.service.v1.Account.deletedAt:type_name -> google.protobuf.Timestamp
	3, // 6: account.service.v1.AccountRole.createdAt:type_name -> google.protobuf.Timestamp
	3, // 7: account.service.v1.AccountRole.updatedAt:type_name -> google.protobuf.Timestamp
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_account_gorm_db_proto_init() }
func file_account_gorm_db_proto_init() {
	if File_account_gorm_db_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_account_gorm_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_account_gorm_db_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_account_gorm_db_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountRole); i {
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
			RawDescriptor: file_account_gorm_db_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_account_gorm_db_proto_goTypes,
		DependencyIndexes: file_account_gorm_db_proto_depIdxs,
		MessageInfos:      file_account_gorm_db_proto_msgTypes,
	}.Build()
	File_account_gorm_db_proto = out.File
	file_account_gorm_db_proto_rawDesc = nil
	file_account_gorm_db_proto_goTypes = nil
	file_account_gorm_db_proto_depIdxs = nil
}
