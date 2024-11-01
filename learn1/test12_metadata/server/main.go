package main

import (
	"context"
	"fmt"
	"micro/learn1/test12_metadata/server/proto"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Server struct {
	proto.UnimplementedCareServer
}

func (s *Server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {

	// trailer 必须在响应返回之后发送，这里注册一个延迟函数来创建 trailer
	defer func() {
		trailer := metadata.Pairs(
			"timeStamp", strconv.Itoa(int(time.Now().Unix())),
		)
		grpc.SetTrailer(ctx, trailer)
	}()

	// 客户端发送请求会带上元数据(如果定义了的话)
	// 执行业务之前必须要 check metadata
	md, ok := metadata.FromIncomingContext(ctx)
	fmt.Printf("md:%#v\n", md) // 元数据默认自带一些 kv

	if !ok { // 请求中没有携带元数据就拒绝接受请求
		return nil, status.Error(codes.Unauthenticated, "invalid request")
	}

	// 接收请求，并检查 token
	vl := md.Get("token")
	if len(vl) == 0 || vl[0] != "yuanshuju.rabbit.qianming" {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	reply := " Hello, " + in.GetName()

	// 定义 header
	header := metadata.New(map[string]string{
		"Addr": "TangShan",
	})
	// 发送 header
	grpc.SendHeader(ctx, header)
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

	// 注册rpc服务
	proto.RegisterCareServer(s, &Server{})

	// 启动rpc服务
	err = s.Serve(listener)
	if err != nil {
		fmt.Printf("s.Serve failed:%v\n", err)
		return
	}
}
