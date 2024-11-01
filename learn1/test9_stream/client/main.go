package main

import (
	"context"
	"fmt"
	"io"
	"micro/learn1/test9_stream/server/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 服务器端流式
// 输出：
//  get reply: Hello, chen
//  get reply: Hello, chen
//  get reply: NiHao, chen
//  get reply: Hope you, chen
//  get reply: Thank you, chen

func main() {

	// 建立连接
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("grpc,NewClient failed:%v\n", err)
		return
	}
	defer conn.Close()

	// 创建客户端
	client := proto.NewGreeterClient(conn)

	// 调用远程服务
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	stream, err := client.LotsOfReplies(ctx, &proto.HelloRequest{Name: "chen"})
	if err != nil {
		fmt.Printf("client.LotsOfReplies failed:%v\n", err)
		return
	}

	// 循环读取数据
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("stream.Recv failed:%v\n", err)
			return
		}
		fmt.Printf("get reply: %s\n", resp.Reply)
	}

}
