package config

import (
	"github.com/google/wire"
)

// ProviderSet for config ...
var ProviderSet = wire.NewSet(NewManagerConfig)
