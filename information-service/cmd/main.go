package main

import (
	"fmt"
	"google.golang.org/grpc"
	information "information-service/internal/information/grpc"
	"information-service/internal/information/handlers"
	languages "information-service/internal/languages/grpc"
	handlers2 "information-service/internal/languages/handlers"
	skills "information-service/internal/skills/grpc"
	handlers3 "information-service/internal/skills/handlers"
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
	cvInfoServer := handlers.NewCVInformationServiceServer(dependencies.CVInformationService, dependencies.Logger)
	information.RegisterInformationServiceServer(grpcServer, cvInfoServer)

	langServer := handlers2.NewLanguageServiceServer(dependencies.LanguageService, dependencies.Logger)
	languages.RegisterLanguagesServiceServer(grpcServer, langServer)

	skillsServer := handlers3.NewSkillServiceServer(dependencies.SkillService, dependencies.Logger)
	skills.RegisterSkillsServiceServer(grpcServer, skillsServer)

	log.Printf("gRPC server is running on port %s", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
