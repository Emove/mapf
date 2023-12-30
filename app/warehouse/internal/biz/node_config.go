package biz

import "mapf/internal/data"

// NodeConfig 节点配置表
type NodeConfig struct {
	*data.Model
	WarehouseId    int    `gorm:"index:idx_warehouse_x_y;not null"`
	NodeRelationId int    `gorm:"index:idx_warehouse_x_y;not null"`
	ConfigItemId   int    `gorm:"index:idx_config_item;not null"`
	ConfigValue    string `gorm:"type:varchar(1023)"`
	IsDefault      int    `gorm:"type:smallint;default:0"`
	Status         int    `gorm:"type:smallint;default:1"`
}

type NodeConfigRepo interface {
}
