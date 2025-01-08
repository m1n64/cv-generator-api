package repositories

import (
	"github.com/google/uuid"
	"information-service/internal/generator/models"
)

type GeneratedPDFRepository interface {
	CreateGeneratedPDF(pdf *models.GeneratedPDF) (*models.GeneratedPDF, error)
	GetGeneratedPDFsByCvID(userID uuid.UUID, cvID uuid.UUID) ([]*models.GeneratedPDF, error)
	GetUserGeneratedPDFs(userID uuid.UUID) ([]*models.GeneratedPDF, error)
	DeleteGeneratedPDF(userID uuid.UUID, cvID uuid.UUID) error
}
