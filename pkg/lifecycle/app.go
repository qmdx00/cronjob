package lifecycle

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// AppInfo application context value.
type AppInfo interface {
	ID() string
	Name() string
	Version() string
	Metadata() map[string]string
}

// App application struct.
type App struct {
	opts   options
	ctx    context.Context
	cancel func()
}

// New return new App
func New(opts ...Option) *App {
	_opts := options{
		ctx:    context.Background(),
		sigs:   []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
	}
	if id, err := uuid.NewUUID(); err == nil {
		_opts.id = id.String()
	}
	for _, o := range opts {
		o(&_opts)
	}
	ctx, cancel := context.WithCancel(_opts.ctx)
	return &App{
		ctx:    ctx,
		cancel: cancel,
		opts:   _opts,
	}
}

// ID returns app instance id.
func (a *App) ID() string { return a.opts.id }

// Name returns service name.
func (a *App) Name() string { return a.opts.name }

// Version returns app version.
func (a *App) Version() string { return a.opts.version }

// Metadata returns service metadata.
func (a *App) Metadata() map[string]string { return a.opts.metadata }

// Run start application.
func (a *App) Run() error {
	var err error
	ctx := NewContext(a.ctx, a)
	group, ctx := errgroup.WithContext(ctx)
	wg := &sync.WaitGroup{}

	// start servers
	for _, srv := range a.opts.servers {
		group.Go(func() error {
			<-ctx.Done()
			return srv.Stop(ctx)
		})

		wg.Add(1)
		group.Go(func() error {
			wg.Done()
			return srv.Start(ctx)
		})
	}
	// waiting for all servers start
	wg.Wait()

	// wait for os signal
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, a.opts.sigs...)
	group.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-sigs:
				return a.Stop()
			}
		}
	})
	if err = group.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

// Stop gracefully stops the application.
func (a *App) Stop() error {
	// cancel context
	if a.cancel != nil {
		a.cancel()
	}
	return nil
}

// appKey app context key
type appKey struct{}

// NewContext returns a new Context that carries value.
func NewContext(ctx context.Context, s AppInfo) context.Context {
	return context.WithValue(ctx, appKey{}, s)
}

// FromContext returns the Transport value stored in ctx, if any.
func FromContext(ctx context.Context) (s AppInfo, ok bool) {
	s, ok = ctx.Value(appKey{}).(AppInfo)
	return
}
