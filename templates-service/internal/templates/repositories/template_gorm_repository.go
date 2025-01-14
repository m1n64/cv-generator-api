package repositories

import (
	"cv-templates-service/internal/templates/models"
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

func (r *templateGormRepository) GetDefaultTemplate() (*models.Template, error) {
	var template models.Template
	if err := r.db.Where("is_default = ?", true).First(&template).Error; err != nil {
		return nil, err
	}

	return &template, nil
}
