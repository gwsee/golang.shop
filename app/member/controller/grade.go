package controller

import (
	"github.com/gin-gonic/gin"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/member/model"
	"gwsee.com.api/app/member/service"
	"net/http"
)

func addGrade(c *gin.Context) {
	var grade model.Grade
	if err := c.ShouldBind(&grade); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	auth := authU.GetConfig(c)
	id, err := service.AddGrade(&grade, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "添加成功", "data": id})
}
func delGrade(c *gin.Context) {
	gradeid := c.PostForm("gradeid")
	if gradeid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数缺失"})
		return
	}
	auth := authU.GetConfig(c)
	err := service.DelGrade(gradeid, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "删除成功"})

}
func listGrade(c *gin.Context) {
	var obj model.Grade
	var data common.Data
	_ = c.BindQuery(&data)
	_ = c.BindQuery(&obj)
	data.Code = 1
	auth := authU.GetConfig(c)
	err := service.ListGrade(&obj, &data, auth)
	if err != nil {
		data.Code = 0
		data.Msg = "获取失败" + err.Error()
	}
	c.JSON(http.StatusOK, data)
}
