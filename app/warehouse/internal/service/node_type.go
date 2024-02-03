package service

import (
	"context"
	v1 "mapf/api/warehouse/v1"
	"mapf/app/warehouse/internal/biz"
)

func (s *WarehouseService) CreateNodeType(ctx context.Context, req *v1.CreateNodeTypeRequest) (*v1.SimpleNodeTypeResponse, error) {
	nodeType := &biz.NodeType{Code: req.Code, Name: req.Name}
	nodeType, err := s.ntuc.CreateNodeType(ctx, nodeType)
	if err != nil {
		return nil, err
	}
	var resp *v1.SimpleNodeTypeResponse
	if nodeType != nil {
		resp.NodeType = ConvertNodeType2protobuf(nodeType)
	}
	return resp, nil
}

func (s *WarehouseService) GetNodeTypeById(ctx context.Context, req *v1.GetNodeTypeByIdRequest) (resp *v1.GetNodeTypeResponse, err error) {
	nodeType, err := s.ntuc.GetNodeTypeById(ctx, int(req.Id))
	if err != nil {
		return
	}
	resp = &v1.GetNodeTypeResponse{}
	if nodeType != nil {
		resp.NodeType = ConvertNodeType2protobuf(nodeType)
	}
	return
}

func (s *WarehouseService) UpdateNodeTypeById(ctx context.Context, req *v1.UpdateNodeTypeByIdRequest) (resp *v1.SimpleResponse, err error) {
	nodeType := &biz.NodeType{}
	if req.Code != nil {
		nodeType.Code = *req.Code
	}
	if req.Name != nil {
		nodeType.Name = *req.Name
	}
	if req.IsDefault != nil {
		nodeType.IsDefault = int(*req.IsDefault)
	}
	resp = &v1.SimpleResponse{}
	err = s.ntuc.UpdateNodeTypeById(ctx, int(req.Id), nodeType)
	if err != nil {
		resp.Msg = "更新失败"
		return
	}
	resp.Msg = "更新成功"
	return
}

func (s *WarehouseService) DeleteNodeTypeById(ctx context.Context, req *v1.DeleteNodeTypeByIdRequest) (*v1.SimpleResponse, error) {
	err := s.ntuc.DeleteNodeTypeById(ctx, int(req.Id))
	if err != nil {
		return &v1.SimpleResponse{Msg: "删除失败"}, err
	}
	return &v1.SimpleResponse{Msg: "删除成功"}, nil
}

func ConvertNodeType2protobuf(nodeType *biz.NodeType) *v1.NodeType {
	isDefault := ConvertModelIsDefault2protoIsDefaultEnum(nodeType.IsDefault)
	return &v1.NodeType{
		Id:        uint32(nodeType.ID),
		Code:      nodeType.Code,
		Name:      nodeType.Name,
		IsDefault: &isDefault,
	}
}
