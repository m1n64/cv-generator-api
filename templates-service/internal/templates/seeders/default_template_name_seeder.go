package seeders

import (
	"cv-templates-service/internal/templates/services"
	"cv-templates-service/pkg/utils"
)

type defaultTemplateNameSeeder struct {
	defaultTemplateService *services.DefaultTemplateService
}

func NewDefaultTemplateNameSeeder(defaultTemplateService *services.DefaultTemplateService) utils.Seeder {
	return &defaultTemplateNameSeeder{
		defaultTemplateService: defaultTemplateService,
	}
}

func (s *defaultTemplateNameSeeder) Seed() error {
	return s.defaultTemplateService.UpdateName("Default")
}
