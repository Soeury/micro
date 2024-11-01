package endpoint

import (
	"micro/learn2/test6_summary/proto/multi"
	third "micro/learn2/test6_summary/third_transport"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

// 同 trim
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
