package container

import (
	"fmt"
	"gateway-service/internal/auth/services"
	"gateway-service/internal/cv/grpc/cv"
	services2 "gateway-service/internal/cv/services"
	generator "gateway-service/internal/generator/grpc"
	services5 "gateway-service/internal/generator/services"
	certificates "gateway-service/internal/information/certificates/grpc"
	contacts "gateway-service/internal/information/contacts/grpc"
	educations "gateway-service/internal/information/educations/grpc"
	experiences "gateway-service/internal/information/experiences/grpc"
	information "gateway-service/internal/information/information/grpc"
	languages "gateway-service/internal/information/languages/grpc"
	services3 "gateway-service/internal/information/services"
	skills "gateway-service/internal/information/skills/grpc"
	templates "gateway-service/internal/templates/grpc"
	services4 "gateway-service/internal/templates/services"
	"gateway-service/internal/users/grpc/auth"
)

type GrpcConnections struct {
	AuthClient            auth.AuthServiceClient
	CVClient              cv.CVServiceClient
	MainInfoClient        information.InformationServiceClient
	LanguagesClient       languages.LanguagesServiceClient
	SkillsClient          skills.SkillsServiceClient
	CertificatesClient    certificates.CertificatesServiceClient
	ContactsClient        contacts.ContactsServiceClient
	EducationsClient      educations.EducationServiceClient
	WorkExperiencesClient experiences.ExperiencesServiceClient
	TemplatesClient       templates.TemplateServiceClient
	GenerationsClient     generator.GeneratorServiceClient
}

func NewGrpcConnections() *GrpcConnections {
	fmt.Println("Initializing gRPC connections...")

	authConn := services.GetAuthConnection()
	cvConn := services2.GetCVConnection()
	informationConn := services3.GetInformationConnection()
	templatesConn := services4.GetTemplatesConnection()
	generatorConn := services5.GetGeneratorConnection()

	return &GrpcConnections{
		AuthClient:            auth.NewAuthServiceClient(authConn),
		CVClient:              cv.NewCVServiceClient(cvConn),
		MainInfoClient:        information.NewInformationServiceClient(informationConn),
		LanguagesClient:       languages.NewLanguagesServiceClient(informationConn),
		SkillsClient:          skills.NewSkillsServiceClient(informationConn),
		CertificatesClient:    certificates.NewCertificatesServiceClient(informationConn),
		ContactsClient:        contacts.NewContactsServiceClient(informationConn),
		EducationsClient:      educations.NewEducationServiceClient(informationConn),
		WorkExperiencesClient: experiences.NewExperiencesServiceClient(informationConn),
		TemplatesClient:       templates.NewTemplateServiceClient(templatesConn),
		GenerationsClient:     generator.NewGeneratorServiceClient(generatorConn),
	}
}
