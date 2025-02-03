package handlers

import (
	ai "ai-service/internal/ai/grpc"
	"ai-service/internal/ai/services"
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AiServiceServer struct {
	ai.UnimplementedAiServiceServer
	aiService *services.AiService
	logger    *zap.Logger
}

func NewAiServiceServer(aiService *services.AiService, logger *zap.Logger) *AiServiceServer {
	return &AiServiceServer{
		aiService: aiService,
		logger:    logger,
	}
}

func (s *AiServiceServer) Generate(ctx context.Context, req *ai.GenerateRequest) (*ai.GenerateResponse, error) {
	if req.Prompt == "" || req.ServiceId == "" {
		return nil, status.Error(codes.InvalidArgument, "prompt and service id cannot be empty")
	}

	description, err := s.aiService.GenerateDescription(req.Prompt, req.ServiceId)
	if err != nil {
		s.logger.Error("error generating description", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &ai.GenerateResponse{
		Response: description,
	}, nil
}

func (s *AiServiceServer) StreamGenerate(req *ai.GenerateRequest, stream ai.AiService_StreamGenerateServer) error {
	if req.Prompt == "" || req.ServiceId == "" {
		return status.Error(codes.InvalidArgument, "prompt and service id cannot be empty")
	}

	err := s.aiService.StreamGenerateDescription(req.Prompt, req.ServiceId, stream)
	if err != nil {
		s.logger.Error("error generating description", zap.Error(err))
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

func (s *AiServiceServer) GetServices(ctx context.Context, req *ai.GetServicesRequest) (*ai.GetServicesResponse, error) {
	servicesList, err := s.aiService.GetServices()
	if err != nil {
		s.logger.Error("error getting services", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	var servicesResp []*ai.Services
	for _, service := range servicesList {
		servicesResp = append(servicesResp, &ai.Services{
			ServiceId:   service.ID,
			ServiceName: service.Name,
		})
	}

	return &ai.GetServicesResponse{
		Services: servicesResp,
	}, nil
}
