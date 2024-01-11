package service

import (
	"context"
	v1 "mapf/api/warehouse/v1"
	"mapf/app/warehouse/internal/biz"
	"mapf/internal/data"
)

func (s *WarehouseService) CreateWarehouse(ctx context.Context, req *v1.CreateWarehouseRequest) (*v1.CreateWarehouseResponse, error) {
	warehouse := &biz.Warehouse{
		Name:      req.Name,
		IsDefault: data.IsDefaultFalse,
	}
	warehouse, err := s.wuc.CreateWarehouse(ctx, warehouse)
	return &v1.CreateWarehouseResponse{
		Warehouse: ConvertWarehouse2protobuf(warehouse),
	}, err
}

func (s *WarehouseService) GetWarehouseByName(ctx context.Context, req *v1.GetWarehouseByNameRequest) (*v1.GetWarehouseResponse, error) {
	warehouse, err := s.wuc.GetWarehouseByName(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	resp := &v1.GetWarehouseResponse{Warehouse: ConvertWarehouse2protobuf(warehouse)}
	return resp, nil
}

func (s *WarehouseService) GetWarehouseById(ctx context.Context, req *v1.GetWarehouseByIdRequest) (*v1.GetWarehouseResponse, error) {
	warehouse, err := s.wuc.GetWarehouseById(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	resp := &v1.GetWarehouseResponse{Warehouse: ConvertWarehouse2protobuf(warehouse)}
	return resp, nil
}

func (s *WarehouseService) UpdateWarehouseById(ctx context.Context, req *v1.UpdateWarehouseByIdRequest) (*v1.SimpleResponse, error) {
	warehouse := &biz.Warehouse{
		Name:      req.Name,
		IsDefault: int(req.IsDefault),
	}
	err := s.wuc.UpdateWarehouseById(ctx, int(req.Id), warehouse)
	if err != nil {
		return &v1.SimpleResponse{Msg: "更新失败"}, err
	}
	return &v1.SimpleResponse{Msg: "更新成功"}, err
}

func (s *WarehouseService) UpdateWarehouseStatusById(ctx context.Context, req *v1.UpdateWarehouseStatusByIdRequest) (resp *v1.SimpleResponse, err error) {
	resp = &v1.SimpleResponse{}
	if req.Status == data.EnableStatus {
		err = s.wuc.EnableWarehouseById(ctx, int(req.Id))
		resp.Msg = "启用仓库"
	} else {
		err = s.wuc.DisableWarehouseById(ctx, int(req.Id))
		resp.Msg = "禁用仓库"
	}
	result := "成功"
	if err != nil {
		result = "失败"
		return
	}
	resp.Msg = resp.Msg + result
	return
}

func (s *WarehouseService) DeleteWarehouseById(ctx context.Context, req *v1.DeleteWarehouseByIdRequest) (*v1.SimpleResponse, error) {
	err := s.wuc.DeleteWarehouseById(ctx, int(req.Id))
	if err != nil {
		return &v1.SimpleResponse{Msg: "删除仓库失败"}, err
	}
	return &v1.SimpleResponse{Msg: "删除仓库成功"}, err
}

func ConvertWarehouse2protobuf(warehouse *biz.Warehouse) *v1.Warehouse {
	if warehouse == nil {
		return nil
	}
	return &v1.Warehouse{
		Id:     uint32(warehouse.ID),
		Name:   warehouse.Name,
		Status: ConvertModelStatus2protobufEnableStatusEnum(warehouse.Status),
	}
}
