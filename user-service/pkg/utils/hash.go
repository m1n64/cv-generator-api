package utils

import (
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

var jwtSecret = []byte(os.Getenv("APP_SECRET_KEY"))

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Генерация JWT токена
func GenerateToken(userID string) (string, error) {
	// Устанавливаем срок действия токена
	expirationTime := time.Now().Add(24 * time.Hour)

	// Создаём payload (claims)
	claims := &jwt.MapClaims{
		"user_id": userID,
		"exp":     expirationTime.Unix(),
	}

	// Генерируем токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Проверка JWT токена
func ValidateToken(tokenStr string) (string, error) {
	// Разбираем токен
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Проверяем метод подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})

	if err != nil {
		return "", err
	}

	// Проверяем claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userID, ok := claims["user_id"].(string); ok {
			return userID, nil
		}
	}

	return "", jwt.ErrTokenNotValidYet
}
