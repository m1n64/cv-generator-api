package handlers

import (
	"context"
	"fmt"
	"gateway-service/internal/cv/entities"
	"gateway-service/internal/cv/grpc/cv"
	"gateway-service/internal/cv/services"
	certificates "gateway-service/internal/information/certificates/grpc"
	contacts "gateway-service/internal/information/contacts/grpc"
	educations "gateway-service/internal/information/educations/grpc"
	experiences "gateway-service/internal/information/experiences/grpc"
	information "gateway-service/internal/information/information/grpc"
	languages "gateway-service/internal/information/languages/grpc"
	skills "gateway-service/internal/information/skills/grpc"
	templates "gateway-service/internal/templates/grpc"
	"gateway-service/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"github.com/vmihailenco/msgpack/v5"
	"go.uber.org/zap"
	"net/http"
	"sync"
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

func NewGeneratorHandler(
	rabbitMq *utils.RabbitMQConnection,
	logger *zap.Logger,
	cvClient cv.CVServiceClient,
	informationClient information.InformationServiceClient,
	contactsClient contacts.ContactsServiceClient,
	skillsClient skills.SkillsServiceClient,
	languagesClient languages.LanguagesServiceClient,
	experiencesClient experiences.ExperiencesServiceClient,
	educationsClient educations.EducationServiceClient,
	certificatesClient certificates.CertificatesServiceClient,
	templatesClient templates.TemplateServiceClient,
) *GeneratorHandler {
	return &GeneratorHandler{
		rabbitMq:           rabbitMq,
		logger:             logger,
		cvClient:           cvClient,
		informationClient:  informationClient,
		contactsClient:     contactsClient,
		skillsClient:       skillsClient,
		languagesClient:    languagesClient,
		experiencesClient:  experiencesClient,
		educationsClient:   educationsClient,
		certificatesClient: certificatesClient,
		templatesClient:    templatesClient,
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

	colorName := c.DefaultQuery("color", "blue")
	templateId := c.DefaultQuery("template", "")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var wg sync.WaitGroup
	errorsChan := make(chan error, 10)

	var (
		cvMain         *cv.CVResponse
		cvInformation  *information.InformationResponse
		cvContacts     *contacts.AllContactsResponse
		cvSkills       *skills.AllSkillsResponse
		cvLanguages    *languages.AllLanguagesResponse
		cvExperiences  *experiences.AllExperiencesResponse
		cvEducations   *educations.AllEducationsResponse
		cvCertificates *certificates.AllCertificatesResponse
		template       string
		color          *templates.Color
	)

	runGRPC := func(fn func() error) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					err := fmt.Errorf("panic recovered: %v", r)
					h.logger.Error("Panic in GRPC call", zap.Error(err))
					errorsChan <- err
				}
			}()

			if err := fn(); err != nil {
				errorsChan <- err
			}
		}()
	}

	runGRPC(func() error {
		var err error
		cvMain, err = h.cvClient.GetCVByID(ctx, &cv.GetCVByIDRequest{CvId: cvId})
		return err
	})

	runGRPC(func() error {
		var err error
		cvInformation, err = h.informationClient.GetInformationByCvID(ctx, &information.GetInformationByCvIDRequest{CvId: cvId})
		return err
	})

	runGRPC(func() error {
		var err error
		cvContacts, err = h.contactsClient.GetContacts(ctx, &contacts.GetContactsRequest{CvId: cvId})
		return err
	})

	runGRPC(func() error {
		var err error
		cvSkills, err = h.skillsClient.GetSkills(ctx, &skills.GetSkillsRequest{CvId: cvId})
		return err
	})

	runGRPC(func() error {
		var err error
		cvLanguages, err = h.languagesClient.GetLanguages(ctx, &languages.GetLanguagesRequest{CvId: cvId})
		return err
	})

	runGRPC(func() error {
		var err error
		cvExperiences, err = h.experiencesClient.GetExperiences(ctx, &experiences.GetExperiencesRequest{CvId: cvId})
		return err
	})

	runGRPC(func() error {
		var err error
		cvEducations, err = h.educationsClient.GetEducations(ctx, &educations.GetEducationsRequest{CvId: cvId})
		return err
	})

	runGRPC(func() error {
		var err error
		cvCertificates, err = h.certificatesClient.GetCertificates(ctx, &certificates.GetCertificatesRequest{CvId: cvId})
		return err
	})

	runGRPC(func() error {
		var err error
		if templateId != "" {
			var templateResp *templates.TemplateResponse
			templateResp, err = h.templatesClient.GetTemplateById(ctx, &templates.TemplateByIdRequest{Id: templateId})

			template = templateResp.Template
		} else {
			var templateResp *templates.Template
			templateResp, err = h.templatesClient.GetDefaultTemplate(ctx, &templates.Empty{})

			template = templateResp.Template
		}

		return err
	})

	runGRPC(func() error {
		var err error
		color, err = h.templatesClient.GetColorSchemeByName(ctx, &templates.ColorSchemeByNameRequest{Name: colorName})
		return err
	})

	wg.Wait()

	close(errorsChan)
	for err := range errorsChan {
		if err != nil {
			h.logger.Error("GRPC error", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch data from services"})
			return
		}
	}

	var contactsInfo []*entities.CVContact
	for _, contact := range cvContacts.Contacts {
		contactsInfo = append(contactsInfo, &entities.CVContact{
			Title: contact.Title,
			Link:  contact.Link,
		})
	}

	var skillsInfo []*entities.CVSkill
	for _, skill := range cvSkills.Skills {
		skillsInfo = append(skillsInfo, &entities.CVSkill{
			Name: skill.Name,
		})
	}

	var languagesInfo []*entities.CVLanguage
	for _, language := range cvLanguages.Languages {
		languagesInfo = append(languagesInfo, &entities.CVLanguage{
			Name:  language.Name,
			Level: language.Level,
		})
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

	accentColor := &color.AccentColor
	if err != nil {
		accentColor = nil
	}

	data, err := msgpack.Marshal(&entities.CvInfo{
		UserID:   uuid.MustParse(userId.(string)),
		CvID:     uuid.MustParse(cvId),
		Template: template,
		Color:    accentColor,
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
