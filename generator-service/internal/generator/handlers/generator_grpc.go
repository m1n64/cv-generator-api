package handlers

import (
	"context"
	"cv-generator-service/internal/generator/enums"
	generator "cv-generator-service/internal/generator/grpc"
	"cv-generator-service/internal/generator/models"
	"cv-generator-service/internal/generator/services"
	"cv-generator-service/pkg/utils"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GeneratorServiceServer struct {
	generator.UnimplementedGeneratorServiceServer
	generatorService *services.PdfGeneratorService
	minio            *utils.MinioClient
	logger           *zap.Logger
}

func NewGeneratorServiceServer(generatorService *services.PdfGeneratorService, minio *utils.MinioClient, logger *zap.Logger) *GeneratorServiceServer {
	return &GeneratorServiceServer{
		generatorService: generatorService,
		minio:            minio,
		logger:           logger,
	}
}

func (s *GeneratorServiceServer) GetAllListGenerated(ctx context.Context, req *generator.AllListGeneratedRequest) (*generator.ListGeneratedPdf, error) {
	if uuid.Validate(req.UserId) != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid user id")
	}

	fs, err := s.generatorService.GetUserGeneratedPDFs(uuid.MustParse(req.UserId))
	if err != nil {
		s.logger.Error("error getting user generated pdfs", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	var pdfs []*generator.GeneratedPdf
	for _, f := range fs {
		pdfs = append(pdfs, s.getGeneratedResponse(ctx, f, false))
	}

	return &generator.ListGeneratedPdf{Pdfs: pdfs}, nil
}

func (s *GeneratorServiceServer) GetListGenerated(ctx context.Context, req *generator.GeneratedRequest) (*generator.ListGeneratedPdf, error) {
	if uuid.Validate(req.UserId) != nil || uuid.Validate(req.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid user id or cv id")
	}

	pdf, err := s.generatorService.GetGeneratedPDFsByCvID(uuid.MustParse(req.UserId), uuid.MustParse(req.CvId))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var pdfs []*generator.GeneratedPdf
	for _, f := range pdf {
		pdfs = append(pdfs, s.getGeneratedResponse(ctx, f, false))
	}

	return &generator.ListGeneratedPdf{Pdfs: pdfs}, nil
}

func (s *GeneratorServiceServer) GetGeneratedPDF(ctx context.Context, req *generator.GeneratedPDFRequest) (*generator.GeneratedPdf, error) {
	if uuid.Validate(req.Id) != nil || uuid.Validate(req.UserId) != nil || uuid.Validate(req.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid id, user id or cv id")
	}

	pdf, err := s.generatorService.GetGeneratedPDF(uuid.MustParse(req.Id), uuid.MustParse(req.UserId), uuid.MustParse(req.CvId))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return s.getGeneratedResponse(ctx, pdf, true), nil
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

func (s *GeneratorServiceServer) GetPDFLink(ctx context.Context, req *generator.GeneratedPDFRequest) (*generator.PDFLink, error) {
	if uuid.Validate(req.Id) != nil || uuid.Validate(req.UserId) != nil || uuid.Validate(req.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid user id or cv id")
	}

	pdf, err := s.generatorService.GetGeneratedPDF(uuid.MustParse(req.Id), uuid.MustParse(req.UserId), uuid.MustParse(req.CvId))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if pdf.FileOrigin == nil || pdf.Status != enums.StatusCompleted {
		return nil, status.Error(codes.NotFound, "pdf file not found")
	}

	pdfFile, pdfLink := s.getPdfFile(ctx, pdf.FileOrigin)

	return &generator.PDFLink{
		Id:      pdf.ID.String(),
		Title:   pdf.Title,
		PdfFile: pdfFile,
		PdfUrl:  pdfLink,
	}, nil
}

func (s *GeneratorServiceServer) getGeneratedResponse(ctx context.Context, model *models.GeneratedPdf, generatePdf bool) *generator.GeneratedPdf {
	var pdfFile []byte

	pdfLink := s.getPdfFileUrl(ctx, model.FileOrigin)

	if generatePdf {
		pdfFile = s.getPdfFile(ctx, model.FileOrigin)
	}

	return &generator.GeneratedPdf{
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

func (s *GeneratorServiceServer) getPdfFile(ctx context.Context, fileOrigin *string) []byte {
	var pdfFile []byte

	if fileOrigin != nil {
		var err error
		pdfFile, err = s.minio.GetFileAsBytes(ctx, *fileOrigin)
		if err != nil {
			s.logger.Error("Error getting PDF file", zap.Error(err))
		}
	}

	return pdfFile
}

func (s *GeneratorServiceServer) getPdfFileUrl(ctx context.Context, fileOrigin *string) *string {
	var pdfLink *string

	if fileOrigin != nil {
		pdfGenLink, err := s.minio.GetFileURL(ctx, *fileOrigin)
		if err != nil {
			s.logger.Error("Error getting PDF link", zap.Error(err))
		}

		pdfLink = &pdfGenLink
	}

	return pdfLink
}
