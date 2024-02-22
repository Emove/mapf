package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"mapf/app/warehouse/internal/biz"
	"mapf/internal/data"
	dataerrors "mapf/internal/errors"
)

type nodeDiagramRepo struct {
	*data.Pager
	data   *Data
	logger *log.Helper
}

func NewNodeDiagramRepo(data *Data, logger log.Logger) (biz.NodeDiagramRepo, error) {
	return &nodeDiagramRepo{
		data:   data,
		logger: log.NewHelper(logger),
	}, nil
}

func (repo *nodeDiagramRepo) CreateNodeDiagram(ctx context.Context, diagram *biz.NodeDiagram) (*biz.NodeDiagram, error) {
	err := repo.data.DB(ctx).Create(diagram).Error
	if err != nil && errors.Is(err, gorm.ErrDuplicatedKey) {
		return nil, dataerrors.NewDuplicatedKeyError("NodeDiagramRelation warehouse_id: [%d], node_id: [%d], reachable_node_id: [%d] Existed", diagram.WarehouseId, diagram.NodeId, diagram.ReachableNodeId)
	}
	return diagram, err
}

func (repo *nodeDiagramRepo) BatchCreateNodeDiagrams(ctx context.Context, diagrams []*biz.NodeDiagram) ([]*biz.NodeDiagram, error) {
	err := repo.data.DB(ctx).Create(diagrams).Error
	return diagrams, err
}

func (repo *nodeDiagramRepo) UpdateNodeDiagram(ctx context.Context, diagram *biz.NodeDiagram) (*biz.NodeDiagram, error) {
	stmt := repo.data.DB(ctx).Model(&biz.NodeDiagram{Model: &data.Model{ID: diagram.ID}})
	result := stmt.Updates(diagram.GetChanges())
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, dataerrors.NewDuplicatedKeyError("NodeDiagramRelation warehouse_id: [%d], node_id: [%d], reachable_node_id: [%d] Existed", diagram.WarehouseId, diagram.NodeId, diagram.ReachableNodeId)
		}
		return nil, result.Error
	}
	return diagram, nil
}

func (repo *nodeDiagramRepo) SelectNodeDiagram(ctx context.Context, query *biz.NodeDiagram, page *data.Page) ([]*biz.NodeDiagram, *data.Page, error) {
	var diagrams []*biz.NodeDiagram
	var err error
	db, page, err := repo.Pager.Prepare(repo.data.DB(ctx), query, page)
	if err != nil {
		return nil, nil, err
	}
	err = db.Find(&diagrams).Error
	return diagrams, page, err
}

func (repo *nodeDiagramRepo) DeleteNodeDiagramById(ctx context.Context, id int) error {
	return repo.data.DB(ctx).
		Model(&biz.NodeDiagram{Model: &data.Model{ID: id}}).
		Updates(biz.NodeDiagram{Model: &data.Model{IsDeleted: data.IsDeletedTrue}}).Error
}
