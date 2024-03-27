package config

type UserSrvConfig struct {
	Host string `mapstructure:"host"`
	Port uint32 `mapstructure:"port"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key"`
}

type SMSConfig struct {
	AccessKeyId     string `mapstructure:"accessKeyId"`
	AccessKeySecret string `mapstructure:"accessKeySecret"`
	Protocol        string `mapstructure:"protocol"`
	RegionId        string `mapstructure:"regionId"`
	Domain          string `mapstructure:"domain"`
	SignName        string `mapstructure:"signName"`
	TemplateCode    string `mapstructure:"templateCode"`
	ApiName         string `mapstructure:"apiName"`
}

type ServerConfig struct {
	Name          string        `mapstructure:"name"`
	Port          uint32        `mapstructure:"port"`
	UserSrvConfig UserSrvConfig `mapstructure:"user-srv"`
	JWTInfo       JWTConfig     `mapstructure:"jwt"`
	SMSConfig     SMSConfig     `mapstructure:"sms"`
}
