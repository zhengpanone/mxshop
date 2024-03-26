package config

type UserSrvConfig struct {
	Host string `mapstructure:"host"`
	Port uint32 `mapstructure:"port"`
}

type ServerConfig struct {
	Name          string        `mapstructure:"name"`
	Port          uint32        `mapstructure:"port"`
	UserSrvConfig UserSrvConfig `mapstructure:"user-srv"`
}
