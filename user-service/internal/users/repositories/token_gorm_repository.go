package repositories

import (
	"gorm.io/gorm"
	"time"
	"user-service/internal/users/models"
)

type tokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) TokenRepository {
	return &tokenRepository{db: db}
}

func (r *tokenRepository) CreateToken(token *models.Token) error {
	return r.db.Create(token).Error
}

func (r *tokenRepository) DeleteTokenByID(id string) error {
	return r.db.Delete(&models.Token{}, "id = ?", id).Error
}

func (r *tokenRepository) DeleteTokenByValue(token string) error {
	return r.db.Delete(&models.Token{}, "token = ?", token).Error
}

func (r *tokenRepository) FindTokenByID(id string) (*models.Token, error) {
	var token models.Token
	if err := r.db.Where("id = ?", id).First(&token).Error; err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *tokenRepository) FindTokenByValue(token string) (*models.Token, error) {
	var tokenModel models.Token
	if err := r.db.Preload("User").Where("token = ?", token).First(&tokenModel).Error; err != nil {
		return nil, err
	}

	return &tokenModel, nil
}

func (r *tokenRepository) DeleteExpiredTokens() error {
	return r.db.Delete(&models.Token{}, "expires_at < ?", time.Now()).Error
}
