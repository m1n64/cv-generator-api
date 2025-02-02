package handlers

import (
	"context"
	templates "cv-templates-service/internal/templates/grpc"
	"cv-templates-service/internal/templates/services"
	"cv-templates-service/pkg/utils"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sort"
)

type TemplateServiceServer struct {
	templates.UnimplementedTemplateServiceServer
	DefaultTemplateService *services.DefaultTemplateService
	TemplatesService       *services.TemplatesService
	ColorService           *services.ColorService
	minio                  *utils.MinioClient
	logger                 *zap.Logger
}

func NewTemplateServiceServer(defaultTemplateService *services.DefaultTemplateService, templatesService *services.TemplatesService, colorService *services.ColorService, minio *utils.MinioClient, logger *zap.Logger) *TemplateServiceServer {
	return &TemplateServiceServer{
		DefaultTemplateService: defaultTemplateService,
		TemplatesService:       templatesService,
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

func (s *TemplateServiceServer) GetTemplates(ctx context.Context, _ *templates.Empty) (*templates.Templates, error) {
	templatesList, err := s.TemplatesService.GetTemplates()
	if err != nil {
		s.logger.Error("Failed to get templates", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to get templates")
	}

	var templatesResp []*templates.TemplateResponse
	for _, template := range templatesList {
		bytes, err := s.minio.GetFileAsBytes(ctx, template.TemplateOrigin)
		if err != nil {
			s.logger.Error("Failed to get template from minio", zap.Error(err))
			return nil, status.Error(codes.Internal, "Failed to get template")
		}

		templatesResp = append(templatesResp, &templates.TemplateResponse{
			Id:       template.ID.String(),
			Template: string(bytes),
			Title:    template.Title,
		})
	}

	return &templates.Templates{Templates: templatesResp}, nil
}

func (s *TemplateServiceServer) GetTemplateById(ctx context.Context, req *templates.TemplateByIdRequest) (*templates.TemplateResponse, error) {
	if uuid.Validate(req.Id) != nil {
		return nil, status.Error(codes.InvalidArgument, "Id is required")
	}

	template, err := s.TemplatesService.GetTemplate(uuid.MustParse(req.Id))
	if err != nil {
		s.logger.Error("Failed to get template", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to get template")
	}

	bytes, err := s.minio.GetFileAsBytes(ctx, template.TemplateOrigin)
	if err != nil {
		s.logger.Error("Failed to get template from minio", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to get template")
	}

	return &templates.TemplateResponse{
		Id:       template.ID.String(),
		Template: string(bytes),
		Title:    template.Title,
	}, nil
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
