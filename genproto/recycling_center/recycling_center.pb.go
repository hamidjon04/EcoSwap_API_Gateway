// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: recycling_center.proto

package recycling_center

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

type ResCenter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name              string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address           string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	AcceptedMaterials []string `protobuf:"bytes,3,rep,name=acceptedMaterials,proto3" json:"acceptedMaterials,omitempty"`
	WorkingHours      string   `protobuf:"bytes,4,opt,name=workingHours,proto3" json:"workingHours,omitempty"`
	ContactNumber     string   `protobuf:"bytes,5,opt,name=contactNumber,proto3" json:"contactNumber,omitempty"`
}

func (x *ResCenter) Reset() {
	*x = ResCenter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recycling_center_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResCenter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResCenter) ProtoMessage() {}

func (x *ResCenter) ProtoReflect() protoreflect.Message {
	mi := &file_recycling_center_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResCenter.ProtoReflect.Descriptor instead.
func (*ResCenter) Descriptor() ([]byte, []int) {
	return file_recycling_center_proto_rawDescGZIP(), []int{0}
}

func (x *ResCenter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResCenter) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ResCenter) GetAcceptedMaterials() []string {
	if x != nil {
		return x.AcceptedMaterials
	}
	return nil
}

func (x *ResCenter) GetWorkingHours() string {
	if x != nil {
		return x.WorkingHours
	}
	return ""
}

func (x *ResCenter) GetContactNumber() string {
	if x != nil {
		return x.ContactNumber
	}
	return ""
}

type ResponceResCenter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name              string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address           string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	AcceptedMaterials []string `protobuf:"bytes,4,rep,name=acceptedMaterials,proto3" json:"acceptedMaterials,omitempty"`
	WorkingHours      string   `protobuf:"bytes,5,opt,name=workingHours,proto3" json:"workingHours,omitempty"`
	ContactNumber     string   `protobuf:"bytes,6,opt,name=contactNumber,proto3" json:"contactNumber,omitempty"`
	CreatedAt         string   `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *ResponceResCenter) Reset() {
	*x = ResponceResCenter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recycling_center_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponceResCenter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponceResCenter) ProtoMessage() {}

func (x *ResponceResCenter) ProtoReflect() protoreflect.Message {
	mi := &file_recycling_center_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponceResCenter.ProtoReflect.Descriptor instead.
func (*ResponceResCenter) Descriptor() ([]byte, []int) {
	return file_recycling_center_proto_rawDescGZIP(), []int{1}
}

func (x *ResponceResCenter) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResponceResCenter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResponceResCenter) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ResponceResCenter) GetAcceptedMaterials() []string {
	if x != nil {
		return x.AcceptedMaterials
	}
	return nil
}

func (x *ResponceResCenter) GetWorkingHours() string {
	if x != nil {
		return x.WorkingHours
	}
	return ""
}

func (x *ResponceResCenter) GetContactNumber() string {
	if x != nil {
		return x.ContactNumber
	}
	return ""
}

func (x *ResponceResCenter) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type FilterField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Material string `protobuf:"bytes,1,opt,name=material,proto3" json:"material,omitempty"`
	Limit    int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset   int32  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *FilterField) Reset() {
	*x = FilterField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recycling_center_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterField) ProtoMessage() {}

func (x *FilterField) ProtoReflect() protoreflect.Message {
	mi := &file_recycling_center_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterField.ProtoReflect.Descriptor instead.
func (*FilterField) Descriptor() ([]byte, []int) {
	return file_recycling_center_proto_rawDescGZIP(), []int{2}
}

func (x *FilterField) GetMaterial() string {
	if x != nil {
		return x.Material
	}
	return ""
}

func (x *FilterField) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FilterField) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type Center struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name              string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address           string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	AcceptedMaterials []string `protobuf:"bytes,4,rep,name=acceptedMaterials,proto3" json:"acceptedMaterials,omitempty"`
	WorkingHours      string   `protobuf:"bytes,5,opt,name=workingHours,proto3" json:"workingHours,omitempty"`
	ContactNumber     string   `protobuf:"bytes,6,opt,name=contactNumber,proto3" json:"contactNumber,omitempty"`
}

func (x *Center) Reset() {
	*x = Center{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recycling_center_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Center) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Center) ProtoMessage() {}

func (x *Center) ProtoReflect() protoreflect.Message {
	mi := &file_recycling_center_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Center.ProtoReflect.Descriptor instead.
func (*Center) Descriptor() ([]byte, []int) {
	return file_recycling_center_proto_rawDescGZIP(), []int{3}
}

func (x *Center) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Center) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Center) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Center) GetAcceptedMaterials() []string {
	if x != nil {
		return x.AcceptedMaterials
	}
	return nil
}

func (x *Center) GetWorkingHours() string {
	if x != nil {
		return x.WorkingHours
	}
	return ""
}

func (x *Center) GetContactNumber() string {
	if x != nil {
		return x.ContactNumber
	}
	return ""
}

type ResAllCenter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Centers []*Center `protobuf:"bytes,1,rep,name=centers,proto3" json:"centers,omitempty"`
	Total   int32     `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Page    int32     `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit   int32     `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ResAllCenter) Reset() {
	*x = ResAllCenter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recycling_center_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResAllCenter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResAllCenter) ProtoMessage() {}

func (x *ResAllCenter) ProtoReflect() protoreflect.Message {
	mi := &file_recycling_center_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResAllCenter.ProtoReflect.Descriptor instead.
func (*ResAllCenter) Descriptor() ([]byte, []int) {
	return file_recycling_center_proto_rawDescGZIP(), []int{4}
}

func (x *ResAllCenter) GetCenters() []*Center {
	if x != nil {
		return x.Centers
	}
	return nil
}

func (x *ResAllCenter) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ResAllCenter) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ResAllCenter) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type Submission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CenterId      string `protobuf:"bytes,1,opt,name=centerId,proto3" json:"centerId,omitempty"`
	UserId        string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	JsonDataItems string `protobuf:"bytes,3,opt,name=jsonDataItems,proto3" json:"jsonDataItems,omitempty"`
}

func (x *Submission) Reset() {
	*x = Submission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recycling_center_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Submission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Submission) ProtoMessage() {}

func (x *Submission) ProtoReflect() protoreflect.Message {
	mi := &file_recycling_center_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Submission.ProtoReflect.Descriptor instead.
func (*Submission) Descriptor() ([]byte, []int) {
	return file_recycling_center_proto_rawDescGZIP(), []int{5}
}

func (x *Submission) GetCenterId() string {
	if x != nil {
		return x.CenterId
	}
	return ""
}

func (x *Submission) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Submission) GetJsonDataItems() string {
	if x != nil {
		return x.JsonDataItems
	}
	return ""
}

type SubmissionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CenterId        string `protobuf:"bytes,2,opt,name=centerId,proto3" json:"centerId,omitempty"`
	UserId          string `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
	JsonDataItems   string `protobuf:"bytes,4,opt,name=jsonDataItems,proto3" json:"jsonDataItems,omitempty"`
	EcoPointsEarned int32  `protobuf:"varint,5,opt,name=ecoPointsEarned,proto3" json:"ecoPointsEarned,omitempty"`
	CreatedAt       string `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *SubmissionResp) Reset() {
	*x = SubmissionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recycling_center_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmissionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmissionResp) ProtoMessage() {}

func (x *SubmissionResp) ProtoReflect() protoreflect.Message {
	mi := &file_recycling_center_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmissionResp.ProtoReflect.Descriptor instead.
func (*SubmissionResp) Descriptor() ([]byte, []int) {
	return file_recycling_center_proto_rawDescGZIP(), []int{6}
}

func (x *SubmissionResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SubmissionResp) GetCenterId() string {
	if x != nil {
		return x.CenterId
	}
	return ""
}

func (x *SubmissionResp) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SubmissionResp) GetJsonDataItems() string {
	if x != nil {
		return x.JsonDataItems
	}
	return ""
}

func (x *SubmissionResp) GetEcoPointsEarned() int32 {
	if x != nil {
		return x.EcoPointsEarned
	}
	return 0
}

func (x *SubmissionResp) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type StatisticField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartDate string `protobuf:"bytes,1,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate   string `protobuf:"bytes,2,opt,name=endDate,proto3" json:"endDate,omitempty"`
}

func (x *StatisticField) Reset() {
	*x = StatisticField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recycling_center_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatisticField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatisticField) ProtoMessage() {}

func (x *StatisticField) ProtoReflect() protoreflect.Message {
	mi := &file_recycling_center_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatisticField.ProtoReflect.Descriptor instead.
func (*StatisticField) Descriptor() ([]byte, []int) {
	return file_recycling_center_proto_rawDescGZIP(), []int{7}
}

func (x *StatisticField) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *StatisticField) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

type Statistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalSwaps         int32  `protobuf:"varint,1,opt,name=TotalSwaps,proto3" json:"TotalSwaps,omitempty"`
	TotalRecycledItems int32  `protobuf:"varint,2,opt,name=TotalRecycledItems,proto3" json:"TotalRecycledItems,omitempty"`
	TotalEarnedPoints  int32  `protobuf:"varint,3,opt,name=TotalEarnedPoints,proto3" json:"TotalEarnedPoints,omitempty"`
	JsonDataTopCateg   string `protobuf:"bytes,4,opt,name=jsonDataTopCateg,proto3" json:"jsonDataTopCateg,omitempty"`
	JsonDateTopCenter  string `protobuf:"bytes,5,opt,name=jsonDateTopCenter,proto3" json:"jsonDateTopCenter,omitempty"`
}

func (x *Statistics) Reset() {
	*x = Statistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recycling_center_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Statistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Statistics) ProtoMessage() {}

func (x *Statistics) ProtoReflect() protoreflect.Message {
	mi := &file_recycling_center_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Statistics.ProtoReflect.Descriptor instead.
func (*Statistics) Descriptor() ([]byte, []int) {
	return file_recycling_center_proto_rawDescGZIP(), []int{8}
}

func (x *Statistics) GetTotalSwaps() int32 {
	if x != nil {
		return x.TotalSwaps
	}
	return 0
}

func (x *Statistics) GetTotalRecycledItems() int32 {
	if x != nil {
		return x.TotalRecycledItems
	}
	return 0
}

func (x *Statistics) GetTotalEarnedPoints() int32 {
	if x != nil {
		return x.TotalEarnedPoints
	}
	return 0
}

func (x *Statistics) GetJsonDataTopCateg() string {
	if x != nil {
		return x.JsonDataTopCateg
	}
	return ""
}

func (x *Statistics) GetJsonDateTopCenter() string {
	if x != nil {
		return x.JsonDateTopCenter
	}
	return ""
}

var File_recycling_center_proto protoreflect.FileDescriptor

var file_recycling_center_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c,
	0x69, 0x6e, 0x67, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x22, 0xb1, 0x01, 0x0a, 0x09, 0x52,
	0x65, 0x73, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x65, 0x64, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x11, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x4d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x48,
	0x6f, 0x75, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6b,
	0x69, 0x6e, 0x67, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xe7,
	0x01, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x43, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x4d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x61,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73,
	0x12, 0x22, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x48, 0x6f, 0x75, 0x72, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x48,
	0x6f, 0x75, 0x72, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x57, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x22, 0xbe, 0x01, 0x0a, 0x06, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x61, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x4d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b,
	0x69, 0x6e, 0x67, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x24, 0x0a, 0x0d,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x82, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x41, 0x6c, 0x6c, 0x43, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x07, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67,
	0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x07,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x66, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x6a, 0x73, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x6a, 0x73, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22,
	0xc2, 0x01, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x6a, 0x73, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6a,
	0x73, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x28, 0x0a, 0x0f,
	0x65, 0x63, 0x6f, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x45, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x65, 0x63, 0x6f, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x45, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x48, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0xe4,
	0x01, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x77, 0x61, 0x70, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x77, 0x61, 0x70, 0x73, 0x12, 0x2e, 0x0a,
	0x12, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x64, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x2c, 0x0a,
	0x11, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x45,
	0x61, 0x72, 0x6e, 0x65, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x6a,
	0x73, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x6f, 0x70, 0x43, 0x61, 0x74, 0x65, 0x67, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6a, 0x73, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x54,
	0x6f, 0x70, 0x43, 0x61, 0x74, 0x65, 0x67, 0x12, 0x2c, 0x0a, 0x11, 0x6a, 0x73, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x6a, 0x73, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x43,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x32, 0xf0, 0x02, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c,
	0x69, 0x6e, 0x67, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x5b, 0x0a, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x12, 0x1b, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x1a,
	0x23, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x43, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12,
	0x1d, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x1a, 0x1e,
	0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x52, 0x65, 0x73, 0x41, 0x6c, 0x6c, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x22, 0x00,
	0x12, 0x53, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x12, 0x1c, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x5f,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x1a, 0x20, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x20, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69,
	0x6e, 0x67, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x1a, 0x1c, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63,
	0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x22, 0x00, 0x42, 0x1c, 0x5a, 0x1a, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x5f,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_recycling_center_proto_rawDescOnce sync.Once
	file_recycling_center_proto_rawDescData = file_recycling_center_proto_rawDesc
)

func file_recycling_center_proto_rawDescGZIP() []byte {
	file_recycling_center_proto_rawDescOnce.Do(func() {
		file_recycling_center_proto_rawDescData = protoimpl.X.CompressGZIP(file_recycling_center_proto_rawDescData)
	})
	return file_recycling_center_proto_rawDescData
}

var file_recycling_center_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_recycling_center_proto_goTypes = []any{
	(*ResCenter)(nil),         // 0: recycling_center.ResCenter
	(*ResponceResCenter)(nil), // 1: recycling_center.ResponceResCenter
	(*FilterField)(nil),       // 2: recycling_center.FilterField
	(*Center)(nil),            // 3: recycling_center.Center
	(*ResAllCenter)(nil),      // 4: recycling_center.ResAllCenter
	(*Submission)(nil),        // 5: recycling_center.Submission
	(*SubmissionResp)(nil),    // 6: recycling_center.SubmissionResp
	(*StatisticField)(nil),    // 7: recycling_center.StatisticField
	(*Statistics)(nil),        // 8: recycling_center.Statistics
}
var file_recycling_center_proto_depIdxs = []int32{
	3, // 0: recycling_center.ResAllCenter.centers:type_name -> recycling_center.Center
	0, // 1: recycling_center.RecyclingCenter.CreateRecyclingCenter:input_type -> recycling_center.ResCenter
	2, // 2: recycling_center.RecyclingCenter.SearchRecyclingCenter:input_type -> recycling_center.FilterField
	5, // 3: recycling_center.RecyclingCenter.ProductDelivery:input_type -> recycling_center.Submission
	7, // 4: recycling_center.RecyclingCenter.GetStatistics:input_type -> recycling_center.StatisticField
	1, // 5: recycling_center.RecyclingCenter.CreateRecyclingCenter:output_type -> recycling_center.ResponceResCenter
	4, // 6: recycling_center.RecyclingCenter.SearchRecyclingCenter:output_type -> recycling_center.ResAllCenter
	6, // 7: recycling_center.RecyclingCenter.ProductDelivery:output_type -> recycling_center.SubmissionResp
	8, // 8: recycling_center.RecyclingCenter.GetStatistics:output_type -> recycling_center.Statistics
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_recycling_center_proto_init() }
func file_recycling_center_proto_init() {
	if File_recycling_center_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_recycling_center_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ResCenter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_recycling_center_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ResponceResCenter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_recycling_center_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*FilterField); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_recycling_center_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Center); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_recycling_center_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ResAllCenter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_recycling_center_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Submission); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_recycling_center_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*SubmissionResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_recycling_center_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*StatisticField); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_recycling_center_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*Statistics); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_recycling_center_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_recycling_center_proto_goTypes,
		DependencyIndexes: file_recycling_center_proto_depIdxs,
		MessageInfos:      file_recycling_center_proto_msgTypes,
	}.Build()
	File_recycling_center_proto = out.File
	file_recycling_center_proto_rawDesc = nil
	file_recycling_center_proto_goTypes = nil
	file_recycling_center_proto_depIdxs = nil
}
