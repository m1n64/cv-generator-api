package routes

import (
	"gateway-service/internal/auth/middlewares"
	middlewares2 "gateway-service/internal/cv/middlewares"
	"gateway-service/internal/information/skills/handlers"
	"github.com/gin-gonic/gin"
)

func CVSkillsRoutes(r *gin.Engine, authMiddleware *middlewares.AuthMiddleware, cvMiddleware *middlewares2.CVMiddleware) {
	skillHandler := handlers.NewSkillsProxy()

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
