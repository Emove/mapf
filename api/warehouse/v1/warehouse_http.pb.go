// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.2
// - protoc             v3.21.12
// source: warehouse/v1/warehouse.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationWarehouseServiceCreateWarehouse = "/warehouse.v1.WarehouseService/CreateWarehouse"
const OperationWarehouseServiceGetWarehouseById = "/warehouse.v1.WarehouseService/GetWarehouseById"
const OperationWarehouseServiceGetWarehouseByName = "/warehouse.v1.WarehouseService/GetWarehouseByName"

type WarehouseServiceHTTPServer interface {
	// CreateWarehouse Creates a warehouse
	CreateWarehouse(context.Context, *CreateWarehouseRequest) (*CreateWarehouseResponse, error)
	GetWarehouseById(context.Context, *GetWarehouseByIdRequest) (*GetWarehouseResponse, error)
	GetWarehouseByName(context.Context, *GetWarehouseByNameRequest) (*GetWarehouseResponse, error)
}

func RegisterWarehouseServiceHTTPServer(s *http.Server, srv WarehouseServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/warehouse", _WarehouseService_CreateWarehouse0_HTTP_Handler(srv))
	r.GET("/warehouse/{name}", _WarehouseService_GetWarehouseByName0_HTTP_Handler(srv))
	r.GET("/warehouse/id/{id}", _WarehouseService_GetWarehouseById0_HTTP_Handler(srv))
}

func _WarehouseService_CreateWarehouse0_HTTP_Handler(srv WarehouseServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateWarehouseRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWarehouseServiceCreateWarehouse)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateWarehouse(ctx, req.(*CreateWarehouseRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateWarehouseResponse)
		return ctx.Result(200, reply)
	}
}

func _WarehouseService_GetWarehouseByName0_HTTP_Handler(srv WarehouseServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetWarehouseByNameRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWarehouseServiceGetWarehouseByName)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetWarehouseByName(ctx, req.(*GetWarehouseByNameRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetWarehouseResponse)
		return ctx.Result(200, reply)
	}
}

func _WarehouseService_GetWarehouseById0_HTTP_Handler(srv WarehouseServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetWarehouseByIdRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWarehouseServiceGetWarehouseById)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetWarehouseById(ctx, req.(*GetWarehouseByIdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetWarehouseResponse)
		return ctx.Result(200, reply)
	}
}

type WarehouseServiceHTTPClient interface {
	CreateWarehouse(ctx context.Context, req *CreateWarehouseRequest, opts ...http.CallOption) (rsp *CreateWarehouseResponse, err error)
	GetWarehouseById(ctx context.Context, req *GetWarehouseByIdRequest, opts ...http.CallOption) (rsp *GetWarehouseResponse, err error)
	GetWarehouseByName(ctx context.Context, req *GetWarehouseByNameRequest, opts ...http.CallOption) (rsp *GetWarehouseResponse, err error)
}

type WarehouseServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewWarehouseServiceHTTPClient(client *http.Client) WarehouseServiceHTTPClient {
	return &WarehouseServiceHTTPClientImpl{client}
}

func (c *WarehouseServiceHTTPClientImpl) CreateWarehouse(ctx context.Context, in *CreateWarehouseRequest, opts ...http.CallOption) (*CreateWarehouseResponse, error) {
	var out CreateWarehouseResponse
	pattern := "/warehouse"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationWarehouseServiceCreateWarehouse))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WarehouseServiceHTTPClientImpl) GetWarehouseById(ctx context.Context, in *GetWarehouseByIdRequest, opts ...http.CallOption) (*GetWarehouseResponse, error) {
	var out GetWarehouseResponse
	pattern := "/warehouse/id/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationWarehouseServiceGetWarehouseById))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WarehouseServiceHTTPClientImpl) GetWarehouseByName(ctx context.Context, in *GetWarehouseByNameRequest, opts ...http.CallOption) (*GetWarehouseResponse, error) {
	var out GetWarehouseResponse
	pattern := "/warehouse/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationWarehouseServiceGetWarehouseByName))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
