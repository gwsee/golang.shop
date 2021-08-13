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
func addRole(c *gin.Context) {
	var role model.ShopRole
	if err := c.ShouldBind(&role); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	auth := authU.GetConfig(c)
	id, err := service.AddRole(&role, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "添加成功", "data": id})
}

//查询
func findRole(c *gin.Context) {
	roleid := c.PostForm("roleid")
	if roleid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ID缺失"})
		return
	}
	var role model.ShopRole
	err := service.FindRole(&role, roleid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": role,
	})
}

//修改
func editRole(c *gin.Context) {
	var role model.ShopRole
	if err := c.ShouldBind(&role); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	auth := authU.GetConfig(c)
	err := service.EditRole(&role, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "修改成功", "data": role})

}

//删除
func delRole(c *gin.Context) {
	roleid := c.PostForm("roleid")
	if roleid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ID缺失"})
		return
	}
	auth := authU.GetConfig(c)
	err := service.DelRole(roleid, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "删除成功"})

}

//状态修改
func setRole(c *gin.Context) {
	roleid := c.PostForm("roleid")
	state := c.PostForm("state")
	if roleid == "" || state == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数缺失"})
		return
	}
	auth := authU.GetConfig(c)
	err := service.SetRole(roleid, state, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "操作成功"})

}

//列表
func listRole(c *gin.Context) {
	var obj model.ShopRole
	var data common.Data
	_ = c.BindQuery(&data)
	_ = c.BindQuery(&obj)
	auth := authU.GetConfig(c)
	err := service.ListRole(&obj, &data, auth)
	data.Code = 1
	if err != nil {
		data.Code = 0
		data.Msg = "获取失败" + err.Error()
	}
	c.JSON(http.StatusOK, data)
}
