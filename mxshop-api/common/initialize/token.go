package initialize

import (
	"github.com/go-redis/redis/v8"
	"github.com/zhengpanone/mxshop/mxshop-api/common/claims"
	commonConfig "github.com/zhengpanone/mxshop/mxshop-api/common/config"
	commonGlobal "github.com/zhengpanone/mxshop/mxshop-api/common/global"
	"time"
)

func InitTokenManager(jwtConfig commonConfig.JWTConfig, rdb *redis.Client) {
	if rdb == nil {
		commonGlobal.Logger.Error("redis client 不能为空,请先初始化 redis")
		panic("redis client 不能为空")

	}
	var tokenManager *claims.TokenManager
	if jwtConfig.IsMulti {
		// 多终端登录，最多允许 3 个 token
		tokenManager = claims.NewTokenManager(rdb, 24*time.Hour, 3)
	} else {
		// 单点登录
		tokenManager = claims.NewTokenManager(rdb, 24*time.Hour, 1)
	}

	commonGlobal.TokenManager = tokenManager

	//// 保存 token
	//_ = multiTokenMgr.SaveToken(1001, "token-abc")
	//_ = multiTokenMgr.SaveToken(1001, "token-def")
	//
	//// 校验 token
	//ok, _ := multiTokenMgr.ValidateToken(1001, "token-abc")
	//fmt.Println("校验结果:", ok)
	//
	//// 删除 token
	//_ = multiTokenMgr.RevokeToken(1001, "token-abc")

}
