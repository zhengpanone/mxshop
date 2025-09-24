package controller

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zhengpanone/mxshop/mxshop-api/common/claims"
	commonGlobal "github.com/zhengpanone/mxshop/mxshop-api/common/global"
	commonMiddleware "github.com/zhengpanone/mxshop/mxshop-api/common/middleware"
	commonpb "github.com/zhengpanone/mxshop/mxshop-api/common/proto/pb"
	commonResponse "github.com/zhengpanone/mxshop/mxshop-api/common/response"
	commonUtils "github.com/zhengpanone/mxshop/mxshop-api/common/utils"
	"github.com/zhengpanone/mxshop/mxshop-api/user-web/forms"
	"github.com/zhengpanone/mxshop/mxshop-api/user-web/global"
	"github.com/zhengpanone/mxshop/mxshop-api/user-web/global/response"
	"github.com/zhengpanone/mxshop/mxshop-api/user-web/utils"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func removeTopStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fields {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

// GetUserList 获取用户列表
//
//	@Summary		用户列表
//	@Description	获取用户列表
//	@Tags			用户管理
//	@Accept			json
//	@Produce		json
//	@Param			x-token	header	string	true	"token令牌"
//	@Param			page	query	int		true	"页码"	default(1)
//	@Param			size	query	int		true	"页面大小"	default(10)
//	@success		200		{array}	response.UserResponse
//	@Router			/v1/user/list [get]
//	@ID				/v1/user/list
func GetUserList(ctx *gin.Context) {

	page, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	rsp, err := global.UserSrvClient.GetUserList(context.Background(), &commonpb.PageInfo{
		PageNum:  uint32(page),
		PageSize: uint32(size),
	})
	if err != nil {
		zap.S().Errorw("[GetUserList]查询【用户列表】失败")
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "用户服务srv")
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
	commonResponse.OkWithData(ctx, reMap)
}

// PasswordLogin
//
//	@Summary		用户登录
//	@Description	用户账号密码登录
//	@Tags			用户管理
//	@Accept			json
//	@Produce		json
//	@Param			request	body		forms.PasswordLoginForm	true	"请求参数"
//	@success		200		{object}	utils.Response{data=interface{}}
//	@Router			/v1/user/pwd_login [post]
func PasswordLogin(ctx *gin.Context) {
	passwordLogin := forms.PasswordLoginForm{}
	if err := ctx.ShouldBind(&passwordLogin); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	if global.ServerConfig.EnableCaptcha {
		if !utils.VerifyCaptcha(passwordLogin.CaptchaId, passwordLogin.Captcha) {
			commonResponse.ErrorWithCodeAndMsg(ctx, http.StatusBadRequest, "验证码不正确")
			return
		}
	}

	// 登录的逻辑
	if rsp, err := global.UserSrvClient.GetUserByMobile(context.Background(), &commonpb.MobileRequest{
		Mobile: passwordLogin.Mobile,
	}); err != nil {
		zap.S().Errorw("用户登录失败失败" + err.Error())
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "用户srv")
		return
	} else {
		// 只是查询到用户，没有检查密码
		if passRsp, passErr := global.UserSrvClient.CheckPassword(context.Background(), &commonpb.PasswordCheckInfo{
			EncryptedPassword: rsp.Password,
			Password:          passwordLogin.Password,
		}); passErr != nil {
			commonResponse.ErrorWithCodeAndMsg(ctx, http.StatusInternalServerError, "登录失败")
			return
		} else {
			if passRsp.Success {
				// 生成Token
				j := commonMiddleware.NewJWT(global.ServerConfig.JWTInfo.SigningKey)
				roleId, _ := strconv.Atoi(rsp.Role)
				customClaims := claims.CustomClaims{
					ID:          strconv.Itoa(int(rsp.Id)),
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
				token, err := j.CreateToken(customClaims)
				if err != nil {
					commonResponse.ErrorWithCodeAndMsg(ctx, http.StatusInternalServerError, "生成token失败")
					return
				}
				commonResponse.OkWithData(ctx, gin.H{
					"id":         rsp.Id,
					"nick_name":  rsp.Nickname,
					"token":      token,
					"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
				})
				return
			} else {
				commonResponse.ErrorWithCodeAndMsg(ctx, http.StatusBadRequest, "登录失败")
				return
			}
		}
	}

}

// LogOut 用户退出
func LogOut(ctx *gin.Context) {
	token := commonMiddleware.ExtractToken(ctx)
	if token == "" {
		commonResponse.ErrorWithCodeAndMsg(ctx, http.StatusBadRequest, "退出失败")
		return
	}
	j := commonMiddleware.NewJWT(global.ServerConfig.JWTInfo.SigningKey)
	tokenClaims, err := j.ParseToken(token)
	if err != nil {
		commonResponse.ErrorWithCodeAndMsg(ctx, http.StatusBadRequest, "token不合法")
		return
	}
	// 将 Token 加入黑名单
	_ = commonGlobal.TokenManager.BlacklistToken(token, tokenClaims.ExpiresAt.Time)
	commonResponse.OkWithMsg(ctx, "登出成功")
}

// Register 用户注册
//
//	@Summary		用户注册
//	@Description	用户注册
//	@Tags			用户管理
//	@Accept			json
//	@Produce		json
//	@Param			request	body		forms.RegisterForm	true	"请求参数"
//	@success		200		{object}	utils.Response{data=interface{}}
//	@Router			/v1/user/register [post]
func Register(ctx *gin.Context) {
	registerForm := forms.RegisterForm{}
	if err := ctx.ShouldBind(&registerForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	// 验证码校验
	code, err := commonGlobal.RedisClient.Get(context.Background(), registerForm.Mobile).Result()
	if err != nil {
		commonResponse.ErrorWithCodeAndMsg(ctx, http.StatusInternalServerError, "服务器内部错误"+err.Error())
		return
	}
	if code != registerForm.Code {
		commonResponse.ErrorWithCodeAndMsg(ctx, http.StatusBadRequest, "验证码不正确")
		return
	}
	user, err := global.UserSrvClient.CreateUser(context.Background(), &commonpb.CreateUserInfo{
		Mobile:   registerForm.Mobile,
		Nickname: registerForm.Mobile,
		Password: registerForm.Password,
	})
	if err != nil {
		zap.S().Errorf("[Register]新建用户失败：%s", err.Error())
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "用户服务srv")
		return
	}

	// 生成Token
	j := commonMiddleware.NewJWT(global.ServerConfig.JWTInfo.SigningKey)
	roleId, _ := strconv.Atoi(user.Role)
	claims := claims.CustomClaims{
		ID:          string(user.Id),
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
		commonResponse.ErrorWithCodeAndMsg(ctx, http.StatusInternalServerError, "生成token失败")
		return
	}
	data := gin.H{
		"id":         user.Id,
		"nick_name":  user.Nickname,
		"token":      token,
		"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
	}
	commonResponse.OkWithData(ctx, data)
}

func Ping(ctx *gin.Context) {
	commonResponse.Ok(ctx)
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
