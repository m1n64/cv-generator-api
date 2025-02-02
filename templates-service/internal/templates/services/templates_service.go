package services

import (
	"cv-templates-service/internal/templates/models"
	"cv-templates-service/internal/templates/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TemplatesService struct {
	templateRepo repositories.TemplateRepository
	db           *gorm.DB
}

func NewTemplatesService(templateRepo repositories.TemplateRepository, db *gorm.DB) *TemplatesService {
	return &TemplatesService{
		templateRepo: templateRepo,
		db:           db,
	}
}

func (s *TemplatesService) GetTemplates() ([]*models.Template, error) {
	return s.templateRepo.GetTemplates()
}

func (s *TemplatesService) GetTemplate(id uuid.UUID) (*models.Template, error) {
	return s.templateRepo.GetTemplate(id)
}

func (s *TemplatesService) CreateTemplate(filePath string, title string, isPremium bool) (*models.Template, error) {
	template := &models.Template{
		TemplateOrigin: filePath,
		Title:          title,
		IsPremium:      isPremium,
		IsDefault:      false,
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		template, err = s.templateRepo.CreateTemplate(template)

		return err
	})

	return template, err
}
