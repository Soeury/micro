package endpoint

import (
	"micro/learn2/test6_summary/proto/trim"
	third "micro/learn2/test6_summary/third_transport"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

// 创建一个基于 grpc.Client 客户端的 endpoint.Endpoint
// 需要区分客户端包装 endpoint 和 自定义的 endpoint
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
