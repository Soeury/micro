package main

import (
	"context"
	"flag"
	"fmt"
	"micro/learn1/test18_resolver/server/proto"
	"net"

	"google.golang.org/grpc"
)

// gRPC 名称解析 + 负载均衡
//   缺点：服务写死了，服务器挂了之后，每次都需要手动去更改服务  ->  服务注册 + 服务发现

// 服务器端写完了之后，分别在本机 8080 , 8081 , 8082 端口开启服务，之后运行编写好的客户端

// 指针类型的数据 ------ 从命令行获取参数
var port = flag.Int("port", 8080, "server port")

type Server struct {
	proto.UnimplementedHandsServer
	Addr string
}

func (s *Server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {

	reply := "Hello " + in.GetName() + " ," + "serving on localhost:" + s.Addr + "..."
	return &proto.HelloResponse{Reply: reply}, nil
}

func main() {

	flag.Parse()
	addr := fmt.Sprintf("localhost:%d", *port)

	// 启动服务
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("net.listen failed:%v\n", err)
		return
	}

	// 创建 rpc 服务
	s := grpc.NewServer()

	// 注册 rpc 服务
	proto.RegisterHandsServer(s, &Server{Addr: addr})

	// 启动 rpc 服务
	err = s.Serve(listener)
	if err != nil {
		fmt.Printf("s.serve failed:%v\n", err)
		return
	}
}
