package storage

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"goods-web/global"
	"io"
)

// MinIO 实现了 ObjectStorage 接口，用于操作 MinIO 服务
type MinIO struct {
	// 可以添加 MinIO 相关配置字段
	client *minio.Client
}

// NewMinIOClient 创建一个连接到 MinIO 的对象存储客户端
func NewMinIOClient(endpoint, accessKey, secretKey string, useSSL bool) (*MinIO, error) {
	// 初始化 MinIO 客户端
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		global.Logger.Fatal(fmt.Sprintf("Failed to initialize MinIO client:%s", err))
		return nil, err
	}

	return &MinIO{
		client: minioClient,
	}, nil
}

// Upload 实现了在 MinIO 上上传对象的方法
func (m *MinIO) Upload(ctx context.Context, bucketName, objectName string, reader io.Reader, objectSize int64, contentType string) (interface{}, error) {
	// 使用 MinIO 客户端上传对象
	object, err := m.client.PutObject(ctx, bucketName, objectName, reader, objectSize, minio.PutObjectOptions{
		ContentType: contentType,
	})
	return object, err
}

// Download 实现了从 MinIO 下载对象的方法
func (m *MinIO) Download(ctx context.Context, bucketName, objectName string) (io.ReadCloser, error) {
	// 使用 MinIO 客户端下载对象
	obj, err := m.client.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Delete 实现了从 MinIO 删除对象的方法
func (m *MinIO) Delete(ctx context.Context, bucketName, objectName string) error {
	// 使用 MinIO 客户端删除对象
	err := m.client.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{})
	return err
}
