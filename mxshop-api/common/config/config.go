package config

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
