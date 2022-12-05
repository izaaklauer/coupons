// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: coupons/v1/coupons.proto

package couponsv1

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

// CouponsServiceClient is the client API for CouponsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CouponsServiceClient interface {
	// HelloWorld is the default rpc. Feel free to delete.
	HelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error)
}

type couponsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCouponsServiceClient(cc grpc.ClientConnInterface) CouponsServiceClient {
	return &couponsServiceClient{cc}
}

func (c *couponsServiceClient) HelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error) {
	out := new(HelloWorldResponse)
	err := c.cc.Invoke(ctx, "/coupons.v1.CouponsService/HelloWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CouponsServiceServer is the server API for CouponsService service.
// All implementations must embed UnimplementedCouponsServiceServer
// for forward compatibility
type CouponsServiceServer interface {
	// HelloWorld is the default rpc. Feel free to delete.
	HelloWorld(context.Context, *HelloWorldRequest) (*HelloWorldResponse, error)
	mustEmbedUnimplementedCouponsServiceServer()
}

// UnimplementedCouponsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCouponsServiceServer struct {
}

func (UnimplementedCouponsServiceServer) HelloWorld(context.Context, *HelloWorldRequest) (*HelloWorldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloWorld not implemented")
}
func (UnimplementedCouponsServiceServer) mustEmbedUnimplementedCouponsServiceServer() {
}

// UnsafeCouponsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CouponsServiceServer will
// result in compilation errors.
type UnsafeCouponsServiceServer interface {
	mustEmbedUnimplementedCouponsServiceServer()
}

func RegisterCouponsServiceServer(s grpc.ServiceRegistrar, srv CouponsServiceServer) {
	s.RegisterService(&CouponsService_ServiceDesc, srv)
}

func _CouponsService_HelloWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloWorldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponsServiceServer).HelloWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coupons.v1.CouponsService/HelloWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponsServiceServer).HelloWorld(ctx, req.(*HelloWorldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CouponsService_ServiceDesc is the grpc.ServiceDesc for CouponsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CouponsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "coupons.v1.CouponsService",
	HandlerType: (*CouponsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloWorld",
			Handler:    _CouponsService_HelloWorld_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coupons/v1/coupons.proto",
}
