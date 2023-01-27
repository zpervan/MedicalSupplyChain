package main

import (
	"context"
	"fmt"
	"log"
	"net"
	ms "server/proto"

	"server/core"

	"google.golang.org/grpc"
)

type server struct {
	ms.UnimplementedMedicalSuppliesServer
}

func (s *server) FetchAll(ctx context.Context, in *ms.Request) (*ms.Response, error) {
	log.Println("Received request from client")

	message := fmt.Sprintf("Hey, my dear friend %s", in.Data)

	return &ms.Response{Data: message}, nil
}

func main() {
	logger := core.NewLogger()
	logger.Info("Initializing server...")

	lis, err := net.Listen("tcp", ":50000")
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to listen: %v", err))
	}

	sopts := []grpc.ServerOption{}

	s := grpc.NewServer(sopts...)
	ms.RegisterMedicalSuppliesServer(s, &server{})

	logger.Info("Starting server...")
	if err := s.Serve(lis); err != nil {
		logger.Fatal(fmt.Sprintf("failed to serve: %v", err))
	}
}
