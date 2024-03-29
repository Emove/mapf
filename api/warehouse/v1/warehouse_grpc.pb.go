// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: warehouse/v1/warehouse.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	v1 "mapf/api/common/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	WarehouseService_CreateWarehouse_FullMethodName           = "/warehouse.v1.WarehouseService/CreateWarehouse"
	WarehouseService_GetWarehouseByName_FullMethodName        = "/warehouse.v1.WarehouseService/GetWarehouseByName"
	WarehouseService_GetWarehouseById_FullMethodName          = "/warehouse.v1.WarehouseService/GetWarehouseById"
	WarehouseService_UpdateWarehouseById_FullMethodName       = "/warehouse.v1.WarehouseService/UpdateWarehouseById"
	WarehouseService_UpdateWarehouseStatusById_FullMethodName = "/warehouse.v1.WarehouseService/UpdateWarehouseStatusById"
	WarehouseService_DeleteWarehouseById_FullMethodName       = "/warehouse.v1.WarehouseService/DeleteWarehouseById"
	WarehouseService_CreateNodeType_FullMethodName            = "/warehouse.v1.WarehouseService/CreateNodeType"
	WarehouseService_GetNodeTypeById_FullMethodName           = "/warehouse.v1.WarehouseService/GetNodeTypeById"
	WarehouseService_UpdateNodeTypeById_FullMethodName        = "/warehouse.v1.WarehouseService/UpdateNodeTypeById"
	WarehouseService_DeleteNodeTypeById_FullMethodName        = "/warehouse.v1.WarehouseService/DeleteNodeTypeById"
	WarehouseService_CreateNodes_FullMethodName               = "/warehouse.v1.WarehouseService/CreateNodes"
	WarehouseService_GetNodeById_FullMethodName               = "/warehouse.v1.WarehouseService/GetNodeById"
	WarehouseService_GetNodesByWarehouseId_FullMethodName     = "/warehouse.v1.WarehouseService/GetNodesByWarehouseId"
	WarehouseService_CreateNodeConfigItem_FullMethodName      = "/warehouse.v1.WarehouseService/CreateNodeConfigItem"
	WarehouseService_GetNodeConfigItemById_FullMethodName     = "/warehouse.v1.WarehouseService/GetNodeConfigItemById"
	WarehouseService_SelectNodeConfigItems_FullMethodName     = "/warehouse.v1.WarehouseService/SelectNodeConfigItems"
	WarehouseService_UpdateNodeConfigItem_FullMethodName      = "/warehouse.v1.WarehouseService/UpdateNodeConfigItem"
	WarehouseService_CreateAffixNode_FullMethodName           = "/warehouse.v1.WarehouseService/CreateAffixNode"
	WarehouseService_GetAffixNodeById_FullMethodName          = "/warehouse.v1.WarehouseService/GetAffixNodeById"
	WarehouseService_UpdateAffixNode_FullMethodName           = "/warehouse.v1.WarehouseService/UpdateAffixNode"
)

// WarehouseServiceClient is the client API for WarehouseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WarehouseServiceClient interface {
	// -------------------------------------------------  Warehouse  -----------------------------------------------------
	CreateWarehouse(ctx context.Context, in *CreateWarehouseRequest, opts ...grpc.CallOption) (*SimpleWarehouseResponse, error)
	GetWarehouseByName(ctx context.Context, in *GetWarehouseByNameRequest, opts ...grpc.CallOption) (*SimpleWarehouseResponse, error)
	GetWarehouseById(ctx context.Context, in *GetWarehouseByIdRequest, opts ...grpc.CallOption) (*SimpleWarehouseResponse, error)
	UpdateWarehouseById(ctx context.Context, in *UpdateWarehouseByIdRequest, opts ...grpc.CallOption) (*v1.SimpleResponse, error)
	UpdateWarehouseStatusById(ctx context.Context, in *UpdateWarehouseStatusByIdRequest, opts ...grpc.CallOption) (*v1.SimpleResponse, error)
	DeleteWarehouseById(ctx context.Context, in *DeleteWarehouseByIdRequest, opts ...grpc.CallOption) (*v1.SimpleResponse, error)
	// -------------------------------------------------  NodeType  ------------------------------------------------------
	CreateNodeType(ctx context.Context, in *CreateNodeTypeRequest, opts ...grpc.CallOption) (*SimpleNodeTypeResponse, error)
	GetNodeTypeById(ctx context.Context, in *GetNodeTypeByIdRequest, opts ...grpc.CallOption) (*SimpleNodeTypeResponse, error)
	UpdateNodeTypeById(ctx context.Context, in *UpdateNodeTypeByIdRequest, opts ...grpc.CallOption) (*v1.SimpleResponse, error)
	DeleteNodeTypeById(ctx context.Context, in *DeleteNodeTypeByIdRequest, opts ...grpc.CallOption) (*v1.SimpleResponse, error)
	// -------------------------------------------------  Node  ----------------------------------------------------------
	CreateNodes(ctx context.Context, in *CreateNodesRequest, opts ...grpc.CallOption) (*v1.SimpleResponse, error)
	GetNodeById(ctx context.Context, in *GetNodeByIdRequest, opts ...grpc.CallOption) (*GetNodeResponse, error)
	GetNodesByWarehouseId(ctx context.Context, in *GetNodesByWarehouseIdRequest, opts ...grpc.CallOption) (*GetNodesResponse, error)
	// -------------------------------------------------  NodeConfigItem  ------------------------------------------------
	CreateNodeConfigItem(ctx context.Context, in *CreateNodeConfigItemRequest, opts ...grpc.CallOption) (*SimpleNodeConfigItemResponse, error)
	GetNodeConfigItemById(ctx context.Context, in *GetNodeConfigItemByIdRequest, opts ...grpc.CallOption) (*SimpleNodeConfigItemResponse, error)
	SelectNodeConfigItems(ctx context.Context, in *SelectNodeConfigItemRequest, opts ...grpc.CallOption) (*SelectNodeConfigItemResponse, error)
	UpdateNodeConfigItem(ctx context.Context, in *UpdateNodeConfigItemRequest, opts ...grpc.CallOption) (*v1.SimpleResponse, error)
	// -------------------------------------------------  AffixNode  -----------------------------------------------------
	CreateAffixNode(ctx context.Context, in *CreateAffixNodeRequest, opts ...grpc.CallOption) (*SimpleAffixNodeResponse, error)
	GetAffixNodeById(ctx context.Context, in *GetAffixNodeByIdRequest, opts ...grpc.CallOption) (*SimpleAffixNodeResponse, error)
	UpdateAffixNode(ctx context.Context, in *UpdateAffixNodeRequest, opts ...grpc.CallOption) (*v1.SimpleResponse, error)
}

type warehouseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWarehouseServiceClient(cc grpc.ClientConnInterface) WarehouseServiceClient {
	return &warehouseServiceClient{cc}
}

func (c *warehouseServiceClient) CreateWarehouse(ctx context.Context, in *CreateWarehouseRequest, opts ...grpc.CallOption) (*SimpleWarehouseResponse, error) {
	out := new(SimpleWarehouseResponse)
	err := c.cc.Invoke(ctx, WarehouseService_CreateWarehouse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseServiceClient) GetWarehouseByName(ctx context.Context, in *GetWarehouseByNameRequest, opts ...grpc.CallOption) (*SimpleWarehouseResponse, error) {
	out := new(SimpleWarehouseResponse)
	err := c.cc.Invoke(ctx, WarehouseService_GetWarehouseByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseServiceClient) GetWarehouseById(ctx context.Context, in *GetWarehouseByIdRequest, opts ...grpc.CallOption) (*SimpleWarehouseResponse, error) {
	out := new(SimpleWarehouseResponse)
	err := c.cc.Invoke(ctx, WarehouseService_GetWarehouseById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseServiceClient) UpdateWarehouseById(ctx context.Context, in *UpdateWarehouseByIdRequest, opts ...grpc.CallOption) (*v1.SimpleResponse, error) {
	out := new(v1.SimpleResponse)
	err := c.cc.Invoke(ctx, WarehouseService_UpdateWarehouseById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseServiceClient) UpdateWarehouseStatusById(ctx context.Context, in *UpdateWarehouseStatusByIdRequest, opts ...grpc.CallOption) (*v1.SimpleResponse, error) {
	out := new(v1.SimpleResponse)
	err := c.cc.Invoke(ctx, WarehouseService_UpdateWarehouseStatusById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseServiceClient) DeleteWarehouseById(ctx context.Context, in *DeleteWarehouseByIdRequest, opts ...grpc.CallOption) (*v1.SimpleResponse, error) {
	out := new(v1.SimpleResponse)
	err := c.cc.Invoke(ctx, WarehouseService_DeleteWarehouseById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseServiceClient) CreateNodeType(ctx context.Context, in *CreateNodeTypeRequest, opts ...grpc.CallOption) (*SimpleNodeTypeResponse, error) {
	out := new(SimpleNodeTypeResponse)
	err := c.cc.Invoke(ctx, WarehouseService_CreateNodeType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseServiceClient) GetNodeTypeById(ctx context.Context, in *GetNodeTypeByIdRequest, opts ...grpc.CallOption) (*SimpleNodeTypeResponse, error) {
	out := new(SimpleNodeTypeResponse)
	err := c.cc.Invoke(ctx, WarehouseService_GetNodeTypeById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseServiceClient) UpdateNodeTypeById(ctx context.Context, in *UpdateNodeTypeByIdRequest, opts ...grpc.CallOption) (*v1.SimpleResponse, error) {
	out := new(v1.SimpleResponse)
	err := c.cc.Invoke(ctx, WarehouseService_UpdateNodeTypeById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseServiceClient) DeleteNodeTypeById(ctx context.Context, in *DeleteNodeTypeByIdRequest, opts ...grpc.CallOption) (*v1.SimpleResponse, error) {
	out := new(v1.SimpleResponse)
	err := c.cc.Invoke(ctx, WarehouseService_DeleteNodeTypeById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseServiceClient) CreateNodes(ctx context.Context, in *CreateNodesRequest, opts ...grpc.CallOption) (*v1.SimpleResponse, error) {
	out := new(v1.SimpleResponse)
	err := c.cc.Invoke(ctx, WarehouseService_CreateNodes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseServiceClient) GetNodeById(ctx context.Context, in *GetNodeByIdRequest, opts ...grpc.CallOption) (*GetNodeResponse, error) {
	out := new(GetNodeResponse)
	err := c.cc.Invoke(ctx, WarehouseService_GetNodeById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseServiceClient) GetNodesByWarehouseId(ctx context.Context, in *GetNodesByWarehouseIdRequest, opts ...grpc.CallOption) (*GetNodesResponse, error) {
	out := new(GetNodesResponse)
	err := c.cc.Invoke(ctx, WarehouseService_GetNodesByWarehouseId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseServiceClient) CreateNodeConfigItem(ctx context.Context, in *CreateNodeConfigItemRequest, opts ...grpc.CallOption) (*SimpleNodeConfigItemResponse, error) {
	out := new(SimpleNodeConfigItemResponse)
	err := c.cc.Invoke(ctx, WarehouseService_CreateNodeConfigItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseServiceClient) GetNodeConfigItemById(ctx context.Context, in *GetNodeConfigItemByIdRequest, opts ...grpc.CallOption) (*SimpleNodeConfigItemResponse, error) {
	out := new(SimpleNodeConfigItemResponse)
	err := c.cc.Invoke(ctx, WarehouseService_GetNodeConfigItemById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseServiceClient) SelectNodeConfigItems(ctx context.Context, in *SelectNodeConfigItemRequest, opts ...grpc.CallOption) (*SelectNodeConfigItemResponse, error) {
	out := new(SelectNodeConfigItemResponse)
	err := c.cc.Invoke(ctx, WarehouseService_SelectNodeConfigItems_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseServiceClient) UpdateNodeConfigItem(ctx context.Context, in *UpdateNodeConfigItemRequest, opts ...grpc.CallOption) (*v1.SimpleResponse, error) {
	out := new(v1.SimpleResponse)
	err := c.cc.Invoke(ctx, WarehouseService_UpdateNodeConfigItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseServiceClient) CreateAffixNode(ctx context.Context, in *CreateAffixNodeRequest, opts ...grpc.CallOption) (*SimpleAffixNodeResponse, error) {
	out := new(SimpleAffixNodeResponse)
	err := c.cc.Invoke(ctx, WarehouseService_CreateAffixNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseServiceClient) GetAffixNodeById(ctx context.Context, in *GetAffixNodeByIdRequest, opts ...grpc.CallOption) (*SimpleAffixNodeResponse, error) {
	out := new(SimpleAffixNodeResponse)
	err := c.cc.Invoke(ctx, WarehouseService_GetAffixNodeById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseServiceClient) UpdateAffixNode(ctx context.Context, in *UpdateAffixNodeRequest, opts ...grpc.CallOption) (*v1.SimpleResponse, error) {
	out := new(v1.SimpleResponse)
	err := c.cc.Invoke(ctx, WarehouseService_UpdateAffixNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WarehouseServiceServer is the server API for WarehouseService service.
// All implementations must embed UnimplementedWarehouseServiceServer
// for forward compatibility
type WarehouseServiceServer interface {
	// -------------------------------------------------  Warehouse  -----------------------------------------------------
	CreateWarehouse(context.Context, *CreateWarehouseRequest) (*SimpleWarehouseResponse, error)
	GetWarehouseByName(context.Context, *GetWarehouseByNameRequest) (*SimpleWarehouseResponse, error)
	GetWarehouseById(context.Context, *GetWarehouseByIdRequest) (*SimpleWarehouseResponse, error)
	UpdateWarehouseById(context.Context, *UpdateWarehouseByIdRequest) (*v1.SimpleResponse, error)
	UpdateWarehouseStatusById(context.Context, *UpdateWarehouseStatusByIdRequest) (*v1.SimpleResponse, error)
	DeleteWarehouseById(context.Context, *DeleteWarehouseByIdRequest) (*v1.SimpleResponse, error)
	// -------------------------------------------------  NodeType  ------------------------------------------------------
	CreateNodeType(context.Context, *CreateNodeTypeRequest) (*SimpleNodeTypeResponse, error)
	GetNodeTypeById(context.Context, *GetNodeTypeByIdRequest) (*SimpleNodeTypeResponse, error)
	UpdateNodeTypeById(context.Context, *UpdateNodeTypeByIdRequest) (*v1.SimpleResponse, error)
	DeleteNodeTypeById(context.Context, *DeleteNodeTypeByIdRequest) (*v1.SimpleResponse, error)
	// -------------------------------------------------  Node  ----------------------------------------------------------
	CreateNodes(context.Context, *CreateNodesRequest) (*v1.SimpleResponse, error)
	GetNodeById(context.Context, *GetNodeByIdRequest) (*GetNodeResponse, error)
	GetNodesByWarehouseId(context.Context, *GetNodesByWarehouseIdRequest) (*GetNodesResponse, error)
	// -------------------------------------------------  NodeConfigItem  ------------------------------------------------
	CreateNodeConfigItem(context.Context, *CreateNodeConfigItemRequest) (*SimpleNodeConfigItemResponse, error)
	GetNodeConfigItemById(context.Context, *GetNodeConfigItemByIdRequest) (*SimpleNodeConfigItemResponse, error)
	SelectNodeConfigItems(context.Context, *SelectNodeConfigItemRequest) (*SelectNodeConfigItemResponse, error)
	UpdateNodeConfigItem(context.Context, *UpdateNodeConfigItemRequest) (*v1.SimpleResponse, error)
	// -------------------------------------------------  AffixNode  -----------------------------------------------------
	CreateAffixNode(context.Context, *CreateAffixNodeRequest) (*SimpleAffixNodeResponse, error)
	GetAffixNodeById(context.Context, *GetAffixNodeByIdRequest) (*SimpleAffixNodeResponse, error)
	UpdateAffixNode(context.Context, *UpdateAffixNodeRequest) (*v1.SimpleResponse, error)
	mustEmbedUnimplementedWarehouseServiceServer()
}

// UnimplementedWarehouseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWarehouseServiceServer struct {
}

func (UnimplementedWarehouseServiceServer) CreateWarehouse(context.Context, *CreateWarehouseRequest) (*SimpleWarehouseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWarehouse not implemented")
}
func (UnimplementedWarehouseServiceServer) GetWarehouseByName(context.Context, *GetWarehouseByNameRequest) (*SimpleWarehouseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWarehouseByName not implemented")
}
func (UnimplementedWarehouseServiceServer) GetWarehouseById(context.Context, *GetWarehouseByIdRequest) (*SimpleWarehouseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWarehouseById not implemented")
}
func (UnimplementedWarehouseServiceServer) UpdateWarehouseById(context.Context, *UpdateWarehouseByIdRequest) (*v1.SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWarehouseById not implemented")
}
func (UnimplementedWarehouseServiceServer) UpdateWarehouseStatusById(context.Context, *UpdateWarehouseStatusByIdRequest) (*v1.SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWarehouseStatusById not implemented")
}
func (UnimplementedWarehouseServiceServer) DeleteWarehouseById(context.Context, *DeleteWarehouseByIdRequest) (*v1.SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWarehouseById not implemented")
}
func (UnimplementedWarehouseServiceServer) CreateNodeType(context.Context, *CreateNodeTypeRequest) (*SimpleNodeTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNodeType not implemented")
}
func (UnimplementedWarehouseServiceServer) GetNodeTypeById(context.Context, *GetNodeTypeByIdRequest) (*SimpleNodeTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeTypeById not implemented")
}
func (UnimplementedWarehouseServiceServer) UpdateNodeTypeById(context.Context, *UpdateNodeTypeByIdRequest) (*v1.SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNodeTypeById not implemented")
}
func (UnimplementedWarehouseServiceServer) DeleteNodeTypeById(context.Context, *DeleteNodeTypeByIdRequest) (*v1.SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNodeTypeById not implemented")
}
func (UnimplementedWarehouseServiceServer) CreateNodes(context.Context, *CreateNodesRequest) (*v1.SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNodes not implemented")
}
func (UnimplementedWarehouseServiceServer) GetNodeById(context.Context, *GetNodeByIdRequest) (*GetNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeById not implemented")
}
func (UnimplementedWarehouseServiceServer) GetNodesByWarehouseId(context.Context, *GetNodesByWarehouseIdRequest) (*GetNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodesByWarehouseId not implemented")
}
func (UnimplementedWarehouseServiceServer) CreateNodeConfigItem(context.Context, *CreateNodeConfigItemRequest) (*SimpleNodeConfigItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNodeConfigItem not implemented")
}
func (UnimplementedWarehouseServiceServer) GetNodeConfigItemById(context.Context, *GetNodeConfigItemByIdRequest) (*SimpleNodeConfigItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeConfigItemById not implemented")
}
func (UnimplementedWarehouseServiceServer) SelectNodeConfigItems(context.Context, *SelectNodeConfigItemRequest) (*SelectNodeConfigItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectNodeConfigItems not implemented")
}
func (UnimplementedWarehouseServiceServer) UpdateNodeConfigItem(context.Context, *UpdateNodeConfigItemRequest) (*v1.SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNodeConfigItem not implemented")
}
func (UnimplementedWarehouseServiceServer) CreateAffixNode(context.Context, *CreateAffixNodeRequest) (*SimpleAffixNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAffixNode not implemented")
}
func (UnimplementedWarehouseServiceServer) GetAffixNodeById(context.Context, *GetAffixNodeByIdRequest) (*SimpleAffixNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAffixNodeById not implemented")
}
func (UnimplementedWarehouseServiceServer) UpdateAffixNode(context.Context, *UpdateAffixNodeRequest) (*v1.SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAffixNode not implemented")
}
func (UnimplementedWarehouseServiceServer) mustEmbedUnimplementedWarehouseServiceServer() {}

// UnsafeWarehouseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WarehouseServiceServer will
// result in compilation errors.
type UnsafeWarehouseServiceServer interface {
	mustEmbedUnimplementedWarehouseServiceServer()
}

func RegisterWarehouseServiceServer(s grpc.ServiceRegistrar, srv WarehouseServiceServer) {
	s.RegisterService(&WarehouseService_ServiceDesc, srv)
}

func _WarehouseService_CreateWarehouse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWarehouseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServiceServer).CreateWarehouse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarehouseService_CreateWarehouse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServiceServer).CreateWarehouse(ctx, req.(*CreateWarehouseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarehouseService_GetWarehouseByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWarehouseByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServiceServer).GetWarehouseByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarehouseService_GetWarehouseByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServiceServer).GetWarehouseByName(ctx, req.(*GetWarehouseByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarehouseService_GetWarehouseById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWarehouseByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServiceServer).GetWarehouseById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarehouseService_GetWarehouseById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServiceServer).GetWarehouseById(ctx, req.(*GetWarehouseByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarehouseService_UpdateWarehouseById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWarehouseByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServiceServer).UpdateWarehouseById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarehouseService_UpdateWarehouseById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServiceServer).UpdateWarehouseById(ctx, req.(*UpdateWarehouseByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarehouseService_UpdateWarehouseStatusById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWarehouseStatusByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServiceServer).UpdateWarehouseStatusById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarehouseService_UpdateWarehouseStatusById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServiceServer).UpdateWarehouseStatusById(ctx, req.(*UpdateWarehouseStatusByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarehouseService_DeleteWarehouseById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWarehouseByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServiceServer).DeleteWarehouseById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarehouseService_DeleteWarehouseById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServiceServer).DeleteWarehouseById(ctx, req.(*DeleteWarehouseByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarehouseService_CreateNodeType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNodeTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServiceServer).CreateNodeType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarehouseService_CreateNodeType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServiceServer).CreateNodeType(ctx, req.(*CreateNodeTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarehouseService_GetNodeTypeById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeTypeByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServiceServer).GetNodeTypeById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarehouseService_GetNodeTypeById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServiceServer).GetNodeTypeById(ctx, req.(*GetNodeTypeByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarehouseService_UpdateNodeTypeById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNodeTypeByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServiceServer).UpdateNodeTypeById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarehouseService_UpdateNodeTypeById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServiceServer).UpdateNodeTypeById(ctx, req.(*UpdateNodeTypeByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarehouseService_DeleteNodeTypeById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNodeTypeByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServiceServer).DeleteNodeTypeById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarehouseService_DeleteNodeTypeById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServiceServer).DeleteNodeTypeById(ctx, req.(*DeleteNodeTypeByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarehouseService_CreateNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServiceServer).CreateNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarehouseService_CreateNodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServiceServer).CreateNodes(ctx, req.(*CreateNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarehouseService_GetNodeById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServiceServer).GetNodeById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarehouseService_GetNodeById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServiceServer).GetNodeById(ctx, req.(*GetNodeByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarehouseService_GetNodesByWarehouseId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodesByWarehouseIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServiceServer).GetNodesByWarehouseId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarehouseService_GetNodesByWarehouseId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServiceServer).GetNodesByWarehouseId(ctx, req.(*GetNodesByWarehouseIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarehouseService_CreateNodeConfigItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNodeConfigItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServiceServer).CreateNodeConfigItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarehouseService_CreateNodeConfigItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServiceServer).CreateNodeConfigItem(ctx, req.(*CreateNodeConfigItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarehouseService_GetNodeConfigItemById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeConfigItemByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServiceServer).GetNodeConfigItemById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarehouseService_GetNodeConfigItemById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServiceServer).GetNodeConfigItemById(ctx, req.(*GetNodeConfigItemByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarehouseService_SelectNodeConfigItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectNodeConfigItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServiceServer).SelectNodeConfigItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarehouseService_SelectNodeConfigItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServiceServer).SelectNodeConfigItems(ctx, req.(*SelectNodeConfigItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarehouseService_UpdateNodeConfigItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNodeConfigItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServiceServer).UpdateNodeConfigItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarehouseService_UpdateNodeConfigItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServiceServer).UpdateNodeConfigItem(ctx, req.(*UpdateNodeConfigItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarehouseService_CreateAffixNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAffixNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServiceServer).CreateAffixNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarehouseService_CreateAffixNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServiceServer).CreateAffixNode(ctx, req.(*CreateAffixNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarehouseService_GetAffixNodeById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAffixNodeByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServiceServer).GetAffixNodeById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarehouseService_GetAffixNodeById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServiceServer).GetAffixNodeById(ctx, req.(*GetAffixNodeByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarehouseService_UpdateAffixNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAffixNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServiceServer).UpdateAffixNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarehouseService_UpdateAffixNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServiceServer).UpdateAffixNode(ctx, req.(*UpdateAffixNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WarehouseService_ServiceDesc is the grpc.ServiceDesc for WarehouseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WarehouseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "warehouse.v1.WarehouseService",
	HandlerType: (*WarehouseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWarehouse",
			Handler:    _WarehouseService_CreateWarehouse_Handler,
		},
		{
			MethodName: "GetWarehouseByName",
			Handler:    _WarehouseService_GetWarehouseByName_Handler,
		},
		{
			MethodName: "GetWarehouseById",
			Handler:    _WarehouseService_GetWarehouseById_Handler,
		},
		{
			MethodName: "UpdateWarehouseById",
			Handler:    _WarehouseService_UpdateWarehouseById_Handler,
		},
		{
			MethodName: "UpdateWarehouseStatusById",
			Handler:    _WarehouseService_UpdateWarehouseStatusById_Handler,
		},
		{
			MethodName: "DeleteWarehouseById",
			Handler:    _WarehouseService_DeleteWarehouseById_Handler,
		},
		{
			MethodName: "CreateNodeType",
			Handler:    _WarehouseService_CreateNodeType_Handler,
		},
		{
			MethodName: "GetNodeTypeById",
			Handler:    _WarehouseService_GetNodeTypeById_Handler,
		},
		{
			MethodName: "UpdateNodeTypeById",
			Handler:    _WarehouseService_UpdateNodeTypeById_Handler,
		},
		{
			MethodName: "DeleteNodeTypeById",
			Handler:    _WarehouseService_DeleteNodeTypeById_Handler,
		},
		{
			MethodName: "CreateNodes",
			Handler:    _WarehouseService_CreateNodes_Handler,
		},
		{
			MethodName: "GetNodeById",
			Handler:    _WarehouseService_GetNodeById_Handler,
		},
		{
			MethodName: "GetNodesByWarehouseId",
			Handler:    _WarehouseService_GetNodesByWarehouseId_Handler,
		},
		{
			MethodName: "CreateNodeConfigItem",
			Handler:    _WarehouseService_CreateNodeConfigItem_Handler,
		},
		{
			MethodName: "GetNodeConfigItemById",
			Handler:    _WarehouseService_GetNodeConfigItemById_Handler,
		},
		{
			MethodName: "SelectNodeConfigItems",
			Handler:    _WarehouseService_SelectNodeConfigItems_Handler,
		},
		{
			MethodName: "UpdateNodeConfigItem",
			Handler:    _WarehouseService_UpdateNodeConfigItem_Handler,
		},
		{
			MethodName: "CreateAffixNode",
			Handler:    _WarehouseService_CreateAffixNode_Handler,
		},
		{
			MethodName: "GetAffixNodeById",
			Handler:    _WarehouseService_GetAffixNodeById_Handler,
		},
		{
			MethodName: "UpdateAffixNode",
			Handler:    _WarehouseService_UpdateAffixNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "warehouse/v1/warehouse.proto",
}
