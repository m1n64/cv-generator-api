package handlers

import (
	"context"
	"gateway-service/internal/cv/entities"
	"gateway-service/internal/cv/grpc/cv"
	"gateway-service/internal/cv/services"
	certificates "gateway-service/internal/information/certificates/grpc"
	contacts "gateway-service/internal/information/contacts/grpc"
	educations "gateway-service/internal/information/educations/grpc"
	experiences "gateway-service/internal/information/experiences/grpc"
	information "gateway-service/internal/information/information/grpc"
	languages "gateway-service/internal/information/languages/grpc"
	services2 "gateway-service/internal/information/services"
	skills "gateway-service/internal/information/skills/grpc"
	templates "gateway-service/internal/templates/grpc"
	services3 "gateway-service/internal/templates/services"
	"gateway-service/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type GeneratorHandler struct {
	rabbitMq           *utils.RabbitMQConnection
	logger             *zap.Logger
	cvClient           cv.CVServiceClient
	informationClient  information.InformationServiceClient
	contactsClient     contacts.ContactsServiceClient
	skillsClient       skills.SkillsServiceClient
	languagesClient    languages.LanguagesServiceClient
	experiencesClient  experiences.ExperiencesServiceClient
	educationsClient   educations.EducationServiceClient
	certificatesClient certificates.CertificatesServiceClient
	templatesClient    templates.TemplateServiceClient
}

type GenerateCVResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewGeneratorHandler(rabbitMq *utils.RabbitMQConnection, logger *zap.Logger) *GeneratorHandler {
	cvConnection := services.GetCVConnection()
	informationConnection := services2.GetInformationConnection()
	templateConnection := services3.GetTemplatesConnection()

	return &GeneratorHandler{
		rabbitMq:           rabbitMq,
		logger:             logger,
		cvClient:           cv.NewCVServiceClient(cvConnection),
		informationClient:  information.NewInformationServiceClient(informationConnection),
		contactsClient:     contacts.NewContactsServiceClient(informationConnection),
		skillsClient:       skills.NewSkillsServiceClient(informationConnection),
		languagesClient:    languages.NewLanguagesServiceClient(informationConnection),
		experiencesClient:  experiences.NewExperiencesServiceClient(informationConnection),
		educationsClient:   educations.NewEducationServiceClient(informationConnection),
		certificatesClient: certificates.NewCertificatesServiceClient(informationConnection),
		templatesClient:    templates.NewTemplateServiceClient(templateConnection),
	}
}

func (h *GeneratorHandler) GenerateCV(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user_id not found"})
		return
	}

	cvId, err := services.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	cvMain, err := h.cvClient.GetCVByID(ctx, &cv.GetCVByIDRequest{
		CvId: cvId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cv not found"})
		return
	}

	cvInformation, err := h.informationClient.GetInformationByCvID(ctx, &information.GetInformationByCvIDRequest{
		CvId: cvId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cvContacts, err := h.contactsClient.GetContacts(ctx, &contacts.GetContactsRequest{
		CvId: cvId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var contactsInfo []*entities.CVContact
	for _, contact := range cvContacts.Contacts {
		contactsInfo = append(contactsInfo, &entities.CVContact{
			Title: contact.Title,
			Link:  contact.Link,
		})
	}

	cvSkills, err := h.skillsClient.GetSkills(ctx, &skills.GetSkillsRequest{
		CvId: cvId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var skillsInfo []*entities.CVSkill
	for _, skill := range cvSkills.Skills {
		skillsInfo = append(skillsInfo, &entities.CVSkill{
			Name: skill.Name,
		})
	}

	cvLanguages, err := h.languagesClient.GetLanguages(ctx, &languages.GetLanguagesRequest{
		CvId: cvId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var languagesInfo []*entities.CVLanguage
	for _, language := range cvLanguages.Languages {
		languagesInfo = append(languagesInfo, &entities.CVLanguage{
			Name:  language.Name,
			Level: language.Level,
		})
	}

	cvExperiences, err := h.experiencesClient.GetExperiences(ctx, &experiences.GetExperiencesRequest{
		CvId: cvId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var experiencesInfo []*entities.CVWorkExperience
	for _, experience := range cvExperiences.WorkExperiences {
		experiencesInfo = append(experiencesInfo, &entities.CVWorkExperience{
			Company:     experience.Company,
			Position:    experience.Position,
			StartDate:   experience.StartDate,
			EndDate:     experience.EndDate,
			Description: experience.Description,
			Location:    experience.Location,
		})
	}

	cvEducations, err := h.educationsClient.GetEducations(ctx, &educations.GetEducationsRequest{
		CvId: cvId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var educationsInfo []*entities.CVEducation
	for _, education := range cvEducations.Educations {
		educationsInfo = append(educationsInfo, &entities.CVEducation{
			Institution: education.Institution,
			Degree:      education.Degree,
			StartDate:   education.StartDate,
			EndDate:     education.EndDate,
			Description: education.Description,
			Location:    education.Location,
			Faculty:     education.Faculty,
		})
	}

	cvCertificates, err := h.certificatesClient.GetCertificates(ctx, &certificates.GetCertificatesRequest{
		CvId: cvId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var certificatesInfo []*entities.CVCertificate
	for _, certificate := range cvCertificates.Certificates {
		certificatesInfo = append(certificatesInfo, &entities.CVCertificate{
			Title:       certificate.Title,
			Vendor:      certificate.Vendor,
			StartDate:   certificate.StartDate,
			EndDate:     certificate.EndDate,
			Description: certificate.Description,
		})
	}

	template, err := h.templatesClient.GetDefaultTemplate(ctx, &templates.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data, err := json.Marshal(&entities.CvInfo{
		UserID:   uuid.MustParse(userId.(string)),
		CvID:     uuid.MustParse(cvId),
		Template: template.Template,
		CV: entities.CV{
			Title: cvMain.Name,
		},
		Information: &entities.CVInformation{
			FullName:  cvInformation.FullName,
			Photo:     &cvInformation.PhotoFile,
			Position:  cvInformation.Position,
			Location:  cvInformation.Location,
			Biography: cvInformation.Biography,
		},
		Contacts:        contactsInfo,
		Skills:          skillsInfo,
		Languages:       languagesInfo,
		WorkExperiences: experiencesInfo,
		Educations:      educationsInfo,
		Certificates:    certificatesInfo,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.rabbitMq.Channel.Publish(
		"",
		utils.PDFGenerateQueue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.logger.Error("Failed to publish message to queue: ", zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, &GenerateCVResponse{
		Success: true,
		Message: utils.Translate("cv.generate.in.queue"),
	})
}
