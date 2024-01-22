package subscriber

import "github.com/google/wire"

// ProviderSet is Subscriber providers.
var ProviderSet = wire.NewSet(NewCreateNodeEventSubscriber)
