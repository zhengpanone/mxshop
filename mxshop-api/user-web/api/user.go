package api

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
	"time"
	"user-web/forms"
	"user-web/global"
	"user-web/global/response"
	"user-web/middlewares"
	"user-web/models"
	"user-web/proto"
	"user-web/utils"
)

func removeTopStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fields {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func GetUserList(ctx *gin.Context) {

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "10"))
	rsp, err := global.UserSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		Page: uint32(page),
		Size: uint32(size),
	})
	if err != nil {
		zap.S().Errorw("[GetUserList]查询【用户列表】失败")
		HandleGrpcErrorToHttp(err, ctx, "用户服务srv")
		return
	}
	/*claims, _ := ctx.Get("claims")
	customClaims := claims.(*models.CustomClaims)
	zap.S().Infof("当前访问用户：%d", customClaims.ID)*/
	//zap.S().Debugf("获取用户列表页")
	result := make([]response.UserResponse, 0)
	for _, value := range rsp.Data {
		user := response.UserResponse{
			Id:       value.Id,
			NickName: value.Nickname,
			Birthday: response.JsonTime(time.Unix(int64(value.Birthday), 0)),
			Gender:   value.Gender,
			Mobile:   value.Mobile,
		}
		result = append(result, user)
	}
	ctx.JSON(http.StatusOK, result)
}

// PasswordLogin
// @Summary 用户登录
// @Description 用户账号密码登录
// @Tags  用户管理
// @Accept  json
// @Produce json
// @Param  request body  forms.PasswordLoginForm true "请求参数"
// @success 200  {object} utils.Response{data=interface{}}
// @Router  /v1/user/pwd_login [post]
func PasswordLogin(ctx *gin.Context) {
	passwordLogin := forms.PasswordLoginForm{}
	if err := ctx.ShouldBind(&passwordLogin); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	if global.ServerConfig.EnableCaptcha {
		if !utils.VerifyCaptcha(passwordLogin.CaptchaId, passwordLogin.Captcha) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": "验证码不正确",
			})
			return
		}
	}

	// 登录的逻辑
	if rsp, err := global.UserSrvClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
		Mobile: passwordLogin.Mobile,
	}); err != nil {
		zap.S().Errorw("用户登录失败失败" + err.Error())
		HandleGrpcErrorToHttp(err, ctx, "用户srv")
		return
	} else {
		// 只是查询到用户，没有检查密码
		if passRsp, passErr := global.UserSrvClient.CheckPassword(context.Background(), &proto.PasswordCheckInfo{
			EncryptedPassword: rsp.Password,
			Password:          passwordLogin.Password,
		}); passErr != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]string{
				"msg": "登录失败",
			})
		} else {
			if passRsp.Success {
				// 生成Token
				j := middlewares.NewJWT()
				roleId, _ := strconv.Atoi(rsp.Role)
				claims := models.CustomClaims{
					ID:          uint(rsp.Id),
					NickName:    rsp.Nickname,
					AuthorityId: uint(roleId),
					RegisteredClaims: jwt.RegisteredClaims{
						NotBefore: jwt.NewNumericDate(time.Now()),                     //签名生效时间
						ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 签名过期时间
						IssuedAt:  jwt.NewNumericDate(time.Now()),
						Issuer:    "mxshop",
						Subject:   "user-identifier",
						Audience:  jwt.ClaimStrings{"your-audience"},
					},
				}
				token, err := j.CreateToken(claims)
				if err != nil {
					utils.ErrorWithCodeAndMsg(ctx, http.StatusInternalServerError, "生成token失败")
					/*ctx.JSON(http.StatusInternalServerError, gin.H{
						"msg": "生成token失败",
					})*/
					return
				}
				utils.OkWithData(ctx, gin.H{
					"id":         rsp.Id,
					"nick_name":  rsp.Nickname,
					"token":      token,
					"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
				})
				/*ctx.JSON(http.StatusOK, gin.H{
					"id":         rsp.Id,
					"nick_name":  rsp.Nickname,
					"token":      token,
					"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
				})*/
				return
			} else {
				utils.ErrorWithCodeAndMsg(ctx, http.StatusBadRequest, "登录失败")
				/*ctx.JSON(http.StatusBadRequest, map[string]string{
					"msg": "登录失败",
				})*/
				return
			}
		}
	}

}

// Register 用户注册
func Register(ctx *gin.Context) {
	registerForm := forms.RegisterForm{}
	if err := ctx.ShouldBind(&registerForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	// 验证码校验
	code, err := global.RedisClient.Get(context.Background(), registerForm.Mobile).Result()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误" + err.Error(),
		})
		return
	}
	if code != registerForm.Code {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码不正确",
		})
		return
	}
	user, err := global.UserSrvClient.CreateUser(context.Background(), &proto.CreateUserInfo{
		Mobile:   registerForm.Mobile,
		Nickname: registerForm.Mobile,
		Password: registerForm.Password,
	})
	if err != nil {
		zap.S().Errorf("[Register]新建用户失败：%s", err.Error())
		HandleGrpcErrorToHttp(err, ctx, "用户服务srv")
		return
	}

	// 生成Token
	j := middlewares.NewJWT()
	roleId, _ := strconv.Atoi(user.Role)
	claims := models.CustomClaims{
		ID:          uint(user.Id),
		NickName:    user.Nickname,
		AuthorityId: uint(roleId),
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),                     //签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 签名过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "mxshop",
			Subject:   "user-identifier",
			Audience:  jwt.ClaimStrings{"your-audience"},
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成token失败",
		})
		return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"id":         user.Id,
		"nick_name":  user.Nickname,
		"token":      token,
		"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
	})

}

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func HandleValidatorError(ctx *gin.Context, err error) {
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
		"error": removeTopStruct(errs.Translate(global.Trans)),
	})
	return
}
