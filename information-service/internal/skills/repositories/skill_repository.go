package repositories

import (
	"github.com/google/uuid"
	"information-service/internal/skills/models"
)

type SkillRepository interface {
	CreateSkill(skill *models.Skill) (*models.Skill, error)
	GetSkill(id uuid.UUID, cvID uuid.UUID) (*models.Skill, error)
	GetSkillsByCvID(cvID uuid.UUID) ([]*models.Skill, error)
	UpdateSkill(id uuid.UUID, skill *models.Skill) (*models.Skill, error)
	DeleteSkill(id uuid.UUID, cvID uuid.UUID) error
	DeleteSkillsByCvID(cvID uuid.UUID) error
}
