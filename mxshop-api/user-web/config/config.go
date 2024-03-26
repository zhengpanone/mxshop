package config

type UserSrvConfig struct {
	Host string `mapstructure:"host"`
	Port uint32 `mapstructure:"port"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key"`
}

type ServerConfig struct {
	Name          string        `mapstructure:"name"`
	Port          uint32        `mapstructure:"port"`
	UserSrvConfig UserSrvConfig `mapstructure:"user-srv"`
	JWTInfo       JWTConfig     `mapstructure:"jwt"`
}
