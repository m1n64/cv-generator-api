package services

import (
	"ai-service/internal/ai/entites"
	"ai-service/internal/ai/enums"
	"ai-service/pkg/utils"
	"github.com/goccy/go-json"
	"github.com/streadway/amqp"
	"time"
)

type AiAnalyticsService struct {
	rabbitMq *utils.RabbitMQConnection
}

func NewAiAnalyticsService(rabbitMq *utils.RabbitMQConnection) *AiAnalyticsService {
	return &AiAnalyticsService{
		rabbitMq: rabbitMq,
	}
}

func (s *AiAnalyticsService) SendAiAnalyticsGenerateEvent(prompt string, response string, service string) error {
	analyticEntity := entites.AiAnalyticsQueueMessage{
		Action:   enums.AiAnalyticsActionGenerate,
		Prompt:   prompt,
		Response: response,
		Service:  service,
		SendAt:   time.Now().Format(time.RFC3339),
	}

	body, err := json.Marshal(analyticEntity)
	if err != nil {
		return err
	}

	err = s.rabbitMq.Channel.Publish(
		"",
		utils.AnalyticQueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	return err
}
