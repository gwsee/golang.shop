package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/shop/model"
	"gwsee.com.api/app/shop/service"
	userMo "gwsee.com.api/app/user/model"
	"net/http"
)

//这个只是增加门店和用户之间的关系
func addUser(c *gin.Context) {
	var user model.ShopUser
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	auth := authU.GetConfig(c)
	id, err := service.AddUser(&user, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "添加成功", "data": id})
}

//这个只是新增用户然后与用户做绑定关系
func addAccount(c *gin.Context) {
	var obj model.ShopUserDetail
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	auth := authU.GetConfig(c)
	id, err := service.AddAccount(&obj, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "添加成功", "data": id})
}

//这个是新增门店然后与新增用户 两者做绑定关系
func register(c *gin.Context) {
	var data struct {
		*model.Shop
		*userMo.User
		Code        string `json:"captcha"`
		IsAccounted string `json:"is_accounted"`
	}
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	fmt.Println(data.User.UserPassword, "绑定的密码")
	fmt.Println(data.User)
	//绑定参数 然后丢给service进行处理
	auth := authU.GetConfig(c)
	err := service.Register(data.Shop, data.User, data.IsAccounted, data.Code, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 2,
		"msg":  "注册成功",
		"data": nil,
	})
}

//查询
func findUser(c *gin.Context) {
	id := c.PostForm("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ID缺失"})
		return
	}
	var user model.ShopUser
	err := service.FindUser(&user, id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": user,
	})
}

//修改
func editUser(c *gin.Context) {
	var user model.ShopUser
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	auth := authU.GetConfig(c)
	err := service.EditUser(&user, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "修改成功", "data": user})

}

//删除
func delUser(c *gin.Context) {
	id := c.PostForm("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "缺失"})
		return
	}
	auth := authU.GetConfig(c)
	err := service.DelUser(id, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "删除成功"})

}

//状态修改
func setUser(c *gin.Context) {
	id := c.PostForm("id")
	state := c.PostForm("state")
	if id == "" || state == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数缺失"})
		return
	}
	auth := authU.GetConfig(c)
	err := service.SetUser(id, state, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "操作成功"})

}

//列表
func listUser(c *gin.Context) {
	var obj model.ShopUserDetail
	var data common.Data
	_ = c.BindQuery(&data)
	_ = c.BindQuery(&obj)
	auth := authU.GetConfig(c)
	err := service.ListUser(&obj, &data, auth)
	data.Code = 1
	if err != nil {
		data.Code = 0
		data.Msg = "获取失败" + err.Error()
	}
	c.JSON(http.StatusOK, data)
}
