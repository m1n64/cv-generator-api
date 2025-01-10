package handlers

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	generator "information-service/internal/generator/grpc"
	"information-service/internal/generator/models"
	"information-service/internal/generator/services"
	"information-service/pkg/utils"
)

type GeneratorServiceServer struct {
	generator.UnimplementedGeneratorServiceServer
	generatorService *services.PdfGeneratorService
	minio            *utils.MinioClient
	logger           *zap.Logger
}

func NewGeneratorServiceServer(generatorService *services.PdfGeneratorService, logger *zap.Logger) *GeneratorServiceServer {
	return &GeneratorServiceServer{
		generatorService: generatorService,
		logger:           logger,
	}
}

func (s *GeneratorServiceServer) GetAllListGenerated(ctx context.Context, req *generator.AllListGeneratedRequest) (*generator.ListGeneratedPDF, error) {
	if uuid.Validate(req.UserId) != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid user id")
	}

	fs, err := s.generatorService.GetUserGeneratedPDFs(uuid.MustParse(req.UserId))
	if err != nil {
		s.logger.Error("error getting user generated pdfs", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	var pdfs []*generator.GeneratedPDF
	for _, f := range fs {
		pdfs = append(pdfs, s.getGeneratedResponse(ctx, f))
	}

	return &generator.ListGeneratedPDF{Pdfs: pdfs}, nil
}

func (s *GeneratorServiceServer) GetListGenerated(ctx context.Context, req *generator.GeneratedRequest) (*generator.ListGeneratedPDF, error) {
	if uuid.Validate(req.UserId) != nil || uuid.Validate(req.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid user id or cv id")
	}

	pdf, err := s.generatorService.GetGeneratedPDFsByCvID(uuid.MustParse(req.UserId), uuid.MustParse(req.CvId))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var pdfs []*generator.GeneratedPDF
	for _, f := range pdf {
		pdfs = append(pdfs, s.getGeneratedResponse(ctx, f))
	}

	return &generator.ListGeneratedPDF{Pdfs: pdfs}, nil
}

func (s *GeneratorServiceServer) GetGeneratedPDF(ctx context.Context, req *generator.GeneratedPDFRequest) (*generator.GeneratedPDF, error) {
	if uuid.Validate(req.Id) != nil || uuid.Validate(req.UserId) != nil || uuid.Validate(req.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid id, user id or cv id")
	}

	pdf, err := s.generatorService.GetGeneratedPDF(uuid.MustParse(req.Id), uuid.MustParse(req.UserId), uuid.MustParse(req.CvId))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return s.getGeneratedResponse(ctx, pdf), nil
}

func (s *GeneratorServiceServer) DeleteGenerated(ctx context.Context, req *generator.GeneratedPDFRequest) (*generator.DeleteGeneratedResponse, error) {
	if uuid.Validate(req.Id) != nil || uuid.Validate(req.UserId) != nil || uuid.Validate(req.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid id, user id or cv id")
	}

	err := s.generatorService.DeleteGeneratedPDF(uuid.MustParse(req.Id), uuid.MustParse(req.UserId), uuid.MustParse(req.CvId))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &generator.DeleteGeneratedResponse{Success: true}, nil
}

func (s *GeneratorServiceServer) getGeneratedResponse(ctx context.Context, model *models.GeneratedPDF) *generator.GeneratedPDF {
	pdfFile, err := s.minio.GetFileAsBytes(ctx, model.FileOrigin)
	if err != nil {
		s.logger.Error("error getting pdf file", zap.Error(err))
		return nil
	}

	pdfLink, err := s.minio.GetFileURL(ctx, model.FileOrigin)
	if err != nil {
		s.logger.Error("error getting pdf link", zap.Error(err))
		return nil
	}

	return &generator.GeneratedPDF{
		Id:        model.ID.String(),
		CvId:      model.CvID.String(),
		UserId:    model.UserID.String(),
		Title:     model.Title,
		PdfFile:   pdfFile,
		PdfUrl:    pdfLink,
		Status:    model.Status.String(),
		CreatedAt: model.CreatedAt.String(),
		UpdatedAt: model.UpdatedAt.String(),
	}
}
