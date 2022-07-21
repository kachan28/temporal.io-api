// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/api/enums/v1/workflow.proto

package enums

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
	strconv "strconv"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Defines how new runs of a workflow with a particular ID may or may not be allowed. Note that
// it is *never* valid to have two actively running instances of the same workflow id.
type WorkflowIdReusePolicy int32

const (
	WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED WorkflowIdReusePolicy = 0
	// Allow starting a workflow execution using the same workflow id.
	WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE WorkflowIdReusePolicy = 1
	// Allow starting a workflow execution using the same workflow id, only when the last
	// execution's final state is one of [terminated, cancelled, timed out, failed].
	WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY WorkflowIdReusePolicy = 2
	// Do not permit re-use of the workflow id for this workflow. Future start workflow requests
	// could potentially change the policy, allowing re-use of the workflow id.
	WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE WorkflowIdReusePolicy = 3
	// If a workflow is running using the same workflow ID, terminate it and start a new one.
	// If no running workflow, then the behavior is the same as ALLOW_DUPLICATE
	WORKFLOW_ID_REUSE_POLICY_TERMINATE_IF_RUNNING WorkflowIdReusePolicy = 4
)

var WorkflowIdReusePolicy_name = map[int32]string{
	0: "Unspecified",
	1: "AllowDuplicate",
	2: "AllowDuplicateFailedOnly",
	3: "RejectDuplicate",
	4: "TerminateIfRunning",
}

var WorkflowIdReusePolicy_value = map[string]int32{
	"Unspecified":              0,
	"AllowDuplicate":           1,
	"AllowDuplicateFailedOnly": 2,
	"RejectDuplicate":          3,
	"TerminateIfRunning":       4,
}

func (WorkflowIdReusePolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_939fa9511cc117f0, []int{0}
}

// Defines how child workflows will react to their parent completing
type ParentClosePolicy int32

const (
	PARENT_CLOSE_POLICY_UNSPECIFIED ParentClosePolicy = 0
	// The child workflow will also terminate
	PARENT_CLOSE_POLICY_TERMINATE ParentClosePolicy = 1
	// The child workflow will do nothing
	PARENT_CLOSE_POLICY_ABANDON ParentClosePolicy = 2
	// Cancellation will be requested of the child workflow
	PARENT_CLOSE_POLICY_REQUEST_CANCEL ParentClosePolicy = 3
)

var ParentClosePolicy_name = map[int32]string{
	0: "Unspecified",
	1: "Terminate",
	2: "Abandon",
	3: "RequestCancel",
}

var ParentClosePolicy_value = map[string]int32{
	"Unspecified":   0,
	"Terminate":     1,
	"Abandon":       2,
	"RequestCancel": 3,
}

func (ParentClosePolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_939fa9511cc117f0, []int{1}
}

type ContinueAsNewInitiator int32

const (
	CONTINUE_AS_NEW_INITIATOR_UNSPECIFIED ContinueAsNewInitiator = 0
	// The workflow itself requested to continue as new
	CONTINUE_AS_NEW_INITIATOR_WORKFLOW ContinueAsNewInitiator = 1
	// The workflow continued as new because it is retrying
	CONTINUE_AS_NEW_INITIATOR_RETRY ContinueAsNewInitiator = 2
	// The workflow continued as new because cron has triggered a new execution
	CONTINUE_AS_NEW_INITIATOR_CRON_SCHEDULE ContinueAsNewInitiator = 3
)

var ContinueAsNewInitiator_name = map[int32]string{
	0: "Unspecified",
	1: "Workflow",
	2: "Retry",
	3: "CronSchedule",
}

var ContinueAsNewInitiator_value = map[string]int32{
	"Unspecified":  0,
	"Workflow":     1,
	"Retry":        2,
	"CronSchedule": 3,
}

func (ContinueAsNewInitiator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_939fa9511cc117f0, []int{2}
}

// (-- api-linter: core::0216::synonyms=disabled
//     aip.dev/not-precedent: There is WorkflowExecutionState already in another package. --)
type WorkflowExecutionStatus int32

const (
	WORKFLOW_EXECUTION_STATUS_UNSPECIFIED WorkflowExecutionStatus = 0
	// Value 1 is hardcoded in SQL persistence.
	WORKFLOW_EXECUTION_STATUS_RUNNING          WorkflowExecutionStatus = 1
	WORKFLOW_EXECUTION_STATUS_COMPLETED        WorkflowExecutionStatus = 2
	WORKFLOW_EXECUTION_STATUS_FAILED           WorkflowExecutionStatus = 3
	WORKFLOW_EXECUTION_STATUS_CANCELED         WorkflowExecutionStatus = 4
	WORKFLOW_EXECUTION_STATUS_TERMINATED       WorkflowExecutionStatus = 5
	WORKFLOW_EXECUTION_STATUS_CONTINUED_AS_NEW WorkflowExecutionStatus = 6
	WORKFLOW_EXECUTION_STATUS_TIMED_OUT        WorkflowExecutionStatus = 7
)

var WorkflowExecutionStatus_name = map[int32]string{
	0: "Unspecified",
	1: "Running",
	2: "Completed",
	3: "Failed",
	4: "Canceled",
	5: "Terminated",
	6: "ContinuedAsNew",
	7: "TimedOut",
}

var WorkflowExecutionStatus_value = map[string]int32{
	"Unspecified":    0,
	"Running":        1,
	"Completed":      2,
	"Failed":         3,
	"Canceled":       4,
	"Terminated":     5,
	"ContinuedAsNew": 6,
	"TimedOut":       7,
}

func (WorkflowExecutionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_939fa9511cc117f0, []int{3}
}

type PendingActivityState int32

const (
	PENDING_ACTIVITY_STATE_UNSPECIFIED      PendingActivityState = 0
	PENDING_ACTIVITY_STATE_SCHEDULED        PendingActivityState = 1
	PENDING_ACTIVITY_STATE_STARTED          PendingActivityState = 2
	PENDING_ACTIVITY_STATE_CANCEL_REQUESTED PendingActivityState = 3
)

var PendingActivityState_name = map[int32]string{
	0: "Unspecified",
	1: "Scheduled",
	2: "Started",
	3: "CancelRequested",
}

var PendingActivityState_value = map[string]int32{
	"Unspecified":     0,
	"Scheduled":       1,
	"Started":         2,
	"CancelRequested": 3,
}

func (PendingActivityState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_939fa9511cc117f0, []int{4}
}

type PendingWorkflowTaskState int32

const (
	PENDING_WORKFLOW_TASK_STATE_UNSPECIFIED PendingWorkflowTaskState = 0
	PENDING_WORKFLOW_TASK_STATE_SCHEDULED   PendingWorkflowTaskState = 1
	PENDING_WORKFLOW_TASK_STATE_STARTED     PendingWorkflowTaskState = 2
)

var PendingWorkflowTaskState_name = map[int32]string{
	0: "Unspecified",
	1: "Scheduled",
	2: "Started",
}

var PendingWorkflowTaskState_value = map[string]int32{
	"Unspecified": 0,
	"Scheduled":   1,
	"Started":     2,
}

func (PendingWorkflowTaskState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_939fa9511cc117f0, []int{5}
}

type HistoryEventFilterType int32

const (
	HISTORY_EVENT_FILTER_TYPE_UNSPECIFIED HistoryEventFilterType = 0
	HISTORY_EVENT_FILTER_TYPE_ALL_EVENT   HistoryEventFilterType = 1
	HISTORY_EVENT_FILTER_TYPE_CLOSE_EVENT HistoryEventFilterType = 2
)

var HistoryEventFilterType_name = map[int32]string{
	0: "Unspecified",
	1: "AllEvent",
	2: "CloseEvent",
}

var HistoryEventFilterType_value = map[string]int32{
	"Unspecified": 0,
	"AllEvent":    1,
	"CloseEvent":  2,
}

func (HistoryEventFilterType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_939fa9511cc117f0, []int{6}
}

type RetryState int32

const (
	RETRY_STATE_UNSPECIFIED              RetryState = 0
	RETRY_STATE_IN_PROGRESS              RetryState = 1
	RETRY_STATE_NON_RETRYABLE_FAILURE    RetryState = 2
	RETRY_STATE_TIMEOUT                  RetryState = 3
	RETRY_STATE_MAXIMUM_ATTEMPTS_REACHED RetryState = 4
	RETRY_STATE_RETRY_POLICY_NOT_SET     RetryState = 5
	RETRY_STATE_INTERNAL_SERVER_ERROR    RetryState = 6
	RETRY_STATE_CANCEL_REQUESTED         RetryState = 7
)

var RetryState_name = map[int32]string{
	0: "Unspecified",
	1: "InProgress",
	2: "NonRetryableFailure",
	3: "Timeout",
	4: "MaximumAttemptsReached",
	5: "RetryPolicyNotSet",
	6: "InternalServerError",
	7: "CancelRequested",
}

var RetryState_value = map[string]int32{
	"Unspecified":            0,
	"InProgress":             1,
	"NonRetryableFailure":    2,
	"Timeout":                3,
	"MaximumAttemptsReached": 4,
	"RetryPolicyNotSet":      5,
	"InternalServerError":    6,
	"CancelRequested":        7,
}

func (RetryState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_939fa9511cc117f0, []int{7}
}

type TimeoutType int32

const (
	TIMEOUT_TYPE_UNSPECIFIED       TimeoutType = 0
	TIMEOUT_TYPE_START_TO_CLOSE    TimeoutType = 1
	TIMEOUT_TYPE_SCHEDULE_TO_START TimeoutType = 2
	TIMEOUT_TYPE_SCHEDULE_TO_CLOSE TimeoutType = 3
	TIMEOUT_TYPE_HEARTBEAT         TimeoutType = 4
)

var TimeoutType_name = map[int32]string{
	0: "Unspecified",
	1: "StartToClose",
	2: "ScheduleToStart",
	3: "ScheduleToClose",
	4: "Heartbeat",
}

var TimeoutType_value = map[string]int32{
	"Unspecified":     0,
	"StartToClose":    1,
	"ScheduleToStart": 2,
	"ScheduleToClose": 3,
	"Heartbeat":       4,
}

func (TimeoutType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_939fa9511cc117f0, []int{8}
}

func init() {
	proto.RegisterEnum("temporal.api.enums.v1.WorkflowIdReusePolicy", WorkflowIdReusePolicy_name, WorkflowIdReusePolicy_value)
	proto.RegisterEnum("temporal.api.enums.v1.ParentClosePolicy", ParentClosePolicy_name, ParentClosePolicy_value)
	proto.RegisterEnum("temporal.api.enums.v1.ContinueAsNewInitiator", ContinueAsNewInitiator_name, ContinueAsNewInitiator_value)
	proto.RegisterEnum("temporal.api.enums.v1.WorkflowExecutionStatus", WorkflowExecutionStatus_name, WorkflowExecutionStatus_value)
	proto.RegisterEnum("temporal.api.enums.v1.PendingActivityState", PendingActivityState_name, PendingActivityState_value)
	proto.RegisterEnum("temporal.api.enums.v1.PendingWorkflowTaskState", PendingWorkflowTaskState_name, PendingWorkflowTaskState_value)
	proto.RegisterEnum("temporal.api.enums.v1.HistoryEventFilterType", HistoryEventFilterType_name, HistoryEventFilterType_value)
	proto.RegisterEnum("temporal.api.enums.v1.RetryState", RetryState_name, RetryState_value)
	proto.RegisterEnum("temporal.api.enums.v1.TimeoutType", TimeoutType_name, TimeoutType_value)
}

func init() {
	proto.RegisterFile("temporal/api/enums/v1/workflow.proto", fileDescriptor_939fa9511cc117f0)
}

var fileDescriptor_939fa9511cc117f0 = []byte{
	// 970 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x96, 0xbf, 0x6f, 0xe4, 0x44,
	0x14, 0xc7, 0x77, 0xbc, 0xb9, 0x9c, 0x34, 0x08, 0x69, 0x30, 0xe4, 0x87, 0xc8, 0xe1, 0xfc, 0xbe,
	0x5c, 0x7c, 0xdc, 0x46, 0x11, 0x14, 0x68, 0xa9, 0x26, 0xf6, 0xdb, 0x64, 0x38, 0xef, 0xd8, 0x8c,
	0xc7, 0xc9, 0x85, 0x66, 0xb4, 0xe4, 0xcc, 0xc9, 0xba, 0x64, 0xbd, 0xda, 0xf5, 0x26, 0xa4, 0xa3,
	0xa7, 0xa1, 0xa3, 0xa0, 0xa0, 0xa1, 0x40, 0x54, 0x34, 0x57, 0xd2, 0x53, 0xa6, 0xbc, 0x92, 0x6c,
	0x84, 0x84, 0xa8, 0xee, 0x4f, 0x40, 0xb6, 0xd7, 0xd1, 0x66, 0x6f, 0xbd, 0xa2, 0xb3, 0x66, 0x3e,
	0x33, 0xf3, 0xde, 0xf7, 0x7d, 0xdf, 0x78, 0xf0, 0x46, 0x12, 0x9e, 0x75, 0xe2, 0x6e, 0xeb, 0x74,
	0xa7, 0xd5, 0x89, 0x76, 0xc2, 0x76, 0xff, 0xac, 0xb7, 0x73, 0xbe, 0xbb, 0x73, 0x11, 0x77, 0x5f,
	0x7e, 0x73, 0x1a, 0x5f, 0xd4, 0x3a, 0xdd, 0x38, 0x89, 0xf5, 0xb9, 0x82, 0xaa, 0xb5, 0x3a, 0x51,
	0x2d, 0xa3, 0x6a, 0xe7, 0xbb, 0xe6, 0xf7, 0x1a, 0x9e, 0x3b, 0x1a, 0x92, 0xec, 0xb9, 0x08, 0xfb,
	0xbd, 0xd0, 0x8b, 0x4f, 0xa3, 0x93, 0x4b, 0xfd, 0x11, 0xde, 0x38, 0x72, 0xc5, 0xd3, 0x86, 0xe3,
	0x1e, 0x29, 0x66, 0x2b, 0x01, 0x81, 0x0f, 0xca, 0x73, 0x1d, 0x66, 0x1d, 0xab, 0x80, 0xfb, 0x1e,
	0x58, 0xac, 0xc1, 0xc0, 0x26, 0x15, 0xfd, 0x63, 0xfc, 0xa8, 0x94, 0xa4, 0x4e, 0x3a, 0x6a, 0x07,
	0x9e, 0xc3, 0x2c, 0x2a, 0x81, 0x20, 0xfd, 0x33, 0xfc, 0xe9, 0xff, 0xa5, 0x55, 0x83, 0x32, 0x07,
	0x6c, 0xe5, 0x72, 0xe7, 0x98, 0x68, 0xfa, 0x13, 0xbc, 0x5d, 0xba, 0x52, 0xc0, 0x17, 0x60, 0xc9,
	0x91, 0x83, 0xaa, 0xfa, 0x2e, 0x7e, 0x52, 0x8a, 0x4b, 0x10, 0x4d, 0xc6, 0xd3, 0x23, 0x58, 0x43,
	0x89, 0x80, 0x73, 0xc6, 0xf7, 0xc9, 0x8c, 0xf9, 0x0b, 0xc2, 0xef, 0x79, 0xad, 0x6e, 0xd8, 0x4e,
	0xac, 0xd3, 0xf8, 0x56, 0x89, 0x75, 0xbc, 0xec, 0x51, 0x01, 0x5c, 0x2a, 0xcb, 0x71, 0xcb, 0x44,
	0x58, 0xc5, 0x1f, 0x4d, 0x82, 0x6e, 0x0f, 0x22, 0x48, 0x5f, 0xc6, 0x4b, 0x93, 0x10, 0xba, 0x47,
	0xb9, 0xed, 0x72, 0xa2, 0xe9, 0x0f, 0xf1, 0xda, 0x24, 0x40, 0xc0, 0x97, 0x01, 0xf8, 0x52, 0x59,
	0x94, 0x5b, 0xe0, 0x90, 0xaa, 0xf9, 0x07, 0xc2, 0xf3, 0x56, 0xdc, 0x4e, 0xa2, 0x76, 0x3f, 0xa4,
	0x3d, 0x1e, 0x5e, 0xb0, 0x76, 0x94, 0x44, 0xad, 0x24, 0xee, 0xea, 0xdb, 0x78, 0xd3, 0x72, 0xb9,
	0x64, 0x3c, 0x00, 0x45, 0x7d, 0xc5, 0xe1, 0x48, 0x31, 0xce, 0x24, 0xa3, 0xd2, 0x15, 0x63, 0x11,
	0x3f, 0xc4, 0x6b, 0xe5, 0x68, 0xa1, 0x1c, 0x41, 0x69, 0xfa, 0xe5, 0x9c, 0x00, 0x29, 0xd2, 0xda,
	0x3c, 0xc6, 0x5b, 0xe5, 0x90, 0x25, 0x5c, 0xae, 0x7c, 0xeb, 0x00, 0xec, 0xc0, 0x01, 0x52, 0x35,
	0xff, 0xd6, 0xf0, 0x42, 0x61, 0x3a, 0xf8, 0x36, 0x3c, 0xe9, 0x27, 0x51, 0xdc, 0xf6, 0x93, 0x56,
	0xd2, 0xef, 0xa5, 0x09, 0xdc, 0x56, 0x0d, 0x9e, 0x81, 0x15, 0x48, 0x96, 0x2e, 0x96, 0x54, 0x06,
	0xfe, 0x58, 0x02, 0x9b, 0x78, 0xb5, 0x1c, 0x2d, 0x8a, 0x8a, 0xf4, 0x2d, 0xbc, 0x5e, 0x8e, 0x59,
	0x6e, 0xd3, 0x73, 0x40, 0x82, 0x4d, 0x34, 0x7d, 0x03, 0xaf, 0x94, 0x83, 0xb9, 0x15, 0x49, 0x35,
	0x95, 0x6d, 0xca, 0x76, 0x59, 0x89, 0xc0, 0x26, 0x33, 0x77, 0xfa, 0xe7, 0x2d, 0xee, 0xd6, 0x16,
	0x36, 0xb9, 0xa7, 0xd7, 0xb0, 0x39, 0x2d, 0xc0, 0x5c, 0x55, 0x7b, 0x28, 0x2b, 0x99, 0x9d, 0x9e,
	0x90, 0x64, 0xcd, 0xb4, 0x63, 0x02, 0x49, 0xee, 0x9b, 0xaf, 0x10, 0xfe, 0xc0, 0x0b, 0xdb, 0xcf,
	0xa3, 0xf6, 0x0b, 0x7a, 0x92, 0x44, 0xe7, 0x51, 0x72, 0x99, 0xaa, 0x1c, 0x66, 0x46, 0x03, 0x6e,
	0x33, 0xbe, 0xaf, 0xa8, 0x25, 0xd9, 0x21, 0x93, 0xc7, 0xd9, 0x7a, 0x18, 0x53, 0x78, 0x03, 0xaf,
	0x94, 0x70, 0x45, 0x35, 0x6d, 0x82, 0xf4, 0x35, 0x6c, 0x94, 0x51, 0x92, 0x8a, 0x5c, 0xdb, 0xc7,
	0x78, 0xab, 0x84, 0xc9, 0x25, 0x2b, 0x4c, 0x9e, 0x4a, 0x6c, 0xfe, 0x84, 0xf0, 0xe2, 0x30, 0xee,
	0xc2, 0x26, 0xb2, 0xd5, 0x7b, 0x99, 0xc7, 0x3e, 0xb2, 0xd3, 0xad, 0x0a, 0x92, 0xfa, 0x4f, 0x27,
	0x26, 0xb0, 0x8d, 0x37, 0xa7, 0xc1, 0xa3, 0x59, 0x6c, 0xe1, 0xf5, 0xa9, 0x68, 0x91, 0x8a, 0xf9,
	0x23, 0xc2, 0xf3, 0x07, 0x51, 0x2f, 0x89, 0xbb, 0x97, 0x70, 0x1e, 0xb6, 0x93, 0x46, 0x74, 0x9a,
	0x84, 0x5d, 0x79, 0xd9, 0x09, 0xd3, 0xe3, 0x0e, 0x98, 0x2f, 0x5d, 0x71, 0xac, 0xe0, 0x30, 0xed,
	0xe3, 0x06, 0x73, 0x24, 0x08, 0x25, 0x8f, 0xbd, 0xf1, 0xc8, 0xb6, 0xf0, 0x7a, 0x39, 0x4a, 0x1d,
	0x27, 0x1f, 0x25, 0x68, 0xfa, 0x9e, 0xf9, 0x3d, 0x91, 0xa3, 0x9a, 0xf9, 0xb3, 0x86, 0xb1, 0x08,
	0x93, 0xee, 0xb0, 0xca, 0x4b, 0x78, 0x21, 0x6b, 0xcf, 0x89, 0xca, 0x8c, 0x4d, 0x32, 0xae, 0x3c,
	0xe1, 0xee, 0x0b, 0xf0, 0x7d, 0x82, 0xd2, 0xce, 0x1a, 0x9d, 0xe4, 0x2e, 0xcf, 0x1b, 0x9d, 0xee,
	0x39, 0xf9, 0xa5, 0x1c, 0x08, 0x20, 0x9a, 0xbe, 0x80, 0xdf, 0x1f, 0xc5, 0x52, 0xeb, 0xa5, 0xc6,
	0xab, 0xa6, 0xde, 0x1f, 0x9d, 0x68, 0xd2, 0x67, 0xac, 0x19, 0x34, 0x15, 0x95, 0x12, 0x9a, 0x9e,
	0xf4, 0x95, 0x00, 0x9a, 0x2a, 0x4f, 0x66, 0x52, 0x87, 0x8d, 0x92, 0xf9, 0xf7, 0xf0, 0xde, 0xe3,
	0xae, 0x54, 0x3e, 0x48, 0x72, 0x6f, 0x3c, 0x1e, 0xc6, 0x25, 0x08, 0x4e, 0x1d, 0xe5, 0x83, 0x38,
	0x04, 0xa1, 0x40, 0x08, 0x57, 0x90, 0x59, 0x7d, 0x05, 0x3f, 0x18, 0xc5, 0xde, 0x72, 0xd6, 0x7d,
	0xf3, 0x77, 0x84, 0xdf, 0x91, 0xd1, 0x59, 0x18, 0xf7, 0x93, 0xac, 0x60, 0x0f, 0xf0, 0xe2, 0x30,
	0xea, 0x49, 0x35, 0x5a, 0xc6, 0x4b, 0x77, 0x66, 0x33, 0x0f, 0x28, 0xe9, 0xe6, 0xb2, 0xe7, 0xce,
	0xbf, 0x0b, 0x0c, 0xfd, 0x94, 0x32, 0x19, 0x4c, 0xb4, 0xa9, 0x4c, 0xbe, 0x4f, 0x55, 0xff, 0x10,
	0xcf, 0xdf, 0x61, 0x0e, 0x80, 0x0a, 0xb9, 0x07, 0x54, 0x92, 0x99, 0xbd, 0x57, 0xe8, 0xea, 0xda,
	0xa8, 0xbc, 0xbe, 0x36, 0x2a, 0x6f, 0xae, 0x0d, 0xf4, 0xdd, 0xc0, 0x40, 0xbf, 0x0e, 0x0c, 0xf4,
	0xe7, 0xc0, 0x40, 0x57, 0x03, 0x03, 0xfd, 0x35, 0x30, 0xd0, 0x3f, 0x03, 0xa3, 0xf2, 0x66, 0x60,
	0xa0, 0x1f, 0x6e, 0x8c, 0xca, 0xd5, 0x8d, 0x51, 0x79, 0x7d, 0x63, 0x54, 0xf0, 0x62, 0x14, 0xd7,
	0x26, 0xfe, 0xf2, 0xf7, 0xde, 0x2d, 0x7a, 0xca, 0x4b, 0x1f, 0x06, 0x1e, 0xfa, 0x6a, 0xf5, 0xc5,
	0x08, 0x1a, 0xc5, 0x77, 0x9e, 0x11, 0x9f, 0x67, 0x1f, 0xbf, 0x69, 0x73, 0xb2, 0x00, 0x68, 0x27,
	0xaa, 0x41, 0xb6, 0xd7, 0xe1, 0xee, 0xbf, 0xda, 0x62, 0x31, 0x5e, 0xaf, 0xd3, 0x4e, 0x54, 0xaf,
	0x67, 0x33, 0xf5, 0xfa, 0xe1, 0xee, 0xd7, 0xb3, 0xd9, 0xbb, 0xe3, 0x93, 0xff, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xb9, 0x99, 0x48, 0x19, 0x9f, 0x08, 0x00, 0x00,
}

func (x WorkflowIdReusePolicy) String() string {
	s, ok := WorkflowIdReusePolicy_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x ParentClosePolicy) String() string {
	s, ok := ParentClosePolicy_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x ContinueAsNewInitiator) String() string {
	s, ok := ContinueAsNewInitiator_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x WorkflowExecutionStatus) String() string {
	s, ok := WorkflowExecutionStatus_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x PendingActivityState) String() string {
	s, ok := PendingActivityState_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x PendingWorkflowTaskState) String() string {
	s, ok := PendingWorkflowTaskState_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x HistoryEventFilterType) String() string {
	s, ok := HistoryEventFilterType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x RetryState) String() string {
	s, ok := RetryState_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x TimeoutType) String() string {
	s, ok := TimeoutType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
