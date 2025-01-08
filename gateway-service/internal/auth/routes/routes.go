package routes

import (
	"gateway-service/internal/auth/handlers"
	"gateway-service/internal/auth/middlewares"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, authMiddleware *middlewares.AuthMiddleware) {
	authHandler := handlers.NewUserProxy()

	authGroup := r.Group("/auth")
	authGroup.POST("/register", authHandler.RegisterHandler)
	authGroup.POST("/login", authHandler.LoginHandler)

	userGroup := r.Group("/user")
	userGroup.Use(authMiddleware.ValidateToken())
	userGroup.GET("/info", authHandler.UserInfoHandler)
}
