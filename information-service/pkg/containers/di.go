package containers

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"information-service/internal/information/handlers"
	"information-service/internal/information/repositories"
	"information-service/internal/information/services"
	"information-service/pkg/utils"
	"log"
)

type Dependencies struct {
	DB                   *gorm.DB
	RedisClient          *redis.Client
	RedisAdapter         *utils.RedisAdapter
	InformationRepo      repositories.InformationRepository
	CVInformationService *services.CVInformationService
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

	// Services
	cvService := services.NewCVInformationService(informationRepo, db)

	// Dependencies
	return &Dependencies{
		DB:                   db,
		RedisClient:          redisClient,
		RedisAdapter:         redisAdapter,
		InformationRepo:      informationRepo,
		CVInformationService: cvService,
		RabbitMQ:             rabbitMQ,
	}, nil
}

func InitializeQueuesConsumer(dependencies *Dependencies) {
	cvDelHandler := handlers.NewDeleteCVInfoHandler(dependencies.CVInformationService)

	err := utils.ListenToQueue(utils.DeleteCVQueueName, cvDelHandler.HandleDeleteCVMessage)
	if err != nil {
		log.Fatalf("Error starting listener for delete_cv_queue: %v", err)
	}
}
