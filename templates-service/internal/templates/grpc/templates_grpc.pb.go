// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.4
// source: proto/templates.proto

package templates

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
	TemplateService_GetDefaultTemplate_FullMethodName   = "/templates.TemplateService/GetDefaultTemplate"
	TemplateService_GetColorScheme_FullMethodName       = "/templates.TemplateService/GetColorScheme"
	TemplateService_GetColorSchemeByName_FullMethodName = "/templates.TemplateService/GetColorSchemeByName"
)

// TemplateServiceClient is the client API for TemplateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TemplateServiceClient interface {
	GetDefaultTemplate(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Template, error)
	GetColorScheme(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ColorScheme, error)
	GetColorSchemeByName(ctx context.Context, in *ColorSchemeByNameRequest, opts ...grpc.CallOption) (*Color, error)
}

type templateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTemplateServiceClient(cc grpc.ClientConnInterface) TemplateServiceClient {
	return &templateServiceClient{cc}
}

func (c *templateServiceClient) GetDefaultTemplate(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Template, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Template)
	err := c.cc.Invoke(ctx, TemplateService_GetDefaultTemplate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateServiceClient) GetColorScheme(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ColorScheme, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ColorScheme)
	err := c.cc.Invoke(ctx, TemplateService_GetColorScheme_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateServiceClient) GetColorSchemeByName(ctx context.Context, in *ColorSchemeByNameRequest, opts ...grpc.CallOption) (*Color, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Color)
	err := c.cc.Invoke(ctx, TemplateService_GetColorSchemeByName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TemplateServiceServer is the server API for TemplateService service.
// All implementations must embed UnimplementedTemplateServiceServer
// for forward compatibility.
type TemplateServiceServer interface {
	GetDefaultTemplate(context.Context, *Empty) (*Template, error)
	GetColorScheme(context.Context, *Empty) (*ColorScheme, error)
	GetColorSchemeByName(context.Context, *ColorSchemeByNameRequest) (*Color, error)
	mustEmbedUnimplementedTemplateServiceServer()
}

// UnimplementedTemplateServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTemplateServiceServer struct{}

func (UnimplementedTemplateServiceServer) GetDefaultTemplate(context.Context, *Empty) (*Template, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefaultTemplate not implemented")
}
func (UnimplementedTemplateServiceServer) GetColorScheme(context.Context, *Empty) (*ColorScheme, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetColorScheme not implemented")
}
func (UnimplementedTemplateServiceServer) GetColorSchemeByName(context.Context, *ColorSchemeByNameRequest) (*Color, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetColorSchemeByName not implemented")
}
func (UnimplementedTemplateServiceServer) mustEmbedUnimplementedTemplateServiceServer() {}
func (UnimplementedTemplateServiceServer) testEmbeddedByValue()                         {}

// UnsafeTemplateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TemplateServiceServer will
// result in compilation errors.
type UnsafeTemplateServiceServer interface {
	mustEmbedUnimplementedTemplateServiceServer()
}

func RegisterTemplateServiceServer(s grpc.ServiceRegistrar, srv TemplateServiceServer) {
	// If the following call pancis, it indicates UnimplementedTemplateServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TemplateService_ServiceDesc, srv)
}

func _TemplateService_GetDefaultTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServiceServer).GetDefaultTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TemplateService_GetDefaultTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServiceServer).GetDefaultTemplate(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TemplateService_GetColorScheme_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServiceServer).GetColorScheme(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TemplateService_GetColorScheme_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServiceServer).GetColorScheme(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TemplateService_GetColorSchemeByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ColorSchemeByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServiceServer).GetColorSchemeByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TemplateService_GetColorSchemeByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServiceServer).GetColorSchemeByName(ctx, req.(*ColorSchemeByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TemplateService_ServiceDesc is the grpc.ServiceDesc for TemplateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TemplateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "templates.TemplateService",
	HandlerType: (*TemplateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDefaultTemplate",
			Handler:    _TemplateService_GetDefaultTemplate_Handler,
		},
		{
			MethodName: "GetColorScheme",
			Handler:    _TemplateService_GetColorScheme_Handler,
		},
		{
			MethodName: "GetColorSchemeByName",
			Handler:    _TemplateService_GetColorSchemeByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/templates.proto",
}
