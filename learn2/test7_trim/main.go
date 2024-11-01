package main

import (
	"context"
	"flag"
	"fmt"
	"micro/learn2/test7_trim/consul"
	"micro/learn2/test7_trim/trim"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// 这个服务作为 test6_summary 中的第三方服务
type Server struct {
	trim.UnimplementedTrimServer
	Addr string
}

var port = 8082

func (s *Server) TrimSpace(ctx context.Context, in *trim.TrimRequest) (*trim.TrimResponse, error) {

	str := in.GetWord1()
	ret := strings.ReplaceAll(str, " ", "")
	fmt.Printf("old str:%s , new str:%s\n", str, ret)
	return &trim.TrimResponse{Ret1: ret}, nil
}

func main() {

	flag.IntVar(&port, "port", 8082, "server port")
	flag.Parse()
	fmt.Printf("-->port:%d\n", port)

	// 1. 启动服务
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "0.0.0.0", port))
	if err != nil {
		fmt.Printf("net.Listen failed:%v\n", err)
		return
	}
	defer listener.Close()

	// 2. 创建 RPC 服务
	s := grpc.NewServer()

	// 获取本地 ip
	uip, err := consul.GetLocalIP()
	if err != nil {
		fmt.Printf("consul.GetLocalIP failed:%v\n", err)
		return
	}
	addr := fmt.Sprintf("%s:%d", uip.String(), port)

	// 3. 注册 RPC 服务
	trim.RegisterTrimServer(s, &Server{Addr: addr})

	// 4. 注册健康检查服务
	healthpb.RegisterHealthServer(s, health.NewServer())

	// 5. 调用一些定义好的 consul 操作:  连接 + 注册服务 + 健康检查

	// ------ 连接 ------
	cc, err := consul.NewConsul()
	if err != nil {
		fmt.Printf("consul.NewConsul failed:%v\n", err)
		return
	}

	// ------ 注册 ------
	err = cc.RegisterService("learn2_test7_trim", uip.String(), port)
	serviceID := fmt.Sprintf("%s-%s-%d", "learn2_test7_trim", uip.String(), port)
	if err != nil {
		fmt.Printf("cc.RegisterService failed:%s\n", err)
		return
	}

	// 6. 启动 RPC 服务
	go func() {
		err = s.Serve(listener)
		if err != nil {
			fmt.Printf("s.serve failed:%v\n", err)
			return
		}
	}()

	// 7.实现服务退出后自动注销
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("waiting server quit...\n")
	<-quit

	err = cc.Deregister(serviceID)
	fmt.Printf("service deregister success!")
	if err != nil {
		fmt.Printf("service deregidter failed:%v\n", err)
		return
	}
}
