package repositories

import (
	"cv-service/internal/cv/models"
	"github.com/google/uuid"
)

type CVRepository interface {
	// Создаёт новое резюме
	CreateCV(cv *models.CV) error

	// Возвращает список всех резюме для пользователя по UserID
	GetAllCVsByUserID(userID uuid.UUID) ([]models.CV, error)

	// Возвращает резюме по ID
	GetCVByID(ID uuid.UUID) (*models.CV, error)

	// Удаляет резюме по ID
	DeleteCVByID(ID uuid.UUID) error

	// Обновляет резюме по ID
	UpdateCVByID(ID uuid.UUID, updatedCV *models.CV) error

	// Получает оригинальный ID (внутренний) по ID
	GetOriginalIDByExternalID(externalID uuid.UUID, userID uuid.UUID) (uuid.UUID, error)
}
