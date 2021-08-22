// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/qmdx00/crobjob/internal/worker/config"
	"github.com/qmdx00/crobjob/internal/worker/job"
	"github.com/qmdx00/crobjob/internal/worker/handler"
	"github.com/qmdx00/crobjob/internal/worker/log"
	"github.com/qmdx00/crobjob/internal/worker/server"
	"github.com/qmdx00/crobjob/pkg/lifecycle"
)

func initApp() (*lifecycle.App, func(), error) {
	panic(wire.Build(server.ProviderSet, handler.ProviderSet, job.ProviderSet, config.ProviderSet, log.ProviderSet, newApp))
}
