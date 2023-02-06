package web

import (
	"context"
	"fmt"
	"server/core"
	"server/database"
	ms "server/proto"
)

type Server struct {
	Log      *core.Logger
	Database *database.Database
	ms.UnimplementedMedicalSuppliesServer
}

func NewServer(log *core.Logger, database *database.Database) *Server {
	server := &Server{}

	server.Log = log
	server.Database = database

	return server
}

func (s *Server) TestRequest(ctx context.Context, in *ms.Request) (*ms.Response, error) {
	s.Log.Info("received request from client")

	message := fmt.Sprintf("Hey, my dear friend %s", in.Data)

	return &ms.Response{Data: message, Error: &ms.Error{Code: ms.ErrorCode_OK, Message: ""}}, nil
}

func (s *Server) CreateUser(ctx context.Context, in *ms.User) (*ms.Response, error) {
	s.Log.Info("creating new user")

	err := s.Database.CreateUser(in)
	if err != nil {
		return &ms.Response{Error: &ms.Error{Code: ms.ErrorCode_INTERNAL_ERROR, Message: err.Error()}}, err
	}
	
	return &ms.Response{}, nil
}
