package controller

import (
	"github.com/gin-gonic/gin"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/system/model"
	"gwsee.com.api/app/system/service"
	"net/http"
)

//新增
func addMenu(c *gin.Context) {
	var menu model.SystemMenu
	if err := c.ShouldBind(&menu); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	auth := authU.GetConfig(c)
	id, err := service.AddMenu(&menu, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "添加成功", "data": id})

}

//查询
func findMenu(c *gin.Context) {
	menuid := c.PostForm("menuid")
	if menuid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "menuid缺失"})
		return
	}
	var menu model.SystemMenu
	err := service.FindMenu(&menu, menuid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": menu,
	})
}

//修改
func editMenu(c *gin.Context) {
	var menu model.SystemMenu
	if err := c.ShouldBind(&menu); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	auth := authU.GetConfig(c)
	err := service.EditMenu(&menu, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "修改成功", "data": 0})
}

//删除
func delMenu(c *gin.Context) {
	menuid := c.PostForm("menuid")
	if menuid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "menuid缺失"})
		return
	}
	auth := authU.GetConfig(c)
	err := service.DelMenu(menuid, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "删除成功"})
}

//状态修改
func setMenu(c *gin.Context) {
	menuid := c.PostForm("menuid")
	state := c.PostForm("state")
	if menuid == "" || state == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数缺失"})
		return
	}
	auth := authU.GetConfig(c)
	err := service.SetMenu(menuid, state, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "操作成功"})
}

//列表
func listMenu(c *gin.Context) {
	var obj model.MenuUrl
	var data common.Data
	data.Code = 1
	_ = c.BindQuery(&data)
	_ = c.BindQuery(&obj)
	err := service.TreeMenu(&obj, &data)
	if err != nil {
		data.Code = 0
		data.Msg = "获取失败" + err.Error()
	}
	c.JSON(http.StatusOK, data)
}

func userMenu(c *gin.Context) {
	var id uint64
	var ids []uint64
	var first model.MenuUrl
	var menu []model.MenuUrl
	var data []model.MenuData
	menuid, _ := c.GetQuery("menuid")
	id, ids, first, menu, data, err := service.UserMenu(menuid)
	if err != nil {
		c.JSON(200, gin.H{"code": 0, "msg": "获取失败" + err.Error()})
	}
	c.JSON(200, gin.H{"code": 1, "msg": "操作成功", "data": gin.H{"id": id, "ids": ids, "first": first, "menu": menu, "data": data}})
}
