// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: middleware/middleware.proto

package middleware

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	kiae "github.com/kiaedev/kiae/api/kiae"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type Claim_Policy int32

const (
	Claim_REUSE  Claim_Policy = 0
	Claim_CREATE Claim_Policy = 1
)

// Enum value maps for Claim_Policy.
var (
	Claim_Policy_name = map[int32]string{
		0: "REUSE",
		1: "CREATE",
	}
	Claim_Policy_value = map[string]int32{
		"REUSE":  0,
		"CREATE": 1,
	}
)

func (x Claim_Policy) Enum() *Claim_Policy {
	p := new(Claim_Policy)
	*p = x
	return p
}

func (x Claim_Policy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Claim_Policy) Descriptor() protoreflect.EnumDescriptor {
	return file_middleware_middleware_proto_enumTypes[0].Descriptor()
}

func (Claim_Policy) Type() protoreflect.EnumType {
	return &file_middleware_middleware_proto_enumTypes[0]
}

func (x Claim_Policy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Claim_Policy.Descriptor instead.
func (Claim_Policy) EnumDescriptor() ([]byte, []int) {
	return file_middleware_middleware_proto_rawDescGZIP(), []int{4, 0}
}

type Claim_Status int32

const (
	Claim_UNBOUND Claim_Status = 0
	Claim_BOUND   Claim_Status = 1
)

// Enum value maps for Claim_Status.
var (
	Claim_Status_name = map[int32]string{
		0: "UNBOUND",
		1: "BOUND",
	}
	Claim_Status_value = map[string]int32{
		"UNBOUND": 0,
		"BOUND":   1,
	}
)

func (x Claim_Status) Enum() *Claim_Status {
	p := new(Claim_Status)
	*p = x
	return p
}

func (x Claim_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Claim_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_middleware_middleware_proto_enumTypes[1].Descriptor()
}

func (Claim_Status) Type() protoreflect.EnumType {
	return &file_middleware_middleware_proto_enumTypes[1]
}

func (x Claim_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Claim_Status.Descriptor instead.
func (Claim_Status) EnumDescriptor() ([]byte, []int) {
	return file_middleware_middleware_proto_rawDescGZIP(), []int{4, 1}
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_middleware_middleware_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_middleware_middleware_proto_msgTypes[0]
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
	return file_middleware_middleware_proto_rawDescGZIP(), []int{0}
}

func (x *ListRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Instance `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total int64       `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_middleware_middleware_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_middleware_middleware_proto_msgTypes[1]
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
	return file_middleware_middleware_proto_rawDescGZIP(), []int{1}
}

func (x *ListResponse) GetItems() []*Instance {
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

type ClaimsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *ClaimsRequest) Reset() {
	*x = ClaimsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_middleware_middleware_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimsRequest) ProtoMessage() {}

func (x *ClaimsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_middleware_middleware_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimsRequest.ProtoReflect.Descriptor instead.
func (*ClaimsRequest) Descriptor() ([]byte, []int) {
	return file_middleware_middleware_proto_rawDescGZIP(), []int{2}
}

func (x *ClaimsRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type ClaimsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Claim `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total int64    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ClaimsResponse) Reset() {
	*x = ClaimsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_middleware_middleware_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimsResponse) ProtoMessage() {}

func (x *ClaimsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_middleware_middleware_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimsResponse.ProtoReflect.Descriptor instead.
func (*ClaimsResponse) Descriptor() ([]byte, []int) {
	return file_middleware_middleware_proto_rawDescGZIP(), []int{3}
}

func (x *ClaimsResponse) GetItems() []*Claim {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ClaimsResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type Claim struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id,omitempty"` // @gotags: bson:"_id,omitempty"
	Appid     string                 `protobuf:"bytes,2,opt,name=appid,proto3" json:"appid,omitempty"`
	Type      string                 `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Name      string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Policy    Claim_Policy           `protobuf:"varint,5,opt,name=policy,proto3,enum=middleware.Claim_Policy" json:"policy,omitempty"`
	Instance  string                 `protobuf:"bytes,6,opt,name=instance,proto3" json:"instance,omitempty"`
	Status    Claim_Status           `protobuf:"varint,7,opt,name=status,proto3,enum=middleware.Claim_Status" json:"status,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,101,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,103,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Claim) Reset() {
	*x = Claim{}
	if protoimpl.UnsafeEnabled {
		mi := &file_middleware_middleware_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Claim) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Claim) ProtoMessage() {}

func (x *Claim) ProtoReflect() protoreflect.Message {
	mi := &file_middleware_middleware_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Claim.ProtoReflect.Descriptor instead.
func (*Claim) Descriptor() ([]byte, []int) {
	return file_middleware_middleware_proto_rawDescGZIP(), []int{4}
}

func (x *Claim) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Claim) GetAppid() string {
	if x != nil {
		return x.Appid
	}
	return ""
}

func (x *Claim) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Claim) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Claim) GetPolicy() Claim_Policy {
	if x != nil {
		return x.Policy
	}
	return Claim_REUSE
}

func (x *Claim) GetInstance() string {
	if x != nil {
		return x.Instance
	}
	return ""
}

func (x *Claim) GetStatus() Claim_Status {
	if x != nil {
		return x.Status
	}
	return Claim_UNBOUND
}

func (x *Claim) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Claim) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Instance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id,omitempty"` // @gotags: bson:"_id,omitempty"
	Name       string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Env        string                 `protobuf:"bytes,3,opt,name=env,proto3" json:"env,omitempty"`
	Type       string                 `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Properties map[string]string      `protobuf:"bytes,6,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Status     kiae.OpStatus          `protobuf:"varint,7,opt,name=status,proto3,enum=kiae.OpStatus" json:"status,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,101,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,103,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Instance) Reset() {
	*x = Instance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_middleware_middleware_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Instance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Instance) ProtoMessage() {}

func (x *Instance) ProtoReflect() protoreflect.Message {
	mi := &file_middleware_middleware_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Instance.ProtoReflect.Descriptor instead.
func (*Instance) Descriptor() ([]byte, []int) {
	return file_middleware_middleware_proto_rawDescGZIP(), []int{5}
}

func (x *Instance) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Instance) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Instance) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

func (x *Instance) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Instance) GetProperties() map[string]string {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *Instance) GetStatus() kiae.OpStatus {
	if x != nil {
		return x.Status
	}
	return kiae.OpStatus(0)
}

func (x *Instance) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Instance) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_middleware_middleware_proto protoreflect.FileDescriptor

var file_middleware_middleware_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2f, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x6b, 0x69, 0x61, 0x65, 0x2f, 0x6f, 0x70, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6b, 0x69, 0x61,
	0x65, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x50, 0x0a, 0x0c, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x23, 0x0a,
	0x0d, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x4f, 0x0a, 0x0e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x22, 0xd1, 0x03, 0x0a, 0x05, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70,
	0x70, 0x69, 0x64, 0x12, 0x55, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x41, 0xfa, 0x42, 0x3e, 0x72, 0x3c, 0x52, 0x05, 0x4d, 0x79, 0x53, 0x51, 0x4c, 0x52,
	0x07, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x44, 0x42, 0x52, 0x05, 0x52, 0x65, 0x64, 0x69, 0x73, 0x52,
	0x09, 0x4d, 0x65, 0x6d, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x52, 0x05, 0x4b, 0x61, 0x66, 0x6b,
	0x61, 0x52, 0x08, 0x52, 0x61, 0x62, 0x62, 0x69, 0x74, 0x4d, 0x51, 0x52, 0x07, 0x69, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30,
	0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x43, 0x6c, 0x61, 0x69,
	0x6d, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x65, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x67, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x1f, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x09,
	0x0a, 0x05, 0x52, 0x45, 0x55, 0x53, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45,
	0x41, 0x54, 0x45, 0x10, 0x01, 0x22, 0x20, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x42, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x42, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x22, 0xba, 0x03, 0x0a, 0x08, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x55, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x41, 0xfa, 0x42, 0x3e, 0x72, 0x3c, 0x52,
	0x05, 0x4d, 0x79, 0x53, 0x51, 0x4c, 0x52, 0x07, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x44, 0x42, 0x52,
	0x05, 0x52, 0x65, 0x64, 0x69, 0x73, 0x52, 0x09, 0x4d, 0x65, 0x6d, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x64, 0x52, 0x05, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x52, 0x08, 0x52, 0x61, 0x62, 0x62, 0x69, 0x74,
	0x4d, 0x51, 0x52, 0x07, 0x69, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x44, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x70, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x6b, 0x69, 0x61, 0x65, 0x2e, 0x4f,
	0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x65, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x67, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x1a, 0x3d, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x32, 0xb5, 0x06, 0x0a, 0x11, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x17, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x73, 0x12, 0x54, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x1a, 0x14, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18,
	0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x73, 0x12, 0x78, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x14, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x1a, 0x14, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x42,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3c, 0x3a, 0x01, 0x2a, 0x5a, 0x1d, 0x3a, 0x01, 0x2a, 0x32, 0x18,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x1a, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x73, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x53, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x6b,
	0x69, 0x61, 0x65, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x2a, 0x18, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x62, 0x0a, 0x06, 0x43, 0x6c, 0x61, 0x69, 0x6d,
	0x73, 0x12, 0x19, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x43,
	0x6c, 0x61, 0x69, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b,
	0x12, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x12, 0x59, 0x0a, 0x0b, 0x43,
	0x6c, 0x61, 0x69, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x1a, 0x11, 0x2e,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d,
	0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2f,
	0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x12, 0x83, 0x01, 0x0a, 0x0b, 0x43, 0x6c, 0x61, 0x69, 0x6d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x1a, 0x11, 0x2e, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x22, 0x4e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x48, 0x3a, 0x01, 0x2a, 0x5a, 0x23, 0x3a, 0x01, 0x2a, 0x32, 0x1e, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x1a, 0x1e, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5e, 0x0a, 0x0b,
	0x43, 0x6c, 0x61, 0x69, 0x6d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x6b, 0x69,
	0x61, 0x65, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x2a, 0x1e, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x28, 0x5a, 0x26,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x69, 0x61, 0x65, 0x64,
	0x65, 0x76, 0x2f, 0x6b, 0x69, 0x61, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_middleware_middleware_proto_rawDescOnce sync.Once
	file_middleware_middleware_proto_rawDescData = file_middleware_middleware_proto_rawDesc
)

func file_middleware_middleware_proto_rawDescGZIP() []byte {
	file_middleware_middleware_proto_rawDescOnce.Do(func() {
		file_middleware_middleware_proto_rawDescData = protoimpl.X.CompressGZIP(file_middleware_middleware_proto_rawDescData)
	})
	return file_middleware_middleware_proto_rawDescData
}

var file_middleware_middleware_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_middleware_middleware_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_middleware_middleware_proto_goTypes = []interface{}{
	(Claim_Policy)(0),             // 0: middleware.Claim.Policy
	(Claim_Status)(0),             // 1: middleware.Claim.Status
	(*ListRequest)(nil),           // 2: middleware.ListRequest
	(*ListResponse)(nil),          // 3: middleware.ListResponse
	(*ClaimsRequest)(nil),         // 4: middleware.ClaimsRequest
	(*ClaimsResponse)(nil),        // 5: middleware.ClaimsResponse
	(*Claim)(nil),                 // 6: middleware.Claim
	(*Instance)(nil),              // 7: middleware.Instance
	nil,                           // 8: middleware.Instance.PropertiesEntry
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
	(kiae.OpStatus)(0),            // 10: kiae.OpStatus
	(*kiae.IdRequest)(nil),        // 11: kiae.IdRequest
	(*emptypb.Empty)(nil),         // 12: google.protobuf.Empty
}
var file_middleware_middleware_proto_depIdxs = []int32{
	7,  // 0: middleware.ListResponse.items:type_name -> middleware.Instance
	6,  // 1: middleware.ClaimsResponse.items:type_name -> middleware.Claim
	0,  // 2: middleware.Claim.policy:type_name -> middleware.Claim.Policy
	1,  // 3: middleware.Claim.status:type_name -> middleware.Claim.Status
	9,  // 4: middleware.Claim.created_at:type_name -> google.protobuf.Timestamp
	9,  // 5: middleware.Claim.updated_at:type_name -> google.protobuf.Timestamp
	8,  // 6: middleware.Instance.properties:type_name -> middleware.Instance.PropertiesEntry
	10, // 7: middleware.Instance.status:type_name -> kiae.OpStatus
	9,  // 8: middleware.Instance.created_at:type_name -> google.protobuf.Timestamp
	9,  // 9: middleware.Instance.updated_at:type_name -> google.protobuf.Timestamp
	2,  // 10: middleware.MiddlewareService.List:input_type -> middleware.ListRequest
	7,  // 11: middleware.MiddlewareService.Create:input_type -> middleware.Instance
	7,  // 12: middleware.MiddlewareService.Update:input_type -> middleware.Instance
	11, // 13: middleware.MiddlewareService.Delete:input_type -> kiae.IdRequest
	4,  // 14: middleware.MiddlewareService.Claims:input_type -> middleware.ClaimsRequest
	6,  // 15: middleware.MiddlewareService.ClaimCreate:input_type -> middleware.Claim
	6,  // 16: middleware.MiddlewareService.ClaimUpdate:input_type -> middleware.Claim
	11, // 17: middleware.MiddlewareService.ClaimDelete:input_type -> kiae.IdRequest
	3,  // 18: middleware.MiddlewareService.List:output_type -> middleware.ListResponse
	7,  // 19: middleware.MiddlewareService.Create:output_type -> middleware.Instance
	7,  // 20: middleware.MiddlewareService.Update:output_type -> middleware.Instance
	12, // 21: middleware.MiddlewareService.Delete:output_type -> google.protobuf.Empty
	5,  // 22: middleware.MiddlewareService.Claims:output_type -> middleware.ClaimsResponse
	6,  // 23: middleware.MiddlewareService.ClaimCreate:output_type -> middleware.Claim
	6,  // 24: middleware.MiddlewareService.ClaimUpdate:output_type -> middleware.Claim
	12, // 25: middleware.MiddlewareService.ClaimDelete:output_type -> google.protobuf.Empty
	18, // [18:26] is the sub-list for method output_type
	10, // [10:18] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_middleware_middleware_proto_init() }
func file_middleware_middleware_proto_init() {
	if File_middleware_middleware_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_middleware_middleware_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_middleware_middleware_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_middleware_middleware_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimsRequest); i {
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
		file_middleware_middleware_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimsResponse); i {
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
		file_middleware_middleware_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Claim); i {
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
		file_middleware_middleware_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Instance); i {
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
			RawDescriptor: file_middleware_middleware_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_middleware_middleware_proto_goTypes,
		DependencyIndexes: file_middleware_middleware_proto_depIdxs,
		EnumInfos:         file_middleware_middleware_proto_enumTypes,
		MessageInfos:      file_middleware_middleware_proto_msgTypes,
	}.Build()
	File_middleware_middleware_proto = out.File
	file_middleware_middleware_proto_rawDesc = nil
	file_middleware_middleware_proto_goTypes = nil
	file_middleware_middleware_proto_depIdxs = nil
}
