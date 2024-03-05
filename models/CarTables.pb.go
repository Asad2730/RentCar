// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: CarTables.proto

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

type Car struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AcHeater  bool   `protobuf:"varint,3,opt,name=ac_heater,json=acHeater,proto3" json:"ac_heater,omitempty"`
	CreatedAt string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt string `protobuf:"bytes,6,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Car) Reset() {
	*x = Car{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CarTables_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Car) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Car) ProtoMessage() {}

func (x *Car) ProtoReflect() protoreflect.Message {
	mi := &file_CarTables_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Car.ProtoReflect.Descriptor instead.
func (*Car) Descriptor() ([]byte, []int) {
	return file_CarTables_proto_rawDescGZIP(), []int{0}
}

func (x *Car) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Car) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Car) GetAcHeater() bool {
	if x != nil {
		return x.AcHeater
	}
	return false
}

func (x *Car) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Car) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Car) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type BodyType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Icon      string `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	CreatedAt string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt string `protobuf:"bytes,6,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	CarId     int32  `protobuf:"varint,7,opt,name=carId,proto3" json:"carId,omitempty"`
}

func (x *BodyType) Reset() {
	*x = BodyType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CarTables_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BodyType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BodyType) ProtoMessage() {}

func (x *BodyType) ProtoReflect() protoreflect.Message {
	mi := &file_CarTables_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BodyType.ProtoReflect.Descriptor instead.
func (*BodyType) Descriptor() ([]byte, []int) {
	return file_CarTables_proto_rawDescGZIP(), []int{1}
}

func (x *BodyType) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BodyType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BodyType) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *BodyType) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *BodyType) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *BodyType) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

func (x *BodyType) GetCarId() int32 {
	if x != nil {
		return x.CarId
	}
	return 0
}

type Model struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Year      int32  `protobuf:"varint,2,opt,name=year,proto3" json:"year,omitempty"`
	CreatedAt string `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt string `protobuf:"bytes,5,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	CarId     int32  `protobuf:"varint,6,opt,name=carId,proto3" json:"carId,omitempty"`
}

func (x *Model) Reset() {
	*x = Model{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CarTables_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Model) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Model) ProtoMessage() {}

func (x *Model) ProtoReflect() protoreflect.Message {
	mi := &file_CarTables_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Model.ProtoReflect.Descriptor instead.
func (*Model) Descriptor() ([]byte, []int) {
	return file_CarTables_proto_rawDescGZIP(), []int{2}
}

func (x *Model) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Model) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Model) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Model) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Model) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

func (x *Model) GetCarId() int32 {
	if x != nil {
		return x.CarId
	}
	return 0
}

type Manufacturer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Icon      string `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	CreatedAt string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt string `protobuf:"bytes,6,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	CarId     int32  `protobuf:"varint,7,opt,name=carId,proto3" json:"carId,omitempty"`
}

func (x *Manufacturer) Reset() {
	*x = Manufacturer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CarTables_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Manufacturer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Manufacturer) ProtoMessage() {}

func (x *Manufacturer) ProtoReflect() protoreflect.Message {
	mi := &file_CarTables_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Manufacturer.ProtoReflect.Descriptor instead.
func (*Manufacturer) Descriptor() ([]byte, []int) {
	return file_CarTables_proto_rawDescGZIP(), []int{3}
}

func (x *Manufacturer) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Manufacturer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Manufacturer) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *Manufacturer) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Manufacturer) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Manufacturer) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

func (x *Manufacturer) GetCarId() int32 {
	if x != nil {
		return x.CarId
	}
	return 0
}

type EngineType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Icon      string `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	CreatedAt string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt string `protobuf:"bytes,6,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	CarId     int32  `protobuf:"varint,7,opt,name=carId,proto3" json:"carId,omitempty"`
}

func (x *EngineType) Reset() {
	*x = EngineType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CarTables_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EngineType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EngineType) ProtoMessage() {}

func (x *EngineType) ProtoReflect() protoreflect.Message {
	mi := &file_CarTables_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EngineType.ProtoReflect.Descriptor instead.
func (*EngineType) Descriptor() ([]byte, []int) {
	return file_CarTables_proto_rawDescGZIP(), []int{4}
}

func (x *EngineType) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EngineType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EngineType) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *EngineType) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *EngineType) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *EngineType) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

func (x *EngineType) GetCarId() int32 {
	if x != nil {
		return x.CarId
	}
	return 0
}

type Color struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Color     string `protobuf:"bytes,2,opt,name=color,proto3" json:"color,omitempty"`
	CreatedAt string `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt string `protobuf:"bytes,5,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	CarId     int32  `protobuf:"varint,7,opt,name=carId,proto3" json:"carId,omitempty"`
}

func (x *Color) Reset() {
	*x = Color{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CarTables_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Color) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Color) ProtoMessage() {}

func (x *Color) ProtoReflect() protoreflect.Message {
	mi := &file_CarTables_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Color.ProtoReflect.Descriptor instead.
func (*Color) Descriptor() ([]byte, []int) {
	return file_CarTables_proto_rawDescGZIP(), []int{5}
}

func (x *Color) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Color) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *Color) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Color) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Color) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

func (x *Color) GetCarId() int32 {
	if x != nil {
		return x.CarId
	}
	return 0
}

type CarResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car          *Car          `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
	BodyType     *BodyType     `protobuf:"bytes,2,opt,name=bodyType,proto3" json:"bodyType,omitempty"`
	Color        *Color        `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
	EngineType   *EngineType   `protobuf:"bytes,4,opt,name=engineType,proto3" json:"engineType,omitempty"`
	Manufacturer *Manufacturer `protobuf:"bytes,5,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	Model        *Model        `protobuf:"bytes,6,opt,name=model,proto3" json:"model,omitempty"`
}

func (x *CarResult) Reset() {
	*x = CarResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CarTables_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarResult) ProtoMessage() {}

func (x *CarResult) ProtoReflect() protoreflect.Message {
	mi := &file_CarTables_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarResult.ProtoReflect.Descriptor instead.
func (*CarResult) Descriptor() ([]byte, []int) {
	return file_CarTables_proto_rawDescGZIP(), []int{6}
}

func (x *CarResult) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

func (x *CarResult) GetBodyType() *BodyType {
	if x != nil {
		return x.BodyType
	}
	return nil
}

func (x *CarResult) GetColor() *Color {
	if x != nil {
		return x.Color
	}
	return nil
}

func (x *CarResult) GetEngineType() *EngineType {
	if x != nil {
		return x.EngineType
	}
	return nil
}

func (x *CarResult) GetManufacturer() *Manufacturer {
	if x != nil {
		return x.Manufacturer
	}
	return nil
}

func (x *CarResult) GetModel() *Model {
	if x != nil {
		return x.Model
	}
	return nil
}

type CarList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cars []*CarResult `protobuf:"bytes,1,rep,name=cars,proto3" json:"cars,omitempty"`
}

func (x *CarList) Reset() {
	*x = CarList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CarTables_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarList) ProtoMessage() {}

func (x *CarList) ProtoReflect() protoreflect.Message {
	mi := &file_CarTables_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarList.ProtoReflect.Descriptor instead.
func (*CarList) Descriptor() ([]byte, []int) {
	return file_CarTables_proto_rawDescGZIP(), []int{7}
}

func (x *CarList) GetCars() []*CarResult {
	if x != nil {
		return x.Cars
	}
	return nil
}

var File_CarTables_proto protoreflect.FileDescriptor

var file_CarTables_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x43, 0x61, 0x72, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0xa3, 0x01, 0x0a, 0x03, 0x43, 0x61,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x63, 0x5f, 0x68, 0x65, 0x61, 0x74,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x63, 0x48, 0x65, 0x61, 0x74,
	0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0xb5, 0x01, 0x0a, 0x08, 0x42, 0x6f, 0x64, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x69, 0x63, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x63, 0x61, 0x72, 0x49, 0x64, 0x22, 0x9e, 0x01, 0x0a, 0x05, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x63, 0x61, 0x72, 0x49, 0x64, 0x22, 0xb9, 0x01, 0x0a, 0x0c, 0x4d, 0x61, 0x6e,
	0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f,
	0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x61, 0x72, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63,
	0x61, 0x72, 0x49, 0x64, 0x22, 0xb7, 0x01, 0x0a, 0x0a, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x49,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x61, 0x72, 0x49, 0x64, 0x22, 0xa0,
	0x01, 0x0a, 0x05, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x61, 0x72, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x61, 0x72, 0x49,
	0x64, 0x22, 0x90, 0x02, 0x0a, 0x09, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x1d, 0x0a, 0x03, 0x63, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x03, 0x63, 0x61, 0x72, 0x12, 0x2c,
	0x0a, 0x08, 0x62, 0x6f, 0x64, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x08, 0x62, 0x6f, 0x64, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x05,
	0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x12, 0x32, 0x0a, 0x0a, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x45,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x72, 0x52, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12,
	0x23, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x05, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x22, 0x30, 0x0a, 0x07, 0x43, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x25, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x04, 0x63, 0x61, 0x72, 0x73, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CarTables_proto_rawDescOnce sync.Once
	file_CarTables_proto_rawDescData = file_CarTables_proto_rawDesc
)

func file_CarTables_proto_rawDescGZIP() []byte {
	file_CarTables_proto_rawDescOnce.Do(func() {
		file_CarTables_proto_rawDescData = protoimpl.X.CompressGZIP(file_CarTables_proto_rawDescData)
	})
	return file_CarTables_proto_rawDescData
}

var file_CarTables_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_CarTables_proto_goTypes = []interface{}{
	(*Car)(nil),          // 0: models.Car
	(*BodyType)(nil),     // 1: models.BodyType
	(*Model)(nil),        // 2: models.Model
	(*Manufacturer)(nil), // 3: models.Manufacturer
	(*EngineType)(nil),   // 4: models.EngineType
	(*Color)(nil),        // 5: models.Color
	(*CarResult)(nil),    // 6: models.CarResult
	(*CarList)(nil),      // 7: models.CarList
}
var file_CarTables_proto_depIdxs = []int32{
	0, // 0: models.CarResult.car:type_name -> models.Car
	1, // 1: models.CarResult.bodyType:type_name -> models.BodyType
	5, // 2: models.CarResult.color:type_name -> models.Color
	4, // 3: models.CarResult.engineType:type_name -> models.EngineType
	3, // 4: models.CarResult.manufacturer:type_name -> models.Manufacturer
	2, // 5: models.CarResult.model:type_name -> models.Model
	6, // 6: models.CarList.cars:type_name -> models.CarResult
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_CarTables_proto_init() }
func file_CarTables_proto_init() {
	if File_CarTables_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CarTables_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Car); i {
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
		file_CarTables_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BodyType); i {
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
		file_CarTables_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Model); i {
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
		file_CarTables_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Manufacturer); i {
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
		file_CarTables_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EngineType); i {
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
		file_CarTables_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Color); i {
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
		file_CarTables_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarResult); i {
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
		file_CarTables_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarList); i {
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
			RawDescriptor: file_CarTables_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CarTables_proto_goTypes,
		DependencyIndexes: file_CarTables_proto_depIdxs,
		MessageInfos:      file_CarTables_proto_msgTypes,
	}.Build()
	File_CarTables_proto = out.File
	file_CarTables_proto_rawDesc = nil
	file_CarTables_proto_goTypes = nil
	file_CarTables_proto_depIdxs = nil
}
