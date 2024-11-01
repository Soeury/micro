package main

import (
	"context"
	"fmt"
	"time"

	"micro/learn1/test8_review/client/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// 基于 oneof , wrappers , field_mask 的一个小练习

func main() {

	// 建立连接
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("grpc.NewClient failed:%v\n", err)
		return
	}
	defer conn.Close()

	// 创建客户端
	client := proto.NewSaleClient(conn)

	// 调用远程服务
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// 定义某一个产品
	product := proto.Things{
		Pid:    1,
		Name:   "rabbit toy",
		Grade:  "three",
		Number: &wrapperspb.Int64Value{Value: 50},
		Money: &proto.Things_Price{
			Purchase: &wrapperspb.DoubleValue{Value: 81.95},
			Sell:     &wrapperspb.DoubleValue{Value: 98.63},
			Profit:   &wrapperspb.DoubleValue{Value: 16.68},
		},
		PayWay: &proto.Things_WeChat{
			WeChat: "use wechat to pay the money ",
		},
	}

	resp, err := client.Products(ctx, &product)
	switch err {
	case nil:
		fmt.Printf("resp:%s\n", string(resp.Data.Value))
	default:
		fmt.Printf("client.Products failed:%v\n", err)
		return
	}

	// 定义需要更新的字段
	paths := []string{"Grade", "Number", "Money.Sell", "Money.Profit"}
	updata := proto.UpdateThings{
		Op: "admin",
		Thing: &proto.Things{
			Grade:  "four",
			Number: &wrapperspb.Int64Value{Value: 98},
			Money: &proto.Things_Price{
				Sell:   &wrapperspb.DoubleValue{Value: 112.67},
				Profit: &wrapperspb.DoubleValue{Value: 30.72},
			},
		},
		UpdateMask: &fieldmaskpb.FieldMask{Paths: paths},
	}

	resp, err = client.Update(ctx, &updata)
	switch err {
	case nil:
		fmt.Printf("response:%s\n", string(resp.Data.Value))
	default:
		fmt.Printf("client.Update failed:%v\n", err)
		return
	}
}
