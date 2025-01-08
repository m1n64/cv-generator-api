package repositories

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"information-service/internal/languages/models"
)

type languageGormRepository struct {
	db *gorm.DB
}

func NewLanguageGormRepository(db *gorm.DB) LanguageRepository {
	return &languageGormRepository{
		db: db,
	}
}

func (r *languageGormRepository) CreateLanguage(language *models.Language) (*models.Language, error) {
	if err := r.db.Create(language).Error; err != nil {
		return nil, err
	}

	return language, nil
}

func (r *languageGormRepository) GetLanguagesByCvID(cvID uuid.UUID) ([]*models.Language, error) {
	var languages []*models.Language
	if err := r.db.Where("cv_id = ?", cvID).Find(&languages).Error; err != nil {
		return nil, err
	}

	return languages, nil
}

func (r *languageGormRepository) GetLanguage(id uuid.UUID, cvID uuid.UUID) (*models.Language, error) {
	var language models.Language
	if err := r.db.Where("id = ? AND cv_id = ?", id, cvID).First(&language).Error; err != nil {
		return nil, err
	}

	return &language, nil
}

func (r *languageGormRepository) UpdateLanguage(id uuid.UUID, language *models.Language) (*models.Language, error) {
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

func (r *languageGormRepository) DeleteLanguageByCvID(id uuid.UUID, cvID uuid.UUID) error {
	return r.db.Where("id = ? AND cv_id = ?", id, cvID).Delete(&models.Language{}).Error
}

func (r *languageGormRepository) DeleteLanguagesByCvID(cvID uuid.UUID) error {
	return r.db.Where("cv_id = ?", cvID).Delete(&models.Language{}).Error
}
