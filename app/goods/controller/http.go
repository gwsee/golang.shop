package controller

import (
	"github.com/gin-gonic/gin"
	authU "gwsee.com.api/app/auth/user"
)

func LoadApi(r *gin.Engine) {
	//门店的
	goods := r.Group("/goods", authU.NeedShop)
	{
		//商品类型
		goods.GET("/listclassify", listClassify)
		goods.POST("/addclassify", addClassify)
		goods.POST("/delclassify", delClassify)
		//商品操作
		goods.GET("/listcommons", listCommons)
		goods.GET("/findcommons", findCommons)
		goods.POST("/addcommons", addCommons)
		goods.POST("/setcommons", setCommons)
		goods.POST("/delcommons", delCommons)
	}
	goodsM := r.Group("/m/goods", authU.NeedShopSn)
	{
		//门店商品列表
		goodsM.GET("/listcommons", listCommonsM)
		goodsM.GET("/findcommons", findCommonsM)
	}

}
