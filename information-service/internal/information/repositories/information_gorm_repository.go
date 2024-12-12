package repositories

import (
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

func (r *informationRepository) CreateOrUpdateInformation(information *models.Information, cvID *uuid.UUID) error {
	if cvID != nil {
		information.CvID = *cvID
	}

	return nil
}

func (r *informationRepository) GetInformationByCvID(cvID uuid.UUID) (*models.Information, error) {
	return nil, nil
}

func (r *informationRepository) DeleteInformationByCvID(cvID uuid.UUID) error {
	return nil
}
