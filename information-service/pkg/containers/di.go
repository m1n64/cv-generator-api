package containers

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
	repositories4 "information-service/internal/certificates/repositories"
	services4 "information-service/internal/certificates/services"
	"information-service/internal/general/handlers"
	"information-service/internal/information/repositories"
	"information-service/internal/information/services"
	repositories2 "information-service/internal/languages/repositories"
	services2 "information-service/internal/languages/services"
	repositories3 "information-service/internal/skills/repositories"
	services3 "information-service/internal/skills/services"
	"information-service/pkg/utils"
	"log"
)

type Dependencies struct {
	DB                   *gorm.DB
	RedisClient          *redis.Client
	RedisAdapter         *utils.RedisAdapter
	Logger               *zap.Logger
	InformationRepo      repositories.InformationRepository
	LanguageRepo         repositories2.LanguageRepository
	CertificateRepo      repositories4.CertificateRepository
	SkillRepo            repositories3.SkillRepository
	CVInformationService *services.CVInformationService
	LanguageService      *services2.LanguageService
	SkillService         *services3.SkillService
	CertificateService   *services4.CertificateService
	RabbitMQ             *utils.RabbitMQConnection
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

	// Repositories
	informationRepo := repositories.NewInformationRepository(db)
	langRepo := repositories2.NewLanguageRepository(db)
	skillRepo := repositories3.NewSkillRepository(db)
	certificatesRepo := repositories4.NewCertificateRepository(db)

	// Services
	langAnalyticsService := services2.NewLanguageAnalyticsService(rabbitMQ, logger)
	skillAnalyticsService := services3.NewSkillsAnalyticsService(rabbitMQ, logger)

	cvService := services.NewCVInformationService(informationRepo, db)
	languageService := services2.NewLanguageService(langRepo, langAnalyticsService, db)
	skillService := services3.NewSkillService(skillRepo, skillAnalyticsService, db)
	certificatesService := services4.NewCertificateService(certificatesRepo, db)

	// Dependencies
	return &Dependencies{
		DB:                   db,
		RedisClient:          redisClient,
		RedisAdapter:         redisAdapter,
		Logger:               logger,
		InformationRepo:      informationRepo,
		LanguageRepo:         langRepo,
		SkillRepo:            skillRepo,
		CertificateRepo:      certificatesRepo,
		CVInformationService: cvService,
		LanguageService:      languageService,
		SkillService:         skillService,
		RabbitMQ:             rabbitMQ,
		CertificateService:   certificatesService,
	}, nil
}

func InitializeQueuesConsumer(dependencies *Dependencies) {
	cvDelHandler := handlers.NewDeleteCVInfoHandler(
		dependencies.Logger,
		dependencies.CVInformationService,
		dependencies.LanguageService,
		dependencies.SkillService,
	)

	err := utils.ListenToQueue(utils.DeleteCVQueueName, cvDelHandler.HandleDeleteCVMessage)
	if err != nil {
		log.Fatalf("Error starting listener for delete_cv_queue: %v", err)
	}
}
