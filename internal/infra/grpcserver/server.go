package grpcserver

import (
	"userservice/internal/domain"
	userservicegrpc "userservice/internal/infra/grpcserver/.gen"
)

type Server struct {
	userservicegrpc.UnimplementedUserServiceServer
	userservice domain.Service
}

func NewServer(userService domain.Service) *Server {
	return &Server{
		userservice: userService,
	}
}
