package config

// redis 配置项
type RedisMgr struct {
	RedisConfig `ini:"redis"`
}

type RedisConfig struct {
	Hostname string `ini:"hostname"`
	Hostport int    `ini:"hostport"`
	Password string `ini:"password"`
}
