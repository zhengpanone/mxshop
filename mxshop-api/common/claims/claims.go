package claims

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	ID          uint
	NickName    string
	AuthorityId uint // 1普通用户、2管理员
	jwt.RegisteredClaims
}
