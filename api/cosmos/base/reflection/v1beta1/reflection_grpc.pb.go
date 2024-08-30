// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: cosmos/base/reflection/v1beta1/reflection.proto

package reflectionv1beta1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ReflectionService_ListAllInterfaces_FullMethodName   = "/cosmos.base.reflection.v1beta1.ReflectionService/ListAllInterfaces"
	ReflectionService_ListImplementations_FullMethodName = "/cosmos.base.reflection.v1beta1.ReflectionService/ListImplementations"
)

// ReflectionServiceClient is the client API for ReflectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// ReflectionService defines a service for interface reflection.
type ReflectionServiceClient interface {
	// ListAllInterfaces lists all the interfaces registered in the interface
	// registry.
	ListAllInterfaces(ctx context.Context, in *ListAllInterfacesRequest, opts ...grpc.CallOption) (*ListAllInterfacesResponse, error)
	// ListImplementations list all the concrete types that implement a given
	// interface.
	ListImplementations(ctx context.Context, in *ListImplementationsRequest, opts ...grpc.CallOption) (*ListImplementationsResponse, error)
}

type reflectionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReflectionServiceClient(cc grpc.ClientConnInterface) ReflectionServiceClient {
	return &reflectionServiceClient{cc}
}

func (c *reflectionServiceClient) ListAllInterfaces(ctx context.Context, in *ListAllInterfacesRequest, opts ...grpc.CallOption) (*ListAllInterfacesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAllInterfacesResponse)
	err := c.cc.Invoke(ctx, ReflectionService_ListAllInterfaces_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reflectionServiceClient) ListImplementations(ctx context.Context, in *ListImplementationsRequest, opts ...grpc.CallOption) (*ListImplementationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListImplementationsResponse)
	err := c.cc.Invoke(ctx, ReflectionService_ListImplementations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReflectionServiceServer is the server API for ReflectionService service.
// All implementations must embed UnimplementedReflectionServiceServer
// for forward compatibility.
//
// ReflectionService defines a service for interface reflection.
type ReflectionServiceServer interface {
	// ListAllInterfaces lists all the interfaces registered in the interface
	// registry.
	ListAllInterfaces(context.Context, *ListAllInterfacesRequest) (*ListAllInterfacesResponse, error)
	// ListImplementations list all the concrete types that implement a given
	// interface.
	ListImplementations(context.Context, *ListImplementationsRequest) (*ListImplementationsResponse, error)
	mustEmbedUnimplementedReflectionServiceServer()
}

// UnimplementedReflectionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedReflectionServiceServer struct{}

func (UnimplementedReflectionServiceServer) ListAllInterfaces(context.Context, *ListAllInterfacesRequest) (*ListAllInterfacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllInterfaces not implemented")
}
func (UnimplementedReflectionServiceServer) ListImplementations(context.Context, *ListImplementationsRequest) (*ListImplementationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListImplementations not implemented")
}
func (UnimplementedReflectionServiceServer) mustEmbedUnimplementedReflectionServiceServer() {}
func (UnimplementedReflectionServiceServer) testEmbeddedByValue()                           {}

// UnsafeReflectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReflectionServiceServer will
// result in compilation errors.
type UnsafeReflectionServiceServer interface {
	mustEmbedUnimplementedReflectionServiceServer()
}

func RegisterReflectionServiceServer(s grpc.ServiceRegistrar, srv ReflectionServiceServer) {
	// If the following call pancis, it indicates UnimplementedReflectionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ReflectionService_ServiceDesc, srv)
}

func _ReflectionService_ListAllInterfaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAllInterfacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReflectionServiceServer).ListAllInterfaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReflectionService_ListAllInterfaces_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReflectionServiceServer).ListAllInterfaces(ctx, req.(*ListAllInterfacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReflectionService_ListImplementations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListImplementationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReflectionServiceServer).ListImplementations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReflectionService_ListImplementations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReflectionServiceServer).ListImplementations(ctx, req.(*ListImplementationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReflectionService_ServiceDesc is the grpc.ServiceDesc for ReflectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReflectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.base.reflection.v1beta1.ReflectionService",
	HandlerType: (*ReflectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAllInterfaces",
			Handler:    _ReflectionService_ListAllInterfaces_Handler,
		},
		{
			MethodName: "ListImplementations",
			Handler:    _ReflectionService_ListImplementations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/base/reflection/v1beta1/reflection.proto",
}