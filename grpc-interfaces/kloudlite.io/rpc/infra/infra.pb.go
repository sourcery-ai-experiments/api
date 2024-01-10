// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.2
// source: infra.proto

package infra

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

type GetVpnDeviceIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceName  string `protobuf:"bytes,1,opt,name=deviceName,proto3" json:"deviceName,omitempty"`
	AccountName string `protobuf:"bytes,2,opt,name=accountName,proto3" json:"accountName,omitempty"`
	ClusterName string `protobuf:"bytes,3,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
}

func (x *GetVpnDeviceIn) Reset() {
	*x = GetVpnDeviceIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVpnDeviceIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVpnDeviceIn) ProtoMessage() {}

func (x *GetVpnDeviceIn) ProtoReflect() protoreflect.Message {
	mi := &file_infra_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVpnDeviceIn.ProtoReflect.Descriptor instead.
func (*GetVpnDeviceIn) Descriptor() ([]byte, []int) {
	return file_infra_proto_rawDescGZIP(), []int{0}
}

func (x *GetVpnDeviceIn) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *GetVpnDeviceIn) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *GetVpnDeviceIn) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

type GetVpnDeviceOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VpnDevice []byte `protobuf:"bytes,1,opt,name=vpnDevice,proto3" json:"vpnDevice,omitempty"`
	WgConfig  []byte `protobuf:"bytes,2,opt,name=wgConfig,proto3" json:"wgConfig,omitempty"`
}

func (x *GetVpnDeviceOut) Reset() {
	*x = GetVpnDeviceOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVpnDeviceOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVpnDeviceOut) ProtoMessage() {}

func (x *GetVpnDeviceOut) ProtoReflect() protoreflect.Message {
	mi := &file_infra_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVpnDeviceOut.ProtoReflect.Descriptor instead.
func (*GetVpnDeviceOut) Descriptor() ([]byte, []int) {
	return file_infra_proto_rawDescGZIP(), []int{1}
}

func (x *GetVpnDeviceOut) GetVpnDevice() []byte {
	if x != nil {
		return x.VpnDevice
	}
	return nil
}

func (x *GetVpnDeviceOut) GetWgConfig() []byte {
	if x != nil {
		return x.WgConfig
	}
	return nil
}

type DeleteVpnDeviceIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountName string `protobuf:"bytes,1,opt,name=accountName,proto3" json:"accountName,omitempty"`
	Id          string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteVpnDeviceIn) Reset() {
	*x = DeleteVpnDeviceIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteVpnDeviceIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteVpnDeviceIn) ProtoMessage() {}

func (x *DeleteVpnDeviceIn) ProtoReflect() protoreflect.Message {
	mi := &file_infra_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteVpnDeviceIn.ProtoReflect.Descriptor instead.
func (*DeleteVpnDeviceIn) Descriptor() ([]byte, []int) {
	return file_infra_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteVpnDeviceIn) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *DeleteVpnDeviceIn) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteVpnDeviceOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DeleteVpnDeviceOut) Reset() {
	*x = DeleteVpnDeviceOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteVpnDeviceOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteVpnDeviceOut) ProtoMessage() {}

func (x *DeleteVpnDeviceOut) ProtoReflect() protoreflect.Message {
	mi := &file_infra_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteVpnDeviceOut.ProtoReflect.Descriptor instead.
func (*DeleteVpnDeviceOut) Descriptor() ([]byte, []int) {
	return file_infra_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteVpnDeviceOut) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type UpsertVpnDeviceIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	VpnDevice   []byte `protobuf:"bytes,2,opt,name=vpnDevice,proto3" json:"vpnDevice,omitempty"`
	AccountName string `protobuf:"bytes,3,opt,name=accountName,proto3" json:"accountName,omitempty"`
	ClusterName string `protobuf:"bytes,4,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
}

func (x *UpsertVpnDeviceIn) Reset() {
	*x = UpsertVpnDeviceIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertVpnDeviceIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertVpnDeviceIn) ProtoMessage() {}

func (x *UpsertVpnDeviceIn) ProtoReflect() protoreflect.Message {
	mi := &file_infra_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertVpnDeviceIn.ProtoReflect.Descriptor instead.
func (*UpsertVpnDeviceIn) Descriptor() ([]byte, []int) {
	return file_infra_proto_rawDescGZIP(), []int{4}
}

func (x *UpsertVpnDeviceIn) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpsertVpnDeviceIn) GetVpnDevice() []byte {
	if x != nil {
		return x.VpnDevice
	}
	return nil
}

func (x *UpsertVpnDeviceIn) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *UpsertVpnDeviceIn) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

type UpsertVpnDeviceOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VpnDevice []byte `protobuf:"bytes,1,opt,name=vpnDevice,proto3" json:"vpnDevice,omitempty"`
	WgConfig  []byte `protobuf:"bytes,2,opt,name=wgConfig,proto3" json:"wgConfig,omitempty"`
}

func (x *UpsertVpnDeviceOut) Reset() {
	*x = UpsertVpnDeviceOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertVpnDeviceOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertVpnDeviceOut) ProtoMessage() {}

func (x *UpsertVpnDeviceOut) ProtoReflect() protoreflect.Message {
	mi := &file_infra_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertVpnDeviceOut.ProtoReflect.Descriptor instead.
func (*UpsertVpnDeviceOut) Descriptor() ([]byte, []int) {
	return file_infra_proto_rawDescGZIP(), []int{5}
}

func (x *UpsertVpnDeviceOut) GetVpnDevice() []byte {
	if x != nil {
		return x.VpnDevice
	}
	return nil
}

func (x *UpsertVpnDeviceOut) GetWgConfig() []byte {
	if x != nil {
		return x.WgConfig
	}
	return nil
}

type GetClusterIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	UserName    string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	UserEmail   string `protobuf:"bytes,3,opt,name=userEmail,proto3" json:"userEmail,omitempty"`
	AccountName string `protobuf:"bytes,4,opt,name=accountName,proto3" json:"accountName,omitempty"`
	ClusterName string `protobuf:"bytes,5,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
}

func (x *GetClusterIn) Reset() {
	*x = GetClusterIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterIn) ProtoMessage() {}

func (x *GetClusterIn) ProtoReflect() protoreflect.Message {
	mi := &file_infra_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterIn.ProtoReflect.Descriptor instead.
func (*GetClusterIn) Descriptor() ([]byte, []int) {
	return file_infra_proto_rawDescGZIP(), []int{6}
}

func (x *GetClusterIn) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetClusterIn) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *GetClusterIn) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *GetClusterIn) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *GetClusterIn) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

type GetClusterOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageQueueTopic string `protobuf:"bytes,1,opt,name=messageQueueTopic,proto3" json:"messageQueueTopic,omitempty"`
	DnsHost           string `protobuf:"bytes,2,opt,name=dnsHost,proto3" json:"dnsHost,omitempty"`
	IACJobName        string `protobuf:"bytes,3,opt,name=IACJobName,proto3" json:"IACJobName,omitempty"`
	IACJobNamespace   string `protobuf:"bytes,4,opt,name=IACJobNamespace,proto3" json:"IACJobNamespace,omitempty"`
}

func (x *GetClusterOut) Reset() {
	*x = GetClusterOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterOut) ProtoMessage() {}

func (x *GetClusterOut) ProtoReflect() protoreflect.Message {
	mi := &file_infra_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterOut.ProtoReflect.Descriptor instead.
func (*GetClusterOut) Descriptor() ([]byte, []int) {
	return file_infra_proto_rawDescGZIP(), []int{7}
}

func (x *GetClusterOut) GetMessageQueueTopic() string {
	if x != nil {
		return x.MessageQueueTopic
	}
	return ""
}

func (x *GetClusterOut) GetDnsHost() string {
	if x != nil {
		return x.DnsHost
	}
	return ""
}

func (x *GetClusterOut) GetIACJobName() string {
	if x != nil {
		return x.IACJobName
	}
	return ""
}

func (x *GetClusterOut) GetIACJobNamespace() string {
	if x != nil {
		return x.IACJobNamespace
	}
	return ""
}

type GetNodepoolIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	UserName     string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	UserEmail    string `protobuf:"bytes,3,opt,name=userEmail,proto3" json:"userEmail,omitempty"`
	AccountName  string `protobuf:"bytes,4,opt,name=accountName,proto3" json:"accountName,omitempty"`
	ClusterName  string `protobuf:"bytes,5,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
	NodepoolName string `protobuf:"bytes,6,opt,name=nodepoolName,proto3" json:"nodepoolName,omitempty"`
}

func (x *GetNodepoolIn) Reset() {
	*x = GetNodepoolIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodepoolIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodepoolIn) ProtoMessage() {}

func (x *GetNodepoolIn) ProtoReflect() protoreflect.Message {
	mi := &file_infra_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodepoolIn.ProtoReflect.Descriptor instead.
func (*GetNodepoolIn) Descriptor() ([]byte, []int) {
	return file_infra_proto_rawDescGZIP(), []int{8}
}

func (x *GetNodepoolIn) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetNodepoolIn) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *GetNodepoolIn) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *GetNodepoolIn) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *GetNodepoolIn) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *GetNodepoolIn) GetNodepoolName() string {
	if x != nil {
		return x.NodepoolName
	}
	return ""
}

type GetNodepoolOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IACJobName      string `protobuf:"bytes,1,opt,name=IACJobName,proto3" json:"IACJobName,omitempty"`
	IACJobNamespace string `protobuf:"bytes,2,opt,name=IACJobNamespace,proto3" json:"IACJobNamespace,omitempty"`
}

func (x *GetNodepoolOut) Reset() {
	*x = GetNodepoolOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodepoolOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodepoolOut) ProtoMessage() {}

func (x *GetNodepoolOut) ProtoReflect() protoreflect.Message {
	mi := &file_infra_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodepoolOut.ProtoReflect.Descriptor instead.
func (*GetNodepoolOut) Descriptor() ([]byte, []int) {
	return file_infra_proto_rawDescGZIP(), []int{9}
}

func (x *GetNodepoolOut) GetIACJobName() string {
	if x != nil {
		return x.IACJobName
	}
	return ""
}

func (x *GetNodepoolOut) GetIACJobNamespace() string {
	if x != nil {
		return x.IACJobNamespace
	}
	return ""
}

var File_infra_proto protoreflect.FileDescriptor

var file_infra_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x56, 0x70, 0x6e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x12,
	0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x4b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x56, 0x70, 0x6e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x4f, 0x75, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x70, 0x6e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x76, 0x70, 0x6e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x77, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x22, 0x45, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x70, 0x6e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x56, 0x70, 0x6e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x75, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x85, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74,
	0x56, 0x70, 0x6e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x76,
	0x70, 0x6e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x76, 0x70, 0x6e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4e, 0x0a,
	0x12, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x56, 0x70, 0x6e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x4f, 0x75, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x70, 0x6e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x76, 0x70, 0x6e, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x77, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xa4, 0x01,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0xa1, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x6e, 0x73, 0x48, 0x6f, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x6e, 0x73, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x49, 0x41, 0x43, 0x4a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x49, 0x41, 0x43, 0x4a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28,
	0x0a, 0x0f, 0x49, 0x41, 0x43, 0x4a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x49, 0x41, 0x43, 0x4a, 0x6f, 0x62, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0xc9, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x4e, 0x6f, 0x64, 0x65, 0x70, 0x6f, 0x6f, 0x6c, 0x49, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x70, 0x6f, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x70, 0x6f, 0x6f, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x5a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x70,
	0x6f, 0x6f, 0x6c, 0x4f, 0x75, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x41, 0x43, 0x4a, 0x6f, 0x62,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x49, 0x41, 0x43, 0x4a,
	0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x49, 0x41, 0x43, 0x4a, 0x6f, 0x62,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x49, 0x41, 0x43, 0x4a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x32, 0x8f, 0x02, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x12, 0x2b, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x1a, 0x0e, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x12, 0x2e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4e, 0x6f,
	0x64, 0x65, 0x70, 0x6f, 0x6f, 0x6c, 0x12, 0x0e, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65,
	0x70, 0x6f, 0x6f, 0x6c, 0x49, 0x6e, 0x1a, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65,
	0x70, 0x6f, 0x6f, 0x6c, 0x4f, 0x75, 0x74, 0x12, 0x31, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x56, 0x70,
	0x6e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x70, 0x6e,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x1a, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x70,
	0x6e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x75, 0x74, 0x12, 0x3a, 0x0a, 0x0f, 0x55, 0x70,
	0x73, 0x65, 0x72, 0x74, 0x56, 0x70, 0x6e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x2e,
	0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x56, 0x70, 0x6e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x6e, 0x1a, 0x13, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x56, 0x70, 0x6e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x4f, 0x75, 0x74, 0x12, 0x3a, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x56, 0x70, 0x6e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x56, 0x70, 0x6e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x1a, 0x13, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x70, 0x6e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4f,
	0x75, 0x74, 0x42, 0x18, 0x5a, 0x16, 0x6b, 0x6c, 0x6f, 0x75, 0x64, 0x6c, 0x69, 0x74, 0x65, 0x2e,
	0x69, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_proto_rawDescOnce sync.Once
	file_infra_proto_rawDescData = file_infra_proto_rawDesc
)

func file_infra_proto_rawDescGZIP() []byte {
	file_infra_proto_rawDescOnce.Do(func() {
		file_infra_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_proto_rawDescData)
	})
	return file_infra_proto_rawDescData
}

var file_infra_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_infra_proto_goTypes = []interface{}{
	(*GetVpnDeviceIn)(nil),     // 0: GetVpnDeviceIn
	(*GetVpnDeviceOut)(nil),    // 1: GetVpnDeviceOut
	(*DeleteVpnDeviceIn)(nil),  // 2: DeleteVpnDeviceIn
	(*DeleteVpnDeviceOut)(nil), // 3: DeleteVpnDeviceOut
	(*UpsertVpnDeviceIn)(nil),  // 4: UpsertVpnDeviceIn
	(*UpsertVpnDeviceOut)(nil), // 5: UpsertVpnDeviceOut
	(*GetClusterIn)(nil),       // 6: GetClusterIn
	(*GetClusterOut)(nil),      // 7: GetClusterOut
	(*GetNodepoolIn)(nil),      // 8: GetNodepoolIn
	(*GetNodepoolOut)(nil),     // 9: GetNodepoolOut
}
var file_infra_proto_depIdxs = []int32{
	6, // 0: Infra.GetCluster:input_type -> GetClusterIn
	8, // 1: Infra.GetNodepool:input_type -> GetNodepoolIn
	0, // 2: Infra.GetVpnDevice:input_type -> GetVpnDeviceIn
	4, // 3: Infra.UpsertVpnDevice:input_type -> UpsertVpnDeviceIn
	2, // 4: Infra.DeleteVpnDevice:input_type -> DeleteVpnDeviceIn
	7, // 5: Infra.GetCluster:output_type -> GetClusterOut
	9, // 6: Infra.GetNodepool:output_type -> GetNodepoolOut
	1, // 7: Infra.GetVpnDevice:output_type -> GetVpnDeviceOut
	5, // 8: Infra.UpsertVpnDevice:output_type -> UpsertVpnDeviceOut
	3, // 9: Infra.DeleteVpnDevice:output_type -> DeleteVpnDeviceOut
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_infra_proto_init() }
func file_infra_proto_init() {
	if File_infra_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVpnDeviceIn); i {
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
		file_infra_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVpnDeviceOut); i {
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
		file_infra_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteVpnDeviceIn); i {
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
		file_infra_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteVpnDeviceOut); i {
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
		file_infra_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertVpnDeviceIn); i {
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
		file_infra_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertVpnDeviceOut); i {
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
		file_infra_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterIn); i {
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
		file_infra_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterOut); i {
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
		file_infra_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodepoolIn); i {
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
		file_infra_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodepoolOut); i {
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
			RawDescriptor: file_infra_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_infra_proto_goTypes,
		DependencyIndexes: file_infra_proto_depIdxs,
		MessageInfos:      file_infra_proto_msgTypes,
	}.Build()
	File_infra_proto = out.File
	file_infra_proto_rawDesc = nil
	file_infra_proto_goTypes = nil
	file_infra_proto_depIdxs = nil
}
