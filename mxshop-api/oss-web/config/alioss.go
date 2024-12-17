package config

import (
	"context"
	s3config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// OSSStorage OSS 实现了 ObjectStorage 接口，用于操作 OSS 服务
type OSSStorage struct {
	// 可以添加 AWS S3 相关配置字段
	client *s3.Client
}

// NewOSSStorage 创建一个连接到 aws 的对象存储客户端
func NewOSSStorage(region string) (*S3Storage, error) {
	s3cfg, err := s3config.LoadDefaultConfig(context.TODO(), s3config.WithRegion(region))
	if err != nil {
		return nil, err
	}
	client := s3.NewFromConfig(s3cfg)
	return &S3Storage{client: client}, nil
}
