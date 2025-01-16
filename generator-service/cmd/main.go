package main

import (
	generator "cv-generator-service/internal/generator/grpc"
	"cv-generator-service/internal/generator/handlers"
	health "cv-generator-service/internal/health/grpc"
	handlers2 "cv-generator-service/internal/health/handlers"
	"cv-generator-service/pkg/containers"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println("CV PDF Generator service started!")

	dependencies, _ := containers.InitializeDependencies()
	go containers.InitializeQueuesConsumer(dependencies)

	port := os.Getenv("SERVICE_PORT")
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()

	generatorServer := handlers.NewGeneratorServiceServer(dependencies.PdfGeneratorService, dependencies.MinioClient, dependencies.Logger)
	generator.RegisterGeneratorServiceServer(grpcServer, generatorServer)

	healthServer := handlers2.NewHealthServiceServer(dependencies.DB, dependencies.RedisClient)
	health.RegisterHealthServiceServer(grpcServer, healthServer)

	log.Printf("gRPC server is running on port %s", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
