package main

import (
	"context"
	"fmt"
	"micro/learn1/test13_error/server/proto"
	"net"
	"sync"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	proto.UnimplementedCareServer
	mu    sync.Mutex
	count map[string]int
}

func (s *Server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {

	// map 并发不安全，这里先上锁，再解锁，也可以使用特殊类型的 sync.Map 来进行，但是使用条件会比较多
	s.mu.Lock()
	defer s.mu.Unlock()

	s.count[in.GetName()] += 2 // 记录数据调用次数
	if s.count[in.GetName()] > 1 {
		st := status.New(codes.ResourceExhausted, "request limit exceeded")

		// 这里的步骤是: 注册更详细的错误信息
		dt, err := st.WithDetails(
			&errdetails.QuotaFailure{
				Violations: []*errdetails.QuotaFailure_Violation{{
					Subject:     fmt.Sprintf("name:%s", in.Name),
					Description: "limit every name just can use once!",
				}},
			},
		)

		// dt , err := st.WithDetails failed, can just return st.Err()
		if err != nil {
			return nil, st.Err()
		}
		return nil, dt.Err()
	}

	reply := "Hello, " + in.GetName()
	return &proto.HelloResponse{Reply: reply}, nil
}

func main() {
	// 创建服务
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("net.Listen failed:%v\n", err)
		return
	}

	// 创建rpc服务
	s := grpc.NewServer()

	// 注册rpc服务，这里需要初始化 count
	proto.RegisterCareServer(s, &Server{count: make(map[string]int)})

	// 启动rpc服务
	err = s.Serve(listener)
	if err != nil {
		fmt.Printf("s.Serve failed:%v\n", err)
		return
	}
}
