package containers

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"information-service/internal/generator/repositories"
	"information-service/internal/generator/services"
	"information-service/pkg/utils"
	"os"
)

type Dependencies struct {
	DB               *gorm.DB
	RedisClient      *redis.Client
	RedisAdapter     *utils.RedisAdapter
	Logger           *zap.Logger
	MinioClient      *utils.MinioClient
	RabbitMQ         *utils.RabbitMQConnection
	GeneratedPDFRepo repositories.GeneratedPDFRepository

	GeneratedPDFService *services.GeneratedPDFService
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

	minioClient := utils.NewMinioClient(os.Getenv("MINIO_ENDPOINT"), os.Getenv("MINIO_ROOT_USER"), os.Getenv("MINIO_ROOT_PASSWORD"), "cv-pdf", os.Getenv("MINIO_SECURE") == "true")

	// Repositories
	generatedPDFRepo := repositories.NewGeneratedPDFGormRepository(db)

	// Services
	generatorPDFService := services.NewGeneratedPDFService(generatedPDFRepo, db)

	// Dependencies
	return &Dependencies{
		DB:                  db,
		RedisClient:         redisClient,
		RedisAdapter:        redisAdapter,
		Logger:              logger,
		MinioClient:         minioClient,
		RabbitMQ:            rabbitMQ,
		GeneratedPDFRepo:    generatedPDFRepo,
		GeneratedPDFService: generatorPDFService,
	}, nil
}

func InitializeQueuesConsumer(dependencies *Dependencies) {
}
