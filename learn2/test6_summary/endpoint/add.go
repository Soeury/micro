package endpoint

import (
	"context"
	"micro/learn2/test6_summary/service"

	"github.com/go-kit/kit/endpoint"
)

// endpoint 对外提供的方法的包装
// endpoint 本身是一个自定义的函数类型，具有参数格式的要求

type AddRequest struct {
	Num1 int64 `json:"num1"`
	Num2 int64 `json:"num2"`
}

type AddResponse struct {
	Ret int64  `json:"ret"`
	Err string `json:"err,omitempty"`
}

func MakeAddEndpoint(srv service.ComputerService) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(AddRequest)
		v, err := srv.Add(ctx, req.Num1, req.Num2)
		if err != nil {
			return AddResponse{Ret: v, Err: err.Error()}, nil // 响应数据里面包装了错误，这里就不需要返回
		}
		return AddResponse{Ret: v}, nil
	}
}
