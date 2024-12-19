package controller

import (
	"common/utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"oss-web/global"
	"oss-web/utils/storage"
	"strings"
	"time"
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
	fileUrl := fmt.Sprintf("%s/%s", global.OSSConfig.Endpoint, fileName)

	// verifySignature and response to client
	if storage.VerifySignature(bytePublicKey, byteMD5, byteAuthorization) {
		// do something you want according to callback_body ...
		result := gin.H{
			"url": fileUrl,
		}
		utils.OkWithData(ctx, result) // response OK : 200
	} else {
		utils.ErrorWithCodeAndMsg(ctx, http.StatusBadRequest, "上传文件失败")
	}
}

// UploadFile 文件上传接口
func UploadFile(c *gin.Context) {
	// 获取上传的文件
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file"})
		return
	}
	defer file.Close()

	// 生成文件在 MinIO 中的路径和名称
	objectName := fmt.Sprintf("%d-%s", time.Now().Unix(), header.Filename)
	contentType := header.Header.Get("Content-Type")

	// 上传文件到 MinIO
	ctx := context.Background()
	object, err := global.OSSClient.MinIOClient.Upload(ctx, global.OSSConfig.Bucket, objectName, file, header.Size, contentType)
	zap.S().Info("Upload File Success", object)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}

	// 构建文件的访问 URL
	fileURL := fmt.Sprintf("http://%s/%s/%s", global.OSSConfig.Endpoint, global.OSSConfig.Bucket, objectName)

	c.JSON(http.StatusOK, gin.H{
		"message":  "File uploaded successfully",
		"file_url": fileURL,
	})
}
