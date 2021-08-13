package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/user/model"
	"gwsee.com.api/app/user/service"
)

/**
在数据库里面查询token然后 赋值给user
*/
//func findToken(c *gin.Context) {
//	token := c.GetHeader("AccessToken")
//	_shopSn := c.DefaultQuery("_shopSn","")
//	if token == "" {
//		//fmt.Println("未找到AccessToken")
//		token = c.DefaultQuery("_AccessToken","")//websocket导致的
//		if token == "" {
//			return
//		}
//	}
//	//这里判断 是redis  还是 mysql 进行访问token信息 如果是mysql 就是最下面那个
//
//	//数据库方式进行读取token
//	err := service.FindToken(&auth.GlobalToken, token)
//	if err != nil {
//		return
//	}
//	var shopList []shop.Shop
//	//fmt.Println(auth.GlobalToken)
//	//fmt.Println("auth.GlobalToken")
//	if auth.GlobalToken.User != "" {
//		_ = json.Unmarshal([]byte(auth.GlobalToken.User), &auth.User)
//		_ = json.Unmarshal([]byte(auth.GlobalToken.Shop), &shopList)
//
//		auth.ShopSn = auth.GlobalToken.ShopSn
//		if _shopSn != "" {
//			for _, v := range shopList {
//				if v.ShopSn == _shopSn {
//					auth.ShopSn = _shopSn
//					break
//				}
//			}
//		}
//		for _, v := range shopList {
//			if v.ShopSn == auth.ShopSn {
//				auth.GlobalShop = v
//			}
//		}
//		//更新数据库信息 go
//		//go EditToken(c)
//	}
//	//在每次运行结束之后把个人信息清除掉
//	return
//}

//func  (g *GlobalConfig) InitGlobalConfig(c *gin.Context)  {
//	token := c.GetHeader("AccessToken")
//	_shopSn := c.DefaultQuery("_shopSn","")
//	if token == "" {
//		token = c.DefaultQuery("_AccessToken","")//websocket导致的
//		if token == "" {
//			return
//		}
//	}
//	//g = new(GlobalConfig)
//	err := service.FindToken(&g.token, token)
//	if err != nil {
//		return
//	}
//	var shopList []shop.Shop
//	if g.token.User != "" {
//		_ = json.Unmarshal([]byte(g.token.User),&g.user)
//		_ = json.Unmarshal([]byte(g.token.Shop),&shopList)
//		g.shopSn = g.shop.ShopSn
//		if _shopSn != "" {
//			for _, v := range shopList {
//				if v.ShopSn == _shopSn {
//					g.shopSn = _shopSn
//					break
//				}
//			}
//		}
//		for _, v := range shopList {
//			if v.ShopSn == g.shopSn {
//				g.shop = v
//			}
//		}
//		//更新数据库信息 go
//		//go EditToken(c)
//	}
//	//在每次运行结束之后把个人信息清除掉
//	return
//
//}
/**
在数据库里面修改token信息token然后 赋值给user
*/
func editToken(c *gin.Context) {
	auth := authU.GetConfig(c)
	token := c.GetHeader("AccessToken")
	userStr, _ := json.Marshal(auth.User)
	/**
		data.Shop = auth.GlobalToken.Shop
	data.ShopSn = auth.ShopSn
	data.Num = auth.GlobalToken.Num + 1
	*/
	var userToken = model.UserToken{
		User:   string(userStr),
		ShopSn: auth.ShopSn,
		Shop:   auth.Token.Shop,
		Num:    auth.Token.Num + 1,
	}
	_ = service.EditToken(&userToken, token, auth)
	//fmt.Println("修改信息:", err,userToken)
}
