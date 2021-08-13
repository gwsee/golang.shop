package controller

import (
	"github.com/gin-gonic/gin"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/system/model"
	"gwsee.com.api/app/system/service"
	"net/http"
)

//列表
func send(c *gin.Context) {
	//存入数据库
	var obj model.SystemMessage
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	auth := authU.GetConfig(c)
	_, err := service.AddMessage(&obj, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "验证码发送失败：" + err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 2, "msg": "验证码发送成功"})
	}
}

//列表
func valid(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "验证码匹配成功"})
	return
}
