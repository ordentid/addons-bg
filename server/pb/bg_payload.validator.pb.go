// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bg_payload.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Task) Validate() error {
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
func (this *TransactionTask) Validate() error {
	if this.Transaction != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Transaction); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Transaction", err)
		}
	}
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	return nil
}
func (this *HealthCheckRequest) Validate() error {
	return nil
}
func (this *HealthCheckResponse) Validate() error {
	return nil
}
func (this *GetThirdPartyIDRequest) Validate() error {
	return nil
}
func (this *GetThirdPartyIDResponse) Validate() error {
	return nil
}
func (this *GetTransactionTaskRequest) Validate() error {
	return nil
}
func (this *GetTransactionTaskResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	if this.Pagination != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Pagination); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Pagination", err)
		}
	}
	return nil
}
func (this *CreateTransactionTaskThirdParty) Validate() error {
	return nil
}
func (this *CreateTransactionTaskRequest) Validate() error {
	for _, item := range this.ThirdParty {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ThirdParty", err)
			}
		}
	}
	return nil
}
func (this *CreateTransactionTaskResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetTransactionRequest) Validate() error {
	if this.Transaction != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Transaction); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Transaction", err)
		}
	}
	return nil
}
func (this *GetTransactionResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	if this.Pagination != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Pagination); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Pagination", err)
		}
	}
	return nil
}
func (this *GetTransactionDetailRequest) Validate() error {
	return nil
}
func (this *GetTransactionDetailResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateTransactionRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateTransactionResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *UpdateTransactionRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *UpdateTransactionResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
