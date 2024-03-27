package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mxshop-api/user-web/forms"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/global/response"
	"mxshop-api/user-web/middlewares"
	"mxshop-api/user-web/models"
	"mxshop-api/user-web/proto"
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

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	// 将grpc的code转换成http的code
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})

			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "用户服务不可用",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "其他错误" + e.Message(),
				})

			}

		}

	}
}

func GetUserList(ctx *gin.Context) {

	// 拨号连接用户grpc服务
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserSrvConfig.Host, global.ServerConfig.UserSrvConfig.Port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList]连接【用户服务失败】", "msg", err.Error())
	}
	// 生成grpc的client并调用接口
	userSrvClient := proto.NewUserClient(userConn)
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "10"))
	rsp, err := userSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		Page: uint32(page),
		Size: uint32(size),
	})
	if err != nil {
		zap.S().Errorw("[GetUserList]查询【用户列表】失败")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	claims, _ := ctx.Get("claims")
	customClaims := claims.(*models.CustomClaims)
	zap.S().Infof("当前访问用户：%d", customClaims.ID)
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

func PasswordLogin(ctx *gin.Context) {
	passwordLogin := forms.PasswordLoginForm{}
	if err := ctx.ShouldBind(&passwordLogin); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	if !VerifyCaptcha(passwordLogin.CaptchaId, passwordLogin.Captcha) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码不正确",
		})
		return
	}
	// 拨号连接用户grpc服务器
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserSrvConfig.Host, global.ServerConfig.UserSrvConfig.Port), grpc.WithInsecure())

	if err != nil {
		zap.S().Errorw("[Get]连接【用户服务失败】", "msg", err.Error())
	}
	// 生成grpc的client并调用接口
	userSevClient := proto.NewUserClient(userConn)
	// 登录的逻辑
	if rsp, err := userSevClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
		Mobile: passwordLogin.Mobile,
	}); err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusBadRequest, map[string]string{
					"mobile": "用户不存在",
				})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]string{
					"msg": "登录失败" + e.Message(),
				})
			}
			return
		}
	} else {
		// 只是查询到用户，没有检查密码
		if passRsp, passErr := userSevClient.CheckPassword(context.Background(), &proto.PasswordCheckInfo{
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
					ctx.JSON(http.StatusInternalServerError, gin.H{
						"msg": "生成token失败",
					})
					return
				}
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"id":         rsp.Id,
					"nick_name":  rsp.Nickname,
					"token":      token,
					"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
				})
			} else {
				ctx.JSON(http.StatusBadRequest, map[string]string{
					"msg": "登录失败",
				})
			}
		}
	}

}

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func HandleValidatorError(ctx *gin.Context, err error) {
	// 如何返回错误信息
	errs, ok := err.(validator.ValidationErrors)
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
