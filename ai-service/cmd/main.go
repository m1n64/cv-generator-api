package main

import (
	ai "ai-service/internal/ai/grpc"
	"ai-service/internal/ai/handlers"
	health "ai-service/internal/health/grpc"
	handlers2 "ai-service/internal/health/handlers"
	"ai-service/pkg/containers"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println("Templates service started!")

	dependencies, _ := containers.InitializeDependencies()
	go containers.InitializeQueuesConsumer(dependencies)

	port := os.Getenv("SERVICE_PORT")
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()

	aiServer := handlers.NewAiServiceServer(dependencies.AiService, dependencies.Logger)
	ai.RegisterAiServiceServer(grpcServer, aiServer)

	healthServer := handlers2.NewHealthServiceServer(dependencies.DB, dependencies.RedisClient)
	health.RegisterHealthServiceServer(grpcServer, healthServer)

	log.Printf("gRPC server is running on port %s", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
