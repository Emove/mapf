package publisher

import "github.com/google/wire"

// ProviderSet is Publisher providers.
var ProviderSet = wire.NewSet(NewCreateNodeEventPublisher)
