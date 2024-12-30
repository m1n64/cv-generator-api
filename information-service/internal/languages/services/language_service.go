package services

import (
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"information-service/internal/languages/models"
	repositories2 "information-service/internal/languages/repositories"
)

type LanguageService struct {
	languageRepo             repositories2.LanguageRepository
	languageAnalyticsService *LanguageAnalyticsService
	db                       *gorm.DB
}

func NewLanguageService(languageRepo repositories2.LanguageRepository, languageAnalyticsService *LanguageAnalyticsService, db *gorm.DB) *LanguageService {
	return &LanguageService{
		languageRepo:             languageRepo,
		languageAnalyticsService: languageAnalyticsService,
		db:                       db,
	}
}

func (s *LanguageService) GetLanguagesByCvID(cvID uuid.UUID) ([]*models.Language, error) {
	langList, err := s.languageRepo.GetLanguagesByCvID(cvID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return langList, nil
}

func (s *LanguageService) GetLanguage(id uuid.UUID, cvID uuid.UUID) (*models.Language, error) {
	language, err := s.languageRepo.GetLanguage(id, cvID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return language, nil
}

func (s *LanguageService) CreateLanguage(cvID uuid.UUID, name string, level string) (*models.Language, error) {
	lang := &models.Language{
		CvID:  cvID,
		Name:  name,
		Level: level,
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		lang, err = s.languageRepo.CreateLanguage(lang)
		if err != nil {
			return err
		}

		s.languageAnalyticsService.SendCreateEvent(lang)

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return lang, nil
}

func (s *LanguageService) UpdateLanguage(langID uuid.UUID, cvID uuid.UUID, name string, level string) (*models.Language, error) {
	lang := &models.Language{
		CvID:  cvID,
		Name:  name,
		Level: level,
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		lang, err = s.languageRepo.UpdateLanguage(langID, lang)
		if err != nil {
			return err
		}

		s.languageAnalyticsService.SendUpdateEvent(lang)

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return lang, nil
}

func (s *LanguageService) DeleteLanguage(langID uuid.UUID, cvID uuid.UUID) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		return s.languageRepo.DeleteLanguageByCvID(langID, cvID)
	})

	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

func (s *LanguageService) DeleteLanguagesByCvID(cvID uuid.UUID) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		return s.languageRepo.DeleteLanguagesByCvID(cvID)
	})

	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}
