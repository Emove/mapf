// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: warehouse/v1/warehouse.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	v1 "mapf/api/common/data/v1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Symbols defined in public import of common/data/v1/data.proto.

type EnableStatus = v1.EnableStatus

const EnableStatus_DISABLE = v1.EnableStatus_DISABLE
const EnableStatus_ENABLE = v1.EnableStatus_ENABLE

var EnableStatus_name = v1.EnableStatus_name
var EnableStatus_value = v1.EnableStatus_value

type DefaultStatus = v1.DefaultStatus

const DefaultStatus_NOT_DEFAULT = v1.DefaultStatus_NOT_DEFAULT
const DefaultStatus_IS_DEFAULT = v1.DefaultStatus_IS_DEFAULT

var DefaultStatus_name = v1.DefaultStatus_name
var DefaultStatus_value = v1.DefaultStatus_value

type ValidStatus = v1.ValidStatus

const ValidStatus_INVALID = v1.ValidStatus_INVALID
const ValidStatus_VALID = v1.ValidStatus_VALID

var ValidStatus_name = v1.ValidStatus_name
var ValidStatus_value = v1.ValidStatus_value

type CreateWarehouseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string                   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	NodeTypes []*CreateNodeTypeRequest `protobuf:"bytes,2,rep,name=node_types,json=nodeTypes,proto3" json:"node_types,omitempty"`
}

func (x *CreateWarehouseRequest) Reset() {
	*x = CreateWarehouseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_v1_warehouse_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWarehouseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWarehouseRequest) ProtoMessage() {}

func (x *CreateWarehouseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_v1_warehouse_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWarehouseRequest.ProtoReflect.Descriptor instead.
func (*CreateWarehouseRequest) Descriptor() ([]byte, []int) {
	return file_warehouse_v1_warehouse_proto_rawDescGZIP(), []int{0}
}

func (x *CreateWarehouseRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateWarehouseRequest) GetNodeTypes() []*CreateNodeTypeRequest {
	if x != nil {
		return x.NodeTypes
	}
	return nil
}

type CreateWarehouseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Warehouse *Warehouse  `protobuf:"bytes,1,opt,name=warehouse,proto3" json:"warehouse,omitempty"`
	NodeTypes []*NodeType `protobuf:"bytes,2,rep,name=node_types,json=nodeTypes,proto3" json:"node_types,omitempty"`
}

func (x *CreateWarehouseResponse) Reset() {
	*x = CreateWarehouseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_v1_warehouse_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWarehouseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWarehouseResponse) ProtoMessage() {}

func (x *CreateWarehouseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_v1_warehouse_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWarehouseResponse.ProtoReflect.Descriptor instead.
func (*CreateWarehouseResponse) Descriptor() ([]byte, []int) {
	return file_warehouse_v1_warehouse_proto_rawDescGZIP(), []int{1}
}

func (x *CreateWarehouseResponse) GetWarehouse() *Warehouse {
	if x != nil {
		return x.Warehouse
	}
	return nil
}

func (x *CreateWarehouseResponse) GetNodeTypes() []*NodeType {
	if x != nil {
		return x.NodeTypes
	}
	return nil
}

type GetWarehouseByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetWarehouseByNameRequest) Reset() {
	*x = GetWarehouseByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_v1_warehouse_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWarehouseByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWarehouseByNameRequest) ProtoMessage() {}

func (x *GetWarehouseByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_v1_warehouse_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWarehouseByNameRequest.ProtoReflect.Descriptor instead.
func (*GetWarehouseByNameRequest) Descriptor() ([]byte, []int) {
	return file_warehouse_v1_warehouse_proto_rawDescGZIP(), []int{2}
}

func (x *GetWarehouseByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetWarehouseByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetWarehouseByIdRequest) Reset() {
	*x = GetWarehouseByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_v1_warehouse_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWarehouseByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWarehouseByIdRequest) ProtoMessage() {}

func (x *GetWarehouseByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_v1_warehouse_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWarehouseByIdRequest.ProtoReflect.Descriptor instead.
func (*GetWarehouseByIdRequest) Descriptor() ([]byte, []int) {
	return file_warehouse_v1_warehouse_proto_rawDescGZIP(), []int{3}
}

func (x *GetWarehouseByIdRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetWarehouseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Warehouse *Warehouse `protobuf:"bytes,1,opt,name=warehouse,proto3" json:"warehouse,omitempty"`
}

func (x *GetWarehouseResponse) Reset() {
	*x = GetWarehouseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_v1_warehouse_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWarehouseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWarehouseResponse) ProtoMessage() {}

func (x *GetWarehouseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_v1_warehouse_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWarehouseResponse.ProtoReflect.Descriptor instead.
func (*GetWarehouseResponse) Descriptor() ([]byte, []int) {
	return file_warehouse_v1_warehouse_proto_rawDescGZIP(), []int{4}
}

func (x *GetWarehouseResponse) GetWarehouse() *Warehouse {
	if x != nil {
		return x.Warehouse
	}
	return nil
}

type CreateNodeTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateNodeTypeRequest) Reset() {
	*x = CreateNodeTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_v1_warehouse_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNodeTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNodeTypeRequest) ProtoMessage() {}

func (x *CreateNodeTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_v1_warehouse_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNodeTypeRequest.ProtoReflect.Descriptor instead.
func (*CreateNodeTypeRequest) Descriptor() ([]byte, []int) {
	return file_warehouse_v1_warehouse_proto_rawDescGZIP(), []int{5}
}

func (x *CreateNodeTypeRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateNodeTypeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Warehouse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint32          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status v1.EnableStatus `protobuf:"varint,3,opt,name=status,proto3,enum=common.data.v1.EnableStatus" json:"status,omitempty"`
}

func (x *Warehouse) Reset() {
	*x = Warehouse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_v1_warehouse_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Warehouse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Warehouse) ProtoMessage() {}

func (x *Warehouse) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_v1_warehouse_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Warehouse.ProtoReflect.Descriptor instead.
func (*Warehouse) Descriptor() ([]byte, []int) {
	return file_warehouse_v1_warehouse_proto_rawDescGZIP(), []int{6}
}

func (x *Warehouse) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Warehouse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Warehouse) GetStatus() v1.EnableStatus {
	if x != nil {
		return x.Status
	}
	return v1.EnableStatus(0)
}

type NodeType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	WarehouseId uint32            `protobuf:"varint,2,opt,name=warehouse_id,json=warehouseId,proto3" json:"warehouse_id,omitempty"`
	Code        string            `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Name        string            `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	IsDefault   *v1.DefaultStatus `protobuf:"varint,5,opt,name=is_default,json=isDefault,proto3,enum=common.data.v1.DefaultStatus,oneof" json:"is_default,omitempty"`
}

func (x *NodeType) Reset() {
	*x = NodeType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_v1_warehouse_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeType) ProtoMessage() {}

func (x *NodeType) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_v1_warehouse_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeType.ProtoReflect.Descriptor instead.
func (*NodeType) Descriptor() ([]byte, []int) {
	return file_warehouse_v1_warehouse_proto_rawDescGZIP(), []int{7}
}

func (x *NodeType) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NodeType) GetWarehouseId() uint32 {
	if x != nil {
		return x.WarehouseId
	}
	return 0
}

func (x *NodeType) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *NodeType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NodeType) GetIsDefault() v1.DefaultStatus {
	if x != nil && x.IsDefault != nil {
		return *x.IsDefault
	}
	return v1.DefaultStatus(0)
}

var File_warehouse_v1_warehouse_proto protoreflect.FileDescriptor

var file_warehouse_v1_warehouse_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x77,
	0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b,
	0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x03, 0x18,
	0x40, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x77, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x22, 0x87, 0x01, 0x0a, 0x17,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x77, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x77, 0x61, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x52, 0x09, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x35,
	0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x6e, 0x6f, 0x64, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x22, 0x3a, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x57, 0x61, 0x72, 0x65,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x03, 0x18, 0x40, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x32, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x28,
	0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4d, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x57, 0x61, 0x72, 0x65,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a,
	0x09, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x09, 0x77, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x22, 0x55, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f,
	0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06,
	0x72, 0x04, 0x10, 0x03, 0x18, 0x40, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72,
	0x04, 0x10, 0x03, 0x18, 0x40, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x65, 0x0a, 0x09, 0x57,
	0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0xb7, 0x01, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x69, 0x73,
	0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52,
	0x09, 0x69, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x32, 0x82, 0x03, 0x0a,
	0x10, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x75, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x12, 0x24, 0x2e, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x77, 0x61, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x77,
	0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x7c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x57,
	0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27,
	0x2e, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2f,
	0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x79, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x57, 0x61, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x25, 0x2e, 0x77, 0x61, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x61, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f,
	0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2f, 0x69, 0x64, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x42, 0x1a, 0x5a, 0x18, 0x6d, 0x61, 0x70, 0x66, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x50, 0x02, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_warehouse_v1_warehouse_proto_rawDescOnce sync.Once
	file_warehouse_v1_warehouse_proto_rawDescData = file_warehouse_v1_warehouse_proto_rawDesc
)

func file_warehouse_v1_warehouse_proto_rawDescGZIP() []byte {
	file_warehouse_v1_warehouse_proto_rawDescOnce.Do(func() {
		file_warehouse_v1_warehouse_proto_rawDescData = protoimpl.X.CompressGZIP(file_warehouse_v1_warehouse_proto_rawDescData)
	})
	return file_warehouse_v1_warehouse_proto_rawDescData
}

var file_warehouse_v1_warehouse_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_warehouse_v1_warehouse_proto_goTypes = []interface{}{
	(*CreateWarehouseRequest)(nil),    // 0: warehouse.v1.CreateWarehouseRequest
	(*CreateWarehouseResponse)(nil),   // 1: warehouse.v1.CreateWarehouseResponse
	(*GetWarehouseByNameRequest)(nil), // 2: warehouse.v1.GetWarehouseByNameRequest
	(*GetWarehouseByIdRequest)(nil),   // 3: warehouse.v1.GetWarehouseByIdRequest
	(*GetWarehouseResponse)(nil),      // 4: warehouse.v1.GetWarehouseResponse
	(*CreateNodeTypeRequest)(nil),     // 5: warehouse.v1.CreateNodeTypeRequest
	(*Warehouse)(nil),                 // 6: warehouse.v1.Warehouse
	(*NodeType)(nil),                  // 7: warehouse.v1.NodeType
	(v1.EnableStatus)(0),              // 8: common.data.v1.EnableStatus
	(v1.DefaultStatus)(0),             // 9: common.data.v1.DefaultStatus
}
var file_warehouse_v1_warehouse_proto_depIdxs = []int32{
	5, // 0: warehouse.v1.CreateWarehouseRequest.node_types:type_name -> warehouse.v1.CreateNodeTypeRequest
	6, // 1: warehouse.v1.CreateWarehouseResponse.warehouse:type_name -> warehouse.v1.Warehouse
	7, // 2: warehouse.v1.CreateWarehouseResponse.node_types:type_name -> warehouse.v1.NodeType
	6, // 3: warehouse.v1.GetWarehouseResponse.warehouse:type_name -> warehouse.v1.Warehouse
	8, // 4: warehouse.v1.Warehouse.status:type_name -> common.data.v1.EnableStatus
	9, // 5: warehouse.v1.NodeType.is_default:type_name -> common.data.v1.DefaultStatus
	0, // 6: warehouse.v1.WarehouseService.CreateWarehouse:input_type -> warehouse.v1.CreateWarehouseRequest
	2, // 7: warehouse.v1.WarehouseService.GetWarehouseByName:input_type -> warehouse.v1.GetWarehouseByNameRequest
	3, // 8: warehouse.v1.WarehouseService.GetWarehouseById:input_type -> warehouse.v1.GetWarehouseByIdRequest
	1, // 9: warehouse.v1.WarehouseService.CreateWarehouse:output_type -> warehouse.v1.CreateWarehouseResponse
	4, // 10: warehouse.v1.WarehouseService.GetWarehouseByName:output_type -> warehouse.v1.GetWarehouseResponse
	4, // 11: warehouse.v1.WarehouseService.GetWarehouseById:output_type -> warehouse.v1.GetWarehouseResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_warehouse_v1_warehouse_proto_init() }
func file_warehouse_v1_warehouse_proto_init() {
	if File_warehouse_v1_warehouse_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_warehouse_v1_warehouse_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWarehouseRequest); i {
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
		file_warehouse_v1_warehouse_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWarehouseResponse); i {
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
		file_warehouse_v1_warehouse_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWarehouseByNameRequest); i {
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
		file_warehouse_v1_warehouse_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWarehouseByIdRequest); i {
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
		file_warehouse_v1_warehouse_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWarehouseResponse); i {
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
		file_warehouse_v1_warehouse_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNodeTypeRequest); i {
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
		file_warehouse_v1_warehouse_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Warehouse); i {
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
		file_warehouse_v1_warehouse_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeType); i {
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
	file_warehouse_v1_warehouse_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_warehouse_v1_warehouse_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_warehouse_v1_warehouse_proto_goTypes,
		DependencyIndexes: file_warehouse_v1_warehouse_proto_depIdxs,
		MessageInfos:      file_warehouse_v1_warehouse_proto_msgTypes,
	}.Build()
	File_warehouse_v1_warehouse_proto = out.File
	file_warehouse_v1_warehouse_proto_rawDesc = nil
	file_warehouse_v1_warehouse_proto_goTypes = nil
	file_warehouse_v1_warehouse_proto_depIdxs = nil
}
