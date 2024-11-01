package main

import (
	"context"
	"fmt"
	"micro/learn1/test16_gateway/server/proto"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime" // 注意v2版本
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GRPC-Gateway 能帮助你同时提供 gRPC 和 RESTful 风格的 API
// GRPC-Gateway 是 Google protocol buffers 编译器 protoc 的一个插件
// 它读取 Protobuf 服务定义并生成一个反向代理服务器，该服务器将 RESTful HTTP API 转换为 gRPC
// 该服务器是根据服务定义中的 google.api.http 注释生成的
// 简单来说就是 gateway 能够将 rpc 服务以 http 的方式暴露出来
// 步骤如下：

// -1. 安装 grpc_gateway 工具：
//    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2

// -2. 在 .proto 包中导入包： import "google/api/annotations.proto"

// -3. 在 .proto 文件中编写注释：在 service 中我们需要的 rpc 服务后的 {} 中编写
//    option (google.api.http) = {
//        post: "path"
//        body: "*"
//    }

// -4. 将代码加入到我们生成pb.go 和 grpc.pb.go 的路径上并生成文件
//    --go-gateway_out=. --go-gateway_opt=paths=source_relative

// -5. 在 server.go 中添加代码，并将 s.Serve(listener) 放入一个 goroutine 中去执行，防止堵塞

// ---------- 总结 ----------------
// 1. 编写注释
// 2. 生成代码
// 3. server 端编写包装代码 ： 5 个步骤

// ----------- 注意 ---------------
// 1. server端导入的包："github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
// 2. s.Serve(listener) 需要单独开启 goroutine 执行

type Server struct {
	proto.UnimplementedLoveServer
}

func (s *Server) SayLove(ctx context.Context, in *proto.LoveRequest) (*proto.LoveResponse, error) {

	reply := "I Love you, " + in.GetName()
	return &proto.LoveResponse{Reply: reply}, nil
}

func main() {

	// 启动服务
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("net.Listen failed:%v\n", err)
		return
	}

	// 创建 rpc 服务
	s := grpc.NewServer()

	// 注册 rpc 服务
	proto.RegisterLoveServer(s, &Server{})

	// 启动 rpc 服务 这里使用 goroutine 防止阻塞
	go func() {
		err = s.Serve(listener)
		if err != nil {
			fmt.Printf("s.serve failed:%v\n", err)
			return
		}
	}()

	// 1. 创建一个连接 rpc 服务的客户端连接
	// grpc-gateway 就是通过这个来代理请求
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("failed to connect server")
		return
	}

	// 2. 创建一个新的 http 请求多路复用器
	gwmux := runtime.NewServeMux()

	// 3. 在多路复用器上处理定义好的服务
	err = proto.RegisterLoveHandler(context.Background(), gwmux, conn) // 注意这里是 handler 不是 server
	if err != nil {
		fmt.Printf("failed to register gateway")
		return
	}

	// 4. 创建一个新的服务： 在 8090 端口提供 gateway 服务
	gwServer := &http.Server{
		Addr:    "localhost:8090",
		Handler: gwmux,
	}
	fmt.Println("serving grpc-gateway on 'localhost:8090'")

	// 5. 启动服务
	err = gwServer.ListenAndServe()
	if err != nil {
		fmt.Printf("gwserver failed:%v\n", err)
		return
	}

	// 之后就可以使用定义好的 url 路径使用 postman 发送请求来检查服务器端代码的编写了
}
