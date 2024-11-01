package transport

import (
	"context"
	"encoding/json"
	"micro/learn2/test6_summary/endpoint"
	"micro/learn2/test6_summary/middleware"
	"micro/learn2/test6_summary/service"
	"net/http"

	"github.com/gin-gonic/gin"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"golang.org/x/time/rate"
)

// httptransport "github.com/go-kit/kit/transport/http"
// http json

func DecodeAddRequest(ctx context.Context, request *http.Request) (interface{}, error) {

	var req endpoint.AddRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeAppendRequest(ctx context.Context, request *http.Request) (interface{}, error) {

	var req endpoint.AppendRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {

	return json.NewEncoder(w).Encode(response)
}

// 创建 http 处理器
func NewHttpServer(srv service.ComputerService, logger log.Logger) http.Handler {

	// 添加中间件
	// 感受: 感觉和 gin 差不多，都是在处理实际业务之前会处理好这些中间件
	add := endpoint.MakeAddEndpoint(srv)
	add = middleware.LoggerMiddleWare(logger)(add)
	add = middleware.RateMiddleWare(rate.NewLimiter(1, 1))(add)

	append := endpoint.MakeAppendEndpoint(srv)
	append = middleware.LoggerMiddleWare(logger)(append)
	append = middleware.RateMiddleWare(rate.NewLimiter(1, 1))(append)

	addHandler := httptransport.NewServer(
		add,
		DecodeAddRequest,
		EncodeResponse,
	)

	appendHandler := httptransport.NewServer(
		append,
		DecodeAppendRequest,
		EncodeResponse,
	)

	// use http
	// http.Handle("/add", addHandler)
	// http.Handle("/append", appendHandler)

	// use gin
	r := gin.Default()

	// gin.WrapH 将 http 处理函数转换为 gin 处理函数
	r.POST("/add", gin.WrapH(addHandler))
	r.POST("/append", gin.WrapH(appendHandler))
	return r
}
