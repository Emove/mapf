package biz

import (
	"context"
	"mapf/internal/data"
)

// NodeConfigItem 节点配置项表
type NodeConfigItem struct {
	*data.Model
	WarehouseId    int    `gorm:"uniqueIndex:uidx_warehouse_node_type_code;not null"`
	NodeTypeId     int    `gorm:"uniqueIndex:uidx_warehouse_node_type_code;not null"`
	Code           string `gorm:"type:varchar(63);uniqueIndex:uidx_warehouse_node_type_code;not null"`
	Name           string `gorm:"type:varchar(63);not null"`
	ValueType      string `gorm:"type:varchar(31);not null"`
	DefaultValue   string `gorm:"type:varchar(63);not null"`
	OptionalValues string `gorm:"type:varchar(1023)"`
	Status         int    `gorm:"type:smallint;default:1;not null"`
}

func (item *NodeConfigItem) SetCode(code string) {
	item.Code = code
	item.Update("code", code)
}

func (item *NodeConfigItem) SetName(name string) {
	item.Name = name
	item.Update("name", name)
}

func (item *NodeConfigItem) SetValueType(valueType string) error {
	// FIXME validate value type
	item.ValueType = valueType
	item.Update("value_type", valueType)
	return nil
}

func (item *NodeConfigItem) SetDefaultValue(defaultValue string) error {
	// FIXME validate value
	item.DefaultValue = defaultValue
	item.Update("default_value", defaultValue)
	return nil
}

func (item *NodeConfigItem) SetOptionalValues(optionalValues string) error {
	// FIXME validate optional values
	item.OptionalValues = optionalValues
	item.Update("optional_values", optionalValues)
	return nil
}

func (*NodeConfigItem) TableName() string {
	return "node_config_time"
}

type NodeConfigItemRepo interface {
	// CreateNodeConfigItem 创建仓库节点配置项
	CreateNodeConfigItem(context.Context, *NodeConfigItem) (*NodeConfigItem, error)
	// UpdateNodeConfigItem 更新节点配置项
	UpdateNodeConfigItem(context.Context, *NodeConfigItem) (*NodeConfigItem, error)
	// GetNodeConfigItemByWarehouseIdAndNodeTypeId 查询仓库节点配置项，节点类型可为空
	GetNodeConfigItemByWarehouseIdAndNodeTypeId(context.Context, int, *int) (items []*NodeConfigItem, err error)
}
