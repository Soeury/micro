package middleware

import (
	"context"

	"github.com/go-kit/log"

	"github.com/go-kit/kit/endpoint"
)

// transport 层中间件
func LoggerMiddleWare(logger log.Logger) endpoint.Middleware {

	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Log("msg", "start endpoint")
			defer logger.Log("msg", "end endpoint")
			return next(ctx, request)
		}
	}
}
