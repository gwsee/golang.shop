package controller

import (
	"github.com/gin-gonic/gin"
	"gwsee.com.api/app/activity/model"
	"gwsee.com.api/app/activity/service"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"net/http"
)

func addActivity(c *gin.Context) {
	var activity model.ActivityData
	if err := c.ShouldBind(&activity); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	auth := authU.GetConfig(c)
	id, err := service.AddActivity(&activity, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "添加成功", "data": id})

}
func delActivity(c *gin.Context) {
	activityid := c.PostForm("activityid")
	if activityid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数缺失"})
		return
	}
	auth := authU.GetConfig(c)
	err := service.DelActivity(activityid, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "删除成功"})
}
func listActivity(c *gin.Context) {
	var obj model.Activity
	var data common.Data
	_ = c.BindQuery(&data)
	_ = c.BindQuery(&obj)
	data.Code = 1
	auth := authU.GetConfig(c)
	err := service.ListActivity(&obj, &data, auth)
	if err != nil {
		data.Code = 0
		data.Msg = "获取失败" + err.Error()
	}
	c.JSON(http.StatusOK, data)
}
func findActivity(c *gin.Context) {
	activityid := c.Query("activityid")
	if activityid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ID缺失"})
		return
	}
	var info model.ActivityData
	err := service.FindActivity(activityid, &info)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{"code": 1, "msg": "", "data": info})
}
func setActivity(c *gin.Context) {
	activityid := c.PostForm("activityid")
	state := c.PostForm("state")
	if activityid == "" || state == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数缺失"})
		return
	}
	auth := authU.GetConfig(c)
	err := service.SetActivity(activityid, state, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "操作成功"})
}
