package subscriber

import (
	"mapf/app/warehouse/internal/event"
	internalevent "mapf/internal/event"
)

func NewCreateNodeEventSubscriber(config internalevent.Configuration) (internalevent.Subscriber, func(), error) {
	return internalevent.NewRabbitMqSubscriber(config, event.CreateNodeExchange, event.CreateNodeKey, event.CreateNodeQueue)
}
