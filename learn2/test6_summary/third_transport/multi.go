package thirdTransport

import (
	"context"
	"micro/learn2/test6_summary/proto/multi"
)

type MultiRequest struct {
	Figure1 int64
	Figure2 int64
}

type MultiResponse struct {
	Res1 int64
	Res2 int64
}

func EncodeGRPCMultiRequest(_ context.Context, request interface{}) (interface{}, error) {

	req := request.(MultiRequest)
	return &multi.MultiRequest{Figure1: req.Figure1, Figure2: req.Figure2}, nil
}

func DecodeGRPCMultiResponse(_ context.Context, response interface{}) (interface{}, error) {

	resp := response.(*multi.MultiResponse)
	return MultiResponse{Res1: resp.GetRes1(), Res2: resp.GetRes2()}, nil
}
