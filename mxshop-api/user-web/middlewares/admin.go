package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user-web/models"
)

func IsAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, _ := ctx.Get("claims")
		customClaims := claims.(*models.CustomClaims)
		if customClaims.AuthorityId != 2 {
			ctx.JSON(http.StatusForbidden, gin.H{"msg": "无权限"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
