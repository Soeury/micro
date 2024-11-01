package main

import (
	"context"
	"fmt"
	"micro/learn1/test9_stream/server/proto"
	"net"

	"google.golang.org/grpc"
)

// 服务器端流式实现
//
//		-1. 编写 protobuffer 文件： rpc name1 (name2) returns (stream name3);
//		-2. 生成文件
//		-3. 写业务代码： stream流数据通过参数传递给指定方法，返回值只有error，需要使用 stream.Send(data)发送数据
//
//	 应用场景：股票分析
type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {

	reply := "Hello " + in.GetName()
	return &proto.HelloResponse{Reply: reply}, nil
}

func (s *Server) LotsOfReplies(in *proto.HelloRequest, stream proto.Greeter_LotsOfRepliesServer) error {

	words := []string{
		"Hello, ",
		"NiHao, ",
		"Hope you, ",
		"Thank you, ",
	}

	// 流式发送
	for _, word := range words {
		reply := &proto.HelloResponse{
			Reply: word + in.GetName(),
		}

		err := stream.Send(reply)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {

	// 启动服务
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("net.Listen failed:%v\n", err)
		return
	}

	// 创建rpc服务
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
