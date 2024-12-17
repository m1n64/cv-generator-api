// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.24.4
// source: proto/skills.proto

package skills

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

type GetSkillsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CvId string `protobuf:"bytes,1,opt,name=cv_id,json=cvId,proto3" json:"cv_id,omitempty"`
}

func (x *GetSkillsRequest) Reset() {
	*x = GetSkillsRequest{}
	mi := &file_proto_skills_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSkillsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSkillsRequest) ProtoMessage() {}

func (x *GetSkillsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_skills_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSkillsRequest.ProtoReflect.Descriptor instead.
func (*GetSkillsRequest) Descriptor() ([]byte, []int) {
	return file_proto_skills_proto_rawDescGZIP(), []int{0}
}

func (x *GetSkillsRequest) GetCvId() string {
	if x != nil {
		return x.CvId
	}
	return ""
}

type GetSkillByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CvId string `protobuf:"bytes,2,opt,name=cv_id,json=cvId,proto3" json:"cv_id,omitempty"`
}

func (x *GetSkillByIDRequest) Reset() {
	*x = GetSkillByIDRequest{}
	mi := &file_proto_skills_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSkillByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSkillByIDRequest) ProtoMessage() {}

func (x *GetSkillByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_skills_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSkillByIDRequest.ProtoReflect.Descriptor instead.
func (*GetSkillByIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_skills_proto_rawDescGZIP(), []int{1}
}

func (x *GetSkillByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetSkillByIDRequest) GetCvId() string {
	if x != nil {
		return x.CvId
	}
	return ""
}

type AllSkillsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Skills []*SkillResponse `protobuf:"bytes,1,rep,name=skills,proto3" json:"skills,omitempty"`
}

func (x *AllSkillsResponse) Reset() {
	*x = AllSkillsResponse{}
	mi := &file_proto_skills_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AllSkillsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllSkillsResponse) ProtoMessage() {}

func (x *AllSkillsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_skills_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllSkillsResponse.ProtoReflect.Descriptor instead.
func (*AllSkillsResponse) Descriptor() ([]byte, []int) {
	return file_proto_skills_proto_rawDescGZIP(), []int{2}
}

func (x *AllSkillsResponse) GetSkills() []*SkillResponse {
	if x != nil {
		return x.Skills
	}
	return nil
}

type CreateSkillRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CvId string `protobuf:"bytes,1,opt,name=cv_id,json=cvId,proto3" json:"cv_id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateSkillRequest) Reset() {
	*x = CreateSkillRequest{}
	mi := &file_proto_skills_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSkillRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSkillRequest) ProtoMessage() {}

func (x *CreateSkillRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_skills_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSkillRequest.ProtoReflect.Descriptor instead.
func (*CreateSkillRequest) Descriptor() ([]byte, []int) {
	return file_proto_skills_proto_rawDescGZIP(), []int{3}
}

func (x *CreateSkillRequest) GetCvId() string {
	if x != nil {
		return x.CvId
	}
	return ""
}

func (x *CreateSkillRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateSkillByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CvId string `protobuf:"bytes,2,opt,name=cv_id,json=cvId,proto3" json:"cv_id,omitempty"`
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateSkillByIDRequest) Reset() {
	*x = UpdateSkillByIDRequest{}
	mi := &file_proto_skills_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateSkillByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSkillByIDRequest) ProtoMessage() {}

func (x *UpdateSkillByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_skills_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSkillByIDRequest.ProtoReflect.Descriptor instead.
func (*UpdateSkillByIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_skills_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateSkillByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateSkillByIDRequest) GetCvId() string {
	if x != nil {
		return x.CvId
	}
	return ""
}

func (x *UpdateSkillByIDRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteSkillByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CvId string `protobuf:"bytes,2,opt,name=cv_id,json=cvId,proto3" json:"cv_id,omitempty"`
}

func (x *DeleteSkillByIDRequest) Reset() {
	*x = DeleteSkillByIDRequest{}
	mi := &file_proto_skills_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteSkillByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSkillByIDRequest) ProtoMessage() {}

func (x *DeleteSkillByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_skills_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSkillByIDRequest.ProtoReflect.Descriptor instead.
func (*DeleteSkillByIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_skills_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteSkillByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteSkillByIDRequest) GetCvId() string {
	if x != nil {
		return x.CvId
	}
	return ""
}

type SkillResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CvId      string `protobuf:"bytes,2,opt,name=cv_id,json=cvId,proto3" json:"cv_id,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *SkillResponse) Reset() {
	*x = SkillResponse{}
	mi := &file_proto_skills_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SkillResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkillResponse) ProtoMessage() {}

func (x *SkillResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_skills_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkillResponse.ProtoReflect.Descriptor instead.
func (*SkillResponse) Descriptor() ([]byte, []int) {
	return file_proto_skills_proto_rawDescGZIP(), []int{6}
}

func (x *SkillResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SkillResponse) GetCvId() string {
	if x != nil {
		return x.CvId
	}
	return ""
}

func (x *SkillResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SkillResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *SkillResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type DeleteSkillByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteSkillByIDResponse) Reset() {
	*x = DeleteSkillByIDResponse{}
	mi := &file_proto_skills_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteSkillByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSkillByIDResponse) ProtoMessage() {}

func (x *DeleteSkillByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_skills_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSkillByIDResponse.ProtoReflect.Descriptor instead.
func (*DeleteSkillByIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_skills_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteSkillByIDResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_skills_proto protoreflect.FileDescriptor

var file_proto_skills_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x22, 0x27, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x13, 0x0a, 0x05, 0x63, 0x76, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x76, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x6b, 0x69, 0x6c,
	0x6c, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x13, 0x0a, 0x05,
	0x63, 0x76, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x76, 0x49,
	0x64, 0x22, 0x42, 0x0a, 0x11, 0x41, 0x6c, 0x6c, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e,
	0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x22, 0x3d, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x05, 0x63,
	0x76, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x76, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x51, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6b,
	0x69, 0x6c, 0x6c, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x13,
	0x0a, 0x05, 0x63, 0x76, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x76, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3d, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x13, 0x0a, 0x05, 0x63, 0x76, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x76, 0x49, 0x64, 0x22, 0x86, 0x01, 0x0a, 0x0d, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x63, 0x76, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x76, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x33, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x32, 0xf5, 0x02, 0x0a, 0x0d, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x73, 0x12, 0x18, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x41, 0x6c, 0x6c, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1b, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x1a, 0x2e, 0x73, 0x6b,
	0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73,
	0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52,
	0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x42, 0x79, 0x49,
	0x44, 0x12, 0x1e, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x48, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6b, 0x69, 0x6c,
	0x6c, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1e, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1f, 0x5a, 0x1d,
	0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x6b, 0x69, 0x6c, 0x6c,
	0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_skills_proto_rawDescOnce sync.Once
	file_proto_skills_proto_rawDescData = file_proto_skills_proto_rawDesc
)

func file_proto_skills_proto_rawDescGZIP() []byte {
	file_proto_skills_proto_rawDescOnce.Do(func() {
		file_proto_skills_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_skills_proto_rawDescData)
	})
	return file_proto_skills_proto_rawDescData
}

var file_proto_skills_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_skills_proto_goTypes = []any{
	(*GetSkillsRequest)(nil),        // 0: skills.GetSkillsRequest
	(*GetSkillByIDRequest)(nil),     // 1: skills.GetSkillByIDRequest
	(*AllSkillsResponse)(nil),       // 2: skills.AllSkillsResponse
	(*CreateSkillRequest)(nil),      // 3: skills.CreateSkillRequest
	(*UpdateSkillByIDRequest)(nil),  // 4: skills.UpdateSkillByIDRequest
	(*DeleteSkillByIDRequest)(nil),  // 5: skills.DeleteSkillByIDRequest
	(*SkillResponse)(nil),           // 6: skills.SkillResponse
	(*DeleteSkillByIDResponse)(nil), // 7: skills.DeleteSkillByIDResponse
}
var file_proto_skills_proto_depIdxs = []int32{
	6, // 0: skills.AllSkillsResponse.skills:type_name -> skills.SkillResponse
	0, // 1: skills.SkillsService.GetSkills:input_type -> skills.GetSkillsRequest
	1, // 2: skills.SkillsService.GetSkillByID:input_type -> skills.GetSkillByIDRequest
	3, // 3: skills.SkillsService.CreateSkill:input_type -> skills.CreateSkillRequest
	5, // 4: skills.SkillsService.DeleteSkillByID:input_type -> skills.DeleteSkillByIDRequest
	4, // 5: skills.SkillsService.UpdateSkillByID:input_type -> skills.UpdateSkillByIDRequest
	2, // 6: skills.SkillsService.GetSkills:output_type -> skills.AllSkillsResponse
	6, // 7: skills.SkillsService.GetSkillByID:output_type -> skills.SkillResponse
	6, // 8: skills.SkillsService.CreateSkill:output_type -> skills.SkillResponse
	7, // 9: skills.SkillsService.DeleteSkillByID:output_type -> skills.DeleteSkillByIDResponse
	6, // 10: skills.SkillsService.UpdateSkillByID:output_type -> skills.SkillResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_skills_proto_init() }
func file_proto_skills_proto_init() {
	if File_proto_skills_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_skills_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_skills_proto_goTypes,
		DependencyIndexes: file_proto_skills_proto_depIdxs,
		MessageInfos:      file_proto_skills_proto_msgTypes,
	}.Build()
	File_proto_skills_proto = out.File
	file_proto_skills_proto_rawDesc = nil
	file_proto_skills_proto_goTypes = nil
	file_proto_skills_proto_depIdxs = nil
}
