package main

import (
	"context"
	"fmt"
	"micro/learn1/test18_resolver/client/proto"
	create "micro/learn1/test18_resolver/create_resolver"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 名称解析器， name resolver ，map[name][]string{ips} , 接收后端服务的名称，返回对应的 ip 列表
// 名称解析定义好之后就可以进行负载均衡了
//    gRPC-go 内置负载均衡支持有:
//        -1. pick_first (默认值)  :  pick_first 会尝试连接取到的第一个服务端地址，如果连接成功，则将其用于所有 RPC
//        -2. round_robin 两种策略 :  循环依次向每个 server 发送一个 RPC

// 使用轮询作为负载均衡策略：
//    在建立连接时加上 : grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy" : "round_robin"}`)

func main() {

	// 建立连接 (所有请求全部打在  127.0.0.1:8080 上)
	// 注意这里： grpc.Dial() 已经废弃使用

	/*
		// -1. DNS 域名解析
		conn, err := grpc.NewClient("dns:///localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Printf("grpc.Newclient failed:%v\n", err)
			return
		}
	*/

	/*
		// -2. 使用 consul 作为注册中心，导入包：_ "github.com/mbobakov/grpc-consul-resolver"
		conn, err := grpc.NewClient(
			"consul://192.168.1.11:8500/hello?wait=14s",
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			fmt.Printf("grpc.Newclient failed:%v\n", err)
			return
		}
	*/

	// -3. 自定义解析器(这里注意服务器端打开不同的端口...)
	conn, err := grpc.NewClient(
		"rabbit:///resolver.kedudu.com",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy" : "round_robin"}`),
		grpc.WithResolvers(&create.RabbitResolverBuilder{}),
	)
	if err != nil {
		fmt.Printf("grpc.Newclient failed:%v\n", err)
		return
	}

	// 创建客户端
	client := proto.NewHandsClient(conn)

	// 调用远程服务
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour*240)
	defer cancel()
	reqData := &proto.HelloRequest{Name: "rabbit"}

	// 模拟多次请求
	for i := 0; i < 10; i++ {
		resp, err := client.SayHello(ctx, reqData)
		if err != nil {
			fmt.Printf("client.Sayhello failed:%v\n", err)
			return
		}
		fmt.Printf("resp:%s\n", resp.GetReply())
	}
}
