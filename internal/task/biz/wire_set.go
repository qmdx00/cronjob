package biz

import (
	"github.com/google/wire"
)

// ProviderSet for business ...
var ProviderSet = wire.NewSet(NewTaskBusiness)
