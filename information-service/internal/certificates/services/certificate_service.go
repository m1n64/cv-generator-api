package services

import (
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"information-service/internal/certificates/models"
	"information-service/internal/certificates/repositories"
	"time"
)

type CertificateService struct {
	certificateRepo repositories.CertificateRepository
	db              *gorm.DB
}

func NewCertificateService(certificateRepo repositories.CertificateRepository, db *gorm.DB) *CertificateService {
	return &CertificateService{
		certificateRepo: certificateRepo,
		db:              db,
	}
}

func (s *CertificateService) GetCertificates(cvID uuid.UUID) ([]*models.Certificate, error) {
	certificates, err := s.certificateRepo.GetCertificates(cvID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return certificates, nil
}

func (s *CertificateService) GetCertificateByID(id uuid.UUID, cvID uuid.UUID) (*models.Certificate, error) {
	certificate, err := s.certificateRepo.GetCertificate(id, cvID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return certificate, nil
}

func (s *CertificateService) CreateCertificate(cvID uuid.UUID, title string, vendor string, startDate time.Time, endDate *time.Time, description *string) (*models.Certificate, error) {
	certificate := &models.Certificate{
		CvID:        cvID,
		Title:       title,
		Vendor:      vendor,
		StartDate:   startDate,
		EndDate:     endDate,
		Description: description,
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		certificate, err = s.certificateRepo.CreateCertificate(certificate)

		return err
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return certificate, nil
}

func (s *CertificateService) UpdateCertificate(id uuid.UUID, cvID uuid.UUID, title string, vendor string, startDate time.Time, endDate *time.Time, description *string) (*models.Certificate, error) {
	certificate := &models.Certificate{
		ID:          id,
		CvID:        cvID,
		Title:       title,
		Vendor:      vendor,
		StartDate:   startDate,
		EndDate:     endDate,
		Description: description,
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		certificate, err = s.certificateRepo.UpdateCertificate(id, certificate)

		return err
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return certificate, nil
}

func (s *CertificateService) DeleteCertificate(id uuid.UUID, cvID uuid.UUID) error {
	err := s.certificateRepo.DeleteCertificate(id, cvID)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (s *CertificateService) DeleteCertificates(cvID uuid.UUID) error {
	return s.certificateRepo.DeleteCertificatesByCvID(cvID)
}
