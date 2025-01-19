package routes

import (
	"gateway-service/internal/auth/middlewares"
	templates "gateway-service/internal/templates/grpc"
	"gateway-service/internal/templates/handlers"
	"github.com/gin-gonic/gin"
)

func TemplatesRoutes(r *gin.Engine, authMiddleware *middlewares.AuthMiddleware, templatesClient templates.TemplateServiceClient) {
	templatesHandler := handlers.NewTemplateProxyHandler(templatesClient)

	templatesRoutes := r.Group("/templates")

	templatesRoutes.Use(authMiddleware.ValidateToken())
	templatesRoutes.GET("/default", templatesHandler.GetDefaultTemplate)
	templatesRoutes.GET("/colors", templatesHandler.GetColors)
}
