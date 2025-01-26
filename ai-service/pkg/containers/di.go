package containers

import (
	"ai-service/internal/ai/repositories"
	"ai-service/internal/ai/services"
	"ai-service/pkg/utils"
	deepseek "github.com/cohesion-org/deepseek-go"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

type Dependencies struct {
	DB                  *gorm.DB
	RedisClient         *redis.Client
	RedisAdapter        *utils.RedisAdapter
	Logger              *zap.Logger
	RabbitMQ            *utils.RabbitMQConnection
	DeepSeek            *deepseek.Client
	AiServiceRepository repositories.AiServicesRepository
	AiService           *services.AiService
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
	aiServiceRepository := repositories.NewAiServicesMapRepository()

	// Services
	deepSeek := deepseek.NewClient(os.Getenv("DEEPSEEK_TOKEN"))
	aiService := services.NewAiService(aiServiceRepository, deepSeek)

	// Dependencies
	return &Dependencies{
		DB:                  db,
		RedisClient:         redisClient,
		RedisAdapter:        redisAdapter,
		Logger:              logger,
		RabbitMQ:            rabbitMQ,
		DeepSeek:            deepSeek,
		AiServiceRepository: aiServiceRepository,
		AiService:           aiService,
	}, nil
}

func InitializeQueuesConsumer(dependencies *Dependencies) {
}
