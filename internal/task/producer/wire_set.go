package producer

import "github.com/google/wire"

// ProviderSet for producer ...
var ProviderSet = wire.NewSet(NewProducer)
