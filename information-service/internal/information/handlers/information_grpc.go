package handlers

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	services2 "information-service/internal/general/services"
	information "information-service/internal/information/grpc"
	"information-service/internal/information/models"
	"information-service/internal/information/services"
)

type CVInformationServiceServer struct {
	information.UnimplementedInformationServiceServer
	infoService *services.CVInformationService
	fileService *services2.FileService
	logger      *zap.Logger
}

func NewCVInformationServiceServer(infoService *services.CVInformationService, fileService *services2.FileService, logger *zap.Logger) *CVInformationServiceServer {
	return &CVInformationServiceServer{
		infoService: infoService,
		fileService: fileService,
		logger:      logger,
	}
}

func (s *CVInformationServiceServer) CreateOrUpdateInformation(ctx context.Context, req *information.CreateOrUpdateInformationRequest) (*information.InformationResponse, error) {
	if req.FullName == "" || uuid.Validate(req.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "full_name and cv_id is required and must be a valid uuid")
	}

	cvID := uuid.MustParse(req.CvId)

	var photoFileID *string
	if req.Photo != nil {
		fileID, err := s.fileService.SaveFile(ctx, cvID, req.Photo)
		if err != nil {
			s.logger.Info(fmt.Sprintf("Error uploading file: %s", err.Error()))
			return nil, err
		}
		photoFileID = fileID
	}

	info, err := s.infoService.CreateOrUpdateCV(cvID, req.FullName, photoFileID, req.Position, req.Location, req.Biography)
	if err != nil {
		s.logger.Info(fmt.Sprintf("Error creating or updating information: %s", err.Error()))
		return nil, err
	}

	return s.getResponse(ctx, info), nil
}

func (s *CVInformationServiceServer) GetInformationByCvID(ctx context.Context, req *information.GetInformationByCvIDRequest) (*information.InformationResponse, error) {
	if uuid.Validate(req.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "cv_id is required and must be a valid uuid")
	}

	cvID := uuid.MustParse(req.CvId)

	info, err := s.infoService.GetCVInformation(cvID)
	if err != nil {
		s.logger.Info(fmt.Sprintf("Error getting information: %s", err.Error()))
		return nil, err
	}

	return s.getResponse(ctx, info), nil
}

func (s *CVInformationServiceServer) DeleteInformationByCvID(ctx context.Context, req *information.DeleteInformationByCvIDRequest) (*information.DeleteInformationByCvIDResponse, error) {
	if uuid.Validate(req.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "cv_id is required and must be a valid uuid")
	}

	cvID := uuid.MustParse(req.CvId)

	err := s.infoService.DeleteInformation(cvID)
	if err != nil {
		s.logger.Info(fmt.Sprintf("Error deleting information: %s", err.Error()))
		return nil, err
	}

	return &information.DeleteInformationByCvIDResponse{
		Success: true,
	}, nil
}

func (s *CVInformationServiceServer) getPhotoFromInformation(ctx context.Context, info *models.Information) ([]byte, string) {
	if info.PhotoFileID == nil {
		return nil, ""
	}

	imageFile, err := s.fileService.GetFileAsBytes(ctx, *info.PhotoFileID)

	if err != nil {
		imageFile = nil
	}

	fileUrl, err := s.fileService.GetFileURL(ctx, *info.PhotoFileID)
	if err != nil {
		fileUrl = ""
	}

	return imageFile, fileUrl
}

func (s *CVInformationServiceServer) getResponse(ctx context.Context, info *models.Information) *information.InformationResponse {
	imageFile, fileUrl := s.getPhotoFromInformation(ctx, info)

	return &information.InformationResponse{
		Id:        info.ID.String(),
		CvId:      info.CvID.String(),
		FullName:  info.FullName,
		PhotoFile: imageFile,
		PhotoUrl:  &fileUrl,
		Position:  info.Position,
		Location:  info.Location,
		Biography: info.Biography,
	}
}
