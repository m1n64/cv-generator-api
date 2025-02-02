package repositories

import (
	"cv-templates-service/internal/templates/models"
	"github.com/google/uuid"
)

type TemplateRepository interface {
	CreateDefaultTemplate(template *models.Template) (*models.Template, error)
	CreateTemplate(template *models.Template) (*models.Template, error)
	GetDefaultTemplate() (*models.Template, error)
	GetTemplates() ([]*models.Template, error)
	GetTemplate(id uuid.UUID) (*models.Template, error)
	UpdateNameDefault(name string) error
}
