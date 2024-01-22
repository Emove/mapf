package service

import (
	"context"
	v1 "mapf/api/warehouse/v1"
	"mapf/app/warehouse/internal/biz"
)

func (s *WarehouseService) CreateNodes(ctx context.Context, req *v1.CreateNodesRequest) (*v1.SimpleResponse, error) {
	err := s.nuc.CreateNodes(ctx, int(req.WarehouseId), int(req.Start), int(req.End), int(req.AuxAix), biz.NodeTupleType(req.NodeTupleType))
	if err != nil {
		return &v1.SimpleResponse{Msg: "创建节点失败"}, err
	}
	return &v1.SimpleResponse{Msg: "节点创建中"}, nil
}

func (s *WarehouseService) GetNodeById(ctx context.Context, req *v1.GetNodeByIdRequest) (resp *v1.GetNodeResponse, err error) {
	node, err := s.nuc.GetNodeById(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	resp = &v1.GetNodeResponse{}
	if node != nil {
		resp.Node = ConvertNode2protobuf(node)
	}
	return
}

func (s *WarehouseService) GetNodesByWarehouseId(ctx context.Context, req *v1.GetNodesByWarehouseIdRequest) (resp *v1.GetNodesResponse, err error) {
	nodes, err := s.nuc.GetNodesByWarehouseId(ctx, int(req.WarehouseId))
	if err != nil {
		return nil, err
	}
	resp = &v1.GetNodesResponse{}
	if len(nodes) > 0 {
		var protobufNodes []*v1.Node
		for _, node := range nodes {
			protobufNodes = append(protobufNodes, ConvertNode2protobuf(node))
		}
		resp.Nodes = protobufNodes
	}
	return resp, nil
}

func ConvertNode2protobuf(node *biz.Node) *v1.Node {
	return &v1.Node{
		Id:          uint32(node.ID),
		WarehouseId: uint32(node.WarehouseId),
		X:           uint32(node.X),
		Y:           uint32(node.Y),
	}
}
