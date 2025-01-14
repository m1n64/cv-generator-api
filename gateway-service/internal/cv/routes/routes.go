package routes

import (
	middlewares2 "gateway-service/internal/auth/middlewares"
	"gateway-service/internal/cv/grpc/cv"
	"gateway-service/internal/cv/handlers"
	"gateway-service/internal/cv/middlewares"
	"gateway-service/pkg/container"
	"gateway-service/pkg/utils"
	"github.com/gin-gonic/gin"
)

func CVRoutes(r *gin.Engine, authMiddleware *middlewares2.AuthMiddleware, cvMiddleware *middlewares.CVMiddleware, cvClient cv.CVServiceClient) {
	cvHandler := handlers.NewCVProxy(cvClient)

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

func CVGeneratorRoutes(r *gin.Engine, authMiddleware *middlewares2.AuthMiddleware, cvMiddleware *middlewares.CVMiddleware, grpcClients *container.GrpcConnections) {
	cvGeneratorHandler := handlers.NewGeneratorHandler(
		utils.GetRabbitMQInstance(),
		utils.GetLogger(),
		grpcClients.CVClient,
		grpcClients.MainInfoClient,
		grpcClients.ContactsClient,
		grpcClients.SkillsClient,
		grpcClients.LanguagesClient,
		grpcClients.WorkExperiencesClient,
		grpcClients.EducationsClient,
		grpcClients.CertificatesClient,
		grpcClients.TemplatesClient,
	)

	cvGeneratorGroup := r.Group("/cv/generate/:cv_id")
	cvGeneratorGroup.Use(authMiddleware.ValidateToken())
	cvGeneratorGroup.Use(cvMiddleware.GetCVOriginalID())
	cvGeneratorGroup.POST("/", cvGeneratorHandler.GenerateCV)
}
