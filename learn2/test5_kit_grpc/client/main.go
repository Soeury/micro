package main

import (
	"context"
	"fmt"
	"micro/learn2/test5_kit_grpc/client/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// 连接
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("grpc.NewClient failed:%v\n", err)
		return
	}
	defer conn.Close()

	// 创建客户端
	client := proto.NewAddClient(conn)

	// 调用 rpc 服务
	// sum
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour*100)
	defer cancel()

	resp, err := client.Sum(ctx, &proto.SumRequest{A: 10, B: 2})
	if err != nil {
		fmt.Printf("client.Sum failed:%v\n", err)
		return
	}

	fmt.Printf("resp1: %d\n", resp.GetValue())

	// append
	resp2, err := client.Append(ctx, &proto.AppendRequest{A: "name:", B: "rabbit"})
	if err != nil {
		fmt.Printf("client.Append failed:%v\n", err)
		return
	}

	fmt.Printf("resp2: %s\n", resp2.GetValue())
}
