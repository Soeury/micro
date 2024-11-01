package main

import (
	"context"
	"fmt"
	"micro/learn1/test15_unary/server/proto"
	"net"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// 服务器端拦截器 ：
//
//	实现：unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)

// 服务端拦截器实现：检查客户端传过来的 token 是否有效
func unaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "missing metadata")
	}

	if !Valid(md["authorization"]) {
		return nil, status.Errorf(codes.InvalidArgument, "invalid token")
	}

	m, err := handler(ctx, req) // 这里应该是实际处理 rpc 请求
	if err != nil {
		fmt.Printf("rpc failed with err:%v\n", err)
	}
	return m, err
}

// token 认定
func Valid(authorization []string) bool {

	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	return token == "yuanshuju.rabbit.qianming"
}

type Server struct {
	proto.UnimplementedFriendServer
}

func (s *Server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {

	// 接收元数据
	md, ok := metadata.FromIncomingContext(ctx)

	// 检查元数据是否存在
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "invalid request")
	}

	// check token
	vl := md.Get("token")
	if len(vl) == 0 || vl[0] != "yuanshuju.rabbit.qianming" {
		return nil, status.Error(codes.Unauthenticated, "invalid request")
	}

	// 返回响应
	reply := "Hello, " + in.GetName()
	return &proto.HelloResponse{Reply: reply}, nil
}

func main() {

	// 启动服务
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("net.Listen failed:%v\n", err)
		return
	}

	// 加载证书 (注意切换到当前路径下)
	creds, err := credentials.NewServerTLSFromFile("./server/certs/server.crt", "./server/certs/server.key")
	if err != nil {
		fmt.Printf("credentials.NewServerTLSFromFile failed:%v\n", err)
		return
	}

	// 创建rpc服务 + 注册证书 + 注册拦截器
	s := grpc.NewServer(
		grpc.Creds(creds),
		grpc.UnaryInterceptor(unaryInterceptor),
	)

	// 注册rpc服务
	proto.RegisterFriendServer(s, &Server{})

	// 启动rpc服务
	err = s.Serve(listener)
	if err != nil {
		fmt.Printf("s.Serve failed:%v\n", err)
		return
	}
}
