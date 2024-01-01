package biz

import "mapf/internal/data"

type AffixNode struct {
	*data.Model
	NodeId     int    `gorm:"index:idx_node;not null"`
	NodeTypeId int    `gorm:"index:idx_node_type;not null"`
	Name       string `gorm:"type:varchar(64);not null"`
	Remark     string `gorm:"type:varchar(255)"`
	Status     int    `gorm:"type:smallint;default:1;not null"`
}

type AffixNodeRepo interface {
}
