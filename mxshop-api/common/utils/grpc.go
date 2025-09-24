package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/zhengpanone/mxshop/mxshop-api/common/response"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strings"
)

func removeTopStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fields {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func HandleValidatorError(ctx *gin.Context, translator ut.Translator, err error) {
	// 如何返回错误信息
	var errs validator.ValidationErrors
	ok := errors.As(err, &errs)
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(errs.Translate(translator)),
	})
	return
}

func HandleGrpcErrorToHttp(err error, c *gin.Context, srvName string) {
	// 将grpc的code转换成http的code
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				response.ErrorWithCodeAndMsg(c, http.StatusNotFound, srvName+":"+e.Message())
				return
			case codes.Internal:
				response.ErrorWithCodeAndMsg(c, http.StatusInternalServerError, srvName+":"+"内部错误"+e.Message())
				return
			case codes.InvalidArgument:
				response.ErrorWithCodeAndMsg(c, http.StatusBadRequest, srvName+":"+"参数错误"+e.Message())
				return
			case codes.Unavailable:
				response.ErrorWithCodeAndMsg(c, http.StatusInternalServerError, srvName+":"+"不可用")
				return
			default:
				response.ErrorWithCodeAndMsg(c, http.StatusInternalServerError, srvName+":"+"其他错误"+e.Message())
				return
			}
		}
	}
}
