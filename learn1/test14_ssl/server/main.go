package main

import (
	"context"
	"fmt"
	"micro/learn1/test14_ssl/server/proto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// rpc 加密认证： 3个步骤：   生成 .key 文件    编写 .conf 文件    生成 .crt 文件
//  生成文件注意切换路径
//  -1. openssl ecparam -genkey -name secp384r1 -out server.key 生成 server.key 文件
//  -2. 编写 server.conf 文件
//  -3. openssl req -nodes -new -x509 -sha256 -days 3650 -config server.conf -extensions 'req_ext' -key server.key -out server.crt

// 使用方法：
//   server :
//     生成证书 ：              creds, err := credentials.NewServerTLSFromFile(path1 , path2)
//     加入到创建的rpc服务中 ：  s := grpc.NewServer(grpc.Creds(creds))

type Server struct {
	proto.UnimplementedFriendServer
}

func (s *Server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {

	reply := "Hello, my name is " + in.GetName()
	return &proto.HelloResponse{Reply: reply}, nil
}

func main() {

	// 启动服务
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("net.Listen failed:%v\n", err)
		return
	}

	// 加载证书
	creds, err := credentials.NewServerTLSFromFile("./certs/server.crt", "./certs/server.key")
	if err != nil {
		fmt.Printf("credentials.NewServerTLSFromFile failed:%v\n", err)
		return
	}

	// 创建rpc服务 + 辅助证书
	s := grpc.NewServer(grpc.Creds(creds))

	// 注册rpc服务
	proto.RegisterFriendServer(s, &Server{})

	// 启动rpc服务
	err = s.Serve(listener)
	if err != nil {
		fmt.Printf("s.Serve failed:%v\n", err)
		return
	}
}
