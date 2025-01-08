package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"information-service/internal/generator/models"
)

type generatedPDFGormRepository struct {
	db *gorm.DB
}

func NewGeneratedPDFGormRepository(db *gorm.DB) GeneratedPDFRepository {
	return &generatedPDFGormRepository{
		db: db,
	}
}

func (r *generatedPDFGormRepository) CreateGeneratedPDF(pdf *models.GeneratedPDF) (*models.GeneratedPDF, error) {
	if err := r.db.Create(pdf).Error; err != nil {
		return nil, err
	}

	return pdf, nil
}

func (r *generatedPDFGormRepository) GetGeneratedPDFsByCvID(userID uuid.UUID, cvID uuid.UUID) ([]*models.GeneratedPDF, error) {
	var pdfs []*models.GeneratedPDF
	if err := r.db.Where("user_id = ? AND cv_id = ?", userID, cvID).Order("created_at DESC").Find(&pdfs).Error; err != nil {
		return nil, err
	}

	return pdfs, nil
}

func (r *generatedPDFGormRepository) GetUserGeneratedPDFs(userID uuid.UUID) ([]*models.GeneratedPDF, error) {
	var pdfs []*models.GeneratedPDF
	if err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&pdfs).Error; err != nil {
		return nil, err
	}

	return pdfs, nil
}

func (r *generatedPDFGormRepository) DeleteGeneratedPDF(userID uuid.UUID, cvID uuid.UUID) error {
	return r.db.Where("user_id = ? AND cv_id = ?", userID, cvID).Delete(&models.GeneratedPDF{}).Error
}
