package services

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"information-service/internal/languages/models"
	repositories2 "information-service/internal/languages/repositories"
	"information-service/pkg/utils"
)

type LanguageService struct {
	languageRepo repositories2.LanguageRepository
	db           *gorm.DB
}

func NewLanguageService(languageRepo repositories2.LanguageRepository, db *gorm.DB) *LanguageService {
	return &LanguageService{
		languageRepo: languageRepo,
		db:           db,
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
	model := &models.Language{
		CvID:  cvID,
		Name:  name,
		Level: level,
	}

	var lang *models.Language

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		lang, err = s.languageRepo.CreateLanguage(model)
		if err != nil {
			return err
		}

		message := utils.LanguageAnalyticQueueMessage{
			LangID:   lang.ID,
			CvID:     lang.CvID,
			Action:   "lang_create",
			DateTime: lang.CreatedAt,
			Detail:   "",
			Language: lang.Name,
			Level:    lang.Level,
		}

		body, err := json.Marshal(message)
		if err == nil {
			rabbit := utils.GetRabbitMQInstance()
			err = rabbit.Channel.Publish(
				"",
				utils.AnalyticQueueName,
				false,
				false,
				amqp.Publishing{
					ContentType: "application/json",
					Body:        body,
				},
			)

			if err != nil {
				utils.GetLogger().Error(fmt.Sprintf("Error publishing message: %v", err))
			}
		}

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return lang, nil
}

func (s *LanguageService) UpdateLanguage(langID uuid.UUID, cvID uuid.UUID, name string, level string) (*models.Language, error) {
	model := &models.Language{
		CvID:  cvID,
		Name:  name,
		Level: level,
	}

	var lang *models.Language

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		lang, err = s.languageRepo.UpdateLanguage(langID, model)
		if err != nil {
			return err
		}

		message := utils.LanguageAnalyticQueueMessage{
			LangID:   lang.ID,
			CvID:     lang.CvID,
			Action:   "lang_update",
			DateTime: lang.UpdatedAt,
			Detail:   "",
			Language: lang.Name,
			Level:    lang.Level,
		}

		body, err := json.Marshal(message)
		if err == nil {
			rabbit := utils.GetRabbitMQInstance()
			err = rabbit.Channel.Publish(
				"",
				utils.AnalyticQueueName,
				false,
				false,
				amqp.Publishing{
					ContentType: "application/json",
					Body:        body,
				},
			)

			if err != nil {
				utils.GetLogger().Error(fmt.Sprintf("Error publishing message: %v", err))
			}
		}

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
