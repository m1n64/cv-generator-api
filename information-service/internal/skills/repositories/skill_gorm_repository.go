package repositories

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"information-service/internal/skills/models"
)

type skillGormRepository struct {
	db *gorm.DB
}

func NewSkillGormRepository(db *gorm.DB) SkillRepository {
	return &skillGormRepository{
		db: db,
	}
}

func (r *skillGormRepository) CreateSkill(skill *models.Skill) (*models.Skill, error) {
	if err := r.db.Create(skill).Error; err != nil {
		return nil, err
	}
	return skill, nil
}

func (r *skillGormRepository) GetSkillsByCvID(cvID uuid.UUID) ([]*models.Skill, error) {
	var skills []*models.Skill
	if err := r.db.Where("cv_id = ?", cvID).Find(&skills).Error; err != nil {
		return nil, err
	}

	return skills, nil
}

func (r *skillGormRepository) GetSkill(id uuid.UUID, cvID uuid.UUID) (*models.Skill, error) {
	var skill models.Skill
	if err := r.db.Where("id = ? AND cv_id = ?", id, cvID).First(&skill).Error; err != nil {
		return nil, err
	}

	return &skill, nil
}

func (r *skillGormRepository) UpdateSkill(id uuid.UUID, skill *models.Skill) (*models.Skill, error) {
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

func (r *skillGormRepository) DeleteSkill(id uuid.UUID, cvID uuid.UUID) error {
	return r.db.Where("id = ? AND cv_id = ?", id, cvID).Delete(&models.Skill{}).Error
}

func (r *skillGormRepository) DeleteSkillsByCvID(cvID uuid.UUID) error {
	return r.db.Where("cv_id = ?", cvID).Delete(&models.Skill{}).Error
}
