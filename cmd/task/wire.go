// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/qmdx00/crobjob/internal/task/server"
	"github.com/qmdx00/crobjob/pkg/lifecycle"
)

func initApp() (*lifecycle.App, func(), error) {
	panic(wire.Build(server.ProviderSet, newApp))
}
