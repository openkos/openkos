// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: entry/entry.proto

package entry

import (
	kiae "github.com/kiaedev/kiae/api/kiae"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Entry_Scope int32

const (
	Entry_ALL     Entry_Scope = 0
	Entry_PARTIAL Entry_Scope = 1
)

// Enum value maps for Entry_Scope.
var (
	Entry_Scope_name = map[int32]string{
		0: "ALL",
		1: "PARTIAL",
	}
	Entry_Scope_value = map[string]int32{
		"ALL":     0,
		"PARTIAL": 1,
	}
)

func (x Entry_Scope) Enum() *Entry_Scope {
	p := new(Entry_Scope)
	*p = x
	return p
}

func (x Entry_Scope) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Entry_Scope) Descriptor() protoreflect.EnumDescriptor {
	return file_entry_entry_proto_enumTypes[0].Descriptor()
}

func (Entry_Scope) Type() protoreflect.EnumType {
	return &file_entry_entry_proto_enumTypes[0]
}

func (x Entry_Scope) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Entry_Scope.Descriptor instead.
func (Entry_Scope) EnumDescriptor() ([]byte, []int) {
	return file_entry_entry_proto_rawDescGZIP(), []int{3, 0}
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid string `protobuf:"bytes,1,opt,name=appid,proto3" json:"appid,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entry_entry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entry_entry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_entry_entry_proto_rawDescGZIP(), []int{0}
}

func (x *ListRequest) GetAppid() string {
	if x != nil {
		return x.Appid
	}
	return ""
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Entry `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total int64    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entry_entry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_entry_entry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_entry_entry_proto_rawDescGZIP(), []int{1}
}

func (x *ListResponse) GetItems() []*Entry {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload    *Entry                 `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entry_entry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entry_entry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_entry_entry_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRequest) GetPayload() *Entry {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *UpdateRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id,omitempty"` // @gotags: bson:"_id,omitempty"
	Appid     string                 `protobuf:"bytes,2,opt,name=appid,proto3" json:"appid,omitempty"`
	Gateway   string                 `protobuf:"bytes,4,opt,name=gateway,proto3" json:"gateway,omitempty"`
	Host      string                 `protobuf:"bytes,5,opt,name=host,proto3" json:"host,omitempty"`
	Scope     Entry_Scope            `protobuf:"varint,6,opt,name=scope,proto3,enum=entry.Entry_Scope" json:"scope,omitempty"`
	RouteIds  []string               `protobuf:"bytes,7,rep,name=route_ids,json=routeIds,proto3" json:"route_ids,omitempty"`
	Status    kiae.OpStatus          `protobuf:"varint,8,opt,name=status,proto3,enum=kiae.OpStatus" json:"status,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,101,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,103,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entry_entry_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_entry_entry_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_entry_entry_proto_rawDescGZIP(), []int{3}
}

func (x *Entry) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Entry) GetAppid() string {
	if x != nil {
		return x.Appid
	}
	return ""
}

func (x *Entry) GetGateway() string {
	if x != nil {
		return x.Gateway
	}
	return ""
}

func (x *Entry) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Entry) GetScope() Entry_Scope {
	if x != nil {
		return x.Scope
	}
	return Entry_ALL
}

func (x *Entry) GetRouteIds() []string {
	if x != nil {
		return x.RouteIds
	}
	return nil
}

func (x *Entry) GetStatus() kiae.OpStatus {
	if x != nil {
		return x.Status
	}
	return kiae.OpStatus(0)
}

func (x *Entry) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Entry) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_entry_entry_proto protoreflect.FileDescriptor

var file_entry_entry_proto_rawDesc = []byte{
	0x0a, 0x11, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x6b, 0x69, 0x61, 0x65, 0x2f, 0x6f,
	0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12,
	0x6b, 0x69, 0x61, 0x65, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x23, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x22, 0x48, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x22, 0x74, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x26, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0xdf, 0x02, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x64, 0x73, 0x12, 0x26, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x6b,
	0x69, 0x61, 0x65, 0x2e, 0x4f, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x67, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x1d, 0x0a, 0x05, 0x53, 0x63,
	0x6f, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x50, 0x41, 0x52, 0x54, 0x49, 0x41, 0x4c, 0x10, 0x01, 0x32, 0xc1, 0x03, 0x0a, 0x0c, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x12, 0x2e, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70,
	0x73, 0x2f, 0x7b, 0x61, 0x70, 0x70, 0x69, 0x64, 0x7d, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x4d, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x1a, 0x0c, 0x2e, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22,
	0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x7b, 0x61,
	0x70, 0x70, 0x69, 0x64, 0x7d, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x3a, 0x01, 0x2a,
	0x12, 0xaf, 0x01, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0c, 0x2e, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22,
	0x80, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x7a, 0x1a, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x7b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e,
	0x61, 0x70, 0x70, 0x69, 0x64, 0x7d, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x7b,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x69, 0x64, 0x7d, 0x3a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x5a, 0x3c, 0x32, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x70, 0x70, 0x73, 0x2f, 0x7b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x61, 0x70,
	0x70, 0x69, 0x64, 0x7d, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x69, 0x64, 0x7d, 0x3a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x59, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x6b,
	0x69, 0x61, 0x65, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x2a, 0x1e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x2f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x23, 0x5a,
	0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x69, 0x61, 0x65,
	0x64, 0x65, 0x76, 0x2f, 0x6b, 0x69, 0x61, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entry_entry_proto_rawDescOnce sync.Once
	file_entry_entry_proto_rawDescData = file_entry_entry_proto_rawDesc
)

func file_entry_entry_proto_rawDescGZIP() []byte {
	file_entry_entry_proto_rawDescOnce.Do(func() {
		file_entry_entry_proto_rawDescData = protoimpl.X.CompressGZIP(file_entry_entry_proto_rawDescData)
	})
	return file_entry_entry_proto_rawDescData
}

var file_entry_entry_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_entry_entry_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_entry_entry_proto_goTypes = []interface{}{
	(Entry_Scope)(0),              // 0: entry.Entry.Scope
	(*ListRequest)(nil),           // 1: entry.ListRequest
	(*ListResponse)(nil),          // 2: entry.ListResponse
	(*UpdateRequest)(nil),         // 3: entry.UpdateRequest
	(*Entry)(nil),                 // 4: entry.Entry
	(*fieldmaskpb.FieldMask)(nil), // 5: google.protobuf.FieldMask
	(kiae.OpStatus)(0),            // 6: kiae.OpStatus
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*kiae.IdRequest)(nil),        // 8: kiae.IdRequest
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_entry_entry_proto_depIdxs = []int32{
	4,  // 0: entry.ListResponse.items:type_name -> entry.Entry
	4,  // 1: entry.UpdateRequest.payload:type_name -> entry.Entry
	5,  // 2: entry.UpdateRequest.update_mask:type_name -> google.protobuf.FieldMask
	0,  // 3: entry.Entry.scope:type_name -> entry.Entry.Scope
	6,  // 4: entry.Entry.status:type_name -> kiae.OpStatus
	7,  // 5: entry.Entry.created_at:type_name -> google.protobuf.Timestamp
	7,  // 6: entry.Entry.updated_at:type_name -> google.protobuf.Timestamp
	1,  // 7: entry.EntryService.List:input_type -> entry.ListRequest
	4,  // 8: entry.EntryService.Create:input_type -> entry.Entry
	3,  // 9: entry.EntryService.Update:input_type -> entry.UpdateRequest
	8,  // 10: entry.EntryService.Delete:input_type -> kiae.IdRequest
	2,  // 11: entry.EntryService.List:output_type -> entry.ListResponse
	4,  // 12: entry.EntryService.Create:output_type -> entry.Entry
	4,  // 13: entry.EntryService.Update:output_type -> entry.Entry
	9,  // 14: entry.EntryService.Delete:output_type -> google.protobuf.Empty
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_entry_entry_proto_init() }
func file_entry_entry_proto_init() {
	if File_entry_entry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_entry_entry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_entry_entry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_entry_entry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_entry_entry_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry); i {
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
			RawDescriptor: file_entry_entry_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_entry_entry_proto_goTypes,
		DependencyIndexes: file_entry_entry_proto_depIdxs,
		EnumInfos:         file_entry_entry_proto_enumTypes,
		MessageInfos:      file_entry_entry_proto_msgTypes,
	}.Build()
	File_entry_entry_proto = out.File
	file_entry_entry_proto_rawDesc = nil
	file_entry_entry_proto_goTypes = nil
	file_entry_entry_proto_depIdxs = nil
}
