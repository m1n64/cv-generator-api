package services

import (
	"ai-service/pkg/utils"
	"github.com/goccy/go-json"
	"go.uber.org/zap"
	"os"
	"sync/atomic"
	"time"
)

type AiConfig struct {
	BasePrompt       string  `json:"base_prompt"`
	Temperature      float32 `json:"temperature"`
	TopP             float32 `json:"top_p"`
	FrequencyPenalty float32 `json:"frequency_penalty"`
}

type ConfigManager struct {
	config atomic.Value
	path   string
	hash   string
	logger *zap.Logger
}

func NewConfigManager(path string, logger *zap.Logger) (*ConfigManager, error) {
	manager := &ConfigManager{path: path, logger: logger}
	if err := manager.reload(); err != nil {
		logger.Error("error reloading config", zap.Error(err))
		return nil, err
	}
	go manager.watch()

	manager.logger.Info("config manager initialized", zap.String("path", path))

	return manager, nil
}

func (m *ConfigManager) GetConfig() *AiConfig {
	return m.config.Load().(*AiConfig)
}

func (m *ConfigManager) reload() error {
	file, err := os.ReadFile(m.path)
	if err != nil {
		return err
	}

	newHash := utils.HashData(file)
	if newHash == m.hash {
		return nil
	}

	var newConfig AiConfig
	if err := m.parseConfig(file, &newConfig); err != nil {
		return err
	}

	m.config.Store(&newConfig)
	m.hash = newHash

	m.logger.Info("config reloaded", zap.String("path", m.path))

	return nil
}

func (m *ConfigManager) watch() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	m.logger.Info("config manager watch started", zap.String("path", m.path))

	for range ticker.C {
		if err := m.reload(); err != nil {
			m.logger.Error("error reloading config", zap.Error(err))
		}
	}
}

func (m *ConfigManager) parseConfig(data []byte, config *AiConfig) error {
	return json.Unmarshal(data, config)
}
