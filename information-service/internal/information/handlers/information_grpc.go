package handlers

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	information "information-service/internal/information/grpc"
	"information-service/internal/information/services"
)

type CVInformationServiceServer struct {
	information.UnimplementedInformationServiceServer
	infoService *services.CVInformationService
	logger      *zap.Logger
}

func NewCVInformationServiceServer(infoService *services.CVInformationService, logger *zap.Logger) *CVInformationServiceServer {
	return &CVInformationServiceServer{
		infoService: infoService,
		logger:      logger,
	}
}

func (s *CVInformationServiceServer) CreateOrUpdateInformation(ctx context.Context, req *information.CreateOrUpdateInformationRequest) (*information.InformationResponse, error) {
	if req.FullName == "" || uuid.Validate(req.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "full_name and cv_id is required and must be a valid uuid")
	}

	cvID := uuid.MustParse(req.CvId)

	info, err := s.infoService.CreateOrUpdateCV(cvID, req.FullName, req.PhotoFileId, req.Position, req.Location, req.Biography)
	if err != nil {
		s.logger.Info(fmt.Sprintf("Error creating or updating information: %s", err.Error()))
		return nil, err
	}

	return &information.InformationResponse{
		Id:          info.ID.String(),
		CvId:        info.CvID.String(),
		FullName:    info.FullName,
		PhotoFileId: info.PhotoFileID,
		Position:    info.Position,
		Location:    info.Location,
		Biography:   info.Biography,
	}, nil
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

	return &information.InformationResponse{
		Id:          info.ID.String(),
		CvId:        info.CvID.String(),
		FullName:    info.FullName,
		PhotoFileId: info.PhotoFileID,
		Position:    info.Position,
		Location:    info.Location,
		Biography:   info.Biography,
	}, nil
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
