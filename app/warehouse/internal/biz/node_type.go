package biz

import (
	"context"
	"mapf/internal/data"
)

type NodeType struct {
	*data.Model
	WarehouseId int    `gorm:"uniqueIndex:uidx_warehouse_code;not null"`
	Code        string `gorm:"uniqueIndex:uidx_warehouse_code;type:varchar(63);not null"`
	Name        string `gorm:"type:varchar(63);not null"`
	IsDefault   int    `gorm:"type:smallint;default:0;not null"`
}

func (*NodeType) TableName() string {
	return "node_type"
}

type NodeTypeRepo interface {
	CreateNodeType(ctx context.Context, nodeType *NodeType) (*NodeType, error)
	BatchCreateNodeTypes(ctx context.Context, nodeTypes []*NodeType) ([]*NodeType, error)
}
