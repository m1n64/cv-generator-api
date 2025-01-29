package routes

import (
	"gateway-service/internal/auth/middlewares"
	middlewares2 "gateway-service/internal/cv/middlewares"
	generator "gateway-service/internal/generator/grpc"
	handlers2 "gateway-service/internal/generator/handlers"
	"github.com/gin-gonic/gin"
)

func GeneratorRoutes(r *gin.Engine, authMiddleware *middlewares.AuthMiddleware, cvMiddleware *middlewares2.CVMiddleware, generatorClient generator.GeneratorServiceClient) {
	generatorHandler := handlers2.NewGeneratorProxy(generatorClient)

	generatorGroup := r.Group("/generator/cv")
	generatorGroup.Use(authMiddleware.ValidateToken())
	generatorGroup.GET("", generatorHandler.GetAllGeneratedPdfs)

	generatorWithIDMiddleware := generatorGroup.Group("/:cv_id")
	generatorWithIDMiddleware.Use(cvMiddleware.GetCVOriginalID())
	generatorWithIDMiddleware.GET("", generatorHandler.GetGeneratedPdfByCV)

	generatorWithGenID := generatorWithIDMiddleware.Group("/:id")
	generatorWithGenID.GET("", generatorHandler.GeneratePdf)
	generatorWithGenID.DELETE("", generatorHandler.DeleteGeneratedPdf)
	generatorWithGenID.GET("/link", generatorHandler.DownloadGeneratedPdf)
}
