package controller

import (
	"github.com/gin-gonic/gin"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/system/model"
	"gwsee.com.api/app/system/service"
	"net/http"
)

// 应该赋值一个默认的信息给他
func findWebsite(c *gin.Context) {
	host := c.PostForm("host")
	if host == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ID缺失"})
		return
	}
	var web model.SystemWebsite
	err := service.FindWebsite(&web, host)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": web,
	})

}
func listWebsite(c *gin.Context) {
	var obj model.SystemWebsite
	var data common.Data
	data.Code = 1
	_ = c.BindQuery(&data)
	_ = c.BindQuery(&obj)
	err := service.ListWebsite(&obj, &data)
	if err != nil {
		data.Code = 0
		data.Msg = "获取失败" + err.Error()
	}
	c.JSON(http.StatusOK, data)
}
func delWebsite(c *gin.Context) {
	webId := c.PostForm("id")
	if webId == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ID缺失"})
		return
	}
	auth := authU.GetConfig(c)
	err := service.DelWebsite(webId, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "删除成功"})
}
func addWebsite(c *gin.Context) {
	var web model.SystemWebsite
	if err := c.ShouldBind(&web); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	auth := authU.GetConfig(c)
	id, err := service.AddWebsite(&web, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "添加成功", "data": id})
}
