package service

import (
	"cv-service/internal/cv/models"
	mocks "cv-service/internal/cv/tests/mocks"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/google/uuid"
)

// @TODO: FIX FCKN DATABASE MOCK!
func TestCVService_CreateCV(t *testing.T) {
	db, mockDB, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		t.Fatalf("failed to initialize GORM: %v", err)
	}

	// Данные для теста
	userID := uuid.New()
	cvID := uuid.New()       // UUID для поля ID
	externalID := uuid.New() // UUID для ExternalID
	cvName := "Test CV"

	// Настраиваем модель с заполнением ID
	cvModel := &models.CV{
		ID:         cvID,
		ExternalID: externalID,
		UserID:     userID,
		Title:      cvName,
	}

	mockDB.ExpectBegin() // Начало транзакции

	mockDB.ExpectExec("^INSERT INTO .*").
		WithArgs(
			cvModel.ID,         // ID
			cvModel.ExternalID, // External ID
			cvModel.UserID,     // User ID
			cvModel.Title,      // Title
			sqlmock.AnyArg(),   // created_at
			sqlmock.AnyArg(),   // updated_at
			sqlmock.AnyArg(),
		).
		WillReturnResult(sqlmock.NewResult(1, 1)) // Указываем успешный результат вставки

	mockDB.ExpectCommit() // Завершение транзакции

	// Моки
	repoMock := new(mocks.CVRepositoryMock)
	redisMock := new(mocks.RedisClientMock)

	svc := NewCVService(repoMock, redisMock, gormDB)

	repoMock.On("CreateCV", mock.AnythingOfType("*models.CV")).Return(nil)
	redisMock.On("Set", mock.Anything, mock.AnythingOfType("string"), mock.Anything, time.Hour*24).Return(nil)

	// Тестируем
	result, err := svc.CreateCV(userID, cvName)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, cvName, result.Title)
	assert.Equal(t, userID, result.UserID)

	if err := mockDB.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
