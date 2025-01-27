package routes

import (
	"gateway-service/internal/auth/handlers"
	"gateway-service/internal/auth/middlewares"
	"gateway-service/internal/users/grpc/auth"
	"gateway-service/pkg/utils"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, authMiddleware *middlewares.AuthMiddleware, client auth.AuthServiceClient, aesEncryptor *utils.AESEncryptor) {
	authHandler := handlers.NewUserProxy(client, aesEncryptor)

	authGroup := r.Group("/auth")
	authGroup.POST("/register", authHandler.RegisterHandler)
	authGroup.POST("/login", authHandler.LoginHandler)
	authGroup.DELETE("/logout", authHandler.LogoutHandler)

	userGroup := r.Group("/user")
	userGroup.Use(authMiddleware.ValidateToken())
	userGroup.GET("/info", authHandler.UserInfoHandler)
}
