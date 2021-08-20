package handler

import (
	"github.com/google/wire"
	"github.com/qmdx00/crobjob/internal/worker/server"
)

// ProviderSet for handler ...
var ProviderSet = wire.NewSet(server.NewTaskConsumer, NewTaskHandler)
