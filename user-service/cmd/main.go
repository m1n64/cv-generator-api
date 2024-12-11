package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"user-service/internal/users/grpc/auth"
	"user-service/internal/users/handlers"
	"user-service/internal/users/repositories"
	"user-service/internal/users/services"
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

	port := os.Getenv("SERVICE_PORT")
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	db := utils.GetDBConnection()

	userRepo := repositories.NewUserRepository(db)
	tokenRepo := repositories.NewTokenRepository(db)

	authService := services.NewAuthService(userRepo, tokenRepo)

	grpcServer := grpc.NewServer()
	authServiceServer := handlers.NewAuthServiceServer(authService)
	auth.RegisterAuthServiceServer(grpcServer, authServiceServer)

	log.Printf("gRPC server is running on port %s", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
