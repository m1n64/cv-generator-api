package services

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"information-service/internal/languages/models"
	"information-service/pkg/utils"
)

type LanguageAnalyticsService struct {
	rabbitMq *utils.RabbitMQConnection
	logger   *zap.Logger
}

func NewLanguageAnalyticsService(rabbitMq *utils.RabbitMQConnection, logger *zap.Logger) *LanguageAnalyticsService {
	return &LanguageAnalyticsService{
		rabbitMq: rabbitMq,
		logger:   logger,
	}
}

func (s *LanguageAnalyticsService) SendCreateEvent(lang *models.Language) {
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
		err = s.rabbitMq.Channel.Publish(
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
			s.logger.Error(fmt.Sprintf("Error publishing message: %v", err))
		}
	}
}

func (s *LanguageAnalyticsService) SendUpdateEvent(lang *models.Language) {
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
		err = s.rabbitMq.Channel.Publish(
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
			s.logger.Error(fmt.Sprintf("Error publishing message: %v", err))
		}
	}
}
