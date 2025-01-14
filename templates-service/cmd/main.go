package main

import (
	templates "cv-templates-service/internal/templates/grpc"
	"cv-templates-service/internal/templates/handlers"
	"cv-templates-service/pkg/containers"
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

	templateServer := handlers.NewTemplateServiceServer(dependencies.DefaultTemplateService, dependencies.MinioClient, dependencies.Logger)
	templates.RegisterTemplateServiceServer(grpcServer, templateServer)

	log.Printf("gRPC server is running on port %s", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
