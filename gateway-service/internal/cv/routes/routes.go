package routes

import (
	middlewares2 "gateway-service/internal/auth/middlewares"
	"gateway-service/internal/cv/handlers"
	"gateway-service/internal/cv/middlewares"
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
