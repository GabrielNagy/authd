// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: authd.proto

package authd

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

const (
	Pam_TestPam_FullMethodName = "/pam/TestPam"
)

// PamClient is the client API for Pam service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PamClient interface {
	TestPam(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StringResponse, error)
}

type pamClient struct {
	cc grpc.ClientConnInterface
}

func NewPamClient(cc grpc.ClientConnInterface) PamClient {
	return &pamClient{cc}
}

func (c *pamClient) TestPam(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StringResponse, error) {
	out := new(StringResponse)
	err := c.cc.Invoke(ctx, Pam_TestPam_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PamServer is the server API for Pam service.
// All implementations must embed UnimplementedPamServer
// for forward compatibility
type PamServer interface {
	TestPam(context.Context, *Empty) (*StringResponse, error)
	mustEmbedUnimplementedPamServer()
}

// UnimplementedPamServer must be embedded to have forward compatible implementations.
type UnimplementedPamServer struct {
}

func (UnimplementedPamServer) TestPam(context.Context, *Empty) (*StringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestPam not implemented")
}
func (UnimplementedPamServer) mustEmbedUnimplementedPamServer() {}

// UnsafePamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PamServer will
// result in compilation errors.
type UnsafePamServer interface {
	mustEmbedUnimplementedPamServer()
}

func RegisterPamServer(s grpc.ServiceRegistrar, srv PamServer) {
	s.RegisterService(&Pam_ServiceDesc, srv)
}

func _Pam_TestPam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PamServer).TestPam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pam_TestPam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PamServer).TestPam(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Pam_ServiceDesc is the grpc.ServiceDesc for Pam service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Pam_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pam",
	HandlerType: (*PamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestPam",
			Handler:    _Pam_TestPam_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authd.proto",
}

const (
	Nss_TestNSS_FullMethodName = "/nss/TestNSS"
)

// NssClient is the client API for Nss service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NssClient interface {
	TestNSS(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StringResponse, error)
}

type nssClient struct {
	cc grpc.ClientConnInterface
}

func NewNssClient(cc grpc.ClientConnInterface) NssClient {
	return &nssClient{cc}
}

func (c *nssClient) TestNSS(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StringResponse, error) {
	out := new(StringResponse)
	err := c.cc.Invoke(ctx, Nss_TestNSS_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NssServer is the server API for Nss service.
// All implementations must embed UnimplementedNssServer
// for forward compatibility
type NssServer interface {
	TestNSS(context.Context, *Empty) (*StringResponse, error)
	mustEmbedUnimplementedNssServer()
}

// UnimplementedNssServer must be embedded to have forward compatible implementations.
type UnimplementedNssServer struct {
}

func (UnimplementedNssServer) TestNSS(context.Context, *Empty) (*StringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestNSS not implemented")
}
func (UnimplementedNssServer) mustEmbedUnimplementedNssServer() {}

// UnsafeNssServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NssServer will
// result in compilation errors.
type UnsafeNssServer interface {
	mustEmbedUnimplementedNssServer()
}

func RegisterNssServer(s grpc.ServiceRegistrar, srv NssServer) {
	s.RegisterService(&Nss_ServiceDesc, srv)
}

func _Nss_TestNSS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NssServer).TestNSS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nss_TestNSS_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NssServer).TestNSS(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Nss_ServiceDesc is the grpc.ServiceDesc for Nss service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Nss_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nss",
	HandlerType: (*NssServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestNSS",
			Handler:    _Nss_TestNSS_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authd.proto",
}
