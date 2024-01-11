package biz

import "mapf/internal/data"

type NodeRelation struct {
	*data.Model
	WarehouseId int `gorm:"index:idx_warehouse_node_affix;not null"`
	NodeId      int `gorm:"index:idx_warehouse_node_affix;not null"`
	AffixNodeId int `gorm:"index:idx_warehouse_name_affix;not null"`
	NodeTypeId  int `gorm:"not null"`
	Status      int `gorm:"type:smallint;default:1;not null"`
}

func (*NodeRelation) TableName() string {
	return "node_relation"
}

type NodeRelationRepo interface {
}
