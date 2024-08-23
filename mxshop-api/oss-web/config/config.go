package config

import commonConfig "mxshop-api/common/config"

type OssConfig struct {
	ApiKey      string `mapstructure:"key" json:"key" yaml:"key"`
	ApiSecret   string `mapstructure:"secret" json:"secret" yaml:"secret"`
	Host        string `mapstructure:"host" json:"host" yaml:"host"`
	CallBackUrl string `mapstructure:"callback_url" json:"callback_url" yaml:"callback_url"`
	UploadDir   string `mapstructure:"upload_url" json:"upload_url" yaml:"upload_url"`
}

type ServerConfig struct {
	Name      string                   `mapstructure:"name" json:"name" yaml:"name"`
	Host      string                   `mapstructure:"host" json:"host" yaml:"host"`
	Port      uint32                   `mapstructure:"port" json:"port" yaml:"port"`
	Tags      []string                 `mapstructure:"tags" json:"tags" yaml:"tags"`
	JWTInfo   commonConfig.JWTConfig   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Consul    commonConfig.Consul      `mapstructure:"consul" json:"consul" yaml:"consul"`
	Nacos     commonConfig.NacosConfig `mapstructure:"nacos" json:"nacos" yaml:"nacos"`
	OssInfo   OssConfig                `mapstructure:"oss" json:"oss" yaml:"oss"`
	LogConfig commonConfig.LogConfig   `mapstructure:"log" json:"log" yaml:"log"` // 日志配置
}
