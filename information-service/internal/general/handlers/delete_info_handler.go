package handlers

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"information-service/internal/information/services"
	services2 "information-service/internal/languages/services"
	services3 "information-service/internal/skills/services"
	"information-service/pkg/utils"
)

type DeleteCVMessage struct {
	CvID string `json:"cv_id"`
}

type DeleteCVInfoHandler struct {
	infoService     *services.CVInformationService
	languageService *services2.LanguageService
	skillService    *services3.SkillService
}

func NewDeleteCVInfoHandler(infoService *services.CVInformationService, languageService *services2.LanguageService, skillService *services3.SkillService) *DeleteCVInfoHandler {
	return &DeleteCVInfoHandler{
		infoService:     infoService,
		languageService: languageService,
		skillService:    skillService,
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

	cvID := uuid.MustParse(deleteMsg.CvID)

	logger.Info(fmt.Sprintf("Received delete CV message: %s", deleteMsg.CvID))

	err = h.infoService.DeleteInformation(cvID)
	if err != nil {
		logger.Error(fmt.Sprintf("Error deleting CV: %v", err))
	}

	err = h.languageService.DeleteLanguagesByCvID(cvID)
	if err != nil {
		logger.Error(fmt.Sprintf("Error deleting CV: %v", err))
	}

	/*err = h.skillService.DeleteSkillsByCvID(cvID)
	if err != nil {
		logger.Error(fmt.Sprintf("Error deleting CV: %v", err))
	}*/

	logger.Info(fmt.Sprintf("Deleted CV: %s", deleteMsg.CvID))
}
