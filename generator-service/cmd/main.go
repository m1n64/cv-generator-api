package main

import (
	"fmt"
	"google.golang.org/grpc"
	generator "information-service/internal/generator/grpc"
	"information-service/internal/generator/handlers"
	"information-service/pkg/containers"
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

	generatorServer := handlers.NewGeneratorServiceServer(dependencies.PdfGeneratorService, dependencies.Logger)
	generator.RegisterGeneratorServiceServer(grpcServer, generatorServer)

	log.Printf("gRPC server is running on port %s", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
