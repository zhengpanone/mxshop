package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGenerateUploadToken(t *testing.T) {
	// 设置 Gin 为测试模式
	gin.SetMode(gin.TestMode)
	tests := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "Valid Token",
			expectedStatus: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建一个模拟的 gin.Context
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			// 调用 GenerateUploadToken 函数
			GenerateUploadToken(ctx)

			// 检查状态码是否符合预期
			assert.Equal(t, tt.expectedStatus, w.Code)

			fmt.Println(w.Body.String())
			// 检查响应头是否正确
			assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
			assert.Equal(t, "POST", w.Header().Get("Access-Control-Allow-Methods"))
		})
	}
}
