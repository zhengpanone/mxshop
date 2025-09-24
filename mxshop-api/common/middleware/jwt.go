package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zhengpanone/mxshop/mxshop-api/common/claims"
	commonGlobal "github.com/zhengpanone/mxshop/mxshop-api/common/global"
	commonResponse "github.com/zhengpanone/mxshop/mxshop-api/common/response"
	"net/http"
	"strings"
)

func JWTAuth(signingKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 这里jwt鉴权取头部信息x-token，登录时返回token信息
		token := ExtractToken(c)
		if token == "" {
			commonResponse.ErrorWithAppErr(c, commonGlobal.ErrUnauthorized)
			c.Abort()
			return
		}

		j := NewJWT(signingKey)
		// parseToken解析token包含的信息
		tokenClaims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, commonGlobal.ErrTokenExpired) {
				c.JSON(http.StatusUnauthorized, map[string]string{
					"msg": "授权已过期",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusUnauthorized, "未登录")
			c.Abort()
			return
		}

		// 检查 Token 是否在黑名单中
		result, err := commonGlobal.TokenManager.ValidateToken(tokenClaims.ID, token)
		if !result {
			c.JSON(http.StatusUnauthorized, "token已经失效")
			c.Abort()
			return
		}
		// 将用户信息存储到上下文
		c.Set("claims", tokenClaims)
		c.Set("userId", tokenClaims.ID)
		c.Set("username", tokenClaims.NickName)
		c.Next()
	}
}

func ExtractToken(c *gin.Context) string {
	// 1. 自定义 header: x-token
	if token := c.GetHeader("x-token"); token != "" {
		return token
	}

	// 2. 标准 Authorization header: Bearer <token>
	if authHeader := c.GetHeader("Authorization"); authHeader != "" {
		if strings.HasPrefix(authHeader, "Bearer ") {
			return strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
		}
		return authHeader // fallback: return whole header if no Bearer
	}

	// 3. Cookie
	if cookieToken, err := c.Cookie("token"); err == nil && cookieToken != "" {
		return cookieToken
	}

	// 4. Query 参数 ?token=xxx
	if queryToken := c.Query("token"); queryToken != "" {
		return queryToken
	}
	return ""
}

type JWT struct {
	SigningKey []byte
}

func NewJWT(signingKey string) *JWT {
	return &JWT{
		[]byte(signingKey), // 可以设置过期时间
	}

}

// CreateToken 创建一个token
func (j *JWT) CreateToken(claims claims.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(j.SigningKey)
	if err == nil {
		_ = commonGlobal.TokenManager.SaveToken(claims.ID, tokenString)
	}
	return tokenString, err
}

func (j *JWT) ParseToken(tokenString string) (*claims.CustomClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &claims.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if token != nil {
		if clams, ok := token.Claims.(*claims.CustomClaims); ok && token.Valid {
			return clams, nil
		}
		return nil, commonGlobal.ErrTokenInvalid
	} else {
		return nil, commonGlobal.ErrTokenInvalid
	}

}
