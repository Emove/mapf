package biz

import (
	"context"
	"mapf/internal/data"
)

type NodeRelation struct {
	*data.Model
	WarehouseId int `gorm:"uniqueIndex:uidx_warehouse_node_affix;not null"`
	NodeId      int `gorm:"uniqueIndex:uidx_warehouse_node_affix;not null"`
	AffixNodeId int `gorm:"uniqueIndex:uidx_warehouse_name_affix;not null"`
	NodeTypeId  int `gorm:"not null"`
	Status      int `gorm:"type:smallint;default:1;not null"`
}

func (*NodeRelation) TableName() string {
	return "node_relation"
}

type NodeRelationRepo interface {
	// CreateNodeRelation 创建节点关联关系
	CreateNodeRelation(context.Context, *NodeRelation) (*NodeRelation, error)
	// GetNodeRelationById 根据ID获取节点关联关系
	GetNodeRelationById(context.Context, int) (*NodeRelation, error)
	// UpdateNodeRelation 更新节点关联关系
	UpdateNodeRelation(context.Context, *NodeRelation) (*NodeRelation, error)
	// SelectNodeRelation 查询节点关联关系。可分页
	SelectNodeRelation(context.Context, *NodeRelation, *data.Page) ([]*NodeRelation, *data.Page, error)
	// DeleteNodeRelation 删除节点关联关系
	DeleteNodeRelation(context.Context, int) error
}
