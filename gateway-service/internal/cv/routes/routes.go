package routes

import (
	middlewares2 "gateway-service/internal/auth/middlewares"
	"gateway-service/internal/cv/handlers"
	"gateway-service/internal/cv/middlewares"
	"gateway-service/pkg/utils"
	"github.com/gin-gonic/gin"
)

func CVRoutes(r *gin.Engine, authMiddleware *middlewares2.AuthMiddleware, cvMiddleware *middlewares.CVMiddleware) {
	cvHandler := handlers.NewCVProxy()

	cvGroup := r.Group("/cv")
	cvGroup.Use(authMiddleware.ValidateToken())
	cvGroup.POST("/", cvHandler.CreateCVHandler)
	cvGroup.GET("/", cvHandler.GetCVListHandler)

	cvWithIDMiddleware := cvGroup.Group("/:cv_id")
	cvWithIDMiddleware.Use(cvMiddleware.GetCVOriginalID())
	cvWithIDMiddleware.GET("/", cvHandler.GetCVHandler)
	cvWithIDMiddleware.POST("/", cvHandler.UpdateCVHandler)
	cvWithIDMiddleware.DELETE("/", cvHandler.DeleteCV)
}

func CVGeneratorRoutes(r *gin.Engine, authMiddleware *middlewares2.AuthMiddleware, cvMiddleware *middlewares.CVMiddleware) {
	cvGeneratorHandler := handlers.NewGeneratorHandler(utils.GetRabbitMQInstance(), utils.GetLogger())

	cvGeneratorGroup := r.Group("/cv/generate/:cv_id")
	cvGeneratorGroup.Use(authMiddleware.ValidateToken())
	cvGeneratorGroup.Use(cvMiddleware.GetCVOriginalID())
	cvGeneratorGroup.POST("/", cvGeneratorHandler.GenerateCV)
}
