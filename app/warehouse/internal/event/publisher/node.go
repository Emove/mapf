package publisher

import (
	"mapf/app/warehouse/internal/event"
	internalevent "mapf/internal/event"
)

// CreateNodeEvent 创建节点事件
type CreateNodeEvent struct {
	WarehouseId   int `json:"warehouse_id"`
	Start         int `json:"start"`
	End           int `json:"end"`
	AuxAix        int `json:"aux_aix"`
	NodeTupleType int `json:"node_tuple_type"`
}

func NewCreateNodeEventPublisher(config internalevent.Configuration) (internalevent.Publisher, func(), error) {
	return internalevent.NewRabbitMQPublisher(config, event.CreateNodeExchange, event.CreateNodeKey, event.CreateNodeQueue)
}
