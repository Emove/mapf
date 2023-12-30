package biz

import "mapf/internal/data"

type Node struct {
	*data.Model
	WarehouseId int `gorm:"uniqueIndex:uidx_warehouse_x_y;not null"`
	X           int `gorm:"uniqueIndex:uidx_warehouse_x_y;not null"`
	Y           int `gorm:"uniqueIndex:uidx_warehouse_x_y;not null"`
}

type NodeRepo interface {
}
