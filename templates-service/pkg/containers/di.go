package containers

import (
	"cv-templates-service/internal/templates/repositories"
	"cv-templates-service/internal/templates/seeders"
	"cv-templates-service/internal/templates/services"
	"cv-templates-service/pkg/utils"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

type Dependencies struct {
	DB                     *gorm.DB
	RedisClient            *redis.Client
	RedisAdapter           *utils.RedisAdapter
	Logger                 *zap.Logger
	MinioClient            *utils.MinioClient
	RabbitMQ               *utils.RabbitMQConnection
	TemplateRepo           repositories.TemplateRepository
	ColorRepo              repositories.ColorRepository
	DefaultTemplateService *services.DefaultTemplateService
	ColorService           *services.ColorService
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

	minioClient := utils.NewMinioClient(os.Getenv("MINIO_ENDPOINT"), os.Getenv("MINIO_ROOT_USER"), os.Getenv("MINIO_ROOT_PASSWORD"), "html-templates", os.Getenv("MINIO_SECURE") == "true")

	// Repositories
	templateRepo := repositories.NewTemplateGormRepository(db)
	colorRepo := repositories.NewColorMapRepository()

	// Services
	defaultTemplateService := services.NewDefaultTemplateService(templateRepo, db)
	colorService := services.NewColorService(colorRepo)

	if os.Getenv("FIRST_START") == "true" {
		utils.SeedDB([]utils.Seeder{
			seeders.NewTemplateSeeder(defaultTemplateService, minioClient),
		})

		utils.UpdateEnvValue("FIRST_START", "false", ".env")
		os.Setenv("FIRST_START", "false")
	}

	// Dependencies
	return &Dependencies{
		DB:                     db,
		RedisClient:            redisClient,
		RedisAdapter:           redisAdapter,
		Logger:                 logger,
		MinioClient:            minioClient,
		RabbitMQ:               rabbitMQ,
		TemplateRepo:           templateRepo,
		ColorRepo:              colorRepo,
		DefaultTemplateService: defaultTemplateService,
		ColorService:           colorService,
	}, nil
}

func InitializeQueuesConsumer(dependencies *Dependencies) {
}
