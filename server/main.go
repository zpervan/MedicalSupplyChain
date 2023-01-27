package main

import (
	"fmt"
	"net"

	"server/core"
	"server/web"
	ms "server/proto"

	"google.golang.org/grpc"
)

func main() {
	logger := core.NewLogger()
	logger.Info("Initializing server...")

	listener, err := net.Listen("tcp", ":50000")
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to listen: %v", err))
	}

	serverOptions := []grpc.ServerOption{}

	grpcServer := grpc.NewServer(serverOptions...)
	ms.RegisterMedicalSuppliesServer(grpcServer, web.NewServer(logger))

	logger.Info("Starting server...")
	
	err = grpcServer.Serve(listener)
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to serve: %v", err))
	}
}
