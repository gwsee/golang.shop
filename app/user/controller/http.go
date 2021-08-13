package controller

import (
	"github.com/gin-gonic/gin"
	authU "gwsee.com.api/app/auth/user"
)

func LoadApi(r *gin.Engine) {
	r.GET("/user/info", authU.NeedUser, info)
	r.GET("/user/test", test) //用于测试异步执行是否需要关闭数据库（想办法做到没有线程的时候才关闭数据库）
	r.POST("/user/edit", edit)
	r.POST("/user/editps", editPs)
	r.POST("/login/login", login)
	r.POST("/login/register", register) //目前没地方用到
	// r.POST("/login/registers", RegisterS)
}
