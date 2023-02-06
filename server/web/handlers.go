package web

import (
	"context"
	"fmt"
	gpb "github.com/golang/protobuf/ptypes/empty"
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

func (s *Server) TestRequest(_ context.Context, in *ms.Request) (*ms.DummyResponse, error) {
	s.Log.Info("received request from client")

	message := fmt.Sprintf("Hey, my dear friend %s", in.Data)

	return &ms.DummyResponse{Data: message}, nil
}

func (s *Server) CreateUser(_ context.Context, in *ms.User) (*ms.UserResponse, error) {
	s.Log.Info("creating new user")

	err := s.Database.CreateUser(in)
	if err != nil {
		return &ms.UserResponse{Error: &ms.Error{Code: ms.ErrorCode_INTERNAL_ERROR, Message: err.Error()}}, err
	}

	return &ms.UserResponse{}, nil
}

func (s *Server) FetchUsers(_ context.Context, _ *gpb.Empty) (*ms.UserResponse, error) {
	s.Log.Info("fetching all users")
	users, err := s.Database.FetchUsers()

	if err != nil {
		errorMessage := fmt.Sprintf("could not fetch all users. reason: %s", err.Error())
		return &ms.UserResponse{Error: &ms.Error{Code: ms.ErrorCode_INTERNAL_ERROR, Message: errorMessage}}, err
	}

	return &ms.UserResponse{Users: users}, nil
}
