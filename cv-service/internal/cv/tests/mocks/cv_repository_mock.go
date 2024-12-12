package mocks

import (
	"cv-service/internal/cv/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type CVRepositoryMock struct {
	mock.Mock
}

func (m *CVRepositoryMock) CreateCV(cv *models.CV) error {
	args := m.Called(cv)
	return args.Error(0)
}

func (m *CVRepositoryMock) GetAllCVsByUserID(userID uuid.UUID) ([]models.CV, error) {
	args := m.Called(userID)
	return args.Get(0).([]models.CV), args.Error(1)
}

func (m *CVRepositoryMock) GetCVByID(cvID uuid.UUID) (*models.CV, error) {
	args := m.Called(cvID)
	return args.Get(0).(*models.CV), args.Error(1)
}

func (m *CVRepositoryMock) DeleteCVByID(cvID uuid.UUID) error {
	args := m.Called(cvID)
	return args.Error(0)
}

func (m *CVRepositoryMock) UpdateCVByID(cvID uuid.UUID, updatedCV *models.CV) error {
	args := m.Called(cvID, updatedCV)
	return args.Error(0)
}

func (m *CVRepositoryMock) GetOriginalIDByExternalID(externalID uuid.UUID, userID uuid.UUID) (uuid.UUID, error) {
	args := m.Called(externalID, userID)
	return args.Get(0).(uuid.UUID), args.Error(1)
}
