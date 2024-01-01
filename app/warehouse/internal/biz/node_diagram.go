package biz

import (
	"mapf/internal/data"
)

// NodeDiagram 节点连通关系
type NodeDiagram struct {
	*data.Model
	WarehouseId     int `gorm:"index:idx_warehouse_node;index:idx_warehouse_relation;not null"`
	NodeId          int `gorm:"index:idx_warehouse_node;not null"`
	ReachableNodeId int `gorm:"index:idx_warehouse_relation;not null"`
	Distance        int `gorm:"default:0;not null"` //Millimeter
	Status          int `gorm:"type:smallint;default:1;not null"`
}

type NodeDiagramRepo interface {
}
