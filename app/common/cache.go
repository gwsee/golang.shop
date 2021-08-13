package common

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"gwsee.com.api/config"
)

//缓存管理
// 缓存管理可以自动调节是用 内存还是redis
var CacheConfig = new(config.CacheMgr)

//初始化确定是加载那种缓存方式  通用的
func InitCache() {
	err := ini.MapTo(CacheConfig, "./config/cache.ini")
	if err != nil {
		// c.JSON(200, gin.H{"code": 0, "msg": err.Error()})
		gin.Logger()
		return
	}
	if CacheConfig.Name == "memory" {
		InitMemory()
	} else if CacheConfig.Name == "redis" {
		InitRedis()
	}
	return
}
