package handlers

import (
	"go.uber.org/zap"
	certificates "information-service/internal/certificates/grpc"
	"information-service/internal/certificates/services"
)

type CertificateServiceServer struct {
	certificates.UnimplementedCertificatesServiceServer
	certificateService *services.CertificateService
	logger             *zap.Logger
}

func NewCertificateServiceServer(certificateService *services.CertificateService, logger *zap.Logger) *CertificateServiceServer {
	return &CertificateServiceServer{
		certificateService: certificateService,
		logger:             logger,
	}
}
