package transport

import (
	"context"
	"errors"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	tracelog "github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
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

// JaegerOption ...
func JaegerOption(tracer opentracing.Tracer, log *zap.Logger) grpc.ServerOption {
	return grpc.UnaryInterceptor(serverInterceptor(tracer, log))
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

// serverInterceptor ...
func serverInterceptor(tracer opentracing.Tracer, log *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var parentContext context.Context

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.Pairs()
		}

		spanContext, err := tracer.Extract(opentracing.TextMap, MDCarrier{md})
		if err != nil && !errors.Is(err, opentracing.ErrSpanContextNotFound) {
			log.Error("extract from metadata err", zap.Error(err))
		} else {
			span := tracer.StartSpan(
				info.FullMethod,
				ext.RPCServerOption(spanContext),
				opentracing.Tag{Key: string(ext.Component), Value: "gRPC server"},
				ext.SpanKindRPCServer,
			)
			defer span.Finish()

			parentContext = opentracing.ContextWithSpan(ctx, span)
		}

		return handler(parentContext, req)
	}
}

// ClientInterceptor for grpc client ...
func ClientInterceptor(_ context.Context, tracer opentracing.Tracer) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		span, _ := opentracing.StartSpanFromContext(
			ctx,
			"call gRPC",
			opentracing.Tag{Key: string(ext.Component), Value: "gRPC client"},
			ext.SpanKindRPCClient,
		)
		defer span.Finish()

		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.Pairs()
		} else {
			md = md.Copy()
		}

		mdCarrier := MDCarrier{md}
		err := tracer.Inject(span.Context(), opentracing.TextMap, mdCarrier)
		if err != nil {
			tracelog.String("inject error", err.Error())
			return err
		}

		newCtx := metadata.NewOutgoingContext(ctx, md)
		err = invoker(newCtx, method, req, reply, cc, opts...)
		if err != nil {
			tracelog.String("call error", err.Error())
			return err
		}
		return nil
	}
}
