// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.28.2
// source: learn1/test8_review/client/proto/store.proto

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

// SaleClient is the client API for Sale service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SaleClient interface {
	Products(ctx context.Context, in *Things, opts ...grpc.CallOption) (*Response, error)
	Update(ctx context.Context, in *UpdateThings, opts ...grpc.CallOption) (*Response, error)
}

type saleClient struct {
	cc grpc.ClientConnInterface
}

func NewSaleClient(cc grpc.ClientConnInterface) SaleClient {
	return &saleClient{cc}
}

func (c *saleClient) Products(ctx context.Context, in *Things, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Sale/Products", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *saleClient) Update(ctx context.Context, in *UpdateThings, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Sale/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SaleServer is the server API for Sale service.
// All implementations must embed UnimplementedSaleServer
// for forward compatibility
type SaleServer interface {
	Products(context.Context, *Things) (*Response, error)
	Update(context.Context, *UpdateThings) (*Response, error)
	mustEmbedUnimplementedSaleServer()
}

// UnimplementedSaleServer must be embedded to have forward compatible implementations.
type UnimplementedSaleServer struct {
}

func (UnimplementedSaleServer) Products(context.Context, *Things) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Products not implemented")
}
func (UnimplementedSaleServer) Update(context.Context, *UpdateThings) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSaleServer) mustEmbedUnimplementedSaleServer() {}

// UnsafeSaleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SaleServer will
// result in compilation errors.
type UnsafeSaleServer interface {
	mustEmbedUnimplementedSaleServer()
}

func RegisterSaleServer(s grpc.ServiceRegistrar, srv SaleServer) {
	s.RegisterService(&Sale_ServiceDesc, srv)
}

func _Sale_Products_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Things)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SaleServer).Products(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Sale/Products",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SaleServer).Products(ctx, req.(*Things))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sale_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateThings)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SaleServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Sale/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SaleServer).Update(ctx, req.(*UpdateThings))
	}
	return interceptor(ctx, in, info, handler)
}

// Sale_ServiceDesc is the grpc.ServiceDesc for Sale service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sale_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Sale",
	HandlerType: (*SaleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Products",
			Handler:    _Sale_Products_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Sale_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "learn1/test8_review/client/proto/store.proto",
}
