package repositories

import (
	"github.com/google/uuid"
	"information-service/internal/contacts/models"
)

type ContactRepository interface {
	CreateContact(contact *models.Contact) (*models.Contact, error)
	GetContact(id uuid.UUID, cvID uuid.UUID) (*models.Contact, error)
	GetContactsByCvID(cvID uuid.UUID) ([]*models.Contact, error)
	UpdateContact(id uuid.UUID, contact *models.Contact) (*models.Contact, error)
	DeleteContact(id uuid.UUID, cvID uuid.UUID) error
	DeleteContactsByCvID(cvID uuid.UUID) error
}
