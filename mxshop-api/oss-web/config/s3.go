package config

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	s3config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
)

// S3Storage AWS S3 实现了 ObjectStorage 接口，用于操作 AWS S3 服务
type S3Storage struct {
	// 可以添加 AWS S3 相关配置字段
	client *s3.Client
}

// NewS3Storage 创建一个连接到 aws 的对象存储客户端
func NewS3Storage(region string) (*S3Storage, error) {
	s3cfg, err := s3config.LoadDefaultConfig(context.TODO(), s3config.WithRegion(region))
	if err != nil {
		return nil, err
	}
	client := s3.NewFromConfig(s3cfg)
	return &S3Storage{client: client}, nil
}

// Upload 实现了在 AWS S3 上上传对象的方法
func (s *S3Storage) Upload(ctx context.Context, bucketName, objectName string, reader io.Reader, objectSize int64, contentType string) (interface{}, error) {
	// 实现 AWS S3 的上传逻辑
	object, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:        aws.String(bucketName),
		Key:           aws.String(objectName),
		Body:          reader,
		ContentType:   aws.String(contentType),
		ContentLength: aws.Int64(objectSize),
	})
	return object, err
}

// Download 实现了从 AWS S3 下载对象的方法
func (s *S3Storage) Download(ctx context.Context, bucketName, objectName string) (io.ReadCloser, error) {
	// 实现 AWS S3 的下载逻辑
	object, err := s.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectName),
	})
	if err != nil {
		return nil, err
	}
	return object.Body, err
}

// Delete 实现了从 AWS S3 删除对象的方法
func (s *S3Storage) Delete(ctx context.Context, bucketName, objectName string) error {
	// 实现 AWS S3 的删除逻辑
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectName),
	})
	return err
}
