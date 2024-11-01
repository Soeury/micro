package main

import (
	"context"
	"fmt"
	"micro/learn1/test15_unary/client/proto"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

// 补充：当一个结构体 (struc) 实现了接口中声明的所有方法时，我们就说这个结构体实现了该接口

//  客户端拦截器： 普通拦截器 + 流拦截器 (中间件，在注册证书时一起注册，类似于 gin 中的中间件)
//  功能：日志记录、身份验证/授权、指标收集以及许多其他可以跨 RPC 共享的功能

// 客户端普通拦截器实现：UnaryClientInterceptor
// func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error

// 重写
type PerRPCCredentials interface {
	GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
	RequireTransportSecurity() bool
}

type OAuth2Credentials struct {
	tokenSource oauth2.TokenSource
}

func NewOAuth2Credentials(tokenSource oauth2.TokenSource) credentials.PerRPCCredentials {
	return &OAuth2Credentials{tokenSource: tokenSource}
}

// GetRequestMetadata 实现 grpc.PerRPCCredentials 接口
func (o *OAuth2Credentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	token, err := o.tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("get OAuth2 token failed: %v", err)
	}
	return map[string]string{
		"authorization": token.Type() + " " + token.AccessToken,
	}, nil
}

// RequireTransportSecurity 实现 grpc.PerRPCCredentials 接口
func (o *OAuth2Credentials) RequireTransportSecurity() bool {
	return true
}

// 定义客户端普通拦截器 实现未传入 token 时补充加入
func unaryClientInterceptor(
	ctx context.Context,
	method string,
	req, reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {

	// 检查调用选项 opts 中是否已包含认证信息
	var credentials bool
	for _, o := range opts {
		if _, ok := o.(grpc.PerRPCCredsCallOption); ok {
			credentials = true
			break
		}
	}

	// 创建一个新的 OAuth2 认证令牌源
	if !credentials {
		config := &oauth2.Config{}
		tokenSource := config.TokenSource(ctx, &oauth2.Token{
			AccessToken: "yuanshuju.rabbit.qianming",
		})
		opts = append(opts, grpc.PerRPCCredentials(NewOAuth2Credentials(tokenSource)))
	}

	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	end := time.Now() // 实际执行 rpc 调用
	fmt.Printf("RPC: %s, cost time: %s, err: %v\n", method, end.Sub(start), err)
	return err
}

func main() {

	// 加载证书 + 注意切换到当前路径下
	path := "./client/certs/server.crt"
	creds, err := credentials.NewClientTLSFromFile(path, "kedudu.com")
	if err != nil {
		fmt.Printf("credentials.NewClientTLSFromFile failed:%v\n", err)
		return
	}

	// 建立连接 + 注册拦截器 + 注册ssl证书
	url := "localhost:8080"
	conn, err := grpc.NewClient(
		url,
		grpc.WithTransportCredentials(creds),
		grpc.WithUnaryInterceptor(unaryClientInterceptor),
	)
	if err != nil {
		fmt.Printf("grpc.NewClient failed:%v\n", err)
		return
	}
	defer conn.Close()

	// 创建客户端
	client := proto.NewFriendClient(conn)

	// 调用rpc服务
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*200)
	defer cancel()

	// 定义需要传送的元数据
	md := metadata.Pairs("token", "yuanshuju.rabbit.qianming")
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := client.SayHello(ctx, &proto.HelloRequest{Name: "Rabbit"})
	if err != nil {
		fmt.Printf("client.SayHello failed:%v\n", err)
		return
	}

	fmt.Printf("resp:%v\n", resp.GetReply())
}
