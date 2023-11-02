package grpcserver

import (
	"context"
	"fmt"

	pb "userservice/internal/infra/grpcserver/.gen"
)

func (s *Server) GetUserById(ctx context.Context, input *pb.UserRequest) (output *pb.User, err error) {
	userID := input.Id
	resp, err := s.userservice.GetUserByID(context.Background(), userID)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	output = &pb.User{
		Id:      resp.Id,
		Fname:   resp.Fname,
		City:    resp.City,
		Phone:   resp.Phone,
		Height:  resp.Height,
		Married: resp.Married,
	}
	return
}

func (s *Server) GetUsersByIds(input *pb.UsersRequest, stream pb.UserService_GetUsersByIdsServer) error {
	for _, userID := range input.Ids {
		user, err := s.GetUserById(stream.Context(), &pb.UserRequest{Id: userID})
		if err == nil {
			if err := stream.Send(user); err != nil {
				return err
			}
		}
	}
	return nil
}
