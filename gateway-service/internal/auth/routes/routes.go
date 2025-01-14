package routes

import (
	"gateway-service/internal/auth/handlers"
	"gateway-service/internal/auth/middlewares"
	"gateway-service/internal/users/grpc/auth"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, authMiddleware *middlewares.AuthMiddleware, client auth.AuthServiceClient) {
	authHandler := handlers.NewUserProxy(client)

	authGroup := r.Group("/auth")
	authGroup.POST("/register", authHandler.RegisterHandler)
	authGroup.POST("/login", authHandler.LoginHandler)

	userGroup := r.Group("/user")
	userGroup.Use(authMiddleware.ValidateToken())
	userGroup.GET("/info", authHandler.UserInfoHandler)
}
