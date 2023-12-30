// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: api/warehouse/v1/warehouse.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Warehouse_CreateWarehouse_FullMethodName = "/warehouse.v1.Warehouse/CreateWarehouse"
)

// WarehouseClient is the client API for Warehouse service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WarehouseClient interface {
	// Creates a warehouse
	CreateWarehouse(ctx context.Context, in *CreateWarehouseRequest, opts ...grpc.CallOption) (*CreateWarehouseResponse, error)
}

type warehouseClient struct {
	cc grpc.ClientConnInterface
}

func NewWarehouseClient(cc grpc.ClientConnInterface) WarehouseClient {
	return &warehouseClient{cc}
}

func (c *warehouseClient) CreateWarehouse(ctx context.Context, in *CreateWarehouseRequest, opts ...grpc.CallOption) (*CreateWarehouseResponse, error) {
	out := new(CreateWarehouseResponse)
	err := c.cc.Invoke(ctx, Warehouse_CreateWarehouse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WarehouseServer is the server API for Warehouse service.
// All implementations must embed UnimplementedWarehouseServer
// for forward compatibility
type WarehouseServer interface {
	// Creates a warehouse
	CreateWarehouse(context.Context, *CreateWarehouseRequest) (*CreateWarehouseResponse, error)
	mustEmbedUnimplementedWarehouseServer()
}

// UnimplementedWarehouseServer must be embedded to have forward compatible implementations.
type UnimplementedWarehouseServer struct {
}

func (UnimplementedWarehouseServer) CreateWarehouse(context.Context, *CreateWarehouseRequest) (*CreateWarehouseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWarehouse not implemented")
}
func (UnimplementedWarehouseServer) mustEmbedUnimplementedWarehouseServer() {}

// UnsafeWarehouseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WarehouseServer will
// result in compilation errors.
type UnsafeWarehouseServer interface {
	mustEmbedUnimplementedWarehouseServer()
}

func RegisterWarehouseServer(s grpc.ServiceRegistrar, srv WarehouseServer) {
	s.RegisterService(&Warehouse_ServiceDesc, srv)
}

func _Warehouse_CreateWarehouse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWarehouseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServer).CreateWarehouse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Warehouse_CreateWarehouse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServer).CreateWarehouse(ctx, req.(*CreateWarehouseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Warehouse_ServiceDesc is the grpc.ServiceDesc for Warehouse service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Warehouse_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "warehouse.v1.Warehouse",
	HandlerType: (*WarehouseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWarehouse",
			Handler:    _Warehouse_CreateWarehouse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/warehouse/v1/warehouse.proto",
}