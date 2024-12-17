package repositories

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"information-service/internal/languages/models"
)

type languageRepository struct {
	db *gorm.DB
}

func NewLanguageRepository(db *gorm.DB) LanguageRepository {
	return &languageRepository{
		db: db,
	}
}

func (r *languageRepository) CreateLanguage(language *models.Language) (*models.Language, error) {
	if err := r.db.Create(language).Error; err != nil {
		return nil, err
	}

	return language, nil
}

func (r *languageRepository) GetLanguagesByCvID(cvID uuid.UUID) ([]*models.Language, error) {
	var languages []*models.Language
	if err := r.db.Where("cv_id = ?", cvID).Find(&languages).Error; err != nil {
		return nil, err
	}

	return languages, nil
}

func (r *languageRepository) GetLanguage(id uuid.UUID, cvID uuid.UUID) (*models.Language, error) {
	var language models.Language
	if err := r.db.Where("id = ? AND cv_id = ?", id, cvID).First(&language).Error; err != nil {
		return nil, err
	}

	return &language, nil
}

func (r *languageRepository) UpdateLanguage(id uuid.UUID, language *models.Language) (*models.Language, error) {
	var existingLanguage models.Language
	if err := r.db.First(&existingLanguage, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	if err := r.db.Model(&existingLanguage).Updates(language).Error; err != nil {
		return nil, err
	}

	return &existingLanguage, nil
}

func (r *languageRepository) DeleteLanguageByCvID(id uuid.UUID, cvID uuid.UUID) error {
	return r.db.Where("id = ? AND cv_id = ?", id, cvID).Delete(&models.Language{}).Error
}
