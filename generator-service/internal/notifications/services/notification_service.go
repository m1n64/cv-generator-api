package services

import (
	"cv-generator-service/internal/notifications/entities"
	"cv-generator-service/internal/notifications/enums"
	"cv-generator-service/pkg/utils"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

type NotificationService struct {
	rabbitmq *utils.RabbitMQConnection
	logger   *zap.Logger
}

func NewNotificationService(rabbitmq *utils.RabbitMQConnection, logger *zap.Logger) *NotificationService {
	return &NotificationService{
		rabbitmq: rabbitmq,
		logger:   logger,
	}
}

func (s *NotificationService) SendSuccess(userID uuid.UUID, message string) {
	event := entities.NotificationEvent{
		Type:    enums.TypeSuccess,
		UserID:  userID.String(),
		Message: message,
	}

	s.sendEvent(event)
}

func (s *NotificationService) SendError(userID uuid.UUID, err error) {
	event := entities.NotificationEvent{
		Type:    enums.TypeError,
		UserID:  userID.String(),
		Message: err.Error(),
	}

	s.sendEvent(event)
}

func (s *NotificationService) sendEvent(event entities.NotificationEvent) {
	body, err := json.Marshal(event)
	if err != nil {
		s.logger.Error("Error marshal success notification", zap.Error(err))
		return
	}

	err = s.rabbitmq.Channel.Publish(
		"",
		utils.GatewayEventsQueue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		s.logger.Error("Error send success notification", zap.Error(err))
		return
	}
}
