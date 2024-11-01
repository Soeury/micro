package thirdTransport

import (
	"context"
	"micro/learn2/test6_summary/proto/trim"
)

type TrimRequest struct {
	Word1 string
}

type TrimResponse struct {
	Ret1 string
}

// trim
// 将 内部的的请求结构体转换成 grpc 请求的结构体，因为是内部作为 grpc 客户端，来调用外部编写好的 grpc 服务
func EncodeGRPCTrimRequest(_ context.Context, request interface{}) (interface{}, error) {

	req := request.(TrimRequest)
	return &trim.TrimRequest{Word1: req.Word1}, nil
}

// 外部 grpc 返回响应后，包装内部响应的结构体形式，便于内部调用
func DecodeGRPCTrimResponse(_ context.Context, response interface{}) (interface{}, error) {

	resp := response.(*trim.TrimResponse)
	return TrimResponse{Ret1: resp.GetRet1()}, nil
}
