package main

import (
	"Farhan-Backend-POS/database"
	grpcServiceAuth "Farhan-Backend-POS/modules/auth/delivery-handler/grpc"
	grpcServiceBakery "Farhan-Backend-POS/modules/bakery/delivery-handler/grpc"
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
	proto.RegisterUserServiceServer(grpcServer, &grpcServiceAuth.UserServiceServer{})
	// bakery services
	proto.RegisterBakeryPOSServiceServer(grpcServer, &grpcServiceBakery.BakeryProductServiceServer{})

	fmt.Println("gRPC Server is running at :50051")

	// Start server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
