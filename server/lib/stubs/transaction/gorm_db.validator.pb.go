// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gorm_db.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Currency) Validate() error {
	if this.CreatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", err)
		}
	}
	if this.UpdatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UpdatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UpdatedAt", err)
		}
	}
	return nil
}
func (this *Transaction) Validate() error {
	if this.ScheduleDate != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ScheduleDate); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ScheduleDate", err)
		}
	}
	if this.DeletedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DeletedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DeletedAt", err)
		}
	}
	if this.CreatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", err)
		}
	}
	if this.UpdatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UpdatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UpdatedAt", err)
		}
	}
	return nil
}
func (this *Scheduler) Validate() error {
	if this.Transaction != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Transaction); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Transaction", err)
		}
	}
	if this.ScheduleDate != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ScheduleDate); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ScheduleDate", err)
		}
	}
	if this.CreatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", err)
		}
	}
	return nil
}
func (this *HostToHostConfig) Validate() error {
	return nil
}
