package config

import (
	commonConfig "github.com/zhengpanone/mxshop/mxshop-api/common/config"
)

type GoodsSrvConfig struct {
	Host string `mapstructure:"host" yaml:"host"`
	Port uint32 `mapstructure:"port" yaml:"port"`
	Name string `mapstructure:"name" yaml:"name"`
}

type System struct {
	UseRedis bool `mapstructure:"use-redis" json:"user-redis" yaml:"use-redis"` // 使用redis
}

type ServerConfig struct {
	Name           string                    `mapstructure:"name" json:"name" yaml:"name"`
	Host           string                    `mapstructure:"host" json:"host" yaml:"host"`
	Port           uint32                    `mapstructure:"port" json:"port" yaml:"port"`
	Tags           []string                  `mapstructure:"tags" json:"tags" yaml:"tags"`
	System         System                    `mapstructure:"system" json:"system" yaml:"system"`
	RedisConfig    commonConfig.RedisConfig  `mapstructure:"redis" json:"redis" yaml:"redis"`
	GoodsSrvConfig GoodsSrvConfig            `mapstructure:"user-srv" json:"user-srv" yaml:"user-srv"`
	JWTInfo        commonConfig.JWTConfig    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Consul         commonConfig.Consul       `mapstructure:"consul" json:"consul" yaml:"consul"`
	Nacos          commonConfig.NacosConfig  `mapstructure:"nacos" json:"nacos" yaml:"nacos"`
	Jaeger         commonConfig.JaegerConfig `mapstructure:"jaeger" json:"jaeger" yaml:"jaeger"`
	LogConfig      commonConfig.LogConfig    `mapstructure:"log" json:"log" yaml:"log"` // 日志配置
}
