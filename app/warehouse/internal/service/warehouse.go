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
	var nodeTypes []*biz.NodeType
	for _, nodeType := range req.NodeTypes {
		nodeTypes = append(nodeTypes, &biz.NodeType{
			Code:      nodeType.Code,
			Name:      nodeType.Name,
			IsDefault: data.IsDefaultFalse,
		})
	}
	warehouse, nodeTypes, err := s.uc.CreateWarehouse(ctx, warehouse, nodeTypes)
	if err != nil {
		return nil, err
	}
	pbNodeTypes := make([]*v1.NodeType, len(nodeTypes))
	for i, nodeType := range nodeTypes {
		pbNodeTypes[i] = &v1.NodeType{
			Id:          uint32(nodeType.ID),
			WarehouseId: uint32(nodeType.WarehouseId),
			Code:        nodeType.Code,
			Name:        nodeType.Name,
		}
	}
	return &v1.CreateWarehouseResponse{
		Warehouse: ConvertWarehouseToProtobuf(warehouse),
		NodeTypes: pbNodeTypes,
	}, err
}

func (s *WarehouseService) GetWarehouseByName(ctx context.Context, req *v1.GetWarehouseByNameRequest) (*v1.GetWarehouseResponse, error) {
	warehouse, err := s.uc.GetWarehouseByName(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	resp := &v1.GetWarehouseResponse{Warehouse: ConvertWarehouseToProtobuf(warehouse)}
	return resp, nil
}

func (s *WarehouseService) GetWarehouseById(ctx context.Context, req *v1.GetWarehouseByIdRequest) (*v1.GetWarehouseResponse, error) {
	warehouse, err := s.uc.GetWarehouseById(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	resp := &v1.GetWarehouseResponse{Warehouse: ConvertWarehouseToProtobuf(warehouse)}
	return resp, nil
}

func ConvertWarehouseToProtobuf(warehouse *biz.Warehouse) *v1.Warehouse {
	if warehouse == nil {
		return nil
	}
	return &v1.Warehouse{
		Id:     uint32(warehouse.ID),
		Name:   warehouse.Name,
		Status: ConvertWarehouseStatusToWarehouseStatusEnum(warehouse.Status),
	}
}

func ConvertWarehouseStatusToWarehouseStatusEnum(warehouseStatus int) v1.EnableStatus {
	if warehouseStatus == data.DisableStatus {
		return v1.EnableStatus_DISABLE
	}
	return v1.EnableStatus_ENABLE
}
