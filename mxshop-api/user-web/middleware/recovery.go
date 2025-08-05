package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/zhengpanone/mxshop/mxshop-api/user-web/global"
	"go.uber.org/zap"
	"net/http"
)

func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if errs, ok := recovered.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"data": errs.Translate(global.Trans),
				"msg":  "入参校验失败",
			})
			return
		}
		if err, ok := recovered.(error); ok {
			zap.S().Errorw(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"data": err.Error(),
				"msg":  "入参校验失败",
			})
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	})
}
