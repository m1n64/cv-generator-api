package services

import (
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"information-service/internal/information/models"
	"information-service/internal/information/repositories"
)

type CVInformationService struct {
	informationRepo repositories.InformationRepository
	db              *gorm.DB
}

func NewCVInformationService(informationRepo repositories.InformationRepository, db *gorm.DB) *CVInformationService {
	return &CVInformationService{
		informationRepo: informationRepo,
		db:              db,
	}
}

func (s *CVInformationService) CreateOrUpdateCV(cvID uuid.UUID, fullName string, photoFileID *string, position *string, location *string, biography *string) (*models.Information, error) {
	model := &models.Information{
		CvID:        cvID,
		FullName:    fullName,
		PhotoFileID: photoFileID,
		Position:    position,
		Location:    location,
		Biography:   biography,
	}

	var info *models.Information
	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		info, err = s.informationRepo.CreateOrUpdateInformation(model)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return info, nil
}

func (s *CVInformationService) GetCVInformation(cvID uuid.UUID) (*models.Information, error) {
	info, err := s.informationRepo.GetInformationByCvID(cvID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return info, nil
}

func (s *CVInformationService) DeleteInformation(cvID uuid.UUID) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		err := s.informationRepo.DeleteInformationByCvID(cvID)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}
