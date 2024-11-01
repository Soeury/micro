package service

import (
	"context"
	"time"

	"github.com/go-kit/log"
)

// 服务层日志中间件 : 定义一个新类型把原先的 service 接口和一个额外的 logger 包装起来，并实现这个接口即可
// (经典加一层...感觉有点麻烦......)
type LogMiddleWare struct {
	logger log.Logger
	next   ComputerService
}

// 下面重写新定义的结构体内部的原始接口
func (m LogMiddleWare) Add(ctx context.Context, num1 int64, num2 int64) (ret int64, err error) {

	// go-kit 内部的日志库是 kv 类型的
	// 记录日志
	defer func(begin time.Time) {
		m.logger.Log(
			"method", "add",
			"num1", num1,
			"num2", num2,
			"output", ret,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	ret, err = m.next.Add(ctx, num1, num2)
	return
}

func (m LogMiddleWare) Append(ctx context.Context, str1 string, str2 string) (ret string, err error) {

	// 这里的 defer 实现的匿名函数指定了参数类型，参数的传入在后面的()内部实现
	defer func(begin time.Time) {
		m.logger.Log(
			"method", "append",
			"str1", str1,
			"str2", str2,
			"output", ret,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	ret, err = m.next.Append(ctx, str1, str2)
	return
}

// NewLogMiddleWare 使用新定义的结构体 创建一个带日志的 ComputerService
func NewLogMiddleWare(logger log.Logger, next ComputerService) ComputerService {

	return &LogMiddleWare{
		logger: logger,
		next:   next,
	}
}
