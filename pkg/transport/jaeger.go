package transport

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	tracelog "github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"io"
)

type MDCarrier struct {
	metadata.MD
}

func (m MDCarrier) Set(key, val string) {
	m.MD[key] = append(m.MD[key], val)
}

func (m MDCarrier) ForeachKey(handler func(key, val string) error) error {
	for k, strs := range m.MD {
		for _, v := range strs {
			if err := handler(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}

// JaegerServerOption ...
func JaegerServerOption(tracer opentracing.Tracer) grpc.ServerOption {
	return grpc.UnaryInterceptor(serverInterceptor(tracer))
}

// JaegerClientOption ...
func JaegerClientOption(tracer opentracing.Tracer) grpc.DialOption {
	return grpc.WithUnaryInterceptor(clientInterceptor(tracer))
}

// NewJaegerTracer ...
func NewJaegerTracer(service string, agentEndpoint string) (opentracing.Tracer, io.Closer, error) {
	cfg := config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: agentEndpoint,
		},
	}

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		return nil, nil, err
	}

	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, nil
}

// serverInterceptor for grpc server ...
func serverInterceptor(tracer opentracing.Tracer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.Pairs()
		}
		spanContext, err := tracer.Extract(
			opentracing.TextMap,
			MDCarrier{md},
		)

		if err != nil && err != opentracing.ErrSpanContextNotFound {
			grpclog.Errorf("extract from metadata err: %v", err)
		} else {
			span := tracer.StartSpan(
				info.FullMethod,
				ext.RPCServerOption(spanContext),
				opentracing.Tag{Key: string(ext.Component), Value: "gRPC Server"},
				ext.SpanKindRPCServer,
			)
			defer span.Finish()

			ctx = opentracing.ContextWithSpan(ctx, span)
		}

		return handler(ctx, req)

	}
}

// clientInterceptor for grpc client ...
func clientInterceptor(tracer opentracing.Tracer) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		var err error
		var parentCtx opentracing.SpanContext
		parentSpan := opentracing.SpanFromContext(ctx)
		if parentSpan != nil {
			parentCtx = parentSpan.Context()
		}
		span := tracer.StartSpan(
			method,
			opentracing.ChildOf(parentCtx),
			opentracing.Tag{Key: string(ext.Component), Value: "gRPC Client"},
			ext.SpanKindRPCClient,
		)
		defer span.Finish()

		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.Pairs()
		} else {
			md = md.Copy()
		}

		if err = tracer.Inject(
			span.Context(),
			opentracing.TextMap,
			MDCarrier{md},
		); err != nil {
			tracelog.String("inject error", err.Error())
		}

		newCtx := metadata.NewOutgoingContext(ctx, md)
		if err = invoker(newCtx, method, req, reply, cc, opts...); err != nil {
			tracelog.String("call error", err.Error())
		}

		return err
	}
}
