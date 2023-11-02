package main

import (
	"log"
	"net"
	"userservice/internal/infra/grpcserver"
	userPb "userservice/internal/infra/grpcserver/.gen"
	"userservice/internal/services/user"

	"google.golang.org/grpc"
)

func main() {
	log.Println("Initiating service ..")

	// create instance of userservice
	userservice := user.NewService()
	log.Println("Created an instance of  userservice ..")

	grpcURL := ":8088"
	lis, err := net.Listen("tcp", grpcURL)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userServer := grpcserver.NewServer(userservice)

	log.Printf("Requesting for GRPC server to listen on port %s\n", grpcURL)

	userPb.RegisterUserServiceServer(grpcServer, userServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
