package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gwsee.com.api/app/common"
	shop "gwsee.com.api/app/shop/model"
	user "gwsee.com.api/app/user/model"
	"net/http"
)

const GlobalConfigKey = "_GlobalConfig"

type GlobalConfig struct {
	User   *user.User
	Shop   *shop.Shop
	Token  *user.UserToken
	ShopSn string
}

func InitGlobalConfig(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	var g GlobalConfig
	var gUser user.User
	var gToken user.UserToken
	var shopList []shop.Shop
	token := c.GetHeader("AccessToken")
	_shopSn := c.DefaultQuery("_shopSn", "")
	if token == "" {
		token = c.DefaultQuery("_AccessToken", "") //websocket导致的
		if token == "" {
			appid := c.GetHeader("AccessAppid")
			if appid != "" {
				var w shop.ShopWechat
				findWechat(&w, appid)
				if w.ShopSn != "" {
					g.Token = &gToken
					g.User = &gUser
					g.ShopSn = w.ShopSn
					c.Set(GlobalConfigKey, g)
					return
				}
			}
			return
		}
	}

	err := findToken(&gToken, token) // 这个是否可以优化 --类似这种感觉不应该调用里面的东西
	if err != nil {
		return
	}
	if gToken.User != "" {
		_ = json.Unmarshal([]byte(gToken.User), &gUser)
		_ = json.Unmarshal([]byte(gToken.Shop), &shopList)
		g.ShopSn = gToken.ShopSn
		if _shopSn != "" {
			for _, v := range shopList {
				if v.ShopSn == _shopSn {
					g.ShopSn = _shopSn
					break
				}
			}
		}
		for _, v := range shopList {
			if v.ShopSn == g.ShopSn {
				g.Shop = &v
				break
			}
		}
		//更新数据库信息 go
		//go EditToken(c)
	}
	//在每次运行结束之后把个人信息清除掉
	g.Token = &gToken
	g.User = &gUser
	c.Set(GlobalConfigKey, g)
	return

}

//查询登陆信息
func findToken(info *user.UserToken, token string) (err error) {
	if common.CacheConfig.Name == "redis" {
		//redis方式进行数据读取
		str, err := common.GetByRedis(token)
		if err != nil || str == "" {
			return err
		}
		_ = json.Unmarshal([]byte(str), &info)
		//然后更新redis信息 --不更新吧
	} else {
		sqlstr := "select * from  user_token where token =?"
		err = common.FindTable(info, sqlstr, token)
	}
	return
}
func findWechat(we *shop.ShopWechat, key string) (err error) {
	sqlstr := "select * from shop_wechat where wechat_appkey = ?"
	err = common.FindTable(we, sqlstr, key)
	return
}

//---------------------------------------数据验证
func NeedShopSn(c *gin.Context) { //appid都没能查到对应的数据
	g := parseGlobalConfig(c)
	if g.User.UserId < 1 {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "请登录"})
		c.Abort()
		return
	}
	if g.ShopSn == "" {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "请登录"})
		c.Abort()
		return
	}
}

//必须要用户登录
func NeedUser(c *gin.Context) {
	g := parseGlobalConfig(c)
	user := g.User
	if user.UserId < 1 {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "请登录"})
		c.Abort()
		return
	}
}

//需要拥有店铺角色
func NeedShop(c *gin.Context) {
	g := parseGlobalConfig(c)
	user := g.User
	shop := g.Shop
	if user.UserId < 1 {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "请登录"})
		c.Abort()
		return
	}
	if shop.ShopId < 1 {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "请登录"})
		c.Abort()
		return
	}
}

//需要拥有店铺平台管理员权限
func NeedAdmin(c *gin.Context) {
	g := parseGlobalConfig(c)
	user := g.User
	shop := g.Shop
	if user.UserId < 1 {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "请登录"})
		c.Abort()
		return
	}
	if shop.ShopId != 1 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "无权操作"})
		c.Abort()
		return
	}
}

//---------------------------------------数据获取
func GetConfig(c *gin.Context) *GlobalConfig {
	g := parseGlobalConfig(c)
	return g
}
func parseGlobalConfig(c *gin.Context) *GlobalConfig {
	if globalConfig, exists := c.Get(GlobalConfigKey); !exists {
		return nil
	} else {
		g := globalConfig.(GlobalConfig)
		return &g
	}
}
