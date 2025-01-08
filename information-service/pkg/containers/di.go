package containers

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
	repositories4 "information-service/internal/certificates/repositories"
	services4 "information-service/internal/certificates/services"
	repositories5 "information-service/internal/contacts/repositories"
	services6 "information-service/internal/contacts/services"
	repositories6 "information-service/internal/educations/repositories"
	services7 "information-service/internal/educations/services"
	repositories7 "information-service/internal/experiences/repositories"
	services8 "information-service/internal/experiences/services"
	"information-service/internal/general/handlers"
	services5 "information-service/internal/general/services"
	"information-service/internal/information/repositories"
	"information-service/internal/information/services"
	repositories2 "information-service/internal/languages/repositories"
	services2 "information-service/internal/languages/services"
	repositories3 "information-service/internal/skills/repositories"
	services3 "information-service/internal/skills/services"
	"information-service/pkg/utils"
	"log"
	"os"
)

type Dependencies struct {
	DB                    *gorm.DB
	RedisClient           *redis.Client
	RedisAdapter          *utils.RedisAdapter
	Logger                *zap.Logger
	MinioClient           *utils.MinioClient
	InformationRepo       repositories.InformationRepository
	LanguageRepo          repositories2.LanguageRepository
	CertificateRepo       repositories4.CertificateRepository
	SkillRepo             repositories3.SkillRepository
	ContactRepo           repositories5.ContactRepository
	EducationRepo         repositories6.EducationRepository
	WorkExperienceRepo    repositories7.WorkExperienceRepository
	CVInformationService  *services.CVInformationService
	FileService           *services5.FileService
	LanguageService       *services2.LanguageService
	SkillService          *services3.SkillService
	CertificateService    *services4.CertificateService
	ContactService        *services6.ContactService
	EducationService      *services7.EducationService
	WorkExperienceService *services8.WorkExperienceService
	RabbitMQ              *utils.RabbitMQConnection
}

func InitializeDependencies() (*Dependencies, error) {
	utils.InitLogs()
	utils.LoadEnv()
	utils.CreateRedisConn()
	utils.InitDBConnection()
	utils.StartMigrations()
	utils.InitValidator()
	rabbitMQ := utils.ConnectRabbitMQ()
	utils.InitializeQueues()

	logger := utils.GetLogger()

	db := utils.GetDBConnection()

	_, redisClient := utils.GetRedisConn()

	redisAdapter := utils.NewRedisAdapter(redisClient)

	minioClient := utils.NewMinioClient(os.Getenv("MINIO_ENDPOINT"), os.Getenv("MINIO_ROOT_USER"), os.Getenv("MINIO_ROOT_PASSWORD"), "cv-information", os.Getenv("MINIO_SECURE") == "true")

	// Repositories
	informationRepo := repositories.NewInformationGormRepository(db)
	langRepo := repositories2.NewLanguageGormRepository(db)
	skillRepo := repositories3.NewSkillGormRepository(db)
	certificatesRepo := repositories4.NewCertificateGormRepository(db)
	contactRepo := repositories5.NewContactGormRepository(db)
	educationRepo := repositories6.NewEducationGormRepository(db)
	workExperienceRepo := repositories7.NewWorkExperienceGormRepository(db)

	// Services
	langAnalyticsService := services2.NewLanguageAnalyticsService(rabbitMQ, logger)
	skillAnalyticsService := services3.NewSkillsAnalyticsService(rabbitMQ, logger)
	fileService := services5.NewFileService(minioClient, logger)

	cvService := services.NewCVInformationService(informationRepo, db)
	languageService := services2.NewLanguageService(langRepo, langAnalyticsService, db)
	skillService := services3.NewSkillService(skillRepo, skillAnalyticsService, db)
	certificatesService := services4.NewCertificateService(certificatesRepo, db)
	contactService := services6.NewContactService(contactRepo, db)
	educationService := services7.NewEducationService(educationRepo, db)
	workExperienceService := services8.NewWorkExperienceService(workExperienceRepo, db)

	// Dependencies
	return &Dependencies{
		DB:                    db,
		RedisClient:           redisClient,
		RedisAdapter:          redisAdapter,
		Logger:                logger,
		MinioClient:           minioClient,
		InformationRepo:       informationRepo,
		LanguageRepo:          langRepo,
		SkillRepo:             skillRepo,
		CertificateRepo:       certificatesRepo,
		ContactRepo:           contactRepo,
		EducationRepo:         educationRepo,
		WorkExperienceRepo:    workExperienceRepo,
		CVInformationService:  cvService,
		FileService:           fileService,
		LanguageService:       languageService,
		SkillService:          skillService,
		RabbitMQ:              rabbitMQ,
		CertificateService:    certificatesService,
		ContactService:        contactService,
		EducationService:      educationService,
		WorkExperienceService: workExperienceService,
	}, nil
}

func InitializeQueuesConsumer(dependencies *Dependencies) {
	cvDelHandler := handlers.NewDeleteCVInfoHandler(
		dependencies.Logger,
		dependencies.CVInformationService,
		dependencies.LanguageService,
		dependencies.SkillService,
		dependencies.CertificateService,
		dependencies.ContactService,
		dependencies.EducationService,
		dependencies.WorkExperienceService,
	)

	err := utils.ListenToQueue(utils.DeleteCVQueueName, cvDelHandler.HandleDeleteCVMessage)
	if err != nil {
		log.Fatalf("Error starting listener for delete_cv_queue: %v", err)
	}
}
