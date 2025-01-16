package handlers

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"time"
	health "user-service/internal/health/grpc"
)

type HealthServiceServer struct {
	health.UnimplementedHealthServiceServer
	db    *gorm.DB
	redis *redis.Client
}

func NewHealthServiceServer(db *gorm.DB, redis *redis.Client) *HealthServiceServer {
	return &HealthServiceServer{
		db:    db,
		redis: redis,
	}
}

func (s *HealthServiceServer) Check(ctx context.Context, req *health.CheckRequest) (*health.CheckResponse, error) {
	dbStatus := true
	redisStatus := true

	if db, err := s.db.DB(); err != nil || db.Ping() != nil {
		dbStatus = false
	}

	status := s.redis.Ping(ctx)
	if status.Err() != nil {
		redisStatus = false
	}

	return &health.CheckResponse{
		ServiceName: req.Service,
		Status:      true,
		StatusDb:    dbStatus,
		StatusRedis: redisStatus,
		Timestamp:   time.Now().Format(time.RFC3339),
	}, nil
}
