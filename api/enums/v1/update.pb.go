// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/api/enums/v1/update.proto

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

type UpdateWorkflowResultAccessStyle int32

const (
	UPDATE_WORKFLOW_RESULT_ACCESS_STYLE_UNSPECIFIED UpdateWorkflowResultAccessStyle = 0
	// Indicates that the update response _must_ be included as part of the gRPC
	// response body
	UPDATE_WORKFLOW_RESULT_ACCESS_STYLE_REQUIRE_INLINE UpdateWorkflowResultAccessStyle = 1
)

var UpdateWorkflowResultAccessStyle_name = map[int32]string{
	0: "Unspecified",
	1: "RequireInline",
}

var UpdateWorkflowResultAccessStyle_value = map[string]int32{
	"Unspecified":   0,
	"RequireInline": 1,
}

func (UpdateWorkflowResultAccessStyle) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_62be6b87506ae277, []int{0}
}

type UpdateWorkflowRejectionDurabilityPreference int32

const (
	// The workflow expresses no preference as to the durability of the
	// rejection.
	UPDATE_WORKFLOW_REJECTION_DURABILITY_PREFERENCE_UNSPECIFIED UpdateWorkflowRejectionDurabilityPreference = 0
	// Used by a workflow to indicate that no workflow state mutation occurred
	// while processing the update and that it wishes that the update not be
	// made durable (and thus not take up space in workflow history).
	UPDATE_WORKFLOW_REJECTION_DURABILITY_PREFERENCE_BYPASS UpdateWorkflowRejectionDurabilityPreference = 1
)

var UpdateWorkflowRejectionDurabilityPreference_name = map[int32]string{
	0: "Unspecified",
	1: "Bypass",
}

var UpdateWorkflowRejectionDurabilityPreference_value = map[string]int32{
	"Unspecified": 0,
	"Bypass":      1,
}

func (UpdateWorkflowRejectionDurabilityPreference) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_62be6b87506ae277, []int{1}
}

func init() {
	proto.RegisterEnum("temporal.api.enums.v1.UpdateWorkflowResultAccessStyle", UpdateWorkflowResultAccessStyle_name, UpdateWorkflowResultAccessStyle_value)
	proto.RegisterEnum("temporal.api.enums.v1.UpdateWorkflowRejectionDurabilityPreference", UpdateWorkflowRejectionDurabilityPreference_name, UpdateWorkflowRejectionDurabilityPreference_value)
}

func init() {
	proto.RegisterFile("temporal/api/enums/v1/update.proto", fileDescriptor_62be6b87506ae277)
}

var fileDescriptor_62be6b87506ae277 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xbf, 0xae, 0xd3, 0x30,
	0x14, 0xc6, 0xe3, 0x3b, 0x30, 0x98, 0xa5, 0x8a, 0x74, 0xa5, 0x3b, 0x19, 0xc1, 0x78, 0x91, 0x12,
	0x85, 0x4a, 0x1d, 0xd2, 0x01, 0xe5, 0x8f, 0x2b, 0x19, 0xa2, 0x34, 0xd8, 0x49, 0xab, 0xb2, 0x58,
	0x69, 0x70, 0x91, 0x21, 0xad, 0xa3, 0xfc, 0x29, 0xea, 0xc6, 0x13, 0x20, 0x1e, 0x03, 0x75, 0xe5,
	0x25, 0x18, 0x3b, 0x76, 0xa4, 0xe9, 0x82, 0x98, 0xfa, 0x08, 0x88, 0xa0, 0x4a, 0x50, 0x18, 0xb8,
	0x9b, 0xe5, 0xf3, 0x3b, 0xdf, 0xf9, 0x8e, 0xbe, 0x03, 0x1f, 0xd5, 0x62, 0x59, 0xa8, 0x32, 0xcd,
	0xcd, 0xb4, 0x90, 0xa6, 0x58, 0x35, 0xcb, 0xca, 0x5c, 0x5b, 0x66, 0x53, 0xbc, 0x4a, 0x6b, 0x61,
	0x14, 0xa5, 0xaa, 0x95, 0x7e, 0x7d, 0x66, 0x8c, 0xb4, 0x90, 0x46, 0xc7, 0x18, 0x6b, 0xeb, 0xf6,
	0x03, 0x80, 0x0f, 0x92, 0x8e, 0x9b, 0xaa, 0xf2, 0xed, 0x22, 0x57, 0xef, 0xa8, 0xa8, 0x9a, 0xbc,
	0x76, 0xb2, 0x4c, 0x54, 0x15, 0xab, 0x37, 0xb9, 0xd0, 0xfb, 0xd0, 0x4c, 0x22, 0xdf, 0x89, 0x31,
	0x9f, 0x8e, 0xe9, 0xf3, 0x51, 0x30, 0x9e, 0x72, 0x8a, 0x59, 0x12, 0xc4, 0xdc, 0xf1, 0x3c, 0xcc,
	0x18, 0x67, 0xf1, 0x2c, 0xc0, 0x3c, 0x09, 0x59, 0x84, 0x3d, 0x32, 0x22, 0xd8, 0xef, 0x69, 0xfa,
	0x00, 0x3e, 0xf9, 0x9f, 0x26, 0x8a, 0x5f, 0x24, 0x84, 0x62, 0x4e, 0xc2, 0x80, 0x84, 0xb8, 0x07,
	0x6e, 0xb7, 0x00, 0x3e, 0xbe, 0x34, 0xf4, 0x46, 0x64, 0xb5, 0x54, 0x2b, 0xbf, 0x29, 0xd3, 0xb9,
	0xcc, 0x65, 0xbd, 0x89, 0x4a, 0xb1, 0x10, 0xa5, 0x58, 0x65, 0x42, 0x7f, 0x0a, 0x87, 0x7f, 0xcf,
	0x79, 0x86, 0xbd, 0x98, 0x8c, 0x43, 0xee, 0x27, 0xd4, 0x71, 0x49, 0x40, 0xe2, 0x19, 0x8f, 0x28,
	0x1e, 0x61, 0x8a, 0x43, 0xef, 0xd2, 0xa8, 0x0d, 0x07, 0x77, 0x15, 0x70, 0x67, 0x91, 0xc3, 0x58,
	0x0f, 0xb8, 0x9f, 0xc1, 0xee, 0x80, 0xb4, 0xfd, 0x01, 0x69, 0xa7, 0x03, 0x02, 0xef, 0x5b, 0x04,
	0x3e, 0xb5, 0x08, 0x7c, 0x69, 0x11, 0xd8, 0xb5, 0x08, 0x7c, 0x6d, 0x11, 0xf8, 0xd6, 0x22, 0xed,
	0xd4, 0x22, 0xf0, 0xf1, 0x88, 0xb4, 0xdd, 0x11, 0x69, 0xfb, 0x23, 0xd2, 0xe0, 0x8d, 0x54, 0xc6,
	0x3f, 0xe3, 0x70, 0xef, 0xff, 0x5a, 0x3d, 0xfa, 0x19, 0x59, 0x04, 0x5e, 0x3e, 0x7c, 0xfd, 0x1b,
	0x28, 0xd5, 0x1f, 0xf1, 0x0e, 0xbb, 0xc7, 0xf6, 0xea, 0x3a, 0x3e, 0x03, 0x4e, 0x21, 0x0d, 0xdc,
	0x29, 0x4d, 0xac, 0xef, 0x57, 0x37, 0xe7, 0x7f, 0xdb, 0x76, 0x0a, 0x69, 0xdb, 0x5d, 0xc5, 0xb6,
	0x27, 0xd6, 0xfc, 0x5e, 0x77, 0x11, 0xfd, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x35, 0x8b, 0x62,
	0xb2, 0x37, 0x02, 0x00, 0x00,
}

func (x UpdateWorkflowResultAccessStyle) String() string {
	s, ok := UpdateWorkflowResultAccessStyle_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x UpdateWorkflowRejectionDurabilityPreference) String() string {
	s, ok := UpdateWorkflowRejectionDurabilityPreference_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
