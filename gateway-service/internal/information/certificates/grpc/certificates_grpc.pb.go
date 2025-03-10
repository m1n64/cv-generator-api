// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.4
// source: proto/certificates.proto

package certificates

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
	CertificatesService_GetCertificates_FullMethodName       = "/certificates.CertificatesService/GetCertificates"
	CertificatesService_GetCertificateByID_FullMethodName    = "/certificates.CertificatesService/GetCertificateByID"
	CertificatesService_CreateCertificate_FullMethodName     = "/certificates.CertificatesService/CreateCertificate"
	CertificatesService_DeleteCertificateByID_FullMethodName = "/certificates.CertificatesService/DeleteCertificateByID"
	CertificatesService_UpdateCertificateByID_FullMethodName = "/certificates.CertificatesService/UpdateCertificateByID"
)

// CertificatesServiceClient is the client API for CertificatesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CertificatesServiceClient interface {
	GetCertificates(ctx context.Context, in *GetCertificatesRequest, opts ...grpc.CallOption) (*AllCertificatesResponse, error)
	GetCertificateByID(ctx context.Context, in *GetCertificateByIDRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
	CreateCertificate(ctx context.Context, in *CreateCertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
	DeleteCertificateByID(ctx context.Context, in *DeleteCertificateByIDRequest, opts ...grpc.CallOption) (*DeleteCertificateByIDResponse, error)
	UpdateCertificateByID(ctx context.Context, in *UpdateCertificateByIDRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
}

type certificatesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCertificatesServiceClient(cc grpc.ClientConnInterface) CertificatesServiceClient {
	return &certificatesServiceClient{cc}
}

func (c *certificatesServiceClient) GetCertificates(ctx context.Context, in *GetCertificatesRequest, opts ...grpc.CallOption) (*AllCertificatesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AllCertificatesResponse)
	err := c.cc.Invoke(ctx, CertificatesService_GetCertificates_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesServiceClient) GetCertificateByID(ctx context.Context, in *GetCertificateByIDRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, CertificatesService_GetCertificateByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesServiceClient) CreateCertificate(ctx context.Context, in *CreateCertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, CertificatesService_CreateCertificate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesServiceClient) DeleteCertificateByID(ctx context.Context, in *DeleteCertificateByIDRequest, opts ...grpc.CallOption) (*DeleteCertificateByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCertificateByIDResponse)
	err := c.cc.Invoke(ctx, CertificatesService_DeleteCertificateByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesServiceClient) UpdateCertificateByID(ctx context.Context, in *UpdateCertificateByIDRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, CertificatesService_UpdateCertificateByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificatesServiceServer is the server API for CertificatesService service.
// All implementations must embed UnimplementedCertificatesServiceServer
// for forward compatibility.
type CertificatesServiceServer interface {
	GetCertificates(context.Context, *GetCertificatesRequest) (*AllCertificatesResponse, error)
	GetCertificateByID(context.Context, *GetCertificateByIDRequest) (*CertificateResponse, error)
	CreateCertificate(context.Context, *CreateCertificateRequest) (*CertificateResponse, error)
	DeleteCertificateByID(context.Context, *DeleteCertificateByIDRequest) (*DeleteCertificateByIDResponse, error)
	UpdateCertificateByID(context.Context, *UpdateCertificateByIDRequest) (*CertificateResponse, error)
	mustEmbedUnimplementedCertificatesServiceServer()
}

// UnimplementedCertificatesServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCertificatesServiceServer struct{}

func (UnimplementedCertificatesServiceServer) GetCertificates(context.Context, *GetCertificatesRequest) (*AllCertificatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCertificates not implemented")
}
func (UnimplementedCertificatesServiceServer) GetCertificateByID(context.Context, *GetCertificateByIDRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCertificateByID not implemented")
}
func (UnimplementedCertificatesServiceServer) CreateCertificate(context.Context, *CreateCertificateRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCertificate not implemented")
}
func (UnimplementedCertificatesServiceServer) DeleteCertificateByID(context.Context, *DeleteCertificateByIDRequest) (*DeleteCertificateByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCertificateByID not implemented")
}
func (UnimplementedCertificatesServiceServer) UpdateCertificateByID(context.Context, *UpdateCertificateByIDRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCertificateByID not implemented")
}
func (UnimplementedCertificatesServiceServer) mustEmbedUnimplementedCertificatesServiceServer() {}
func (UnimplementedCertificatesServiceServer) testEmbeddedByValue()                             {}

// UnsafeCertificatesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CertificatesServiceServer will
// result in compilation errors.
type UnsafeCertificatesServiceServer interface {
	mustEmbedUnimplementedCertificatesServiceServer()
}

func RegisterCertificatesServiceServer(s grpc.ServiceRegistrar, srv CertificatesServiceServer) {
	// If the following call pancis, it indicates UnimplementedCertificatesServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CertificatesService_ServiceDesc, srv)
}

func _CertificatesService_GetCertificates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCertificatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServiceServer).GetCertificates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificatesService_GetCertificates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServiceServer).GetCertificates(ctx, req.(*GetCertificatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificatesService_GetCertificateByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCertificateByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServiceServer).GetCertificateByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificatesService_GetCertificateByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServiceServer).GetCertificateByID(ctx, req.(*GetCertificateByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificatesService_CreateCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServiceServer).CreateCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificatesService_CreateCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServiceServer).CreateCertificate(ctx, req.(*CreateCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificatesService_DeleteCertificateByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCertificateByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServiceServer).DeleteCertificateByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificatesService_DeleteCertificateByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServiceServer).DeleteCertificateByID(ctx, req.(*DeleteCertificateByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificatesService_UpdateCertificateByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCertificateByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServiceServer).UpdateCertificateByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificatesService_UpdateCertificateByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServiceServer).UpdateCertificateByID(ctx, req.(*UpdateCertificateByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CertificatesService_ServiceDesc is the grpc.ServiceDesc for CertificatesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CertificatesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "certificates.CertificatesService",
	HandlerType: (*CertificatesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCertificates",
			Handler:    _CertificatesService_GetCertificates_Handler,
		},
		{
			MethodName: "GetCertificateByID",
			Handler:    _CertificatesService_GetCertificateByID_Handler,
		},
		{
			MethodName: "CreateCertificate",
			Handler:    _CertificatesService_CreateCertificate_Handler,
		},
		{
			MethodName: "DeleteCertificateByID",
			Handler:    _CertificatesService_DeleteCertificateByID_Handler,
		},
		{
			MethodName: "UpdateCertificateByID",
			Handler:    _CertificatesService_UpdateCertificateByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/certificates.proto",
}
