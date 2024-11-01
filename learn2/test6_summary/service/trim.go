package service

import (
	"context"
	third "micro/learn2/test6_summary/third_transport"

	"github.com/go-kit/kit/endpoint"
)

// 调用第三方的服务也是通过包装成 中间件的形式来解决的
type WithTrimMiddleWare struct {
	next        ComputerService
	trimService endpoint.Endpoint
}

// Add 不变
func (tm WithTrimMiddleWare) Add(ctx context.Context, num1 int64, num2 int64) (ret int64, err error) {

	ret, err = tm.next.Add(ctx, num1, num2)
	return
}

func (tm WithTrimMiddleWare) Append(ctx context.Context, str1 string, str2 string) (ret string, err error) {

	// 这里需要新的逻辑处理
	// 当我们调用新第三方服务时
	//  -1. 发起请求，对数据进行处理(调用第三方服务/依赖第三方服务)
	//  -2. 拿到数据
	//  -3. 数据断言
	resp1, err := tm.trimService(ctx, third.TrimRequest{Word1: str1})
	if err != nil {
		return "", err
	}

	resp2, err := tm.trimService(ctx, third.TrimRequest{Word1: str2})
	if err != nil {
		return "", err
	}

	trim1 := resp1.(third.TrimResponse)
	trim2 := resp2.(third.TrimResponse)

	ret, err = tm.next.Append(ctx, trim1.Ret1, trim2.Ret1)
	return
}

// NewWithTrimMiddleWare 使用新定义的结构体 创建一个调用了第三方服务的 ComputerService
func NewWithTrimMiddleWare(trimService endpoint.Endpoint, next ComputerService) ComputerService {

	return &WithTrimMiddleWare{
		next:        next,
		trimService: trimService,
	}
}
