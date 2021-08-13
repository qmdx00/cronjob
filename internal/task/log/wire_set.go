package log

import (
	"github.com/google/wire"
	"github.com/qmdx00/crobjob/pkg/log"
)

var ProviderSet = wire.NewSet(log.NewHelper, NewLogger)
