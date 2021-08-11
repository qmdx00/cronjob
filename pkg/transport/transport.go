package transport

import "context"

// Server ...
type Server interface {
	Start(context.Context) error
	Stop(context.Context) error
}
