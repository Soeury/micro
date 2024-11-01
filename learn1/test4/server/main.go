package main

import (
	"context"
	"fmt"
	"micro/learn1/test4/server/proto"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) Add(ctx context.Context, in *proto.AddRequest) (*proto.AddResponse, error) {
	ret := in.Num1 + in.Num2
	return &proto.AddResponse{Ret: ret}, nil
}

func main() {

	// 启动服务
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("net.Dial failed: %v\n", err)
		return
	}

	// 创建RPC服务
	s := grpc.NewServer()

	// 注册服务 :  server + 对象
	proto.RegisterGreeterServer(s, &server{})

	// 启动服务
	err = s.Serve(listener)
	if err != nil {
		fmt.Printf("s.Serve failed: %v\n", err)
		return
	}
}
