// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.4
// source: cv.proto

package cv

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
	CVService_CreateCV_FullMethodName          = "/cv.CVService/CreateCV"
	CVService_GetAllCVsByUserID_FullMethodName = "/cv.CVService/GetAllCVsByUserID"
	CVService_GetCVByID_FullMethodName         = "/cv.CVService/GetCVByID"
	CVService_DeleteCVByID_FullMethodName      = "/cv.CVService/DeleteCVByID"
	CVService_UpdateCV_FullMethodName          = "/cv.CVService/UpdateCV"
	CVService_GetOriginalID_FullMethodName     = "/cv.CVService/GetOriginalID"
)

// CVServiceClient is the client API for CVService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// CVService содержит методы для работы с резюме
type CVServiceClient interface {
	// Метод для создания резюме
	CreateCV(ctx context.Context, in *CreateCVRequest, opts ...grpc.CallOption) (*CVResponse, error)
	// Метод для получения всех резюме пользователя
	GetAllCVsByUserID(ctx context.Context, in *GetAllCVsByUserIDRequest, opts ...grpc.CallOption) (*GetAllCVsResponse, error)
	// Метод для получения резюме по ID
	GetCVByID(ctx context.Context, in *GetCVByIDRequest, opts ...grpc.CallOption) (*CVResponse, error)
	// Метод для удаления резюме
	DeleteCVByID(ctx context.Context, in *DeleteCVByIDRequest, opts ...grpc.CallOption) (*DeleteCVByIDResponse, error)
	// Метод для редактирования резюме
	UpdateCV(ctx context.Context, in *UpdateCVRequest, opts ...grpc.CallOption) (*CVResponse, error)
	// Метод для получения оригинального ID для CV
	GetOriginalID(ctx context.Context, in *GetOriginalIDRequest, opts ...grpc.CallOption) (*GetOriginalIDResponse, error)
}

type cVServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCVServiceClient(cc grpc.ClientConnInterface) CVServiceClient {
	return &cVServiceClient{cc}
}

func (c *cVServiceClient) CreateCV(ctx context.Context, in *CreateCVRequest, opts ...grpc.CallOption) (*CVResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CVResponse)
	err := c.cc.Invoke(ctx, CVService_CreateCV_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cVServiceClient) GetAllCVsByUserID(ctx context.Context, in *GetAllCVsByUserIDRequest, opts ...grpc.CallOption) (*GetAllCVsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllCVsResponse)
	err := c.cc.Invoke(ctx, CVService_GetAllCVsByUserID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cVServiceClient) GetCVByID(ctx context.Context, in *GetCVByIDRequest, opts ...grpc.CallOption) (*CVResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CVResponse)
	err := c.cc.Invoke(ctx, CVService_GetCVByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cVServiceClient) DeleteCVByID(ctx context.Context, in *DeleteCVByIDRequest, opts ...grpc.CallOption) (*DeleteCVByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCVByIDResponse)
	err := c.cc.Invoke(ctx, CVService_DeleteCVByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cVServiceClient) UpdateCV(ctx context.Context, in *UpdateCVRequest, opts ...grpc.CallOption) (*CVResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CVResponse)
	err := c.cc.Invoke(ctx, CVService_UpdateCV_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cVServiceClient) GetOriginalID(ctx context.Context, in *GetOriginalIDRequest, opts ...grpc.CallOption) (*GetOriginalIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOriginalIDResponse)
	err := c.cc.Invoke(ctx, CVService_GetOriginalID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CVServiceServer is the server API for CVService service.
// All implementations must embed UnimplementedCVServiceServer
// for forward compatibility.
//
// CVService содержит методы для работы с резюме
type CVServiceServer interface {
	// Метод для создания резюме
	CreateCV(context.Context, *CreateCVRequest) (*CVResponse, error)
	// Метод для получения всех резюме пользователя
	GetAllCVsByUserID(context.Context, *GetAllCVsByUserIDRequest) (*GetAllCVsResponse, error)
	// Метод для получения резюме по ID
	GetCVByID(context.Context, *GetCVByIDRequest) (*CVResponse, error)
	// Метод для удаления резюме
	DeleteCVByID(context.Context, *DeleteCVByIDRequest) (*DeleteCVByIDResponse, error)
	// Метод для редактирования резюме
	UpdateCV(context.Context, *UpdateCVRequest) (*CVResponse, error)
	// Метод для получения оригинального ID для CV
	GetOriginalID(context.Context, *GetOriginalIDRequest) (*GetOriginalIDResponse, error)
	mustEmbedUnimplementedCVServiceServer()
}

// UnimplementedCVServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCVServiceServer struct{}

func (UnimplementedCVServiceServer) CreateCV(context.Context, *CreateCVRequest) (*CVResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCV not implemented")
}
func (UnimplementedCVServiceServer) GetAllCVsByUserID(context.Context, *GetAllCVsByUserIDRequest) (*GetAllCVsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCVsByUserID not implemented")
}
func (UnimplementedCVServiceServer) GetCVByID(context.Context, *GetCVByIDRequest) (*CVResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCVByID not implemented")
}
func (UnimplementedCVServiceServer) DeleteCVByID(context.Context, *DeleteCVByIDRequest) (*DeleteCVByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCVByID not implemented")
}
func (UnimplementedCVServiceServer) UpdateCV(context.Context, *UpdateCVRequest) (*CVResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCV not implemented")
}
func (UnimplementedCVServiceServer) GetOriginalID(context.Context, *GetOriginalIDRequest) (*GetOriginalIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOriginalID not implemented")
}
func (UnimplementedCVServiceServer) mustEmbedUnimplementedCVServiceServer() {}
func (UnimplementedCVServiceServer) testEmbeddedByValue()                   {}

// UnsafeCVServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CVServiceServer will
// result in compilation errors.
type UnsafeCVServiceServer interface {
	mustEmbedUnimplementedCVServiceServer()
}

func RegisterCVServiceServer(s grpc.ServiceRegistrar, srv CVServiceServer) {
	// If the following call pancis, it indicates UnimplementedCVServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CVService_ServiceDesc, srv)
}

func _CVService_CreateCV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCVRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CVServiceServer).CreateCV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CVService_CreateCV_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CVServiceServer).CreateCV(ctx, req.(*CreateCVRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CVService_GetAllCVsByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllCVsByUserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CVServiceServer).GetAllCVsByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CVService_GetAllCVsByUserID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CVServiceServer).GetAllCVsByUserID(ctx, req.(*GetAllCVsByUserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CVService_GetCVByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCVByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CVServiceServer).GetCVByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CVService_GetCVByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CVServiceServer).GetCVByID(ctx, req.(*GetCVByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CVService_DeleteCVByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCVByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CVServiceServer).DeleteCVByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CVService_DeleteCVByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CVServiceServer).DeleteCVByID(ctx, req.(*DeleteCVByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CVService_UpdateCV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCVRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CVServiceServer).UpdateCV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CVService_UpdateCV_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CVServiceServer).UpdateCV(ctx, req.(*UpdateCVRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CVService_GetOriginalID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOriginalIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CVServiceServer).GetOriginalID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CVService_GetOriginalID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CVServiceServer).GetOriginalID(ctx, req.(*GetOriginalIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CVService_ServiceDesc is the grpc.ServiceDesc for CVService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CVService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cv.CVService",
	HandlerType: (*CVServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCV",
			Handler:    _CVService_CreateCV_Handler,
		},
		{
			MethodName: "GetAllCVsByUserID",
			Handler:    _CVService_GetAllCVsByUserID_Handler,
		},
		{
			MethodName: "GetCVByID",
			Handler:    _CVService_GetCVByID_Handler,
		},
		{
			MethodName: "DeleteCVByID",
			Handler:    _CVService_DeleteCVByID_Handler,
		},
		{
			MethodName: "UpdateCV",
			Handler:    _CVService_UpdateCV_Handler,
		},
		{
			MethodName: "GetOriginalID",
			Handler:    _CVService_GetOriginalID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cv.proto",
}
