package repositories

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"information-service/internal/skills/models"
)

type skillRepository struct {
	db *gorm.DB
}

func NewSkillRepository(db *gorm.DB) SkillRepository {
	return &skillRepository{
		db: db,
	}
}

func (r *skillRepository) CreateSkill(skill *models.Skill) (*models.Skill, error) {
	if err := r.db.Create(skill).Error; err != nil {
		return nil, err
	}
	return skill, nil
}

func (r *skillRepository) GetSkillsByCvID(cvID uuid.UUID) ([]*models.Skill, error) {
	var skills []*models.Skill
	if err := r.db.Where("cv_id = ?", cvID).Find(&skills).Error; err != nil {
		return nil, err
	}

	return skills, nil
}

func (r *skillRepository) GetSkill(id uuid.UUID, cvID uuid.UUID) (*models.Skill, error) {
	var skill models.Skill
	if err := r.db.Where("id = ? AND cv_id = ?", id, cvID).First(&skill).Error; err != nil {
		return nil, err
	}

	return &skill, nil
}

func (r *skillRepository) UpdateSkill(id uuid.UUID, skill *models.Skill) (*models.Skill, error) {
	var existingSkill models.Skill
	if err := r.db.First(&existingSkill, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	if err := r.db.Model(&existingSkill).Updates(skill).Error; err != nil {
		return nil, err
	}
	return &existingSkill, nil
}

func (r *skillRepository) DeleteSkill(id uuid.UUID, cvID uuid.UUID) error {
	return r.db.Where("id = ? AND cv_id = ?", id, cvID).Delete(&models.Skill{}).Error
}

func (r *skillRepository) DeleteSkillsByCvID(cvID uuid.UUID) error {
	return r.db.Where("cv_id = ?", cvID).Delete(&models.Skill{}).Error
}
