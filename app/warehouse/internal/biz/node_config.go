package biz

import (
	"context"
	"mapf/internal/data"
)

// NodeConfig 节点配置表
type NodeConfig struct {
	*data.Model
	WarehouseId    int    `gorm:"uniqueIndex:uidx_warehouse_x_y;not null"`
	NodeRelationId int    `gorm:"uniqueIndex:uidx_warehouse_x_y;not null"`
	ConfigItemId   int    `gorm:"uniqueIndex:uidx_config_item;not null"`
	Value          string `gorm:"type:varchar(1023)"`
	IsDefault      int    `gorm:"type:smallint;default:2;not null"`
	Status         int    `gorm:"type:smallint;default:1;not null"`
}

func (*NodeConfig) TableName() string {
	return "node_config"
}

type NodeConfigRepo interface {
	// CreateNodeConfig 创建节点配置信息
	CreateNodeConfig(context.Context, *NodeConfig) (*NodeConfig, error)
	// UpdateNodeConfig 更新节点配置信息
	UpdateNodeConfig(context.Context, *NodeConfig) (*NodeConfig, error)
	// GetNodeConfigById 根据ID获取节点配置信息
	GetNodeConfigById(context.Context, int) (*NodeConfig, error)
	// SelectNodeConfig 查询节点配置信息，可分页
	SelectNodeConfig(context.Context, *NodeConfig, *data.Page) ([]*NodeConfig, *data.Page, error)
	// DeleteNodeConfigById 删除节点配置信息
	DeleteNodeConfigById(context.Context, int) error
}

type NodeConfigUsecase struct {
}

func (uc *NodeConfigUsecase) UpdateNodeConfig() {

	// TODO 更新节点配置之后，应该自动更新地图
}
