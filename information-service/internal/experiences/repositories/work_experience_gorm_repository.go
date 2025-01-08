package repositories

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"information-service/internal/experiences/models"
)

type workExperienceGormRepository struct {
	db *gorm.DB
}

func NewWorkExperienceGormRepository(db *gorm.DB) WorkExperienceRepository {
	return &workExperienceGormRepository{
		db: db,
	}
}

func (r *workExperienceGormRepository) CreateWorkExperience(experience *models.WorkExperience) (*models.WorkExperience, error) {
	if err := r.db.Create(experience).Error; err != nil {
		return nil, err
	}

	return experience, nil
}

func (r *workExperienceGormRepository) GetWorkExperience(id uuid.UUID, cvID uuid.UUID) (*models.WorkExperience, error) {
	var experience models.WorkExperience
	if err := r.db.Where("id = ? AND cv_id = ?", id, cvID).First(&experience).Error; err != nil {
		return nil, err
	}

	return &experience, nil
}

func (r *workExperienceGormRepository) GetWorkExperiencesByCvID(cvID uuid.UUID) ([]*models.WorkExperience, error) {
	var experiences []*models.WorkExperience
	if err := r.db.Where("cv_id = ?", cvID).Order("start_date DESC").Find(&experiences).Error; err != nil {
		return nil, err
	}

	return experiences, nil
}

func (r *workExperienceGormRepository) UpdateWorkExperience(id uuid.UUID, experience *models.WorkExperience) (*models.WorkExperience, error) {
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

func (r *workExperienceGormRepository) DeleteWorkExperiencesByCvID(cvID uuid.UUID) error {
	return r.db.Where("cv_id = ?", cvID).Delete(&models.WorkExperience{}).Error
}

func (r *workExperienceGormRepository) DeleteWorkExperience(id uuid.UUID, cvID uuid.UUID) error {
	return r.db.Where("id = ? AND cv_id = ?", id, cvID).Delete(&models.WorkExperience{}).Error
}
