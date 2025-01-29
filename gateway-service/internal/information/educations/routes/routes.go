package routes

import (
	"gateway-service/internal/auth/middlewares"
	middlewares2 "gateway-service/internal/cv/middlewares"
	educations "gateway-service/internal/information/educations/grpc"
	"gateway-service/internal/information/educations/handlers"
	"github.com/gin-gonic/gin"
)

func CVEducationsRoutes(r *gin.Engine, authMiddleware *middlewares.AuthMiddleware, cvMiddleware *middlewares2.CVMiddleware, educationsClient educations.EducationServiceClient) {
	educationHandler := handlers.NewEducationsProxy(educationsClient)

	educationsGroup := r.Group("/information/educations/:cv_id")
	educationsGroup.Use(authMiddleware.ValidateToken())
	educationsGroup.Use(cvMiddleware.GetCVOriginalID())
	educationsGroup.GET("", educationHandler.GetEducations)
	educationsGroup.POST("", educationHandler.CreateEducation)

	educationIdGroup := educationsGroup.Group("/:id")
	educationIdGroup.GET("", educationHandler.GetEducation)
	educationIdGroup.POST("", educationHandler.UpdateEducation)
	educationIdGroup.DELETE("", educationHandler.DeleteEducation)
}
