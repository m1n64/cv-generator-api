package repositories

import (
	"cv-templates-service/internal/templates/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type templateGormRepository struct {
	db *gorm.DB
}

func NewTemplateGormRepository(db *gorm.DB) TemplateRepository {
	return &templateGormRepository{
		db: db,
	}
}

func (r *templateGormRepository) CreateDefaultTemplate(template *models.Template) (*models.Template, error) {
	template.IsDefault = true

	if err := r.db.Create(template).Error; err != nil {
		return nil, err
	}

	return template, nil
}

func (r *templateGormRepository) GetTemplates() ([]*models.Template, error) {
	var templates []*models.Template
	if err := r.db.Find(&templates).Error; err != nil {
		return nil, err
	}

	return templates, nil
}

func (r *templateGormRepository) GetTemplate(id uuid.UUID) (*models.Template, error) {
	var template models.Template
	if err := r.db.Where("id = ?", id).First(&template).Error; err != nil {
		return nil, err
	}

	return &template, nil
}

func (r *templateGormRepository) CreateTemplate(template *models.Template) (*models.Template, error) {
	if err := r.db.Create(template).Error; err != nil {
		return nil, err
	}

	return template, nil
}

func (r *templateGormRepository) GetDefaultTemplate() (*models.Template, error) {
	var template models.Template
	if err := r.db.Where("is_default = ?", true).First(&template).Error; err != nil {
		return nil, err
	}

	return &template, nil
}

func (r *templateGormRepository) UpdateNameDefault(name string) error {
	return r.db.Model(&models.Template{}).Where("is_default = ?", true).Update("title", name).Error
}
