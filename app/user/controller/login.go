package controller

import (
	"github.com/gin-gonic/gin"
	authU "gwsee.com.api/app/auth/user"
	shopMo "gwsee.com.api/app/shop/model"
	shopSer "gwsee.com.api/app/shop/service"
	userMo "gwsee.com.api/app/user/model"
	"gwsee.com.api/app/user/service"
	"net/http"
	"strconv"
)

//只有这个登录的地方
/**
	登陆简介
	1：登陆来源有 门店 也有 客户
    2：验证方式有 账密 电话号码与验证码 电话号码与unionid

*/
func login(c *gin.Context) {
	//入参处理
	var obj struct {
		Acc       string `json:"account"`
		Pws       string `json:"password" `
		Mobile    string `json:"mobile"`
		Captcha   string `json:"captcha"`
		LoginType string `json:"_loginType"`
		UnionId   string `json:"unionid"`
		Type      string `json:"type"`
	}
	err := c.ShouldBindJSON(&obj)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	var user userMo.User
	if obj.Type == "mobile" {
		user, err = service.CheckCode(obj.Mobile, obj.Captcha)
	} else if obj.Type == "unionid" {
		user, err = service.CheckUnicode(obj.Mobile, obj.UnionId)
	} else {
		user, err = service.CheckAccount(obj.Acc, obj.Pws)
	}
	//检查账密是否正确

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	var shopList []*shopMo.Shop
	shopSn := ""
	if obj.LoginType == "_shop" {
		//处理门店信息
		err = shopSer.ListShopUser(&shopList, "shop.is_del=0 and shop_user.is_del=0 and user_id="+strconv.FormatUint(user.UserId, 10))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "该账户未绑定门店",
			})
			return
		}
		for _, v := range shopList {
			if shopSn == "" {
				shopSn = v.ShopSn
				break
			}
		}
		if shopSn == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "未申请门店",
			})
			return
		}
	} else if obj.LoginType == "_wxApp" {
		var shop shopMo.Shop
		auth := authU.GetConfig(c)
		err = shopSer.Find(&shop, auth.ShopSn)
		if shop.ShopSn == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "数据异常",
			})
			return
		}
		shopList = append(shopList, &shop)
		shopSn = auth.ShopSn
	}
	//登录成功获取token
	token, err := service.AddToken(&user, &shopList, shopSn)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "token异常",
		})
	}
	msg := "登录成功"
	if user.UserName != "" {
		msg = "欢迎" + user.UserName + "回来！"
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   2,
		"msg":    msg,
		"_token": token,
		"user":   user,
	})
}

// 所有注册 都只需要 姓名(昵称) 账号 密码 电话号码
func register(c *gin.Context) {
	var data userMo.User
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	if data.UserName == "" || data.UserAccount == "" || data.UserPassword == "" || data.UserMobile == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数不全",
		})
		return
	}
	auth := authU.GetConfig(c)
	_, err := service.AddUser(&data, auth, true)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 2,
		"msg":  "注册成功",
		"data": nil,
	})
}
func loginOut() {

}
