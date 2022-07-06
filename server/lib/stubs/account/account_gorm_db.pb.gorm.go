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
	CreatedAt *time.Time
	DeletedAt *time.Time
	Id        uint64 `gorm:"primary_key;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"default:user;not null"`
	UpdatedAt *time.Time
	Username  string `gorm:"unique;not null"`
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
	to.Id = m.Id
	to.Username = m.Username
	to.Password = m.Password
	to.Role = m.Role
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
	to.Id = m.Id
	to.Username = m.Username
	to.Password = m.Password
	to.Role = m.Role
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

type AccountORM struct {
	AccessLevel      string     `gorm:"column:AccessLevel;not null"`
	AccountAlias     string     `gorm:"column:AccountAlias;type:jsonb;default:[];not null"`
	AccountCurrency  string     `gorm:"column:AccountCurrency;not null"`
	AccountID        uint64     `gorm:"column:AccountID;primary_key;not null;auto_increment"`
	AccountName      string     `gorm:"column:AccountName;not null"`
	AccountNumber    string     `gorm:"column:AccountNumber;not null"`
	AccountStatus    string     `gorm:"column:AccountStatus;not null"`
	AccountType      string     `gorm:"column:AccountType;not null"`
	CompanyID        uint64     `gorm:"column:CompanyID;not null"`
	CreatedAt        *time.Time `gorm:"column:CreatedAt"`
	CreatedByID      uint64     `gorm:"column:CreatedByID;not null"`
	DeletedAt        *time.Time `gorm:"column:DeletedAt"`
	DeletedByID      uint64     `gorm:"column:DeletedByID;not null"`
	Disabled         bool       `gorm:"column:Disabled"`
	IsOwnedByCompany bool       `gorm:"column:IsOwnedByCompany;not null"`
	RoleID           uint64     `gorm:"column:RoleID"`
	UpdatedAt        *time.Time `gorm:"column:UpdatedAt"`
	UpdatedByID      uint64     `gorm:"column:UpdatedByID;not null"`
}

// TableName overrides the default tablename generated by GORM
func (AccountORM) TableName() string {
	return "bank_accounts"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Account) ToORM(ctx context.Context) (AccountORM, error) {
	to := AccountORM{}
	var err error
	if prehook, ok := interface{}(m).(AccountWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.AccountID = m.AccountID
	to.CompanyID = m.CompanyID
	to.AccountNumber = m.AccountNumber
	to.AccountAlias = m.AccountAlias
	to.AccountName = m.AccountName
	to.AccountType = m.AccountType
	to.AccountStatus = m.AccountStatus
	to.AccountCurrency = m.AccountCurrency
	to.AccessLevel = m.AccessLevel
	to.IsOwnedByCompany = m.IsOwnedByCompany
	to.CreatedByID = m.CreatedByID
	to.UpdatedByID = m.UpdatedByID
	to.DeletedByID = m.DeletedByID
	to.RoleID = m.RoleID
	to.Disabled = m.Disabled
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
	if posthook, ok := interface{}(m).(AccountWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *AccountORM) ToPB(ctx context.Context) (Account, error) {
	to := Account{}
	var err error
	if prehook, ok := interface{}(m).(AccountWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.AccountID = m.AccountID
	to.CompanyID = m.CompanyID
	to.AccountNumber = m.AccountNumber
	to.AccountAlias = m.AccountAlias
	to.AccountName = m.AccountName
	to.AccountType = m.AccountType
	to.AccountStatus = m.AccountStatus
	to.AccountCurrency = m.AccountCurrency
	to.AccessLevel = m.AccessLevel
	to.IsOwnedByCompany = m.IsOwnedByCompany
	to.CreatedByID = m.CreatedByID
	to.UpdatedByID = m.UpdatedByID
	to.DeletedByID = m.DeletedByID
	to.RoleID = m.RoleID
	to.Disabled = m.Disabled
	if m.CreatedAt != nil {
		to.CreatedAt = timestamppb.New(*m.CreatedAt)
	}
	if m.UpdatedAt != nil {
		to.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}
	if m.DeletedAt != nil {
		to.DeletedAt = timestamppb.New(*m.DeletedAt)
	}
	if posthook, ok := interface{}(m).(AccountWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Account the arg will be the target, the caller the one being converted from

// AccountBeforeToORM called before default ToORM code
type AccountWithBeforeToORM interface {
	BeforeToORM(context.Context, *AccountORM) error
}

// AccountAfterToORM called after default ToORM code
type AccountWithAfterToORM interface {
	AfterToORM(context.Context, *AccountORM) error
}

// AccountBeforeToPB called before default ToPB code
type AccountWithBeforeToPB interface {
	BeforeToPB(context.Context, *Account) error
}

// AccountAfterToPB called after default ToPB code
type AccountWithAfterToPB interface {
	AfterToPB(context.Context, *Account) error
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
	if ormObj.Id == 0 {
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
	if ormObj.Id == 0 {
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
		if ormObj.Id == 0 {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&UserORM{})).(UserORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&UserORM{}).Error
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
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
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
	pbReadRes, err := DefaultReadUser(ctx, &User{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
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
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"Username" {
			patchee.Username = patcher.Username
			continue
		}
		if f == prefix+"Password" {
			patchee.Password = patcher.Password
			continue
		}
		if f == prefix+"Role" {
			patchee.Role = patcher.Role
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
	db = db.Order("id")
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

// DefaultCreateAccount executes a basic gorm create call
func DefaultCreateAccount(ctx context.Context, in *Account, db *gorm.DB) (*Account, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AccountORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AccountORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type AccountORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type AccountORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadAccount(ctx context.Context, in *Account, db *gorm.DB) (*Account, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.AccountID == 0 {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(AccountORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &AccountORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AccountORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := AccountORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(AccountORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type AccountORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type AccountORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type AccountORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteAccount(ctx context.Context, in *Account, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.AccountID == 0 {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(AccountORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&AccountORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(AccountORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type AccountORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type AccountORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteAccountSet(ctx context.Context, in []*Account, db *gorm.DB) error {
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
		if ormObj.AccountID == 0 {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.AccountID)
	}
	if hook, ok := (interface{}(&AccountORM{})).(AccountORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("account_id in (?)", keys).Delete(&AccountORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&AccountORM{})).(AccountORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type AccountORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*Account, *gorm.DB) (*gorm.DB, error)
}
type AccountORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*Account, *gorm.DB) error
}

// DefaultStrictUpdateAccount clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateAccount(ctx context.Context, in *Account, db *gorm.DB) (*Account, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateAccount")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &AccountORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("AccountID=?", ormObj.AccountID).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(AccountORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(AccountORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AccountORMWithAfterStrictUpdateSave); ok {
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

type AccountORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type AccountORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type AccountORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchAccount executes a basic gorm update call with patch behavior
func DefaultPatchAccount(ctx context.Context, in *Account, updateMask *field_mask.FieldMask, db *gorm.DB) (*Account, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj Account
	var err error
	if hook, ok := interface{}(&pbObj).(AccountWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&pbObj).(AccountWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskAccount(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(AccountWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateAccount(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(AccountWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type AccountWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *Account, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type AccountWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *Account, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type AccountWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *Account, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type AccountWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *Account, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetAccount executes a bulk gorm update call with patch behavior
func DefaultPatchSetAccount(ctx context.Context, objects []*Account, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*Account, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*Account, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchAccount(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskAccount patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskAccount(ctx context.Context, patchee *Account, patcher *Account, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*Account, error) {
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
		if f == prefix+"AccountID" {
			patchee.AccountID = patcher.AccountID
			continue
		}
		if f == prefix+"CompanyID" {
			patchee.CompanyID = patcher.CompanyID
			continue
		}
		if f == prefix+"AccountNumber" {
			patchee.AccountNumber = patcher.AccountNumber
			continue
		}
		if f == prefix+"AccountAlias" {
			patchee.AccountAlias = patcher.AccountAlias
			continue
		}
		if f == prefix+"AccountName" {
			patchee.AccountName = patcher.AccountName
			continue
		}
		if f == prefix+"AccountType" {
			patchee.AccountType = patcher.AccountType
			continue
		}
		if f == prefix+"AccountStatus" {
			patchee.AccountStatus = patcher.AccountStatus
			continue
		}
		if f == prefix+"AccountCurrency" {
			patchee.AccountCurrency = patcher.AccountCurrency
			continue
		}
		if f == prefix+"AccessLevel" {
			patchee.AccessLevel = patcher.AccessLevel
			continue
		}
		if f == prefix+"IsOwnedByCompany" {
			patchee.IsOwnedByCompany = patcher.IsOwnedByCompany
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
		if f == prefix+"RoleID" {
			patchee.RoleID = patcher.RoleID
			continue
		}
		if f == prefix+"Disabled" {
			patchee.Disabled = patcher.Disabled
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

// DefaultListAccount executes a gorm list call
func DefaultListAccount(ctx context.Context, db *gorm.DB) ([]*Account, error) {
	in := Account{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AccountORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &AccountORM{}, &Account{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AccountORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("AccountID")
	ormResponse := []AccountORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AccountORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*Account{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type AccountORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type AccountORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type AccountORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]AccountORM) error
}
