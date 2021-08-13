package config

type DbMgr struct {
	MysqlConfig `ini:"mysql"`
}

// mysql 配置项
type MysqlConfig struct {
	Hostname string `ini:"hostname"`
	Database string `ini:"database"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	Hostport int    `ini:"hostport"`
	Maxopen  int    `ini:"maxopen"`
	Maxidle  int    `ini:"maxidle"`
}
