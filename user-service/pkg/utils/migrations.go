package utils

import (
	"gorm.io/gorm"
	"user-service/internal/users/models"
)

func StartMigrations() {
	db := GetDBConnection()

	removeUniqueIndex(db)

	db.AutoMigrate(&models.User{}, &models.Token{})
}

func removeUniqueIndex(db *gorm.DB) error {
	err := db.Exec("ALTER TABLE users DROP CONSTRAINT IF EXISTS uni_users_email;").Error
	if err != nil {
		return err
	}
	return db.Exec("DROP INDEX IF EXISTS uni_users_email;").Error
}
