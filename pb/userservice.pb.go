// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: pb/userservice.proto

package pb

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Fname     string  `protobuf:"bytes,2,opt,name=fname,proto3" json:"fname,omitempty"`
	City      string  `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Phone     string  `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Height    float64 `protobuf:"fixed64,5,opt,name=height,proto3" json:"height,omitempty"`
	IsMarried bool    `protobuf:"varint,6,opt,name=isMarried,proto3" json:"isMarried,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_userservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_pb_userservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_pb_userservice_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetFname() string {
	if x != nil {
		return x.Fname
	}
	return ""
}

func (x *User) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *User) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *User) GetHeight() float64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *User) GetIsMarried() bool {
	if x != nil {
		return x.IsMarried
	}
	return false
}

type GetUserByIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserByIDReq) Reset() {
	*x = GetUserByIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_userservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIDReq) ProtoMessage() {}

func (x *GetUserByIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_userservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIDReq.ProtoReflect.Descriptor instead.
func (*GetUserByIDReq) Descriptor() ([]byte, []int) {
	return file_pb_userservice_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserByIDReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetUserByIDRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserByIDRes) Reset() {
	*x = GetUserByIDRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_userservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByIDRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIDRes) ProtoMessage() {}

func (x *GetUserByIDRes) ProtoReflect() protoreflect.Message {
	mi := &file_pb_userservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIDRes.ProtoReflect.Descriptor instead.
func (*GetUserByIDRes) Descriptor() ([]byte, []int) {
	return file_pb_userservice_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserByIDRes) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetUserListByIDsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []uint64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GetUserListByIDsReq) Reset() {
	*x = GetUserListByIDsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_userservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserListByIDsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserListByIDsReq) ProtoMessage() {}

func (x *GetUserListByIDsReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_userservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserListByIDsReq.ProtoReflect.Descriptor instead.
func (*GetUserListByIDsReq) Descriptor() ([]byte, []int) {
	return file_pb_userservice_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserListByIDsReq) GetIds() []uint64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type GetUserListByIDsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserList []*User  `protobuf:"bytes,1,rep,name=userList,proto3" json:"userList,omitempty"`
	NotFound []uint64 `protobuf:"varint,2,rep,packed,name=notFound,proto3" json:"notFound,omitempty"`
}

func (x *GetUserListByIDsRes) Reset() {
	*x = GetUserListByIDsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_userservice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserListByIDsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserListByIDsRes) ProtoMessage() {}

func (x *GetUserListByIDsRes) ProtoReflect() protoreflect.Message {
	mi := &file_pb_userservice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserListByIDsRes.ProtoReflect.Descriptor instead.
func (*GetUserListByIDsRes) Descriptor() ([]byte, []int) {
	return file_pb_userservice_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserListByIDsRes) GetUserList() []*User {
	if x != nil {
		return x.UserList
	}
	return nil
}

func (x *GetUserListByIDsRes) GetNotFound() []uint64 {
	if x != nil {
		return x.NotFound
	}
	return nil
}

type SearchUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fname                  string  `protobuf:"bytes,1,opt,name=fname,proto3" json:"fname,omitempty"`
	City                   string  `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Phone                  string  `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Height                 float64 `protobuf:"fixed64,4,opt,name=height,proto3" json:"height,omitempty"`
	IsMarried              bool    `protobuf:"varint,5,opt,name=isMarried,proto3" json:"isMarried,omitempty"`
	IsMarriedFilterApplied bool    `protobuf:"varint,6,opt,name=isMarriedFilterApplied,proto3" json:"isMarriedFilterApplied,omitempty"`
}

func (x *SearchUserReq) Reset() {
	*x = SearchUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_userservice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUserReq) ProtoMessage() {}

func (x *SearchUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_userservice_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUserReq.ProtoReflect.Descriptor instead.
func (*SearchUserReq) Descriptor() ([]byte, []int) {
	return file_pb_userservice_proto_rawDescGZIP(), []int{5}
}

func (x *SearchUserReq) GetFname() string {
	if x != nil {
		return x.Fname
	}
	return ""
}

func (x *SearchUserReq) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *SearchUserReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *SearchUserReq) GetHeight() float64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *SearchUserReq) GetIsMarried() bool {
	if x != nil {
		return x.IsMarried
	}
	return false
}

func (x *SearchUserReq) GetIsMarriedFilterApplied() bool {
	if x != nil {
		return x.IsMarriedFilterApplied
	}
	return false
}

type SearchUserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserList []*User `protobuf:"bytes,1,rep,name=userList,proto3" json:"userList,omitempty"`
}

func (x *SearchUserRes) Reset() {
	*x = SearchUserRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_userservice_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUserRes) ProtoMessage() {}

func (x *SearchUserRes) ProtoReflect() protoreflect.Message {
	mi := &file_pb_userservice_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUserRes.ProtoReflect.Descriptor instead.
func (*SearchUserRes) Descriptor() ([]byte, []int) {
	return file_pb_userservice_proto_rawDescGZIP(), []int{6}
}

func (x *SearchUserRes) GetUserList() []*User {
	if x != nil {
		return x.UserList
	}
	return nil
}

var File_pb_userservice_proto protoreflect.FileDescriptor

var file_pb_userservice_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x8c, 0x01, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x66, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69,
	0x73, 0x4d, 0x61, 0x72, 0x72, 0x69, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x69, 0x73, 0x4d, 0x61, 0x72, 0x72, 0x69, 0x65, 0x64, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x12, 0x1c, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x27, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x73, 0x52,
	0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52,
	0x03, 0x69, 0x64, 0x73, 0x22, 0x57, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x73, 0x52, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x04, 0x52, 0x08, 0x6e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0xbd, 0x01,
	0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12,
	0x14, 0x0a, 0x05, 0x66, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x66, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x4d, 0x61, 0x72,
	0x72, 0x69, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x4d, 0x61,
	0x72, 0x72, 0x69, 0x65, 0x64, 0x12, 0x36, 0x0a, 0x16, 0x69, 0x73, 0x4d, 0x61, 0x72, 0x72, 0x69,
	0x65, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x69, 0x73, 0x4d, 0x61, 0x72, 0x72, 0x69, 0x65, 0x64,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x22, 0x35, 0x0a,
	0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x24,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x32, 0xc4, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x79, 0x49, 0x44, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x46, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44,
	0x73, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x73,
	0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_userservice_proto_rawDescOnce sync.Once
	file_pb_userservice_proto_rawDescData = file_pb_userservice_proto_rawDesc
)

func file_pb_userservice_proto_rawDescGZIP() []byte {
	file_pb_userservice_proto_rawDescOnce.Do(func() {
		file_pb_userservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_userservice_proto_rawDescData)
	})
	return file_pb_userservice_proto_rawDescData
}

var file_pb_userservice_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pb_userservice_proto_goTypes = []interface{}{
	(*User)(nil),                // 0: pb.User
	(*GetUserByIDReq)(nil),      // 1: pb.GetUserByIDReq
	(*GetUserByIDRes)(nil),      // 2: pb.GetUserByIDRes
	(*GetUserListByIDsReq)(nil), // 3: pb.GetUserListByIDsReq
	(*GetUserListByIDsRes)(nil), // 4: pb.GetUserListByIDsRes
	(*SearchUserReq)(nil),       // 5: pb.SearchUserReq
	(*SearchUserRes)(nil),       // 6: pb.SearchUserRes
}
var file_pb_userservice_proto_depIdxs = []int32{
	0, // 0: pb.GetUserByIDRes.user:type_name -> pb.User
	0, // 1: pb.GetUserListByIDsRes.userList:type_name -> pb.User
	0, // 2: pb.SearchUserRes.userList:type_name -> pb.User
	1, // 3: pb.UserService.GetUserByID:input_type -> pb.GetUserByIDReq
	3, // 4: pb.UserService.GetUserListByIDs:input_type -> pb.GetUserListByIDsReq
	5, // 5: pb.UserService.SearchUser:input_type -> pb.SearchUserReq
	2, // 6: pb.UserService.GetUserByID:output_type -> pb.GetUserByIDRes
	4, // 7: pb.UserService.GetUserListByIDs:output_type -> pb.GetUserListByIDsRes
	6, // 8: pb.UserService.SearchUser:output_type -> pb.SearchUserRes
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pb_userservice_proto_init() }
func file_pb_userservice_proto_init() {
	if File_pb_userservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_userservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_pb_userservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByIDReq); i {
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
		file_pb_userservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByIDRes); i {
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
		file_pb_userservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserListByIDsReq); i {
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
		file_pb_userservice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserListByIDsRes); i {
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
		file_pb_userservice_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchUserReq); i {
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
		file_pb_userservice_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchUserRes); i {
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
			RawDescriptor: file_pb_userservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_userservice_proto_goTypes,
		DependencyIndexes: file_pb_userservice_proto_depIdxs,
		MessageInfos:      file_pb_userservice_proto_msgTypes,
	}.Build()
	File_pb_userservice_proto = out.File
	file_pb_userservice_proto_rawDesc = nil
	file_pb_userservice_proto_goTypes = nil
	file_pb_userservice_proto_depIdxs = nil
}
