package handlers

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	services4 "information-service/internal/certificates/services"
	services5 "information-service/internal/contacts/services"
	services7 "information-service/internal/educations/services"
	services6 "information-service/internal/experiences/services"
	"information-service/internal/information/services"
	services2 "information-service/internal/languages/services"
	services3 "information-service/internal/skills/services"
)

type DeleteCVMessage struct {
	CvID string `json:"cv_id"`
}

type deletionFunc func(uuid.UUID) error

type DeleteCVInfoHandler struct {
	logger                *zap.Logger
	infoService           *services.CVInformationService
	languageService       *services2.LanguageService
	skillService          *services3.SkillService
	certificatesService   *services4.CertificateService
	contactService        *services5.ContactService
	educationService      *services7.EducationService
	workExperienceService *services6.WorkExperienceService
}

func NewDeleteCVInfoHandler(
	logger *zap.Logger,
	infoService *services.CVInformationService,
	languageService *services2.LanguageService,
	skillService *services3.SkillService,
	certificateService *services4.CertificateService,
	contactService *services5.ContactService,
	educationService *services7.EducationService,
	workExperienceService *services6.WorkExperienceService,
) *DeleteCVInfoHandler {
	return &DeleteCVInfoHandler{
		logger:                logger,
		infoService:           infoService,
		languageService:       languageService,
		skillService:          skillService,
		certificatesService:   certificateService,
		contactService:        contactService,
		educationService:      educationService,
		workExperienceService: workExperienceService,
	}
}

func (h *DeleteCVInfoHandler) HandleDeleteCVMessage(msg amqp.Delivery) {
	var deleteMsg DeleteCVMessage
	if err := json.Unmarshal(msg.Body, &deleteMsg); err != nil {
		h.logger.Error(fmt.Sprintf("Error unmarshalling delete CV message: %v", err))
		return
	}

	if uuid.Validate(deleteMsg.CvID) != nil {
		h.logger.Error("cv_id is required and must be a valid uuid")
		return
	}

	cvID := uuid.MustParse(deleteMsg.CvID)
	h.logger.Info(fmt.Sprintf("Received delete CV message: %s", deleteMsg.CvID))

	deletionSteps := map[string]deletionFunc{
		"information":  h.infoService.DeleteInformation,
		"languages":    h.languageService.DeleteLanguagesByCvID,
		"skills":       h.skillService.DeleteSkillsByCvID,
		"certificates": h.certificatesService.DeleteCertificates,
		"contacts":     h.contactService.DeleteContactsByCvID,
		"educations":   h.educationService.DeleteEducationsByCvID,
		"experiences":  h.workExperienceService.DeleteWorkExperiencesByCvID,
	}

	for name, deleteFunc := range deletionSteps {
		err := deleteFunc(cvID)
		h.handleError(err, fmt.Sprintf("Error deleting %s for CV", name))
	}

	h.logger.Info(fmt.Sprintf("Deleted CV: %s", deleteMsg.CvID))
}

func (h *DeleteCVInfoHandler) handleError(err error, message string) {
	if err != nil {
		h.logger.Error(fmt.Sprintf("%s: %v", message, err))
	}
}
