package biz

import "mapf/internal/data"

type NodeTag struct {
	*data.Model
	AffixNodeId int    `gorm:"index:idx_affix;not null"`
	Name        string `gorm:"type:varchar(63);not null"`
	Value       string `gorm:"type:varchar(127);not null'"`
	Status      int    `gorm:"type:smallint;default:1;not null"`
}

func (*NodeTag) TableName() string {
	return "node_tag"
}

type NodeTagRepo interface {
}
