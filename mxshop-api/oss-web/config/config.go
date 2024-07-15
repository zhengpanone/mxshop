package config

type JWTConfig struct {
	SigningKey string `mapstructure:"key" yaml:"key"`
}

type Consul struct {
	Host string `mapstructure:"host" yaml:"host"`
	Port uint32 `mpastructure:"port" yaml:"port"`
}

type OssConfig struct {
	ApiKey      string `mapstructure:"key" json:"key" yaml:"key"`
	ApiSecret   string `mapstructure:"secret" json:"secret" yaml:"secret"`
	Host        string `mapstructure:"host" json:"host" yaml:"host"`
	CallBackUrl string `mapstructure:"callback_url" json:"callback_url" yaml:"callback_url"`
	UploadDir   string `mapstructure:"upload_url" json:"upload_url" yaml:"upload_url"`
}

type ServerConfig struct {
	Name    string      `mapstructure:"name" json:"name" yaml:"name"`
	Host    string      `mapstructure:"host" json:"host" yaml:"host"`
	Port    uint32      `mapstructure:"port" json:"port" yaml:"port"`
	Tags    []string    `mapstructure:"tags" json:"tags" yaml:"tags"`
	JWTInfo JWTConfig   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Consul  Consul      `mapstructure:"consul" json:"consul" yaml:"consul"`
	Nacos   NacosConfig `mapstructure:"nacos" json:"nacos" yaml:"nacos"`
	OssInfo OssConfig   `mapstructure:"oss" json:"oss" yaml:"oss"`
}

type NacosConfig struct {
	Host        string `mapstructure:"host"`
	Port        uint64 `mastructure:"port"`
	ContextPath string `mastructure:"context-path"`
	Namespace   string `mastructure:"namespace"`
	User        string `mastructure:"user"`
	Password    string `mastructure:"password"`
	DataId      string `mastructure:"dataId"`
	Group       string `mastructure:"group"`
}
