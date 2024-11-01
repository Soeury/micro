package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// 使用 go 内置的 net/rpc 包
// 注册 rpc 需要解决的三个问题:
//   -1. 注册服务对象
//   -2. 实现数据参数和返回值的处理
//   -3. 确定传输协议和数据格式

// rpc 也支持 tcp 协议而不使用 HTTP 协议
// rpc 默认使用 gob 协议对数据进行传输，局限性比较高，也可以改用 json 协议传输数据

type Args struct {
	X int
	Y int
}

// 此时，客户端想要调用服务器端的某个服务，就需要先连接到server端，再执行调用
func main() {

	// 建立HTTP连接
	client, err := rpc.DialHTTP("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	/*

	   // 基于TCP协议的连接
	   client , err := rpc.Dial("tcp" , "localhost:8080")
	   if err != nil {
	       log.Fatal(err)
	   }

	*/

	/*

		// 基于TCP协议并且使用JSON传输数据的RPC
		conn , err := net.Dial("tcp" , "localhost:8080")
		if err := nil {
			log.Fatal(err)
		}

		client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	*/

	// 同步调用: 使用变量 client.Call()
	args := &Args{10, 20}
	var reply int
	err = client.Call("ServiceA.Add", args, &reply)
	if err != nil {
		log.Fatal("ServiceA.Add:", err)
	}
	fmt.Printf("ServiceA.Add: %d+%d=%d\n", args.X, args.Y, reply)

	// 异步调用: 使用通道 client.Go()
	args = &Args{10, 20}
	var reply2 int
	divCall := client.Go("ServiceA.Add", args, &reply2, nil)
	replyCall := <-divCall.Done // 接收调用的结果
	fmt.Println(replyCall.Error)
	fmt.Println(reply2)
}
