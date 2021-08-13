package controller

import (
	"github.com/gin-gonic/gin"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/member/model"
	"gwsee.com.api/app/member/service"
	"net/http"
)

func addGradeTemplate(c *gin.Context) {
	var template model.GradeTemplate
	if err := c.ShouldBind(&template); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	auth := authU.GetConfig(c)
	id, err := service.AddGradeTemplate(&template, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "添加成功", "data": id})

}
func delGradeTemplate(c *gin.Context) {
	id := c.PostForm("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数缺失"})
		return
	}
	auth := authU.GetConfig(c)
	err := service.DelGradeTemplate(id, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "删除成功"})

}
func listGradeTemplate(c *gin.Context) {
	var obj model.GradeTemplate
	var data common.Data
	data.Code = 1
	_ = c.BindQuery(&data)
	_ = c.BindQuery(&obj)
	auth := authU.GetConfig(c)
	err := service.ListGradeTemplate(&obj, &data, auth)
	if err != nil {
		data.Code = 0
		data.Msg = "获取失败" + err.Error()
	}
	c.JSON(http.StatusOK, data)
}
