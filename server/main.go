package main

import (
	"fmt"
	"net"
	"os"

	"server/core"
	"server/database"
	ms "server/proto"
	"server/web"

	"google.golang.org/grpc"
)

func main() {
	logger := core.NewLogger()
	logger.Info("initializing server...")

	var psqlInfo string

	if os.Getenv("DATABASE_HOST") != "" {
		// Login data acquired from Docker container environment
		psqlInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	} else {
		// Non-Docker (local) PSQL login data
		psqlInfo = "host=localhost port=5432 user=user password=password dbname=msc_db sslmode=disable"
	}

	database, err := database.NewDatabase(logger, psqlInfo)
	if err != nil {
		logger.Fatal("cannot connect to database. reason: " + err.Error())
	}

	listener, err := net.Listen("tcp", ":3500")
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to listen: %v", err))
	}

	serverOptions := []grpc.ServerOption{}

	grpcServer := grpc.NewServer(serverOptions...)
	ms.RegisterMedicalSuppliesServer(grpcServer, web.NewServer(logger, database))

	logger.Info("starting server...")

	err = grpcServer.Serve(listener)
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to serve: %v", err))
	}
}
