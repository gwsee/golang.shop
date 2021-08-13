package controller

import (
	"github.com/gin-gonic/gin"
	authU "gwsee.com.api/app/auth/user"
)

func LoadApi(r *gin.Engine) {
	//message
	shop := r.Group("/shop", authU.NeedShop)
	{
		//role
		shop.GET("/listrole", listRole)
		shop.POST("/addrole", addRole)
		shop.POST("/delrole", delRole)
		//depart
		shop.GET("/listdepart", listDepart)
		shop.POST("/adddepart", addDepart)
		shop.POST("/deldepart", delDepart)
		//user
		shop.GET("/listuser", listUser)
		shop.POST("/adduser", addUser)
		shop.POST("/addaccount", addAccount)
		shop.POST("/deluser", delUser)
		shop.POST("/register", register)

		//
		shop.GET("/list", list)
		shop.GET("/find", find)
		shop.POST("/add", add)
		shop.POST("/del", del)
		//门店类目 品牌管理
		shop.GET("/listclassify", listClassify)
		shop.POST("/addclassify", addClassify)
		shop.POST("/delclassify", delClassify)

		//店铺基础设置挡
		shop.GET("/listsettingdefault", listSettingDefault)
		shop.POST("/addsettingdefault", addSettingDefault)
		shop.POST("/delsettingdefault", setSettingDefault)

		//店铺的服务类型
		shop.GET("/listservicetype", listServiceType)
		shop.POST("/addservicetype", addServiceType)
		shop.POST("/delservicetype", delServiceType)
		shop.POST("/setservicetype", setServiceType)
		//店铺企业微信管理
		shop.GET("/listwechat", listWechat)
		shop.POST("/addwechat", addWechat)
		shop.POST("/delwechat", delWechat)
	}

	//只需要appkey就可以查到的
	r.POST("/shop/findwechat", findWechat)
}
