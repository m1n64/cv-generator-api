package handlers

import (
	"context"
	"cv-templates-service/internal/templates/seeders"
	"cv-templates-service/pkg/containers"
	infrastructure "cv-templates-service/pkg/infrastructure/grpc"
	"cv-templates-service/pkg/infrastructure/models"
	"cv-templates-service/pkg/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type SeederServiceServer struct {
	infrastructure.UnimplementedSeederServiceServer
	db      *gorm.DB
	logger  *zap.Logger
	seeders map[string]utils.Seeder
}

func NewSeederServiceServer(dependencies *containers.Dependencies) *SeederServiceServer {
	return &SeederServiceServer{
		db:     dependencies.DB,
		logger: dependencies.Logger,
		seeders: map[string]utils.Seeder{
			"template":     seeders.NewTemplateSeeder(dependencies.DefaultTemplateService, dependencies.MinioClient),
			"default-name": seeders.NewDefaultTemplateNameSeeder(dependencies.DefaultTemplateService),
			"ru-template":  seeders.NewRuLangTemplate(dependencies.TemplateService, dependencies.MinioClient),
		},
	}
}

func (s *SeederServiceServer) SeedByName(ctx context.Context, req *infrastructure.SeedByNameRequest) (*infrastructure.Empty, error) {
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}

	if seeder, ok := s.seeders[req.Name]; ok {
		var count int64
		err := s.db.Model(&models.Migration{}).Where("name = ?", req.Name).Count(&count).Error
		if err != nil {
			return nil, status.Error(codes.Internal, "Error counting migrations: "+err.Error())
		}

		if count > 0 {
			s.logger.Info("Seeder already executed", zap.String("name", req.Name))
			return &infrastructure.Empty{}, nil
		}

		if err := seeder.Seed(); err != nil {
			s.logger.Error("Error executing seeder", zap.String("name", req.Name), zap.Error(err))
			return nil, status.Error(codes.Internal, err.Error())
		}

		newMigration := models.Migration{
			Name: req.Name,
		}

		if err := s.db.Create(&newMigration).Error; err != nil {
			s.logger.Error("Error creating migration", zap.Error(err))
			return nil, status.Error(codes.Internal, "Error creating migration: "+err.Error())
		}

		s.logger.Info("Seeder executed successfully", zap.String("name", req.Name))
		return &infrastructure.Empty{}, nil
	}

	return nil, status.Errorf(codes.NotFound, "Seeder %s not found", req.Name)
}
