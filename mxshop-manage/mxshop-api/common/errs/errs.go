package errs

import "fmt"

// AppError 定义错误结构体
type AppError struct {
	Code    int
	Message string
}

// 实现 error 接口
func (e *AppError) Error() string {
	return fmt.Sprintf("code=%d, msg=%s", e.Code, e.Message)
}

// New 快速构造
func New(code int, message string) *AppError {
	return &AppError{Code: code, Message: message}
}
