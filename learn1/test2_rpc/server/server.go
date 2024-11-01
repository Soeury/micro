package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	X int
	Y int
}

// 自定义一个结构类型
type ServiceA struct{}

// Add 为自定义的结构体对象 ServiceA 定义一个可导出的 Add 方法
// 思考: 所有的返回值都为error , 我们真正需要返回的值应该定义为参数，用指针的形式取出值
func (s *ServiceA) Add(args *Args, reply *int) error {
	*reply = args.X + args.Y
	return nil
}

// main 函数为将上面自定义的结构体类型注册为一个服务
// 之后就可以使用 rpc 调用了
func main() {

	service := new(ServiceA) // new一个对象
	rpc.Register(service)    // 为对象注册rpc服务
	rpc.HandleHTTP()         // 基于HTTP协议

	listener, err := net.Listen("tcp", ":8080") // 创建监听器
	if err != nil {
		log.Fatal(err)
	}
	http.Serve(listener, nil) // 启动服务器

	/*

		// 基于TCP协议的RPC
		service := new(ServiceA)
		rpc.Register(service)
		listener , err := net.Listen("tcp" , ":8080")

		if err != nil {
			log.Fatal(err)
		}

		for {
			conn , _ := listener.Accept()
			rpc.ServeConn(conn)
		}

	*/

	/*

		// 基于TCP协议并且使用JSON传输数据的RPC
		service := new(ServiceA)
		rpc.Register(service)
		listener , err := net.Listen("tcp" , ":8080")

		if err != nil {
			log.Fatal(err)
		}

		for {
			conn , _ := listener.Accept()
			rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
		}

	*/
}
