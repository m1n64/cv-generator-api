package repositories

import (
	"cv-service/internal/cv/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type cvRepository struct {
	db *gorm.DB
}

func NewCVRepository(db *gorm.DB) CVRepository {
	return &cvRepository{
		db: db,
	}
}

func (r *cvRepository) CreateCV(cv *models.CV) error {
	return r.db.Create(cv).Error
}

func (r *cvRepository) GetAllCVsByUserID(userID uuid.UUID) ([]models.CV, error) {
	var cvs []models.CV

	if err := r.db.Where("user_id = ?", userID).Find(&cvs).Error; err != nil {
		return nil, err
	}

	return cvs, nil
}

func (r *cvRepository) GetCVByID(ID uuid.UUID) (*models.CV, error) {
	var cv models.CV

	if err := r.db.Where("id = ?", ID).First(&cv).Error; err != nil {
		return nil, err
	}

	return &cv, nil
}

func (r *cvRepository) DeleteCVByID(ID uuid.UUID) error {
	if err := r.db.Where("id = ?", ID).Delete(&models.CV{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *cvRepository) UpdateCVByID(ID uuid.UUID, updatedCV *models.CV) error {
	if err := r.db.Model(&models.CV{}).Where("id = ?", ID).Updates(updatedCV).Error; err != nil {
		return err
	}

	return nil
}

func (r *cvRepository) GetOriginalIDByExternalID(externalID uuid.UUID, userID uuid.UUID) (uuid.UUID, error) {
	var cv models.CV

	if err := r.db.Where("external_id = ? AND user_id = ?", externalID, userID).First(&cv).Error; err != nil {
		return uuid.Nil, err
	}

	return cv.ID, nil
}
