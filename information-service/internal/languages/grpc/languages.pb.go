// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.24.4
// source: proto/languages.proto

package languages

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetLanguagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CvId string `protobuf:"bytes,1,opt,name=cv_id,json=cvId,proto3" json:"cv_id,omitempty"`
}

func (x *GetLanguagesRequest) Reset() {
	*x = GetLanguagesRequest{}
	mi := &file_proto_languages_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLanguagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLanguagesRequest) ProtoMessage() {}

func (x *GetLanguagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_languages_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLanguagesRequest.ProtoReflect.Descriptor instead.
func (*GetLanguagesRequest) Descriptor() ([]byte, []int) {
	return file_proto_languages_proto_rawDescGZIP(), []int{0}
}

func (x *GetLanguagesRequest) GetCvId() string {
	if x != nil {
		return x.CvId
	}
	return ""
}

type GetLanguageByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CvId string `protobuf:"bytes,2,opt,name=cv_id,json=cvId,proto3" json:"cv_id,omitempty"`
}

func (x *GetLanguageByIDRequest) Reset() {
	*x = GetLanguageByIDRequest{}
	mi := &file_proto_languages_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLanguageByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLanguageByIDRequest) ProtoMessage() {}

func (x *GetLanguageByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_languages_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLanguageByIDRequest.ProtoReflect.Descriptor instead.
func (*GetLanguageByIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_languages_proto_rawDescGZIP(), []int{1}
}

func (x *GetLanguageByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetLanguageByIDRequest) GetCvId() string {
	if x != nil {
		return x.CvId
	}
	return ""
}

type AllLanguagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Languages []*LanguageResponse `protobuf:"bytes,1,rep,name=languages,proto3" json:"languages,omitempty"`
}

func (x *AllLanguagesResponse) Reset() {
	*x = AllLanguagesResponse{}
	mi := &file_proto_languages_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AllLanguagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllLanguagesResponse) ProtoMessage() {}

func (x *AllLanguagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_languages_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllLanguagesResponse.ProtoReflect.Descriptor instead.
func (*AllLanguagesResponse) Descriptor() ([]byte, []int) {
	return file_proto_languages_proto_rawDescGZIP(), []int{2}
}

func (x *AllLanguagesResponse) GetLanguages() []*LanguageResponse {
	if x != nil {
		return x.Languages
	}
	return nil
}

type CreateLanguageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CvId  string `protobuf:"bytes,1,opt,name=cv_id,json=cvId,proto3" json:"cv_id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Level string `protobuf:"bytes,3,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *CreateLanguageRequest) Reset() {
	*x = CreateLanguageRequest{}
	mi := &file_proto_languages_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateLanguageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLanguageRequest) ProtoMessage() {}

func (x *CreateLanguageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_languages_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLanguageRequest.ProtoReflect.Descriptor instead.
func (*CreateLanguageRequest) Descriptor() ([]byte, []int) {
	return file_proto_languages_proto_rawDescGZIP(), []int{3}
}

func (x *CreateLanguageRequest) GetCvId() string {
	if x != nil {
		return x.CvId
	}
	return ""
}

func (x *CreateLanguageRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateLanguageRequest) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

type UpdateLanguageByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CvId  string `protobuf:"bytes,2,opt,name=cv_id,json=cvId,proto3" json:"cv_id,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Level string `protobuf:"bytes,4,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *UpdateLanguageByIDRequest) Reset() {
	*x = UpdateLanguageByIDRequest{}
	mi := &file_proto_languages_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateLanguageByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLanguageByIDRequest) ProtoMessage() {}

func (x *UpdateLanguageByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_languages_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLanguageByIDRequest.ProtoReflect.Descriptor instead.
func (*UpdateLanguageByIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_languages_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateLanguageByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateLanguageByIDRequest) GetCvId() string {
	if x != nil {
		return x.CvId
	}
	return ""
}

func (x *UpdateLanguageByIDRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateLanguageByIDRequest) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

type DeleteLanguageByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CvId string `protobuf:"bytes,2,opt,name=cv_id,json=cvId,proto3" json:"cv_id,omitempty"`
}

func (x *DeleteLanguageByIDRequest) Reset() {
	*x = DeleteLanguageByIDRequest{}
	mi := &file_proto_languages_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteLanguageByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLanguageByIDRequest) ProtoMessage() {}

func (x *DeleteLanguageByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_languages_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLanguageByIDRequest.ProtoReflect.Descriptor instead.
func (*DeleteLanguageByIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_languages_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteLanguageByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteLanguageByIDRequest) GetCvId() string {
	if x != nil {
		return x.CvId
	}
	return ""
}

type LanguageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CvId      string `protobuf:"bytes,2,opt,name=cv_id,json=cvId,proto3" json:"cv_id,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Level     string `protobuf:"bytes,4,opt,name=level,proto3" json:"level,omitempty"`
	CreatedAt string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *LanguageResponse) Reset() {
	*x = LanguageResponse{}
	mi := &file_proto_languages_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LanguageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LanguageResponse) ProtoMessage() {}

func (x *LanguageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_languages_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LanguageResponse.ProtoReflect.Descriptor instead.
func (*LanguageResponse) Descriptor() ([]byte, []int) {
	return file_proto_languages_proto_rawDescGZIP(), []int{6}
}

func (x *LanguageResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LanguageResponse) GetCvId() string {
	if x != nil {
		return x.CvId
	}
	return ""
}

func (x *LanguageResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LanguageResponse) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *LanguageResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *LanguageResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type DeleteLanguageByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteLanguageByIDResponse) Reset() {
	*x = DeleteLanguageByIDResponse{}
	mi := &file_proto_languages_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteLanguageByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLanguageByIDResponse) ProtoMessage() {}

func (x *DeleteLanguageByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_languages_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLanguageByIDResponse.ProtoReflect.Descriptor instead.
func (*DeleteLanguageByIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_languages_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteLanguageByIDResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_languages_proto protoreflect.FileDescriptor

var file_proto_languages_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x73, 0x22, 0x2a, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x05, 0x63, 0x76, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x76, 0x49, 0x64, 0x22, 0x3d,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x63, 0x76, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x76, 0x49, 0x64, 0x22, 0x51, 0x0a,
	0x14, 0x41, 0x6c, 0x6c, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73,
	0x22, 0x56, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x05, 0x63, 0x76, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x76, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x6a, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x63, 0x76, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x76, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x22, 0x40, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x13, 0x0a, 0x05, 0x63, 0x76, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x76, 0x49, 0x64, 0x22, 0x9f, 0x01, 0x0a, 0x10, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x63,
	0x76, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x76, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x36, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x32, 0xc3, 0x03, 0x0a, 0x10, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x41, 0x6c, 0x6c, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x42, 0x79, 0x49, 0x44, 0x12, 0x21, 0x2e, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x20, 0x2e, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x42, 0x79, 0x49, 0x44,
	0x12, 0x24, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a,
	0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x42,
	0x79, 0x49, 0x44, 0x12, 0x24, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x25, 0x5a, 0x23, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x3b, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_languages_proto_rawDescOnce sync.Once
	file_proto_languages_proto_rawDescData = file_proto_languages_proto_rawDesc
)

func file_proto_languages_proto_rawDescGZIP() []byte {
	file_proto_languages_proto_rawDescOnce.Do(func() {
		file_proto_languages_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_languages_proto_rawDescData)
	})
	return file_proto_languages_proto_rawDescData
}

var file_proto_languages_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_languages_proto_goTypes = []any{
	(*GetLanguagesRequest)(nil),        // 0: languages.GetLanguagesRequest
	(*GetLanguageByIDRequest)(nil),     // 1: languages.GetLanguageByIDRequest
	(*AllLanguagesResponse)(nil),       // 2: languages.AllLanguagesResponse
	(*CreateLanguageRequest)(nil),      // 3: languages.CreateLanguageRequest
	(*UpdateLanguageByIDRequest)(nil),  // 4: languages.UpdateLanguageByIDRequest
	(*DeleteLanguageByIDRequest)(nil),  // 5: languages.DeleteLanguageByIDRequest
	(*LanguageResponse)(nil),           // 6: languages.LanguageResponse
	(*DeleteLanguageByIDResponse)(nil), // 7: languages.DeleteLanguageByIDResponse
}
var file_proto_languages_proto_depIdxs = []int32{
	6, // 0: languages.AllLanguagesResponse.languages:type_name -> languages.LanguageResponse
	0, // 1: languages.LanguagesService.GetLanguages:input_type -> languages.GetLanguagesRequest
	1, // 2: languages.LanguagesService.GetLanguageByID:input_type -> languages.GetLanguageByIDRequest
	3, // 3: languages.LanguagesService.CreateLanguage:input_type -> languages.CreateLanguageRequest
	5, // 4: languages.LanguagesService.DeleteLanguageByID:input_type -> languages.DeleteLanguageByIDRequest
	4, // 5: languages.LanguagesService.UpdateLanguageByID:input_type -> languages.UpdateLanguageByIDRequest
	2, // 6: languages.LanguagesService.GetLanguages:output_type -> languages.AllLanguagesResponse
	6, // 7: languages.LanguagesService.GetLanguageByID:output_type -> languages.LanguageResponse
	6, // 8: languages.LanguagesService.CreateLanguage:output_type -> languages.LanguageResponse
	7, // 9: languages.LanguagesService.DeleteLanguageByID:output_type -> languages.DeleteLanguageByIDResponse
	6, // 10: languages.LanguagesService.UpdateLanguageByID:output_type -> languages.LanguageResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_languages_proto_init() }
func file_proto_languages_proto_init() {
	if File_proto_languages_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_languages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_languages_proto_goTypes,
		DependencyIndexes: file_proto_languages_proto_depIdxs,
		MessageInfos:      file_proto_languages_proto_msgTypes,
	}.Build()
	File_proto_languages_proto = out.File
	file_proto_languages_proto_rawDesc = nil
	file_proto_languages_proto_goTypes = nil
	file_proto_languages_proto_depIdxs = nil
}