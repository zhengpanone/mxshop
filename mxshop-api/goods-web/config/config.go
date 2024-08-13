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
