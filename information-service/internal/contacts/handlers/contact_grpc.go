package handlers

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	contacts "information-service/internal/contacts/grpc"
	"information-service/internal/contacts/models"
	"information-service/internal/contacts/services"
)

type ContactServiceServer struct {
	contacts.UnimplementedContactsServiceServer
	contactService *services.ContactService
	logger         *zap.Logger
}

func NewContactServiceServer(contactService *services.ContactService, logger *zap.Logger) *ContactServiceServer {
	return &ContactServiceServer{
		contactService: contactService,
		logger:         logger,
	}
}

func (s *ContactServiceServer) GetContacts(ctx context.Context, req *contacts.GetContactsRequest) (*contacts.AllContactsResponse, error) {
	if uuid.Validate(req.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid cv id")
	}

	contactsList, err := s.contactService.GetContacts(uuid.MustParse(req.CvId))
	if err != nil {
		s.logger.Error("error getting contacts", zap.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	var contactsResponse []*contacts.ContactResponse
	for _, contact := range contactsList {
		contactsResponse = append(contactsResponse, s.getContactResponse(contact))
	}

	return &contacts.AllContactsResponse{
		Contacts: contactsResponse,
	}, nil
}

func (s *ContactServiceServer) GetContactByID(ctx context.Context, req *contacts.GetContactByIDRequest) (*contacts.ContactResponse, error) {
	if uuid.Validate(req.Id) != nil || uuid.Validate(req.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid id and cv_id")
	}

	contact, err := s.contactService.GetContact(uuid.MustParse(req.Id), uuid.MustParse(req.CvId))
	if err != nil {
		s.logger.Error("error getting contact", zap.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return s.getContactResponse(contact), nil
}

func (s *ContactServiceServer) CreateContact(ctx context.Context, req *contacts.CreateContactRequest) (*contacts.ContactResponse, error) {
	if uuid.Validate(req.CvId) != nil || req.Title == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid cv id")
	}

	contact, err := s.contactService.CreateContact(uuid.MustParse(req.CvId), req.Title, req.Link)
	if err != nil {
		s.logger.Error("error creating contact", zap.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return s.getContactResponse(contact), nil
}

func (s *ContactServiceServer) UpdateContactByID(ctx context.Context, req *contacts.UpdateContactByIDRequest) (*contacts.ContactResponse, error) {
	if uuid.Validate(req.Id) != nil || uuid.Validate(req.CvId) != nil || req.Title == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid id and cv id")
	}

	contact, err := s.contactService.UpdateContact(uuid.MustParse(req.Id), uuid.MustParse(req.CvId), req.Title, req.Link)
	if err != nil {
		s.logger.Error("error updating contact", zap.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return s.getContactResponse(contact), nil
}

func (s *ContactServiceServer) DeleteContactByID(ctx context.Context, req *contacts.DeleteContactByIDRequest) (*contacts.DeleteContactByIDResponse, error) {
	if uuid.Validate(req.Id) != nil || uuid.Validate(req.CvId) != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid id and cv id")
	}

	err := s.contactService.DeleteContact(uuid.MustParse(req.Id), uuid.MustParse(req.CvId))
	if err != nil {
		s.logger.Error("error deleting contact", zap.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &contacts.DeleteContactByIDResponse{
		Success: true,
	}, nil
}

func (s *ContactServiceServer) getContactResponse(contact *models.Contact) *contacts.ContactResponse {
	return &contacts.ContactResponse{
		Id:        contact.ID.String(),
		CvId:      contact.CvID.String(),
		Title:     contact.Title,
		Link:      contact.Link,
		CreatedAt: contact.CreatedAt.String(),
		UpdatedAt: contact.UpdatedAt.String(),
	}
}
