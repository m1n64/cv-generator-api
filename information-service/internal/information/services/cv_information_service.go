package services

import (
	"gorm.io/gorm"
	"information-service/internal/information/repositories"
)

type CVInformationService struct {
	informationRepo repositories.InformationRepository
	db              *gorm.DB
}

func NewCVInformationService(informationRepo repositories.InformationRepository, db *gorm.DB) CVInformationService {
	return CVInformationService{
		informationRepo: informationRepo,
		db:              db,
	}
}
