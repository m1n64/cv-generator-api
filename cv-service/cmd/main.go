package main

import (
	"cv-service/internal/cv/grpc/cv"
	"cv-service/internal/cv/handlers"
	"cv-service/internal/cv/repositories"
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

	port := os.Getenv("SERVICE_PORT")
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	db := utils.GetDBConnection()

	cvRepo := repositories.NewCVRepository(db)

	grpcServer := grpc.NewServer()
	cvService := handlers.NewCVServiceServer(cvRepo)
	cv.RegisterCVServiceServer(grpcServer, cvService)

	log.Printf("gRPC server is running on port %s", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
