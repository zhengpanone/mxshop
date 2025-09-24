package config

import "time"

// LogConfig 日志配置
type LogConfig struct {
	// 日志级别
	Level string `mapstructure:"level"`
	// 志文件的位置
	Filename string `mapstructure:"filename"`
	// 在进行切割之前，日志文件的最大大小（以MB为单位）
	MaxSize int `mapstructure:"max_size"`
	// 保留旧文件的最大天数
	MaxAge int `mapstructure:"max_age"`
	// 保留旧文件的最大个数
	MaxBackups int `mapstructure:"max_backups"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" yaml:"key"`
	IsMulti    bool   `mapstructure:"is_multi" yaml:"is_multi"` // 是否多终端登录
	MaxToken   int64  `mapstructure:"max_token" yaml:"max_token"`
	TTL        int64  `mapstructure:"ttl" yaml:"ttl"`
	TTLUnit    string `mapstructure:"ttl_unit" yaml:"ttl_unit"`
}

type JaegerConfig struct {
	Host string `mapstructure:"host" yaml:"host"`
	Port uint32 `mapstructure:"port" yaml:"port"`
}

type Consul struct {
	Host string `mapstructure:"host" yaml:"host"`
	Port uint32 `mapstructure:"port" yaml:"port"`
}

type NacosConfig struct {
	Host        string `mapstructure:"host"`
	Port        uint64 `mapstructure:"port"`
	ContextPath string `mapstructure:"context-path"`
	Namespace   string `mapstructure:"namespace"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	DataId      string `mapstructure:"dataId"`
	Group       string `mapstructure:"group"`
}

type RedisConfig struct {
	Host        string        `mapstructure:"host" yaml:"host"`
	Port        uint32        `mapstructure:"port" yaml:"port"`
	Password    string        `mapstructure:"password" yaml:"password"`
	Database    int           `mapstructure:"database" yaml:"database"`
	DialTimeout time.Duration `mapstructure:"dialTimeout" yaml:"dialTimeout"`
}
