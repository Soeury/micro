// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.28.2
// source: learn1/test12_metadata/server/proto/hello.proto

package proto

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

// CareClient is the client API for Care service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CareClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type careClient struct {
	cc grpc.ClientConnInterface
}

func NewCareClient(cc grpc.ClientConnInterface) CareClient {
	return &careClient{cc}
}

func (c *careClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/proto.Care/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CareServer is the server API for Care service.
// All implementations must embed UnimplementedCareServer
// for forward compatibility
type CareServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
	mustEmbedUnimplementedCareServer()
}

// UnimplementedCareServer must be embedded to have forward compatible implementations.
type UnimplementedCareServer struct {
}

func (UnimplementedCareServer) SayHello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedCareServer) mustEmbedUnimplementedCareServer() {}

// UnsafeCareServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CareServer will
// result in compilation errors.
type UnsafeCareServer interface {
	mustEmbedUnimplementedCareServer()
}

func RegisterCareServer(s grpc.ServiceRegistrar, srv CareServer) {
	s.RegisterService(&Care_ServiceDesc, srv)
}

func _Care_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CareServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Care/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CareServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Care_ServiceDesc is the grpc.ServiceDesc for Care service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Care_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Care",
	HandlerType: (*CareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Care_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "learn1/test12_metadata/server/proto/hello.proto",
}
