// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.24.4
// source: proto/educations.proto

package educations

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

type GetEducationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CvId string `protobuf:"bytes,1,opt,name=cv_id,json=cvId,proto3" json:"cv_id,omitempty"`
}

func (x *GetEducationsRequest) Reset() {
	*x = GetEducationsRequest{}
	mi := &file_proto_educations_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEducationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEducationsRequest) ProtoMessage() {}

func (x *GetEducationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_educations_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEducationsRequest.ProtoReflect.Descriptor instead.
func (*GetEducationsRequest) Descriptor() ([]byte, []int) {
	return file_proto_educations_proto_rawDescGZIP(), []int{0}
}

func (x *GetEducationsRequest) GetCvId() string {
	if x != nil {
		return x.CvId
	}
	return ""
}

type GetEducationByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CvId string `protobuf:"bytes,2,opt,name=cv_id,json=cvId,proto3" json:"cv_id,omitempty"`
}

func (x *GetEducationByIDRequest) Reset() {
	*x = GetEducationByIDRequest{}
	mi := &file_proto_educations_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEducationByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEducationByIDRequest) ProtoMessage() {}

func (x *GetEducationByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_educations_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEducationByIDRequest.ProtoReflect.Descriptor instead.
func (*GetEducationByIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_educations_proto_rawDescGZIP(), []int{1}
}

func (x *GetEducationByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetEducationByIDRequest) GetCvId() string {
	if x != nil {
		return x.CvId
	}
	return ""
}

type AllEducationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Educations []*EducationResponse `protobuf:"bytes,1,rep,name=educations,proto3" json:"educations,omitempty"`
}

func (x *AllEducationsResponse) Reset() {
	*x = AllEducationsResponse{}
	mi := &file_proto_educations_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AllEducationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllEducationsResponse) ProtoMessage() {}

func (x *AllEducationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_educations_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllEducationsResponse.ProtoReflect.Descriptor instead.
func (*AllEducationsResponse) Descriptor() ([]byte, []int) {
	return file_proto_educations_proto_rawDescGZIP(), []int{2}
}

func (x *AllEducationsResponse) GetEducations() []*EducationResponse {
	if x != nil {
		return x.Educations
	}
	return nil
}

type CreateEducationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CvId        string  `protobuf:"bytes,1,opt,name=cv_id,json=cvId,proto3" json:"cv_id,omitempty"`
	Institution string  `protobuf:"bytes,2,opt,name=institution,proto3" json:"institution,omitempty"`
	StartDate   string  `protobuf:"bytes,3,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate     *string `protobuf:"bytes,4,opt,name=end_date,json=endDate,proto3,oneof" json:"end_date,omitempty"`
	Location    string  `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
	Description *string `protobuf:"bytes,6,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Faculty     string  `protobuf:"bytes,7,opt,name=faculty,proto3" json:"faculty,omitempty"`
	Degree      *string `protobuf:"bytes,8,opt,name=degree,proto3,oneof" json:"degree,omitempty"`
}

func (x *CreateEducationRequest) Reset() {
	*x = CreateEducationRequest{}
	mi := &file_proto_educations_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateEducationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEducationRequest) ProtoMessage() {}

func (x *CreateEducationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_educations_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEducationRequest.ProtoReflect.Descriptor instead.
func (*CreateEducationRequest) Descriptor() ([]byte, []int) {
	return file_proto_educations_proto_rawDescGZIP(), []int{3}
}

func (x *CreateEducationRequest) GetCvId() string {
	if x != nil {
		return x.CvId
	}
	return ""
}

func (x *CreateEducationRequest) GetInstitution() string {
	if x != nil {
		return x.Institution
	}
	return ""
}

func (x *CreateEducationRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *CreateEducationRequest) GetEndDate() string {
	if x != nil && x.EndDate != nil {
		return *x.EndDate
	}
	return ""
}

func (x *CreateEducationRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *CreateEducationRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *CreateEducationRequest) GetFaculty() string {
	if x != nil {
		return x.Faculty
	}
	return ""
}

func (x *CreateEducationRequest) GetDegree() string {
	if x != nil && x.Degree != nil {
		return *x.Degree
	}
	return ""
}

type UpdateEducationByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CvId        string  `protobuf:"bytes,2,opt,name=cv_id,json=cvId,proto3" json:"cv_id,omitempty"`
	Institution string  `protobuf:"bytes,3,opt,name=institution,proto3" json:"institution,omitempty"`
	StartDate   string  `protobuf:"bytes,4,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate     *string `protobuf:"bytes,5,opt,name=end_date,json=endDate,proto3,oneof" json:"end_date,omitempty"`
	Location    string  `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	Description *string `protobuf:"bytes,7,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Faculty     string  `protobuf:"bytes,8,opt,name=faculty,proto3" json:"faculty,omitempty"`
	Degree      *string `protobuf:"bytes,9,opt,name=degree,proto3,oneof" json:"degree,omitempty"`
}

func (x *UpdateEducationByIDRequest) Reset() {
	*x = UpdateEducationByIDRequest{}
	mi := &file_proto_educations_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateEducationByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEducationByIDRequest) ProtoMessage() {}

func (x *UpdateEducationByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_educations_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEducationByIDRequest.ProtoReflect.Descriptor instead.
func (*UpdateEducationByIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_educations_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateEducationByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateEducationByIDRequest) GetCvId() string {
	if x != nil {
		return x.CvId
	}
	return ""
}

func (x *UpdateEducationByIDRequest) GetInstitution() string {
	if x != nil {
		return x.Institution
	}
	return ""
}

func (x *UpdateEducationByIDRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *UpdateEducationByIDRequest) GetEndDate() string {
	if x != nil && x.EndDate != nil {
		return *x.EndDate
	}
	return ""
}

func (x *UpdateEducationByIDRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *UpdateEducationByIDRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *UpdateEducationByIDRequest) GetFaculty() string {
	if x != nil {
		return x.Faculty
	}
	return ""
}

func (x *UpdateEducationByIDRequest) GetDegree() string {
	if x != nil && x.Degree != nil {
		return *x.Degree
	}
	return ""
}

type DeleteEducationByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CvId string `protobuf:"bytes,2,opt,name=cv_id,json=cvId,proto3" json:"cv_id,omitempty"`
}

func (x *DeleteEducationByIDRequest) Reset() {
	*x = DeleteEducationByIDRequest{}
	mi := &file_proto_educations_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteEducationByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEducationByIDRequest) ProtoMessage() {}

func (x *DeleteEducationByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_educations_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEducationByIDRequest.ProtoReflect.Descriptor instead.
func (*DeleteEducationByIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_educations_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteEducationByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteEducationByIDRequest) GetCvId() string {
	if x != nil {
		return x.CvId
	}
	return ""
}

type EducationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CvId        string  `protobuf:"bytes,2,opt,name=cv_id,json=cvId,proto3" json:"cv_id,omitempty"`
	Institution string  `protobuf:"bytes,3,opt,name=institution,proto3" json:"institution,omitempty"`
	StartDate   string  `protobuf:"bytes,4,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate     *string `protobuf:"bytes,5,opt,name=end_date,json=endDate,proto3,oneof" json:"end_date,omitempty"`
	Location    string  `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	Description *string `protobuf:"bytes,7,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Faculty     string  `protobuf:"bytes,8,opt,name=faculty,proto3" json:"faculty,omitempty"`
	Degree      *string `protobuf:"bytes,9,opt,name=degree,proto3,oneof" json:"degree,omitempty"`
	CreatedAt   string  `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string  `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *EducationResponse) Reset() {
	*x = EducationResponse{}
	mi := &file_proto_educations_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EducationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EducationResponse) ProtoMessage() {}

func (x *EducationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_educations_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EducationResponse.ProtoReflect.Descriptor instead.
func (*EducationResponse) Descriptor() ([]byte, []int) {
	return file_proto_educations_proto_rawDescGZIP(), []int{6}
}

func (x *EducationResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EducationResponse) GetCvId() string {
	if x != nil {
		return x.CvId
	}
	return ""
}

func (x *EducationResponse) GetInstitution() string {
	if x != nil {
		return x.Institution
	}
	return ""
}

func (x *EducationResponse) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *EducationResponse) GetEndDate() string {
	if x != nil && x.EndDate != nil {
		return *x.EndDate
	}
	return ""
}

func (x *EducationResponse) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *EducationResponse) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *EducationResponse) GetFaculty() string {
	if x != nil {
		return x.Faculty
	}
	return ""
}

func (x *EducationResponse) GetDegree() string {
	if x != nil && x.Degree != nil {
		return *x.Degree
	}
	return ""
}

func (x *EducationResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *EducationResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type DeleteEducationByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteEducationByIDResponse) Reset() {
	*x = DeleteEducationByIDResponse{}
	mi := &file_proto_educations_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteEducationByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEducationByIDResponse) ProtoMessage() {}

func (x *DeleteEducationByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_educations_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEducationByIDResponse.ProtoReflect.Descriptor instead.
func (*DeleteEducationByIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_educations_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteEducationByIDResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_educations_proto protoreflect.FileDescriptor

var file_proto_educations_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0x2b, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x45, 0x64, 0x75, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x05,
	0x63, 0x76, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x76, 0x49,
	0x64, 0x22, 0x3e, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x13, 0x0a, 0x05,
	0x63, 0x76, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x76, 0x49,
	0x64, 0x22, 0x56, 0x0a, 0x15, 0x41, 0x6c, 0x6c, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x65, 0x64,
	0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x64, 0x75, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x65,
	0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xb0, 0x02, 0x0a, 0x16, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x05, 0x63, 0x76, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x76, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x73,
	0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x08, 0x65, 0x6e,
	0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07,
	0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a,
	0x07, 0x66, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x66, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x06, 0x64, 0x65, 0x67, 0x72, 0x65,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x06, 0x64, 0x65, 0x67, 0x72, 0x65,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x22, 0xc4, 0x02, 0x0a,
	0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x63,
	0x76, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x76, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x1e, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x1b,
	0x0a, 0x06, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02,
	0x52, 0x06, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f,
	0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x64, 0x65, 0x67,
	0x72, 0x65, 0x65, 0x22, 0x41, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x64, 0x75,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x13, 0x0a, 0x05, 0x63, 0x76, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x76, 0x49, 0x64, 0x22, 0xf9, 0x02, 0x0a, 0x11, 0x45, 0x64, 0x75, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x13, 0x0a, 0x05,
	0x63, 0x76, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x76, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x1e, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x79,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12,
	0x1b, 0x0a, 0x06, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x02, 0x52, 0x06, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x65,
	0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x64, 0x65, 0x67, 0x72,
	0x65, 0x65, 0x22, 0x37, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x64, 0x75, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xdc, 0x03, 0x0a, 0x10,
	0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x54, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x20, 0x2e, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x41, 0x6c, 0x6c, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x64, 0x75,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x44, 0x12, 0x23, 0x2e, 0x65, 0x64, 0x75,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x64, 0x75, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x64, 0x75,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54,
	0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x22, 0x2e, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x64,
	0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x44, 0x12, 0x26, 0x2e, 0x65, 0x64,
	0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45,
	0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x13,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x79, 0x49, 0x44, 0x12, 0x26, 0x2e, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x65, 0x64,
	0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x27, 0x5a, 0x25, 0x2e, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_educations_proto_rawDescOnce sync.Once
	file_proto_educations_proto_rawDescData = file_proto_educations_proto_rawDesc
)

func file_proto_educations_proto_rawDescGZIP() []byte {
	file_proto_educations_proto_rawDescOnce.Do(func() {
		file_proto_educations_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_educations_proto_rawDescData)
	})
	return file_proto_educations_proto_rawDescData
}

var file_proto_educations_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_educations_proto_goTypes = []any{
	(*GetEducationsRequest)(nil),        // 0: educations.GetEducationsRequest
	(*GetEducationByIDRequest)(nil),     // 1: educations.GetEducationByIDRequest
	(*AllEducationsResponse)(nil),       // 2: educations.AllEducationsResponse
	(*CreateEducationRequest)(nil),      // 3: educations.CreateEducationRequest
	(*UpdateEducationByIDRequest)(nil),  // 4: educations.UpdateEducationByIDRequest
	(*DeleteEducationByIDRequest)(nil),  // 5: educations.DeleteEducationByIDRequest
	(*EducationResponse)(nil),           // 6: educations.EducationResponse
	(*DeleteEducationByIDResponse)(nil), // 7: educations.DeleteEducationByIDResponse
}
var file_proto_educations_proto_depIdxs = []int32{
	6, // 0: educations.AllEducationsResponse.educations:type_name -> educations.EducationResponse
	0, // 1: educations.EducationService.GetEducations:input_type -> educations.GetEducationsRequest
	1, // 2: educations.EducationService.GetEducationByID:input_type -> educations.GetEducationByIDRequest
	3, // 3: educations.EducationService.CreateEducation:input_type -> educations.CreateEducationRequest
	5, // 4: educations.EducationService.DeleteEducationByID:input_type -> educations.DeleteEducationByIDRequest
	4, // 5: educations.EducationService.UpdateEducationByID:input_type -> educations.UpdateEducationByIDRequest
	2, // 6: educations.EducationService.GetEducations:output_type -> educations.AllEducationsResponse
	6, // 7: educations.EducationService.GetEducationByID:output_type -> educations.EducationResponse
	6, // 8: educations.EducationService.CreateEducation:output_type -> educations.EducationResponse
	7, // 9: educations.EducationService.DeleteEducationByID:output_type -> educations.DeleteEducationByIDResponse
	6, // 10: educations.EducationService.UpdateEducationByID:output_type -> educations.EducationResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_educations_proto_init() }
func file_proto_educations_proto_init() {
	if File_proto_educations_proto != nil {
		return
	}
	file_proto_educations_proto_msgTypes[3].OneofWrappers = []any{}
	file_proto_educations_proto_msgTypes[4].OneofWrappers = []any{}
	file_proto_educations_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_educations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_educations_proto_goTypes,
		DependencyIndexes: file_proto_educations_proto_depIdxs,
		MessageInfos:      file_proto_educations_proto_msgTypes,
	}.Build()
	File_proto_educations_proto = out.File
	file_proto_educations_proto_rawDesc = nil
	file_proto_educations_proto_goTypes = nil
	file_proto_educations_proto_depIdxs = nil
}