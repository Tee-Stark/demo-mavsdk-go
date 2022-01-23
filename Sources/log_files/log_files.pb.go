// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.15.2
// source: log_files.proto

package log_files

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

// Possible results returned for calibration commands
type LogFilesResult_Result int32

const (
	LogFilesResult_RESULT_UNKNOWN          LogFilesResult_Result = 0 // Unknown result
	LogFilesResult_RESULT_SUCCESS          LogFilesResult_Result = 1 // Request succeeded
	LogFilesResult_RESULT_NEXT             LogFilesResult_Result = 2 // Progress update
	LogFilesResult_RESULT_NO_LOGFILES      LogFilesResult_Result = 3 // No log files found
	LogFilesResult_RESULT_TIMEOUT          LogFilesResult_Result = 4 // A timeout happened
	LogFilesResult_RESULT_INVALID_ARGUMENT LogFilesResult_Result = 5 // Invalid argument
	LogFilesResult_RESULT_FILE_OPEN_FAILED LogFilesResult_Result = 6 // File open failed
	LogFilesResult_RESULT_NO_SYSTEM        LogFilesResult_Result = 7 // No system is connected
)

// Enum value maps for LogFilesResult_Result.
var (
	LogFilesResult_Result_name = map[int32]string{
		0: "RESULT_UNKNOWN",
		1: "RESULT_SUCCESS",
		2: "RESULT_NEXT",
		3: "RESULT_NO_LOGFILES",
		4: "RESULT_TIMEOUT",
		5: "RESULT_INVALID_ARGUMENT",
		6: "RESULT_FILE_OPEN_FAILED",
		7: "RESULT_NO_SYSTEM",
	}
	LogFilesResult_Result_value = map[string]int32{
		"RESULT_UNKNOWN":          0,
		"RESULT_SUCCESS":          1,
		"RESULT_NEXT":             2,
		"RESULT_NO_LOGFILES":      3,
		"RESULT_TIMEOUT":          4,
		"RESULT_INVALID_ARGUMENT": 5,
		"RESULT_FILE_OPEN_FAILED": 6,
		"RESULT_NO_SYSTEM":        7,
	}
)

func (x LogFilesResult_Result) Enum() *LogFilesResult_Result {
	p := new(LogFilesResult_Result)
	*p = x
	return p
}

func (x LogFilesResult_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogFilesResult_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_log_files_proto_enumTypes[0].Descriptor()
}

func (LogFilesResult_Result) Type() protoreflect.EnumType {
	return &file_log_files_proto_enumTypes[0]
}

func (x LogFilesResult_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogFilesResult_Result.Descriptor instead.
func (LogFilesResult_Result) EnumDescriptor() ([]byte, []int) {
	return file_log_files_proto_rawDescGZIP(), []int{6, 0}
}

type GetEntriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetEntriesRequest) Reset() {
	*x = GetEntriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_files_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEntriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEntriesRequest) ProtoMessage() {}

func (x *GetEntriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_log_files_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEntriesRequest.ProtoReflect.Descriptor instead.
func (*GetEntriesRequest) Descriptor() ([]byte, []int) {
	return file_log_files_proto_rawDescGZIP(), []int{0}
}

type GetEntriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogFilesResult *LogFilesResult `protobuf:"bytes,1,opt,name=log_files_result,json=logFilesResult,proto3" json:"log_files_result,omitempty"`
	Entries        []*Entry        `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"` // List of entries
}

func (x *GetEntriesResponse) Reset() {
	*x = GetEntriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_files_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEntriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEntriesResponse) ProtoMessage() {}

func (x *GetEntriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_log_files_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEntriesResponse.ProtoReflect.Descriptor instead.
func (*GetEntriesResponse) Descriptor() ([]byte, []int) {
	return file_log_files_proto_rawDescGZIP(), []int{1}
}

func (x *GetEntriesResponse) GetLogFilesResult() *LogFilesResult {
	if x != nil {
		return x.LogFilesResult
	}
	return nil
}

func (x *GetEntriesResponse) GetEntries() []*Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type SubscribeDownloadLogFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entry *Entry `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"` // Entry of the log file to download.
	Path  string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`   // Path of where to download log file to.
}

func (x *SubscribeDownloadLogFileRequest) Reset() {
	*x = SubscribeDownloadLogFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_files_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeDownloadLogFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeDownloadLogFileRequest) ProtoMessage() {}

func (x *SubscribeDownloadLogFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_log_files_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeDownloadLogFileRequest.ProtoReflect.Descriptor instead.
func (*SubscribeDownloadLogFileRequest) Descriptor() ([]byte, []int) {
	return file_log_files_proto_rawDescGZIP(), []int{2}
}

func (x *SubscribeDownloadLogFileRequest) GetEntry() *Entry {
	if x != nil {
		return x.Entry
	}
	return nil
}

func (x *SubscribeDownloadLogFileRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type DownloadLogFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogFilesResult *LogFilesResult `protobuf:"bytes,1,opt,name=log_files_result,json=logFilesResult,proto3" json:"log_files_result,omitempty"`
	Progress       *ProgressData   `protobuf:"bytes,2,opt,name=progress,proto3" json:"progress,omitempty"` // Progress if result is progress
}

func (x *DownloadLogFileResponse) Reset() {
	*x = DownloadLogFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_files_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadLogFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadLogFileResponse) ProtoMessage() {}

func (x *DownloadLogFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_log_files_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadLogFileResponse.ProtoReflect.Descriptor instead.
func (*DownloadLogFileResponse) Descriptor() ([]byte, []int) {
	return file_log_files_proto_rawDescGZIP(), []int{3}
}

func (x *DownloadLogFileResponse) GetLogFilesResult() *LogFilesResult {
	if x != nil {
		return x.LogFilesResult
	}
	return nil
}

func (x *DownloadLogFileResponse) GetProgress() *ProgressData {
	if x != nil {
		return x.Progress
	}
	return nil
}

//
// Progress data coming when downloading a log file.
type ProgressData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Progress float32 `protobuf:"fixed32,1,opt,name=progress,proto3" json:"progress,omitempty"` // Progress from 0 to 1
}

func (x *ProgressData) Reset() {
	*x = ProgressData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_files_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProgressData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgressData) ProtoMessage() {}

func (x *ProgressData) ProtoReflect() protoreflect.Message {
	mi := &file_log_files_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgressData.ProtoReflect.Descriptor instead.
func (*ProgressData) Descriptor() ([]byte, []int) {
	return file_log_files_proto_rawDescGZIP(), []int{4}
}

func (x *ProgressData) GetProgress() float32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

// Log file entry type.
type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                // ID of the log file, to specify a file to be downloaded
	Date      string `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`                             // Date of the log file in UTC in ISO 8601 format "yyyy-mm-ddThh:mm:ssZ"
	SizeBytes uint32 `protobuf:"varint,3,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"` // Size of file in bytes
}

func (x *Entry) Reset() {
	*x = Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_files_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_log_files_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_log_files_proto_rawDescGZIP(), []int{5}
}

func (x *Entry) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Entry) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Entry) GetSizeBytes() uint32 {
	if x != nil {
		return x.SizeBytes
	}
	return 0
}

// Result type.
type LogFilesResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result    LogFilesResult_Result `protobuf:"varint,1,opt,name=result,proto3,enum=mavsdk.rpc.log_files.LogFilesResult_Result" json:"result,omitempty"` // Result enum value
	ResultStr string                `protobuf:"bytes,2,opt,name=result_str,json=resultStr,proto3" json:"result_str,omitempty"`                           // Human-readable English string describing the result
}

func (x *LogFilesResult) Reset() {
	*x = LogFilesResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_files_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogFilesResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogFilesResult) ProtoMessage() {}

func (x *LogFilesResult) ProtoReflect() protoreflect.Message {
	mi := &file_log_files_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogFilesResult.ProtoReflect.Descriptor instead.
func (*LogFilesResult) Descriptor() ([]byte, []int) {
	return file_log_files_proto_rawDescGZIP(), []int{6}
}

func (x *LogFilesResult) GetResult() LogFilesResult_Result {
	if x != nil {
		return x.Result
	}
	return LogFilesResult_RESULT_UNKNOWN
}

func (x *LogFilesResult) GetResultStr() string {
	if x != nil {
		return x.ResultStr
	}
	return ""
}

var File_log_files_proto protoreflect.FileDescriptor

var file_log_files_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6c, 0x6f, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f,
	0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x1a, 0x14, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x5f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x13, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x9b, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x10, 0x6c, 0x6f, 0x67,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x6c, 0x6f, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0e, 0x6c, 0x6f, 0x67, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x35, 0x0a, 0x07, 0x65, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x61, 0x76,
	0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x22, 0x68, 0x0a, 0x1f, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x6c, 0x6f, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0xa9, 0x01, 0x0a, 0x17, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x10, 0x6c, 0x6f, 0x67, 0x5f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f,
	0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0e, 0x6c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3e, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64,
	0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x22, 0x33, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x42, 0x07, 0x82, 0xb5, 0x18, 0x03, 0x4e, 0x61,
	0x4e, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x22, 0x4a, 0x0a, 0x05, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x69, 0x7a, 0x65,
	0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x69,
	0x7a, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0xb4, 0x02, 0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x46,
	0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x43, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x6d, 0x61, 0x76,
	0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x74, 0x72, 0x22, 0xbd,
	0x01, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x53,
	0x55, 0x4c, 0x54, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x12, 0x0a,
	0x0e, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10,
	0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x4e, 0x45, 0x58, 0x54,
	0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x4e, 0x4f, 0x5f,
	0x4c, 0x4f, 0x47, 0x46, 0x49, 0x4c, 0x45, 0x53, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45,
	0x53, 0x55, 0x4c, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x04, 0x12, 0x1b,
	0x0a, 0x17, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x05, 0x12, 0x1b, 0x0a, 0x17, 0x52,
	0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x5f,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x06, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x45, 0x53, 0x55,
	0x4c, 0x54, 0x5f, 0x4e, 0x4f, 0x5f, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x10, 0x07, 0x32, 0x83,
	0x02, 0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x61, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x12, 0x27, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f,
	0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6d, 0x61, 0x76, 0x73,
	0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x8c, 0x01, 0x0a, 0x18, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x6f, 0x67, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x35, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x6c, 0x6f, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x6f, 0x67, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6d, 0x61, 0x76, 0x73,
	0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x08, 0x80, 0xb5, 0x18, 0x00, 0x88, 0xb5,
	0x18, 0x01, 0x30, 0x01, 0x42, 0x1c, 0x42, 0x0d, 0x4c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x0b, 0x2e, 0x3b, 0x6c, 0x6f, 0x67, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_log_files_proto_rawDescOnce sync.Once
	file_log_files_proto_rawDescData = file_log_files_proto_rawDesc
)

func file_log_files_proto_rawDescGZIP() []byte {
	file_log_files_proto_rawDescOnce.Do(func() {
		file_log_files_proto_rawDescData = protoimpl.X.CompressGZIP(file_log_files_proto_rawDescData)
	})
	return file_log_files_proto_rawDescData
}

var file_log_files_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_log_files_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_log_files_proto_goTypes = []interface{}{
	(LogFilesResult_Result)(0),              // 0: mavsdk.rpc.log_files.LogFilesResult.Result
	(*GetEntriesRequest)(nil),               // 1: mavsdk.rpc.log_files.GetEntriesRequest
	(*GetEntriesResponse)(nil),              // 2: mavsdk.rpc.log_files.GetEntriesResponse
	(*SubscribeDownloadLogFileRequest)(nil), // 3: mavsdk.rpc.log_files.SubscribeDownloadLogFileRequest
	(*DownloadLogFileResponse)(nil),         // 4: mavsdk.rpc.log_files.DownloadLogFileResponse
	(*ProgressData)(nil),                    // 5: mavsdk.rpc.log_files.ProgressData
	(*Entry)(nil),                           // 6: mavsdk.rpc.log_files.Entry
	(*LogFilesResult)(nil),                  // 7: mavsdk.rpc.log_files.LogFilesResult
}
var file_log_files_proto_depIdxs = []int32{
	7, // 0: mavsdk.rpc.log_files.GetEntriesResponse.log_files_result:type_name -> mavsdk.rpc.log_files.LogFilesResult
	6, // 1: mavsdk.rpc.log_files.GetEntriesResponse.entries:type_name -> mavsdk.rpc.log_files.Entry
	6, // 2: mavsdk.rpc.log_files.SubscribeDownloadLogFileRequest.entry:type_name -> mavsdk.rpc.log_files.Entry
	7, // 3: mavsdk.rpc.log_files.DownloadLogFileResponse.log_files_result:type_name -> mavsdk.rpc.log_files.LogFilesResult
	5, // 4: mavsdk.rpc.log_files.DownloadLogFileResponse.progress:type_name -> mavsdk.rpc.log_files.ProgressData
	0, // 5: mavsdk.rpc.log_files.LogFilesResult.result:type_name -> mavsdk.rpc.log_files.LogFilesResult.Result
	1, // 6: mavsdk.rpc.log_files.LogFilesService.GetEntries:input_type -> mavsdk.rpc.log_files.GetEntriesRequest
	3, // 7: mavsdk.rpc.log_files.LogFilesService.SubscribeDownloadLogFile:input_type -> mavsdk.rpc.log_files.SubscribeDownloadLogFileRequest
	2, // 8: mavsdk.rpc.log_files.LogFilesService.GetEntries:output_type -> mavsdk.rpc.log_files.GetEntriesResponse
	4, // 9: mavsdk.rpc.log_files.LogFilesService.SubscribeDownloadLogFile:output_type -> mavsdk.rpc.log_files.DownloadLogFileResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_log_files_proto_init() }
func file_log_files_proto_init() {
	if File_log_files_proto != nil {
		return
	}
	file_mavsdk_options_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_log_files_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEntriesRequest); i {
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
		file_log_files_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEntriesResponse); i {
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
		file_log_files_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeDownloadLogFileRequest); i {
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
		file_log_files_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadLogFileResponse); i {
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
		file_log_files_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProgressData); i {
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
		file_log_files_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry); i {
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
		file_log_files_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogFilesResult); i {
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
			RawDescriptor: file_log_files_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_log_files_proto_goTypes,
		DependencyIndexes: file_log_files_proto_depIdxs,
		EnumInfos:         file_log_files_proto_enumTypes,
		MessageInfos:      file_log_files_proto_msgTypes,
	}.Build()
	File_log_files_proto = out.File
	file_log_files_proto_rawDesc = nil
	file_log_files_proto_goTypes = nil
	file_log_files_proto_depIdxs = nil
}
