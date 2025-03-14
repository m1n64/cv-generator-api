package routes

import (
	"gateway-service/internal/auth/middlewares"
	middlewares2 "gateway-service/internal/cv/middlewares"
	certificates "gateway-service/internal/information/certificates/grpc"
	"gateway-service/internal/information/certificates/handlers"
	"github.com/gin-gonic/gin"
)

func CVCertificatesRoutes(r *gin.Engine, authMiddleware *middlewares.AuthMiddleware, cvMiddleware *middlewares2.CVMiddleware, certificatesClient certificates.CertificatesServiceClient) {
	certsHandler := handlers.NewCertificatesProxy(certificatesClient)

	certsGroup := r.Group("/information/certificates/:cv_id")

	certsGroup.Use(authMiddleware.ValidateToken())
	certsGroup.Use(cvMiddleware.GetCVOriginalID())
	certsGroup.GET("", certsHandler.GetCertificates)
	certsGroup.POST("", certsHandler.CreateCertificate)

	certsIdGroup := certsGroup.Group("/:id")
	certsIdGroup.GET("", certsHandler.GetCertificate)
	certsIdGroup.POST("", certsHandler.UpdateCertificate)
	certsIdGroup.DELETE("", certsHandler.DeleteCertificate)
}
