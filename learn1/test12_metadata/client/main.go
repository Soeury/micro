package main

import (
	"context"
	"fmt"
	"micro/learn1/test12_metadata/client/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// metadata:
// GRPC 中使用 metadata 进行 client-server 数据传输
//  metadata 中存放一些和业务代码无关的内容
//  用户可以指定存放在metadata中的内容，比如一些认证的数据
//  如果客户端不指定内容的话，metadata 也会默认存在一些数据

// server 也可以发送元数据给 client，通过 header 和 trailer
// client 想要接收 server 返回的 header 和 trailer 必须发起rpc调用之前定义好这两个变量

func main() {

	// 建立连接
	addr := "localhost:8080"
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("grpc.NewClient failed:%v\n", err)
		return
	}
	defer conn.Close()

	// 建立客户端
	client := proto.NewCareClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*200)
	defer cancel()

	// 普通调用rpc服务
	// 在这之前可以带上元数据: 注意这里的写法!
	md := metadata.Pairs(
		"token", "yuanshuju.rabbit.qianming",
	)

	// 将元数据加入到 context 中
	ctx = metadata.NewOutgoingContext(ctx, md)
	// 这里 header , trailer 需要在调用之前定义好定义好，并且添加到 rpc 调用中
	var header metadata.MD
	var trailer metadata.MD
	resp, err := client.SayHello(
		ctx,
		&proto.HelloRequest{Name: "chen"},
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		fmt.Printf("client.SayHello failed:%v\n", err)
		return
	}

	// 调用了服务之后就可以拿到 header , trailer
	fmt.Printf("resp:%s\n", resp.GetReply())
	fmt.Printf("header: %+v\n", header)
	fmt.Printf("trailer: %+v\n", trailer) // 获取 trailer
}
