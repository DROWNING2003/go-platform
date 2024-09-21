// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.24.4
// source: api/configure/configure/configure_configure.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// 查询指定模板
type GetConfigureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerId uint32 `protobuf:"varint,1,opt,name=serverId,proto3" json:"serverId,omitempty"`
	EnvId    uint32 `protobuf:"varint,2,opt,name=envId,proto3" json:"envId,omitempty"`
}

func (x *GetConfigureRequest) Reset() {
	*x = GetConfigureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_configure_configure_configure_configure_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigureRequest) ProtoMessage() {}

func (x *GetConfigureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_configure_configure_configure_configure_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigureRequest.ProtoReflect.Descriptor instead.
func (*GetConfigureRequest) Descriptor() ([]byte, []int) {
	return file_api_configure_configure_configure_configure_proto_rawDescGZIP(), []int{0}
}

func (x *GetConfigureRequest) GetServerId() uint32 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *GetConfigureRequest) GetEnvId() uint32 {
	if x != nil {
		return x.EnvId
	}
	return 0
}

type GetConfigureReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ServerId    uint32 `protobuf:"varint,2,opt,name=serverId,proto3" json:"serverId,omitempty"`
	EnvId       uint32 `protobuf:"varint,3,opt,name=envId,proto3" json:"envId,omitempty"`
	Content     string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Version     string `protobuf:"bytes,6,opt,name=version,proto3" json:"version,omitempty"`
	Format      string `protobuf:"bytes,7,opt,name=format,proto3" json:"format,omitempty"`
	CreatedAt   uint32 `protobuf:"varint,8,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *GetConfigureReply) Reset() {
	*x = GetConfigureReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_configure_configure_configure_configure_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigureReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigureReply) ProtoMessage() {}

func (x *GetConfigureReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_configure_configure_configure_configure_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigureReply.ProtoReflect.Descriptor instead.
func (*GetConfigureReply) Descriptor() ([]byte, []int) {
	return file_api_configure_configure_configure_configure_proto_rawDescGZIP(), []int{1}
}

func (x *GetConfigureReply) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetConfigureReply) GetServerId() uint32 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *GetConfigureReply) GetEnvId() uint32 {
	if x != nil {
		return x.EnvId
	}
	return 0
}

func (x *GetConfigureReply) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *GetConfigureReply) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetConfigureReply) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *GetConfigureReply) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *GetConfigureReply) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

// 对比配置
type CompareConfigureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerId uint32 `protobuf:"varint,1,opt,name=serverId,proto3" json:"serverId,omitempty"`
	EnvId    uint32 `protobuf:"varint,2,opt,name=envId,proto3" json:"envId,omitempty"`
}

func (x *CompareConfigureRequest) Reset() {
	*x = CompareConfigureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_configure_configure_configure_configure_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompareConfigureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompareConfigureRequest) ProtoMessage() {}

func (x *CompareConfigureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_configure_configure_configure_configure_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompareConfigureRequest.ProtoReflect.Descriptor instead.
func (*CompareConfigureRequest) Descriptor() ([]byte, []int) {
	return file_api_configure_configure_configure_configure_proto_rawDescGZIP(), []int{2}
}

func (x *CompareConfigureRequest) GetServerId() uint32 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *CompareConfigureRequest) GetEnvId() uint32 {
	if x != nil {
		return x.EnvId
	}
	return 0
}

// 对比配置详情
type Compare struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Key  string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Old  string `protobuf:"bytes,3,opt,name=old,proto3" json:"old,omitempty"`
	Cur  string `protobuf:"bytes,4,opt,name=cur,proto3" json:"cur,omitempty"`
}

func (x *Compare) Reset() {
	*x = Compare{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_configure_configure_configure_configure_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Compare) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Compare) ProtoMessage() {}

func (x *Compare) ProtoReflect() protoreflect.Message {
	mi := &file_api_configure_configure_configure_configure_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Compare.ProtoReflect.Descriptor instead.
func (*Compare) Descriptor() ([]byte, []int) {
	return file_api_configure_configure_configure_configure_proto_rawDescGZIP(), []int{3}
}

func (x *Compare) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Compare) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Compare) GetOld() string {
	if x != nil {
		return x.Old
	}
	return ""
}

func (x *Compare) GetCur() string {
	if x != nil {
		return x.Cur
	}
	return ""
}

type CompareConfigureReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Compare `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *CompareConfigureReply) Reset() {
	*x = CompareConfigureReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_configure_configure_configure_configure_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompareConfigureReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompareConfigureReply) ProtoMessage() {}

func (x *CompareConfigureReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_configure_configure_configure_configure_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompareConfigureReply.ProtoReflect.Descriptor instead.
func (*CompareConfigureReply) Descriptor() ([]byte, []int) {
	return file_api_configure_configure_configure_configure_proto_rawDescGZIP(), []int{4}
}

func (x *CompareConfigureReply) GetList() []*Compare {
	if x != nil {
		return x.List
	}
	return nil
}

// 修改配置
type UpdateConfigureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerId    uint32 `protobuf:"varint,1,opt,name=serverId,proto3" json:"serverId,omitempty"`
	EnvId       uint32 `protobuf:"varint,2,opt,name=envId,proto3" json:"envId,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *UpdateConfigureRequest) Reset() {
	*x = UpdateConfigureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_configure_configure_configure_configure_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateConfigureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateConfigureRequest) ProtoMessage() {}

func (x *UpdateConfigureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_configure_configure_configure_configure_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateConfigureRequest.ProtoReflect.Descriptor instead.
func (*UpdateConfigureRequest) Descriptor() ([]byte, []int) {
	return file_api_configure_configure_configure_configure_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateConfigureRequest) GetServerId() uint32 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *UpdateConfigureRequest) GetEnvId() uint32 {
	if x != nil {
		return x.EnvId
	}
	return 0
}

func (x *UpdateConfigureRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type UpdateConfigureReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateConfigureReply) Reset() {
	*x = UpdateConfigureReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_configure_configure_configure_configure_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateConfigureReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateConfigureReply) ProtoMessage() {}

func (x *UpdateConfigureReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_configure_configure_configure_configure_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateConfigureReply.ProtoReflect.Descriptor instead.
func (*UpdateConfigureReply) Descriptor() ([]byte, []int) {
	return file_api_configure_configure_configure_configure_proto_rawDescGZIP(), []int{6}
}

// 监听获取配置
type WatchConfigureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server string `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *WatchConfigureRequest) Reset() {
	*x = WatchConfigureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_configure_configure_configure_configure_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchConfigureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchConfigureRequest) ProtoMessage() {}

func (x *WatchConfigureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_configure_configure_configure_configure_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchConfigureRequest.ProtoReflect.Descriptor instead.
func (*WatchConfigureRequest) Descriptor() ([]byte, []int) {
	return file_api_configure_configure_configure_configure_proto_rawDescGZIP(), []int{7}
}

func (x *WatchConfigureRequest) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

func (x *WatchConfigureRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type WatchConfigureReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Format  string `protobuf:"bytes,2,opt,name=format,proto3" json:"format,omitempty"`
}

func (x *WatchConfigureReply) Reset() {
	*x = WatchConfigureReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_configure_configure_configure_configure_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchConfigureReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchConfigureReply) ProtoMessage() {}

func (x *WatchConfigureReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_configure_configure_configure_configure_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchConfigureReply.ProtoReflect.Descriptor instead.
func (*WatchConfigureReply) Descriptor() ([]byte, []int) {
	return file_api_configure_configure_configure_configure_proto_rawDescGZIP(), []int{8}
}

func (x *WatchConfigureReply) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *WatchConfigureReply) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

var File_api_configure_configure_configure_configure_proto protoreflect.FileDescriptor

var file_api_configure_configure_configure_configure_proto_rawDesc = []byte{
	0x0a, 0x31, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x24, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x59, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x2a, 0x02, 0x20, 0x00, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x05, 0x65, 0x6e, 0x76, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x05, 0x65, 0x6e, 0x76, 0x49, 0x64, 0x22, 0xe1, 0x01,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6e, 0x76, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x65, 0x6e, 0x76, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x5d, 0x0a, 0x17, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x08,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x05, 0x65, 0x6e, 0x76, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x05, 0x65, 0x6e, 0x76, 0x49, 0x64,
	0x22, 0x53, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6f, 0x6c, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x75, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x63, 0x75, 0x72, 0x22, 0x5a, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x41,
	0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x22, 0x89, 0x01, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x08,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x05, 0x65, 0x6e, 0x76, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x05, 0x65, 0x6e, 0x76, 0x49, 0x64,
	0x12, 0x2b, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x40,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x16, 0x0a,
	0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x57, 0x0a, 0x15, 0x57, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x1d, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x47,
	0x0a, 0x13, 0x57, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x42, 0x3b, 0x0a, 0x24, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x42,
	0x08, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x07, 0x2e, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_configure_configure_configure_configure_proto_rawDescOnce sync.Once
	file_api_configure_configure_configure_configure_proto_rawDescData = file_api_configure_configure_configure_configure_proto_rawDesc
)

func file_api_configure_configure_configure_configure_proto_rawDescGZIP() []byte {
	file_api_configure_configure_configure_configure_proto_rawDescOnce.Do(func() {
		file_api_configure_configure_configure_configure_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_configure_configure_configure_configure_proto_rawDescData)
	})
	return file_api_configure_configure_configure_configure_proto_rawDescData
}

var file_api_configure_configure_configure_configure_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_configure_configure_configure_configure_proto_goTypes = []interface{}{
	(*GetConfigureRequest)(nil),     // 0: configure.api.configure.configure.v1.GetConfigureRequest
	(*GetConfigureReply)(nil),       // 1: configure.api.configure.configure.v1.GetConfigureReply
	(*CompareConfigureRequest)(nil), // 2: configure.api.configure.configure.v1.CompareConfigureRequest
	(*Compare)(nil),                 // 3: configure.api.configure.configure.v1.Compare
	(*CompareConfigureReply)(nil),   // 4: configure.api.configure.configure.v1.CompareConfigureReply
	(*UpdateConfigureRequest)(nil),  // 5: configure.api.configure.configure.v1.UpdateConfigureRequest
	(*UpdateConfigureReply)(nil),    // 6: configure.api.configure.configure.v1.UpdateConfigureReply
	(*WatchConfigureRequest)(nil),   // 7: configure.api.configure.configure.v1.WatchConfigureRequest
	(*WatchConfigureReply)(nil),     // 8: configure.api.configure.configure.v1.WatchConfigureReply
}
var file_api_configure_configure_configure_configure_proto_depIdxs = []int32{
	3, // 0: configure.api.configure.configure.v1.CompareConfigureReply.list:type_name -> configure.api.configure.configure.v1.Compare
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_configure_configure_configure_configure_proto_init() }
func file_api_configure_configure_configure_configure_proto_init() {
	if File_api_configure_configure_configure_configure_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_configure_configure_configure_configure_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigureRequest); i {
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
		file_api_configure_configure_configure_configure_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigureReply); i {
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
		file_api_configure_configure_configure_configure_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompareConfigureRequest); i {
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
		file_api_configure_configure_configure_configure_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Compare); i {
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
		file_api_configure_configure_configure_configure_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompareConfigureReply); i {
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
		file_api_configure_configure_configure_configure_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateConfigureRequest); i {
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
		file_api_configure_configure_configure_configure_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateConfigureReply); i {
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
		file_api_configure_configure_configure_configure_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchConfigureRequest); i {
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
		file_api_configure_configure_configure_configure_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchConfigureReply); i {
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
			RawDescriptor: file_api_configure_configure_configure_configure_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_configure_configure_configure_configure_proto_goTypes,
		DependencyIndexes: file_api_configure_configure_configure_configure_proto_depIdxs,
		MessageInfos:      file_api_configure_configure_configure_configure_proto_msgTypes,
	}.Build()
	File_api_configure_configure_configure_configure_proto = out.File
	file_api_configure_configure_configure_configure_proto_rawDesc = nil
	file_api_configure_configure_configure_configure_proto_goTypes = nil
	file_api_configure_configure_configure_configure_proto_depIdxs = nil
}
