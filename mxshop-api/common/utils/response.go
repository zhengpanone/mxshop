package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
)

const (
	SUCCESS      = http.StatusOK
	ERROR        = -1
	TOKEN_EXPIRE = -2
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
	// Errors 错误提示，如 xx字段不能为空等
	Errors []ErrorItem `json:"errors"`
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
		Errors: []ErrorItem{},
	}
}

func ResultJson(ctx *gin.Context, code int, msg string, errors []ErrorItem, data interface{}) {
	response := New()
	response.Code = code
	response.Msg = msg
	response.Errors = errors
	response.Data = data
	ctx.JSON(http.StatusOK, response)
	return
}

func Ok(ctx *gin.Context) {
	ResultJson(ctx, SUCCESS, "请求成功", nil, map[string]interface{}{})
}

func OkWithMsg(ctx *gin.Context, msg string) {
	ResultJson(ctx, SUCCESS, msg, nil, nil)
}

func OkWithData(ctx *gin.Context, data interface{}) {
	ResultJson(ctx, SUCCESS, "请求成功", nil, data)
}

func OkWithDataAndMsg(ctx *gin.Context, data interface{}, msg string) {
	ResultJson(ctx, SUCCESS, msg, nil, data)
}

// ErrorWithMsg 错误信息
func ErrorWithMsg(ctx *gin.Context, msg string) {
	ResultJson(ctx, ERROR, msg, nil, map[string]interface{}{})
}

// ErrorWithCodeAndMsg 错误码和信息
func ErrorWithCodeAndMsg(ctx *gin.Context, code int, msg string) {
	ResultJson(ctx, code, msg, nil, map[string]interface{}{})
}

// ErrorWithMsgAndErrors 错误信息
func ErrorWithMsgAndErrors(ctx *gin.Context, msg string, errors []ErrorItem) {
	ResultJson(ctx, ERROR, msg, errors, map[string]interface{}{})
}

func ErrorWithToken(ctx *gin.Context, msg string) {
	ResultJson(ctx, TOKEN_EXPIRE, msg, nil, map[string]interface{}{})
}

func NoAuth(ctx *gin.Context, message string) {
	ResultJson(ctx, http.StatusUnauthorized, message, nil, nil)
}
