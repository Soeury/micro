package service

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

// 服务层限流中间件：
type RateLimitMiddleWare struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           ComputerService
}

func (rl RateLimitMiddleWare) Add(ctx context.Context, num1 int64, num2 int64) (ret int64, err error) {

	defer func(begin time.Time) {
		lvs := []string{"method", "add", "error", fmt.Sprint(err != nil)}
		rl.requestCount.With(lvs...).Add(1)
		rl.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
		rl.countResult.Observe(float64(ret))
	}(time.Now())

	ret, err = rl.next.Add(ctx, num1, num2)
	return
}

func (rl RateLimitMiddleWare) Append(ctx context.Context, str1 string, str2 string) (ret string, err error) {

	defer func(begin time.Time) {
		lvs := []string{"method", "append", "error", "false"}
		rl.requestCount.With(lvs...).Add(1)
		rl.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	ret, err = rl.next.Append(ctx, str1, str2)
	return
}

// NewRateLimitMiddleWare 使用新定义的结构体 创建一个带限流的 ComputerService
func NewRateLimitMiddleWare(
	requestCount metrics.Counter,
	requestLatency metrics.Histogram,
	countResult metrics.Histogram,
	next ComputerService,
) ComputerService {

	return &RateLimitMiddleWare{
		requestCount:   requestCount,
		requestLatency: requestLatency,
		countResult:    countResult,
		next:           next,
	}
}
