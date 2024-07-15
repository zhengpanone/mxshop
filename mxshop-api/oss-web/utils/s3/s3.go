package s3

import "io"

// ObjectStorage 定义了通用的对象存储接口
type ObjectStorage interface {
	Upload(bucketName, objectName string, reader io.Reader, contentType string) error
	Download(bucketName, objectName string) (io.ReadCloser, error)
	Delete(bucketName, objectName string) error
}

// AWS S3 实现了 ObjectStorage 接口，用于操作 AWS S3 服务
type AWSS3 struct {
	// 可以添加 AWS S3 相关配置字段
}

// Upload 实现了在 AWS S3 上上传对象的方法
func (a *AWSS3) Upload(bucketName, objectName string, reader io.Reader, contentType string) error {
	// 实现 AWS S3 的上传逻辑
	return nil
}

// Download 实现了从 AWS S3 下载对象的方法
func (a *AWSS3) Download(bucketName, objectName string) (io.ReadCloser, error) {
	// 实现 AWS S3 的下载逻辑
	return nil, nil
}

// Delete 实现了从 AWS S3 删除对象的方法
func (a *AWSS3) Delete(bucketName, objectName string) error {
	// 实现 AWS S3 的删除逻辑
	return nil
}
