package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
)

var Ctx = context.Background()

var Rdb *redis.Client

// CreateRedisConn создает подключение к редису
func CreateRedisConn() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: "",
		DB:       0,
	})
}

// GetRedisConn возвращает контекст и подключение к редису
func GetRedisConn() (context.Context, *redis.Client) {
	return Ctx, Rdb
}
