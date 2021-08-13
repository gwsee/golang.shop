package controller

import (
	"github.com/gin-gonic/gin"
	authU "gwsee.com.api/app/auth/user"
)

func LoadApi(r *gin.Engine) {
	activity := r.Group("/activity", authU.NeedShop)
	{
		activity.POST("addactivity", addActivity)
		activity.POST("delactivity", delActivity)
		activity.GET("listactivity", listActivity)
		activity.GET("findactivity", findActivity)
		activity.POST("setactivity", setActivity)

		activity.GET("listlog", listLog)
	}
}
