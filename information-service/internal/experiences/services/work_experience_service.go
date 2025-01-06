package services

import (
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"information-service/internal/experiences/models"
	"information-service/internal/experiences/repositories"
	"time"
)

type WorkExperienceService struct {
	experienceRepository repositories.WorkExperienceRepository
	db                   *gorm.DB
}

func NewWorkExperienceService(experienceRepository repositories.WorkExperienceRepository, db *gorm.DB) *WorkExperienceService {
	return &WorkExperienceService{
		experienceRepository: experienceRepository,
		db:                   db,
	}
}

func (s *WorkExperienceService) GetWorkExperiencesByCvID(cvID uuid.UUID) ([]*models.WorkExperience, error) {
	expList, err := s.experienceRepository.GetWorkExperiencesByCvID(cvID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return expList, nil
}

func (s *WorkExperienceService) GetWorkExperience(id uuid.UUID, cvID uuid.UUID) (*models.WorkExperience, error) {
	edu, err := s.experienceRepository.GetWorkExperience(id, cvID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return edu, nil
}

func (s *WorkExperienceService) CreateWorkExperience(cvID uuid.UUID, company string, position string, startDate time.Time, endDate *time.Time, location string, description string) (*models.WorkExperience, error) {
	exp := &models.WorkExperience{
		CvID:        cvID,
		Company:     company,
		Position:    position,
		StartDate:   startDate,
		EndDate:     endDate,
		Location:    location,
		Description: description,
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		exp, err = s.experienceRepository.CreateWorkExperience(exp)

		return err
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return exp, nil
}

func (s *WorkExperienceService) UpdateWorkExperience(id uuid.UUID, cvID uuid.UUID, company string, position string, startDate time.Time, endDate *time.Time, location string, description string) (*models.WorkExperience, error) {
	exp := &models.WorkExperience{
		ID:          id,
		CvID:        cvID,
		Company:     company,
		Position:    position,
		StartDate:   startDate,
		EndDate:     endDate,
		Location:    location,
		Description: description,
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		exp, err = s.experienceRepository.UpdateWorkExperience(id, exp)

		return err
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return exp, nil
}

func (s *WorkExperienceService) DeleteWorkExperience(id uuid.UUID, cvID uuid.UUID) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		return s.experienceRepository.DeleteWorkExperience(id, cvID)
	})

	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

func (s *WorkExperienceService) DeleteWorkExperiencesByCvID(cvID uuid.UUID) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		return s.experienceRepository.DeleteWorkExperiencesByCvID(cvID)
	})

	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}
