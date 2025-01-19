package main

import (
	"context"
	generator "cv-generator-service/internal/generator/grpc"
	"cv-generator-service/internal/generator/handlers"
	health "cv-generator-service/internal/health/grpc"
	handlers2 "cv-generator-service/internal/health/handlers"
	"cv-generator-service/pkg/containers"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
	"os"
)

func tokenAuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "Missing metadata")
	}

	tokens := md["authorization"]
	if len(tokens) == 0 || tokens[0] != "Bearer "+os.Getenv("GRPC_TOKEN") {
		return nil, grpc.Errorf(codes.Unauthenticated, "Invalid token")
	}

	return handler(ctx, req)
}

func main() {
	fmt.Println("CV PDF Generator service started!")

	dependencies, _ := containers.InitializeDependencies()
	go containers.InitializeQueuesConsumer(dependencies)

	port := os.Getenv("SERVICE_PORT")
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(tokenAuthInterceptor),
	)

	generatorServer := handlers.NewGeneratorServiceServer(dependencies.PdfGeneratorService, dependencies.MinioClient, dependencies.Logger)
	generator.RegisterGeneratorServiceServer(grpcServer, generatorServer)

	healthServer := handlers2.NewHealthServiceServer(dependencies.DB, dependencies.RedisClient)
	health.RegisterHealthServiceServer(grpcServer, healthServer)

	log.Printf("gRPC server is running on port %s", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
