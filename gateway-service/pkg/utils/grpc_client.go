package utils

import (
	"google.golang.org/grpc"
	"log"
)

func CreateGRPCClient(address string) *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	return conn
}
