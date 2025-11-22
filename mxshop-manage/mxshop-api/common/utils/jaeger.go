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
		// 采样配置 将采样频率设置为 1，每一个 span 都记录，方便查看测试结果
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,                             // 开启本地 `span` 日志
			LocalAgentHostPort: fmt.Sprintf("%s:%d", host, port), // Jaeger agent address
		},
	}

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger)) // 使用 NullLogger 以避免 `span` 频繁打印
	if err != nil {
		log.Fatalf("Could not initialize jaeger tracer: %s", err.Error())
	}
	return tracer, closer
}
