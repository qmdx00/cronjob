package middleware

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
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
