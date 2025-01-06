package repositories

import (
	"github.com/google/uuid"
	"information-service/internal/experiences/models"
)

type WorkExperienceRepository interface {
	CreateWorkExperience(experience *models.WorkExperience) (*models.WorkExperience, error)
	GetWorkExperience(id uuid.UUID, cvID uuid.UUID) (*models.WorkExperience, error)
	GetWorkExperiencesByCvID(cvID uuid.UUID) ([]*models.WorkExperience, error)
	UpdateWorkExperience(id uuid.UUID, experience *models.WorkExperience) (*models.WorkExperience, error)
	DeleteWorkExperience(id uuid.UUID, cvID uuid.UUID) error
	DeleteWorkExperiencesByCvID(cvID uuid.UUID) error
}
