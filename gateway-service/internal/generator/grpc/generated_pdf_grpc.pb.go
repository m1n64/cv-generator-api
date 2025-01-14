// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.4
// source: proto/generated_pdf.proto

package generator

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
	GeneratorService_GetAllListGenerated_FullMethodName = "/generator.GeneratorService/GetAllListGenerated"
	GeneratorService_GetListGenerated_FullMethodName    = "/generator.GeneratorService/GetListGenerated"
	GeneratorService_GetGeneratedPDF_FullMethodName     = "/generator.GeneratorService/GetGeneratedPDF"
	GeneratorService_DeleteGenerated_FullMethodName     = "/generator.GeneratorService/DeleteGenerated"
	GeneratorService_GetPDFLink_FullMethodName          = "/generator.GeneratorService/GetPDFLink"
)

// GeneratorServiceClient is the client API for GeneratorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeneratorServiceClient interface {
	GetAllListGenerated(ctx context.Context, in *AllListGeneratedRequest, opts ...grpc.CallOption) (*ListGeneratedPdf, error)
	GetListGenerated(ctx context.Context, in *GeneratedRequest, opts ...grpc.CallOption) (*ListGeneratedPdf, error)
	GetGeneratedPDF(ctx context.Context, in *GeneratedPDFRequest, opts ...grpc.CallOption) (*GeneratedPdf, error)
	DeleteGenerated(ctx context.Context, in *GeneratedPDFRequest, opts ...grpc.CallOption) (*DeleteGeneratedResponse, error)
	GetPDFLink(ctx context.Context, in *GeneratedPDFRequest, opts ...grpc.CallOption) (*PDFLink, error)
}

type generatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGeneratorServiceClient(cc grpc.ClientConnInterface) GeneratorServiceClient {
	return &generatorServiceClient{cc}
}

func (c *generatorServiceClient) GetAllListGenerated(ctx context.Context, in *AllListGeneratedRequest, opts ...grpc.CallOption) (*ListGeneratedPdf, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListGeneratedPdf)
	err := c.cc.Invoke(ctx, GeneratorService_GetAllListGenerated_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generatorServiceClient) GetListGenerated(ctx context.Context, in *GeneratedRequest, opts ...grpc.CallOption) (*ListGeneratedPdf, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListGeneratedPdf)
	err := c.cc.Invoke(ctx, GeneratorService_GetListGenerated_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generatorServiceClient) GetGeneratedPDF(ctx context.Context, in *GeneratedPDFRequest, opts ...grpc.CallOption) (*GeneratedPdf, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GeneratedPdf)
	err := c.cc.Invoke(ctx, GeneratorService_GetGeneratedPDF_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generatorServiceClient) DeleteGenerated(ctx context.Context, in *GeneratedPDFRequest, opts ...grpc.CallOption) (*DeleteGeneratedResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteGeneratedResponse)
	err := c.cc.Invoke(ctx, GeneratorService_DeleteGenerated_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generatorServiceClient) GetPDFLink(ctx context.Context, in *GeneratedPDFRequest, opts ...grpc.CallOption) (*PDFLink, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PDFLink)
	err := c.cc.Invoke(ctx, GeneratorService_GetPDFLink_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeneratorServiceServer is the server API for GeneratorService service.
// All implementations must embed UnimplementedGeneratorServiceServer
// for forward compatibility.
type GeneratorServiceServer interface {
	GetAllListGenerated(context.Context, *AllListGeneratedRequest) (*ListGeneratedPdf, error)
	GetListGenerated(context.Context, *GeneratedRequest) (*ListGeneratedPdf, error)
	GetGeneratedPDF(context.Context, *GeneratedPDFRequest) (*GeneratedPdf, error)
	DeleteGenerated(context.Context, *GeneratedPDFRequest) (*DeleteGeneratedResponse, error)
	GetPDFLink(context.Context, *GeneratedPDFRequest) (*PDFLink, error)
	mustEmbedUnimplementedGeneratorServiceServer()
}

// UnimplementedGeneratorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGeneratorServiceServer struct{}

func (UnimplementedGeneratorServiceServer) GetAllListGenerated(context.Context, *AllListGeneratedRequest) (*ListGeneratedPdf, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllListGenerated not implemented")
}
func (UnimplementedGeneratorServiceServer) GetListGenerated(context.Context, *GeneratedRequest) (*ListGeneratedPdf, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListGenerated not implemented")
}
func (UnimplementedGeneratorServiceServer) GetGeneratedPDF(context.Context, *GeneratedPDFRequest) (*GeneratedPdf, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGeneratedPDF not implemented")
}
func (UnimplementedGeneratorServiceServer) DeleteGenerated(context.Context, *GeneratedPDFRequest) (*DeleteGeneratedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGenerated not implemented")
}
func (UnimplementedGeneratorServiceServer) GetPDFLink(context.Context, *GeneratedPDFRequest) (*PDFLink, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPDFLink not implemented")
}
func (UnimplementedGeneratorServiceServer) mustEmbedUnimplementedGeneratorServiceServer() {}
func (UnimplementedGeneratorServiceServer) testEmbeddedByValue()                          {}

// UnsafeGeneratorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeneratorServiceServer will
// result in compilation errors.
type UnsafeGeneratorServiceServer interface {
	mustEmbedUnimplementedGeneratorServiceServer()
}

func RegisterGeneratorServiceServer(s grpc.ServiceRegistrar, srv GeneratorServiceServer) {
	// If the following call pancis, it indicates UnimplementedGeneratorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GeneratorService_ServiceDesc, srv)
}

func _GeneratorService_GetAllListGenerated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllListGeneratedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneratorServiceServer).GetAllListGenerated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GeneratorService_GetAllListGenerated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneratorServiceServer).GetAllListGenerated(ctx, req.(*AllListGeneratedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeneratorService_GetListGenerated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeneratedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneratorServiceServer).GetListGenerated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GeneratorService_GetListGenerated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneratorServiceServer).GetListGenerated(ctx, req.(*GeneratedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeneratorService_GetGeneratedPDF_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeneratedPDFRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneratorServiceServer).GetGeneratedPDF(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GeneratorService_GetGeneratedPDF_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneratorServiceServer).GetGeneratedPDF(ctx, req.(*GeneratedPDFRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeneratorService_DeleteGenerated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeneratedPDFRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneratorServiceServer).DeleteGenerated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GeneratorService_DeleteGenerated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneratorServiceServer).DeleteGenerated(ctx, req.(*GeneratedPDFRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeneratorService_GetPDFLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeneratedPDFRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneratorServiceServer).GetPDFLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GeneratorService_GetPDFLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneratorServiceServer).GetPDFLink(ctx, req.(*GeneratedPDFRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GeneratorService_ServiceDesc is the grpc.ServiceDesc for GeneratorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GeneratorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "generator.GeneratorService",
	HandlerType: (*GeneratorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllListGenerated",
			Handler:    _GeneratorService_GetAllListGenerated_Handler,
		},
		{
			MethodName: "GetListGenerated",
			Handler:    _GeneratorService_GetListGenerated_Handler,
		},
		{
			MethodName: "GetGeneratedPDF",
			Handler:    _GeneratorService_GetGeneratedPDF_Handler,
		},
		{
			MethodName: "DeleteGenerated",
			Handler:    _GeneratorService_DeleteGenerated_Handler,
		},
		{
			MethodName: "GetPDFLink",
			Handler:    _GeneratorService_GetPDFLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/generated_pdf.proto",
}
