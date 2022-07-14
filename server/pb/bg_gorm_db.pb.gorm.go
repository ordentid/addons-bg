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

type ThirdPartyORM struct {
	CreatedAt    *time.Time `gorm:"not null"`
	Id           uint64     `gorm:"primary_key;not null"`
	Name         string     `gorm:"not null"`
	ThirdPartyID uint64     `gorm:"not null"`
	UpdatedAt    *time.Time `gorm:"not null"`
}

// TableName overrides the default tablename generated by GORM
func (ThirdPartyORM) TableName() string {
	return "third_parties"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *ThirdParty) ToORM(ctx context.Context) (ThirdPartyORM, error) {
	to := ThirdPartyORM{}
	var err error
	if prehook, ok := interface{}(m).(ThirdPartyWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.ThirdPartyID = m.ThirdPartyID
	to.Name = m.Name
	if m.CreatedAt != nil {
		t := m.CreatedAt.AsTime()
		to.CreatedAt = &t
	}
	if m.UpdatedAt != nil {
		t := m.UpdatedAt.AsTime()
		to.UpdatedAt = &t
	}
	if posthook, ok := interface{}(m).(ThirdPartyWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *ThirdPartyORM) ToPB(ctx context.Context) (ThirdParty, error) {
	to := ThirdParty{}
	var err error
	if prehook, ok := interface{}(m).(ThirdPartyWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.ThirdPartyID = m.ThirdPartyID
	to.Name = m.Name
	if m.CreatedAt != nil {
		to.CreatedAt = timestamppb.New(*m.CreatedAt)
	}
	if m.UpdatedAt != nil {
		to.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}
	if posthook, ok := interface{}(m).(ThirdPartyWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type ThirdParty the arg will be the target, the caller the one being converted from

// ThirdPartyBeforeToORM called before default ToORM code
type ThirdPartyWithBeforeToORM interface {
	BeforeToORM(context.Context, *ThirdPartyORM) error
}

// ThirdPartyAfterToORM called after default ToORM code
type ThirdPartyWithAfterToORM interface {
	AfterToORM(context.Context, *ThirdPartyORM) error
}

// ThirdPartyBeforeToPB called before default ToPB code
type ThirdPartyWithBeforeToPB interface {
	BeforeToPB(context.Context, *ThirdParty) error
}

// ThirdPartyAfterToPB called after default ToPB code
type ThirdPartyWithAfterToPB interface {
	AfterToPB(context.Context, *ThirdParty) error
}

type TransactionORM struct {
	Amount          float64 `gorm:"not null"`
	ApplicantName   string  `gorm:"not null"`
	BeneficiaryName string  `gorm:"not null"`
	BgStatus        int32   `gorm:"not null"`
	BgType          int32   `gorm:"not null"`
	ChannelID       uint64
	ChannelName     string
	ClaimPeriod     uint64     `gorm:"not null"`
	ClosingDate     string     `gorm:"not null"`
	CompanyID       uint64     `gorm:"not null"`
	CreatedAt       *time.Time `gorm:"not null"`
	CreatedByID     uint64     `gorm:"not null"`
	Currency        string     `gorm:"not null"`
	DocumentPath    string
	EffectiveDate   string `gorm:"not null"`
	ExpiryDate      string `gorm:"not null"`
	Id              uint64 `gorm:"primary_key;not null"`
	IssueDate       string `gorm:"not null"`
	MaturityDate    string
	ReferenceNo     string     `gorm:"not null"`
	RegistrationNo  string     `gorm:"not null"`
	Remark          string     `gorm:"not null"`
	Status          int32      `gorm:"not null"`
	ThirdPartyID    uint64     `gorm:"not null"`
	ThirdPartyName  string     `gorm:"not null"`
	TransactionID   uint64     `gorm:"not null"`
	UpdatedAt       *time.Time `gorm:"not null"`
	UpdatedByID     uint64     `gorm:"not null"`
}

// TableName overrides the default tablename generated by GORM
func (TransactionORM) TableName() string {
	return "transactions"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Transaction) ToORM(ctx context.Context) (TransactionORM, error) {
	to := TransactionORM{}
	var err error
	if prehook, ok := interface{}(m).(TransactionWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.TransactionID = m.TransactionID
	to.ThirdPartyID = m.ThirdPartyID
	to.ThirdPartyName = m.ThirdPartyName
	to.ReferenceNo = m.ReferenceNo
	to.RegistrationNo = m.RegistrationNo
	to.ApplicantName = m.ApplicantName
	to.BeneficiaryName = m.BeneficiaryName
	to.IssueDate = m.IssueDate
	to.EffectiveDate = m.EffectiveDate
	to.ExpiryDate = m.ExpiryDate
	to.ClaimPeriod = m.ClaimPeriod
	to.ClosingDate = m.ClosingDate
	to.Currency = m.Currency
	to.Amount = m.Amount
	to.Remark = m.Remark
	to.BgStatus = int32(m.BgStatus)
	to.ChannelID = m.ChannelID
	to.ChannelName = m.ChannelName
	to.BgType = int32(m.BgType)
	to.DocumentPath = m.DocumentPath
	to.CompanyID = m.CompanyID
	to.Status = int32(m.Status)
	to.MaturityDate = m.MaturityDate
	to.CreatedByID = m.CreatedByID
	to.UpdatedByID = m.UpdatedByID
	if m.CreatedAt != nil {
		t := m.CreatedAt.AsTime()
		to.CreatedAt = &t
	}
	if m.UpdatedAt != nil {
		t := m.UpdatedAt.AsTime()
		to.UpdatedAt = &t
	}
	if posthook, ok := interface{}(m).(TransactionWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *TransactionORM) ToPB(ctx context.Context) (Transaction, error) {
	to := Transaction{}
	var err error
	if prehook, ok := interface{}(m).(TransactionWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.TransactionID = m.TransactionID
	to.ThirdPartyID = m.ThirdPartyID
	to.ThirdPartyName = m.ThirdPartyName
	to.ReferenceNo = m.ReferenceNo
	to.RegistrationNo = m.RegistrationNo
	to.ApplicantName = m.ApplicantName
	to.BeneficiaryName = m.BeneficiaryName
	to.IssueDate = m.IssueDate
	to.EffectiveDate = m.EffectiveDate
	to.ExpiryDate = m.ExpiryDate
	to.ClaimPeriod = m.ClaimPeriod
	to.ClosingDate = m.ClosingDate
	to.Currency = m.Currency
	to.Amount = m.Amount
	to.Remark = m.Remark
	to.BgStatus = BgStatus(m.BgStatus)
	to.ChannelID = m.ChannelID
	to.ChannelName = m.ChannelName
	to.BgType = BgType(m.BgType)
	to.DocumentPath = m.DocumentPath
	to.CompanyID = m.CompanyID
	to.Status = TransactionStatus(m.Status)
	to.MaturityDate = m.MaturityDate
	to.CreatedByID = m.CreatedByID
	to.UpdatedByID = m.UpdatedByID
	if m.CreatedAt != nil {
		to.CreatedAt = timestamppb.New(*m.CreatedAt)
	}
	if m.UpdatedAt != nil {
		to.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}
	if posthook, ok := interface{}(m).(TransactionWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Transaction the arg will be the target, the caller the one being converted from

// TransactionBeforeToORM called before default ToORM code
type TransactionWithBeforeToORM interface {
	BeforeToORM(context.Context, *TransactionORM) error
}

// TransactionAfterToORM called after default ToORM code
type TransactionWithAfterToORM interface {
	AfterToORM(context.Context, *TransactionORM) error
}

// TransactionBeforeToPB called before default ToPB code
type TransactionWithBeforeToPB interface {
	BeforeToPB(context.Context, *Transaction) error
}

// TransactionAfterToPB called after default ToPB code
type TransactionWithAfterToPB interface {
	AfterToPB(context.Context, *Transaction) error
}

// DefaultCreateThirdParty executes a basic gorm create call
func DefaultCreateThirdParty(ctx context.Context, in *ThirdParty, db *gorm.DB) (*ThirdParty, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ThirdPartyORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ThirdPartyORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type ThirdPartyORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ThirdPartyORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadThirdParty(ctx context.Context, in *ThirdParty, db *gorm.DB) (*ThirdParty, error) {
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
	if hook, ok := interface{}(&ormObj).(ThirdPartyORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &ThirdPartyORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ThirdPartyORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := ThirdPartyORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(ThirdPartyORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type ThirdPartyORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ThirdPartyORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ThirdPartyORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteThirdParty(ctx context.Context, in *ThirdParty, db *gorm.DB) error {
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
	if hook, ok := interface{}(&ormObj).(ThirdPartyORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&ThirdPartyORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(ThirdPartyORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type ThirdPartyORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ThirdPartyORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteThirdPartySet(ctx context.Context, in []*ThirdParty, db *gorm.DB) error {
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
	if hook, ok := (interface{}(&ThirdPartyORM{})).(ThirdPartyORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&ThirdPartyORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&ThirdPartyORM{})).(ThirdPartyORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type ThirdPartyORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*ThirdParty, *gorm.DB) (*gorm.DB, error)
}
type ThirdPartyORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*ThirdParty, *gorm.DB) error
}

// DefaultStrictUpdateThirdParty clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateThirdParty(ctx context.Context, in *ThirdParty, db *gorm.DB) (*ThirdParty, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateThirdParty")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &ThirdPartyORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(ThirdPartyORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(ThirdPartyORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ThirdPartyORMWithAfterStrictUpdateSave); ok {
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

type ThirdPartyORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ThirdPartyORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ThirdPartyORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchThirdParty executes a basic gorm update call with patch behavior
func DefaultPatchThirdParty(ctx context.Context, in *ThirdParty, updateMask *field_mask.FieldMask, db *gorm.DB) (*ThirdParty, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj ThirdParty
	var err error
	if hook, ok := interface{}(&pbObj).(ThirdPartyWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadThirdParty(ctx, &ThirdParty{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(ThirdPartyWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskThirdParty(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(ThirdPartyWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateThirdParty(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(ThirdPartyWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type ThirdPartyWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *ThirdParty, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type ThirdPartyWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *ThirdParty, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type ThirdPartyWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *ThirdParty, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type ThirdPartyWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *ThirdParty, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetThirdParty executes a bulk gorm update call with patch behavior
func DefaultPatchSetThirdParty(ctx context.Context, objects []*ThirdParty, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*ThirdParty, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*ThirdParty, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchThirdParty(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskThirdParty patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskThirdParty(ctx context.Context, patchee *ThirdParty, patcher *ThirdParty, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*ThirdParty, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	var updatedCreatedAt bool
	var updatedUpdatedAt bool
	for i, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"ThirdPartyID" {
			patchee.ThirdPartyID = patcher.ThirdPartyID
			continue
		}
		if f == prefix+"Name" {
			patchee.Name = patcher.Name
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
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListThirdParty executes a gorm list call
func DefaultListThirdParty(ctx context.Context, db *gorm.DB) ([]*ThirdParty, error) {
	in := ThirdParty{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ThirdPartyORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &ThirdPartyORM{}, &ThirdParty{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ThirdPartyORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []ThirdPartyORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ThirdPartyORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*ThirdParty{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type ThirdPartyORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ThirdPartyORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ThirdPartyORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]ThirdPartyORM) error
}

// DefaultCreateTransaction executes a basic gorm create call
func DefaultCreateTransaction(ctx context.Context, in *Transaction, db *gorm.DB) (*Transaction, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type TransactionORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadTransaction(ctx context.Context, in *Transaction, db *gorm.DB) (*Transaction, error) {
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
	if hook, ok := interface{}(&ormObj).(TransactionORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &TransactionORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := TransactionORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(TransactionORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type TransactionORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteTransaction(ctx context.Context, in *Transaction, db *gorm.DB) error {
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
	if hook, ok := interface{}(&ormObj).(TransactionORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&TransactionORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type TransactionORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteTransactionSet(ctx context.Context, in []*Transaction, db *gorm.DB) error {
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
	if hook, ok := (interface{}(&TransactionORM{})).(TransactionORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&TransactionORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&TransactionORM{})).(TransactionORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type TransactionORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*Transaction, *gorm.DB) (*gorm.DB, error)
}
type TransactionORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*Transaction, *gorm.DB) error
}

// DefaultStrictUpdateTransaction clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateTransaction(ctx context.Context, in *Transaction, db *gorm.DB) (*Transaction, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateTransaction")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &TransactionORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(TransactionORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithAfterStrictUpdateSave); ok {
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

type TransactionORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchTransaction executes a basic gorm update call with patch behavior
func DefaultPatchTransaction(ctx context.Context, in *Transaction, updateMask *field_mask.FieldMask, db *gorm.DB) (*Transaction, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj Transaction
	var err error
	if hook, ok := interface{}(&pbObj).(TransactionWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadTransaction(ctx, &Transaction{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(TransactionWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskTransaction(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(TransactionWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateTransaction(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(TransactionWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type TransactionWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *Transaction, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type TransactionWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *Transaction, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type TransactionWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *Transaction, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type TransactionWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *Transaction, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetTransaction executes a bulk gorm update call with patch behavior
func DefaultPatchSetTransaction(ctx context.Context, objects []*Transaction, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*Transaction, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*Transaction, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchTransaction(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskTransaction patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskTransaction(ctx context.Context, patchee *Transaction, patcher *Transaction, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*Transaction, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	var updatedCreatedAt bool
	var updatedUpdatedAt bool
	for i, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"TransactionID" {
			patchee.TransactionID = patcher.TransactionID
			continue
		}
		if f == prefix+"ThirdPartyID" {
			patchee.ThirdPartyID = patcher.ThirdPartyID
			continue
		}
		if f == prefix+"ThirdPartyName" {
			patchee.ThirdPartyName = patcher.ThirdPartyName
			continue
		}
		if f == prefix+"ReferenceNo" {
			patchee.ReferenceNo = patcher.ReferenceNo
			continue
		}
		if f == prefix+"RegistrationNo" {
			patchee.RegistrationNo = patcher.RegistrationNo
			continue
		}
		if f == prefix+"ApplicantName" {
			patchee.ApplicantName = patcher.ApplicantName
			continue
		}
		if f == prefix+"BeneficiaryName" {
			patchee.BeneficiaryName = patcher.BeneficiaryName
			continue
		}
		if f == prefix+"IssueDate" {
			patchee.IssueDate = patcher.IssueDate
			continue
		}
		if f == prefix+"EffectiveDate" {
			patchee.EffectiveDate = patcher.EffectiveDate
			continue
		}
		if f == prefix+"ExpiryDate" {
			patchee.ExpiryDate = patcher.ExpiryDate
			continue
		}
		if f == prefix+"ClaimPeriod" {
			patchee.ClaimPeriod = patcher.ClaimPeriod
			continue
		}
		if f == prefix+"ClosingDate" {
			patchee.ClosingDate = patcher.ClosingDate
			continue
		}
		if f == prefix+"Currency" {
			patchee.Currency = patcher.Currency
			continue
		}
		if f == prefix+"Amount" {
			patchee.Amount = patcher.Amount
			continue
		}
		if f == prefix+"Remark" {
			patchee.Remark = patcher.Remark
			continue
		}
		if f == prefix+"BgStatus" {
			patchee.BgStatus = patcher.BgStatus
			continue
		}
		if f == prefix+"ChannelID" {
			patchee.ChannelID = patcher.ChannelID
			continue
		}
		if f == prefix+"ChannelName" {
			patchee.ChannelName = patcher.ChannelName
			continue
		}
		if f == prefix+"BgType" {
			patchee.BgType = patcher.BgType
			continue
		}
		if f == prefix+"DocumentPath" {
			patchee.DocumentPath = patcher.DocumentPath
			continue
		}
		if f == prefix+"CompanyID" {
			patchee.CompanyID = patcher.CompanyID
			continue
		}
		if f == prefix+"Status" {
			patchee.Status = patcher.Status
			continue
		}
		if f == prefix+"MaturityDate" {
			patchee.MaturityDate = patcher.MaturityDate
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
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListTransaction executes a gorm list call
func DefaultListTransaction(ctx context.Context, db *gorm.DB) ([]*Transaction, error) {
	in := Transaction{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &TransactionORM{}, &Transaction{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []TransactionORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*Transaction{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type TransactionORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]TransactionORM) error
}
