package repositories

import (
	"cv-generator-service/internal/generator/enums"
	"cv-generator-service/internal/generator/models"
	"github.com/google/uuid"
)

type PdfGeneratorRepository interface {
	CreateGeneratedPDF(pdf *models.GeneratedPdf) (*models.GeneratedPdf, error)
	UpdateStatus(id uuid.UUID, status enums.StatusType) error
	UpdateFileOrigin(id uuid.UUID, fileOrigin *string) error
	UpdateGeneratedPDF(id uuid.UUID, pdf *models.GeneratedPdf) (*models.GeneratedPdf, error)
	GetGeneratedPDFsByCvID(userID uuid.UUID, cvID uuid.UUID) ([]*models.GeneratedPdf, error)
	GetUserGeneratedPDFs(userID uuid.UUID) ([]*models.GeneratedPdf, error)
	GetGeneratedPDF(id uuid.UUID, userID uuid.UUID, cvID uuid.UUID) (*models.GeneratedPdf, error)
	DeleteGeneratedPDF(id uuid.UUID, userID uuid.UUID, cvID uuid.UUID) error
}
