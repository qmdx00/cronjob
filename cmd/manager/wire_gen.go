// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/qmdx00/crobjob/internal/manager/biz"
	"github.com/qmdx00/crobjob/internal/manager/log"
	"github.com/qmdx00/crobjob/internal/manager/server"
	"github.com/qmdx00/crobjob/pkg/lifecycle"
)

// Injectors from wire.go:

func initApp() (*lifecycle.App, func(), error) {
	logger := log.NewManagerLogger()
	clientConn, err := biz.NewGRPCConn()
	if err != nil {
		return nil, nil, err
	}
	taskServiceClient := biz.NewTaskServiceClient(clientConn)
	taskBusiness := biz.NewTaskBusiness(logger, taskServiceClient)
	engine := server.NewHTTPRouter(logger, taskBusiness)
	transportServer := server.NewHttpServer(logger, engine)
	app := newApp(transportServer)
	return app, func() {
	}, nil
}