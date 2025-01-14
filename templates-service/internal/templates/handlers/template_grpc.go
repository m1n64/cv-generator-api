package handlers

import (
	"context"
	templates "cv-templates-service/internal/templates/grpc"
	"cv-templates-service/internal/templates/services"
	"cv-templates-service/pkg/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TemplateServiceServer struct {
	templates.UnimplementedTemplateServiceServer
	DefaultTemplateService *services.DefaultTemplateService
	minio                  *utils.MinioClient
	logger                 *zap.Logger
}

func NewTemplateServiceServer(defaultTemplateService *services.DefaultTemplateService, minio *utils.MinioClient, logger *zap.Logger) *TemplateServiceServer {
	return &TemplateServiceServer{
		DefaultTemplateService: defaultTemplateService,
		minio:                  minio,
		logger:                 logger,
	}
}

func (s *TemplateServiceServer) GetDefaultTemplate(ctx context.Context, _ *templates.Empty) (*templates.Template, error) {
	template, err := s.DefaultTemplateService.GetDefaultTemplate()

	if err != nil {
		s.logger.Error("Failed to get default template", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to get default template")
	}

	bytes, err := s.minio.GetFileAsBytes(ctx, template.TemplateOrigin)
	if err != nil {
		s.logger.Error("Failed to get default template from minio", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to get default template")
	}

	return &templates.Template{Template: string(bytes)}, nil
}
