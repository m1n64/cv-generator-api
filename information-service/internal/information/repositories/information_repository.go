package repositories

import (
	"github.com/google/uuid"
	"information-service/internal/information/models"
)

type InformationRepository interface {
	CreateOrUpdateInformation(information *models.Information) error
	GetInformationByCvID(cvID uuid.UUID) (*models.Information, error)
}
