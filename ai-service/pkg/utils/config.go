package utils

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		GetLogger().Sugar().Error("Error loading .env file")
	}
}

func UpdateEnvValue(key, value, filePath string) error {
	inputFile, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open .env file: %w", err)
	}
	defer inputFile.Close()

	var lines []string
	keyFound := false

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, key+"=") {
			lines = append(lines, fmt.Sprintf("%s=%s", key, value))
			keyFound = true
		} else {
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to read .env file: %w", err)
	}

	if !keyFound {
		lines = append(lines, fmt.Sprintf("%s=%s", key, value))
	}

	outputFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to open .env file for writing: %w", err)
	}
	defer outputFile.Close()

	for _, line := range lines {
		_, err := outputFile.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("failed to write to .env file: %w", err)
		}
	}

	return nil
}
