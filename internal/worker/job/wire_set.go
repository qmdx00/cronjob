package job

import "github.com/google/wire"

// ProviderSet for job ...
var ProviderSet = wire.NewSet(NewRootJob)
