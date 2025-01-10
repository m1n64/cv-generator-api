package repositories

import (
	"github.com/google/uuid"
	"information-service/internal/generator/enums"
	"information-service/internal/generator/models"
)

type PdfGeneratorRepository interface {
	CreateGeneratedPDF(pdf *models.GeneratedPdf) (*models.GeneratedPdf, error)
	UpdateStatus(id uuid.UUID, status enums.StatusType) error
	GetGeneratedPDFsByCvID(userID uuid.UUID, cvID uuid.UUID) ([]*models.GeneratedPdf, error)
	GetUserGeneratedPDFs(userID uuid.UUID) ([]*models.GeneratedPdf, error)
	GetGeneratedPDF(id uuid.UUID, userID uuid.UUID, cvID uuid.UUID) (*models.GeneratedPdf, error)
	DeleteGeneratedPDF(id uuid.UUID, userID uuid.UUID, cvID uuid.UUID) error
}
