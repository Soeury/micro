package main

import (
	"context"
	"fmt"
	store "micro/learn1/test17/bookstore"
	data "micro/learn1/test17/data"
	"micro/learn1/test17/proto"
	"net"
	"net/http"
	"strings"

	// 注意 v2 版本
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c" // h2c 适用于需要将 gRPC 服务和 HTTP 服务部署在同一个端口上的场景
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	db, err := data.NewDB()
	if err != nil {
		fmt.Printf("newDB fiailed:%v\n", err)
		return
	}

	// 将 db 注入到服务对象中
	srv := store.Server{
		BS: &data.Bookstore{DB: db},
	}

	// 启动服务
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("net.Listen failed:%v\n", err)
		return
	}

	// 创建rpc 服务
	s := grpc.NewServer()

	// 注册 rpc 服务
	proto.RegisterBookstoreServer(s, &srv)

	// 这里是实现在不同端口处理 grpc 和 http 请求

	// 启动rpc服务
	go func() {
		err := s.Serve(listener)
		if err != nil {
			fmt.Printf("s.serve failed:%v\n", err)
			return
		}
	}()

	// grpc-gateway
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("grpc.newclient failed:%v\n", err)
		return
	}

	gwmux := runtime.NewServeMux()
	err = proto.RegisterBookstoreHandler(context.Background(), gwmux, conn) // 注意这里是 handler 不是 server
	if err != nil {
		fmt.Printf("proto.registerhandler failed:%v\n", err)
		return
	}

	gwserver := &http.Server{
		Addr:    "localhost:8090",
		Handler: gwmux,
	}
	fmt.Printf("gateway serve on 'localhost:8090'")

	err = gwserver.ListenAndServe()
	if err != nil {
		fmt.Printf("listenansserve failed:%v\n", err)
		return
	}

	// 下面的代码实现不了？？？

	/*

		// 实现在同一个端口同时进行处理 rpc 和 http ------- 四个步骤
		// 这里因为没有使用 TLS 加密，所以选择使用 h2c 这个包
		// 需要使用到的两个包:  	"golang.org/x/net/http2"     "golang.org/x/net/http2/h2c"

		// 1. 创建 gate-way mux，并将 gRPC 服务注册到 gwmux
		// runtime.NewServeMux() 是 gRPC-Gateway 库中用于将 gRPC 服务暴露为 RESTful API 的函数
		gwmux := runtime.NewServeMux()
		dops := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
		if err := proto.RegisterBookstoreHandlerFromEndpoint(
			context.Background(), gwmux, "localhost:8090", dops); err != nil {
			fmt.Printf("proto.registerHandlerFromEndpoint failed:%v\n", err)
			return
		}

		// 2. 创建 httpMux
		// 所有 HTTP 请求都会首先被 httpMux 接收，然后根据路径分发到 gwmux 或其他的处理器
		httpMux := http.NewServeMux()
		httpMux.Handle("/", gwmux)

		// 3. 定义 server
		newServer := &http.Server{
			Addr:    "localhost:8090",
			Handler: GrpcHandleFunc(s, httpMux),
		}

		// 4. 启动 server
		err = newServer.Serve(listener)
		if err != nil {
			fmt.Printf("newServer.Serve failed:%v\n", err)
			return
		}

	*/
}

// GrpcHandleFunc 将 rpc 和 http 请求分别调用不同的 handler 处理
// 翻译一下：为了区分 grpc 请求和 http 请求
func GrpcHandleFunc(grpcServer *grpc.Server, otherHandler *http.ServeMux) http.Handler {

	// h2c.NewHandler 来同时支持并识别 HTTP/1.1 和 HTTP/2 (无TLS)
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}
