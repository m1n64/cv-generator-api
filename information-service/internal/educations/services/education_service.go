package services

import (
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"information-service/internal/educations/models"
	"information-service/internal/educations/repositories"
	"time"
)

type EducationService struct {
	educationRepository repositories.EducationRepository
	db                  *gorm.DB
}

func NewEducationService(educationRepository repositories.EducationRepository, db *gorm.DB) *EducationService {
	return &EducationService{
		educationRepository: educationRepository,
		db:                  db,
	}
}

func (s *EducationService) CreateEducation(cvID uuid.UUID, institution string, location string, faculty string, startDate time.Time, endDate *time.Time, degree *string, description *string) (*models.Education, error) {
	education := &models.Education{
		CvID:        cvID,
		Institution: institution,
		Location:    location,
		Faculty:     faculty,
		Degree:      degree,
		StartDate:   startDate,
		EndDate:     endDate,
		Description: description,
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		education, err = s.educationRepository.CreateEducation(education)

		return err
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return education, nil
}

func (s *EducationService) GetEducationsByCvID(cvID uuid.UUID) ([]*models.Education, error) {
	educations, err := s.educationRepository.GetEducationsByCvID(cvID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return educations, nil
}

func (s *EducationService) GetEducation(id uuid.UUID, cvID uuid.UUID) (*models.Education, error) {
	education, err := s.educationRepository.GetEducation(id, cvID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return education, nil
}

func (s *EducationService) UpdateEducation(id uuid.UUID, cvID uuid.UUID, institution string, location string, faculty string, startDate time.Time, endDate *time.Time, degree *string, description *string) (*models.Education, error) {
	education := &models.Education{
		ID:          id,
		CvID:        cvID,
		Institution: institution,
		Location:    location,
		Faculty:     faculty,
		Degree:      degree,
		StartDate:   startDate,
		EndDate:     endDate,
		Description: description,
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		education, err = s.educationRepository.UpdateEducation(id, education)

		return err
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return education, nil
}

func (s *EducationService) DeleteEducation(id uuid.UUID, cvID uuid.UUID) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		return s.educationRepository.DeleteEducation(id, cvID)
	})

	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

func (s *EducationService) DeleteEducationsByCvID(cvID uuid.UUID) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		return s.educationRepository.DeleteEducationsByCvID(cvID)
	})

	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}
