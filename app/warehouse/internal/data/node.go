package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"mapf/app/warehouse/internal/biz"
	"mapf/internal/data"
	"mapf/internal/data/tx"
)

var _ biz.NodeRepo = (*nodeRepo)(nil)

type nodeRepo struct {
	data   *Data
	logger *log.Helper
}

func NewNodeRepo(data *Data, logger log.Logger) (biz.NodeRepo, error) {
	return &nodeRepo{
		data:   data,
		logger: log.NewHelper(logger),
	}, nil
}

// CreateNode 创建节点
func (repo *nodeRepo) CreateNode(ctx context.Context, node *biz.Node) (*biz.Node, error) {
	err := repo.data.DB(ctx).Create(node).Error
	return node, err
}

// BatchCreateNode 批量创建节点
func (repo *nodeRepo) BatchCreateNode(ctx context.Context, nodes []*biz.Node) ([]*biz.Node, error) {
	err := repo.data.InTx(ctx, func(ctx context.Context) error {
		db := tx.GetTxFromContext(ctx)
		if err := db.Create(nodes).Error; err != nil {
			return err
		}
		return nil
	})

	return nodes, err
}

// GetNodeById 通过ID获取节点
func (repo *nodeRepo) GetNodeById(ctx context.Context, id int) (node *biz.Node, err error) {
	err = repo.data.DB(ctx).Where(biz.NodeType{Model: &data.Model{ID: id}}).First(&node).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
}

// GetNodesByIds 通过ID列表获取节点列表
func (repo *nodeRepo) GetNodesByIds(ctx context.Context, ids []int) (nodes []*biz.Node, err error) {
	err = repo.data.DB(ctx).Where(ids).Find(&nodes).Error
	return
}

// GetNodesByWarehouseId 通过仓库ID查询节点
func (repo *nodeRepo) GetNodesByWarehouseId(ctx context.Context, warehouseId int) (nodes []*biz.Node, err error) {
	err = repo.data.DB(ctx).Where(biz.Node{WarehouseId: warehouseId}).Find(&nodes).Error
	return
}

// GetNodeByWarehouseAndCoordinate 根据仓库ID和坐标查询节点
func (repo *nodeRepo) GetNodeByWarehouseAndCoordinate(ctx context.Context, warehouseId, x, y int) (node *biz.Node, err error) {
	err = repo.data.DB(ctx).Where(biz.Node{WarehouseId: warehouseId, X: x, Y: y}).Find(&node).Error
	return
}

// SelectNodeTupleByCoorRange 根据坐标范围查询节点信息
func (repo *nodeRepo) SelectNodeTupleByCoorRange(ctx context.Context, warehouseId, start, end, auxAix int, tupleType biz.NodeTupleType) (biz.NodeTuple, error) {
	db := repo.data.DB(ctx).Where(biz.Node{WarehouseId: warehouseId})
	if tupleType == biz.NodeTupleTypeColumn {
		db.Where("x = ? and y >= ? and y <= ?", auxAix, start, end)
	} else {
		db.Where("x >= ? and x <= ? and y = ?", start, end, auxAix)
	}
	var tuple biz.NodeTuple
	err := db.Find(&tuple).Error
	return tuple, err
}
