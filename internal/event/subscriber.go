package event

import "context"

type Consumer func(ctx context.Context, message []byte) error

type Subscriber interface {
	Subscribe(consumer Consumer) error
}
