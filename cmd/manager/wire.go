// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/qmdx00/crobjob/internal/manager/biz"
	"github.com/qmdx00/crobjob/internal/manager/log"
	"github.com/qmdx00/crobjob/internal/manager/server"
	"github.com/qmdx00/crobjob/pkg/lifecycle"
)

func initApp() (*lifecycle.App, func(), error) {
	panic(wire.Build(server.ProviderSet, biz.ProviderSet, log.ProviderSet, newApp))
}
