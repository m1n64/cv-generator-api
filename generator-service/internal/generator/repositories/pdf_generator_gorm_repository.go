package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"information-service/internal/generator/enums"
	"information-service/internal/generator/models"
)

type pdfGeneratorGormRepository struct {
	db *gorm.DB
}

func NewPdfGeneratorGormRepository(db *gorm.DB) PdfGeneratorRepository {
	return &pdfGeneratorGormRepository{
		db: db,
	}
}

func (r *pdfGeneratorGormRepository) CreateGeneratedPDF(pdf *models.GeneratedPdf) (*models.GeneratedPdf, error) {
	if err := r.db.Create(pdf).Error; err != nil {
		return nil, err
	}

	return pdf, nil
}

func (r *pdfGeneratorGormRepository) UpdateStatus(id uuid.UUID, status enums.StatusType) error {
	return r.db.Model(&models.GeneratedPdf{}).Where("id = ?", id).Update("status", status).Error
}

func (r *pdfGeneratorGormRepository) GetGeneratedPDFsByCvID(userID uuid.UUID, cvID uuid.UUID) ([]*models.GeneratedPdf, error) {
	var pdfs []*models.GeneratedPdf
	if err := r.db.Where("user_id = ? AND cv_id = ?", userID, cvID).Order("created_at DESC").Find(&pdfs).Error; err != nil {
		return nil, err
	}

	return pdfs, nil
}

func (r *pdfGeneratorGormRepository) GetUserGeneratedPDFs(userID uuid.UUID) ([]*models.GeneratedPdf, error) {
	var pdfs []*models.GeneratedPdf
	if err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&pdfs).Error; err != nil {
		return nil, err
	}

	return pdfs, nil
}

func (r *pdfGeneratorGormRepository) GetGeneratedPDF(id uuid.UUID, userID uuid.UUID, cvID uuid.UUID) (*models.GeneratedPdf, error) {
	var pdf *models.GeneratedPdf
	if err := r.db.Where("id = ? AND user_id = ? AND cv_id = ?", id, userID, cvID).First(&pdf).Error; err != nil {
		return nil, err
	}

	return pdf, nil
}

func (r *pdfGeneratorGormRepository) DeleteGeneratedPDF(id uuid.UUID, userID uuid.UUID, cvID uuid.UUID) error {
	return r.db.Where("id = ? AND user_id = ? AND cv_id = ?", id, userID, cvID).Delete(&models.GeneratedPdf{}).Error
}
