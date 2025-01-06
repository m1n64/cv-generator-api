package handlers

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	languages "information-service/internal/languages/grpc"
	"information-service/internal/languages/models"
	"information-service/internal/languages/services"
)

type LanguageServiceServer struct {
	languages.UnimplementedLanguagesServiceServer
	languageService *services.LanguageService
	logger          *zap.Logger
}

func NewLanguageServiceServer(languageService *services.LanguageService, logger *zap.Logger) *LanguageServiceServer {
	return &LanguageServiceServer{
		languageService: languageService,
		logger:          logger,
	}
}

func (server *LanguageServiceServer) GetLanguages(context context.Context, req *languages.GetLanguagesRequest) (*languages.AllLanguagesResponse, error) {
	if uuid.Validate(req.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "cv_id is required and must be a valid uuid")
	}

	languagesList, err := server.languageService.GetLanguagesByCvID(uuid.MustParse(req.CvId))
	if err != nil {
		return nil, err
	}

	var langs []*languages.LanguageResponse
	for _, lang := range languagesList {
		langs = append(langs, server.getLanguageResponse(lang))
	}
	return &languages.AllLanguagesResponse{
		Languages: langs,
	}, nil
}

func (server *LanguageServiceServer) GetLanguageByID(context context.Context, req *languages.GetLanguageByIDRequest) (*languages.LanguageResponse, error) {
	if uuid.Validate(req.Id) != nil || uuid.Validate(req.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "id and cv_id is required and must be a valid uuid")
	}

	id := uuid.MustParse(req.Id)
	cvID := uuid.MustParse(req.CvId)

	language, err := server.languageService.GetLanguage(id, cvID)
	if err != nil {
		return nil, err
	}

	return server.getLanguageResponse(language), nil
}

func (server *LanguageServiceServer) CreateLanguage(context context.Context, req *languages.CreateLanguageRequest) (*languages.LanguageResponse, error) {
	if uuid.Validate(req.CvId) != nil || req.Name == "" || req.Level == "" {
		return nil, status.Error(codes.InvalidArgument, "cv_id, name and level is required and must be a valid uuid")
	}

	language, err := server.languageService.CreateLanguage(uuid.MustParse(req.CvId), req.Name, req.Level)
	if err != nil {
		return nil, err
	}

	return server.getLanguageResponse(language), nil
}

func (server *LanguageServiceServer) UpdateLanguageByID(context context.Context, req *languages.UpdateLanguageByIDRequest) (*languages.LanguageResponse, error) {
	if uuid.Validate(req.Id) != nil || uuid.Validate(req.CvId) != nil || req.Name == "" || req.Level == "" {
		return nil, status.Error(codes.InvalidArgument, "id, cv_id, name and level is required and must be a valid uuid")
	}

	language, err := server.languageService.UpdateLanguage(uuid.MustParse(req.Id), uuid.MustParse(req.CvId), req.Name, req.Level)
	if err != nil {
		return nil, err
	}

	return server.getLanguageResponse(language), nil
}

func (server *LanguageServiceServer) DeleteLanguageByID(context context.Context, req *languages.DeleteLanguageByIDRequest) (*languages.DeleteLanguageByIDResponse, error) {
	if uuid.Validate(req.Id) != nil || uuid.Validate(req.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "id and cv_id is required and must be a valid uuid")
	}

	err := server.languageService.DeleteLanguage(uuid.MustParse(req.Id), uuid.MustParse(req.CvId))
	if err != nil {
		return nil, err
	}

	return &languages.DeleteLanguageByIDResponse{
		Success: true,
	}, nil
}

func (server *LanguageServiceServer) getLanguageResponse(language *models.Language) *languages.LanguageResponse {
	return &languages.LanguageResponse{
		Id:        language.ID.String(),
		CvId:      language.CvID.String(),
		Name:      language.Name,
		Level:     language.Level,
		CreatedAt: language.CreatedAt.String(),
		UpdatedAt: language.UpdatedAt.String(),
	}
}
