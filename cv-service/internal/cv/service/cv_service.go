package service

import (
	"context"
	"cv-service/internal/cv/models"
	"cv-service/internal/cv/repositories"
	"cv-service/pkg/utils"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"time"
)

type CVService struct {
	cvRepo      repositories.CVRepository
	redisClient utils.RedisClient
	db          *gorm.DB
}

func NewCVService(cvRepo repositories.CVRepository, redisClient utils.RedisClient, db *gorm.DB) *CVService {
	return &CVService{
		cvRepo:      cvRepo,
		redisClient: redisClient,
		db:          db,
	}
}

func (s *CVService) CreateCV(userID uuid.UUID, name string) (*models.CV, error) {
	cvModel := &models.CV{
		UserID: userID,
		Title:  name,
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := s.cvRepo.CreateCV(cvModel); err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		ctx := context.Background()
		cacheKey := fmt.Sprintf("cv:original_id:%s:%s", userID.String(), cvModel.ExternalID.String())

		err := s.redisClient.Set(ctx, cacheKey, cvModel.ID.String(), time.Hour*24)
		if err != nil {
			utils.GetLogger().Error(fmt.Sprintf("Error saving CV ID to Redis: %v", err))
		}

		analyticMessage := utils.CVAnalyticQueueMessage{
			UserID:   userID,
			CvID:     cvModel.ExternalID,
			Action:   "cv_create",
			DateTime: cvModel.CreatedAt,
			Detail:   "CV created successfully",
		}

		body, err := json.Marshal(analyticMessage)
		if err != nil {
			utils.GetLogger().Error(fmt.Sprintf("Error marshaling CV analytic message: %v", err))
		}

		rabbit := utils.GetRabbitMQInstance()
		err = rabbit.Channel.Publish(
			"",
			utils.AnalyticQueueName,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			},
		)

		if err != nil {
			utils.GetLogger().Error(fmt.Sprintf("Error publishing CV analytic message: %v", err))
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return cvModel, nil
}

func (s *CVService) GetAllCVsByUserID(userID uuid.UUID) ([]models.CV, error) {
	cvs, err := s.cvRepo.GetAllCVsByUserID(userID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return cvs, nil
}

func (s *CVService) GetCVByID(cvID uuid.UUID) (*models.CV, error) {
	cvModel, err := s.cvRepo.GetCVByID(cvID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return cvModel, nil
}

func (s *CVService) DeleteCVByID(cvID uuid.UUID) error {
	cv, err := s.cvRepo.GetCVByID(cvID)
	if err != nil {
		return err
	}

	err = s.db.Transaction(func(tx *gorm.DB) error {
		ctx := context.Background()
		cacheKey := fmt.Sprintf("cv:original_id:%s:%s", cv.UserID.String(), cv.ExternalID.String())

		err = s.redisClient.Del(ctx, cacheKey)
		if err != nil {
			utils.GetLogger().Error(fmt.Sprintf("Error remove CV ID from Redis: %v", err))
		}

		analyticMessage := utils.CVAnalyticQueueMessage{
			UserID:   cv.UserID,
			CvID:     cv.ExternalID,
			Action:   "cv_delete",
			DateTime: cv.CreatedAt,
			Detail:   "CV created successfully",
		}

		analyticBody, err := json.Marshal(analyticMessage)
		if err != nil {
			utils.GetLogger().Error(fmt.Sprintf("Error marshaling CV analytic message: %v", err))
		}

		rabbit := utils.GetRabbitMQInstance()
		err = rabbit.Channel.Publish(
			"",
			utils.AnalyticQueueName,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        analyticBody,
			},
		)

		if err != nil {
			utils.GetLogger().Error(fmt.Sprintf("Error publishing CV analytic message: %v", err))
		}

		deleteQueueBody := map[string]interface{}{
			"cv_id": cv.ID,
		}

		deleteBody, err := json.Marshal(deleteQueueBody)

		fmt.Println(deleteBody, deleteQueueBody)

		err = rabbit.Channel.Publish(
			"",
			utils.DeleteCVQueueName,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        deleteBody,
			},
		)

		if err != nil {
			utils.GetLogger().Error(fmt.Sprintf("Error publishing CV analytic message: %v", err))
		}

		err = s.cvRepo.DeleteCVByID(cvID)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *CVService) UpdateCV(cvID uuid.UUID, name string) (*models.CV, error) {
	updatedCV := &models.CV{
		Title: name,
	}

	var updatedCVFromDB *models.CV
	err := s.db.Transaction(func(tx *gorm.DB) error {
		err := s.cvRepo.UpdateCVByID(cvID, updatedCV)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		updatedCVFromDB, err = s.cvRepo.GetCVByID(cvID)
		if err != nil {
			return status.Error(codes.Internal, "failed to fetch updated CV")
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return updatedCVFromDB, nil
}

func (s *CVService) GetOriginalID(userID uuid.UUID, cvID uuid.UUID) (*uuid.UUID, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("cv:original_id:%s:%s", userID, cvID)

	cachedValue, err := s.redisClient.Get(ctx, cacheKey)
	if err == nil {
		originalID, parseErr := uuid.Parse(cachedValue)
		if parseErr != nil {
			utils.GetLogger().Error(fmt.Sprintf("Error parsing CV ID from Redis: %v", parseErr))
		} else {
			return &originalID, nil
		}
	} else if !errors.Is(err, redis.Nil) {
		utils.GetLogger().Error(fmt.Sprintf("Error reading CV ID from Redis: %v", err))
	}

	originalId, err := s.cvRepo.GetOriginalIDByExternalID(cvID, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = s.redisClient.Set(ctx, cacheKey, originalId.String(), time.Hour*24)
	if err != nil {
		utils.GetLogger().Error(fmt.Sprintf("Error saving CV ID to Redis: %v", err))
	}

	return &originalId, nil
}
