package main

import (
	"Farhan-Backend-POS/controllers/handler"
	"Farhan-Backend-POS/database"
	"Farhan-Backend-POS/proto"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	database.ConnectUser()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterUserServiceServer(grpcServer, &handler.UserServiceServer{})
	fmt.Println("UserService is running at :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
