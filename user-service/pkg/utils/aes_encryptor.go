package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
)

type AESEncryptor struct {
}

func NewAESEncryptor() AESEncryptor {
	return AESEncryptor{}
}

func (s *AESEncryptor) GetKey() []byte {
	keyBase64 := os.Getenv("LARAVEL_APP_KEY")
	keyBase64 = keyBase64[7:]

	key, err := base64.StdEncoding.DecodeString(keyBase64)

	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
	}
	return key
}

func (s *AESEncryptor) GetIVAndCipher(encoded string) ([]byte, []byte) {
	encStr, _ := base64.StdEncoding.DecodeString(encoded)

	iv := encStr[:aes.BlockSize]
	encBaseStr := encStr[aes.BlockSize:]

	return iv, encBaseStr
}

// Encrypt шифрует текст с использованием AES-256-CBC и заданного ключа и IV.
func (s *AESEncryptor) Encrypt(plainText string, key []byte, iv []byte) (string, error) {
	// Создание нового шифра AES
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %w", err)
	}

	// Проверяем размер блока шифра и длину IV
	if len(iv) != aes.BlockSize {
		return "", fmt.Errorf("invalid IV length: expected %d bytes, got %d bytes", aes.BlockSize, len(iv))
	}

	// Применяем PKCS7 паддинг
	paddedText := s.pkcs7Padding([]byte(plainText), aes.BlockSize)

	// Шифрование данных с использованием режима CBC
	encrypted := make([]byte, len(paddedText))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(encrypted, paddedText)

	// Возвращаем результат в base64 формате (IV + шифротекст)
	finalResult := base64.StdEncoding.EncodeToString(append(iv, encrypted...))

	return finalResult, nil
}

func (s *AESEncryptor) Decrypt(encrypted []byte, key []byte, iv []byte) (string, error) {
	// Создание декодера AES
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %w", err)
	}

	// Проверяем размер блока шифра и длину IV
	if len(iv) != aes.BlockSize {
		return "", fmt.Errorf("invalid IV length: expected %d bytes, got %d bytes", aes.BlockSize, len(iv))
	}

	// Проверяем длину зашифрованных данных
	if len(encrypted)%aes.BlockSize != 0 {
		return "", fmt.Errorf("invalid encrypted data length: must be a multiple of %d", aes.BlockSize)
	}

	// Расшифрование данных с использованием режима CBC
	mode := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(encrypted))
	mode.CryptBlocks(decrypted, encrypted)

	// Удаление возможных паддингов (PKCS7)
	decrypted = s.pkcs7UnPadding(decrypted)

	return string(decrypted), nil
}

// GenerateIV генерирует случайный IV длиной 16 байт (размер блока AES)
func (s *AESEncryptor) GenerateIV() ([]byte, error) {
	iv := make([]byte, aes.BlockSize)
	_, err := rand.Read(iv)
	if err != nil {
		return nil, fmt.Errorf("failed to generate IV: %w", err)
	}
	return iv, nil
}

func (s *AESEncryptor) pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

// Функция для удаления PKCS7 паддинга
func (s *AESEncryptor) pkcs7UnPadding(src []byte) []byte {
	length := len(src)
	padding := int(src[length-1])
	if padding > length {
		return src
	}
	return src[:length-padding]
}
