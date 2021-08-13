package controller

import (
	"github.com/gin-gonic/gin"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/user/model"
	"gwsee.com.api/app/user/service"
	"log"
	"net/http"
	"time"
)

//根据userid 查询指定人信息
func info(c *gin.Context) {
	account := c.DefaultQuery("account", "")
	var info model.User
	err := service.GetInfo(&info, account)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "人员信息",
		"data": info,
	})
}

//修改个人信息--除开密码
func edit(c *gin.Context) {
	var data model.User
	//form(ShouldBind) -- 需要在结构体字段对应 比如 加个form :"XXXX" 或者 直接用名字
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	auth := authU.GetConfig(c)
	err := service.EditUser(&data, auth.User.UserId, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "修改成功", "data": data})
}

//修改个人信息--密码
func editPs(c *gin.Context) {
	var obj struct {
		Acc     string `json:"account" binding:"required"`
		Pws     string `json:"password" binding:"required"`
		Mobile  string `json:"mobile" binding:"required"`
		Captcha string `json:"captcha" binding:"required"`
	}
	err := c.BindJSON(&obj)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	auth := authU.GetConfig(c)
	err = service.EditPs(obj.Pws, obj.Acc, obj.Mobile, obj.Captcha, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "修改成功", "data": nil})
}

//获取当前登录者信息
func myInfo(c *gin.Context) {
	var info model.User
	auth := authU.GetConfig(c)
	account := string(auth.User.UserAccount)
	err := service.GetInfo(&info, account)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "",
		"data": info,
	})
}
func test(c *gin.Context) {
	copyContext := c.Copy() // 使用副本
	common.Wg.Add(1)        //开启一个异步
	go func() {
		defer common.Wg.Done()
		time.Sleep(3 * time.Second)
		account := copyContext.DefaultQuery("account", "")
		var info model.User
		err := service.GetInfo(&info, account)
		if err != nil {
			return
		}
		log.Printf("异步执行：%s,得到的结果是：%v\n", copyContext.Request.URL.Path, info)
	}()
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "测试异步",
	})
}
