package services

import (
	"ai-service/internal/ai/entites"
	ai "ai-service/internal/ai/grpc"
	"ai-service/internal/ai/repositories"
	"context"
	"errors"
	"fmt"
	"github.com/cohesion-org/deepseek-go"
	"github.com/cohesion-org/deepseek-go/constants"
	"io"
	"time"
)

type AiService struct {
	aiServiceRepository repositories.AiServicesRepository
	aiAnalyticsService  *AiAnalyticsService
	deepSeek            *deepseek.Client
	configManager       *ConfigManager
}

func NewAiService(aiServiceRepository repositories.AiServicesRepository, aiAnalyticsService *AiAnalyticsService, deepSeek *deepseek.Client, configManager *ConfigManager) *AiService {
	return &AiService{
		aiServiceRepository: aiServiceRepository,
		aiAnalyticsService:  aiAnalyticsService,
		deepSeek:            deepSeek,
		configManager:       configManager,
	}
}

func (s *AiService) GetServices() ([]*entites.Service, error) {
	return s.aiServiceRepository.GetServices()
}

func (s *AiService) GenerateDescription(prompt string, serviceId string) (string, error) {
	service, err := s.getServiceName(serviceId)
	if err != nil {
		return "", err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	resultChan := make(chan string, 1)
	errorChan := make(chan error, 1)

	go func() {
		completion, err := s.deepSeek.CreateChatCompletion(ctx, &deepseek.ChatCompletionRequest{
			Model:            deepseek.DeepSeekChat,
			Messages:         s.getDeepSeekMessages(prompt, service),
			Temperature:      s.configManager.GetConfig().Temperature,
			TopP:             s.configManager.GetConfig().TopP,
			FrequencyPenalty: s.configManager.GetConfig().FrequencyPenalty,
		})
		if err != nil {
			errorChan <- err
			return
		}
		resultChan <- completion.Choices[0].Message.Content
	}()

	select {
	case res := <-resultChan:
		err = s.aiAnalyticsService.SendAiAnalyticsGenerateEvent(prompt, res, service)
		return res, nil
	case err := <-errorChan:
		return "", err
	case <-ctx.Done():
		return "", errors.New("request to AI service timed out")
	}
}

func (s *AiService) StreamGenerateDescription(prompt string, serviceId string, stream ai.AiService_StreamGenerateServer) error {
	service, err := s.getServiceName(serviceId)
	if err != nil {
		return err
	}

	ctx := context.Background()

	deepSeekStream, err := s.deepSeek.CreateChatCompletionStream(ctx, &deepseek.StreamChatCompletionRequest{
		Model:            deepseek.DeepSeekChat,
		Messages:         s.getDeepSeekMessages(prompt, service),
		Stream:           true,
		Temperature:      s.configManager.GetConfig().Temperature,
		TopP:             s.configManager.GetConfig().TopP,
		FrequencyPenalty: s.configManager.GetConfig().FrequencyPenalty,
	})
	if err != nil {
		return fmt.Errorf("failed to create stream with DeepSeek: %w", err)
	}

	defer deepSeekStream.Close()

	var responseText string
	for {
		response, err := deepSeekStream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return fmt.Errorf("error receiving from DeepSeek stream: %w", err)
		}

		for _, choice := range response.Choices {
			err := stream.Send(&ai.GenerateResponse{
				Response: choice.Delta.Content,
			})
			responseText += choice.Delta.Content
			if err != nil {
				return fmt.Errorf("error sending stream to client: %w", err)
			}
		}
	}

	err = s.aiAnalyticsService.SendAiAnalyticsGenerateEvent(prompt, responseText, service)

	return nil
}

func (s *AiService) getServiceName(serviceId string) (string, error) {
	return s.aiServiceRepository.GetService(serviceId)
}

func (s *AiService) getDeepSeekMessages(prompt string, service string) []deepseek.ChatCompletionMessage {
	return []deepseek.ChatCompletionMessage{
		{Role: constants.ChatMessageRoleSystem, Content: fmt.Sprintf(s.configManager.GetConfig().BasePrompt, service)},
		{Role: constants.ChatMessageRoleUser, Content: prompt},
	}
}
