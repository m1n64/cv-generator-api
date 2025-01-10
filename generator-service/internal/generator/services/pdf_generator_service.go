package services

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"information-service/internal/generator/enums"
	"information-service/internal/generator/models"
	"information-service/internal/generator/repositories"
)

type PdfGeneratorService struct {
	generatedPdfRepo repositories.PdfGeneratorRepository
	db               *gorm.DB
}

func NewPdfGeneratorService(generatedPdfRepo repositories.PdfGeneratorRepository, db *gorm.DB) *PdfGeneratorService {
	return &PdfGeneratorService{
		generatedPdfRepo: generatedPdfRepo,
		db:               db,
	}
}

func (s *PdfGeneratorService) CreateGeneratedPDF(cvID uuid.UUID, userID uuid.UUID, title string, fileOrigin string, status enums.StatusType) (*models.GeneratedPdf, error) {
	model := &models.GeneratedPdf{
		CvID:       cvID,
		UserID:     userID,
		Title:      title,
		FileOrigin: fileOrigin,
		Status:     status,
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		model, err = s.generatedPdfRepo.CreateGeneratedPDF(model)

		return err
	})
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *PdfGeneratorService) UpdateStatus(id uuid.UUID, status enums.StatusType) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		return s.generatedPdfRepo.UpdateStatus(id, status)
	})
}

func (s *PdfGeneratorService) GetGeneratedPDFsByCvID(userID uuid.UUID, cvID uuid.UUID) ([]*models.GeneratedPdf, error) {
	return s.generatedPdfRepo.GetGeneratedPDFsByCvID(userID, cvID)
}

func (s *PdfGeneratorService) GetUserGeneratedPDFs(userID uuid.UUID) ([]*models.GeneratedPdf, error) {
	return s.generatedPdfRepo.GetUserGeneratedPDFs(userID)
}

func (s *PdfGeneratorService) GetGeneratedPDF(id uuid.UUID, userID uuid.UUID, cvID uuid.UUID) (*models.GeneratedPdf, error) {
	return s.generatedPdfRepo.GetGeneratedPDF(id, userID, cvID)
}

func (s *PdfGeneratorService) DeleteGeneratedPDF(id uuid.UUID, userID uuid.UUID, cvID uuid.UUID) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		return s.generatedPdfRepo.DeleteGeneratedPDF(id, userID, cvID)
	})
}
