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
func addDepart(c *gin.Context) {
	var depart model.ShopDepart
	if err := c.ShouldBind(&depart); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	auth := authU.GetConfig(c)
	id, err := service.AddDepart(&depart, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "添加成功", "data": id})
}

//查询
func findDepart(c *gin.Context) {
	urlid := c.PostForm("urlid")
	if urlid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ID缺失"})
		return
	}
	var depart model.ShopDepart
	err := service.FindDepart(&depart, urlid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": depart,
	})
}

//修改
func editDepart(c *gin.Context) {
	var depart model.ShopDepart
	if err := c.ShouldBind(&depart); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	auth := authU.GetConfig(c)
	err := service.EditDepart(&depart, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "修改成功", "data": depart})

}

//删除
func delDepart(c *gin.Context) {
	departid := c.PostForm("departid")
	if departid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数缺失"})
		return
	}
	auth := authU.GetConfig(c)
	err := service.DelDepart(departid, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "删除成功"})

}

//状态修改
func setDepart(c *gin.Context) {
	departid := c.PostForm("departid")
	state := c.PostForm("state")
	if departid == "" || state == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数缺失"})
		return
	}
	auth := authU.GetConfig(c)
	err := service.SetDepart(departid, state, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "操作成功"})

}

//列表
func listDepart(c *gin.Context) {
	var obj model.DepartShop
	var data common.Data
	_ = c.BindQuery(&data)
	_ = c.BindQuery(&obj)
	auth := authU.GetConfig(c)
	if obj.ShopSn == "" {
		obj.ShopSn = auth.ShopSn
	}
	data.Code = 1
	err := service.TreeDepart(&obj, &data)
	if err != nil {
		data.Code = 0
		data.Msg = "获取失败" + err.Error()
	}
	c.JSON(http.StatusOK, data)
}
