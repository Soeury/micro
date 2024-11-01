package endpoint

import (
	"context"
	"micro/learn2/test6_summary/service"

	"github.com/go-kit/kit/endpoint"
)

// endpoint 对外提供的方法的包装
// endpoint 本身是一个自定义的函数类型，具有参数格式的要求

type AppendRequest struct {
	Str1 string `json:"str1"`
	Str2 string `json:"str2"`
}

type AppendResponse struct {
	Ret string `json:"ret"`
	Err string `json:"err,omitempty"`
}

func MakeAppendEndpoint(srv service.ComputerService) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AppendRequest)
		v, err := srv.Append(ctx, req.Str1, req.Str2)
		if err != nil {
			return AppendResponse{Ret: v, Err: err.Error()}, nil // 响应数据里面包装了错误，这里就不需要返回
		}
		return AppendResponse{Ret: v}, nil
	}
}
