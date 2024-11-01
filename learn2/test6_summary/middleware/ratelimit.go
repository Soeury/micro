package middleware

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/time/rate"
)

var ErrRateLimit = errors.New("need rate limit")

// 限流中间件
func RateMiddleWare(limit *rate.Limiter) endpoint.Middleware {

	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			if limit.Allow() {
				return next(ctx, request)
			} else {
				return nil, ErrRateLimit
			}
		}
	}
}
