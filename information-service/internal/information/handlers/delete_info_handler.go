package handlers

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"information-service/internal/information/services"
	"information-service/pkg/utils"
)

type DeleteCVMessage struct {
	CvID string `json:"cv_id"`
}

type DeleteCVInfoHandler struct {
	infoService *services.CVInformationService
}

func NewDeleteCVInfoHandler(infoService *services.CVInformationService) *DeleteCVInfoHandler {
	return &DeleteCVInfoHandler{
		infoService: infoService,
	}
}

func (h *DeleteCVInfoHandler) HandleDeleteCVMessage(msg amqp.Delivery) {
	logger := utils.GetLogger()

	var deleteMsg DeleteCVMessage
	err := json.Unmarshal(msg.Body, &deleteMsg)
	if err != nil {
		logger.Error(fmt.Sprintf("Error unmarshalling delete CV message: %v", err))
		return
	}

	if uuid.Validate(deleteMsg.CvID) != nil {
		logger.Error("cv_id is required and must be a valid uuid")
		return
	}

	logger.Info(fmt.Sprintf("Received delete CV message: %s", deleteMsg.CvID))

	err = h.infoService.DeleteInformation(uuid.MustParse(deleteMsg.CvID))
	if err != nil {
		logger.Error(fmt.Sprintf("Error deleting CV: %v", err))
		return
	}

	logger.Info(fmt.Sprintf("Deleted CV: %s", deleteMsg.CvID))
}
