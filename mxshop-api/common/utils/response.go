package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"github.com/zhengpanone/mxshop/mxshop-api/common/errs"
	"github.com/zhengpanone/mxshop/mxshop-api/common/global"

	"net/http"
)

type Response struct {
	// 业务状态码
	Code int `json:"code"`
	// 提示信息
	Msg string `json:"msg"`
	// 响应数据
	Data interface{} `json:"data"`
	// Meta 源数据,存储如请求ID,分页等信息
	Meta Meta `json:"meta"`
}

// Meta 元数据
type Meta struct {
	RequestId string `json:"request_id"`
	// TODO 还可以集成分页信息
}

// ErrorItem 错误项
type ErrorItem struct {
	Key   string `json:"key"`
	Value string `json:"error"`
}

func New() *Response {
	requestId := uuid.NewV4()
	return &Response{
		Code: http.StatusOK,
		Msg:  "",
		Data: nil,
		Meta: Meta{
			RequestId: requestId.String(),
		},
	}
}

func ResultJson(ctx *gin.Context, code int, msg string, data interface{}) {
	response := New()
	response.Code = code
	response.Msg = msg
	response.Data = data
	ctx.JSON(http.StatusOK, response)
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
	ResultErrJson(ctx, global.OK, nil)
}

func OkWithMsg(ctx *gin.Context, msg string) {
	ResultJson(ctx, global.CodeSuccess, msg, nil)
}

func OkWithData(ctx *gin.Context, data interface{}) {
	ResultJson(ctx, global.CodeSuccess, "请求成功", data)
}

func OkWithDataAndMsg(ctx *gin.Context, data interface{}, msg string) {
	ResultJson(ctx, global.CodeSuccess, msg, data)
}

// ErrorWithMsg 错误信息
func ErrorWithMsg(ctx *gin.Context, msg string) {
	ResultJson(ctx, global.CodeSuccess, msg, nil)
}

// ErrorWithCodeAndMsg 错误码和信息
func ErrorWithCodeAndMsg(ctx *gin.Context, code int, msg string) {
	ResultJson(ctx, code, msg, nil)
}

// ErrorWithAppErr 错误码和信息
func ErrorWithAppErr(ctx *gin.Context, appErr *errs.AppError) {
	ResultErrJson(ctx, appErr, nil)
}

// ErrorWithCodeAndError 错误信息
func ErrorWithCodeAndError(ctx *gin.Context, code int, error error) {
	ResultJson(ctx, code, error.Error(), nil)
}
