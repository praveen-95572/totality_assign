package main

import (
	"context"
	"log"

	grpc_infra "userservice/internal/infra/grpcserver/.gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8088", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := grpc_infra.NewUserServiceClient(conn)
	response := GetUsersByIds(c)
	log.Printf("Response from Server: %s", response)
}

func GetUserById(c grpc_infra.UserServiceClient) *grpc_infra.User {
	message := grpc_infra.UserRequest{
		Id: 1,
	}
	response, err := c.GetUserById(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling GetUserById: %s", err)
	}
	return response
}

func GetUsersByIds(c grpc_infra.UserServiceClient) grpc_infra.UserService_GetUsersByIdsClient {
	message := grpc_infra.UsersRequest{
		Ids: []int64{1, 2, 3},
	}
	response, err := c.GetUsersByIds(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling GetUserById: %s", err)
	}
	return response
}
