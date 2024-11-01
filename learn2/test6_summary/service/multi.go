package service

import (
	"context"

	third "micro/learn2/test6_summary/third_transport"

	"github.com/go-kit/kit/endpoint"
)

type WithMultiMiddleWares struct {
	next         ComputerService
	multiService endpoint.Endpoint
}

func (mm WithMultiMiddleWares) Add(ctx context.Context, num1 int64, num2 int64) (ret int64, err error) {

	// 1. 发起请求
	// 2. 获取响应】
	// 3. 数据断言
	resp, err := mm.multiService(ctx, third.MultiRequest{Figure1: num1, Figure2: num2})
	if err != nil {
		return 0, err
	}

	multi := resp.(third.MultiResponse)

	ret, err = mm.next.Add(ctx, multi.Res1, multi.Res2)
	return
}

func (mm WithMultiMiddleWares) Append(ctx context.Context, str1 string, str2 string) (ret string, err error) {

	ret, err = mm.next.Append(ctx, str1, str2)
	return
}

func NewWithMultiMiddleWares(multiService endpoint.Endpoint, next ComputerService) ComputerService {

	return &WithMultiMiddleWares{
		next:         next,
		multiService: multiService,
	}
}
