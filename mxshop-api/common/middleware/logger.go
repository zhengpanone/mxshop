package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

// CustomResponseWriter 自定义的ResponseWriter 用于获取响应数据
type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write
func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w CustomResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// GinLogger 日志中间件
func GinLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		// 跳过 /health 路径的日志记录
		if path == "/health" {
			ctx.Next()
			return
		}

		query := ctx.Request.URL.RawQuery
		logger.Info(fmt.Sprintf("请求开始：%s", path),
			zap.String("method", ctx.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", ctx.ClientIP()),
			zap.String("user-agent", ctx.Request.UserAgent()),
		)
		blw := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = blw

		ctx.Set("example", "123456")
		// 让原本执行的逻辑继续执行
		ctx.Next()
		end := time.Since(start)
		logger.Info(fmt.Sprintf("请求结束：%s\n 耗时：%v\n", path, end),
			zap.Int("status", ctx.Writer.Status()),
			zap.String("path", path),
			zap.String("response", blw.body.String()),
			zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Int64("cost", time.Since(start).Milliseconds()),
		)

		status := ctx.Writer.Status()
		fmt.Printf("状态：%v\n", status)
	}
}

func GinRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			// 执行recover() 捕获异常
			if err := recover(); err != nil {
				//检查是否有断开的连接，因为这并不是一个真正需要进行紧急堆栈跟踪的条件。
				var brokenPipe bool
				// 如果异常时net.OpError类型 ==> 转为 对应类型 err.() 类型断言的写法, 判断类型, 转换类型
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				// 将http请求转换成字节切面, 第二个布尔类型的参数表示是否包括请求体, 为false的话表示只包括请求头
				httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
				if brokenPipe {
					logger.Error(ctx.Request.URL.Path, zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					ctx.Error(err.(error))
					ctx.Abort()
					return
				}
				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				ctx.AbortWithStatus(http.StatusInternalServerError)

			}
		}()
		ctx.Next()
	}
}
