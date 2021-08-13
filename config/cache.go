package config

type CacheMgr struct {
	CacheConfig `ini:"cache"`
}
type CacheConfig struct {
	Name string `ini:"name"`
}
