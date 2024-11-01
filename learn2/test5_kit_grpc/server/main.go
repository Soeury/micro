package main

import (
	"context"
	"errors"
	"fmt"
	"micro/learn2/test5_kit_grpc/server/proto"
	"net"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

// 1. service
type AddService interface {
	Sum(context.Context, int, int) (int, error)
	Append(context.Context, string, string) (string, error)
}

type addService struct{}

func (addService) Sum(_ context.Context, a int, b int) (int, error) {

	return a + b, nil
}

func (addService) Append(_ context.Context, a string, b string) (string, error) {

	if a == "" && b == "" {
		return "", errors.New("all strings are empty")
	}

	return a + b, nil
}

type SumRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type SumResponse struct {
	Value int    `json:"value"`
	Err   string `json:"err,omitempty"`
}

type AppendRequest struct {
	A string `json:"a"`
	B string `json:"b"`
}

type AppendResponse struct {
	Value string `json:"value"`
	Err   string `json:"err,omitempty"`
}

// 2. endpoint
// 上面的部分都是为这里的编码做准备
func MakeSumEndpoint(srv AddService) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SumRequest)
		v, err := srv.Sum(ctx, req.A, req.B)
		if err != nil {
			return SumResponse{Value: v, Err: err.Error()}, nil
		}
		return SumResponse{Value: v}, nil
	}
}

func MakeAppendEndpoint(srv AddService) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AppendRequest)
		v, err := srv.Append(ctx, req.A, req.B)
		if err != nil {
			return AppendResponse{Value: v, Err: err.Error()}, nil
		}
		return AppendResponse{Value: v}, nil
	}
}

// 3. transport
// gprc 服务与实现：[对于客户端来说可以调用的]
type GrpcServer struct {
	proto.UnimplementedAddServer
	sum    grpctransport.Handler // 注意这些字段的类型，和之前写的 grpc 不太一样(之前写的这里没有)
	append grpctransport.Handler // 使用特定的处理函数，目的是为了将 grpc 与 go-kit 连接起来
}

func (s *GrpcServer) Sum(ctx context.Context, in *proto.SumRequest) (*proto.SumResponse, error) {

	// 注意这里的调用 !
	// SerceGRPC 是 Handler 接口中的一个方法
	_, resp, err := s.sum.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*proto.SumResponse), nil
}

func (s *GrpcServer) Append(ctx context.Context, in *proto.AppendRequest) (*proto.AppendResponse, error) {

	// 注意这里的调用 !
	_, resp, err := s.append.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*proto.AppendResponse), nil
}

// 编码 & 解码
// 这些函数是gRPC服务中常见的模式，用于在gRPC消息（由protobuf定义并自动生成的）和自定义Go结构体之间进行转换
// 这样的转换是必要的，因为gRPC消息通常使用protobuf的基本类型（如int64、string等）
// 而服务逻辑可能使用更具体的Go类型（如int、自定义结构体等）
func DecodeGRPCSumRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*proto.SumRequest)
	return SumRequest{A: int(req.A), B: int(req.B)}, nil
}

func DecodeGRPCAppendRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*proto.AppendRequest)
	return AppendRequest{A: req.A, B: req.B}, nil
}

func EncodeGRPCSumResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(SumResponse)
	return &proto.SumResponse{Value: int64(resp.Value)}, nil
}

func EncodeGRPCAppendResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(AppendResponse)
	return &proto.AppendResponse{Value: resp.Value}, nil
}

// 这里是最后的步骤: 将上述的所有内容连接在 grpc 服务内部
func NewGRPCserver(srv AddService) proto.AddServer {

	// 这里返回的 *Server 结构体对象实现了定义的 Handler 接口
	return &GrpcServer{
		sum: grpctransport.NewServer(
			MakeSumEndpoint(srv),
			DecodeGRPCSumRequest,
			EncodeGRPCSumResponse,
		),
		append: grpctransport.NewServer(
			MakeAppendEndpoint(srv),
			DecodeGRPCAppendRequest,
			EncodeGRPCAppendResponse,
		),
	}
}

func main() {

	srv := addService{}      // 定义服务对象，作为创建 GRPC 对象的参数
	gs := NewGRPCserver(srv) // 返回的是实现grpc服务的对象

	// 启动服务
	listner, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("net.Listen failed:%v\n", err)
		return
	}

	// 创建 rpc 服务
	s := grpc.NewServer()

	// 注册 rpc 服务
	proto.RegisterAddServer(s, gs)

	// 启动 rpc 服务
	err = s.Serve(listner)
	if err != nil {
		fmt.Printf("s.serve failed:%v\n", err)
		return
	}
}
