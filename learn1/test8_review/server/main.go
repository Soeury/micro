package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net"

	proto "micro/learn1/test8_review/server/proto" // 注意导入包

	"github.com/iancoleman/strcase"
	fieldmask_utils "github.com/mennanov/fieldmask-utils"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/anypb"
)

// 基于 oneof , wrappers , field_mask 的一个小练习
type Server struct {
	proto.UnimplementedSaleServer
}

func (s *Server) Products(ctx context.Context, in *proto.Things) (*proto.Response, error) {

	// 数据入库...
	// 断言判断使用了 oneof 中的哪个字段
	var reData string
	switch in.PayWay.(type) {
	case *proto.Things_AliPay:
		reData = "use alipay to pay the money"
	case *proto.Things_Cash:
		reData = "use cash to pay the money"
	case *proto.Things_WeChat:
		reData = "use wechat to pay the money"
	}

	// 数据转换成二进制形式
	reByte, err := json.Marshal(reData)

	// 处理错误
	switch err {
	case nil:
		return &proto.Response{
			Code:    200,
			Message: "insert into database success",
			Data:    &anypb.Any{Value: reByte},
		}, nil
	default:
		return &proto.Response{
			Code:    500,
			Message: "server busy",
			Data:    nil,
		}, err
	}
}

func (s *Server) Update(ctx context.Context, in *proto.UpdateThings) (*proto.Response, error) {

	mask, err := fieldmask_utils.MaskFromProtoFieldMask(in.UpdateMask, strcase.ToCamel)
	var dst = make(map[string]interface{}) // 这里必须使用 make 的方式， 创建并且初始化，否则报错
	fieldmask_utils.StructToMap(mask, in.Thing, dst)

	// 转换成 []byte 类型，放入 anypb.Any{} 中，将地址赋值给定义的字段变量
	data, _ := json.Marshal(dst)
	switch err {
	case nil:
		return &proto.Response{
			Code:    200,
			Message: "update success",
			Data:    &anypb.Any{Value: data},
		}, nil
	default:
		return &proto.Response{
			Code:    500,
			Message: "busy server",
			Data:    nil,
		}, err
	}
}

func main() {

	// 启动服务
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("net.Listen failed:%v\n", err)
		return
	}

	// 创建rpc服务
	s := grpc.NewServer()

	// 注册rpc服务 ： 服务 + 对象
	proto.RegisterSaleServer(s, &Server{})

	// 启动 rpc 服务
	err = s.Serve(listener)
	if err != nil {
		fmt.Printf("s.serve failed:%v\n", err)
		return
	}
}
