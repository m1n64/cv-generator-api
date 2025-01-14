package routes

import (
	"gateway-service/internal/auth/middlewares"
	middlewares2 "gateway-service/internal/cv/middlewares"
	skills "gateway-service/internal/information/skills/grpc"
	"gateway-service/internal/information/skills/handlers"
	"github.com/gin-gonic/gin"
)

func CVSkillsRoutes(r *gin.Engine, authMiddleware *middlewares.AuthMiddleware, cvMiddleware *middlewares2.CVMiddleware, skillsClient skills.SkillsServiceClient) {
	skillHandler := handlers.NewSkillsProxy(skillsClient)

	skillGroup := r.Group("/information/skills/:cv_id")
	skillGroup.Use(authMiddleware.ValidateToken())
	skillGroup.Use(cvMiddleware.GetCVOriginalID())
	skillGroup.GET("/", skillHandler.GetSkills)
	skillGroup.POST("/", skillHandler.CreateSkill)

	skillIDGroup := skillGroup.Group("/:id")
	skillIDGroup.GET("/", skillHandler.GetSkill)
	skillIDGroup.POST("/", skillHandler.UpdateSkill)
	skillIDGroup.DELETE("/", skillHandler.DeleteSkill)
}
