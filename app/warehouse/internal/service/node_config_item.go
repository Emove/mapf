package service

import (
	"context"
	"encoding/json"
	v1 "mapf/api/warehouse/v1"
	"mapf/app/warehouse/internal/biz"
	"mapf/internal/data"
)

func (s *WarehouseService) CreateNodeConfigItem(ctx context.Context, req *v1.CreateNodeConfigItemRequest) (*v1.SimpleNodeConfigItemResponse, error) {
	var optionalValues []string
	for _, optionalValue := range req.OptionalValues {
		optionalValues = append(optionalValues, optionalValue)
	}
	ovs, err := json.Marshal(optionalValues)
	if err != nil {
		return nil, err
	}
	item := &biz.NodeConfigItem{
		WarehouseId:    int(req.WarehouseId),
		NodeTypeId:     int(req.NodeTypeId),
		Code:           req.Code,
		Name:           req.Name,
		ValueType:      req.ValueType,
		DefaultValue:   req.DefaultValue,
		OptionalValues: string(ovs),
	}
	item, err = s.nciuc.CreateNodeConfigItem(ctx, item)
	if err != nil {
		return nil, err
	}
	return &v1.SimpleNodeConfigItemResponse{Item: ConvertNodeConfigItem2protobuf(item)}, nil
}

func (s *WarehouseService) GetNodeConfigItemById(ctx context.Context, req *v1.GetNodeConfigItemByIdRequest) (*v1.SimpleNodeConfigItemResponse, error) {
	item, err := s.nciuc.GetNodeConfigItemById(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	resp := &v1.SimpleNodeConfigItemResponse{}
	if item != nil {
		resp.Item = ConvertNodeConfigItem2protobuf(item)
	}
	return resp, nil
}

func (s *WarehouseService) SelectNodeConfigItems(ctx context.Context, req *v1.SelectNodeConfigItemRequest) (*v1.SelectNodeConfigItemResponse, error) {
	query := &biz.NodeConfigItem{
		WarehouseId: int(req.WarehouseId),
		Code:        *req.Code,
		Name:        *req.Name,
		ValueType:   *req.ValueType,
	}
	if req.NodeTypeId != nil {
		nodeTypeId := *req.NodeTypeId
		query.NodeTypeId = int(nodeTypeId)
	}
	if req.Status != nil {
		query.Status = ConvertProtobufEnableStatus2modelStatus(*req.Status)
	}
	var page *data.Page
	if req.Page != nil {
		page.Size = int64(req.Page.Size)
		page.Num = int64(req.Page.Num)
	}
	items, page, err := s.nciuc.SelectNodeConfigItem(ctx, query, page)
	if err != nil {
		return nil, err
	}
	resp := &v1.SelectNodeConfigItemResponse{}
	for _, item := range items {
		resp.Items = append(resp.Items, ConvertNodeConfigItem2protobuf(item))
	}
	if page != nil {
		resp.Page = &v1.Page{
			Size:  uint32(page.Size),
			Num:   uint32(page.Num),
			Total: uint32(page.Total),
		}
	}
	return resp, nil
}

func (s *WarehouseService) UpdateNodeConfigItem(ctx context.Context, req *v1.UpdateNodeConfigItemRequest) (*v1.SimpleResponse, error) {
	item := &biz.NodeConfigItem{
		Model:        &data.Model{ID: int(req.Id)},
		Code:         *req.Code,
		Name:         *req.Name,
		ValueType:    *req.ValueType,
		DefaultValue: *req.DefaultValue,
	}
	if req.OptionalValues != nil {
		var ovs []string
		for _, value := range req.OptionalValues {
			ovs = append(ovs, value)
		}
		ovsStr, err := json.Marshal(ovs)
		if err != nil {
			return nil, err
		}
		item.OptionalValues = string(ovsStr)
	}
	_, err := s.nciuc.UpdateNodeConfigItem(ctx, item)
	resp := &v1.SimpleResponse{
		Msg: "更新失败",
	}
	if err != nil {
		return resp, err
	}
	resp.Msg = "更新成功"
	return resp, nil
}

func ConvertNodeConfigItem2protobuf(item *biz.NodeConfigItem) *v1.NodeConfigItem {
	var optionalValues []string
	_ = json.Unmarshal([]byte(item.OptionalValues), &optionalValues)
	return &v1.NodeConfigItem{
		Id:             uint32(item.ID),
		WarehouseId:    uint32(item.WarehouseId),
		NodeTypeId:     uint32(item.NodeTypeId),
		Code:           item.Code,
		Name:           item.Name,
		ValueType:      item.ValueType,
		DefaultValue:   item.DefaultValue,
		OptionalValues: optionalValues,
		Status:         ConvertModelStatus2protobufEnableStatusEnum(item.Status),
	}
}
