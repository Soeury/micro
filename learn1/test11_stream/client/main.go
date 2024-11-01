package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"micro/learn1/test11_stream/client/proto"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// 建立连接
	addr := "localhost:8080"
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("grpc.NewClient failed:%v\n", err)
		return
	}

	// 创建客户端
	client := proto.NewGreeterClient(conn)

	// 调用远程服务
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*200)
	defer cancel()

	// 双向流模式:
	stream, err := client.BiYing(ctx)
	if err != nil {
		fmt.Printf("client.biying failed: %v\n", err)
		return
	}

	// 定义一个结构体类型的通道
	waitc := make(chan struct{})

	// 协程接收服务端返回的响应
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				fmt.Printf("stream.Recv failed:%v\n", err)
				return
			}

			fmt.Printf("AI: %s\n", resp.GetReply())
		}
	}()

	// 从标准输入生成读对象，缓冲读取，
	reader := bufio.NewReader(os.Stdin)
	for {
		// 读取用户在终端输入的字符串
		cmd, _ := reader.ReadString('\n')

		// trimSpace 函数去除 cmd 字符串前后的空白字符
		// 确保即使用户输入了额外的空格或按了多次回车键，程序也能正确处理命令
		cmd = strings.TrimSpace(cmd)
		if len(cmd) == 0 {
			continue
		}

		// 这里表示 quit 不区分大小写
		if strings.ToUpper(cmd) == "QUIT" {
			break
		}

		// 将获取到的数据发送至服务端
		if err := stream.Send(&proto.HelloRequest{Name: cmd}); err != nil {
			fmt.Printf("c.BidiHello stream.Send(%v) failed: %v", cmd, err)
		}
	}
	stream.CloseSend()
	<-waitc
}
