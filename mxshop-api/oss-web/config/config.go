package config

import (
	commonConfig "common/config"
)

type OSSClients struct {
	MinIOClient MinIO
	S3Client    S3Storage
	// TODO
}

// StorageType 定义支持的存储类型
type StorageType string

// 支持多种对象存储类型（如 OSS, S3, MinIO）
const (
	StorageTypeMinIO StorageType = "minio"
	StorageTypeS3    StorageType = "s3"
	StorageTypeOSS   StorageType = "oss"
)

type OssConfig struct {
	AccessKey   string      `mapstructure:"access_key" json:"access_key" yaml:"access_key"`
	SecretKey   string      `mapstructure:"secret_key" json:"secret_key" yaml:"secret_key"`       // 访问密钥的Secret
	Endpoint    string      `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`             // 服务器地址或服务端点
	CallbackURL string      `mapstructure:"callback_url" json:"callback_url" yaml:"callback_url"` // 回调地址
	UploadDir   string      `mapstructure:"upload_dir" json:"upload_dir" yaml:"upload_dir"`       // 上传基础目录
	StorageType StorageType `mapstructure:"storage_type" json:"storage_type" yaml:"storage_type"` // 存储类型：minio/s3/oss
	Bucket      string      `mapstructure:"bucket" json:"bucket" yaml:"bucket"`                   // 存储桶名称
	Region      string      `mapstructure:"region" json:"region" yaml:"region"`                   // S3 或其他服务可能需要的区域配置
	BucketName  string      `mapstructure:"bucket_name" json:"bucket_name" yaml:"bucket_name"`    // 存储桶名称
	UseSSL      bool        `mapstructure:"use_ssl" json:"use_ssl" yaml:"use_ssl"`                // 是否使用 HTTPS
}

type ServerConfig struct {
	Name      string                    `mapstructure:"name" json:"name" yaml:"name"`
	Host      string                    `mapstructure:"host" json:"host" yaml:"host"`
	Port      uint32                    `mapstructure:"port" json:"port" yaml:"port"`
	Tags      []string                  `mapstructure:"tags" json:"tags" yaml:"tags"`
	JWTInfo   commonConfig.JWTConfig    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Consul    commonConfig.Consul       `mapstructure:"consul" json:"consul" yaml:"consul"`
	Nacos     commonConfig.NacosConfig  `mapstructure:"nacos" json:"nacos" yaml:"nacos"`
	OssInfo   map[StorageType]OssConfig `mapstructure:"oss" json:"oss" yaml:"oss"`
	LogConfig commonConfig.LogConfig    `mapstructure:"log" json:"log" yaml:"log"` // 日志配置
}
