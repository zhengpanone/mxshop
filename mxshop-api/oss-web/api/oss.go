package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"oss-web/global"
	"oss-web/utils/storage"
	"strings"
)

// GenerateUploadToken 生成上传凭证
func GenerateUploadToken(ctx *gin.Context) {
	response := storage.GetPolicyToken()
	ctx.Header("Access-Control-Allow-Methods", "POST")
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.String(200, response)

}

func HandlerRequest(ctx *gin.Context) {
	fmt.Println("\nHandle Post Request ... ")
	// Get PublicKey bytes
	bytePublicKey, err := storage.GetPublicKey(ctx)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	// Get Authorization bytes : decode from Base64String
	byteAuthorization, err := storage.GetAuthorization(ctx)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	// Get MD5 bytes from Newly Constructed Authorization String.
	byteMD5, bodyStr, err := storage.GetMD5FromNewAuthString(ctx)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	decodeUrl, err := url.QueryUnescape(bodyStr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(decodeUrl)
	params := make(map[string]string)
	split := strings.Split(decodeUrl, "&")
	for _, v := range split {
		dataList := strings.Split(v, "=")
		fmt.Println(v)
		params[dataList[0]] = dataList[1]
	}
	fileName := params["filename"]
	fileUrl := fmt.Sprintf("%s/%s", global.ServerConfig.OssInfo.Host, fileName)

	// verifySignature and response to client
	if storage.VerifySignature(bytePublicKey, byteMD5, byteAuthorization) {
		// do something you want according to callback_body ...
		ctx.JSON(http.StatusOK, gin.H{
			"url": fileUrl,
		})
		//utils.ResponseSuccess(ctx)  // response OK : 200
	} else {
		ctx.Status(http.StatusBadRequest) // response FAILED : 400
	}

}
