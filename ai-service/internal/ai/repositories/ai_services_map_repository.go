package repositories

import (
	"ai-service/internal/ai/entites"
	"errors"
)

type aiServicesMapRepository struct {
	serviceMap map[string]string
}

var services = []*entites.Service{
	{ID: "info", Name: "Information"},
	{ID: "certs", Name: "Certificate"},
	{ID: "edu", Name: "Education"},
	{ID: "exp", Name: "Work Experience"},
}

func NewAiServicesMapRepository() AiServicesRepository {
	serviceMap := make(map[string]string)
	for _, service := range services {
		serviceMap[service.ID] = service.Name
	}

	return &aiServicesMapRepository{
		serviceMap: serviceMap,
	}
}

func (r *aiServicesMapRepository) GetServices() ([]*entites.Service, error) {
	return services, nil
}

func (r *aiServicesMapRepository) GetService(serviceId string) (string, error) {
	if name, exists := r.serviceMap[serviceId]; exists {
		return name, nil
	}

	return "", errors.New("service not found")
}
