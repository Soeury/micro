package main

import (
	"context"
	"flag"
	"fmt"
	"micro/learn2/test8_multi/consul"
	"micro/learn2/test8_multi/multi"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type Server struct {
	multi.UnimplementedMultiServer
	Addr string
}

var port = 8083

// Multiply 实现传入的两个数分别翻倍后相加
func (s *Server) Multiply(ctx context.Context, in *multi.MultiRequest) (*multi.MultiResponse, error) {

	res1 := in.GetFigure1() * 2
	res2 := in.GetFigure2() * 2
	fmt.Printf("old data1:%d , new data1:%d\n", in.GetFigure1(), res1)
	fmt.Printf("old data2:%d , new data2:%d\n", in.GetFigure2(), res2)
	return &multi.MultiResponse{Res1: res1, Res2: res2}, nil
}

func main() {

	flag.IntVar(&port, "port", 8083, "server port")
	flag.Parse()
	fmt.Printf("-->port:%d\n", port)

	// 启动服务 , 这里记得改 addr
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "0.0.0.0", port))
	if err != nil {
		fmt.Printf("net.Listen failed:%v\n", err)
		return
	}
	defer listener.Close()

	// 创建 RPC 服务
	s := grpc.NewServer()

	// 获取本地ip
	uip, err := consul.GetLocalAddr()
	if err != nil {
		fmt.Printf("consul.GetLocalAddr faile:%s\n", err)
		return
	}
	addr := fmt.Sprintf("%s-%d", uip.String(), port)
	serviceID := fmt.Sprintf("%s-%s-%d", "learn2_test8_multi", uip.String(), port)

	// 注册 RPC 服务
	multi.RegisterMultiServer(s, &Server{Addr: addr})

	// 注册健康检查服务
	healthpb.RegisterHealthServer(s, health.NewServer())

	// consul 相关操作： 连接 + 注册服务 + 健康检查 + 退出撤销

	// ---- 连接 ----
	cc, err := consul.ConnectToConsul()
	if err != nil {
		fmt.Printf("consul.ConnectToConsul failed:%v\n", err)
		return
	}

	// ---- 注册 ----
	err = cc.RegisterService("learn2_test8-multi", uip.String(), port)
	if err != nil {
		fmt.Printf("consul.RegisterService failed:%s\n", err)
		return
	}

	// 启动 rpc 服务
	go func() {
		err = s.Serve(listener)
		if err != nil {
			fmt.Printf("s.serve failed:%v\n", err)
			return
		}
	}()

	// 退出服务后撤销 consul 的注册
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("wainting serice quit...\n")
	<-quit

	err = cc.DeregisterService(serviceID)
	if err != nil {
		fmt.Printf("ccDeregisterService failed:%v\n", err)
		return
	}
	fmt.Printf("service deregister success!")
}
