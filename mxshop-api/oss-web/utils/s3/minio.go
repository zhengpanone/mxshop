package s3

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
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
		fmt.Println("Failed to initialize MinIO client:", err)
		return nil, err
	}

	return &MinIO{
		client: minioClient,
	}, nil
}

func GetPolicyToken() {

}

// Upload 实现了在 MinIO 上上传对象的方法
func (m *MinIO) Upload(bucketName, objectName string, reader io.Reader, contentType string) error {
	// 使用 MinIO 客户端上传对象
	_, err := m.client.PutObject(context.Background(), bucketName, objectName, reader, -1, minio.PutObjectOptions{
		ContentType: contentType,
	})
	return err
}

// Download 实现了从 MinIO 下载对象的方法
func (m *MinIO) Download(bucketName, objectName string) (io.ReadCloser, error) {
	// 使用 MinIO 客户端下载对象
	obj, err := m.client.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Delete 实现了从 MinIO 删除对象的方法
func (m *MinIO) Delete(bucketName, objectName string) error {
	// 使用 MinIO 客户端删除对象
	err := m.client.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{})
	return err
}
