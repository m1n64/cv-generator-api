package repositories

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"information-service/internal/educations/models"
)

type educationGormRepository struct {
	db *gorm.DB
}

func NewEducationGormRepository(db *gorm.DB) EducationRepository {
	return &educationGormRepository{
		db: db,
	}
}

func (r *educationGormRepository) CreateEducation(education *models.Education) (*models.Education, error) {
	if err := r.db.Create(education).Error; err != nil {
		return nil, err
	}

	return education, nil
}

func (r *educationGormRepository) GetEducation(id uuid.UUID, cvID uuid.UUID) (*models.Education, error) {
	var contact models.Education
	if err := r.db.Where("id = ? AND cv_id = ?", id, cvID).First(&contact).Error; err != nil {
		return nil, err
	}

	return &contact, nil
}

func (r *educationGormRepository) GetEducationsByCvID(cvID uuid.UUID) ([]*models.Education, error) {
	var contacts []*models.Education
	if err := r.db.Where("cv_id = ?", cvID).Order("start_date DESC").Find(&contacts).Error; err != nil {
		return nil, err
	}

	return contacts, nil
}

func (r *educationGormRepository) UpdateEducation(id uuid.UUID, education *models.Education) (*models.Education, error) {
	var existingEducation models.Education
	if err := r.db.First(&existingEducation, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	if err := r.db.Model(&existingEducation).Updates(education).Error; err != nil {
		return nil, err
	}

	return education, nil
}

func (r *educationGormRepository) DeleteEducation(id uuid.UUID, cvID uuid.UUID) error {
	return r.db.Where("id = ? AND cv_id = ?", id, cvID).Delete(&models.Education{}).Error
}

func (r *educationGormRepository) DeleteEducationsByCvID(cvID uuid.UUID) error {
	return r.db.Where("cv_id = ?", cvID).Delete(&models.Education{}).Error
}
