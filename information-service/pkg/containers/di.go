package containers

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
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
	InformationRepo      repositories.InformationRepository
	LanguageRepo         repositories2.LanguageRepository
	SkillRepo            repositories3.SkillRepository
	CVInformationService *services.CVInformationService
	LanguageService      *services2.LanguageService
	SkillService         *services3.SkillService
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

	db := utils.GetDBConnection()

	_, redisClient := utils.GetRedisConn()

	redisAdapter := utils.NewRedisAdapter(redisClient)

	// Repositories
	informationRepo := repositories.NewInformationRepository(db)
	langRepo := repositories2.NewLanguageRepository(db)
	skillRepo := repositories3.NewSkillRepository(db)

	// Services
	cvService := services.NewCVInformationService(informationRepo, db)
	languageService := services2.NewLanguageService(langRepo, db)
	skillService := services3.NewSkillService(skillRepo, db)

	// Dependencies
	return &Dependencies{
		DB:                   db,
		RedisClient:          redisClient,
		RedisAdapter:         redisAdapter,
		InformationRepo:      informationRepo,
		LanguageRepo:         langRepo,
		SkillRepo:            skillRepo,
		CVInformationService: cvService,
		LanguageService:      languageService,
		SkillService:         skillService,
		RabbitMQ:             rabbitMQ,
	}, nil
}

func InitializeQueuesConsumer(dependencies *Dependencies) {
	cvDelHandler := handlers.NewDeleteCVInfoHandler(
		dependencies.CVInformationService,
		dependencies.LanguageService,
		dependencies.SkillService,
	)

	err := utils.ListenToQueue(utils.DeleteCVQueueName, cvDelHandler.HandleDeleteCVMessage)
	if err != nil {
		log.Fatalf("Error starting listener for delete_cv_queue: %v", err)
	}
}
