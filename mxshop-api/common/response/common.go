package response

import (
	"github.com/gin-gonic/gin"
	"github.com/zhengpanone/mxshop/mxshop-api/common/errs"
	"github.com/zhengpanone/mxshop/mxshop-api/common/global"
	"net/http"
)

// Response  响应结构体
type Response struct {
	Code int         `json:"code"` // 业务状态码
	Msg  string      `json:"msg"`  // 提示信息
	Data interface{} `json:"data"` // 响应数据
}

func New() *Response {
	return &Response{
		Code: global.CodeSuccess,
		Msg:  "",
		Data: nil,
	}
}

func ResultJson(ctx *gin.Context, code int, msg string, data interface{}) {
	response := New()
	response.Code = code
	response.Msg = msg
	response.Data = data
	ctx.JSON(global.CodeSuccess, response)
	return
}

func ResultErrJson(ctx *gin.Context, appErr *errs.AppError, data interface{}) {
	response := New()
	response.Code = appErr.Code
	response.Msg = appErr.Message
	response.Data = data
	ctx.JSON(http.StatusOK, response)
	return
}

func Ok(ctx *gin.Context) {
	ResultJson(ctx, global.CodeSuccess, "请求成功", map[string]interface{}{})
}

func OkWithMsg(ctx *gin.Context, msg string) {
	ResultJson(ctx, global.CodeSuccess, msg, nil)
}

func OkWithData[T any](ctx *gin.Context, data T) {
	ResultJson(ctx, global.CodeSuccess, "请求成功", data)
}

// ErrorWithAppErr 错误码和信息
func ErrorWithAppErr(ctx *gin.Context, appErr *errs.AppError) {
	ResultErrJson(ctx, appErr, nil)
}

// ErrorWithCodeAndMsg 错误码和信息
func ErrorWithCodeAndMsg(ctx *gin.Context, code int, msg string) {
	ResultJson(ctx, code, msg, nil)
}

type PageResult[T any] struct {
	List     []T   `json:"list"`
	Total    int64 `json:"total"`
	Page     int   `json:"page"`
	PageSize int   `json:"pageSize"`
}
