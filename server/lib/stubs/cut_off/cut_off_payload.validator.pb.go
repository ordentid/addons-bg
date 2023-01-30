// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cut_off_payload.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "github.com/mwitkow/go-proto-validators"
	regexp "regexp"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *PaginationResponse) Validate() error {
	return nil
}
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
func (this *Sort) Validate() error {
	return nil
}
func (this *ListProductUsedReq) Validate() error {
	return nil
}
func (this *ListProductUsedRes) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *Module) Validate() error {
	return nil
}

var _regex_ScheduleTime_StartTime = regexp.MustCompile(`^\d{2}:\d{2}$`)
var _regex_ScheduleTime_EndTime = regexp.MustCompile(`^\d{2}:\d{2}$`)

func (this *ScheduleTime) Validate() error {
	if !_regex_ScheduleTime_StartTime.MatchString(this.StartTime) {
		return github_com_mwitkow_go_proto_validators.FieldError("StartTime", fmt.Errorf(`value '%v' must be a string conforming to regex "^\\d{2}:\\d{2}$"`, this.StartTime))
	}
	if !_regex_ScheduleTime_EndTime.MatchString(this.EndTime) {
		return github_com_mwitkow_go_proto_validators.FieldError("EndTime", fmt.Errorf(`value '%v' must be a string conforming to regex "^\\d{2}:\\d{2}$"`, this.EndTime))
	}
	return nil
}

var _regex_CutOffData_ScheduleName = regexp.MustCompile(`^.*\S.*$`)
var _regex_CutOffData_ScheduleDescription = regexp.MustCompile(`^.*\S.*$`)

func (this *CutOffData) Validate() error {
	if !_regex_CutOffData_ScheduleName.MatchString(this.ScheduleName) {
		return github_com_mwitkow_go_proto_validators.FieldError("ScheduleName", fmt.Errorf(`value '%v' must be a string conforming to regex "^.*\\S.*$"`, this.ScheduleName))
	}
	if !_regex_CutOffData_ScheduleDescription.MatchString(this.ScheduleDescription) {
		return github_com_mwitkow_go_proto_validators.FieldError("ScheduleDescription", fmt.Errorf(`value '%v' must be a string conforming to regex "^.*\\S.*$"`, this.ScheduleDescription))
	}
	for _, item := range this.Modules {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Modules", err)
			}
		}
	}
	for _, item := range this.ScheduleTimes {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ScheduleTimes", err)
			}
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
	if this.DeletedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DeletedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DeletedAt", err)
		}
	}
	return nil
}
func (this *CreateCutOffTaskRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *TaskData) Validate() error {
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateCutOffTaskResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CutOffTaskActionRequest) Validate() error {
	return nil
}
func (this *CutOffTaskActionResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *ListCutOffTaskRequest) Validate() error {
	if this.Filter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Filter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Filter", err)
		}
	}
	return nil
}
func (this *ListCutOffTaskRequest_TaskListFilter) Validate() error {
	return nil
}
func (this *ListCutOffTaskResponse) Validate() error {
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
func (this *FileListCutOffTaskRequest) Validate() error {
	if this.Filter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Filter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Filter", err)
		}
	}
	return nil
}
func (this *GetCutOffTaskByIDRequest) Validate() error {
	return nil
}
func (this *GetCutOffTaskByIDResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *ListCutOffDataRequest) Validate() error {
	if this.Filter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Filter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Filter", err)
		}
	}
	return nil
}
func (this *ListCutOffDataRequest_DataListFilter) Validate() error {
	return nil
}
func (this *ListCutOffDataScheduleTodayRequest) Validate() error {
	return nil
}
func (this *ListCutOffDataScheduleTodayRespons) Validate() error {
	if this.TimeExecute != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TimeExecute); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TimeExecute", err)
		}
	}
	return nil
}
func (this *ListCutOffDataScheduleRequest) Validate() error {
	return nil
}
func (this *Dayproduct) Validate() error {
	for _, item := range this.ScheduleTime {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ScheduleTime", err)
			}
		}
	}
	return nil
}
func (this *ScheduledCutOff) Validate() error {
	for _, item := range this.Product {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Product", err)
			}
		}
	}
	return nil
}
func (this *ListCutOffScheduledDataResponse) Validate() error {
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
func (this *ScheduleTimeByTime) Validate() error {
	return nil
}
func (this *ScheduledCutOffByTime) Validate() error {
	for _, item := range this.TimeDay {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("TimeDay", err)
			}
		}
	}
	return nil
}
func (this *ListCutOffScheduledDataByTimeResponse) Validate() error {
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
func (this *ListCutOffDataResponse) Validate() error {
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
func (this *CutOffTemplate) Validate() error {
	return nil
}
func (this *ListCutOffDataTemplateResponse) Validate() error {
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
func (this *GetCutOffDataByIDRequest) Validate() error {
	return nil
}
func (this *GetCutOffDataByIDResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *SaveCutOffDataRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *SaveCutOffDataResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *ArrayString) Validate() error {
	return nil
}

var _regex_HolidaySchedule_Date = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
var _regex_HolidaySchedule_Label = regexp.MustCompile(`^.*\S.*$`)
var _regex_HolidaySchedule_Pattern = regexp.MustCompile(`^(Yearly)|(Monthly)|(yearly)|(monthly)$`)

func (this *HolidaySchedule) Validate() error {
	if !_regex_HolidaySchedule_Date.MatchString(this.Date) {
		return github_com_mwitkow_go_proto_validators.FieldError("Date", fmt.Errorf(`value '%v' must be a string conforming to regex "^\\d{4}-\\d{2}-\\d{2}$"`, this.Date))
	}
	if !_regex_HolidaySchedule_Label.MatchString(this.Label) {
		return github_com_mwitkow_go_proto_validators.FieldError("Label", fmt.Errorf(`value '%v' must be a string conforming to regex "^.*\\S.*$"`, this.Label))
	}
	if !_regex_HolidaySchedule_Pattern.MatchString(this.Pattern) {
		return github_com_mwitkow_go_proto_validators.FieldError("Pattern", fmt.Errorf(`value '%v' must be a string conforming to regex "^(Yearly)|(Monthly)|(yearly)|(monthly)$"`, this.Pattern))
	}
	return nil
}
func (this *HolidayModule) Validate() error {
	return nil
}
func (this *GetHolidayDataRes) Validate() error {
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
	if this.DeletedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DeletedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DeletedAt", err)
		}
	}
	return nil
}

var _regex_HolidayData_ScheduleName = regexp.MustCompile(`^.*\S.*$`)
var _regex_HolidayData_ScheduleDescription = regexp.MustCompile(`^.*\S.*$`)

func (this *HolidayData) Validate() error {
	if !_regex_HolidayData_ScheduleName.MatchString(this.ScheduleName) {
		return github_com_mwitkow_go_proto_validators.FieldError("ScheduleName", fmt.Errorf(`value '%v' must be a string conforming to regex "^.*\\S.*$"`, this.ScheduleName))
	}
	if !_regex_HolidayData_ScheduleDescription.MatchString(this.ScheduleDescription) {
		return github_com_mwitkow_go_proto_validators.FieldError("ScheduleDescription", fmt.Errorf(`value '%v' must be a string conforming to regex "^.*\\S.*$"`, this.ScheduleDescription))
	}
	for _, item := range this.ScheduleDates {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ScheduleDates", err)
			}
		}
	}
	for _, item := range this.Modules {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Modules", err)
			}
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
	if this.DeletedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DeletedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DeletedAt", err)
		}
	}
	return nil
}
func (this *HolidayTaskData) Validate() error {
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateHolidayTaskRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateHolidayTaskResponse) Validate() error {
	return nil
}
func (this *GetHolidayTaskRequest) Validate() error {
	return nil
}
func (this *GetHolidayTaskResponse) Validate() error {
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
func (this *GetHolidayTaskByIDRequest) Validate() error {
	return nil
}
func (this *GetHolidayTaskByIDResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *GetHolidayRequest) Validate() error {
	if this.Holiday != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Holiday); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Holiday", err)
		}
	}
	return nil
}
func (this *GetHolidayResponse) Validate() error {
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
func (this *GetHolidayByFeatureIDRequest) Validate() error {
	return nil
}
func (this *GetHolidayByFeatureIDResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateHolidayRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateHolidayResponse) Validate() error {
	return nil
}
func (this *DeleteHolidayRequest) Validate() error {
	return nil
}
func (this *HolidayTaskActionRequest) Validate() error {
	return nil
}
func (this *HolidayTaskActionResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *GetHolidaySchedulesRequest) Validate() error {
	return nil
}
func (this *GetHolidaySchedulesResponse) Validate() error {
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
func (this *HolidayScheduleRes) Validate() error {
	return nil
}
func (this *DownloadListHolidayTaskRequest) Validate() error {
	return nil
}
func (this *HolidayNotificationData) Validate() error {
	return nil
}
func (this *CutOffNotificationData) Validate() error {
	return nil
}
func (this *GetAvailableTimeRequest) Validate() error {
	return nil
}
func (this *GetAvailableTimeResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *Availabletime) Validate() error {
	return nil
}
