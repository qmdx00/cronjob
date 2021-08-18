package main

import (
	"github.com/qmdx00/crobjob/pkg/lifecycle"
	"github.com/qmdx00/crobjob/pkg/transport"
)

func newApp(server transport.Server) *lifecycle.App {
	return lifecycle.New(
		lifecycle.WithName("cronjob worker"),
		lifecycle.WithVersion("1.0"),
		lifecycle.WithMetadata(map[string]string{}),
		lifecycle.WithServer(server))
}

func main() {
	app, cleanup, err := initApp()
	defer cleanup()

	if err != nil {
		panic(err)
	}

	if err = app.Run(); err != nil {
		panic(err)
	}
}
