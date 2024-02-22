package biz

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"mapf/app/warehouse/internal/event/publisher"
	"mapf/app/warehouse/internal/event/subscriber"
	"mapf/internal/data"
	"mapf/internal/data/tx"
	dataerrors "mapf/internal/errors"
	internalevent "mapf/internal/event"
	"mapf/internal/utils"
	"strings"
)

type AffixNode struct {
	*data.Model
	WarehouseId int    `gorm:"index:idx_warehouse;not null"`
	NodeId      int    `gorm:"uniqueIndex:uidx_node_node_type_name;not null"`
	NodeTypeId  int    `gorm:"uniqueIndex:uidx_node_node_type_name;not null"`
	Name        string `gorm:"uniqueIndex:uidx_node_node_type_name;type:varchar(64);not null"`
	Remark      string `gorm:"type:varchar(255)"`
	Status      int    `gorm:"type:smallint;default:1;not null"`
}

func (*AffixNode) TableName() string {
	return "affix_node"
}

func (an *AffixNode) SetNodeId(nodeId int) {
	an.NodeId = nodeId
	an.Update("node_id", nodeId)
}

func (an *AffixNode) SetNodeTypeId(nodeTypeId int) {
	an.NodeTypeId = nodeTypeId
	an.Update("node_type_id", nodeTypeId)
}

func (an *AffixNode) SetName(name string) {
	an.Name = name
	an.Update("name", name)
}

func (an *AffixNode) SetRemark(remark string) {
	an.Remark = remark
	an.Update("remark", remark)
}

func (an *AffixNode) SetStatus(status int) {
	an.Status = status
	an.Update("status", status)
}

type AffixNodeRepo interface {
	// CreateAffixNode 创建用户节点
	CreateAffixNode(context.Context, *AffixNode) (*AffixNode, error)
	// BatchCreateAffixNode 批量创建用户节点
	BatchCreateAffixNode(context.Context, []*AffixNode) ([]*AffixNode, error)
	// UpdateAffixNode 更新用户节点
	UpdateAffixNode(context.Context, *AffixNode) (*AffixNode, error)
	// GetAffixNodeById 通过ID获取用户节点
	GetAffixNodeById(context.Context, int) (*AffixNode, error)
	// SelectAffixNodes 查询用户节点，可分页
	SelectAffixNodes(context.Context, *AffixNode, *data.Page) ([]*AffixNode, *data.Page, error)
}

type AffixNodeUsecase struct {
	tm                       tx.Transaction
	repo                     AffixNodeRepo
	warehouseRepo            WarehouseRepo
	nodeRepo                 NodeRepo
	nodeTypeRepo             NodeTypeRepo
	logger                   *log.Helper
	createAffixNodePublisher internalevent.Publisher
}

func NewAffixNodeUsecase(
	tm tx.Transaction,
	repo AffixNodeRepo,
	warehouseRepo WarehouseRepo,
	nodeRepo NodeRepo,
	nodeTypeRepo NodeTypeRepo,
	eventCenterConfig internalevent.Configuration,
	logger log.Logger,
) (*AffixNodeUsecase, error) {
	var err error
	var cleans []func()
	defer func() {
		if err != nil {
			for _, clean := range cleans {
				clean()
			}
		}
	}()
	createAffixNodePublisher, clean, err := publisher.NewCreateAffixNodeEventPublisher(eventCenterConfig)
	if err != nil {
		return nil, err
	}
	cleans = append(cleans, clean)
	uc := &AffixNodeUsecase{
		tm:                       tm,
		repo:                     repo,
		warehouseRepo:            warehouseRepo,
		nodeRepo:                 nodeRepo,
		nodeTypeRepo:             nodeTypeRepo,
		createAffixNodePublisher: createAffixNodePublisher,
		logger:                   log.NewHelper(logger),
	}
	createAffixNodeSubscriber, clean, err := subscriber.NewCreateAffixNodeEventSubscriber(eventCenterConfig)
	if err != nil {
		return nil, err
	}
	cleans = append(cleans, clean)
	err = createAffixNodeSubscriber.Subscribe(uc.consumeCreateAffixNodeEvent)
	return uc, err
}

func (uc *AffixNodeUsecase) CreateAffixNode(ctx context.Context, warehouseId, x, y, nodeTypeId int, name, remark string) (*AffixNode, error) {
	_, node, nodeType, err := ValidateWarehouseIdAndNodeAndNodeType(ctx, uc.warehouseRepo, uc.nodeRepo, uc.nodeTypeRepo, warehouseId, x, y, nodeTypeId)
	if err != nil {
		return nil, err
	}
	if !utils.NotBlank(name) {
		name = fmt.Sprintf("%s_%d", nodeType.Name, RoundUpAix(x, y))
	}
	affixNode := &AffixNode{
		WarehouseId: warehouseId,
		NodeId:      node.ID,
		NodeTypeId:  nodeTypeId,
		Name:        name,
		Remark:      remark,
	}
	return uc.repo.CreateAffixNode(ctx, affixNode)
}

func (uc *AffixNodeUsecase) CreateAffixNodes(ctx context.Context, warehouseId, nodeTypeId, start, end, auxAix int, nodeTupleType NodeTupleType) error {
	if err := ValidateWarehouseIdAndNodeType(ctx, uc.warehouseRepo, uc.nodeTypeRepo, warehouseId, nodeTypeId); err != nil {
		return nil
	}
	return uc.createAffixNodePublisher.Publish(context.Background(), &publisher.CreateAffixNodeEvent{
		WarehouseId:   warehouseId,
		NodeTypeId:    nodeTypeId,
		Start:         start,
		End:           end,
		AuxAix:        auxAix,
		NodeTupleType: int(nodeTupleType),
	})
}

func (uc *AffixNodeUsecase) UpdateAffixNode(ctx context.Context, affixNode *AffixNode) (*AffixNode, error) {
	originAffixNode, err := uc.repo.GetAffixNodeById(ctx, affixNode.ID)
	if err != nil {
		return nil, err
	}
	if utils.ValidInt(originAffixNode.NodeId) && originAffixNode.NodeId != affixNode.NodeId {
		node, err := uc.nodeRepo.GetNodeById(ctx, affixNode.NodeId)
		if err != nil {
			return nil, err
		}
		if data.IsDeleted(node.IsDeleted) {
			return nil, dataerrors.NewInvalidArgumentError("given Node: [%d] not found", affixNode.NodeId)
		}
	}
	if utils.ValidInt(originAffixNode.NodeTypeId) && originAffixNode.NodeTypeId != affixNode.NodeTypeId {
		nodeType, err := uc.nodeTypeRepo.GetNodeTypeById(ctx, affixNode.NodeTypeId)
		if err != nil {
			return nil, err
		}
		if data.IsDeleted(nodeType.IsDeleted) {
			return nil, dataerrors.NewInvalidArgumentError("given NodeType: [%d] not found", affixNode.NodeTypeId)
		}
	}
	if utils.NotBlank(affixNode.Name) && originAffixNode.Name != affixNode.Name {
		originAffixNode.SetName(strings.TrimSpace(affixNode.Name))
	}
	if utils.NotBlank(affixNode.Remark) && originAffixNode.Remark != affixNode.Remark {
		originAffixNode.SetRemark(affixNode.Remark)
	}
	return uc.repo.UpdateAffixNode(ctx, originAffixNode)
}

func (uc *AffixNodeUsecase) GetAffixNodeById(ctx context.Context, id int) (*AffixNode, error) {
	return uc.repo.GetAffixNodeById(ctx, id)
}

// TODO
// func (uc *AffixNodeUsecase) SelectAffixNode(ctx context.Context, )

func (uc *AffixNodeUsecase) consumeCreateAffixNodeEvent(ctx context.Context, msg []byte) error {
	var event *publisher.CreateAffixNodeEvent
	if err := json.Unmarshal(msg, &event); err != nil {
		return err
	}
	nodeType, _ := uc.nodeTypeRepo.GetNodeTypeById(ctx, event.NodeTypeId)
	for start := event.Start; start < event.End; {
		end := start + 100
		if end > event.End {
			end = event.End
		}
		nodeTupleType := NodeTupleType(event.NodeTupleType)
		tuple, err := uc.nodeRepo.SelectNodeTupleByCoorRange(ctx, event.WarehouseId, start, end, event.AuxAix, nodeTupleType)
		if err != nil {
			return err
		}
		var nodes []*AffixNode
		for _, node := range tuple {
			nodes = append(nodes, &AffixNode{
				WarehouseId: event.WarehouseId,
				NodeId:      node.ID,
				NodeTypeId:  event.NodeTypeId,
				Name:        fmt.Sprintf("%s_%d", nodeType.Name, RoundUpAix(node.X, node.Y)),
			})
		}
		start = end + 1
		if len(nodes) == 0 {
			continue
		}
		_, err = uc.repo.BatchCreateAffixNode(ctx, nodes)
		if err != nil {
			return err
		}

		// FIXME create node relation
	}
	return nil
}

func ValidateWarehouseIdAndNodeAndNodeType(
	ctx context.Context,
	warehouseRepo WarehouseRepo,
	nodeRepo NodeRepo,
	nodeTypeRepo NodeTypeRepo,
	warehouseId, x, y, nodeTypeId int,
) (*Warehouse, *Node, *NodeType, error) {
	warehouse, err := warehouseRepo.GetWarehouseById(ctx, warehouseId)
	if err != nil {
		return nil, nil, nil, err
	}
	if data.IsDeleted(warehouse.IsDeleted) {
		return nil, nil, nil, dataerrors.NewInvalidArgumentError("given Warehouse: [%d] not found", warehouseId)
	}
	if data.IsDisable(warehouse.Status) {
		return nil, nil, nil, dataerrors.NewInvalidArgumentError("given Warehouse: [%d] is disable", warehouseId)
	}
	node, err := nodeRepo.GetNodeByWarehouseAndCoordinate(ctx, warehouseId, x, y)
	if err != nil {
		return nil, nil, nil, err
	}
	if node == nil {
		return nil, nil, nil, dataerrors.NewResourceNotFoundError("given Node: (%d, %d) not found", x, y)
	}
	nodeType, err := nodeTypeRepo.GetNodeTypeById(ctx, nodeTypeId)
	if err != nil {
		return nil, nil, nil, err
	}
	if data.IsDeleted(nodeType.IsDeleted) {
		return nil, nil, nil, dataerrors.NewInvalidArgumentError("given NodeType: [%d] not found", nodeTypeId)
	}
	return warehouse, node, nodeType, nil
}

func ValidateWarehouseIdAndNodeType(ctx context.Context, warehouseRepo WarehouseRepo, nodeTypeRepo NodeTypeRepo, warehouseId, nodeTypeId int) error {
	warehouse, err := warehouseRepo.GetWarehouseById(ctx, warehouseId)
	if err != nil {
		return err
	}
	if data.IsDeleted(warehouse.IsDeleted) {
		return dataerrors.NewInvalidArgumentError("given Warehouse: [%d] not found", warehouseId)
	}
	if data.IsDisable(warehouse.Status) {
		return dataerrors.NewInvalidArgumentError("given Warehouse: [%d] is disable", warehouseId)
	}
	nodeType, err := nodeTypeRepo.GetNodeTypeById(ctx, nodeTypeId)
	if err != nil {
		return err
	}
	if data.IsDeleted(nodeType.IsDeleted) {
		return dataerrors.NewInvalidArgumentError("given NodeType: [%d] not found", nodeTypeId)
	}
	return nil
}
