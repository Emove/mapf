package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"mapf/internal/data"
	"mapf/internal/data/tx"
	dataerrors "mapf/internal/errors"
	"mapf/internal/utils"
	"strings"
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
	// GetNodeConfigItemById 根据ID查询节点配置项
	GetNodeConfigItemById(context.Context, int) (*NodeConfigItem, error)
	// SelectNodeConfigItem 查询节点配置项，可分页
	SelectNodeConfigItem(context.Context, *NodeConfigItem, *data.Page) ([]*NodeConfigItem, *data.Page, error)
	// GetNodeConfigItemByWarehouseIdAndNodeTypeId 查询仓库节点配置项，节点类型可为空
	GetNodeConfigItemByWarehouseIdAndNodeTypeId(context.Context, int, *int) (items []*NodeConfigItem, err error)
}

type NodeConfigItemUsecase struct {
	tm            tx.Transaction
	repo          NodeConfigItemRepo
	warehouseRepo WarehouseRepo
	nodeTypeRepo  NodeTypeRepo
	logger        *log.Helper
}

func NewNodeConfigItemUsecase(tm tx.Transaction, repo NodeConfigItemRepo, warehouseRepo WarehouseRepo, nodeTypeRepo NodeTypeRepo, logger log.Logger) *NodeConfigItemUsecase {
	return &NodeConfigItemUsecase{tm: tm, repo: repo, warehouseRepo: warehouseRepo, nodeTypeRepo: nodeTypeRepo, logger: log.NewHelper(logger)}
}

// CreateNodeConfigItem 创建仓库节点配置项
func (uc *NodeConfigItemUsecase) CreateNodeConfigItem(ctx context.Context, item *NodeConfigItem) (*NodeConfigItem, error) {
	warehouse, err := uc.warehouseRepo.GetWarehouseById(ctx, item.WarehouseId)
	if err != nil {
		return nil, err
	}
	if data.IsDeleted(warehouse.IsDeleted) {
		return nil, dataerrors.NewInvalidArgumentError("given Warehouse: [%d] not found", item.WarehouseId)
	}
	if data.IsDisable(warehouse.Status) {
		return nil, dataerrors.NewInvalidArgumentError("given Warehouse: [%d] is disable", item.WarehouseId)
	}
	nodeType, err := uc.nodeTypeRepo.GetNodeTypeById(ctx, item.NodeTypeId)
	if err != nil {
		return nil, err
	}
	if data.IsDeleted(nodeType.IsDeleted) {
		return nil, dataerrors.NewInvalidArgumentError("given NodeType: [%d] not found", item.NodeTypeId)
	}
	return uc.repo.CreateNodeConfigItem(ctx, item)
}

// UpdateNodeConfigItem 更新节点配置项
func (uc *NodeConfigItemUsecase) UpdateNodeConfigItem(ctx context.Context, item *NodeConfigItem) (*NodeConfigItem, error) {
	originItem, err := uc.repo.GetNodeConfigItemById(ctx, item.ID)
	if err != nil {
		return nil, err
	}
	if utils.NotBlank(item.Code) && originItem.Code != item.Code {
		originItem.SetCode(strings.TrimSpace(item.Code))
	}
	if utils.NotBlank(item.Name) && originItem.Name != item.Name {
		originItem.SetName(strings.TrimSpace(item.Name))
	}
	if utils.NotBlank(item.Name) && originItem.ValueType != item.ValueType {
		err = originItem.SetValueType(strings.TrimSpace(item.ValueType))
		if err != nil {
			return nil, err
		}
	}
	if utils.NotBlank(item.DefaultValue) && originItem.DefaultValue != item.DefaultValue {
		err = originItem.SetDefaultValue(strings.TrimSpace(item.DefaultValue))
		if err != nil {
			return nil, err
		}
	}
	if utils.NotBlank(item.OptionalValues) && originItem.OptionalValues != item.OptionalValues {
		err = originItem.SetOptionalValues(strings.TrimSpace(item.OptionalValues))
		if err != nil {
			return nil, err
		}
	}
	return uc.repo.UpdateNodeConfigItem(ctx, originItem)
}

// EnableNodeConfigItem 启用节点配置项
func (uc *NodeConfigItemUsecase) EnableNodeConfigItem(ctx context.Context, id int) error {
	return nil
}

// DisableNodeConfigItem 禁用节点配置项
func (uc *NodeConfigItemUsecase) DisableNodeConfigItem(ctx context.Context, id int) error {
	return nil
}

// GetNodeConfigItemById 根据节点配置项获取节点配置项信息
func (uc *NodeConfigItemUsecase) GetNodeConfigItemById(ctx context.Context, id int) (*NodeConfigItem, error) {
	return uc.repo.GetNodeConfigItemById(ctx, id)
}

func (uc *NodeConfigItemUsecase) SelectNodeConfigItem(ctx context.Context, item *NodeConfigItem, page *data.Page) ([]*NodeConfigItem, *data.Page, error) {
	return uc.repo.SelectNodeConfigItem(ctx, item, page)
}

// GetNodeConfigItemByWarehouseIdAndNodeTypeId 根据仓库ID和节点类型ID查询节点配置项
func (uc *NodeConfigItemUsecase) GetNodeConfigItemByWarehouseIdAndNodeTypeId(ctx context.Context, warehouseId int, nodeTypeId *int) ([]*NodeConfigItem, error) {
	return uc.repo.GetNodeConfigItemByWarehouseIdAndNodeTypeId(ctx, warehouseId, nodeTypeId)
}
