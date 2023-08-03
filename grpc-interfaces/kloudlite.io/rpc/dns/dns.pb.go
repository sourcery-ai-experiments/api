// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: dns.proto

package dns

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

type GetAccountDomainsIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=accountId,proto3" json:"accountId,omitempty"`
}

func (x *GetAccountDomainsIn) Reset() {
	*x = GetAccountDomainsIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dns_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountDomainsIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountDomainsIn) ProtoMessage() {}

func (x *GetAccountDomainsIn) ProtoReflect() protoreflect.Message {
	mi := &file_dns_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountDomainsIn.ProtoReflect.Descriptor instead.
func (*GetAccountDomainsIn) Descriptor() ([]byte, []int) {
	return file_dns_proto_rawDescGZIP(), []int{0}
}

func (x *GetAccountDomainsIn) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type GetAccountDomainsOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domains []string `protobuf:"bytes,1,rep,name=domains,proto3" json:"domains,omitempty"`
}

func (x *GetAccountDomainsOut) Reset() {
	*x = GetAccountDomainsOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dns_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountDomainsOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountDomainsOut) ProtoMessage() {}

func (x *GetAccountDomainsOut) ProtoReflect() protoreflect.Message {
	mi := &file_dns_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountDomainsOut.ProtoReflect.Descriptor instead.
func (*GetAccountDomainsOut) Descriptor() ([]byte, []int) {
	return file_dns_proto_rawDescGZIP(), []int{1}
}

func (x *GetAccountDomainsOut) GetDomains() []string {
	if x != nil {
		return x.Domains
	}
	return nil
}

var File_dns_proto protoreflect.FileDescriptor

var file_dns_proto_rawDesc = []byte{
	0x0a, 0x09, 0x64, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73,
	0x49, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x22, 0x30, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x73, 0x4f, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x73, 0x32, 0x47, 0x0a, 0x03, 0x44, 0x4e, 0x53, 0x12, 0x40, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x14,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x73, 0x49, 0x6e, 0x1a, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x4f, 0x75, 0x74, 0x42, 0x16, 0x5a, 0x14, 0x6b,
	0x6c, 0x6f, 0x75, 0x64, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x64, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dns_proto_rawDescOnce sync.Once
	file_dns_proto_rawDescData = file_dns_proto_rawDesc
)

func file_dns_proto_rawDescGZIP() []byte {
	file_dns_proto_rawDescOnce.Do(func() {
		file_dns_proto_rawDescData = protoimpl.X.CompressGZIP(file_dns_proto_rawDescData)
	})
	return file_dns_proto_rawDescData
}

var file_dns_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dns_proto_goTypes = []interface{}{
	(*GetAccountDomainsIn)(nil),  // 0: GetAccountDomainsIn
	(*GetAccountDomainsOut)(nil), // 1: GetAccountDomainsOut
}
var file_dns_proto_depIdxs = []int32{
	0, // 0: DNS.GetAccountDomains:input_type -> GetAccountDomainsIn
	1, // 1: DNS.GetAccountDomains:output_type -> GetAccountDomainsOut
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dns_proto_init() }
func file_dns_proto_init() {
	if File_dns_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dns_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountDomainsIn); i {
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
		file_dns_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountDomainsOut); i {
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
			RawDescriptor: file_dns_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dns_proto_goTypes,
		DependencyIndexes: file_dns_proto_depIdxs,
		MessageInfos:      file_dns_proto_msgTypes,
	}.Build()
	File_dns_proto = out.File
	file_dns_proto_rawDesc = nil
	file_dns_proto_goTypes = nil
	file_dns_proto_depIdxs = nil
}
