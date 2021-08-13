package common

import (
	"fmt"
	"github.com/go-redis/redis"
	"gopkg.in/ini.v1"
	"gwsee.com.api/config"
	"strconv"
	"time"
)

//redis管理
var redisClient *redis.Client
var redisConfig = new(config.RedisMgr)

//启动redis 服务
func InitRedis() {
	//加载配置
	err := ini.MapTo(redisConfig, "./config/redis.ini")
	if err != nil {
		//c.JSON(200, gin.H{"code": 0, "msg": err.Error()})
		fmt.Println("redis配置获取失败", err.Error())
		return
	}
	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Hostname + ":" + strconv.Itoa(redisConfig.Hostport),
		Password: redisConfig.Password, // no password set
		DB:       0,                    // use default DB
	})
	_, err = redisClient.Ping().Result()
	if err != nil {
		//c.JSON(200, gin.H{"code": 0, "msg": err.Error()})
		fmt.Println("redis链接失败", err.Error())
		//c.Abort()
		return
	}
	// defer redisClient.Close() //用完就关闭如果需要启动的话
	//c.Next()
	return
}
func SetByRedis(key string, value interface{}, maxTime time.Duration) (err error) {
	err = redisClient.Set(key, value, maxTime).Err()
	return err
}
func GetByRedis(key string) (val string, err error) {
	val, err = redisClient.Get(key).Result()
	if err == redis.Nil {
		fmt.Println("redis没有查询到数据")
		err = nil
	} else if err != nil {
		return
	}
	return
}
