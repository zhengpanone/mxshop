package global

import (
	"github.com/zhengpanone/mxshop/mxshop-api/common/errs"
	"net/http"
)

// 常量定义
const (
	CodeSuccess       = http.StatusOK // 成功
	CodeERROR         = -1
	CodeBadRequest    = 400  // 请求参数错误
	CodeUnauthorized  = 401  // 未授权
	CodeForbidden     = 403  // 禁止访问
	CodeNotFound      = 404  // 资源不存在
	CodeServerError   = 500  // 服务器错误
	CodeDBError       = 600  // 数据库操作失败
	CodeBusinessError = 700  // 业务逻辑错误
	CodeTokenInvalid  = 1401 // token无效
	CodeTokenExpired  = 1402 // token过期
)

var (
	OK               = errs.New(CodeSuccess, "操作成功")
	ERR              = errs.New(CodeERROR, "操作成功")
	ErrBadRequest    = errs.New(CodeBadRequest, "请求参数错误")
	ErrUnauthorized  = errs.New(CodeUnauthorized, "未授权访问")
	ErrForbidden     = errs.New(CodeForbidden, "禁止访问")
	ErrNotFound      = errs.New(CodeNotFound, "资源未找到")
	ErrInternal      = errs.New(CodeServerError, "服务器内部错误")
	ErrDBError       = errs.New(CodeDBError, "数据库操作失败")
	ErrBusinessError = errs.New(CodeBusinessError, "业务逻辑错误")
	ErrTokenInvalid  = errs.New(CodeTokenInvalid, "token无效")
	ErrTokenExpired  = errs.New(CodeTokenExpired, "token过期")
)
