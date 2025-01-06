package services

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"information-service/internal/skills/models"
	"information-service/pkg/utils"
)

type SkillsAnalyticsService struct {
	rabbitMq *utils.RabbitMQConnection
	logger   *zap.Logger
}

func NewSkillsAnalyticsService(rabbitMq *utils.RabbitMQConnection, logger *zap.Logger) *SkillsAnalyticsService {
	return &SkillsAnalyticsService{
		rabbitMq: rabbitMq,
		logger:   logger,
	}
}

func (s *SkillsAnalyticsService) SendCreateEvent(skill *models.Skill) {
	message := utils.SkillAnalyticQueueMessage{
		StackID:    skill.ID,
		CvID:       skill.CvID,
		Action:     "skill_create",
		DateTime:   skill.CreatedAt,
		Detail:     "",
		Technology: skill.Name,
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

func (s *SkillsAnalyticsService) SendUpdateEvent(skill *models.Skill) {
	message := utils.SkillAnalyticQueueMessage{
		StackID:    skill.ID,
		CvID:       skill.CvID,
		Action:     "skill_update",
		DateTime:   skill.CreatedAt,
		Detail:     "",
		Technology: skill.Name,
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
