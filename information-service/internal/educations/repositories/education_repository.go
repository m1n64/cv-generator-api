package repositories

import (
	"github.com/google/uuid"
	"information-service/internal/educations/models"
)

type EducationRepository interface {
	CreateEducation(education *models.Education) (*models.Education, error)
	GetEducation(id uuid.UUID, cvID uuid.UUID) (*models.Education, error)
	GetEducationsByCvID(cvID uuid.UUID) ([]*models.Education, error)
	UpdateEducation(id uuid.UUID, education *models.Education) (*models.Education, error)
	DeleteEducation(id uuid.UUID, cvID uuid.UUID) error
	DeleteEducationsByCvID(cvID uuid.UUID) error
}
