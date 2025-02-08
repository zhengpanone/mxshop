package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/zhengpanone/mxshop/common/utils"
	"github.com/zhengpanone/mxshop/goods-web/global"
)

// https://github.dev/lixd/i-go/tree/master/apm/trace/gin
// https://www.lixueduan.com/posts/tracing/04-jaeger-gin-grpc/
func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		var parentSpan opentracing.Span

		tracer, closer := utils.InitJaeger(global.ServerConfig.Name, global.ServerConfig.Jaeger.Host, global.ServerConfig.Jaeger.Port)

		defer closer.Close()

		// 直接从 c.Request.Header 中提取 span,如果没有就新建一个
		spCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		if err != nil {
			parentSpan = tracer.StartSpan(c.Request.URL.Path)
			defer parentSpan.Finish()
		} else {
			parentSpan = opentracing.StartSpan(
				c.Request.URL.Path,
				opentracing.ChildOf(spCtx),
				opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
				ext.SpanKindRPCServer,
			)
			defer parentSpan.Finish()
		}
		// 然后存到 g.ctx 中 供后续使用
		c.Set("tracer", tracer)
		c.Set("ctx", opentracing.ContextWithSpan(context.Background(), parentSpan))
		c.Next()

	}
}
