package utils

import (
	"github.com/joho/godotenv"
)

// LoadEnv - загружает переменные окружения из .env
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		GetLogger().Sugar().Error("Error loading .env file")
	}
}
