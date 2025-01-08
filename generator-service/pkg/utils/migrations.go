package utils

import "information-service/internal/generator/models"

func StartMigrations() {
	db := GetDBConnection()
	db.AutoMigrate(&models.GeneratedPDF{})
}
