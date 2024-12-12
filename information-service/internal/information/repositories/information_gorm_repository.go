package repositories

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"information-service/internal/information/models"
)

type informationRepository struct {
	db *gorm.DB
}

func NewInformationRepository(db *gorm.DB) InformationRepository {
	return &informationRepository{
		db: db,
	}
}

func (r *informationRepository) CreateOrUpdateInformation(information *models.Information) (*models.Information, error) {
	var result models.Information

	err := r.db.Where("cv_id = ?", information.CvID).First(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := r.db.Create(information).Error; err != nil {
				return nil, err
			}
			return information, nil
		}

		return nil, err
	}

	if err := r.db.Model(&result).Updates(information).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *informationRepository) GetInformationByCvID(cvID uuid.UUID) (*models.Information, error) {
	var information models.Information

	if err := r.db.Where("cv_id = ?", cvID).First(&information).Error; err != nil {
		return nil, err
	}

	return &information, nil
}

func (r *informationRepository) DeleteInformationByCvID(cvID uuid.UUID) error {
	return r.db.Where("cv_id = ?", cvID).Delete(&models.Information{}).Error
}
