package main

import (
	"context"
	"fmt"
	"micro/learn1/test16_gateway/client/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// 建立连接
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("grpc.newClient failed:%v\n", err)
		return
	}
	defer conn.Close()

	// 创建客户端
	client := proto.NewLoveClient(conn)

	// 调用rpc 服务
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*400)
	defer cancel()

	resp, err := client.SayLove(ctx, &proto.LoveRequest{Name: "rabbit"})
	if err != nil {
		fmt.Printf("client.SayLove failed:%v\n", err)
		return
	}
	fmt.Printf("resp:%s\n", resp.GetReply())
}
