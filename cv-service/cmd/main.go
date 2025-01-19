package main

import (
	"context"
	"cv-service/internal/cv/grpc/cv"
	"cv-service/internal/cv/handlers"
	"cv-service/internal/cv/repositories"
	"cv-service/internal/cv/service"
	health "cv-service/internal/health/grpc"
	handlers2 "cv-service/internal/health/handlers"
	"cv-service/pkg/utils"
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
	fmt.Println("CV service started!")

	utils.InitLogs()
	utils.LoadEnv()
	utils.CreateRedisConn()
	utils.InitDBConnection()
	utils.StartMigrations()
	utils.InitValidator()
	utils.ConnectRabbitMQ()
	utils.InitializeQueues()

	port := os.Getenv("SERVICE_PORT")
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	db := utils.GetDBConnection()

	_, redisConnection := utils.GetRedisConn()

	redisClient := utils.NewRedisAdapter(redisConnection)

	cvRepo := repositories.NewCVGormRepository(db)

	cvService := service.NewCVService(cvRepo, redisClient, db)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(tokenAuthInterceptor),
	)
	cvServiceServer := handlers.NewCVServiceServer(cvService)
	cv.RegisterCVServiceServer(grpcServer, cvServiceServer)

	healthServiceServer := handlers2.NewHealthServiceServer(db, redisConnection)
	health.RegisterHealthServiceServer(grpcServer, healthServiceServer)

	log.Printf("gRPC server is running on port %s", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
