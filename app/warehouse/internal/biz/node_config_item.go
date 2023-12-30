package biz

import "mapf/internal/data"

// NodeConfigItem 节点配置项表
type NodeConfigItem struct {
	*data.Model
	WarehouseId    int    `gorm:"index:idx_warehouse_node_type;not null"`
	NodeTypeId     int    `gorm:"index:idx_warehouse_node_type;not null"`
	Code           string `gorm:"type:varchar(63);not null"`
	Name           string `gorm:"type:varchar(63);not null"`
	ValueType      string `gorm:"type:varchar(31);not null"`
	OptionalValues string `gorm:"type:varchar(1023)"`
	Status         int    `gorm:"type:smallint;default:1"`
}

type NodeConfigItemRepo interface {
}
