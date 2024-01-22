package biz

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"mapf/app/warehouse/internal/event/publisher"
	"mapf/internal/data"
	"mapf/internal/data/tx"
	dataerrors "mapf/internal/errors"
	internalevent "mapf/internal/event"
)

type Node struct {
	*data.Model
	WarehouseId int `gorm:"uniqueIndex:uidx_warehouse_x_y;not null"`
	X           int `gorm:"uniqueIndex:uidx_warehouse_x_y;not null"`
	Y           int `gorm:"uniqueIndex:uidx_warehouse_x_y;not null"`
}

type NodeTupleType int

const (
	NodeTupleTypeColumn = 1
	NodeTupleTypeRow    = 2
)

type NodeTuple []*Node

func (*Node) TableName() string {
	return "node"
}

type NodeRepo interface {
	// CreateNode 创建节点
	CreateNode(context.Context, *Node) (*Node, error)
	// BatchCreateNode 批量创建节点
	BatchCreateNode(context.Context, []*Node) ([]*Node, error)
	// GetNodeById 根据节点ID获取节点信息
	GetNodeById(context.Context, int) (*Node, error)
	// GetNodesByIds 根据节点ID列表批量获取节点信息
	GetNodesByIds(context.Context, []int) ([]*Node, error)
	// GetNodesByWarehouseId 根据仓库ID获取节点信息列表
	GetNodesByWarehouseId(context.Context, int) ([]*Node, error)
	// SelectNodeTupleByCoorRange 根据坐标范围查询节点信息
	SelectNodeTupleByCoorRange(context.Context, int, int, int, int, NodeTupleType) (NodeTuple, error)
}

type NodeUsecase struct {
	tm                  tx.Transaction
	repo                NodeRepo
	warehouseRepo       WarehouseRepo
	createNodePublisher internalevent.Publisher
	logger              *log.Helper
}

func NewNodeUsecase(
	tm tx.Transaction,
	repo NodeRepo,
	warehouseRepo WarehouseRepo,
	createNodePublisher internalevent.Publisher,
	createNodeSubscriber internalevent.Subscriber,
	logger log.Logger,
) (*NodeUsecase, error) {
	uc := &NodeUsecase{
		tm:                  tm,
		repo:                repo,
		warehouseRepo:       warehouseRepo,
		createNodePublisher: createNodePublisher,
		logger:              log.NewHelper(logger),
	}
	err := createNodeSubscriber.Subscribe(uc.consumeCreateNodesEvent)
	return uc, err
}

func (uc *NodeUsecase) CreateNodes(ctx context.Context, warehouseId, start, end, auxAix int, nodeTupleType NodeTupleType) error {
	warehouse, err := uc.warehouseRepo.GetWarehouseById(ctx, warehouseId)
	if err != nil {
		return err
	}
	if data.IsDeleted(warehouse.IsDeleted) {
		return dataerrors.NewInvalidArgumentError("given Warehouse: [%d] not found", warehouseId)
	}
	if data.IsDisable(warehouse.Status) {
		return dataerrors.NewInvalidArgumentError("given Warehouse: [%d] is disable", warehouseId)
	}
	return uc.createNodePublisher.Publish(context.Background(), &publisher.CreateNodeEvent{
		WarehouseId:   warehouseId,
		Start:         start,
		End:           end,
		AuxAix:        auxAix,
		NodeTupleType: int(nodeTupleType),
	})
}

func (uc *NodeUsecase) GetNodeById(ctx context.Context, id int) (*Node, error) {
	return uc.repo.GetNodeById(ctx, id)
}

func (uc *NodeUsecase) GetNodesByIds(ctx context.Context, ids []int) ([]*Node, error) {
	return uc.repo.GetNodesByIds(ctx, ids)
}

func (uc *NodeUsecase) GetNodesByWarehouseId(ctx context.Context, warehouseId int) ([]*Node, error) {
	return uc.repo.GetNodesByWarehouseId(ctx, warehouseId)
}

func (uc *NodeUsecase) consumeCreateNodesEvent(ctx context.Context, msg []byte) error {
	var event *publisher.CreateNodeEvent
	if err := json.Unmarshal(msg, &event); err != nil {
		return err
	}
	for start := event.Start; start < event.End; {
		end := start + 100
		if end > event.End {
			end = event.End
		}
		nodeTupleType := NodeTupleType(event.NodeTupleType)
		nodeTuple := generateNodeTuple(event.WarehouseId, start, end, event.AuxAix, nodeTupleType)
		tuple, err := uc.repo.SelectNodeTupleByCoorRange(ctx, event.WarehouseId, start, end, event.AuxAix, nodeTupleType)
		if err != nil {
			return err
		}
		index := map[int]struct{}{}
		for _, node := range tuple {
			if nodeTupleType == NodeTupleTypeColumn {
				index[node.Y] = struct{}{}
			} else {
				index[node.X] = struct{}{}
			}
		}
		var nodes []*Node
		for _, node := range nodeTuple {
			if nodeTupleType == NodeTupleTypeColumn {
				if _, ok := index[node.Y]; ok {
					continue
				}
			} else {
				if _, ok := index[node.X]; ok {
					continue
				}
			}
			nodes = append(nodes, node)
		}
		start = end + 1
		if len(nodes) == 0 {
			continue
		}
		_, err = uc.repo.BatchCreateNode(ctx, nodes)
		if err != nil {
			return err
		}
	}
	return nil
}

func generateNodeTuple(warehouseId, start, end, auxAix int, tupleType NodeTupleType) NodeTuple {
	var tuple NodeTuple
	for aix := start; aix <= end; aix++ {
		node := &Node{WarehouseId: warehouseId}
		if tupleType == NodeTupleTypeColumn {
			node.X = auxAix
			node.Y = aix
		} else {
			node.X = aix
			node.Y = auxAix
		}
		tuple = append(tuple, node)
	}
	return tuple
}
