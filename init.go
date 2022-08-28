package main

import (
	"context"
	"fmt"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"

	"github.com/loopholelabs/frpc-go-examples/grpc/echo"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var zapLogger *zap.Logger

type svc struct {
	echo.UnimplementedEchoServiceServer
}

func (s *svc) Echo(_ context.Context, req *echo.Request) (*echo.Response, error) {
	res := new(echo.Response)
	res.Message = req.Message
	return res, nil
}

func main() {
	// TODO: replace port 8080 with configure env, See: https://github.com/spf13/viper
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", "8080"))
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_zap.StreamServerInterceptor(zapLogger),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(zapLogger),
		)),
	)
	echo.RegisterEchoServiceServer(server, new(svc))

	err = server.Serve(lis)

	if err != nil {
		panic(err)
	}
}
