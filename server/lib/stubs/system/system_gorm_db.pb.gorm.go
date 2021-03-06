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

type SystemSchemesORM struct {
	CreatedAt   *time.Time `gorm:"not null"`
	CreatedByID uint64     `gorm:"not null"`
	Description string     `gorm:"not null"`
	Id          uint64     `gorm:"primary_key;not null"`
	Key         string     `gorm:"not null"`
	ProductID   string     `gorm:"not null"`
	Type        string     `gorm:"not null"`
	UpdatedAt   *time.Time `gorm:"not null"`
	UpdatedByID uint64     `gorm:"not null"`
	Value       string     `gorm:"not null"`
}

// TableName overrides the default tablename generated by GORM
func (SystemSchemesORM) TableName() string {
	return "system_params"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *SystemSchemes) ToORM(ctx context.Context) (SystemSchemesORM, error) {
	to := SystemSchemesORM{}
	var err error
	if prehook, ok := interface{}(m).(SystemSchemesWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.ProductID = m.ProductID
	to.Key = m.Key
	to.Value = m.Value
	to.Type = m.Type
	to.Description = m.Description
	if m.CreatedAt != nil {
		t := m.CreatedAt.AsTime()
		to.CreatedAt = &t
	}
	to.CreatedByID = m.CreatedByID
	if m.UpdatedAt != nil {
		t := m.UpdatedAt.AsTime()
		to.UpdatedAt = &t
	}
	to.UpdatedByID = m.UpdatedByID
	if posthook, ok := interface{}(m).(SystemSchemesWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *SystemSchemesORM) ToPB(ctx context.Context) (SystemSchemes, error) {
	to := SystemSchemes{}
	var err error
	if prehook, ok := interface{}(m).(SystemSchemesWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.ProductID = m.ProductID
	to.Key = m.Key
	to.Value = m.Value
	to.Type = m.Type
	to.Description = m.Description
	if m.CreatedAt != nil {
		to.CreatedAt = timestamppb.New(*m.CreatedAt)
	}
	to.CreatedByID = m.CreatedByID
	if m.UpdatedAt != nil {
		to.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}
	to.UpdatedByID = m.UpdatedByID
	if posthook, ok := interface{}(m).(SystemSchemesWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type SystemSchemes the arg will be the target, the caller the one being converted from

// SystemSchemesBeforeToORM called before default ToORM code
type SystemSchemesWithBeforeToORM interface {
	BeforeToORM(context.Context, *SystemSchemesORM) error
}

// SystemSchemesAfterToORM called after default ToORM code
type SystemSchemesWithAfterToORM interface {
	AfterToORM(context.Context, *SystemSchemesORM) error
}

// SystemSchemesBeforeToPB called before default ToPB code
type SystemSchemesWithBeforeToPB interface {
	BeforeToPB(context.Context, *SystemSchemes) error
}

// SystemSchemesAfterToPB called after default ToPB code
type SystemSchemesWithAfterToPB interface {
	AfterToPB(context.Context, *SystemSchemes) error
}

// DefaultCreateSystemSchemes executes a basic gorm create call
func DefaultCreateSystemSchemes(ctx context.Context, in *SystemSchemes, db *gorm.DB) (*SystemSchemes, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SystemSchemesORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SystemSchemesORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type SystemSchemesORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type SystemSchemesORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadSystemSchemes(ctx context.Context, in *SystemSchemes, db *gorm.DB) (*SystemSchemes, error) {
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
	if hook, ok := interface{}(&ormObj).(SystemSchemesORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &SystemSchemesORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SystemSchemesORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := SystemSchemesORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(SystemSchemesORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type SystemSchemesORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type SystemSchemesORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type SystemSchemesORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteSystemSchemes(ctx context.Context, in *SystemSchemes, db *gorm.DB) error {
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
	if hook, ok := interface{}(&ormObj).(SystemSchemesORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&SystemSchemesORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(SystemSchemesORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type SystemSchemesORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type SystemSchemesORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteSystemSchemesSet(ctx context.Context, in []*SystemSchemes, db *gorm.DB) error {
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
	if hook, ok := (interface{}(&SystemSchemesORM{})).(SystemSchemesORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&SystemSchemesORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&SystemSchemesORM{})).(SystemSchemesORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type SystemSchemesORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*SystemSchemes, *gorm.DB) (*gorm.DB, error)
}
type SystemSchemesORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*SystemSchemes, *gorm.DB) error
}

// DefaultStrictUpdateSystemSchemes clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateSystemSchemes(ctx context.Context, in *SystemSchemes, db *gorm.DB) (*SystemSchemes, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateSystemSchemes")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &SystemSchemesORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(SystemSchemesORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(SystemSchemesORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SystemSchemesORMWithAfterStrictUpdateSave); ok {
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

type SystemSchemesORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type SystemSchemesORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type SystemSchemesORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchSystemSchemes executes a basic gorm update call with patch behavior
func DefaultPatchSystemSchemes(ctx context.Context, in *SystemSchemes, updateMask *field_mask.FieldMask, db *gorm.DB) (*SystemSchemes, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj SystemSchemes
	var err error
	if hook, ok := interface{}(&pbObj).(SystemSchemesWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadSystemSchemes(ctx, &SystemSchemes{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(SystemSchemesWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskSystemSchemes(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(SystemSchemesWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateSystemSchemes(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(SystemSchemesWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type SystemSchemesWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *SystemSchemes, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type SystemSchemesWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *SystemSchemes, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type SystemSchemesWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *SystemSchemes, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type SystemSchemesWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *SystemSchemes, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetSystemSchemes executes a bulk gorm update call with patch behavior
func DefaultPatchSetSystemSchemes(ctx context.Context, objects []*SystemSchemes, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*SystemSchemes, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*SystemSchemes, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchSystemSchemes(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskSystemSchemes patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskSystemSchemes(ctx context.Context, patchee *SystemSchemes, patcher *SystemSchemes, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*SystemSchemes, error) {
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
		if f == prefix+"ProductID" {
			patchee.ProductID = patcher.ProductID
			continue
		}
		if f == prefix+"Key" {
			patchee.Key = patcher.Key
			continue
		}
		if f == prefix+"Value" {
			patchee.Value = patcher.Value
			continue
		}
		if f == prefix+"Type" {
			patchee.Type = patcher.Type
			continue
		}
		if f == prefix+"Description" {
			patchee.Description = patcher.Description
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
		if f == prefix+"CreatedByID" {
			patchee.CreatedByID = patcher.CreatedByID
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
		if f == prefix+"UpdatedByID" {
			patchee.UpdatedByID = patcher.UpdatedByID
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListSystemSchemes executes a gorm list call
func DefaultListSystemSchemes(ctx context.Context, db *gorm.DB) ([]*SystemSchemes, error) {
	in := SystemSchemes{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SystemSchemesORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &SystemSchemesORM{}, &SystemSchemes{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SystemSchemesORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []SystemSchemesORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(SystemSchemesORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*SystemSchemes{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type SystemSchemesORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type SystemSchemesORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type SystemSchemesORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]SystemSchemesORM) error
}
