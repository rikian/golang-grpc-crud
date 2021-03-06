// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: grpc-proto/grpc.proto

package grpc_crud

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

// GRPC request
type NewUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *NewUser) Reset() {
	*x = NewUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewUser) ProtoMessage() {}

func (x *NewUser) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewUser.ProtoReflect.Descriptor instead.
func (*NewUser) Descriptor() ([]byte, []int) {
	return file_grpc_proto_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *NewUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *NewUser) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type GetAllUsers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Token    string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetAllUsers) Reset() {
	*x = GetAllUsers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUsers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUsers) ProtoMessage() {}

func (x *GetAllUsers) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUsers.ProtoReflect.Descriptor instead.
func (*GetAllUsers) Descriptor() ([]byte, []int) {
	return file_grpc_proto_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllUsers) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetAllUsers) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *GetAllUsers) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type NewDataUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldEmail    string `protobuf:"bytes,1,opt,name=oldEmail,proto3" json:"oldEmail,omitempty"`
	NewEmail    string `protobuf:"bytes,2,opt,name=newEmail,proto3" json:"newEmail,omitempty"`
	OldPassword string `protobuf:"bytes,3,opt,name=oldPassword,proto3" json:"oldPassword,omitempty"`
	NewPassword string `protobuf:"bytes,4,opt,name=newPassword,proto3" json:"newPassword,omitempty"`
	OldStatus   string `protobuf:"bytes,5,opt,name=oldStatus,proto3" json:"oldStatus,omitempty"`
	NewStatus   string `protobuf:"bytes,6,opt,name=newStatus,proto3" json:"newStatus,omitempty"`
}

func (x *NewDataUser) Reset() {
	*x = NewDataUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewDataUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewDataUser) ProtoMessage() {}

func (x *NewDataUser) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewDataUser.ProtoReflect.Descriptor instead.
func (*NewDataUser) Descriptor() ([]byte, []int) {
	return file_grpc_proto_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *NewDataUser) GetOldEmail() string {
	if x != nil {
		return x.OldEmail
	}
	return ""
}

func (x *NewDataUser) GetNewEmail() string {
	if x != nil {
		return x.NewEmail
	}
	return ""
}

func (x *NewDataUser) GetOldPassword() string {
	if x != nil {
		return x.OldPassword
	}
	return ""
}

func (x *NewDataUser) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

func (x *NewDataUser) GetOldStatus() string {
	if x != nil {
		return x.OldStatus
	}
	return ""
}

func (x *NewDataUser) GetNewStatus() string {
	if x != nil {
		return x.NewStatus
	}
	return ""
}

// GRPC response
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Status   string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_grpc_proto_msgTypes[3]
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
	return file_grpc_proto_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UserList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *UserList) Reset() {
	*x = UserList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserList) ProtoMessage() {}

func (x *UserList) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserList.ProtoReflect.Descriptor instead.
func (*UserList) Descriptor() ([]byte, []int) {
	return file_grpc_proto_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *UserList) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_grpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_grpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_grpc_proto_grpc_proto_rawDescGZIP(), []int{5}
}

func (x *Status) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *Status) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_grpc_proto_grpc_proto protoreflect.FileDescriptor

var file_grpc_proto_grpc_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x55, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xc5, 0x01, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x44,
	0x61, 0x74, 0x61, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x6c, 0x64, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x20, 0x0a, 0x0b, 0x6f, 0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x60, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x32, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x0a,
	0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x3a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0xa9, 0x02, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x75, 0x64, 0x12, 0x38,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77,
	0x55, 0x73, 0x65, 0x72, 0x1a, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x1a, 0x14, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65,
	0x77, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x55, 0x73, 0x65, 0x72,
	0x1a, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x42, 0x15, 0x5a,
	0x13, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x72, 0x75, 0x64, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x63, 0x72, 0x75, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_proto_grpc_proto_rawDescOnce sync.Once
	file_grpc_proto_grpc_proto_rawDescData = file_grpc_proto_grpc_proto_rawDesc
)

func file_grpc_proto_grpc_proto_rawDescGZIP() []byte {
	file_grpc_proto_grpc_proto_rawDescOnce.Do(func() {
		file_grpc_proto_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_grpc_proto_rawDescData)
	})
	return file_grpc_proto_grpc_proto_rawDescData
}

var file_grpc_proto_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_grpc_proto_grpc_proto_goTypes = []interface{}{
	(*NewUser)(nil),     // 0: grpc_proto.NewUser
	(*GetAllUsers)(nil), // 1: grpc_proto.GetAllUsers
	(*NewDataUser)(nil), // 2: grpc_proto.NewDataUser
	(*User)(nil),        // 3: grpc_proto.User
	(*UserList)(nil),    // 4: grpc_proto.UserList
	(*Status)(nil),      // 5: grpc_proto.Status
}
var file_grpc_proto_grpc_proto_depIdxs = []int32{
	3, // 0: grpc_proto.UserList.users:type_name -> grpc_proto.User
	0, // 1: grpc_proto.UserCrud.CreateNewUser:input_type -> grpc_proto.NewUser
	1, // 2: grpc_proto.UserCrud.GetUsers:input_type -> grpc_proto.GetAllUsers
	0, // 3: grpc_proto.UserCrud.GetUser:input_type -> grpc_proto.NewUser
	2, // 4: grpc_proto.UserCrud.UpdateUser:input_type -> grpc_proto.NewDataUser
	0, // 5: grpc_proto.UserCrud.DeleteUser:input_type -> grpc_proto.NewUser
	3, // 6: grpc_proto.UserCrud.CreateNewUser:output_type -> grpc_proto.User
	4, // 7: grpc_proto.UserCrud.GetUsers:output_type -> grpc_proto.UserList
	3, // 8: grpc_proto.UserCrud.GetUser:output_type -> grpc_proto.User
	3, // 9: grpc_proto.UserCrud.UpdateUser:output_type -> grpc_proto.User
	5, // 10: grpc_proto.UserCrud.DeleteUser:output_type -> grpc_proto.Status
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpc_proto_grpc_proto_init() }
func file_grpc_proto_grpc_proto_init() {
	if File_grpc_proto_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewUser); i {
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
		file_grpc_proto_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUsers); i {
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
		file_grpc_proto_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewDataUser); i {
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
		file_grpc_proto_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_grpc_proto_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserList); i {
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
		file_grpc_proto_grpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_grpc_proto_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_grpc_proto_goTypes,
		DependencyIndexes: file_grpc_proto_grpc_proto_depIdxs,
		MessageInfos:      file_grpc_proto_grpc_proto_msgTypes,
	}.Build()
	File_grpc_proto_grpc_proto = out.File
	file_grpc_proto_grpc_proto_rawDesc = nil
	file_grpc_proto_grpc_proto_goTypes = nil
	file_grpc_proto_grpc_proto_depIdxs = nil
}
