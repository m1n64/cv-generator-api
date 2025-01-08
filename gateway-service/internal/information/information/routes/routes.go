package routes

import (
	"gateway-service/internal/auth/middlewares"
	middlewares2 "gateway-service/internal/cv/middlewares"
	handlers4 "gateway-service/internal/information/information/handlers"
	"github.com/gin-gonic/gin"
)

func CVInfoRoutes(r *gin.Engine, authMiddleware *middlewares.AuthMiddleware, cvMiddleware *middlewares2.CVMiddleware) {
	cvInfoHandler := handlers4.NewInformationProxy()

	informationGroup := r.Group("/information/main/:cv_id")
	informationGroup.Use(authMiddleware.ValidateToken())
	informationGroup.Use(cvMiddleware.GetCVOriginalID())
	informationGroup.GET("/", cvInfoHandler.GetCVInformation)
	informationGroup.POST("/", cvInfoHandler.CreateOrUpdateCVInformation)
	informationGroup.DELETE("/", cvInfoHandler.DeleteCVInformation)
}
