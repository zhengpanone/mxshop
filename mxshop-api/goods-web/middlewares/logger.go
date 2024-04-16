package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func MyLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Set("example", "123456")
		// 让原本执行的逻辑继续执行
		ctx.Next()

		end := time.Since(start)
		fmt.Printf("耗时：%v\n", end)

		status := ctx.Writer.Status()
		fmt.Printf("状态：%v\n", status)
	}
}
