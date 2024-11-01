package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"micro/learn2/test2_register/server/consul"
	"micro/learn2/test2_register/server/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type Server struct {
	proto.UnimplementedGreeterServer
	Addr string
}

var port = 8080 // 端口默认值

func (s *Server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {

	reply := "Hello " + in.GetName() + " , serving on [" + s.Addr + "]..."
	return &proto.HelloResponse{Reply: reply}, nil
}

func main() {

	flag.IntVar(&port, "port", 8080, "server port")
	flag.Parse() // 解析命令行参数
	fmt.Printf("-->port:%d\n", port)

	// 1. 启动服务
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "0.0.0.0", port))
	if err != nil {
		fmt.Printf("net.Listen failed:%v\n", err)
		return
	}

	// 2. 创建 RPC 服务
	s := grpc.NewServer()

	uip, err := consul.GetOutIP() // 获取本机地址
	if err != nil {
		fmt.Printf("consul.GetOutIP failed:%v\n", err)
		return
	}
	addr := fmt.Sprintf("%s:%d", uip.String(), port)

	// 3. 注册 RPC 服务
	proto.RegisterGreeterServer(s, &Server{Addr: addr})

	// 4. 注册健康检查服务(用来在连接consul时进行健康检查)
	healthpb.RegisterHealthServer(s, health.NewServer())

	// 5. 调用一些定义好的 consul 操作:  连接 + 注册服务 + 健康检查

	// ------ 连接 ------
	cc, err := consul.NewConsul()
	if err != nil {
		fmt.Printf("consul.newConsul failed:%v\n", err)
	}

	// ------ 注册 ------
	err = cc.RegisterService("learn2-test2", uip.String(), port) // 注意这里的 net.IP 和 string 的转换
	serviceID := fmt.Sprintf("%s-%s-%d", "learn2-test2", uip.String(), port)
	if err != nil {
		fmt.Printf("cc.registerService failed:%v\n", err)
		return
	}

	// 6. 启动RPC服务
	go func() {
		err = s.Serve(listener)
		if err != nil {
			fmt.Printf("s.serve failed:%v\n", err)
			return
		}
	}()

	// 7.服务注销(可选)
	//    -1. 手动服务注销:  通道阻塞 + 信号 + gouroutine启动RPC服务
	//    -2. 健康检查不通过注销服务
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	fmt.Printf("waiting quit server...\n")
	<-quit

	// 服务结束后自动手动注销服务
	err = cc.Deregister(serviceID)
	fmt.Printf("Server Deregister seccess!")
	if err != nil {
		fmt.Printf("service Deregister failed:%v\n", err)
		return
	}
}
