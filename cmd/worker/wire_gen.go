// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/qmdx00/crobjob/internal/worker/config"
	"github.com/qmdx00/crobjob/internal/worker/handler"
	"github.com/qmdx00/crobjob/internal/worker/job"
	"github.com/qmdx00/crobjob/internal/worker/log"
	"github.com/qmdx00/crobjob/internal/worker/server"
	"github.com/qmdx00/crobjob/pkg/lifecycle"
)

// Injectors from wire.go:

func initApp() (*lifecycle.App, func(), error) {
	workerConfig := config.NewWorkerConfig()
	logger := log.NewWorkerLogger(workerConfig)
	rootJob := job.NewRootJob(logger)
	receive := handler.NewTaskHandler(logger)
	v, err := server.NewServers(rootJob, workerConfig, logger, receive)
	if err != nil {
		return nil, nil, err
	}
	app := newApp(v)
	return app, func() {
	}, nil
}
