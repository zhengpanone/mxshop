package main

import (
	"fmt"
	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
	"time"
)

func main() {
	cfg := &jaegerConfig.Configuration{
		ServiceName: "mxshop",
		// 采样配置
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: fmt.Sprintf("%s:%d", "127.0.0.1", 6831),
		},
	}
	tracer, closer, err := cfg.NewTracer(jaegerConfig.Logger(jaeger.StdLogger))
	defer closer.Close()
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	span := tracer.StartSpan("go-grpc-web")
	time.Sleep(time.Second)
	defer span.Finish()
}

/*package main

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"log"
	"time"
)

func InitJaeger(serviceName string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "127.0.0.1:6831",
		},
	}

	tracer, closer, errs := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if errs != nil {
		log.Fatalf("Could not initialize jaeger tracer: %s", errs.Error())
	}

	opentracing.SetGlobalTracer(tracer)
	return tracer, closer
}

func main() {
	tracer, closer := InitJaeger("my-golang-app")
	defer closer.Close()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		span := tracer.StartSpan("ping-handler")
		defer span.Finish()

		time.Sleep(1 * time.Second) // 模拟业务逻辑处理
		c.String(200, "pong")
	})

	r.Run(":8080")
}*/
