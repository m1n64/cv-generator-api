package repositories

import "cv-templates-service/internal/templates/models"

type TemplateRepository interface {
	CreateDefaultTemplate(template *models.Template) (*models.Template, error)
	GetDefaultTemplate() (*models.Template, error)
}
