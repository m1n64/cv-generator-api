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

func (r *informationRepository) CreateOrUpdateInformation(information *models.Information) error {
	return nil
}

func (r *informationRepository) GetInformationByCvID(cvID uuid.UUID) (*models.Information, error) {
	return nil, nil
}
