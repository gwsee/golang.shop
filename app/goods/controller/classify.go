package controller

import (
	"github.com/gin-gonic/gin"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/goods/model"
	"gwsee.com.api/app/goods/service"
	"net/http"
)

func addClassify(c *gin.Context) {
	var classify model.Classify
	if err := c.ShouldBind(&classify); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	auth := authU.GetConfig(c)
	id, err := service.AddClassify(&classify, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "添加成功", "data": id})

}
func delClassify(c *gin.Context) {
	classifyid := c.PostForm("classifyid")
	if classifyid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数缺失"})
		return
	}
	auth := authU.GetConfig(c)
	err := service.DelClassify(classifyid, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "删除成功"})

}
func listClassify(c *gin.Context) {
	var obj model.Classify
	var data common.Data
	_ = c.BindQuery(&data)
	_ = c.BindQuery(&obj)
	data.Code = 1
	auth := authU.GetConfig(c)
	err := service.ListClassify(&obj, &data, auth)
	if err != nil {
		data.Code = 0
		data.Msg = "获取失败" + err.Error()
	}
	c.JSON(http.StatusOK, data)
}
