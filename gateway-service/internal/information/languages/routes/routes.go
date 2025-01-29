package routes

import (
	"gateway-service/internal/auth/middlewares"
	middlewares2 "gateway-service/internal/cv/middlewares"
	languages "gateway-service/internal/information/languages/grpc"
	"gateway-service/internal/information/languages/handlers"
	"github.com/gin-gonic/gin"
)

func CVLanguagesRoutes(r *gin.Engine, authMiddleware *middlewares.AuthMiddleware, cvMiddleware *middlewares2.CVMiddleware, languagesClient languages.LanguagesServiceClient) {
	languageHandler := handlers.NewLanguagesProxy(languagesClient)

	languagesGroup := r.Group("/information/languages/:cv_id")
	languagesGroup.Use(authMiddleware.ValidateToken())
	languagesGroup.Use(cvMiddleware.GetCVOriginalID())
	languagesGroup.GET("", languageHandler.GetLanguages)
	languagesGroup.POST("", languageHandler.CreateLanguage)

	languagesIDGroup := languagesGroup.Group("/:id")
	languagesIDGroup.GET("", languageHandler.GetLanguage)
	languagesIDGroup.POST("", languageHandler.UpdateLanguage)
	languagesIDGroup.DELETE("", languageHandler.DeleteLanguage)
}
