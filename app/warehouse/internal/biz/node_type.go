package biz

import "mapf/internal/data"

type NodeType struct {
	*data.Model
	WarehouseId int    `gorm:"uniqueIndex:uidx_warehouse_code;not null"`
	Code        string `gorm:"uniqueIndex:uidx_warehouse_code;type:varchar(63);not null"`
	Name        string `gorm:"type:varchar(63);not null"`
	IsDefault   int    `gorm:"type:smallint;default:0;not null"`
}

type NodeTypeRepo interface {
}
