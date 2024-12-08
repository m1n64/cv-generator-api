package main

import (
	"fmt"
	"gateway-service/internal/auth/handlers"
	"gateway-service/internal/auth/middlewares"
	"gateway-service/pkg/utils"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	fmt.Println("Gateway service started!")

	utils.InitLogs()
	utils.LoadEnv()

	authHandler := handlers.NewUserProxy()
	authMiddleware := middlewares.NewAuthMiddleware()

	r := gin.Default()

	r.POST("/register", authHandler.RegisterHandler)
	r.POST("/login", authHandler.LoginHandler)

	protected := r.Group("/user")
	protected.Use(authMiddleware.ValidateToken())
	protected.GET("/info", authHandler.UserInfoHandler)

	r.Run(fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT")))
	fmt.Println("Gateway service run successfully!")
}
