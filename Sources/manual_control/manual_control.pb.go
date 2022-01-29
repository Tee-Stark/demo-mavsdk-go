// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.2
// source: manual_control.proto

package manual_control

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Possible results returned for manual control requests.
type ManualControlResult_Result int32

const (
	ManualControlResult_RESULT_UNKNOWN            ManualControlResult_Result = 0 // Unknown result
	ManualControlResult_RESULT_SUCCESS            ManualControlResult_Result = 1 // Request was successful
	ManualControlResult_RESULT_NO_SYSTEM          ManualControlResult_Result = 2 // No system is connected
	ManualControlResult_RESULT_CONNECTION_ERROR   ManualControlResult_Result = 3 // Connection error
	ManualControlResult_RESULT_BUSY               ManualControlResult_Result = 4 // Vehicle is busy
	ManualControlResult_RESULT_COMMAND_DENIED     ManualControlResult_Result = 5 // Command refused by vehicle
	ManualControlResult_RESULT_TIMEOUT            ManualControlResult_Result = 6 // Request timed out
	ManualControlResult_RESULT_INPUT_OUT_OF_RANGE ManualControlResult_Result = 7 // Input out of range
	ManualControlResult_RESULT_INPUT_NOT_SET      ManualControlResult_Result = 8 // No Input set
)

// Enum value maps for ManualControlResult_Result.
var (
	ManualControlResult_Result_name = map[int32]string{
		0: "RESULT_UNKNOWN",
		1: "RESULT_SUCCESS",
		2: "RESULT_NO_SYSTEM",
		3: "RESULT_CONNECTION_ERROR",
		4: "RESULT_BUSY",
		5: "RESULT_COMMAND_DENIED",
		6: "RESULT_TIMEOUT",
		7: "RESULT_INPUT_OUT_OF_RANGE",
		8: "RESULT_INPUT_NOT_SET",
	}
	ManualControlResult_Result_value = map[string]int32{
		"RESULT_UNKNOWN":            0,
		"RESULT_SUCCESS":            1,
		"RESULT_NO_SYSTEM":          2,
		"RESULT_CONNECTION_ERROR":   3,
		"RESULT_BUSY":               4,
		"RESULT_COMMAND_DENIED":     5,
		"RESULT_TIMEOUT":            6,
		"RESULT_INPUT_OUT_OF_RANGE": 7,
		"RESULT_INPUT_NOT_SET":      8,
	}
)

func (x ManualControlResult_Result) Enum() *ManualControlResult_Result {
	p := new(ManualControlResult_Result)
	*p = x
	return p
}

func (x ManualControlResult_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ManualControlResult_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_manual_control_proto_enumTypes[0].Descriptor()
}

func (ManualControlResult_Result) Type() protoreflect.EnumType {
	return &file_manual_control_proto_enumTypes[0]
}

func (x ManualControlResult_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ManualControlResult_Result.Descriptor instead.
func (ManualControlResult_Result) EnumDescriptor() ([]byte, []int) {
	return file_manual_control_proto_rawDescGZIP(), []int{6, 0}
}

type StartPositionControlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartPositionControlRequest) Reset() {
	*x = StartPositionControlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manual_control_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartPositionControlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartPositionControlRequest) ProtoMessage() {}

func (x *StartPositionControlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manual_control_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartPositionControlRequest.ProtoReflect.Descriptor instead.
func (*StartPositionControlRequest) Descriptor() ([]byte, []int) {
	return file_manual_control_proto_rawDescGZIP(), []int{0}
}

type StartPositionControlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ManualControlResult *ManualControlResult `protobuf:"bytes,1,opt,name=manual_control_result,json=manualControlResult,proto3" json:"manual_control_result,omitempty"`
}

func (x *StartPositionControlResponse) Reset() {
	*x = StartPositionControlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manual_control_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartPositionControlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartPositionControlResponse) ProtoMessage() {}

func (x *StartPositionControlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_manual_control_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartPositionControlResponse.ProtoReflect.Descriptor instead.
func (*StartPositionControlResponse) Descriptor() ([]byte, []int) {
	return file_manual_control_proto_rawDescGZIP(), []int{1}
}

func (x *StartPositionControlResponse) GetManualControlResult() *ManualControlResult {
	if x != nil {
		return x.ManualControlResult
	}
	return nil
}

type StartAltitudeControlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartAltitudeControlRequest) Reset() {
	*x = StartAltitudeControlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manual_control_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartAltitudeControlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartAltitudeControlRequest) ProtoMessage() {}

func (x *StartAltitudeControlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manual_control_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartAltitudeControlRequest.ProtoReflect.Descriptor instead.
func (*StartAltitudeControlRequest) Descriptor() ([]byte, []int) {
	return file_manual_control_proto_rawDescGZIP(), []int{2}
}

type StartAltitudeControlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ManualControlResult *ManualControlResult `protobuf:"bytes,1,opt,name=manual_control_result,json=manualControlResult,proto3" json:"manual_control_result,omitempty"`
}

func (x *StartAltitudeControlResponse) Reset() {
	*x = StartAltitudeControlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manual_control_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartAltitudeControlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartAltitudeControlResponse) ProtoMessage() {}

func (x *StartAltitudeControlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_manual_control_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartAltitudeControlResponse.ProtoReflect.Descriptor instead.
func (*StartAltitudeControlResponse) Descriptor() ([]byte, []int) {
	return file_manual_control_proto_rawDescGZIP(), []int{3}
}

func (x *StartAltitudeControlResponse) GetManualControlResult() *ManualControlResult {
	if x != nil {
		return x.ManualControlResult
	}
	return nil
}

type SetManualControlInputRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"` // value between -1. to 1. negative -> backwards, positive -> forwards
	Y float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"` // value between -1. to 1. negative -> left, positive -> right
	Z float32 `protobuf:"fixed32,3,opt,name=z,proto3" json:"z,omitempty"` // value between -1. to 1. negative -> down, positive -> up (usually for now, for multicopter 0 to 1 is expected)
	R float32 `protobuf:"fixed32,4,opt,name=r,proto3" json:"r,omitempty"` // value between -1. to 1. negative -> turn anti-clockwise (towards the left), positive -> turn clockwise (towards the right)
}

func (x *SetManualControlInputRequest) Reset() {
	*x = SetManualControlInputRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manual_control_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetManualControlInputRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetManualControlInputRequest) ProtoMessage() {}

func (x *SetManualControlInputRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manual_control_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetManualControlInputRequest.ProtoReflect.Descriptor instead.
func (*SetManualControlInputRequest) Descriptor() ([]byte, []int) {
	return file_manual_control_proto_rawDescGZIP(), []int{4}
}

func (x *SetManualControlInputRequest) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *SetManualControlInputRequest) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *SetManualControlInputRequest) GetZ() float32 {
	if x != nil {
		return x.Z
	}
	return 0
}

func (x *SetManualControlInputRequest) GetR() float32 {
	if x != nil {
		return x.R
	}
	return 0
}

type SetManualControlInputResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ManualControlResult *ManualControlResult `protobuf:"bytes,1,opt,name=manual_control_result,json=manualControlResult,proto3" json:"manual_control_result,omitempty"`
}

func (x *SetManualControlInputResponse) Reset() {
	*x = SetManualControlInputResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manual_control_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetManualControlInputResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetManualControlInputResponse) ProtoMessage() {}

func (x *SetManualControlInputResponse) ProtoReflect() protoreflect.Message {
	mi := &file_manual_control_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetManualControlInputResponse.ProtoReflect.Descriptor instead.
func (*SetManualControlInputResponse) Descriptor() ([]byte, []int) {
	return file_manual_control_proto_rawDescGZIP(), []int{5}
}

func (x *SetManualControlInputResponse) GetManualControlResult() *ManualControlResult {
	if x != nil {
		return x.ManualControlResult
	}
	return nil
}

// Result type.
type ManualControlResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result    ManualControlResult_Result `protobuf:"varint,1,opt,name=result,proto3,enum=mavsdk.rpc.manual_control.ManualControlResult_Result" json:"result,omitempty"` // Result enum value
	ResultStr string                     `protobuf:"bytes,2,opt,name=result_str,json=resultStr,proto3" json:"result_str,omitempty"`                                     // Human-readable English string describing the result
}

func (x *ManualControlResult) Reset() {
	*x = ManualControlResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manual_control_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManualControlResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManualControlResult) ProtoMessage() {}

func (x *ManualControlResult) ProtoReflect() protoreflect.Message {
	mi := &file_manual_control_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManualControlResult.ProtoReflect.Descriptor instead.
func (*ManualControlResult) Descriptor() ([]byte, []int) {
	return file_manual_control_proto_rawDescGZIP(), []int{6}
}

func (x *ManualControlResult) GetResult() ManualControlResult_Result {
	if x != nil {
		return x.Result
	}
	return ManualControlResult_RESULT_UNKNOWN
}

func (x *ManualControlResult) GetResultStr() string {
	if x != nil {
		return x.ResultStr
	}
	return ""
}

var File_manual_control_proto protoreflect.FileDescriptor

var file_manual_control_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x1a, 0x14, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x1b, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x82, 0x01, 0x0a, 0x1c, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x15, 0x6d, 0x61, 0x6e, 0x75, 0x61,
	0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2e, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x13, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x1d, 0x0a, 0x1b, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x41, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x82, 0x01, 0x0a, 0x1c, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x41, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x15, 0x6d,
	0x61, 0x6e, 0x75, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6d, 0x61, 0x76,
	0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x13, 0x6d, 0x61, 0x6e, 0x75,
	0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x56, 0x0a, 0x1c, 0x53, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a,
	0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x7a, 0x12, 0x0c, 0x0a, 0x01, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x72, 0x22, 0x83, 0x01, 0x0a, 0x1d, 0x53, 0x65, 0x74, 0x4d,
	0x61, 0x6e, 0x75, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x15, 0x6d, 0x61, 0x6e,
	0x75, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64,
	0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x13, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xe2, 0x02,
	0x0a, 0x13, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x4d, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2e, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x73,
	0x74, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x53, 0x74, 0x72, 0x22, 0xdc, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12,
	0x0a, 0x0e, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x53, 0x55, 0x43,
	0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54,
	0x5f, 0x4e, 0x4f, 0x5f, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17,
	0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x45, 0x53,
	0x55, 0x4c, 0x54, 0x5f, 0x42, 0x55, 0x53, 0x59, 0x10, 0x04, 0x12, 0x19, 0x0a, 0x15, 0x52, 0x45,
	0x53, 0x55, 0x4c, 0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x44, 0x45, 0x4e,
	0x49, 0x45, 0x44, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f,
	0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x06, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x45, 0x53,
	0x55, 0x4c, 0x54, 0x5f, 0x49, 0x4e, 0x50, 0x55, 0x54, 0x5f, 0x4f, 0x55, 0x54, 0x5f, 0x4f, 0x46,
	0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x07, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x45, 0x53, 0x55,
	0x4c, 0x54, 0x5f, 0x49, 0x4e, 0x50, 0x55, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x54,
	0x10, 0x08, 0x32, 0xc1, 0x03, 0x0a, 0x14, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x89, 0x01, 0x0a, 0x14,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x12, 0x36, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x6d,
	0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x89, 0x01, 0x0a, 0x14, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x41, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x12, 0x36, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61,
	0x6e, 0x75, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x41, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64,
	0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x6c, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x90, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x75, 0x61,
	0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x37, 0x2e,
	0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x75, 0x61,
	0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x53, 0x65, 0x74, 0x4d, 0x61, 0x6e,
	0x75, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2e, 0x53, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x04, 0x80, 0xb5, 0x18, 0x01, 0x42, 0x26, 0x42, 0x12, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x10, 0x2e, 0x3b,
	0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_manual_control_proto_rawDescOnce sync.Once
	file_manual_control_proto_rawDescData = file_manual_control_proto_rawDesc
)

func file_manual_control_proto_rawDescGZIP() []byte {
	file_manual_control_proto_rawDescOnce.Do(func() {
		file_manual_control_proto_rawDescData = protoimpl.X.CompressGZIP(file_manual_control_proto_rawDescData)
	})
	return file_manual_control_proto_rawDescData
}

var file_manual_control_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_manual_control_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_manual_control_proto_goTypes = []interface{}{
	(ManualControlResult_Result)(0),       // 0: mavsdk.rpc.manual_control.ManualControlResult.Result
	(*StartPositionControlRequest)(nil),   // 1: mavsdk.rpc.manual_control.StartPositionControlRequest
	(*StartPositionControlResponse)(nil),  // 2: mavsdk.rpc.manual_control.StartPositionControlResponse
	(*StartAltitudeControlRequest)(nil),   // 3: mavsdk.rpc.manual_control.StartAltitudeControlRequest
	(*StartAltitudeControlResponse)(nil),  // 4: mavsdk.rpc.manual_control.StartAltitudeControlResponse
	(*SetManualControlInputRequest)(nil),  // 5: mavsdk.rpc.manual_control.SetManualControlInputRequest
	(*SetManualControlInputResponse)(nil), // 6: mavsdk.rpc.manual_control.SetManualControlInputResponse
	(*ManualControlResult)(nil),           // 7: mavsdk.rpc.manual_control.ManualControlResult
}
var file_manual_control_proto_depIdxs = []int32{
	7, // 0: mavsdk.rpc.manual_control.StartPositionControlResponse.manual_control_result:type_name -> mavsdk.rpc.manual_control.ManualControlResult
	7, // 1: mavsdk.rpc.manual_control.StartAltitudeControlResponse.manual_control_result:type_name -> mavsdk.rpc.manual_control.ManualControlResult
	7, // 2: mavsdk.rpc.manual_control.SetManualControlInputResponse.manual_control_result:type_name -> mavsdk.rpc.manual_control.ManualControlResult
	0, // 3: mavsdk.rpc.manual_control.ManualControlResult.result:type_name -> mavsdk.rpc.manual_control.ManualControlResult.Result
	1, // 4: mavsdk.rpc.manual_control.ManualControlService.StartPositionControl:input_type -> mavsdk.rpc.manual_control.StartPositionControlRequest
	3, // 5: mavsdk.rpc.manual_control.ManualControlService.StartAltitudeControl:input_type -> mavsdk.rpc.manual_control.StartAltitudeControlRequest
	5, // 6: mavsdk.rpc.manual_control.ManualControlService.SetManualControlInput:input_type -> mavsdk.rpc.manual_control.SetManualControlInputRequest
	2, // 7: mavsdk.rpc.manual_control.ManualControlService.StartPositionControl:output_type -> mavsdk.rpc.manual_control.StartPositionControlResponse
	4, // 8: mavsdk.rpc.manual_control.ManualControlService.StartAltitudeControl:output_type -> mavsdk.rpc.manual_control.StartAltitudeControlResponse
	6, // 9: mavsdk.rpc.manual_control.ManualControlService.SetManualControlInput:output_type -> mavsdk.rpc.manual_control.SetManualControlInputResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_manual_control_proto_init() }
func file_manual_control_proto_init() {
	if File_manual_control_proto != nil {
		return
	}
	
	if !protoimpl.UnsafeEnabled {
		file_manual_control_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartPositionControlRequest); i {
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
		file_manual_control_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartPositionControlResponse); i {
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
		file_manual_control_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartAltitudeControlRequest); i {
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
		file_manual_control_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartAltitudeControlResponse); i {
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
		file_manual_control_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetManualControlInputRequest); i {
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
		file_manual_control_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetManualControlInputResponse); i {
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
		file_manual_control_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManualControlResult); i {
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
			RawDescriptor: file_manual_control_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_manual_control_proto_goTypes,
		DependencyIndexes: file_manual_control_proto_depIdxs,
		EnumInfos:         file_manual_control_proto_enumTypes,
		MessageInfos:      file_manual_control_proto_msgTypes,
	}.Build()
	File_manual_control_proto = out.File
	file_manual_control_proto_rawDesc = nil
	file_manual_control_proto_goTypes = nil
	file_manual_control_proto_depIdxs = nil
}
