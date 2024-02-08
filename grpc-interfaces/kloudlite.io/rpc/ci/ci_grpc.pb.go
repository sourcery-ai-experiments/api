// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: ci.proto

package ci

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CIClient is the client API for CI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CIClient interface {
	CreateHarborProject(ctx context.Context, in *HarborProjectIn, opts ...grpc.CallOption) (*HarborProjectOut, error)
	DeleteHarborProject(ctx context.Context, in *HarborProjectIn, opts ...grpc.CallOption) (*HarborProjectOut, error)
}

type cIClient struct {
	cc grpc.ClientConnInterface
}

func NewCIClient(cc grpc.ClientConnInterface) CIClient {
	return &cIClient{cc}
}

func (c *cIClient) CreateHarborProject(ctx context.Context, in *HarborProjectIn, opts ...grpc.CallOption) (*HarborProjectOut, error) {
	out := new(HarborProjectOut)
	err := c.cc.Invoke(ctx, "/CI/CreateHarborProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cIClient) DeleteHarborProject(ctx context.Context, in *HarborProjectIn, opts ...grpc.CallOption) (*HarborProjectOut, error) {
	out := new(HarborProjectOut)
	err := c.cc.Invoke(ctx, "/CI/DeleteHarborProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CIServer is the server API for CI service.
// All implementations must embed UnimplementedCIServer
// for forward compatibility
type CIServer interface {
	CreateHarborProject(context.Context, *HarborProjectIn) (*HarborProjectOut, error)
	DeleteHarborProject(context.Context, *HarborProjectIn) (*HarborProjectOut, error)
	mustEmbedUnimplementedCIServer()
}

// UnimplementedCIServer must be embedded to have forward compatible implementations.
type UnimplementedCIServer struct {
}

func (UnimplementedCIServer) CreateHarborProject(context.Context, *HarborProjectIn) (*HarborProjectOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHarborProject not implemented")
}
func (UnimplementedCIServer) DeleteHarborProject(context.Context, *HarborProjectIn) (*HarborProjectOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHarborProject not implemented")
}
func (UnimplementedCIServer) mustEmbedUnimplementedCIServer() {}

// UnsafeCIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CIServer will
// result in compilation errors.
type UnsafeCIServer interface {
	mustEmbedUnimplementedCIServer()
}

func RegisterCIServer(s grpc.ServiceRegistrar, srv CIServer) {
	s.RegisterService(&CI_ServiceDesc, srv)
}

func _CI_CreateHarborProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HarborProjectIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CIServer).CreateHarborProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CI/CreateHarborProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CIServer).CreateHarborProject(ctx, req.(*HarborProjectIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _CI_DeleteHarborProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HarborProjectIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CIServer).DeleteHarborProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CI/DeleteHarborProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CIServer).DeleteHarborProject(ctx, req.(*HarborProjectIn))
	}
	return interceptor(ctx, in, info, handler)
}

// CI_ServiceDesc is the grpc.ServiceDesc for CI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CI",
	HandlerType: (*CIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHarborProject",
			Handler:    _CI_CreateHarborProject_Handler,
		},
		{
			MethodName: "DeleteHarborProject",
			Handler:    _CI_DeleteHarborProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ci.proto",
}
