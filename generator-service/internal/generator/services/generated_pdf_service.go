package services

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"information-service/internal/generator/models"
	"information-service/internal/generator/repositories"
)

type GeneratedPDFService struct {
	generatedPdfRepo repositories.GeneratedPDFRepository
	db               *gorm.DB
}

func NewGeneratedPDFService(generatedPdfRepo repositories.GeneratedPDFRepository, db *gorm.DB) *GeneratedPDFService {
	return &GeneratedPDFService{
		generatedPdfRepo: generatedPdfRepo,
		db:               db,
	}
}

func (s *GeneratedPDFService) CreateGeneratedPDF(cvID uuid.UUID, userID uuid.UUID, title string, fileOrigin string) (*models.GeneratedPDF, error) {
	model := &models.GeneratedPDF{
		CvID:       cvID,
		UserID:     userID,
		Title:      title,
		FileOrigin: fileOrigin,
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
