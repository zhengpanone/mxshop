package middleware

import (
	"github.com/gin-gonic/gin"
	commonClaims "github.com/zhengpanone/mxshop/common/claims"
	"net/http"
)

func IsAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, _ := ctx.Get("claims")
		customClaims := claims.(*commonClaims.CustomClaims)
		if customClaims.AuthorityId != 2 {
			ctx.JSON(http.StatusForbidden, gin.H{"msg": "无权限"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
