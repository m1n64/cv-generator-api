package repositories

import (
	"github.com/google/uuid"
	"information-service/internal/languages/models"
)

type LanguageRepository interface {
	CreateLanguage(language *models.Language) (*models.Language, error)
	GetLanguagesByCvID(cvID uuid.UUID) ([]*models.Language, error)
	GetLanguage(id uuid.UUID, cvID uuid.UUID) (*models.Language, error)
	UpdateLanguage(id uuid.UUID, language *models.Language) (*models.Language, error)
	DeleteLanguageByCvID(id uuid.UUID, cvID uuid.UUID) error
	DeleteLanguagesByCvID(cvID uuid.UUID) error
}
