package web

import (
	"context"
	"fmt"
	"server/core"
	"server/database"
	ms "server/proto"
)

type Server struct {
	Log *core.Logger
	Database *database.Database
	ms.UnimplementedMedicalSuppliesServer
}

func NewServer(log *core.Logger, database *database.Database) *Server {
	server := &Server{}
	
	server.Log = log
	server.Database = database

	return server
}

func (s *Server) FetchAll(ctx context.Context, in *ms.Request) (*ms.Response, error) {
	s.Log.Info("Received request from client")

	message := fmt.Sprintf("Hey, my dear friend %s", in.Data)

	return &ms.Response{Data: message}, nil
}
