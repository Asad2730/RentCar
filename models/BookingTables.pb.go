// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: BookingTables.proto

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

type BookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Description    string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Phone          string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	PickupLocation []byte `protobuf:"bytes,4,opt,name=pickup_location,json=pickupLocation,proto3" json:"pickup_location,omitempty"`
	BoId           int32  `protobuf:"varint,5,opt,name=boId,proto3" json:"boId,omitempty"`
	CustomerId     int32  `protobuf:"varint,6,opt,name=customerId,proto3" json:"customerId,omitempty"`
	CarDetails     []byte `protobuf:"bytes,7,opt,name=carDetails,proto3" json:"carDetails,omitempty"`
	Status         string `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt      string `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt      string `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt      string `protobuf:"bytes,11,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *BookingRequest) Reset() {
	*x = BookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BookingTables_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingRequest) ProtoMessage() {}

func (x *BookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_BookingTables_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingRequest.ProtoReflect.Descriptor instead.
func (*BookingRequest) Descriptor() ([]byte, []int) {
	return file_BookingTables_proto_rawDescGZIP(), []int{0}
}

func (x *BookingRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BookingRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BookingRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *BookingRequest) GetPickupLocation() []byte {
	if x != nil {
		return x.PickupLocation
	}
	return nil
}

func (x *BookingRequest) GetBoId() int32 {
	if x != nil {
		return x.BoId
	}
	return 0
}

func (x *BookingRequest) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *BookingRequest) GetCarDetails() []byte {
	if x != nil {
		return x.CarDetails
	}
	return nil
}

func (x *BookingRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *BookingRequest) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *BookingRequest) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *BookingRequest) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type BookingClosing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BookingId       int32  `protobuf:"varint,2,opt,name=bookingId,proto3" json:"bookingId,omitempty"`
	DriverId        int32  `protobuf:"varint,3,opt,name=driverId,proto3" json:"driverId,omitempty"`
	BoId            int32  `protobuf:"varint,4,opt,name=boId,proto3" json:"boId,omitempty"`
	DropoffLocation []byte `protobuf:"bytes,5,opt,name=dropoff_location,json=dropoffLocation,proto3" json:"dropoff_location,omitempty"`
	CreatedAt       string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt       string `protobuf:"bytes,8,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *BookingClosing) Reset() {
	*x = BookingClosing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BookingTables_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingClosing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingClosing) ProtoMessage() {}

func (x *BookingClosing) ProtoReflect() protoreflect.Message {
	mi := &file_BookingTables_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingClosing.ProtoReflect.Descriptor instead.
func (*BookingClosing) Descriptor() ([]byte, []int) {
	return file_BookingTables_proto_rawDescGZIP(), []int{1}
}

func (x *BookingClosing) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BookingClosing) GetBookingId() int32 {
	if x != nil {
		return x.BookingId
	}
	return 0
}

func (x *BookingClosing) GetDriverId() int32 {
	if x != nil {
		return x.DriverId
	}
	return 0
}

func (x *BookingClosing) GetBoId() int32 {
	if x != nil {
		return x.BoId
	}
	return 0
}

func (x *BookingClosing) GetDropoffLocation() []byte {
	if x != nil {
		return x.DropoffLocation
	}
	return nil
}

func (x *BookingClosing) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *BookingClosing) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *BookingClosing) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type BookingDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BookingId int32  `protobuf:"varint,2,opt,name=bookingId,proto3" json:"bookingId,omitempty"`
	DriverId  int32  `protobuf:"varint,3,opt,name=driverId,proto3" json:"driverId,omitempty"`
	CarId     int32  `protobuf:"varint,4,opt,name=carId,proto3" json:"carId,omitempty"`
	CreatedAt string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt string `protobuf:"bytes,7,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *BookingDetail) Reset() {
	*x = BookingDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BookingTables_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingDetail) ProtoMessage() {}

func (x *BookingDetail) ProtoReflect() protoreflect.Message {
	mi := &file_BookingTables_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingDetail.ProtoReflect.Descriptor instead.
func (*BookingDetail) Descriptor() ([]byte, []int) {
	return file_BookingTables_proto_rawDescGZIP(), []int{2}
}

func (x *BookingDetail) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BookingDetail) GetBookingId() int32 {
	if x != nil {
		return x.BookingId
	}
	return 0
}

func (x *BookingDetail) GetDriverId() int32 {
	if x != nil {
		return x.DriverId
	}
	return 0
}

func (x *BookingDetail) GetCarId() int32 {
	if x != nil {
		return x.CarId
	}
	return 0
}

func (x *BookingDetail) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *BookingDetail) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *BookingDetail) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type BookingFeedBack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BookingId   int32  `protobuf:"varint,2,opt,name=bookingId,proto3" json:"bookingId,omitempty"`
	BoId        int32  `protobuf:"varint,3,opt,name=boId,proto3" json:"boId,omitempty"`
	RatingScore int32  `protobuf:"varint,4,opt,name=ratingScore,proto3" json:"ratingScore,omitempty"`
	FeedBack    string `protobuf:"bytes,5,opt,name=feedBack,proto3" json:"feedBack,omitempty"`
	CreatedAt   string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt   string `protobuf:"bytes,8,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *BookingFeedBack) Reset() {
	*x = BookingFeedBack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BookingTables_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingFeedBack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingFeedBack) ProtoMessage() {}

func (x *BookingFeedBack) ProtoReflect() protoreflect.Message {
	mi := &file_BookingTables_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingFeedBack.ProtoReflect.Descriptor instead.
func (*BookingFeedBack) Descriptor() ([]byte, []int) {
	return file_BookingTables_proto_rawDescGZIP(), []int{3}
}

func (x *BookingFeedBack) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BookingFeedBack) GetBookingId() int32 {
	if x != nil {
		return x.BookingId
	}
	return 0
}

func (x *BookingFeedBack) GetBoId() int32 {
	if x != nil {
		return x.BoId
	}
	return 0
}

func (x *BookingFeedBack) GetRatingScore() int32 {
	if x != nil {
		return x.RatingScore
	}
	return 0
}

func (x *BookingFeedBack) GetFeedBack() string {
	if x != nil {
		return x.FeedBack
	}
	return ""
}

func (x *BookingFeedBack) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *BookingFeedBack) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *BookingFeedBack) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type BookingResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookingRequest  *BookingRequest  `protobuf:"bytes,1,opt,name=bookingRequest,proto3" json:"bookingRequest,omitempty"`
	BookingClosing  *BookingClosing  `protobuf:"bytes,2,opt,name=bookingClosing,proto3" json:"bookingClosing,omitempty"`
	BookingDetail   *BookingDetail   `protobuf:"bytes,3,opt,name=bookingDetail,proto3" json:"bookingDetail,omitempty"`
	BookingFeedBack *BookingFeedBack `protobuf:"bytes,4,opt,name=bookingFeedBack,proto3" json:"bookingFeedBack,omitempty"`
}

func (x *BookingResult) Reset() {
	*x = BookingResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BookingTables_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingResult) ProtoMessage() {}

func (x *BookingResult) ProtoReflect() protoreflect.Message {
	mi := &file_BookingTables_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingResult.ProtoReflect.Descriptor instead.
func (*BookingResult) Descriptor() ([]byte, []int) {
	return file_BookingTables_proto_rawDescGZIP(), []int{4}
}

func (x *BookingResult) GetBookingRequest() *BookingRequest {
	if x != nil {
		return x.BookingRequest
	}
	return nil
}

func (x *BookingResult) GetBookingClosing() *BookingClosing {
	if x != nil {
		return x.BookingClosing
	}
	return nil
}

func (x *BookingResult) GetBookingDetail() *BookingDetail {
	if x != nil {
		return x.BookingDetail
	}
	return nil
}

func (x *BookingResult) GetBookingFeedBack() *BookingFeedBack {
	if x != nil {
		return x.BookingFeedBack
	}
	return nil
}

type BookingList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bookings []*BookingResult `protobuf:"bytes,1,rep,name=bookings,proto3" json:"bookings,omitempty"`
}

func (x *BookingList) Reset() {
	*x = BookingList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BookingTables_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingList) ProtoMessage() {}

func (x *BookingList) ProtoReflect() protoreflect.Message {
	mi := &file_BookingTables_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingList.ProtoReflect.Descriptor instead.
func (*BookingList) Descriptor() ([]byte, []int) {
	return file_BookingTables_proto_rawDescGZIP(), []int{5}
}

func (x *BookingList) GetBookings() []*BookingResult {
	if x != nil {
		return x.Bookings
	}
	return nil
}

var File_BookingTables_proto protoreflect.FileDescriptor

var file_BookingTables_proto_rawDesc = []byte{
	0x0a, 0x13, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0xca, 0x02,
	0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x69, 0x63, 0x6b,
	0x75, 0x70, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0e, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x62, 0x6f, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xf6, 0x01, 0x0a, 0x0e, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x6c, 0x6f, 0x73, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x49, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x62, 0x6f, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x64,
	0x72, 0x6f, 0x70, 0x6f, 0x66, 0x66, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x64, 0x72, 0x6f, 0x70, 0x6f, 0x66, 0x66, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0xcc, 0x01, 0x0a, 0x0d, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x63, 0x61, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0xee, 0x01, 0x0a, 0x0f, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x46, 0x65,
	0x65, 0x64, 0x42, 0x61, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x62, 0x6f, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x65,
	0x65, 0x64, 0x42, 0x61, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x65,
	0x65, 0x64, 0x42, 0x61, 0x63, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x8f, 0x02, 0x0a, 0x0d, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3e, 0x0a, 0x0e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x43, 0x6c, 0x6f, 0x73, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x6c,
	0x6f, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x0e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x6c,
	0x6f, 0x73, 0x69, 0x6e, 0x67, 0x12, 0x3b, 0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x0d, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x12, 0x41, 0x0a, 0x0f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x46, 0x65, 0x65,
	0x64, 0x42, 0x61, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x46, 0x65, 0x65, 0x64,
	0x42, 0x61, 0x63, 0x6b, 0x52, 0x0f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x46, 0x65, 0x65,
	0x64, 0x42, 0x61, 0x63, 0x6b, 0x22, 0x40, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x08, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BookingTables_proto_rawDescOnce sync.Once
	file_BookingTables_proto_rawDescData = file_BookingTables_proto_rawDesc
)

func file_BookingTables_proto_rawDescGZIP() []byte {
	file_BookingTables_proto_rawDescOnce.Do(func() {
		file_BookingTables_proto_rawDescData = protoimpl.X.CompressGZIP(file_BookingTables_proto_rawDescData)
	})
	return file_BookingTables_proto_rawDescData
}

var file_BookingTables_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_BookingTables_proto_goTypes = []interface{}{
	(*BookingRequest)(nil),  // 0: models.BookingRequest
	(*BookingClosing)(nil),  // 1: models.BookingClosing
	(*BookingDetail)(nil),   // 2: models.BookingDetail
	(*BookingFeedBack)(nil), // 3: models.BookingFeedBack
	(*BookingResult)(nil),   // 4: models.BookingResult
	(*BookingList)(nil),     // 5: models.BookingList
}
var file_BookingTables_proto_depIdxs = []int32{
	0, // 0: models.BookingResult.bookingRequest:type_name -> models.BookingRequest
	1, // 1: models.BookingResult.bookingClosing:type_name -> models.BookingClosing
	2, // 2: models.BookingResult.bookingDetail:type_name -> models.BookingDetail
	3, // 3: models.BookingResult.bookingFeedBack:type_name -> models.BookingFeedBack
	4, // 4: models.BookingList.bookings:type_name -> models.BookingResult
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_BookingTables_proto_init() }
func file_BookingTables_proto_init() {
	if File_BookingTables_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_BookingTables_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingRequest); i {
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
		file_BookingTables_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingClosing); i {
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
		file_BookingTables_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingDetail); i {
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
		file_BookingTables_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingFeedBack); i {
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
		file_BookingTables_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingResult); i {
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
		file_BookingTables_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingList); i {
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
			RawDescriptor: file_BookingTables_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BookingTables_proto_goTypes,
		DependencyIndexes: file_BookingTables_proto_depIdxs,
		MessageInfos:      file_BookingTables_proto_msgTypes,
	}.Build()
	File_BookingTables_proto = out.File
	file_BookingTables_proto_rawDesc = nil
	file_BookingTables_proto_goTypes = nil
	file_BookingTables_proto_depIdxs = nil
}
