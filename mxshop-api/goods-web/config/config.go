package config

type GoodsSrvConfig struct {
	Host string `mapstructure:"host" yaml:"host"`
	Port uint32 `mapstructure:"port" yaml:"port"`
	Name string `mapstructure:"name" yaml:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" yaml:"key"`
}

type JaegerConfig struct {
	Host string `mapstructure:"host" yaml:"host"`
	Port uint32 `mapstructure:"port" yaml:"port"`
}
type System struct {
	UseRedis bool `mapstructure:"use-redis" json:"user-redis" yaml:"use-redis"` // 使用redis
}

type Consul struct {
	Host string `mapstructure:"host" yaml:"host"`
	Port uint32 `mpastructure:"port" yaml:"port"`
}

type ServerConfig struct {
	Name           string         `mapstructure:"name" json:"name" yaml:"name"`
	Host           string         `mapstructure:"host" json:"host" yaml:"host"`
	Port           uint32         `mapstructure:"port" json:"port" yaml:"port"`
	Tags           []string       `mapstructure:"tags" json:"tags" yaml:"tags"`
	GoodsSrvConfig GoodsSrvConfig `mapstructure:"user-srv" json:"user-srv" yaml:"user-srv"`
	JWTInfo        JWTConfig      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	System         System         `mapstructure:"system" json:"system" yaml:"system"`
	Consul         Consul         `mapstructure:"consul" json:"consul" yaml:"consul"`
	Nacos          NacosConfig    `mapstructure:"nacos" json:"nacos" yaml:"nacos"`
	Jaeger         JaegerConfig   `mapstructure:"jaeger" json:"jaeger" yaml:"jaeger"`
	// 日志配置
	LogConfig LogConfig `mapstructure:"log" json:"log" yaml:"log"`
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
