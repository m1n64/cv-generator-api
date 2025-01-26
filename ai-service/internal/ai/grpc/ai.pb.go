// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v4.24.4
// source: proto/ai.proto

package ai

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetServicesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetServicesRequest) Reset() {
	*x = GetServicesRequest{}
	mi := &file_proto_ai_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetServicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServicesRequest) ProtoMessage() {}

func (x *GetServicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ai_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServicesRequest.ProtoReflect.Descriptor instead.
func (*GetServicesRequest) Descriptor() ([]byte, []int) {
	return file_proto_ai_proto_rawDescGZIP(), []int{0}
}

type GetServicesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Services      []*Services            `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetServicesResponse) Reset() {
	*x = GetServicesResponse{}
	mi := &file_proto_ai_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetServicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServicesResponse) ProtoMessage() {}

func (x *GetServicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ai_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServicesResponse.ProtoReflect.Descriptor instead.
func (*GetServicesResponse) Descriptor() ([]byte, []int) {
	return file_proto_ai_proto_rawDescGZIP(), []int{1}
}

func (x *GetServicesResponse) GetServices() []*Services {
	if x != nil {
		return x.Services
	}
	return nil
}

type Services struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ServiceId     string                 `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	ServiceName   string                 `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Services) Reset() {
	*x = Services{}
	mi := &file_proto_ai_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Services) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Services) ProtoMessage() {}

func (x *Services) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ai_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Services.ProtoReflect.Descriptor instead.
func (*Services) Descriptor() ([]byte, []int) {
	return file_proto_ai_proto_rawDescGZIP(), []int{2}
}

func (x *Services) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *Services) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

type GenerateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Prompt        string                 `protobuf:"bytes,1,opt,name=prompt,proto3" json:"prompt,omitempty"`
	ServiceId     string                 `protobuf:"bytes,2,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateRequest) Reset() {
	*x = GenerateRequest{}
	mi := &file_proto_ai_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateRequest) ProtoMessage() {}

func (x *GenerateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ai_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateRequest.ProtoReflect.Descriptor instead.
func (*GenerateRequest) Descriptor() ([]byte, []int) {
	return file_proto_ai_proto_rawDescGZIP(), []int{3}
}

func (x *GenerateRequest) GetPrompt() string {
	if x != nil {
		return x.Prompt
	}
	return ""
}

func (x *GenerateRequest) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

type GenerateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Response      string                 `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateResponse) Reset() {
	*x = GenerateResponse{}
	mi := &file_proto_ai_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateResponse) ProtoMessage() {}

func (x *GenerateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ai_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateResponse.ProtoReflect.Descriptor instead.
func (*GenerateResponse) Descriptor() ([]byte, []int) {
	return file_proto_ai_proto_rawDescGZIP(), []int{4}
}

func (x *GenerateResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_proto_ai_proto protoreflect.FileDescriptor

var file_proto_ai_proto_rawDesc = string([]byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x61, 0x69, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x28, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x4c, 0x0a, 0x08, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x0f, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72,
	0x6f, 0x6d, 0x70, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0xc1, 0x01, 0x0a, 0x09, 0x41, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x35, 0x0a, 0x08, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e,
	0x61, 0x69, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x69, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x61, 0x69, 0x2e,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x61, 0x69, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x3e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x61, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x61, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x61, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_ai_proto_rawDescOnce sync.Once
	file_proto_ai_proto_rawDescData []byte
)

func file_proto_ai_proto_rawDescGZIP() []byte {
	file_proto_ai_proto_rawDescOnce.Do(func() {
		file_proto_ai_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_ai_proto_rawDesc), len(file_proto_ai_proto_rawDesc)))
	})
	return file_proto_ai_proto_rawDescData
}

var file_proto_ai_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_ai_proto_goTypes = []any{
	(*GetServicesRequest)(nil),  // 0: ai.GetServicesRequest
	(*GetServicesResponse)(nil), // 1: ai.GetServicesResponse
	(*Services)(nil),            // 2: ai.Services
	(*GenerateRequest)(nil),     // 3: ai.GenerateRequest
	(*GenerateResponse)(nil),    // 4: ai.GenerateResponse
}
var file_proto_ai_proto_depIdxs = []int32{
	2, // 0: ai.GetServicesResponse.services:type_name -> ai.Services
	3, // 1: ai.AiService.Generate:input_type -> ai.GenerateRequest
	3, // 2: ai.AiService.StreamGenerate:input_type -> ai.GenerateRequest
	0, // 3: ai.AiService.GetServices:input_type -> ai.GetServicesRequest
	4, // 4: ai.AiService.Generate:output_type -> ai.GenerateResponse
	4, // 5: ai.AiService.StreamGenerate:output_type -> ai.GenerateResponse
	1, // 6: ai.AiService.GetServices:output_type -> ai.GetServicesResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_ai_proto_init() }
func file_proto_ai_proto_init() {
	if File_proto_ai_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_ai_proto_rawDesc), len(file_proto_ai_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ai_proto_goTypes,
		DependencyIndexes: file_proto_ai_proto_depIdxs,
		MessageInfos:      file_proto_ai_proto_msgTypes,
	}.Build()
	File_proto_ai_proto = out.File
	file_proto_ai_proto_goTypes = nil
	file_proto_ai_proto_depIdxs = nil
}
