package main

import (
	"context"
	"fmt"
	"io"
	"micro/learn1/test11_stream/server/proto"
	"net"
	"strings"

	"google.golang.org/grpc"
)

// 双向流式
//
//	应用场景：对话聊天
type Server struct {
	proto.UnimplementedGreeterServer
}

// 以下是一段魔法
// 这里更改了服务器端需要重新连接
func Magic(s string) string {

	s = strings.ReplaceAll(s, "吗", "")
	s = strings.ReplaceAll(s, "吧", "")
	s = strings.ReplaceAll(s, "你", "我")
	s = strings.ReplaceAll(s, "?", "!")
	s = strings.ReplaceAll(s, "哪儿", "这儿")
	s = strings.ReplaceAll(s, "要", "可")
	return s
}

func (s *Server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {

	reply := "Hello " + in.GetName()
	return &proto.HelloResponse{Reply: reply}, nil
}

func (s *Server) BiYing(stream proto.Greeter_BiYingServer) error {

	for {
		data, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		// 服务器端发送流
		reply := Magic(data.GetName())
		if err = stream.Send(&proto.HelloResponse{Reply: reply}); err != nil {
			return err
		}
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
