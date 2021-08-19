// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/qmdx00/crobjob/internal/task/biz"
	"github.com/qmdx00/crobjob/internal/task/config"
	"github.com/qmdx00/crobjob/internal/task/log"
	"github.com/qmdx00/crobjob/internal/task/producer"
	"github.com/qmdx00/crobjob/internal/task/server"
	"github.com/qmdx00/crobjob/pkg/lifecycle"
)

func initApp() (*lifecycle.App, func(), error) {
	panic(wire.Build(server.ProviderSet, biz.ProviderSet, producer.ProviderSet, config.ProviderSet, log.ProviderSet, newApp))
}
