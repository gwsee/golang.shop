package controller

import (
	"github.com/gin-gonic/gin"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/shop/model"
	"gwsee.com.api/app/shop/service"
	"net/http"
)

func addServiceType(c *gin.Context) {
	var obj model.ShopServiceType
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	auth := authU.GetConfig(c)
	id, err := service.AddServiceType(&obj, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "添加成功", "data": id})
}
func delServiceType(c *gin.Context) {
	typeid := c.PostForm("typeid")
	if typeid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ID缺失"})
		return
	}
	auth := authU.GetConfig(c)
	err := service.DelServiceType(typeid, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "删除成功"})
}
func listServiceType(c *gin.Context) {
	var obj model.ShopServiceType
	var data common.Data
	_ = c.BindQuery(&data)
	_ = c.BindQuery(&obj)
	auth := authU.GetConfig(c)
	err := service.ListServiceType(&obj, &data, auth)
	data.Code = 1
	if err != nil {
		data.Code = 0
		data.Msg = "获取失败" + err.Error()
	}
	c.JSON(http.StatusOK, data)
}
func setServiceType(c *gin.Context) {
	typeid := c.PostForm("typeid")
	state := c.PostForm("state")
	if typeid == "" || state == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数缺失"})
		return
	}
	auth := authU.GetConfig(c)
	err := service.SetServiceType(typeid, state, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "操作成功"})

}
