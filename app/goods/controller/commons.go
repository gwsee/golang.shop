package controller

import (
	"github.com/gin-gonic/gin"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/goods/model"
	"gwsee.com.api/app/goods/service"
	model2 "gwsee.com.api/app/system/model"
	"net/http"
)

//门店
func addCommons(c *gin.Context) {
	var commons model.CommonsData
	if err := c.ShouldBind(&commons); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	auth := authU.GetConfig(c)
	id, err := service.AddCommons(&commons, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "添加成功", "data": id})
}
func delCommons(c *gin.Context) {
	commonsid := c.PostForm("commonsid")
	if commonsid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数缺失"})
		return
	}
	auth := authU.GetConfig(c)
	err := service.DelCommons(commonsid, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "删除成功"})
}
func listCommons(c *gin.Context) {
	var obj model.Commons
	var data common.Data
	_ = c.BindQuery(&data)
	_ = c.BindQuery(&obj)
	data.Code = 1
	auth := authU.GetConfig(c)
	err := service.ListCommons(&obj, &data, auth)
	if err != nil {
		data.Code = 0
		data.Msg = "获取失败" + err.Error()
	}
	c.JSON(http.StatusOK, data)
}
func findCommons(c *gin.Context) {
	commonsid := c.Query("commonsid")
	if commonsid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ID缺失"})
		return
	}
	var info model.CommonsData
	var list []*model2.Classify
	var form []*model.CommonsAttributeForm
	var row string
	err := service.FindCommons(&info, &list, &form, &row, commonsid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "data": info, "form": form, "row": row, "classifies": list, "msg": "成功"})
}
func setCommons(c *gin.Context) {
	commonsid := c.PostForm("commonsid")
	state := c.PostForm("state")
	if commonsid == "" || state == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数缺失"})
		return
	}
	auth := authU.GetConfig(c)
	err := service.SetCommons(commonsid, state, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "操作成功"})
}
