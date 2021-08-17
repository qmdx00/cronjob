package middleware

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	tracelog "github.com/opentracing/opentracing-go/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
)


// JaegerServerOption ...
func JaegerServerOption(tracer opentracing.Tracer) grpc.ServerOption {
	return grpc.UnaryInterceptor(serverInterceptor(tracer))
}

// JaegerClientOption ...
func JaegerClientOption(tracer opentracing.Tracer) grpc.DialOption {
	return grpc.WithUnaryInterceptor(clientInterceptor(tracer))
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
