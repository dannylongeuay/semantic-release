// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: pkg/analyzer/commit_analyzer.proto

package analyzer

import (
	semrel "github.com/go-semantic-release/semantic-release/v2/pkg/semrel"
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

type CommitAnalyzerInit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CommitAnalyzerInit) Reset() {
	*x = CommitAnalyzerInit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_analyzer_commit_analyzer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitAnalyzerInit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitAnalyzerInit) ProtoMessage() {}

func (x *CommitAnalyzerInit) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_analyzer_commit_analyzer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitAnalyzerInit.ProtoReflect.Descriptor instead.
func (*CommitAnalyzerInit) Descriptor() ([]byte, []int) {
	return file_pkg_analyzer_commit_analyzer_proto_rawDescGZIP(), []int{0}
}

type CommitAnalyzerName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CommitAnalyzerName) Reset() {
	*x = CommitAnalyzerName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_analyzer_commit_analyzer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitAnalyzerName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitAnalyzerName) ProtoMessage() {}

func (x *CommitAnalyzerName) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_analyzer_commit_analyzer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitAnalyzerName.ProtoReflect.Descriptor instead.
func (*CommitAnalyzerName) Descriptor() ([]byte, []int) {
	return file_pkg_analyzer_commit_analyzer_proto_rawDescGZIP(), []int{1}
}

type CommitAnalyzerVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CommitAnalyzerVersion) Reset() {
	*x = CommitAnalyzerVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_analyzer_commit_analyzer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitAnalyzerVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitAnalyzerVersion) ProtoMessage() {}

func (x *CommitAnalyzerVersion) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_analyzer_commit_analyzer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitAnalyzerVersion.ProtoReflect.Descriptor instead.
func (*CommitAnalyzerVersion) Descriptor() ([]byte, []int) {
	return file_pkg_analyzer_commit_analyzer_proto_rawDescGZIP(), []int{2}
}

type AnalyzeCommits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AnalyzeCommits) Reset() {
	*x = AnalyzeCommits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_analyzer_commit_analyzer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzeCommits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzeCommits) ProtoMessage() {}

func (x *AnalyzeCommits) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_analyzer_commit_analyzer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzeCommits.ProtoReflect.Descriptor instead.
func (*AnalyzeCommits) Descriptor() ([]byte, []int) {
	return file_pkg_analyzer_commit_analyzer_proto_rawDescGZIP(), []int{3}
}

type CommitAnalyzerInit_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config map[string]string `protobuf:"bytes,1,rep,name=config,proto3" json:"config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CommitAnalyzerInit_Request) Reset() {
	*x = CommitAnalyzerInit_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_analyzer_commit_analyzer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitAnalyzerInit_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitAnalyzerInit_Request) ProtoMessage() {}

func (x *CommitAnalyzerInit_Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_analyzer_commit_analyzer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitAnalyzerInit_Request.ProtoReflect.Descriptor instead.
func (*CommitAnalyzerInit_Request) Descriptor() ([]byte, []int) {
	return file_pkg_analyzer_commit_analyzer_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CommitAnalyzerInit_Request) GetConfig() map[string]string {
	if x != nil {
		return x.Config
	}
	return nil
}

type CommitAnalyzerInit_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CommitAnalyzerInit_Response) Reset() {
	*x = CommitAnalyzerInit_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_analyzer_commit_analyzer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitAnalyzerInit_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitAnalyzerInit_Response) ProtoMessage() {}

func (x *CommitAnalyzerInit_Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_analyzer_commit_analyzer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitAnalyzerInit_Response.ProtoReflect.Descriptor instead.
func (*CommitAnalyzerInit_Response) Descriptor() ([]byte, []int) {
	return file_pkg_analyzer_commit_analyzer_proto_rawDescGZIP(), []int{0, 1}
}

func (x *CommitAnalyzerInit_Response) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type CommitAnalyzerName_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CommitAnalyzerName_Request) Reset() {
	*x = CommitAnalyzerName_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_analyzer_commit_analyzer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitAnalyzerName_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitAnalyzerName_Request) ProtoMessage() {}

func (x *CommitAnalyzerName_Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_analyzer_commit_analyzer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitAnalyzerName_Request.ProtoReflect.Descriptor instead.
func (*CommitAnalyzerName_Request) Descriptor() ([]byte, []int) {
	return file_pkg_analyzer_commit_analyzer_proto_rawDescGZIP(), []int{1, 0}
}

type CommitAnalyzerName_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CommitAnalyzerName_Response) Reset() {
	*x = CommitAnalyzerName_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_analyzer_commit_analyzer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitAnalyzerName_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitAnalyzerName_Response) ProtoMessage() {}

func (x *CommitAnalyzerName_Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_analyzer_commit_analyzer_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitAnalyzerName_Response.ProtoReflect.Descriptor instead.
func (*CommitAnalyzerName_Response) Descriptor() ([]byte, []int) {
	return file_pkg_analyzer_commit_analyzer_proto_rawDescGZIP(), []int{1, 1}
}

func (x *CommitAnalyzerName_Response) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CommitAnalyzerVersion_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CommitAnalyzerVersion_Request) Reset() {
	*x = CommitAnalyzerVersion_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_analyzer_commit_analyzer_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitAnalyzerVersion_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitAnalyzerVersion_Request) ProtoMessage() {}

func (x *CommitAnalyzerVersion_Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_analyzer_commit_analyzer_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitAnalyzerVersion_Request.ProtoReflect.Descriptor instead.
func (*CommitAnalyzerVersion_Request) Descriptor() ([]byte, []int) {
	return file_pkg_analyzer_commit_analyzer_proto_rawDescGZIP(), []int{2, 0}
}

type CommitAnalyzerVersion_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *CommitAnalyzerVersion_Response) Reset() {
	*x = CommitAnalyzerVersion_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_analyzer_commit_analyzer_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitAnalyzerVersion_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitAnalyzerVersion_Response) ProtoMessage() {}

func (x *CommitAnalyzerVersion_Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_analyzer_commit_analyzer_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitAnalyzerVersion_Response.ProtoReflect.Descriptor instead.
func (*CommitAnalyzerVersion_Response) Descriptor() ([]byte, []int) {
	return file_pkg_analyzer_commit_analyzer_proto_rawDescGZIP(), []int{2, 1}
}

func (x *CommitAnalyzerVersion_Response) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type AnalyzeCommits_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawCommits []*semrel.RawCommit `protobuf:"bytes,1,rep,name=raw_commits,json=rawCommits,proto3" json:"raw_commits,omitempty"`
}

func (x *AnalyzeCommits_Request) Reset() {
	*x = AnalyzeCommits_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_analyzer_commit_analyzer_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzeCommits_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzeCommits_Request) ProtoMessage() {}

func (x *AnalyzeCommits_Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_analyzer_commit_analyzer_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzeCommits_Request.ProtoReflect.Descriptor instead.
func (*AnalyzeCommits_Request) Descriptor() ([]byte, []int) {
	return file_pkg_analyzer_commit_analyzer_proto_rawDescGZIP(), []int{3, 0}
}

func (x *AnalyzeCommits_Request) GetRawCommits() []*semrel.RawCommit {
	if x != nil {
		return x.RawCommits
	}
	return nil
}

type AnalyzeCommits_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commits []*semrel.Commit `protobuf:"bytes,1,rep,name=commits,proto3" json:"commits,omitempty"`
}

func (x *AnalyzeCommits_Response) Reset() {
	*x = AnalyzeCommits_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_analyzer_commit_analyzer_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzeCommits_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzeCommits_Response) ProtoMessage() {}

func (x *AnalyzeCommits_Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_analyzer_commit_analyzer_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzeCommits_Response.ProtoReflect.Descriptor instead.
func (*AnalyzeCommits_Response) Descriptor() ([]byte, []int) {
	return file_pkg_analyzer_commit_analyzer_proto_rawDescGZIP(), []int{3, 1}
}

func (x *AnalyzeCommits_Response) GetCommits() []*semrel.Commit {
	if x != nil {
		return x.Commits
	}
	return nil
}

var File_pkg_analyzer_commit_analyzer_proto protoreflect.FileDescriptor

var file_pkg_analyzer_commit_analyzer_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x6d, 0x72, 0x65, 0x6c,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe,
	0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65,
	0x72, 0x49, 0x6e, 0x69, 0x74, 0x1a, 0x85, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3f, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a,
	0x65, 0x72, 0x49, 0x6e, 0x69, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x1a, 0x39, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x20, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22,
	0x3f, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x48, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a,
	0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x77, 0x0a, 0x0e, 0x41, 0x6e,
	0x61, 0x6c, 0x79, 0x7a, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x1a, 0x36, 0x0a, 0x07,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x0b, 0x72, 0x61, 0x77, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x52,
	0x61, 0x77, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x0a, 0x72, 0x61, 0x77, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x73, 0x1a, 0x2d, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x21, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x73, 0x32, 0xa6, 0x02, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x41, 0x6e,
	0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x41, 0x0a, 0x04,
	0x49, 0x6e, 0x69, 0x74, 0x12, 0x1b, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x41, 0x6e, 0x61,
	0x6c, 0x79, 0x7a, 0x65, 0x72, 0x49, 0x6e, 0x69, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a,
	0x65, 0x72, 0x49, 0x6e, 0x69, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x41, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x41, 0x6e, 0x61,
	0x6c, 0x79, 0x7a, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4a, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c,
	0x0a, 0x07, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x12, 0x17, 0x2e, 0x41, 0x6e, 0x61, 0x6c,
	0x79, 0x7a, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x41, 0x5a, 0x3f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x65,
	0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x2d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2f, 0x73,
	0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x2d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2f,
	0x76, 0x32, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_analyzer_commit_analyzer_proto_rawDescOnce sync.Once
	file_pkg_analyzer_commit_analyzer_proto_rawDescData = file_pkg_analyzer_commit_analyzer_proto_rawDesc
)

func file_pkg_analyzer_commit_analyzer_proto_rawDescGZIP() []byte {
	file_pkg_analyzer_commit_analyzer_proto_rawDescOnce.Do(func() {
		file_pkg_analyzer_commit_analyzer_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_analyzer_commit_analyzer_proto_rawDescData)
	})
	return file_pkg_analyzer_commit_analyzer_proto_rawDescData
}

var file_pkg_analyzer_commit_analyzer_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_pkg_analyzer_commit_analyzer_proto_goTypes = []interface{}{
	(*CommitAnalyzerInit)(nil),             // 0: CommitAnalyzerInit
	(*CommitAnalyzerName)(nil),             // 1: CommitAnalyzerName
	(*CommitAnalyzerVersion)(nil),          // 2: CommitAnalyzerVersion
	(*AnalyzeCommits)(nil),                 // 3: AnalyzeCommits
	(*CommitAnalyzerInit_Request)(nil),     // 4: CommitAnalyzerInit.Request
	(*CommitAnalyzerInit_Response)(nil),    // 5: CommitAnalyzerInit.Response
	nil,                                    // 6: CommitAnalyzerInit.Request.ConfigEntry
	(*CommitAnalyzerName_Request)(nil),     // 7: CommitAnalyzerName.Request
	(*CommitAnalyzerName_Response)(nil),    // 8: CommitAnalyzerName.Response
	(*CommitAnalyzerVersion_Request)(nil),  // 9: CommitAnalyzerVersion.Request
	(*CommitAnalyzerVersion_Response)(nil), // 10: CommitAnalyzerVersion.Response
	(*AnalyzeCommits_Request)(nil),         // 11: AnalyzeCommits.Request
	(*AnalyzeCommits_Response)(nil),        // 12: AnalyzeCommits.Response
	(*semrel.RawCommit)(nil),               // 13: RawCommit
	(*semrel.Commit)(nil),                  // 14: Commit
}
var file_pkg_analyzer_commit_analyzer_proto_depIdxs = []int32{
	6,  // 0: CommitAnalyzerInit.Request.config:type_name -> CommitAnalyzerInit.Request.ConfigEntry
	13, // 1: AnalyzeCommits.Request.raw_commits:type_name -> RawCommit
	14, // 2: AnalyzeCommits.Response.commits:type_name -> Commit
	4,  // 3: CommitAnalyzerPlugin.Init:input_type -> CommitAnalyzerInit.Request
	7,  // 4: CommitAnalyzerPlugin.Name:input_type -> CommitAnalyzerName.Request
	9,  // 5: CommitAnalyzerPlugin.Version:input_type -> CommitAnalyzerVersion.Request
	11, // 6: CommitAnalyzerPlugin.Analyze:input_type -> AnalyzeCommits.Request
	5,  // 7: CommitAnalyzerPlugin.Init:output_type -> CommitAnalyzerInit.Response
	8,  // 8: CommitAnalyzerPlugin.Name:output_type -> CommitAnalyzerName.Response
	10, // 9: CommitAnalyzerPlugin.Version:output_type -> CommitAnalyzerVersion.Response
	12, // 10: CommitAnalyzerPlugin.Analyze:output_type -> AnalyzeCommits.Response
	7,  // [7:11] is the sub-list for method output_type
	3,  // [3:7] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_analyzer_commit_analyzer_proto_init() }
func file_pkg_analyzer_commit_analyzer_proto_init() {
	if File_pkg_analyzer_commit_analyzer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_analyzer_commit_analyzer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitAnalyzerInit); i {
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
		file_pkg_analyzer_commit_analyzer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitAnalyzerName); i {
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
		file_pkg_analyzer_commit_analyzer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitAnalyzerVersion); i {
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
		file_pkg_analyzer_commit_analyzer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzeCommits); i {
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
		file_pkg_analyzer_commit_analyzer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitAnalyzerInit_Request); i {
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
		file_pkg_analyzer_commit_analyzer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitAnalyzerInit_Response); i {
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
		file_pkg_analyzer_commit_analyzer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitAnalyzerName_Request); i {
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
		file_pkg_analyzer_commit_analyzer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitAnalyzerName_Response); i {
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
		file_pkg_analyzer_commit_analyzer_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitAnalyzerVersion_Request); i {
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
		file_pkg_analyzer_commit_analyzer_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitAnalyzerVersion_Response); i {
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
		file_pkg_analyzer_commit_analyzer_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzeCommits_Request); i {
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
		file_pkg_analyzer_commit_analyzer_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzeCommits_Response); i {
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
			RawDescriptor: file_pkg_analyzer_commit_analyzer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_analyzer_commit_analyzer_proto_goTypes,
		DependencyIndexes: file_pkg_analyzer_commit_analyzer_proto_depIdxs,
		MessageInfos:      file_pkg_analyzer_commit_analyzer_proto_msgTypes,
	}.Build()
	File_pkg_analyzer_commit_analyzer_proto = out.File
	file_pkg_analyzer_commit_analyzer_proto_rawDesc = nil
	file_pkg_analyzer_commit_analyzer_proto_goTypes = nil
	file_pkg_analyzer_commit_analyzer_proto_depIdxs = nil
}
