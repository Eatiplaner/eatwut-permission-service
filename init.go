package main

import (
	"context"
	"eatwut-permission-service/util"
	"fmt"
	"log"
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
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.GRPC_SERVER_PORT))
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

	log.Println(fmt.Sprintf("Grpc Server is listening on port %s", config.GRPC_SERVER_PORT))
	err = server.Serve(lis)

	if err != nil {
		panic(err)
	}
}
