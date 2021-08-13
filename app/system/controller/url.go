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
func addUrl(c *gin.Context) {
	var url model.SystemUrl
	if err := c.ShouldBind(&url); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	auth := authU.GetConfig(c)
	id, err := service.AddUrl(&url, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "添加成功", "data": id})
}

//查询
func findUrl(c *gin.Context) {
	urlid := c.PostForm("urlid")
	if urlid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ID缺失"})
		return
	}
	var url model.SystemUrl
	err := service.FindUrl(&url, urlid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": url,
	})
}

//修改
func editUrl(c *gin.Context) {
	var url model.SystemUrl
	if err := c.ShouldBind(&url); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	err := service.EditUrl(&url)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "修改成功", "data": url})
}

//删除
func delUrl(c *gin.Context) {
	urlid := c.PostForm("urlid")
	if urlid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ID缺失"})
		return
	}
	auth := authU.GetConfig(c)
	err := service.DelUrl(urlid, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "删除成功"})
}

//状态修改
func setUrl(c *gin.Context) {
	urlid := c.PostForm("urlid")
	state := c.PostForm("state")
	if urlid == "" || state == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数缺失"})
		return
	}
	auth := authU.GetConfig(c)
	err := service.SetUrl(urlid, state, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "操作成功"})
}

//列表
func listUrl(c *gin.Context) {
	var obj model.SystemUrl
	var data common.Data
	data.Code = 1
	_ = c.BindQuery(&data)
	_ = c.BindQuery(&obj)
	err := service.ListUrl(&obj, &data)
	if err != nil {
		data.Code = 0
		data.Msg = "获取失败" + err.Error()
	}
	c.JSON(http.StatusOK, data)
}

func funcName(err error) string {
	return err.Error()
}
