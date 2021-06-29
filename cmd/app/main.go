package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/sapawarga/auth/cmd/container"
	"github.com/sapawarga/auth/cmd/db"
	"github.com/sapawarga/auth/config"
	grpcTransport "github.com/sapawarga/auth/transport/grpc"
	"github.com/sapawarga/proto-file/auth"

	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Get()

	logger := initLogger()
	containers := container.Init(db.NewConnection(cfg), logger, "grpc")
	grpcAdd := flag.String("grpc", fmt.Sprintf(":%s", cfg.APP_GRPC_PORT), "gRPC listening address")

	logger.Log("transport", "grpc", "addr", grpcAdd)

	lis, err := net.Listen("tcp", *grpcAdd)
	if err != nil {
		logger.Log("transport", "grpc", "during", "listen", "err", err)
		os.Exit(1)
	}

	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)
	handler := grpcTransport.MakeHandler(containers.AuthService)

	auth.RegisterAuthHandlerServer(server, handler)
	err = server.Serve(lis)
	if err != nil {
		logger.Log("transport", "grpc", "during", "listen", "err", err)
	}
}

func initLogger() log.Logger {
	logger := log.NewJSONLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	return logger
}
