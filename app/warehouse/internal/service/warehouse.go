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
		Warehouse: &v1.Warehouse{
			Id:     uint32(warehouse.ID),
			Name:   warehouse.Name,
			Status: ConvertWarehouseStatusToWarehouseStatusEnum(warehouse.Status),
		},
		NodeTypes: pbNodeTypes,
	}, err
}

func ConvertWarehouseStatusToWarehouseStatusEnum(warehouseStatus int) v1.EnableStatus {
	if warehouseStatus == data.DisableStatus {
		return v1.EnableStatus_DISABLE
	}
	return v1.EnableStatus_ENABLE
}
