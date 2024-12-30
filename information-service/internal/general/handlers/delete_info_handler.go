package handlers

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"information-service/internal/information/services"
	services2 "information-service/internal/languages/services"
	services3 "information-service/internal/skills/services"
)

type DeleteCVMessage struct {
	CvID string `json:"cv_id"`
}

type DeleteCVInfoHandler struct {
	logger          *zap.Logger
	infoService     *services.CVInformationService
	languageService *services2.LanguageService
	skillService    *services3.SkillService
}

func NewDeleteCVInfoHandler(logger *zap.Logger, infoService *services.CVInformationService, languageService *services2.LanguageService, skillService *services3.SkillService) *DeleteCVInfoHandler {
	return &DeleteCVInfoHandler{
		logger:          logger,
		infoService:     infoService,
		languageService: languageService,
		skillService:    skillService,
	}
}

func (h *DeleteCVInfoHandler) HandleDeleteCVMessage(msg amqp.Delivery) {
	var deleteMsg DeleteCVMessage
	err := json.Unmarshal(msg.Body, &deleteMsg)
	if err != nil {
		h.logger.Error(fmt.Sprintf("Error unmarshalling delete CV message: %v", err))
		return
	}

	if uuid.Validate(deleteMsg.CvID) != nil {
		h.logger.Error("cv_id is required and must be a valid uuid")
		return
	}

	cvID := uuid.MustParse(deleteMsg.CvID)

	h.logger.Info(fmt.Sprintf("Received delete CV message: %s", deleteMsg.CvID))

	err = h.infoService.DeleteInformation(cvID)
	if err != nil {
		h.logger.Error(fmt.Sprintf("Error deleting CV: %v", err))
	}

	err = h.languageService.DeleteLanguagesByCvID(cvID)
	if err != nil {
		h.logger.Error(fmt.Sprintf("Error deleting CV: %v", err))
	}

	err = h.skillService.DeleteSkillsByCvID(cvID)
	if err != nil {
		h.logger.Error(fmt.Sprintf("Error deleting CV: %v", err))
	}

	h.logger.Info(fmt.Sprintf("Deleted CV: %s", deleteMsg.CvID))
}
