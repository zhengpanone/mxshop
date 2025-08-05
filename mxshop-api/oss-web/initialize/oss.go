package initialize

import (
	"fmt"
	"github.com/zhengpanone/mxshop/mxshop-api/oss-web/config"
	"github.com/zhengpanone/mxshop/mxshop-api/oss-web/global"
	"go.uber.org/zap"
)

func InitOSS(ossConfigs map[config.StorageType]config.OssConfig) {
	client := global.OSSClient

	for storageType, ossConfig := range ossConfigs {
		switch storageType {
		case config.StorageTypeMinIO:
			global.OSSConfig = &ossConfig
			minioClient, err := config.NewMinIOClient(ossConfig)
			if err != nil {
				zap.S().Errorln("初始化Minio客户端失败,", err.Error())
				panic(err)
			}
			client.MinIOClient = *minioClient
			zap.S().Info("初始化Minio客户端成功")
		case config.StorageTypeOSS:
			// TODO
			//ossClient, err := storage.NewS3Storage()
		case config.StorageTypeS3:
		// TODO
		default:
			// 未知的存储类型
			panic(fmt.Errorf("unsupported storage type: %v", storageType))
		}

	}

}
