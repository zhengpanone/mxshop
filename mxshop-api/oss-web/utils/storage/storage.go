package storage

import (
	"context"
	"crypto"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hash"
	"io"
	"net/http"
	"net/url"
	"oss-web/global"
	"time"
)

var expireTime int64 = 3000

func getGmtIso8601(expireEnd int64) string {
	var tokenExpire = time.Unix(expireEnd, 0).Format("2006-01-02T15:04:05Z")
	return tokenExpire
}

type ConfigStruct struct {
	Expiration string     `json:"expiration"`
	Conditions [][]string `json:"conditions"`
}

type CallbackParam struct {
	CallbackUrl      string `json:"callbackUrl"`
	CallbackBody     string `json:"callbackBody"`
	CallbackBodyType string `json:"callbackBodyType"`
}

type PolicyToken struct {
	AccessKeyId string `json:"accessId"`
	Host        string `json:"host"`
	Expire      int64  `json:"expire"`
	Signature   string `json:"signature"`
	Policy      string `json:"policy"`
	Directory   string `json:"dir"`
	Callback    string `json:"callback"`
}

// GetPolicyToken 生成上传的预签名 URL
func GetPolicyToken() string {
	now := time.Now().Unix()
	expireEnd := now + expireTime
	var tokenExpire = getGmtIso8601(expireEnd)
	//create post policy json
	var config ConfigStruct
	config.Expiration = tokenExpire
	var condition []string
	condition = append(condition, "starts-with")
	condition = append(condition, "$key")
	condition = append(condition, global.ServerConfig.OssInfo.UploadDir)
	condition = append(condition, global.ServerConfig.OssInfo.UploadDir)
	config.Conditions = append(config.Conditions, condition)

	//calculate signature
	result, err := json.Marshal(config)
	deByte := base64.StdEncoding.EncodeToString(result)
	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte("" /*global.ServerConfig.OssInfo.ApiSecret*/))
	_, _ = io.WriteString(h, deByte)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))
	var callbackParam CallbackParam
	callbackParam.CallbackUrl = global.ServerConfig.OssInfo.CallBackUrl
	callbackParam.CallbackBody = "filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}"
	callbackParam.CallbackBodyType = "application/x-www-form-urlencoded"
	callbackStr, err := json.Marshal(callbackParam)
	if err != nil {
		fmt.Println("callback json err:", err)
	}

	callbackBase64 := base64.StdEncoding.EncodeToString(callbackStr)

	var policyToken PolicyToken
	policyToken.AccessKeyId = global.ServerConfig.OssInfo.ApiKey
	policyToken.Host = global.ServerConfig.OssInfo.Host
	policyToken.Expire = expireEnd
	policyToken.Signature = signedStr
	policyToken.Directory = global.ServerConfig.OssInfo.UploadDir
	policyToken.Policy = deByte
	policyToken.Callback = callbackBase64
	response, err := json.Marshal(policyToken)
	if err != nil {
		fmt.Println("json err:", err)
	}
	return string(response)
}

// ObjectStorage 定义了通用的对象存储接口
type ObjectStorage interface {
	Upload(ctx context.Context, bucketName, objectName string, reader io.Reader, objectSize int64, contentType string) (interface{}, error)
	Download(ctx context.Context, bucketName, objectName string) (io.ReadCloser, error)
	Delete(ctx context.Context, bucketName, objectName string) error
}

// GetPublicKey : Get PublicKey bytes from Request.URL
func GetPublicKey(ctx *gin.Context) ([]byte, error) {
	var bytePublicKey []byte
	publicKeyURLBase64 := ctx.GetHeader("publicKey")
	if publicKeyURLBase64 == "" {
		global.Logger.Error("GetPublicKey from Request header failed: No publicKey field")
		return bytePublicKey, errors.New("no publicKey field in request header")
	}
	publicKeyURL, err := base64.StdEncoding.DecodeString(publicKeyURLBase64)
	if err != nil {
		global.Logger.Error("decode publicKey failed", zap.Error(err))
		return bytePublicKey, err
	}
	responsePublicKeyURL, err := http.Get(string(publicKeyURL))
	if err != nil {
		global.Logger.Error("Get PublicKey Content from URL failed : %s ", zap.Error(err))
		return bytePublicKey, err
	}
	bytePublicKey, err = io.ReadAll(responsePublicKeyURL.Body)
	if err != nil {
		global.Logger.Error("Read PublicKey Content from URL failed : %s \n", zap.Error(err))
		return bytePublicKey, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			global.Logger.Error("responsePublicKeyURL Body Close failed : ", zap.Error(err))
		}
	}(responsePublicKeyURL.Body)
	return bytePublicKey, nil
}

// GetAuthorization : decode from Base64String
func GetAuthorization(ctx *gin.Context) ([]byte, error) {
	var byteAuthorization []byte
	strAuthorizationBase64 := ctx.GetHeader("authorization")
	if strAuthorizationBase64 == "" {
		global.Logger.Error("Failed to get authorization field from request header. ")
		return byteAuthorization, errors.New("no authorization field in Request header")
	}
	byteAuthorization, err := base64.StdEncoding.DecodeString(strAuthorizationBase64)
	return byteAuthorization, err
}

// GetMD5FromNewAuthString : Get MD5 bytes from Newly Constructed Authorization String.
func GetMD5FromNewAuthString(ctx *gin.Context) ([]byte, string, error) {
	var byteMD5 []byte
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		global.Logger.Error("Read Request Body failed : %s \n", zap.Error(err))
		return byteMD5, "", err
	}
	strCallbackBody := string(body)
	strURLPathDecode, err := url.PathUnescape(ctx.Request.URL.Path)
	if err != nil {
		global.Logger.Error(fmt.Sprintf("url.PathUnescape failed : URL.Path=%s, error=%s", ctx.Request.URL.Path, err.Error()))
		return byteMD5, "", err
	}
	// Generate New Auth String prepare for MD5
	strAuth := ""
	if ctx.Request.URL.RawQuery == "" {
		strAuth = fmt.Sprintf("%s\n%s", strURLPathDecode, strCallbackBody)
	} else {
		strAuth = fmt.Sprintf("%s?%s\n%s", strURLPathDecode, ctx.Request.URL.RawQuery, strCallbackBody)
	}
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(strAuth))
	byteMD5 = md5Ctx.Sum(nil)
	return byteMD5, strCallbackBody, nil
}

// VerifySignature 需要三个重要的数据信息来进行签名验证： 1>获取公钥PublicKey;  2>生成新的MD5鉴权串;  3>解码Request携带的鉴权串;
//
//	1>获取公钥PublicKey : 从RequestHeader的"x-oss-pub-key-url"字段中获取 URL, 读取URL链接的包含的公钥内容， 进行解码解析， 将其作为rsa.VerifyPKCS1v15的入参。
//	2>生成新的MD5鉴权串 : 把Request中的url中的path部分进行urlDecode， 加上url的query部分， 再加上body， 组合之后进行MD5编码， 得到MD5鉴权字节串。
//	3>解码Request携带的鉴权串 ： 获取RequestHeader的"authorization"字段， 对其进行Base64解码，作为签名验证的鉴权对比串。
//	rsa.VerifyPKCS1v15进行签名验证，返回验证结果。
func VerifySignature(bytePublicKey []byte, byteMd5 []byte, authorization []byte) bool {
	pubBlock, _ := pem.Decode(bytePublicKey)
	if pubBlock == nil {
		global.Logger.Error("Failed to parse PEM block containing the public key")
		return false
	}
	pubInterface, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if pubInterface == nil || err != nil {
		global.Logger.Error("x509.ParsePKIXPublicKey(publicKey) failed : ", zap.Error(err))
		return false
	}
	pub := pubInterface.(*rsa.PublicKey)
	err = rsa.VerifyPKCS1v15(pub, crypto.MD5, byteMd5, authorization)
	if err != nil {
		global.Logger.Error("Failed to verify signature : ", zap.Error(err))
		return false
	}
	global.Logger.Info("Signature Verification is Successful.")
	return true
}
