package routes

import (
	ai "gateway-service/internal/ai/grpc"
	"gateway-service/internal/ai/handlers"
	"gateway-service/internal/auth/middlewares"
	"github.com/gin-gonic/gin"
)

func AIRoutes(r *gin.Engine, middleware *middlewares.AuthMiddleware, aiClient ai.AiServiceClient) {
	aiHandler := handlers.NewAiProxy(aiClient)

	aiGroup := r.Group("/ai")
	aiGroup.Use(middleware.ValidateToken())
	aiGroup.POST("/generate/:type", aiHandler.Generate)
}
