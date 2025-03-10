package routes

import (
	"gateway-service/internal/auth/middlewares"
	middlewares2 "gateway-service/internal/cv/middlewares"
	experiences "gateway-service/internal/information/experiences/grpc"
	"gateway-service/internal/information/experiences/handlers"
	"github.com/gin-gonic/gin"
)

func CVExperiencesRoutes(r *gin.Engine, authMiddleware *middlewares.AuthMiddleware, cvMiddleware *middlewares2.CVMiddleware, experiencesClient experiences.ExperiencesServiceClient) {
	experienceHandler := handlers.NewExperienceProxy(experiencesClient)

	experienceGroup := r.Group("/information/experiences/:cv_id")
	experienceGroup.Use(authMiddleware.ValidateToken())
	experienceGroup.Use(cvMiddleware.GetCVOriginalID())
	experienceGroup.GET("", experienceHandler.GetExperiences)
	experienceGroup.POST("", experienceHandler.CreateExperience)

	experienceIdGroup := experienceGroup.Group("/:id")
	experienceIdGroup.GET("", experienceHandler.GetExperience)
	experienceIdGroup.POST("", experienceHandler.UpdateExperience)
	experienceIdGroup.DELETE("", experienceHandler.DeleteExperience)
}
