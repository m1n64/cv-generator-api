package handlers

import (
	"context"
	templates "cv-templates-service/internal/templates/grpc"
	"cv-templates-service/internal/templates/services"
	"cv-templates-service/pkg/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sort"
)

type TemplateServiceServer struct {
	templates.UnimplementedTemplateServiceServer
	DefaultTemplateService *services.DefaultTemplateService
	ColorService           *services.ColorService
	minio                  *utils.MinioClient
	logger                 *zap.Logger
}

func NewTemplateServiceServer(defaultTemplateService *services.DefaultTemplateService, colorService *services.ColorService, minio *utils.MinioClient, logger *zap.Logger) *TemplateServiceServer {
	return &TemplateServiceServer{
		DefaultTemplateService: defaultTemplateService,
		ColorService:           colorService,
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

func (s *TemplateServiceServer) GetColorScheme(ctx context.Context, _ *templates.Empty) (*templates.ColorScheme, error) {
	colors, err := s.ColorService.GetColorsMap()
	if err != nil {
		s.logger.Error("Failed to get colors", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to get colors")
	}

	var keys []string
	for title := range colors {
		keys = append(keys, title)
	}

	sort.Strings(keys)

	var colorsResp []*templates.Color
	for _, title := range keys {
		colorsResp = append(colorsResp, &templates.Color{
			Title:       title,
			AccentColor: colors[title],
		})
	}

	return &templates.ColorScheme{Colors: colorsResp}, nil
}

func (s *TemplateServiceServer) GetColorSchemeByName(ctx context.Context, req *templates.ColorSchemeByNameRequest) (*templates.Color, error) {
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "Name is required")
	}

	color, err := s.ColorService.GetColor(req.Name)
	if err != nil {
		s.logger.Error("Failed to get color", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to get color")
	}

	return &templates.Color{Title: color.Title, AccentColor: color.AccentColor}, nil
}
