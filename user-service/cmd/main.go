package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"time"
	health "user-service/internal/health/grpc"
	handlers2 "user-service/internal/health/handlers"
	"user-service/internal/users/grpc/auth"
	"user-service/internal/users/handlers"
	"user-service/internal/users/repositories"
	"user-service/internal/users/services"
	"user-service/internal/users/workers"
	"user-service/pkg/utils"
)

func main() {
	fmt.Println("User service started!")

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
	_, redis := utils.GetRedisConn()

	userRepo := repositories.NewUserGormRepository(db)
	tokenRepo := repositories.NewTokenGormRepository(db)

	authService := services.NewAuthService(userRepo, tokenRepo, db)

	go workers.StartRemoveExpiredTokensWorker(tokenRepo, 24*time.Hour)

	grpcServer := grpc.NewServer()
	authServiceServer := handlers.NewAuthServiceServer(authService)
	auth.RegisterAuthServiceServer(grpcServer, authServiceServer)

	healthServiceServer := handlers2.NewHealthServiceServer(db, redis)
	health.RegisterHealthServiceServer(grpcServer, healthServiceServer)

	log.Printf("gRPC server is running on port %s", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
