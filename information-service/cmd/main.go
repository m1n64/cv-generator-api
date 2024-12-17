package main

import (
	"fmt"
	"google.golang.org/grpc"
	information "information-service/internal/information/grpc"
	"information-service/internal/information/handlers"
	languages "information-service/internal/languages/grpc"
	handlers2 "information-service/internal/languages/handlers"
	"information-service/pkg/containers"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println("CV Information service started!")

	dependencies, _ := containers.InitializeDependencies()
	go containers.InitializeQueuesConsumer(dependencies)

	port := os.Getenv("SERVICE_PORT")
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	cvInfoServer := handlers.NewCVInformationServiceServer(dependencies.CVInformationService)
	information.RegisterInformationServiceServer(grpcServer, cvInfoServer)

	langServer := handlers2.NewLanguageServiceServer(dependencies.LanguageService)
	languages.RegisterLanguagesServiceServer(grpcServer, langServer)

	log.Printf("gRPC server is running on port %s", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
