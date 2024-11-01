package main

import (
	"context"
	"fmt"
	"micro/learn1/test3/server/proto"
	"net"

	"google.golang.org/grpc" // 注意使用的 gRPC 包是这个包
)

// 提供服务的结构体对象内部需要实现 proto.UnimplementedGreeterServer
// 保证在没有完全使用对外提供的服务的情况下程序也可以正常运行
type server struct {
	proto.UnimplementedGreeterServer
}

// SayHello 是我们实现的方法，用来对外提供服务
// 这个方法的实现需要参照生成的两个文件中的内容
func (s *server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {

	reply := "hello " + in.GetName()
	return &proto.HelloResponse{Reply: reply}, nil
}

func main() {

	// 启动服务
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("net.Listen failed:%v\n", err)
		return
	}

	// 创建RPC服务
	s := grpc.NewServer()

	// 注册服务
	proto.RegisterGreeterServer(s, &server{})

	// 启动服务
	err = s.Serve(listener)
	if err != nil {
		fmt.Printf("s.Serve failed:%v\n", err)
		return
	}
}
