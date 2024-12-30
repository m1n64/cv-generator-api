package repositories

import (
	"github.com/google/uuid"
	"information-service/internal/certificates/models"
)

type CertificateRepository interface {
	GetCertificates(cvID uuid.UUID) ([]*models.Certificate, error)
	GetCertificate(id uuid.UUID, cvID uuid.UUID) (*models.Certificate, error)
	CreateCertificate(certificate *models.Certificate) (*models.Certificate, error)
	UpdateCertificate(id uuid.UUID, certificate *models.Certificate) (*models.Certificate, error)
	DeleteCertificate(id uuid.UUID, cvID uuid.UUID) error
	DeleteCertificatesByCvID(cvID uuid.UUID) error
}
