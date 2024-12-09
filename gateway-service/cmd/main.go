package main

import (
	"fmt"
	"gateway-service/internal/auth/handlers"
	"gateway-service/internal/auth/middlewares"
	handlers2 "gateway-service/internal/system/handlers"
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

	r.GET("/ping", handlers2.PingHandler)

	r.GET("/documentation/openapi.json", func(c *gin.Context) {
		c.File("./config/swagger/openapi.json")
	})

	authGroup := r.Group("/auth")
	authGroup.POST("/register", authHandler.RegisterHandler)
	authGroup.POST("/login", authHandler.LoginHandler)

	userGroup := r.Group("/user")
	userGroup.Use(authMiddleware.ValidateToken())
	userGroup.GET("/info", authHandler.UserInfoHandler)

	r.Run(fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT")))
	fmt.Println("Gateway service run successfully!")
}
