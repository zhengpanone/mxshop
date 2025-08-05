package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zhengpanone/mxshop/mxshop-api/common/claims"
	"net/http"
	"strings"
	//"order-web/global"
	//"order-web/models"
)

func JWTAuth(signingKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 这里jwt鉴权取头部信息x-token，登录时返回token信息
		token := c.Request.Header.Get("x-token")
		if token == "" {
			// 尝试从 Authorization 中获取 token
			token = c.Request.Header.Get("Authorization")
		}
		if token == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"msg": "请登录",
			})
			c.Abort()
			return
		}
		if strings.Contains(token, "Bearer ") {
			token = strings.Split(token, " ")[1]
		}
		j := NewJWT(signingKey)
		// parseToken解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
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
		c.Set("claims", claims)
		c.Set("userId", claims.ID)
		c.Next()
	}
}

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token:")
)

func NewJWT(signingKey string) *JWT {

	return &JWT{
		[]byte(signingKey), // 可以设置过期时间
	}

}

// 创建一个token
func (j *JWT) CreateToken(claims claims.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
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
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}

}
