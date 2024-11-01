package service

import (
	"fmt"
	"io"
	"micro/learn2/test6_summary/proto/multi"
	"micro/learn2/test6_summary/proto/trim"
	"time"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	third "micro/learn2/test6_summary/third_transport"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	sdconsul "github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
)

// consul
// 从注册中心获取 trim 服务的地址
// 基于 consul 的 trim srevice 服务发现

// Go kit 为不同的服务发现系统（eureka、zookeeper、consul、etcd等）提供适配器
// Endpointer负责监听服务发现系统，并根据需要生成一组相同的端点
//     type Endpointer interface {
//    	   Endpoints() ([]endpoint.Endpoint, error)
//     }

// Go kit 提供了工厂函数 —— Factory
// 它是一个将实例字符串(例如host:port)转换为特定端点的函数。提供多个端点的实例需要多个工厂函数。
// ------ type Factory func(instance string) (endpoint.Endpoint, io.Closer, error)

func GetTrimServiceFromConsul(consulAddr string, logger log.Logger, serviceName string, tags []string) (endpoint.Endpoint, error) {

	// 1. 连接 consul
	conuslConfig := api.DefaultConfig()
	conuslConfig.Address = consulAddr
	cc, err := api.NewClient(conuslConfig)

	if err != nil {
		fmt.Printf("newClient failed:%v\n", err)
		return nil, err
	}

	// 2. 使用 go-kit 提供的适配器
	sdClient := sdconsul.NewClient(cc)
	instancer := sdconsul.NewInstancer(sdClient, logger, serviceName, tags, true)

	// 3. endpoints 返回 endpoints 的集合
	endpointer := sd.NewEndpointer(instancer, Factory_trim, logger)

	// 4. balancer 实现 round_robin 负载均衡(从集合中选出指定的 endpoint)
	balancer := lb.NewRoundRobin(endpointer)

	// 5. retry 重试
	// 重试策略包装负载均衡器，并返回可用的端点。重试策略将重试失败的请求，直到达到最大尝试或超时为止
	retry := lb.Retry(3, time.Second, balancer)
	return retry, nil
}

func GetMultiServiceFromConsul(consulAddr string, logger log.Logger, serviceName string, tags []string) (endpoint.Endpoint, error) {

	// 1. 连接 consul
	cc, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		fmt.Printf("api.NewClient failed:%v\n", err)
		return nil, err
	}

	// 2. 使用 go-kit 提供的适配器
	sdClient := sdconsul.NewClient(cc)
	instancer := sdconsul.NewInstancer(sdClient, logger, serviceName, tags, true)

	// 3. endpoints 返回 endpoints 的集合
	endpointer := sd.NewEndpointer(instancer, Factory_multi, logger)

	// 4. 负载均衡
	balancer := lb.NewRoundRobin(endpointer)

	// 5. retry
	retry := lb.Retry(3, time.Second, balancer)
	return retry, nil
}

// 这里的 conn 是可以通过 close 关闭的，所以可以是 io.Closer 类型
func Factory_trim(instance string) (endpoint.Endpoint, io.Closer, error) {

	conn, err := grpc.NewClient(instance, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("grpc.NewClient failed:%v\n", err)
		return nil, nil, err
	}

	my_endpoint := MakeTrimEndpoint(conn)
	return my_endpoint, conn, nil
}

func Factory_multi(instance string) (endpoint.Endpoint, io.Closer, error) {

	conn, err := grpc.NewClient(instance, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("grpc.newClient failed:%v\n", err)
		return nil, nil, err
	}

	my_endpoint := MakeMultiEndpoint(conn)
	return my_endpoint, conn, nil
}

// 创建一个基于 grpc.Client 客户端的 endpoint.Endpoint
// 需要区分客户端包装 endpoint 和 自定义的 endpoint
// ----- 下面这个函数在 endpoint 层也写了一遍，导入会产生循环导入包的问题。。。。
func MakeTrimEndpoint(conn *grpc.ClientConn) endpoint.Endpoint {

	return grpctransport.NewClient(
		conn,
		"trim.Trim",
		"TrimSpace",
		third.EncodeGRPCTrimRequest,
		third.DecodeGRPCTrimResponse,
		trim.TrimResponse{},
	).Endpoint()
}

// 同样循环导入包的问题......
func MakeMultiEndpoint(conn *grpc.ClientConn) endpoint.Endpoint {

	return grpctransport.NewClient(
		conn,
		"multi.Multi",
		"Multiply",
		third.EncodeGRPCMultiRequest,
		third.DecodeGRPCMultiResponse,
		multi.MultiResponse{}, // GRPC 的响应 注意这里返回类型，是我们这个包下的 multi.MultiResponse{} 不是 test8 中的
	).Endpoint()
}
