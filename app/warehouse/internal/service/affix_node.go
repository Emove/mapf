package service

import (
	"context"
	v1 "mapf/api/warehouse/v1"
	"mapf/app/warehouse/internal/biz"
)

func (s *WarehouseService) CreateAffixNodes(ctx context.Context, req *v1.CreateAffixNodeRequest) (*v1.SimpleAffixNodeResponse, error) {
	return nil, nil
}

func (s *WarehouseService) GetAffixNodeById(ctx context.Context, req *v1.GetAffixNodeByIdRequest) (*v1.SimpleAffixNodeResponse, error) {
	affixNode, err := s.anuc.GetAffixNodeById(ctx, req.Id)
	if err != nil || affixNode == nil {
		return nil, err
	}
	return &v1.SimpleAffixNodeResponse{AffixNode: ConvertAffixNode2protobuf(affixNode)}, nil
}

func (s *WarehouseService) UpdateAffixNodeById(ctx context.Context, req *v1.UpdateAffixNodeRequest) (*v1.SimpleResponse, error) {
	affixNode := &biz.AffixNode{
		NodeTypeId: int(req.NodeTypeId),
		NodeId:     int(req.NodeId),
		Name:       req.Name,
	}
	if req.Remark != nil {
		affixNode.Remark = *req.Remark
	}
	_, err := s.anuc.UpdateAffixNode(ctx, affixNode)
	if err != nil {
		return nil, err
	}
	return &v1.SimpleResponse{Msg: "更新节点信息成功"}, nil
}

func ConvertAffixNode2protobuf(node *biz.AffixNode) *v1.AffixNode {
	return &v1.AffixNode{
		Id:          uint32(node.ID),
		WarehouseId: uint32(node.WarehouseId),
		NodeTypeId:  uint32(node.NodeTypeId),
		NodeId:      uint32(node.NodeId),
		Name:        node.Name,
		Remark:      node.Remark,
		Status:      ConvertModelStatus2protobufEnableStatusEnum(node.Status),
	}
}
