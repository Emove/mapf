package event

import "context"

type Publisher interface {
	Publish(ctx context.Context, msg interface{}) error
}
