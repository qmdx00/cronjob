package lifecycle

import (
	"context"
	"github.com/qmdx00/crobjob/pkg/transport"
	"os"
)

// Option application option.
type Option func(*options)

// options application option struct.
type options struct {
	id       string
	name     string
	version  string
	metadata map[string]string

	ctx  context.Context
	sigs []os.Signal

	servers []transport.Server
}

// WithID with service id.
func WithID(id string) Option {
	return func(o *options) { o.id = id }
}

// WithName with service name.
func WithName(name string) Option {
	return func(o *options) { o.name = name }
}

// WithVersion with service version.
func WithVersion(version string) Option {
	return func(o *options) { o.version = version }
}

// WithMetadata with service metadata.
func WithMetadata(md map[string]string) Option {
	return func(o *options) { o.metadata = md }
}

// WithContext with service context.
func WithContext(ctx context.Context) Option {
	return func(o *options) { o.ctx = ctx }
}

// WithSignal with exit signals.
func WithSignal(sigs ...os.Signal) Option {
	return func(o *options) { o.sigs = sigs }
}

// WithServer with transport servers.
func WithServer(srv ...transport.Server) Option {
	return func(o *options) { o.servers = srv }
}
