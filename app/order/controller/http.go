package controller

import (
	"github.com/gin-gonic/gin"
	authU "gwsee.com.api/app/auth/user"
)

func LoadApi(r *gin.Engine) {
	//message
	shop := r.Group("/order", authU.NeedShop)
	{
		//role
		shop.GET("/listorder", listOrder)
		shop.GET("/listorderservice", listOrderService)
	}
}
