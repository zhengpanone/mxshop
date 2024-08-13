package utils

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"log"
)

func InitJaeger(serviceName string, host string, port uint32) (opentracing.Tracer, io.Closer) {

	// 参数详解 https://www.jaegertracing.io/docs/1.20/sampling/
	cfg := &config.Configuration{
		ServiceName: serviceName,
		// 采样配置
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: fmt.Sprintf("%s:%d", host, port), // Jaeger's agent address
		},
	}

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		log.Fatalf("Could not initialize jaeger tracer: %s", err.Error())
	}

	opentracing.SetGlobalTracer(tracer)
	return tracer, closer
}
