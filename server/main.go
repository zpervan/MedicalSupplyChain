package main

import (
	"context"
	"fmt"
	"log"
	"net"
	 ms "server/proto"

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
	lis, err := net.Listen("tcp", ":50000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	sopts := []grpc.ServerOption{}

	s := grpc.NewServer(sopts...)
	ms.RegisterMedicalSuppliesServer(s, &server{})

	log.Printf("Starting server...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
