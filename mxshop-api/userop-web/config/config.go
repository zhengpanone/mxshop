package config

import commonConfig "common/config"

type SrvConfig struct {
	Name string `mapstructure:"name" yaml:"name"`
}

type System struct {
	UseRedis bool `mapstructure:"use-redis" json:"user-redis" yaml:"use-redis"` // 使用redis
}

type ServerConfig struct {
	Name            string                   `mapstructure:"name" json:"name" yaml:"name"`
	Host            string                   `mapstructure:"host" json:"host" yaml:"host"`
	Port            uint32                   `mapstructure:"port" json:"port" yaml:"port"`
	Tags            []string                 `mapstructure:"tags" json:"tags" yaml:"tags"`
	UserOpSrvConfig SrvConfig                `mapstructure:"userop-srv" json:"userop-srv" yaml:"userop-srv"`
	GoodsSrvConfig  SrvConfig                `mapstructure:"goods-srv" json:"goods-srv" yaml:"goods-srv"`
	JWTInfo         commonConfig.JWTConfig   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	System          System                   `mapstructure:"system" json:"system" yaml:"system"`
	Consul          commonConfig.Consul      `mapstructure:"consul" json:"consul" yaml:"consul"`
	Nacos           commonConfig.NacosConfig `mapstructure:"nacos" json:"nacos" yaml:"nacos"`
	LogConfig       commonConfig.LogConfig   `mapstructure:"log" json:"log" yaml:"log"` // 日志配置
}
