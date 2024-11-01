package main

import (
	"context"
	"fmt"
	"micro/learn1/test3/client/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// grpc 客户端
// 调用服务端 【提供好的服务】
func main() {

	// 建立连接
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("grpc.NewClient failed:%v\n", err)
		return
	}
	defer conn.Close()

	// 创建客户端
	client := proto.NewGreeterClient(conn)

	// 调用远程服务
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var name = "chen"
	resp, err := client.SayHello(ctx, &proto.HelloRequest{Name: name})
	if err != nil {
		fmt.Printf("client.SayHello failed:%v\n", err)
		return
	}

	// 获取数据
	respData := resp.GetReply()
	fmt.Printf("respData: %s", respData) // hello chen
}
