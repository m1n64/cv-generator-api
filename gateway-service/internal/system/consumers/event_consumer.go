package consumers

import (
	"gateway-service/internal/system/entities"
	"gateway-service/internal/system/enums"
	"gateway-service/pkg/utils"
	"github.com/goccy/go-json"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

type eventConsumer struct {
	logger           *zap.Logger
	webSocketManager *utils.WebSocketPrivateManager
}

func NewEventConsumer(logger *zap.Logger, webSocketManager *utils.WebSocketPrivateManager) utils.Consumer {
	return &eventConsumer{
		logger:           logger,
		webSocketManager: webSocketManager,
	}
}

func (c *eventConsumer) Handle(msg amqp.Delivery) {
	var notification entities.NotificationEvent

	err := json.Unmarshal(msg.Body, &notification)
	if err != nil {
		c.logger.Error("Error unmarshalling message", zap.Error(err))
		return
	}

	var message string

	if notification.Type == enums.TypeSuccess {
		message = utils.Translate("pdf.generated.success")
	} else {
		message = utils.Translate("pdf.generated.error")
	}

	notification.Message = utils.Translate(message)

	c.webSocketManager.BroadcastToUser(notification.UserID, notification)
}
