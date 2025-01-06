// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.4
// source: proto/work_experience.proto

package experiences

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ExperiencesService_GetExperiences_FullMethodName       = "/experiences.ExperiencesService/GetExperiences"
	ExperiencesService_GetExperienceByID_FullMethodName    = "/experiences.ExperiencesService/GetExperienceByID"
	ExperiencesService_CreateExperience_FullMethodName     = "/experiences.ExperiencesService/CreateExperience"
	ExperiencesService_DeleteExperienceByID_FullMethodName = "/experiences.ExperiencesService/DeleteExperienceByID"
	ExperiencesService_UpdateExperienceByID_FullMethodName = "/experiences.ExperiencesService/UpdateExperienceByID"
)

// ExperiencesServiceClient is the client API for ExperiencesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExperiencesServiceClient interface {
	GetExperiences(ctx context.Context, in *GetExperiencesRequest, opts ...grpc.CallOption) (*AllExperiencesResponse, error)
	GetExperienceByID(ctx context.Context, in *GetExperienceByIDRequest, opts ...grpc.CallOption) (*ExperienceResponse, error)
	CreateExperience(ctx context.Context, in *CreateExperienceRequest, opts ...grpc.CallOption) (*ExperienceResponse, error)
	DeleteExperienceByID(ctx context.Context, in *DeleteExperienceByIDRequest, opts ...grpc.CallOption) (*DeleteExperienceByIDResponse, error)
	UpdateExperienceByID(ctx context.Context, in *UpdateExperienceByIDRequest, opts ...grpc.CallOption) (*ExperienceResponse, error)
}

type experiencesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExperiencesServiceClient(cc grpc.ClientConnInterface) ExperiencesServiceClient {
	return &experiencesServiceClient{cc}
}

func (c *experiencesServiceClient) GetExperiences(ctx context.Context, in *GetExperiencesRequest, opts ...grpc.CallOption) (*AllExperiencesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AllExperiencesResponse)
	err := c.cc.Invoke(ctx, ExperiencesService_GetExperiences_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *experiencesServiceClient) GetExperienceByID(ctx context.Context, in *GetExperienceByIDRequest, opts ...grpc.CallOption) (*ExperienceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExperienceResponse)
	err := c.cc.Invoke(ctx, ExperiencesService_GetExperienceByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *experiencesServiceClient) CreateExperience(ctx context.Context, in *CreateExperienceRequest, opts ...grpc.CallOption) (*ExperienceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExperienceResponse)
	err := c.cc.Invoke(ctx, ExperiencesService_CreateExperience_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *experiencesServiceClient) DeleteExperienceByID(ctx context.Context, in *DeleteExperienceByIDRequest, opts ...grpc.CallOption) (*DeleteExperienceByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteExperienceByIDResponse)
	err := c.cc.Invoke(ctx, ExperiencesService_DeleteExperienceByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *experiencesServiceClient) UpdateExperienceByID(ctx context.Context, in *UpdateExperienceByIDRequest, opts ...grpc.CallOption) (*ExperienceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExperienceResponse)
	err := c.cc.Invoke(ctx, ExperiencesService_UpdateExperienceByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExperiencesServiceServer is the server API for ExperiencesService service.
// All implementations must embed UnimplementedExperiencesServiceServer
// for forward compatibility.
type ExperiencesServiceServer interface {
	GetExperiences(context.Context, *GetExperiencesRequest) (*AllExperiencesResponse, error)
	GetExperienceByID(context.Context, *GetExperienceByIDRequest) (*ExperienceResponse, error)
	CreateExperience(context.Context, *CreateExperienceRequest) (*ExperienceResponse, error)
	DeleteExperienceByID(context.Context, *DeleteExperienceByIDRequest) (*DeleteExperienceByIDResponse, error)
	UpdateExperienceByID(context.Context, *UpdateExperienceByIDRequest) (*ExperienceResponse, error)
	mustEmbedUnimplementedExperiencesServiceServer()
}

// UnimplementedExperiencesServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedExperiencesServiceServer struct{}

func (UnimplementedExperiencesServiceServer) GetExperiences(context.Context, *GetExperiencesRequest) (*AllExperiencesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExperiences not implemented")
}
func (UnimplementedExperiencesServiceServer) GetExperienceByID(context.Context, *GetExperienceByIDRequest) (*ExperienceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExperienceByID not implemented")
}
func (UnimplementedExperiencesServiceServer) CreateExperience(context.Context, *CreateExperienceRequest) (*ExperienceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExperience not implemented")
}
func (UnimplementedExperiencesServiceServer) DeleteExperienceByID(context.Context, *DeleteExperienceByIDRequest) (*DeleteExperienceByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteExperienceByID not implemented")
}
func (UnimplementedExperiencesServiceServer) UpdateExperienceByID(context.Context, *UpdateExperienceByIDRequest) (*ExperienceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateExperienceByID not implemented")
}
func (UnimplementedExperiencesServiceServer) mustEmbedUnimplementedExperiencesServiceServer() {}
func (UnimplementedExperiencesServiceServer) testEmbeddedByValue()                            {}

// UnsafeExperiencesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExperiencesServiceServer will
// result in compilation errors.
type UnsafeExperiencesServiceServer interface {
	mustEmbedUnimplementedExperiencesServiceServer()
}

func RegisterExperiencesServiceServer(s grpc.ServiceRegistrar, srv ExperiencesServiceServer) {
	// If the following call pancis, it indicates UnimplementedExperiencesServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ExperiencesService_ServiceDesc, srv)
}

func _ExperiencesService_GetExperiences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExperiencesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperiencesServiceServer).GetExperiences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExperiencesService_GetExperiences_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperiencesServiceServer).GetExperiences(ctx, req.(*GetExperiencesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExperiencesService_GetExperienceByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExperienceByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperiencesServiceServer).GetExperienceByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExperiencesService_GetExperienceByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperiencesServiceServer).GetExperienceByID(ctx, req.(*GetExperienceByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExperiencesService_CreateExperience_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExperienceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperiencesServiceServer).CreateExperience(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExperiencesService_CreateExperience_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperiencesServiceServer).CreateExperience(ctx, req.(*CreateExperienceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExperiencesService_DeleteExperienceByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteExperienceByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperiencesServiceServer).DeleteExperienceByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExperiencesService_DeleteExperienceByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperiencesServiceServer).DeleteExperienceByID(ctx, req.(*DeleteExperienceByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExperiencesService_UpdateExperienceByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateExperienceByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperiencesServiceServer).UpdateExperienceByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExperiencesService_UpdateExperienceByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperiencesServiceServer).UpdateExperienceByID(ctx, req.(*UpdateExperienceByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExperiencesService_ServiceDesc is the grpc.ServiceDesc for ExperiencesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExperiencesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "experiences.ExperiencesService",
	HandlerType: (*ExperiencesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetExperiences",
			Handler:    _ExperiencesService_GetExperiences_Handler,
		},
		{
			MethodName: "GetExperienceByID",
			Handler:    _ExperiencesService_GetExperienceByID_Handler,
		},
		{
			MethodName: "CreateExperience",
			Handler:    _ExperiencesService_CreateExperience_Handler,
		},
		{
			MethodName: "DeleteExperienceByID",
			Handler:    _ExperiencesService_DeleteExperienceByID_Handler,
		},
		{
			MethodName: "UpdateExperienceByID",
			Handler:    _ExperiencesService_UpdateExperienceByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/work_experience.proto",
}
