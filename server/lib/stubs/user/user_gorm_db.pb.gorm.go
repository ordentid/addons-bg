package pb

import (
	context "context"
	fmt "fmt"
	gorm1 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	errors "github.com/infobloxopen/protoc-gen-gorm/errors"
	gorm "github.com/jinzhu/gorm"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	strings "strings"
	time "time"
)

type UserORM struct {
	Authority    string     `gorm:"column:Authority"`
	CompanyID    uint64     `gorm:"column:CompanyID"`
	CompanyName  string     `gorm:"column:CompanyName"`
	CreatedAt    *time.Time `gorm:"not null"`
	CreatedByID  uint64     `gorm:"column:CreatedByID;not null"`
	DeletedAt    *time.Time
	DeletedByID  uint64     `gorm:"column:DeletedByID"`
	IdNumber     string     `gorm:"column:IdNumber;not null"`
	IdType       int32      `gorm:"column:IdType;not null"`
	IsBlocked    bool       `gorm:"column:IsBlocked"`
	TokenFcm     string     `gorm:"column:TokenFcm"`
	UpdatedAt    *time.Time `gorm:"not null"`
	UpdatedByID  uint64     `gorm:"column:UpdatedByID;not null"`
	UserID       uint64     `gorm:"column:UserID;primary_key;not null;auto_increment"`
	UserTypeID   uint64     `gorm:"column:UserTypeID;not null"`
	UserTypeName string     `gorm:"column:UserTypeName;not null"`
	Username     string     `gorm:"column:Username;not null"`
}

// TableName overrides the default tablename generated by GORM
func (UserORM) TableName() string {
	return "users"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *User) ToORM(ctx context.Context) (UserORM, error) {
	to := UserORM{}
	var err error
	if prehook, ok := interface{}(m).(UserWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.UserID = m.UserID
	to.Username = m.Username
	to.IdType = int32(m.IdType)
	to.IdNumber = m.IdNumber
	to.UserTypeID = m.UserTypeID
	to.UserTypeName = m.UserTypeName
	to.CompanyID = m.CompanyID
	to.CompanyName = m.CompanyName
	to.IsBlocked = m.IsBlocked
	to.CreatedByID = m.CreatedByID
	to.UpdatedByID = m.UpdatedByID
	to.DeletedByID = m.DeletedByID
	to.Authority = m.Authority
	to.TokenFcm = m.TokenFcm
	if m.CreatedAt != nil {
		t := m.CreatedAt.AsTime()
		to.CreatedAt = &t
	}
	if m.UpdatedAt != nil {
		t := m.UpdatedAt.AsTime()
		to.UpdatedAt = &t
	}
	if m.DeletedAt != nil {
		t := m.DeletedAt.AsTime()
		to.DeletedAt = &t
	}
	if posthook, ok := interface{}(m).(UserWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *UserORM) ToPB(ctx context.Context) (User, error) {
	to := User{}
	var err error
	if prehook, ok := interface{}(m).(UserWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.UserID = m.UserID
	to.Username = m.Username
	to.IdType = User_IDType(m.IdType)
	to.IdNumber = m.IdNumber
	to.UserTypeID = m.UserTypeID
	to.UserTypeName = m.UserTypeName
	to.CompanyID = m.CompanyID
	to.CompanyName = m.CompanyName
	to.IsBlocked = m.IsBlocked
	to.CreatedByID = m.CreatedByID
	to.UpdatedByID = m.UpdatedByID
	to.DeletedByID = m.DeletedByID
	to.Authority = m.Authority
	to.TokenFcm = m.TokenFcm
	if m.CreatedAt != nil {
		to.CreatedAt = timestamppb.New(*m.CreatedAt)
	}
	if m.UpdatedAt != nil {
		to.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}
	if m.DeletedAt != nil {
		to.DeletedAt = timestamppb.New(*m.DeletedAt)
	}
	if posthook, ok := interface{}(m).(UserWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type User the arg will be the target, the caller the one being converted from

// UserBeforeToORM called before default ToORM code
type UserWithBeforeToORM interface {
	BeforeToORM(context.Context, *UserORM) error
}

// UserAfterToORM called after default ToORM code
type UserWithAfterToORM interface {
	AfterToORM(context.Context, *UserORM) error
}

// UserBeforeToPB called before default ToPB code
type UserWithBeforeToPB interface {
	BeforeToPB(context.Context, *User) error
}

// UserAfterToPB called after default ToPB code
type UserWithAfterToPB interface {
	AfterToPB(context.Context, *User) error
}

type UserLoginORM struct {
	Authority    string `gorm:"column:Authority"`
	CompanyID    uint64 `gorm:"column:CompanyID"`
	CompanyName  string `gorm:"column:CompanyName"`
	IsBlocked    bool   `gorm:"column:IsBlocked"`
	TokenFcm     string `gorm:"column:TokenFcm"`
	UserID       uint64 `gorm:"column:UserID;primary_key;not null;auto_increment"`
	UserTypeID   uint64 `gorm:"column:UserTypeID;not null"`
	UserTypeName string `gorm:"column:UserTypeName;not null"`
	Username     string `gorm:"column:Username;not null"`
}

// TableName overrides the default tablename generated by GORM
func (UserLoginORM) TableName() string {
	return "users"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *UserLogin) ToORM(ctx context.Context) (UserLoginORM, error) {
	to := UserLoginORM{}
	var err error
	if prehook, ok := interface{}(m).(UserLoginWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.UserID = m.UserID
	to.Username = m.Username
	to.UserTypeID = m.UserTypeID
	to.UserTypeName = m.UserTypeName
	to.CompanyID = m.CompanyID
	to.CompanyName = m.CompanyName
	to.IsBlocked = m.IsBlocked
	to.Authority = m.Authority
	to.TokenFcm = m.TokenFcm
	if posthook, ok := interface{}(m).(UserLoginWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *UserLoginORM) ToPB(ctx context.Context) (UserLogin, error) {
	to := UserLogin{}
	var err error
	if prehook, ok := interface{}(m).(UserLoginWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.UserID = m.UserID
	to.Username = m.Username
	to.UserTypeID = m.UserTypeID
	to.UserTypeName = m.UserTypeName
	to.CompanyID = m.CompanyID
	to.CompanyName = m.CompanyName
	to.IsBlocked = m.IsBlocked
	to.Authority = m.Authority
	to.TokenFcm = m.TokenFcm
	if posthook, ok := interface{}(m).(UserLoginWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type UserLogin the arg will be the target, the caller the one being converted from

// UserLoginBeforeToORM called before default ToORM code
type UserLoginWithBeforeToORM interface {
	BeforeToORM(context.Context, *UserLoginORM) error
}

// UserLoginAfterToORM called after default ToORM code
type UserLoginWithAfterToORM interface {
	AfterToORM(context.Context, *UserLoginORM) error
}

// UserLoginBeforeToPB called before default ToPB code
type UserLoginWithBeforeToPB interface {
	BeforeToPB(context.Context, *UserLogin) error
}

// UserLoginAfterToPB called after default ToPB code
type UserLoginWithAfterToPB interface {
	AfterToPB(context.Context, *UserLogin) error
}

// DefaultCreateUser executes a basic gorm create call
func DefaultCreateUser(ctx context.Context, in *User, db *gorm.DB) (*User, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type UserORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadUser(ctx context.Context, in *User, db *gorm.DB) (*User, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.UserID == 0 {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &UserORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := UserORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(UserORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type UserORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteUser(ctx context.Context, in *User, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.UserID == 0 {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&UserORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type UserORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteUserSet(ctx context.Context, in []*User, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	var err error
	keys := []uint64{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.UserID == 0 {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.UserID)
	}
	if hook, ok := (interface{}(&UserORM{})).(UserORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("user_id in (?)", keys).Delete(&UserORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&UserORM{})).(UserORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type UserORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*User, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*User, *gorm.DB) error
}

// DefaultStrictUpdateUser clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateUser(ctx context.Context, in *User, db *gorm.DB) (*User, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateUser")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &UserORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("UserID=?", ormObj.UserID).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type UserORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchUser executes a basic gorm update call with patch behavior
func DefaultPatchUser(ctx context.Context, in *User, updateMask *field_mask.FieldMask, db *gorm.DB) (*User, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj User
	var err error
	if hook, ok := interface{}(&pbObj).(UserWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&pbObj).(UserWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskUser(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(UserWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateUser(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(UserWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type UserWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *User, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type UserWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *User, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type UserWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *User, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type UserWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *User, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetUser executes a bulk gorm update call with patch behavior
func DefaultPatchSetUser(ctx context.Context, objects []*User, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*User, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*User, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchUser(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskUser patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskUser(ctx context.Context, patchee *User, patcher *User, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*User, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	var updatedCreatedAt bool
	var updatedUpdatedAt bool
	var updatedDeletedAt bool
	for i, f := range updateMask.Paths {
		if f == prefix+"UserID" {
			patchee.UserID = patcher.UserID
			continue
		}
		if f == prefix+"Username" {
			patchee.Username = patcher.Username
			continue
		}
		if f == prefix+"IdType" {
			patchee.IdType = patcher.IdType
			continue
		}
		if f == prefix+"IdNumber" {
			patchee.IdNumber = patcher.IdNumber
			continue
		}
		if f == prefix+"UserTypeID" {
			patchee.UserTypeID = patcher.UserTypeID
			continue
		}
		if f == prefix+"UserTypeName" {
			patchee.UserTypeName = patcher.UserTypeName
			continue
		}
		if f == prefix+"CompanyID" {
			patchee.CompanyID = patcher.CompanyID
			continue
		}
		if f == prefix+"CompanyName" {
			patchee.CompanyName = patcher.CompanyName
			continue
		}
		if f == prefix+"IsBlocked" {
			patchee.IsBlocked = patcher.IsBlocked
			continue
		}
		if f == prefix+"CreatedByID" {
			patchee.CreatedByID = patcher.CreatedByID
			continue
		}
		if f == prefix+"UpdatedByID" {
			patchee.UpdatedByID = patcher.UpdatedByID
			continue
		}
		if f == prefix+"DeletedByID" {
			patchee.DeletedByID = patcher.DeletedByID
			continue
		}
		if f == prefix+"Authority" {
			patchee.Authority = patcher.Authority
			continue
		}
		if f == prefix+"TokenFcm" {
			patchee.TokenFcm = patcher.TokenFcm
			continue
		}
		if !updatedCreatedAt && strings.HasPrefix(f, prefix+"CreatedAt.") {
			if patcher.CreatedAt == nil {
				patchee.CreatedAt = nil
				continue
			}
			if patchee.CreatedAt == nil {
				patchee.CreatedAt = &timestamppb.Timestamp{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"CreatedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.CreatedAt, patchee.CreatedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"CreatedAt" {
			updatedCreatedAt = true
			patchee.CreatedAt = patcher.CreatedAt
			continue
		}
		if !updatedUpdatedAt && strings.HasPrefix(f, prefix+"UpdatedAt.") {
			if patcher.UpdatedAt == nil {
				patchee.UpdatedAt = nil
				continue
			}
			if patchee.UpdatedAt == nil {
				patchee.UpdatedAt = &timestamppb.Timestamp{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"UpdatedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.UpdatedAt, patchee.UpdatedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"UpdatedAt" {
			updatedUpdatedAt = true
			patchee.UpdatedAt = patcher.UpdatedAt
			continue
		}
		if !updatedDeletedAt && strings.HasPrefix(f, prefix+"DeletedAt.") {
			if patcher.DeletedAt == nil {
				patchee.DeletedAt = nil
				continue
			}
			if patchee.DeletedAt == nil {
				patchee.DeletedAt = &timestamppb.Timestamp{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"DeletedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.DeletedAt, patchee.DeletedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"DeletedAt" {
			updatedDeletedAt = true
			patchee.DeletedAt = patcher.DeletedAt
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListUser executes a gorm list call
func DefaultListUser(ctx context.Context, db *gorm.DB) ([]*User, error) {
	in := User{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &UserORM{}, &User{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("UserID")
	ormResponse := []UserORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*User{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type UserORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]UserORM) error
}

// DefaultCreateUserLogin executes a basic gorm create call
func DefaultCreateUserLogin(ctx context.Context, in *UserLogin, db *gorm.DB) (*UserLogin, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserLoginORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserLoginORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type UserLoginORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserLoginORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadUserLogin(ctx context.Context, in *UserLogin, db *gorm.DB) (*UserLogin, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.UserID == 0 {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(UserLoginORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &UserLoginORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserLoginORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := UserLoginORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(UserLoginORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type UserLoginORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserLoginORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserLoginORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteUserLogin(ctx context.Context, in *UserLogin, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.UserID == 0 {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(UserLoginORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&UserLoginORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(UserLoginORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type UserLoginORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserLoginORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteUserLoginSet(ctx context.Context, in []*UserLogin, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	var err error
	keys := []uint64{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.UserID == 0 {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.UserID)
	}
	if hook, ok := (interface{}(&UserLoginORM{})).(UserLoginORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("user_id in (?)", keys).Delete(&UserLoginORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&UserLoginORM{})).(UserLoginORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type UserLoginORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*UserLogin, *gorm.DB) (*gorm.DB, error)
}
type UserLoginORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*UserLogin, *gorm.DB) error
}

// DefaultStrictUpdateUserLogin clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateUserLogin(ctx context.Context, in *UserLogin, db *gorm.DB) (*UserLogin, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateUserLogin")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &UserLoginORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("UserID=?", ormObj.UserID).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(UserLoginORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(UserLoginORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserLoginORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type UserLoginORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserLoginORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserLoginORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchUserLogin executes a basic gorm update call with patch behavior
func DefaultPatchUserLogin(ctx context.Context, in *UserLogin, updateMask *field_mask.FieldMask, db *gorm.DB) (*UserLogin, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj UserLogin
	var err error
	if hook, ok := interface{}(&pbObj).(UserLoginWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&pbObj).(UserLoginWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskUserLogin(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(UserLoginWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateUserLogin(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(UserLoginWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type UserLoginWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *UserLogin, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type UserLoginWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *UserLogin, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type UserLoginWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *UserLogin, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type UserLoginWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *UserLogin, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetUserLogin executes a bulk gorm update call with patch behavior
func DefaultPatchSetUserLogin(ctx context.Context, objects []*UserLogin, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*UserLogin, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*UserLogin, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchUserLogin(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskUserLogin patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskUserLogin(ctx context.Context, patchee *UserLogin, patcher *UserLogin, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*UserLogin, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"UserID" {
			patchee.UserID = patcher.UserID
			continue
		}
		if f == prefix+"Username" {
			patchee.Username = patcher.Username
			continue
		}
		if f == prefix+"UserTypeID" {
			patchee.UserTypeID = patcher.UserTypeID
			continue
		}
		if f == prefix+"UserTypeName" {
			patchee.UserTypeName = patcher.UserTypeName
			continue
		}
		if f == prefix+"CompanyID" {
			patchee.CompanyID = patcher.CompanyID
			continue
		}
		if f == prefix+"CompanyName" {
			patchee.CompanyName = patcher.CompanyName
			continue
		}
		if f == prefix+"IsBlocked" {
			patchee.IsBlocked = patcher.IsBlocked
			continue
		}
		if f == prefix+"Authority" {
			patchee.Authority = patcher.Authority
			continue
		}
		if f == prefix+"TokenFcm" {
			patchee.TokenFcm = patcher.TokenFcm
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListUserLogin executes a gorm list call
func DefaultListUserLogin(ctx context.Context, db *gorm.DB) ([]*UserLogin, error) {
	in := UserLogin{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserLoginORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &UserLoginORM{}, &UserLogin{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserLoginORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("UserID")
	ormResponse := []UserLoginORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserLoginORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*UserLogin{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type UserLoginORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserLoginORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserLoginORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]UserLoginORM) error
}
