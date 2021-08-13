package controller

import (
	"github.com/gin-gonic/gin"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/shop/model"
	"gwsee.com.api/app/shop/service"
	"net/http"
)

//新增

//新增
func add(c *gin.Context) {
	var shop model.Shop
	if err := c.ShouldBind(&shop); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	auth := authU.GetConfig(c)
	id, err := service.Add(&shop, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "操作成功", "data": id})
}

//查询
func find(c *gin.Context) {
	shopsn := c.Query("shopsn")
	if shopsn == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ID缺失"})
		return
	}
	var shop model.Shop
	err := service.Find(&shop, shopsn)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": shop,
	})
}

//修改
func edit(c *gin.Context) {
	var shop model.Shop
	if err := c.ShouldBind(&shop); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	auth := authU.GetConfig(c)
	err := service.Edit(&shop, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "修改成功", "data": shop})

}

//删除
func del(c *gin.Context) {
	shopid := c.PostForm("shopid")
	if shopid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ID缺失"})
		return
	}
	auth := authU.GetConfig(c)
	err := service.Del(shopid, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "删除成功"})

}

//状态修改
func set(c *gin.Context) {
	shopid := c.PostForm("shopid")
	state := c.PostForm("state")
	if shopid == "" || state == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数缺失"})
		return
	}
	auth := authU.GetConfig(c)
	err := service.Set(shopid, state, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "操作成功"})

}

//列表
func list(c *gin.Context) {
	var obj model.Shop
	var data common.Data
	_ = c.BindQuery(&data)
	_ = c.BindQuery(&obj)
	auth := authU.GetConfig(c)
	err := service.List(&obj, &data, auth)
	data.Code = 1
	if err != nil {
		data.Code = 0
		data.Msg = "获取失败" + err.Error()
	}
	c.JSON(http.StatusOK, data)
}
