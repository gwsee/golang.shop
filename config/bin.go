package config

type BinMgr struct {
	CrontabConfig `ini:"crontab"`
}
type CrontabConfig struct {
	Path string `ini:"path"`
}
