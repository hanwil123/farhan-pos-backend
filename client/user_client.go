package client

import (
	"Farhan-Backend-POS/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var UserClient proto.UserServiceClient
var CategoryClient proto.BakeryPOSServiceClient

func InitializeClient() {
	// Use grpc.NewClient instead of grpc.Dial
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	// defer conn.Close()

	UserClient = proto.NewUserServiceClient(conn)
	CategoryClient = proto.NewBakeryPOSServiceClient(conn)
}
