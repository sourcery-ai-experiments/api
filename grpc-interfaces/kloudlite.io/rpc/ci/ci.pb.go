// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: ci.proto

package ci

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

type PipelineIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId               string            `protobuf:"bytes,11,opt,name=userId,proto3" json:"userId,omitempty"`
	ProjectId            string            `protobuf:"bytes,12,opt,name=projectId,proto3" json:"projectId,omitempty"`
	RepoName             string            `protobuf:"bytes,13,opt,name=RepoName,proto3" json:"RepoName,omitempty"`
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ImageName            string            `protobuf:"bytes,2,opt,name=imageName,proto3" json:"imageName,omitempty"`
	PipelineEnv          string            `protobuf:"bytes,3,opt,name=pipelineEnv,proto3" json:"pipelineEnv,omitempty"`
	GitProvider          string            `protobuf:"bytes,4,opt,name=GitProvider,proto3" json:"GitProvider,omitempty"`
	GitRepoUrl           string            `protobuf:"bytes,5,opt,name=GitRepoUrl,proto3" json:"GitRepoUrl,omitempty"`
	DockerFile           string            `protobuf:"bytes,6,opt,name=DockerFile,proto3" json:"DockerFile,omitempty"`
	ContextDir           string            `protobuf:"bytes,7,opt,name=ContextDir,proto3" json:"ContextDir,omitempty"`
	GithubInstallationId int32             `protobuf:"varint,8,opt,name=GithubInstallationId,proto3" json:"GithubInstallationId,omitempty"`
	GitlabTokenId        string            `protobuf:"bytes,9,opt,name=GitlabTokenId,proto3" json:"GitlabTokenId,omitempty"`
	BuildArgs            map[string]string `protobuf:"bytes,10,rep,name=BuildArgs,proto3" json:"BuildArgs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PipelineIn) Reset() {
	*x = PipelineIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ci_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineIn) ProtoMessage() {}

func (x *PipelineIn) ProtoReflect() protoreflect.Message {
	mi := &file_ci_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineIn.ProtoReflect.Descriptor instead.
func (*PipelineIn) Descriptor() ([]byte, []int) {
	return file_ci_proto_rawDescGZIP(), []int{0}
}

func (x *PipelineIn) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PipelineIn) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *PipelineIn) GetRepoName() string {
	if x != nil {
		return x.RepoName
	}
	return ""
}

func (x *PipelineIn) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PipelineIn) GetImageName() string {
	if x != nil {
		return x.ImageName
	}
	return ""
}

func (x *PipelineIn) GetPipelineEnv() string {
	if x != nil {
		return x.PipelineEnv
	}
	return ""
}

func (x *PipelineIn) GetGitProvider() string {
	if x != nil {
		return x.GitProvider
	}
	return ""
}

func (x *PipelineIn) GetGitRepoUrl() string {
	if x != nil {
		return x.GitRepoUrl
	}
	return ""
}

func (x *PipelineIn) GetDockerFile() string {
	if x != nil {
		return x.DockerFile
	}
	return ""
}

func (x *PipelineIn) GetContextDir() string {
	if x != nil {
		return x.ContextDir
	}
	return ""
}

func (x *PipelineIn) GetGithubInstallationId() int32 {
	if x != nil {
		return x.GithubInstallationId
	}
	return 0
}

func (x *PipelineIn) GetGitlabTokenId() string {
	if x != nil {
		return x.GitlabTokenId
	}
	return ""
}

func (x *PipelineIn) GetBuildArgs() map[string]string {
	if x != nil {
		return x.BuildArgs
	}
	return nil
}

type PipelineOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PipelineId string `protobuf:"bytes,1,opt,name=pipelineId,proto3" json:"pipelineId,omitempty"`
}

func (x *PipelineOutput) Reset() {
	*x = PipelineOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ci_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineOutput) ProtoMessage() {}

func (x *PipelineOutput) ProtoReflect() protoreflect.Message {
	mi := &file_ci_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineOutput.ProtoReflect.Descriptor instead.
func (*PipelineOutput) Descriptor() ([]byte, []int) {
	return file_ci_proto_rawDescGZIP(), []int{1}
}

func (x *PipelineOutput) GetPipelineId() string {
	if x != nil {
		return x.PipelineId
	}
	return ""
}

type HarborProjectIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HarborProjectIn) Reset() {
	*x = HarborProjectIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ci_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HarborProjectIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HarborProjectIn) ProtoMessage() {}

func (x *HarborProjectIn) ProtoReflect() protoreflect.Message {
	mi := &file_ci_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HarborProjectIn.ProtoReflect.Descriptor instead.
func (*HarborProjectIn) Descriptor() ([]byte, []int) {
	return file_ci_proto_rawDescGZIP(), []int{2}
}

func (x *HarborProjectIn) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HarborProjectOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *HarborProjectOut) Reset() {
	*x = HarborProjectOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ci_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HarborProjectOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HarborProjectOut) ProtoMessage() {}

func (x *HarborProjectOut) ProtoReflect() protoreflect.Message {
	mi := &file_ci_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HarborProjectOut.ProtoReflect.Descriptor instead.
func (*HarborProjectOut) Descriptor() ([]byte, []int) {
	return file_ci_proto_rawDescGZIP(), []int{3}
}

func (x *HarborProjectOut) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

var File_ci_proto protoreflect.FileDescriptor

var file_ci_proto_rawDesc = []byte{
	0x0a, 0x08, 0x63, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x04, 0x0a, 0x0a, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x52, 0x65, 0x70, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x52, 0x65, 0x70, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x45, 0x6e, 0x76, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x45, 0x6e, 0x76, 0x12,
	0x20, 0x0a, 0x0b, 0x47, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x47, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x47, 0x69, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x55, 0x72, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x47, 0x69, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x55, 0x72,
	0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x44, 0x69, 0x72, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x44, 0x69,
	0x72, 0x12, 0x32, 0x0a, 0x14, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x14, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x47, 0x69,
	0x74, 0x6c, 0x61, 0x62, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x41, 0x72, 0x67, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x2e, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x41, 0x72, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x3c, 0x0a, 0x0e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x41, 0x72,
	0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x30, 0x0a, 0x0e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x49, 0x64, 0x22, 0x25, 0x0a, 0x0f, 0x48, 0x61, 0x72, 0x62, 0x6f, 0x72, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x10,
	0x48, 0x61, 0x72, 0x62, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4f, 0x75, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xac, 0x01, 0x0a, 0x02, 0x43, 0x49, 0x12,
	0x2e, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x12, 0x0b, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x1a, 0x0f,
	0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12,
	0x3a, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x61, 0x72, 0x62, 0x6f, 0x72, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x10, 0x2e, 0x48, 0x61, 0x72, 0x62, 0x6f, 0x72, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x1a, 0x11, 0x2e, 0x48, 0x61, 0x72, 0x62, 0x6f,
	0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4f, 0x75, 0x74, 0x12, 0x3a, 0x0a, 0x13, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x61, 0x72, 0x62, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x10, 0x2e, 0x48, 0x61, 0x72, 0x62, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x6e, 0x1a, 0x11, 0x2e, 0x48, 0x61, 0x72, 0x62, 0x6f, 0x72, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x4f, 0x75, 0x74, 0x42, 0x15, 0x5a, 0x13, 0x6b, 0x6c, 0x6f, 0x75, 0x64,
	0x6c, 0x69, 0x74, 0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ci_proto_rawDescOnce sync.Once
	file_ci_proto_rawDescData = file_ci_proto_rawDesc
)

func file_ci_proto_rawDescGZIP() []byte {
	file_ci_proto_rawDescOnce.Do(func() {
		file_ci_proto_rawDescData = protoimpl.X.CompressGZIP(file_ci_proto_rawDescData)
	})
	return file_ci_proto_rawDescData
}

var file_ci_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_ci_proto_goTypes = []interface{}{
	(*PipelineIn)(nil),       // 0: PipelineIn
	(*PipelineOutput)(nil),   // 1: PipelineOutput
	(*HarborProjectIn)(nil),  // 2: HarborProjectIn
	(*HarborProjectOut)(nil), // 3: HarborProjectOut
	nil,                      // 4: PipelineIn.BuildArgsEntry
}
var file_ci_proto_depIdxs = []int32{
	4, // 0: PipelineIn.BuildArgs:type_name -> PipelineIn.BuildArgsEntry
	0, // 1: CI.CreatePipeline:input_type -> PipelineIn
	2, // 2: CI.CreateHarborProject:input_type -> HarborProjectIn
	2, // 3: CI.DeleteHarborProject:input_type -> HarborProjectIn
	1, // 4: CI.CreatePipeline:output_type -> PipelineOutput
	3, // 5: CI.CreateHarborProject:output_type -> HarborProjectOut
	3, // 6: CI.DeleteHarborProject:output_type -> HarborProjectOut
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ci_proto_init() }
func file_ci_proto_init() {
	if File_ci_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ci_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineIn); i {
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
		file_ci_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineOutput); i {
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
		file_ci_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HarborProjectIn); i {
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
		file_ci_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HarborProjectOut); i {
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
			RawDescriptor: file_ci_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ci_proto_goTypes,
		DependencyIndexes: file_ci_proto_depIdxs,
		MessageInfos:      file_ci_proto_msgTypes,
	}.Build()
	File_ci_proto = out.File
	file_ci_proto_rawDesc = nil
	file_ci_proto_goTypes = nil
	file_ci_proto_depIdxs = nil
}
