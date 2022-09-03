package server

import (
	"context"

	"github.com/loopholelabs/frpc-go-examples/grpc/echo"
)

type Echo struct {
	echo.UnimplementedEchoServiceServer
}

func (s *Echo) Echo(_ context.Context, req *echo.Request) (*echo.Response, error) {
	res := new(echo.Response)
	res.Message = req.Message
	return res, nil
}
