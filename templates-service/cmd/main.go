package main

import (
	health "cv-templates-service/internal/health/grpc"
	handlers2 "cv-templates-service/internal/health/handlers"
	templates "cv-templates-service/internal/templates/grpc"
	"cv-templates-service/internal/templates/handlers"
	"cv-templates-service/pkg/containers"
	infrastructure "cv-templates-service/pkg/infrastructure/grpc"
	handlers3 "cv-templates-service/pkg/infrastructure/handlers"
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

	templateServer := handlers.NewTemplateServiceServer(dependencies.DefaultTemplateService, dependencies.TemplateService, dependencies.ColorService, dependencies.MinioClient, dependencies.Logger)
	templates.RegisterTemplateServiceServer(grpcServer, templateServer)

	healthServer := handlers2.NewHealthServiceServer(dependencies.DB, dependencies.RedisClient)
	health.RegisterHealthServiceServer(grpcServer, healthServer)

	seedingServer := handlers3.NewSeederServiceServer(dependencies)
	infrastructure.RegisterSeederServiceServer(grpcServer, seedingServer)

	log.Printf("gRPC server is running on port %s", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
