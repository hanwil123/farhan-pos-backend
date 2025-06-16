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
	// Initialize database connections
	database.ConnectUser()
	database.ConnectCategory()
	database.ConnectProduct()

	// Start gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create gRPC server with options
	grpcServer := grpc.NewServer(
		grpc.MaxConcurrentStreams(100),
		grpc.MaxRecvMsgSize(1024*1024*10), // 10MB
	)

	// Register services
	proto.RegisterUserServiceServer(grpcServer, &handler.UserServiceServer{})
	// bakery services
	proto.RegisterBakeryPOSServiceServer(grpcServer, &handler.BakeryProductServiceServer{})

	fmt.Println("gRPC Server is running at :50051")

	// Start server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
