package repositories

import (
	"gorm.io/gorm"
	"user-service/internal/users/models"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetDB() *gorm.DB {
	return r.db
}

func (r *userRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByID(id string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByToken(token string) (*models.User, error) {
	var user models.User
	if err := r.db.Joins("JOIN tokens ON tokens.user_id = users.id").
		Where("tokens.token = ?", token).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
