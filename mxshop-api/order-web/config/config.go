package config

type SrvConfig struct {
	Name string `mapstructure:"name" yaml:"name"`
}
type JWTConfig struct {
	SigningKey string `mapstructure:"key" yaml:"key"`
}

type System struct {
	UseRedis bool `mapstructure:"use-redis" json:"user-redis" yaml:"use-redis"` // 使用redis
}

type Consul struct {
	Host string `mapstructure:"host" yaml:"host"`
	Port uint32 `mapstructure:"port" yaml:"port"`
}

type ServerConfig struct {
	Name               string      `mapstructure:"name" json:"name" yaml:"name"`
	Host               string      `mapstructure:"host" json:"host" yaml:"host"`
	Port               uint32      `mapstructure:"port" json:"port" yaml:"port"`
	Tags               []string    `mapstructure:"tags" json:"tags" yaml:"tags"`
	OrderSrvConfig     SrvConfig   `mapstructure:"order-srv" json:"order-srv" yaml:"order-srv"`
	GoodsSrvConfig     SrvConfig   `mapstructure:"goods-srv" json:"goods-srv" yaml:"goods-srv"`
	InventorySrvConfig SrvConfig   `mapstructure:"inventory-srv" json:"inventory-srv" yaml:"inventory-srv"`
	JWTInfo            JWTConfig   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	System             System      `mapstructure:"system" json:"system" yaml:"system"`
	Consul             Consul      `mapstructure:"consul" json:"consul" yaml:"consul"`
	Nacos              NacosConfig `mapstructure:"nacos" json:"nacos" yaml:"nacos"`
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
