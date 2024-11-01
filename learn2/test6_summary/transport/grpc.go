package transport

import (
	"context"
	"micro/learn2/test6_summary/endpoint"
	"micro/learn2/test6_summary/middleware"
	"micro/learn2/test6_summary/proto/sum"
	"micro/learn2/test6_summary/service"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	"golang.org/x/time/rate"
)

// GRPC

type GrpcServer struct {
	sum.UnimplementedComputerServer
	add    grpctransport.Handler // 注意这些字段的类型，和之前写的 grpc 不太一样(之前写的这里没有)
	append grpctransport.Handler // 使用特定的处理函数，目的是为了将 grpc 与 go-kit 连接起来
}

func (s *GrpcServer) Sum(ctx context.Context, in *sum.AddPRCRequest) (*sum.AddRPCResponse, error) {

	// 注意这里的调用 !
	// SerceGRPC 是 Handler 接口中的一个方法
	_, resp, err := s.add.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*sum.AddRPCResponse), nil
}

func (s *GrpcServer) Append(ctx context.Context, in *sum.AppendRPCRequest) (*sum.AppendRPCResponse, error) {

	// 注意这里的调用 !
	_, resp, err := s.append.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*sum.AppendRPCResponse), nil
}

// 编码 & 解码
// 这些函数是gRPC服务中常见的模式，用于在gRPC消息（由protobuf定义并自动生成的）和自定义Go结构体之间进行转换
// 这样的转换是必要的，因为gRPC消息通常使用protobuf的基本类型（如int64、string等）
// 而服务逻辑可能使用更具体的Go类型（如int、自定义结构体等）

// 外部的 grpc 请求结构体类型，转换成内部的请求结构体格式，因为外部是发起 grpc 请求我们内部的服务
func DecodeGRPCAddRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*sum.AddPRCRequest)
	return endpoint.AddRequest{Num1: req.Num1, Num2: req.Num2}, nil
}

func DecodeGRPCAppendRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*sum.AppendRPCRequest)
	return endpoint.AppendRequest{Str1: req.Str1, Str2: req.Str2}, nil
}

// 内部返回数据，包装成 grpc 响应结构体，返回给外部
func EncodeGRPCAddResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.AddResponse)
	return &sum.AddRPCResponse{Ret: resp.Ret}, nil
}

func EncodeGRPCAppendResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.AppendResponse)
	return &sum.AppendRPCResponse{Ret: resp.Ret}, nil
}

// 返回 grpc 服务
func NewGRPCServer(srv service.ComputerService, logger log.Logger) sum.ComputerServer {

	add := endpoint.MakeAddEndpoint(srv)
	// 添加logger中间件
	// transport 层的日志好像只能使用在 transport 层? ? ?
	add = middleware.LoggerMiddleWare(logger)(add) // 这前面一部分返回的是函数，也可以调用
	// 添加rate中间件
	add = middleware.RateMiddleWare(rate.NewLimiter(1, 1))(add)
	// ......

	append := endpoint.MakeAppendEndpoint(srv)
	append = middleware.LoggerMiddleWare(logger)(append)
	append = middleware.RateMiddleWare(rate.NewLimiter(1, 1))(append)

	return &GrpcServer{
		add: grpctransport.NewServer(
			add,
			DecodeGRPCAddRequest,
			EncodeGRPCAddResponse,
		),

		append: grpctransport.NewServer(
			append,
			DecodeGRPCAppendRequest,
			EncodeGRPCAppendResponse,
		),
	}
}
