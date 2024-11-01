package main

import (
	"context"
	"fmt"
	"time"

	"micro/learn2/test2_register/client/proto"

	_ "github.com/mbobakov/grpc-consul-resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	/*
		// 一些连接consul获取服务address的操作
		// 第一种方式： 手动编写代码
		addr, err := consul.GetConsulAddr()
		if err != nil {
			fmt.Printf("%s\n", addr)
			return
		}

		// 1. 建立连接
		conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Printf("grpc.NewClient failed:%v\n", err)
			return
		}
		defer conn.Close()

	*/

	// 服务发现第二种方式: 匿名导入包 + 按照特定格式编写地址
	conn, err := grpc.NewClient(
		"consul://localhost:8500/learn2-test2?healthy=true",
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy" : "round_robin"}`), // 负载均衡策略：轮询
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Printf("grpc.NewClient failed:%v\n", err)
		return
	}

	// 2. 创建客户端
	client := proto.NewGreeterClient(conn)

	// 3. 调用远程服务
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour*2)
	defer cancel()

	// 模拟10次请求
	for i := 0; i < 10; i++ {
		resp, err := client.SayHello(ctx, &proto.HelloRequest{Name: "rabbit"})
		if err != nil {
			fmt.Printf("client.sayHello failed:%v\n", err)
			return
		}
		fmt.Printf("resp:%s\n", resp.GetReply())
	}
}
