package main

import (
	"context"
	"fmt"
	"micro/learn1/test14_ssl/client/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// rpc 加密认证
// 使用方法：
//    client :
//       -1. 加载证书：      creds, err := credentials.NewClientTLSFromFile(path , url)
//       -2. 服务中加入证书： conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(creds))

func main() {

	// 加载证书
	path := "./learn1/test14_ssl/client/certs/server.crt"
	creds, err := credentials.NewClientTLSFromFile(path, "kedudu.com")
	if err != nil {
		fmt.Printf("credentials.NewClientTLSFromFile failed:%v\n", err)
		return
	}

	// 建立连接
	url := "localhost:8080"
	conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Printf("grpc.NewClient failed:%v\n", err)
		return
	}
	defer conn.Close()

	// 创建客户端
	client := proto.NewFriendClient(conn)

	// 调用rpc服务
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*200)
	defer cancel()

	resp, err := client.SayHello(ctx, &proto.HelloRequest{Name: "Rabbit"})
	if err != nil {
		fmt.Printf("client.SayHello failed:%v\n", err)
		return
	}

	fmt.Printf("resp:%v\n", resp.GetReply())
}
