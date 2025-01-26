package repositories

import "ai-service/internal/ai/entites"

type AiServicesRepository interface {
	GetServices() ([]*entites.Service, error)
	GetService(serviceId string) (string, error)
}
