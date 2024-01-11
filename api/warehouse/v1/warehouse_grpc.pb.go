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
)

// WarehouseServiceClient is the client API for WarehouseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WarehouseServiceClient interface {
	// -------------------------------------------------  Warehouse  -----------------------------------------------------
	CreateWarehouse(ctx context.Context, in *CreateWarehouseRequest, opts ...grpc.CallOption) (*CreateWarehouseResponse, error)
	GetWarehouseByName(ctx context.Context, in *GetWarehouseByNameRequest, opts ...grpc.CallOption) (*GetWarehouseResponse, error)
	GetWarehouseById(ctx context.Context, in *GetWarehouseByIdRequest, opts ...grpc.CallOption) (*GetWarehouseResponse, error)
	UpdateWarehouseById(ctx context.Context, in *UpdateWarehouseByIdRequest, opts ...grpc.CallOption) (*v1.SimpleResponse, error)
	UpdateWarehouseStatusById(ctx context.Context, in *UpdateWarehouseStatusByIdRequest, opts ...grpc.CallOption) (*v1.SimpleResponse, error)
	DeleteWarehouseById(ctx context.Context, in *DeleteWarehouseByIdRequest, opts ...grpc.CallOption) (*v1.SimpleResponse, error)
	// -------------------------------------------------  NodeType  ------------------------------------------------------
	CreateNodeType(ctx context.Context, in *CreateNodeTypeRequest, opts ...grpc.CallOption) (*CreateNodeTypeResponse, error)
	GetNodeTypeById(ctx context.Context, in *GetNodeTypeByIdRequest, opts ...grpc.CallOption) (*GetNodeTypeResponse, error)
	UpdateNodeTypeById(ctx context.Context, in *UpdateNodeTypeByIdRequest, opts ...grpc.CallOption) (*v1.SimpleResponse, error)
	DeleteNodeTypeById(ctx context.Context, in *DeleteNodeTypeByIdRequest, opts ...grpc.CallOption) (*v1.SimpleResponse, error)
}

type warehouseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWarehouseServiceClient(cc grpc.ClientConnInterface) WarehouseServiceClient {
	return &warehouseServiceClient{cc}
}

func (c *warehouseServiceClient) CreateWarehouse(ctx context.Context, in *CreateWarehouseRequest, opts ...grpc.CallOption) (*CreateWarehouseResponse, error) {
	out := new(CreateWarehouseResponse)
	err := c.cc.Invoke(ctx, WarehouseService_CreateWarehouse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseServiceClient) GetWarehouseByName(ctx context.Context, in *GetWarehouseByNameRequest, opts ...grpc.CallOption) (*GetWarehouseResponse, error) {
	out := new(GetWarehouseResponse)
	err := c.cc.Invoke(ctx, WarehouseService_GetWarehouseByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseServiceClient) GetWarehouseById(ctx context.Context, in *GetWarehouseByIdRequest, opts ...grpc.CallOption) (*GetWarehouseResponse, error) {
	out := new(GetWarehouseResponse)
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

func (c *warehouseServiceClient) CreateNodeType(ctx context.Context, in *CreateNodeTypeRequest, opts ...grpc.CallOption) (*CreateNodeTypeResponse, error) {
	out := new(CreateNodeTypeResponse)
	err := c.cc.Invoke(ctx, WarehouseService_CreateNodeType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseServiceClient) GetNodeTypeById(ctx context.Context, in *GetNodeTypeByIdRequest, opts ...grpc.CallOption) (*GetNodeTypeResponse, error) {
	out := new(GetNodeTypeResponse)
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

// WarehouseServiceServer is the server API for WarehouseService service.
// All implementations must embed UnimplementedWarehouseServiceServer
// for forward compatibility
type WarehouseServiceServer interface {
	// -------------------------------------------------  Warehouse  -----------------------------------------------------
	CreateWarehouse(context.Context, *CreateWarehouseRequest) (*CreateWarehouseResponse, error)
	GetWarehouseByName(context.Context, *GetWarehouseByNameRequest) (*GetWarehouseResponse, error)
	GetWarehouseById(context.Context, *GetWarehouseByIdRequest) (*GetWarehouseResponse, error)
	UpdateWarehouseById(context.Context, *UpdateWarehouseByIdRequest) (*v1.SimpleResponse, error)
	UpdateWarehouseStatusById(context.Context, *UpdateWarehouseStatusByIdRequest) (*v1.SimpleResponse, error)
	DeleteWarehouseById(context.Context, *DeleteWarehouseByIdRequest) (*v1.SimpleResponse, error)
	// -------------------------------------------------  NodeType  ------------------------------------------------------
	CreateNodeType(context.Context, *CreateNodeTypeRequest) (*CreateNodeTypeResponse, error)
	GetNodeTypeById(context.Context, *GetNodeTypeByIdRequest) (*GetNodeTypeResponse, error)
	UpdateNodeTypeById(context.Context, *UpdateNodeTypeByIdRequest) (*v1.SimpleResponse, error)
	DeleteNodeTypeById(context.Context, *DeleteNodeTypeByIdRequest) (*v1.SimpleResponse, error)
	mustEmbedUnimplementedWarehouseServiceServer()
}

// UnimplementedWarehouseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWarehouseServiceServer struct {
}

func (UnimplementedWarehouseServiceServer) CreateWarehouse(context.Context, *CreateWarehouseRequest) (*CreateWarehouseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWarehouse not implemented")
}
func (UnimplementedWarehouseServiceServer) GetWarehouseByName(context.Context, *GetWarehouseByNameRequest) (*GetWarehouseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWarehouseByName not implemented")
}
func (UnimplementedWarehouseServiceServer) GetWarehouseById(context.Context, *GetWarehouseByIdRequest) (*GetWarehouseResponse, error) {
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
func (UnimplementedWarehouseServiceServer) CreateNodeType(context.Context, *CreateNodeTypeRequest) (*CreateNodeTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNodeType not implemented")
}
func (UnimplementedWarehouseServiceServer) GetNodeTypeById(context.Context, *GetNodeTypeByIdRequest) (*GetNodeTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeTypeById not implemented")
}
func (UnimplementedWarehouseServiceServer) UpdateNodeTypeById(context.Context, *UpdateNodeTypeByIdRequest) (*v1.SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNodeTypeById not implemented")
}
func (UnimplementedWarehouseServiceServer) DeleteNodeTypeById(context.Context, *DeleteNodeTypeByIdRequest) (*v1.SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNodeTypeById not implemented")
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "warehouse/v1/warehouse.proto",
}
