package main

import (
	"context"
	"eatwut-permission-service/util"
	"fmt"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"

	"github.com/loopholelabs/frpc-go-examples/grpc/echo"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	zapLogger  *zap.Logger
	customFunc grpc_zap.CodeToLevel
)

type svc struct {
	echo.UnimplementedEchoServiceServer
}

func (s *svc) Echo(_ context.Context, req *echo.Request) (*echo.Response, error) {
	res := new(echo.Response)
	res.Message = req.Message
	return res, nil
}

func main() {
	// Init Configuration
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Init connection TCP
	port, err := net.Listen("tcp", fmt.Sprintf(":%s", config.GRPC_SERVER_PORT))
	if err != nil {
		panic(err)
	}

	//Set up Interceptor and initialize gRPC server
	zap, zap_opt := util.SetupZapLogger()
	grpc := grpc.NewServer( // --- ③
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(zap, zap_opt),
		),
	)

	echo.RegisterEchoServiceServer(grpc, new(svc))

	reflection.Register(grpc)
	log.Println(fmt.Sprintf("Grpc Server is listening on port %s", config.GRPC_SERVER_PORT))

	if err := grpc.Serve(port); err != nil {
		log.Fatal(err)
	}
}
