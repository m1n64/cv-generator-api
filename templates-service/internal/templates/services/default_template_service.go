package services

import (
	"cv-templates-service/internal/templates/models"
	"cv-templates-service/internal/templates/repositories"
	"gorm.io/gorm"
)

type DefaultTemplateService struct {
	TemplateRepo repositories.TemplateRepository
	db           *gorm.DB
}

func NewDefaultTemplateService(templateRepo repositories.TemplateRepository, db *gorm.DB) *DefaultTemplateService {
	return &DefaultTemplateService{
		TemplateRepo: templateRepo,
		db:           db,
	}
}

func (s *DefaultTemplateService) CreateDefaultTemplate(fileOrigin string) (*models.Template, error) {
	template := &models.Template{
		TemplateOrigin: fileOrigin,
		IsDefault:      true,
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		template, err = s.TemplateRepo.CreateDefaultTemplate(template)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return template, nil
}

func (s *DefaultTemplateService) GetDefaultTemplate() (*models.Template, error) {
	return s.TemplateRepo.GetDefaultTemplate()
}
