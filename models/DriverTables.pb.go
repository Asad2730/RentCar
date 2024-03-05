// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: DriverTables.proto

package models

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

type Driver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Dob       string `protobuf:"bytes,3,opt,name=dob,proto3" json:"dob,omitempty"`
	CreatedAt string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt string `protobuf:"bytes,6,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Driver) Reset() {
	*x = Driver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DriverTables_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Driver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Driver) ProtoMessage() {}

func (x *Driver) ProtoReflect() protoreflect.Message {
	mi := &file_DriverTables_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Driver.ProtoReflect.Descriptor instead.
func (*Driver) Descriptor() ([]byte, []int) {
	return file_DriverTables_proto_rawDescGZIP(), []int{0}
}

func (x *Driver) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Driver) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Driver) GetDob() string {
	if x != nil {
		return x.Dob
	}
	return ""
}

func (x *Driver) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Driver) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Driver) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address   string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Location  []byte `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	CreatedAt string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt string `protobuf:"bytes,6,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	DriverId  int32  `protobuf:"varint,7,opt,name=driverId,proto3" json:"driverId,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DriverTables_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_DriverTables_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_DriverTables_proto_rawDescGZIP(), []int{1}
}

func (x *Address) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Address) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Address) GetLocation() []byte {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *Address) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Address) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Address) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

func (x *Address) GetDriverId() int32 {
	if x != nil {
		return x.DriverId
	}
	return 0
}

type Document struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DocumentTypeId int32  `protobuf:"varint,2,opt,name=documentTypeId,proto3" json:"documentTypeId,omitempty"`
	DocumentUrl    string `protobuf:"bytes,3,opt,name=documentUrl,proto3" json:"documentUrl,omitempty"`
	CreatedAt      string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt      string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt      string `protobuf:"bytes,6,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	DriverId       int32  `protobuf:"varint,7,opt,name=driverId,proto3" json:"driverId,omitempty"`
}

func (x *Document) Reset() {
	*x = Document{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DriverTables_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Document) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Document) ProtoMessage() {}

func (x *Document) ProtoReflect() protoreflect.Message {
	mi := &file_DriverTables_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Document.ProtoReflect.Descriptor instead.
func (*Document) Descriptor() ([]byte, []int) {
	return file_DriverTables_proto_rawDescGZIP(), []int{2}
}

func (x *Document) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Document) GetDocumentTypeId() int32 {
	if x != nil {
		return x.DocumentTypeId
	}
	return 0
}

func (x *Document) GetDocumentUrl() string {
	if x != nil {
		return x.DocumentUrl
	}
	return ""
}

func (x *Document) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Document) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Document) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

func (x *Document) GetDriverId() int32 {
	if x != nil {
		return x.DriverId
	}
	return 0
}

type DocumentType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type      string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	CreatedAt string `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt string `protobuf:"bytes,5,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	DriverId  int32  `protobuf:"varint,6,opt,name=driverId,proto3" json:"driverId,omitempty"`
}

func (x *DocumentType) Reset() {
	*x = DocumentType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DriverTables_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentType) ProtoMessage() {}

func (x *DocumentType) ProtoReflect() protoreflect.Message {
	mi := &file_DriverTables_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentType.ProtoReflect.Descriptor instead.
func (*DocumentType) Descriptor() ([]byte, []int) {
	return file_DriverTables_proto_rawDescGZIP(), []int{3}
}

func (x *DocumentType) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DocumentType) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DocumentType) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *DocumentType) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *DocumentType) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

func (x *DocumentType) GetDriverId() int32 {
	if x != nil {
		return x.DriverId
	}
	return 0
}

type DocumentResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Document     *Document     `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	DocumentType *DocumentType `protobuf:"bytes,2,opt,name=documentType,proto3" json:"documentType,omitempty"`
}

func (x *DocumentResult) Reset() {
	*x = DocumentResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DriverTables_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentResult) ProtoMessage() {}

func (x *DocumentResult) ProtoReflect() protoreflect.Message {
	mi := &file_DriverTables_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentResult.ProtoReflect.Descriptor instead.
func (*DocumentResult) Descriptor() ([]byte, []int) {
	return file_DriverTables_proto_rawDescGZIP(), []int{4}
}

func (x *DocumentResult) GetDocument() *Document {
	if x != nil {
		return x.Document
	}
	return nil
}

func (x *DocumentResult) GetDocumentType() *DocumentType {
	if x != nil {
		return x.DocumentType
	}
	return nil
}

type DocumentTypeList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DocumentTypeList []*DocumentType `protobuf:"bytes,1,rep,name=documentTypeList,proto3" json:"documentTypeList,omitempty"`
}

func (x *DocumentTypeList) Reset() {
	*x = DocumentTypeList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DriverTables_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentTypeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentTypeList) ProtoMessage() {}

func (x *DocumentTypeList) ProtoReflect() protoreflect.Message {
	mi := &file_DriverTables_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentTypeList.ProtoReflect.Descriptor instead.
func (*DocumentTypeList) Descriptor() ([]byte, []int) {
	return file_DriverTables_proto_rawDescGZIP(), []int{5}
}

func (x *DocumentTypeList) GetDocumentTypeList() []*DocumentType {
	if x != nil {
		return x.DocumentTypeList
	}
	return nil
}

type DriverResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Driver         *Driver         `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	Address        *Address        `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	DocumentResult *DocumentResult `protobuf:"bytes,3,opt,name=documentResult,proto3" json:"documentResult,omitempty"`
}

func (x *DriverResult) Reset() {
	*x = DriverResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DriverTables_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriverResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverResult) ProtoMessage() {}

func (x *DriverResult) ProtoReflect() protoreflect.Message {
	mi := &file_DriverTables_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverResult.ProtoReflect.Descriptor instead.
func (*DriverResult) Descriptor() ([]byte, []int) {
	return file_DriverTables_proto_rawDescGZIP(), []int{6}
}

func (x *DriverResult) GetDriver() *Driver {
	if x != nil {
		return x.Driver
	}
	return nil
}

func (x *DriverResult) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *DriverResult) GetDocumentResult() *DocumentResult {
	if x != nil {
		return x.DocumentResult
	}
	return nil
}

type DriverList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverResult []*DriverResult `protobuf:"bytes,1,rep,name=driverResult,proto3" json:"driverResult,omitempty"`
}

func (x *DriverList) Reset() {
	*x = DriverList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DriverTables_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriverList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverList) ProtoMessage() {}

func (x *DriverList) ProtoReflect() protoreflect.Message {
	mi := &file_DriverTables_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverList.ProtoReflect.Descriptor instead.
func (*DriverList) Descriptor() ([]byte, []int) {
	return file_DriverTables_proto_rawDescGZIP(), []int{7}
}

func (x *DriverList) GetDriverResult() []*DriverResult {
	if x != nil {
		return x.DriverResult
	}
	return nil
}

var File_DriverTables_proto protoreflect.FileDescriptor

var file_DriverTables_proto_rawDesc = []byte{
	0x0a, 0x12, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0x9b, 0x01, 0x0a,
	0x06, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64,
	0x6f, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6f, 0x62, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xc8, 0x01, 0x0a, 0x07, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0xdd, 0x01, 0x0a, 0x08, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0xab, 0x01, 0x0a, 0x0c, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x78, 0x0a, 0x0e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x0c, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0c, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x54, 0x0a,
	0x10, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x40, 0x0a, 0x10, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x10, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0xa1, 0x01, 0x0a, 0x0c, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3e, 0x0a, 0x0e, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x46, 0x0a, 0x0a, 0x44, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0c, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x0c, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_DriverTables_proto_rawDescOnce sync.Once
	file_DriverTables_proto_rawDescData = file_DriverTables_proto_rawDesc
)

func file_DriverTables_proto_rawDescGZIP() []byte {
	file_DriverTables_proto_rawDescOnce.Do(func() {
		file_DriverTables_proto_rawDescData = protoimpl.X.CompressGZIP(file_DriverTables_proto_rawDescData)
	})
	return file_DriverTables_proto_rawDescData
}

var file_DriverTables_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_DriverTables_proto_goTypes = []interface{}{
	(*Driver)(nil),           // 0: models.Driver
	(*Address)(nil),          // 1: models.Address
	(*Document)(nil),         // 2: models.Document
	(*DocumentType)(nil),     // 3: models.DocumentType
	(*DocumentResult)(nil),   // 4: models.DocumentResult
	(*DocumentTypeList)(nil), // 5: models.DocumentTypeList
	(*DriverResult)(nil),     // 6: models.DriverResult
	(*DriverList)(nil),       // 7: models.DriverList
}
var file_DriverTables_proto_depIdxs = []int32{
	2, // 0: models.DocumentResult.document:type_name -> models.Document
	3, // 1: models.DocumentResult.documentType:type_name -> models.DocumentType
	3, // 2: models.DocumentTypeList.documentTypeList:type_name -> models.DocumentType
	0, // 3: models.DriverResult.driver:type_name -> models.Driver
	1, // 4: models.DriverResult.address:type_name -> models.Address
	4, // 5: models.DriverResult.documentResult:type_name -> models.DocumentResult
	6, // 6: models.DriverList.driverResult:type_name -> models.DriverResult
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_DriverTables_proto_init() }
func file_DriverTables_proto_init() {
	if File_DriverTables_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_DriverTables_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Driver); i {
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
		file_DriverTables_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_DriverTables_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Document); i {
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
		file_DriverTables_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentType); i {
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
		file_DriverTables_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentResult); i {
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
		file_DriverTables_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentTypeList); i {
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
		file_DriverTables_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DriverResult); i {
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
		file_DriverTables_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DriverList); i {
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
			RawDescriptor: file_DriverTables_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DriverTables_proto_goTypes,
		DependencyIndexes: file_DriverTables_proto_depIdxs,
		MessageInfos:      file_DriverTables_proto_msgTypes,
	}.Build()
	File_DriverTables_proto = out.File
	file_DriverTables_proto_rawDesc = nil
	file_DriverTables_proto_goTypes = nil
	file_DriverTables_proto_depIdxs = nil
}
