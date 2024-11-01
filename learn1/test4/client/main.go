package main

import (
	"context"
	"fmt"
	"micro/learn1/test4/client/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// 建立连接
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("grpc.NewClient failed: %v\n", err)
		return
	}
	defer conn.Close()

	// 创建客户端
	client := proto.NewGreeterClient(conn)

	// 调用远程服务
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.Add(ctx, &proto.AddRequest{Num1: 10, Num2: 20})
	if err != nil {
		fmt.Printf("client.Add failed: %v\n ", err)
		return
	}

	// 获取响应数据
	respData := resp.GetRet()
	fmt.Printf("respData: %d\n", respData) // 30

}
