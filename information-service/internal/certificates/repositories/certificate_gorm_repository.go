package repositories

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"information-service/internal/certificates/models"
)

type certificateGormRepository struct {
	db *gorm.DB
}

func NewCertificateGormRepository(db *gorm.DB) CertificateRepository {
	return &certificateGormRepository{
		db: db,
	}
}

func (r *certificateGormRepository) GetCertificates(cvID uuid.UUID) ([]*models.Certificate, error) {
	var certificates []*models.Certificate
	if err := r.db.Where("cv_id = ?", cvID).Order("start_date DESC").Find(&certificates).Error; err != nil {
		return nil, err
	}
	return certificates, nil
}

func (r *certificateGormRepository) GetCertificate(id uuid.UUID, cvID uuid.UUID) (*models.Certificate, error) {
	var certificate models.Certificate
	if err := r.db.Where("id = ? AND cv_id = ?", id, cvID).First(&certificate).Error; err != nil {
		return nil, err
	}
	return &certificate, nil
}

func (r *certificateGormRepository) CreateCertificate(certificate *models.Certificate) (*models.Certificate, error) {
	if err := r.db.Create(&certificate).Error; err != nil {
		return nil, err
	}
	return certificate, nil
}

func (r *certificateGormRepository) UpdateCertificate(id uuid.UUID, certificate *models.Certificate) (*models.Certificate, error) {
	var existingCertificate models.Certificate
	if err := r.db.First(&existingCertificate, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	if err := r.db.Model(&existingCertificate).Updates(certificate).Error; err != nil {
		return nil, err
	}

	return &existingCertificate, nil
}

func (r *certificateGormRepository) DeleteCertificate(id uuid.UUID, cvID uuid.UUID) error {
	return r.db.Where("id = ? AND cv_id = ?", id, cvID).Delete(&models.Certificate{}).Error
}

func (r *certificateGormRepository) DeleteCertificatesByCvID(cvID uuid.UUID) error {
	return r.db.Where("cv_id = ?", cvID).Delete(&models.Certificate{}).Error
}
