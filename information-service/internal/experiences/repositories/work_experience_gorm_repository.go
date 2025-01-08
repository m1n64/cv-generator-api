package repositories

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"information-service/internal/experiences/models"
)

type workExperienceRepository struct {
	db *gorm.DB
}

func NewWorkExperienceRepository(db *gorm.DB) WorkExperienceRepository {
	return &workExperienceRepository{
		db: db,
	}
}

func (r *workExperienceRepository) CreateWorkExperience(experience *models.WorkExperience) (*models.WorkExperience, error) {
	if err := r.db.Create(experience).Error; err != nil {
		return nil, err
	}

	return experience, nil
}

func (r *workExperienceRepository) GetWorkExperience(id uuid.UUID, cvID uuid.UUID) (*models.WorkExperience, error) {
	var experience models.WorkExperience
	if err := r.db.Where("id = ? AND cv_id = ?", id, cvID).First(&experience).Error; err != nil {
		return nil, err
	}

	return &experience, nil
}

func (r *workExperienceRepository) GetWorkExperiencesByCvID(cvID uuid.UUID) ([]*models.WorkExperience, error) {
	var experiences []*models.WorkExperience
	if err := r.db.Where("cv_id = ?", cvID).Order("start_date DESC").Find(&experiences).Error; err != nil {
		return nil, err
	}

	return experiences, nil
}

func (r *workExperienceRepository) UpdateWorkExperience(id uuid.UUID, experience *models.WorkExperience) (*models.WorkExperience, error) {
	var existingExperience models.WorkExperience
	if err := r.db.First(&existingExperience, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	if err := r.db.Model(&existingExperience).Updates(experience).Error; err != nil {
		return nil, err
	}

	return experience, nil
}

func (r *workExperienceRepository) DeleteWorkExperiencesByCvID(cvID uuid.UUID) error {
	return r.db.Where("cv_id = ?", cvID).Delete(&models.WorkExperience{}).Error
}

func (r *workExperienceRepository) DeleteWorkExperience(id uuid.UUID, cvID uuid.UUID) error {
	return r.db.Where("id = ? AND cv_id = ?", id, cvID).Delete(&models.WorkExperience{}).Error
}
