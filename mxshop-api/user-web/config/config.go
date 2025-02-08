package config

import (
	commonConfig "github.com/zhengpanone/mxshop/common/config"
	"time"
)

type UserSrvConfig struct {
	//Host string `mapstructure:"host" yaml:"host"`
	//Port uint32 `mapstructure:"port" yaml:"port"`
	Name string `mapstructure:"name" yaml:"name"`
}

type SMSConfig struct {
	AccessKeyId     string `mapstructure:"accessKeyId" yaml:"accessKeyId"`
	AccessKeySecret string `mapstructure:"accessKeySecret" yaml:"accessKeySecret"`
	Protocol        string `mapstructure:"protocol" yaml:"protocol"`
	RegionId        string `mapstructure:"regionId" yaml:"regionId"`
	Domain          string `mapstructure:"domain" yaml:"domain"`
	SignName        string `mapstructure:"signName" yaml:"signName"`
	TemplateCode    string `mapstructure:"templateCode" yaml:"templateCode"`
	ApiName         string `mapstructure:"apiName" yaml:"apiName"`
}

type RedisConfig struct {
	Host        string        `mapstructure:"host" yaml:"host"`
	Port        uint32        `mapstructure:"port" yaml:"port"`
	Password    string        `mapstructure:"password" yaml:"password"`
	DefaultDB   int           `mapstructure:"defaultDB" yaml:"defaultDB"`
	DialTimeout time.Duration `mapstructure:"dialTimeout" yaml:"dialTimeout"`
}

type System struct {
	UseRedis bool `mapstructure:"use-redis" json:"user-redis" yaml:"use-redis"` // 使用redis
}

type ServerConfig struct {
	Name          string                    `mapstructure:"name" json:"name" yaml:"name"`
	Host          string                    `mapstructure:"host" json:"host" yaml:"host"`
	Port          uint32                    `mapstructure:"port" json:"port" yaml:"port"`
	Tags          []string                  `mapstructure:"tags" json:"tags" yaml:"tags"`
	UserSrvConfig UserSrvConfig             `mapstructure:"user-srv" json:"user-srv" yaml:"user-srv"`
	SMSConfig     SMSConfig                 `mapstructure:"sms" json:"smd,omitempty" yaml:"sms"`
	RedisConfig   RedisConfig               `mapstructure:"redis" json:"redis" yaml:"redis"`
	System        System                    `mapstructure:"system" json:"system" yaml:"system"`
	EnableCaptcha bool                      `mapstructure:"enable-captcha" json:"enable-captcha" yaml:"enable-captcha"`
	JWTInfo       commonConfig.JWTConfig    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Consul        commonConfig.Consul       `mapstructure:"consul" json:"consul" yaml:"consul"`
	Nacos         commonConfig.NacosConfig  `mapstructure:"nacos" json:"nacos" yaml:"nacos"`
	LogConfig     commonConfig.LogConfig    `mapstructure:"log" json:"log" yaml:"log"` // 日志配置
	Jaeger        commonConfig.JaegerConfig `mapstructure:"jaeger" json:"jaeger" yaml:"jaeger"`
}
