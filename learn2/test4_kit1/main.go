package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http" // 记得匿名导入包
)

var (
	ErrEmptyStr = errors.New("all strings are empty")
)

// 1.1 业务逻辑抽象定义服务接口
type AddService interface {
	Sum(context.Context, int, int) (int, error)
	AppendStr(context.Context, string, string) (string, error)
}

// 1.2 实现接口: 结构体 + 方法
type addService struct{}

// Add返回两个字段的和
func (addService) Sum(_ context.Context, a int, b int) (int, error) {

	return a + b, nil
}

// AppendStr返回两个字段的拼接
func (addService) AppendStr(_ context.Context, a string, b string) (string, error) {

	if a == "" && b == "" {
		return "", ErrEmptyStr
	}

	return a + b, nil
}

// 在 Go kit 中，主要的消息模式是 RPC
// 因此，我们接口中的每个方法都将被建模为一个远程过程调用
// 1.3 对于每个方法，我们最好都要定义请求和响应结构体，分别捕获所有的输入和输出参数
type SumRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type SumResponse struct {
	Value int    `json:"value"`
	Err   string `json:"err,omitempty"`
}

type AppendStrRequest struct {
	A string `json:"a"`
	B string `json:"b"`
}

type AppendStrResponse struct {
	Value string `json:"value"`
	Err   string `json:"err,omitempty"`
}

// 2. endpoint
// 借助 适配器 将 方法 -> endpoint
func MakeSumEndpoint(srv AddService) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SumRequest)
		v, err := srv.Sum(ctx, req.A, req.B)
		if err != nil {
			return SumResponse{Value: v, Err: err.Error()}, nil
		}
		return SumResponse{Value: v}, nil
	}
}

func MakeAppenStrEndpoint(srv AddService) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AppendStrRequest)
		v, err := srv.AppendStr(ctx, req.A, req.B)
		if err != nil {
			return AppendStrResponse{Value: v, Err: err.Error()}, nil
		}
		return AppendStrResponse{Value: v}, nil
	}
}

//  3. transport
//     3.1 解码 : 请求来了之后根据协议(http , http2)和编码(json , thrift , pb)去解析数据
func DecodeSumRequest(ctx context.Context, request *http.Request) (interface{}, error) {

	var req SumRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeAppendRequest(ctx context.Context, request *http.Request) (interface{}, error) {

	var req AppendStrRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

// 3.2 编码: 把响应数据按照协议和编码返回
func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func main() {

	// 导入包 ： httptransport "github.com/go-kit/kit/transport/http"
	srv := addService{}

	// http json 服务
	sumHandler := httptransport.NewServer(
		MakeSumEndpoint(srv),
		DecodeSumRequest,
		EncodeResponse,
	)

	appendHandler := httptransport.NewServer(
		MakeAppenStrEndpoint(srv),
		DecodeAppendRequest,
		EncodeResponse,
	)

	http.Handle("/sum", sumHandler)
	http.Handle("/append", appendHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Printf("http.ListenAndServe failed:%v\n", err)
		return
	}

	// go-kit 三个步骤:
	//  -1. 定义服务(接口+结构体+方法)，定义请求和响应结构体
	//  -2. 接口中所有的方法转 -> endpoint，函数通常以 Make___ 命名
	//  -3. 请求和响应的编码，解码的定义
}
