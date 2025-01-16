package main

import (
	"cv-service/internal/cv/grpc/cv"
	"cv-service/internal/cv/handlers"
	"cv-service/internal/cv/repositories"
	"cv-service/internal/cv/service"
	health "cv-service/internal/health/grpc"
	handlers2 "cv-service/internal/health/handlers"
	"cv-service/pkg/utils"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

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

	grpcServer := grpc.NewServer()
	cvServiceServer := handlers.NewCVServiceServer(cvService)
	cv.RegisterCVServiceServer(grpcServer, cvServiceServer)

	healthServiceServer := handlers2.NewHealthServiceServer(db, redisConnection)
	health.RegisterHealthServiceServer(grpcServer, healthServiceServer)

	log.Printf("gRPC server is running on port %s", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
