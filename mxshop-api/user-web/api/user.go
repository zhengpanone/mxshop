package api

import (
	"common/claims"
	commonMiddleware "common/middleware"
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

// GetUserList 获取用户列表
// @Summary 用户列表
// @Description 获取用户列表
// @Tags  用户管理
// @Accept  json
// @Produce json
// @Param x-token header string true "token令牌"
// @Param page query int true "页码" default(1)
// @Param size query int true "页面大小" default(10)
// @success 200  {array} response.UserResponse
// @Router  /v1/user/list [get]
// @ID /v1/user/list
func GetUserList(ctx *gin.Context) {

	page, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	rsp, err := global.UserSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		PageNum:  uint32(page),
		PageSize: uint32(size),
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
	reMap := map[string]interface{}{
		"total": rsp.Total,
		"data":  result,
	}
	utils.OkWithData(ctx, reMap)
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
			utils.ErrorWithCodeMsg(ctx, http.StatusBadRequest, "验证码不正确")
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
			utils.ErrorWithCodeMsg(ctx, http.StatusInternalServerError, "登录失败")
			return
		} else {
			if passRsp.Success {
				// 生成Token
				j := commonMiddleware.NewJWT(global.ServerConfig.JWTInfo.SigningKey)
				roleId, _ := strconv.Atoi(rsp.Role)
				claims := claims.CustomClaims{
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
					utils.ErrorWithCodeMsg(ctx, http.StatusInternalServerError, "生成token失败")
					return
				}
				utils.OkWithData(ctx, gin.H{
					"id":         rsp.Id,
					"nick_name":  rsp.Nickname,
					"token":      token,
					"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
				})
				return
			} else {
				utils.ErrorWithCodeMsg(ctx, http.StatusBadRequest, "登录失败")
				return
			}
		}
	}

}

// Register 用户注册
// @Summary 用户注册
// @Description 用户注册
// @Tags  用户管理
// @Accept  json
// @Produce json
// @Param  request body  forms.RegisterForm true "请求参数"
// @success 200  {object} utils.Response{data=interface{}}
// @Router  /v1/user/register [post]
func Register(ctx *gin.Context) {
	registerForm := forms.RegisterForm{}
	if err := ctx.ShouldBind(&registerForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	// 验证码校验
	code, err := global.RedisClient.Get(context.Background(), registerForm.Mobile).Result()
	if err != nil {
		utils.ErrorWithCodeMsg(ctx, http.StatusInternalServerError, "服务器内部错误"+err.Error())
		return
	}
	if code != registerForm.Code {
		utils.ErrorWithCodeMsg(ctx, http.StatusBadRequest, "验证码不正确")
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
	j := commonMiddleware.NewJWT(global.ServerConfig.JWTInfo.SigningKey)
	roleId, _ := strconv.Atoi(user.Role)
	claims := claims.CustomClaims{
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
		utils.ErrorWithCodeMsg(ctx, http.StatusInternalServerError, "生成token失败")
		return
	}
	data := gin.H{
		"id":         user.Id,
		"nick_name":  user.Nickname,
		"token":      token,
		"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
	}
	utils.OkWithData(ctx, data)
}

func Ping(ctx *gin.Context) {
	utils.Ok(ctx)
}

func HandleValidatorError(ctx *gin.Context, err error) {
	// 如何返回错误信息
	var errs validator.ValidationErrors
	ok := errors.As(err, &errs)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(errs.Translate(global.Trans)),
	})
	return
}
