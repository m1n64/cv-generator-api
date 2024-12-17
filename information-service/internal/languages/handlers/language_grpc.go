package handlers

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	languages "information-service/internal/languages/grpc"
	"information-service/internal/languages/services"
)

type LanguageServiceServer struct {
	languages.UnimplementedLanguagesServiceServer
	languageService *services.LanguageService
}

func NewLanguageServiceServer(languageService *services.LanguageService) *LanguageServiceServer {
	return &LanguageServiceServer{
		languageService: languageService,
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
		langs = append(langs, &languages.LanguageResponse{
			Id:        lang.ID.String(),
			CvId:      lang.CvID.String(),
			Name:      lang.Name,
			Level:     lang.Level,
			CreatedAt: lang.CreatedAt.String(),
			UpdatedAt: lang.UpdatedAt.String(),
		})
	}
	return &languages.AllLanguagesResponse{
		Languages: langs,
	}, nil
}
