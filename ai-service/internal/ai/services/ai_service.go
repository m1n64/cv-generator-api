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

const Temperature = 1.35

type AiService struct {
	aiServiceRepository repositories.AiServicesRepository
	deepSeek            *deepseek.Client
}

func NewAiService(aiServiceRepository repositories.AiServicesRepository, deepSeek *deepseek.Client) *AiService {
	return &AiService{
		aiServiceRepository: aiServiceRepository,
		deepSeek:            deepSeek,
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
			Model:       deepseek.DeepSeekChat,
			Messages:    s.getDeepSeekMessages(prompt, service),
			Temperature: Temperature,
		})
		if err != nil {
			errorChan <- err
			return
		}
		resultChan <- completion.Choices[0].Message.Content
	}()

	select {
	case res := <-resultChan:
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
		Model:       deepseek.DeepSeekChat,
		Messages:    s.getDeepSeekMessages(prompt, service),
		Stream:      true,
		Temperature: Temperature,
	})
	if err != nil {
		return fmt.Errorf("failed to create stream with DeepSeek: %w", err)
	}

	defer deepSeekStream.Close()

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
			if err != nil {
				return fmt.Errorf("error sending stream to client: %w", err)
			}
		}
	}

	return nil
}

func (s *AiService) getServiceName(serviceId string) (string, error) {
	return s.aiServiceRepository.GetService(serviceId)
}

func (s *AiService) getDeepSeekMessages(prompt string, service string) []deepseek.ChatCompletionMessage {
	return []deepseek.ChatCompletionMessage{
		{Role: constants.ChatMessageRoleSystem, Content: fmt.Sprintf("Generate description for the '%s' section of a resume. Focus on accomplishments, tools, and impact.", service)},
		{Role: constants.ChatMessageRoleUser, Content: prompt},
	}
}