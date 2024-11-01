package main

import (
	"context"
	"fmt"
	"micro/learn1/test13_error/client/proto"
	"time"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

// grpc 错误处理
//  示例：client 向 server 传递数据实现 server 端只调用一次数据，调用次数超过一次返回指定错误
//  解决：server 在服务对象中注册 map[string]int 计数器，记录数据调用次数
//    注意 map 的使用事项:  -1. 必须先初始化才能使用(make)   -2. 并发不安全，解决办法: (Mutex, sync.Map)

//  打印出来是这样的: 不能理解第二条.....
// quota failure:violations:{subject:"name:chen"  description:"limit every name just can use once!"}
// client.SayHello failed:rpc error: code = ResourceExhausted desc = request limit exceeded

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

	// 调用 rpc 服务
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*200)
	defer cancel()

	resp, err := client.SayHello(ctx, &proto.HelloRequest{Name: "chen"})
	if err != nil {
		// 将自定义错误转化为 *status.status 类型
		st := status.Convert(err)

		// 获取 details
		for _, d := range st.Details() {
			switch info := d.(type) {
			case *errdetails.QuotaFailure:
				fmt.Printf("quota failure:%v\n", info)
			default:
				fmt.Printf("unexpected type:%v\n", info)
			}
		}

		fmt.Printf("client.SayHello failed:%v\n", err) // 这里是怎么拿到 st.Err() 的数据的 ?
		return
	}

	fmt.Printf("resp:%s\n", resp.GetReply())
}
