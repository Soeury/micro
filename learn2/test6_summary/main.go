package main

import (
	"fmt"
	"micro/learn2/test6_summary/proto/sum"
	"micro/learn2/test6_summary/service"
	"micro/learn2/test6_summary/transport"
	"net"
	"net/http"
	"os"

	// transportgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/gin-gonic/gin"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/log"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

// errorGroup : 提供了一个简单的方式来并行运行多个 goroutine，并等待它们全部完成或遇到第一个错误时立即返回
//  -1. Group : 管理一组 goroutine
//  -2. Wait  : 阻塞直到组内所有的 goroutine 都完成，或者其中一个返回了错误
//              如果组内有任何 goroutine 返回错误，Wait 会立即返回该错误，并且取消组内其他还在运行的 goroutine
//  -3. Go    : 组内启动 goroutine

// main 函数做服务的启动
// 启动 consul:
//
//	-1. 终端输入 consul agent -dev
//	-2. 网页URL : http://localhost:8500
func main() {

	// 前置资源初始化
	srv := service.NewService()

	// 初始化带 logger 的 srv
	logger := log.NewJSONLogger(os.Stderr)
	srv = service.NewLogMiddleWare(logger, srv)

	// 初始化带 ratelimit 的 srv
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here

	srv = service.NewRateLimitMiddleWare(requestCount, requestLatency, countResult, srv)

	// 初始化带第三方调用 trim 的 srv
	// conn_trim, err := grpc.NewClient("localhost:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	fmt.Printf("grpc.NewCLient in trim failed:%v\n", err)
	// 	return
	// }
	// defer conn_trim.Close()
	// trimEndpoint := endpoint.MakeTrimEndpoint(conn_trim)

	// // 从 consul 中读取 trim service
	trimEndpoint, err := service.GetTrimServiceFromConsul("localhost:8500", logger, "learn2_test7_trim", nil)
	if err != nil {
		fmt.Printf("service.GetTrimServiceFromConsul failed:%v\n", err)
		return
	}
	srv = service.NewWithTrimMiddleWare(trimEndpoint, srv)

	// 初始化第三方调用 multi 的 srv
	// conn_multi, err := grpc.NewClient("localhost:8083", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	fmt.Printf("grpc.NewCLient in multi failed:%v\n", err)
	// 	return
	// }
	// defer conn_multi.Close()

	// multiEndpoint := endpoint.MakeMultiEndpoint(conn_multi)

	// 从 consul 中读取 trim service
	multiEndpoint, err := service.GetMultiServiceFromConsul("localhost:8500", logger, "learn2_test8-multi", nil)
	if err != nil {
		fmt.Printf("service.GetMultiServiceFromConsul failed:%v\n", err)
		return
	}
	srv = service.NewWithMultiMiddleWares(multiEndpoint, srv)

	// 管理 goroutines :
	var group errgroup.Group

	// 以下两个 ip 是对外提供该服务的 ip , 其余的第三方服务的 ip 是内部调用的，
	// 所以需要访问服务的时候只需要调用下面的两个 ip 地址就可以了
	// grpc
	group.Go(func() error {
		grpcListener, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			return err
		}
		defer grpcListener.Close()

		s := grpc.NewServer()
		logger := log.NewLogfmtLogger(os.Stderr)
		sum.RegisterComputerServer(s, transport.NewGRPCServer(srv, logger))
		return s.Serve(grpcListener)
	})

	// http
	group.Go(func() error {
		httpListener, err := net.Listen("tcp", "localhost:8081")
		if err != nil {
			return err
		}
		defer httpListener.Close()

		logger := log.NewLogfmtLogger(os.Stderr)
		httpHandler := transport.NewHttpServer(srv, logger)

		// 添加 metrics gin路由，注意这里添加的处理函数
		httpHandler.(*gin.Engine).GET("/metrics", gin.WrapH(promhttp.Handler()))
		return http.Serve(httpListener, httpHandler)
	})

	if err := group.Wait(); err != nil {
		fmt.Printf("goroutines exit with err:%v\n", err)
		return
	}
}
