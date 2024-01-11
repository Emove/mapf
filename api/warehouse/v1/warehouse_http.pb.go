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
	v1 "mapf/api/common/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationWarehouseServiceCreateNodeType = "/warehouse.v1.WarehouseService/CreateNodeType"
const OperationWarehouseServiceCreateWarehouse = "/warehouse.v1.WarehouseService/CreateWarehouse"
const OperationWarehouseServiceDeleteNodeTypeById = "/warehouse.v1.WarehouseService/DeleteNodeTypeById"
const OperationWarehouseServiceDeleteWarehouseById = "/warehouse.v1.WarehouseService/DeleteWarehouseById"
const OperationWarehouseServiceGetNodeTypeById = "/warehouse.v1.WarehouseService/GetNodeTypeById"
const OperationWarehouseServiceGetWarehouseById = "/warehouse.v1.WarehouseService/GetWarehouseById"
const OperationWarehouseServiceGetWarehouseByName = "/warehouse.v1.WarehouseService/GetWarehouseByName"
const OperationWarehouseServiceUpdateNodeTypeById = "/warehouse.v1.WarehouseService/UpdateNodeTypeById"
const OperationWarehouseServiceUpdateWarehouseById = "/warehouse.v1.WarehouseService/UpdateWarehouseById"
const OperationWarehouseServiceUpdateWarehouseStatusById = "/warehouse.v1.WarehouseService/UpdateWarehouseStatusById"

type WarehouseServiceHTTPServer interface {
	// CreateNodeType -------------------------------------------------  NodeType  ------------------------------------------------------
	CreateNodeType(context.Context, *CreateNodeTypeRequest) (*CreateNodeTypeResponse, error)
	// CreateWarehouse -------------------------------------------------  Warehouse  -----------------------------------------------------
	CreateWarehouse(context.Context, *CreateWarehouseRequest) (*CreateWarehouseResponse, error)
	DeleteNodeTypeById(context.Context, *DeleteNodeTypeByIdRequest) (*v1.SimpleResponse, error)
	DeleteWarehouseById(context.Context, *DeleteWarehouseByIdRequest) (*v1.SimpleResponse, error)
	GetNodeTypeById(context.Context, *GetNodeTypeByIdRequest) (*GetNodeTypeResponse, error)
	GetWarehouseById(context.Context, *GetWarehouseByIdRequest) (*GetWarehouseResponse, error)
	GetWarehouseByName(context.Context, *GetWarehouseByNameRequest) (*GetWarehouseResponse, error)
	UpdateNodeTypeById(context.Context, *UpdateNodeTypeByIdRequest) (*v1.SimpleResponse, error)
	UpdateWarehouseById(context.Context, *UpdateWarehouseByIdRequest) (*v1.SimpleResponse, error)
	UpdateWarehouseStatusById(context.Context, *UpdateWarehouseStatusByIdRequest) (*v1.SimpleResponse, error)
}

func RegisterWarehouseServiceHTTPServer(s *http.Server, srv WarehouseServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/warehouse", _WarehouseService_CreateWarehouse0_HTTP_Handler(srv))
	r.GET("/warehouse/name/{name}", _WarehouseService_GetWarehouseByName0_HTTP_Handler(srv))
	r.GET("/warehouse/{id}", _WarehouseService_GetWarehouseById0_HTTP_Handler(srv))
	r.PUT("/warehouse/{id}", _WarehouseService_UpdateWarehouseById0_HTTP_Handler(srv))
	r.PUT("/warehouse/{id}/status", _WarehouseService_UpdateWarehouseStatusById0_HTTP_Handler(srv))
	r.DELETE("/warehouse/{id}", _WarehouseService_DeleteWarehouseById0_HTTP_Handler(srv))
	r.POST("/node_type", _WarehouseService_CreateNodeType0_HTTP_Handler(srv))
	r.GET("/node_type/{id}", _WarehouseService_GetNodeTypeById0_HTTP_Handler(srv))
	r.PUT("/node_type/{id}", _WarehouseService_UpdateNodeTypeById0_HTTP_Handler(srv))
	r.DELETE("/node_type/{id}", _WarehouseService_DeleteNodeTypeById0_HTTP_Handler(srv))
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

func _WarehouseService_UpdateWarehouseById0_HTTP_Handler(srv WarehouseServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateWarehouseByIdRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWarehouseServiceUpdateWarehouseById)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateWarehouseById(ctx, req.(*UpdateWarehouseByIdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.SimpleResponse)
		return ctx.Result(200, reply)
	}
}

func _WarehouseService_UpdateWarehouseStatusById0_HTTP_Handler(srv WarehouseServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateWarehouseStatusByIdRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWarehouseServiceUpdateWarehouseStatusById)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateWarehouseStatusById(ctx, req.(*UpdateWarehouseStatusByIdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.SimpleResponse)
		return ctx.Result(200, reply)
	}
}

func _WarehouseService_DeleteWarehouseById0_HTTP_Handler(srv WarehouseServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteWarehouseByIdRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWarehouseServiceDeleteWarehouseById)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteWarehouseById(ctx, req.(*DeleteWarehouseByIdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.SimpleResponse)
		return ctx.Result(200, reply)
	}
}

func _WarehouseService_CreateNodeType0_HTTP_Handler(srv WarehouseServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateNodeTypeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWarehouseServiceCreateNodeType)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateNodeType(ctx, req.(*CreateNodeTypeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateNodeTypeResponse)
		return ctx.Result(200, reply)
	}
}

func _WarehouseService_GetNodeTypeById0_HTTP_Handler(srv WarehouseServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetNodeTypeByIdRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWarehouseServiceGetNodeTypeById)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetNodeTypeById(ctx, req.(*GetNodeTypeByIdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetNodeTypeResponse)
		return ctx.Result(200, reply)
	}
}

func _WarehouseService_UpdateNodeTypeById0_HTTP_Handler(srv WarehouseServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateNodeTypeByIdRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWarehouseServiceUpdateNodeTypeById)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateNodeTypeById(ctx, req.(*UpdateNodeTypeByIdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.SimpleResponse)
		return ctx.Result(200, reply)
	}
}

func _WarehouseService_DeleteNodeTypeById0_HTTP_Handler(srv WarehouseServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteNodeTypeByIdRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationWarehouseServiceDeleteNodeTypeById)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteNodeTypeById(ctx, req.(*DeleteNodeTypeByIdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.SimpleResponse)
		return ctx.Result(200, reply)
	}
}

type WarehouseServiceHTTPClient interface {
	CreateNodeType(ctx context.Context, req *CreateNodeTypeRequest, opts ...http.CallOption) (rsp *CreateNodeTypeResponse, err error)
	CreateWarehouse(ctx context.Context, req *CreateWarehouseRequest, opts ...http.CallOption) (rsp *CreateWarehouseResponse, err error)
	DeleteNodeTypeById(ctx context.Context, req *DeleteNodeTypeByIdRequest, opts ...http.CallOption) (rsp *v1.SimpleResponse, err error)
	DeleteWarehouseById(ctx context.Context, req *DeleteWarehouseByIdRequest, opts ...http.CallOption) (rsp *v1.SimpleResponse, err error)
	GetNodeTypeById(ctx context.Context, req *GetNodeTypeByIdRequest, opts ...http.CallOption) (rsp *GetNodeTypeResponse, err error)
	GetWarehouseById(ctx context.Context, req *GetWarehouseByIdRequest, opts ...http.CallOption) (rsp *GetWarehouseResponse, err error)
	GetWarehouseByName(ctx context.Context, req *GetWarehouseByNameRequest, opts ...http.CallOption) (rsp *GetWarehouseResponse, err error)
	UpdateNodeTypeById(ctx context.Context, req *UpdateNodeTypeByIdRequest, opts ...http.CallOption) (rsp *v1.SimpleResponse, err error)
	UpdateWarehouseById(ctx context.Context, req *UpdateWarehouseByIdRequest, opts ...http.CallOption) (rsp *v1.SimpleResponse, err error)
	UpdateWarehouseStatusById(ctx context.Context, req *UpdateWarehouseStatusByIdRequest, opts ...http.CallOption) (rsp *v1.SimpleResponse, err error)
}

type WarehouseServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewWarehouseServiceHTTPClient(client *http.Client) WarehouseServiceHTTPClient {
	return &WarehouseServiceHTTPClientImpl{client}
}

func (c *WarehouseServiceHTTPClientImpl) CreateNodeType(ctx context.Context, in *CreateNodeTypeRequest, opts ...http.CallOption) (*CreateNodeTypeResponse, error) {
	var out CreateNodeTypeResponse
	pattern := "/node_type"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationWarehouseServiceCreateNodeType))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
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

func (c *WarehouseServiceHTTPClientImpl) DeleteNodeTypeById(ctx context.Context, in *DeleteNodeTypeByIdRequest, opts ...http.CallOption) (*v1.SimpleResponse, error) {
	var out v1.SimpleResponse
	pattern := "/node_type/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationWarehouseServiceDeleteNodeTypeById))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WarehouseServiceHTTPClientImpl) DeleteWarehouseById(ctx context.Context, in *DeleteWarehouseByIdRequest, opts ...http.CallOption) (*v1.SimpleResponse, error) {
	var out v1.SimpleResponse
	pattern := "/warehouse/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationWarehouseServiceDeleteWarehouseById))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WarehouseServiceHTTPClientImpl) GetNodeTypeById(ctx context.Context, in *GetNodeTypeByIdRequest, opts ...http.CallOption) (*GetNodeTypeResponse, error) {
	var out GetNodeTypeResponse
	pattern := "/node_type/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationWarehouseServiceGetNodeTypeById))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WarehouseServiceHTTPClientImpl) GetWarehouseById(ctx context.Context, in *GetWarehouseByIdRequest, opts ...http.CallOption) (*GetWarehouseResponse, error) {
	var out GetWarehouseResponse
	pattern := "/warehouse/{id}"
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
	pattern := "/warehouse/name/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationWarehouseServiceGetWarehouseByName))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WarehouseServiceHTTPClientImpl) UpdateNodeTypeById(ctx context.Context, in *UpdateNodeTypeByIdRequest, opts ...http.CallOption) (*v1.SimpleResponse, error) {
	var out v1.SimpleResponse
	pattern := "/node_type/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationWarehouseServiceUpdateNodeTypeById))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WarehouseServiceHTTPClientImpl) UpdateWarehouseById(ctx context.Context, in *UpdateWarehouseByIdRequest, opts ...http.CallOption) (*v1.SimpleResponse, error) {
	var out v1.SimpleResponse
	pattern := "/warehouse/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationWarehouseServiceUpdateWarehouseById))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WarehouseServiceHTTPClientImpl) UpdateWarehouseStatusById(ctx context.Context, in *UpdateWarehouseStatusByIdRequest, opts ...http.CallOption) (*v1.SimpleResponse, error) {
	var out v1.SimpleResponse
	pattern := "/warehouse/{id}/status"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationWarehouseServiceUpdateWarehouseStatusById))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
