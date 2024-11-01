package main

import (
	"context"
	"fmt"
	"io"
	"micro/learn1/test10_stream/server/proto"
	"net"

	"google.golang.org/grpc"
)

// 客户端流式
type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {

	reply := "Hello " + in.GetName()
	return &proto.HelloResponse{Reply: reply}, nil
}

func (s *Server) LotsOfGreetings(stream proto.Greeter_LotsOfGreetingsServer) error {

	reply := "Hello, "

	for {
		// 接收客户端发来的流式消息
		req, err := stream.Recv()
		if err == io.EOF {
			// 末尾统一回复，发送响应并关闭流
			return stream.SendAndClose(&proto.HelloResponse{
				Reply: reply,
			})
		}
		if err != nil {
			return err
		}
		reply = reply + req.GetName()
	}
}

func main() {

	// 启动服务
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("net.Listen failed: %v\n", err)
		return
	}

	// 创建 rpc 服务
	s := grpc.NewServer()

	// 注册rpc服务
	proto.RegisterGreeterServer(s, &Server{})

	// 启动rpc服务
	err = s.Serve(listener)
	if err != nil {
		fmt.Printf("s.Serve failed:%v\n", err)
		return
	}
}
