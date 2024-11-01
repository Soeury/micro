package main

import (
	"context"
	"fmt"
	"micro/learn1/test10_stream/client/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 客户端流式
// 应用场景：物联网终端向服务器上报数据

func main() {

	// 建立连接
	addr := "localhost:8080"
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("grpc.NewClient failed:%v\n", err)
		return
	}

	// 创建客户端
	client := proto.NewGreeterClient(conn)

	// 调用远程服务
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// 客户端向服务器发送流
	stream, err := client.LotsOfGreetings(ctx)
	if err != nil {
		fmt.Printf("client.LotsOfGretings failed:%v\n", err)
		return
	}

	names := []string{"chen, ", "boliang, ", "rabbit, ", "kaka, "}
	for _, name := range names {
		err := stream.Send(&proto.HelloRequest{Name: name})
		if err != nil {
			fmt.Printf("stream.Send failed:%v\n", err)
			return
		}
	}

	// 关闭流时，获取响应
	resp, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Printf("stream.CloseAndRecv failed:%v\n", err)
		return
	}
	fmt.Printf("response:%s\n", resp.GetReply())
}
