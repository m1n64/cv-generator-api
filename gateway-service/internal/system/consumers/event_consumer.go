package consumers

import (
	"context"
	"gateway-service/internal/cv/grpc/cv"
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
	cvClient         cv.CVServiceClient
}

func NewEventConsumer(logger *zap.Logger, webSocketManager *utils.WebSocketPrivateManager, cvClient cv.CVServiceClient) utils.Consumer {
	return &eventConsumer{
		logger:           logger,
		webSocketManager: webSocketManager,
		cvClient:         cvClient,
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

	cvResp, err := c.cvClient.GetCVByID(context.Background(), &cv.GetCVByIDRequest{CvId: notification.CvID})
	if err != nil {
		c.logger.Error("Error getting cv", zap.Error(err))
	}

	notification.CvID = cvResp.ExternalId
	notification.Message = utils.Translate(message)

	c.webSocketManager.BroadcastToUser(notification.UserID, notification)
}
