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

//对外
/**
1:商品展示页
*/
func listCommonsM(c *gin.Context) {
	var obj model.Commons
	var data common.Data
	_ = c.BindQuery(&data)
	_ = c.BindQuery(&obj)
	data.Code = 1
	auth := authU.GetConfig(c)
	err := service.ListCommonsM(&obj, &data, auth)
	if err != nil {
		data.Code = 0
		data.Msg = "获取失败" + err.Error()
	}
	c.JSON(http.StatusOK, data)
}
func findCommonsM(c *gin.Context) {
	commonsid := c.Query("commonsid")
	if commonsid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ID缺失"})
		return
	}
	var info model.CommonsData
	var list []*model2.Classify
	var form []*model.CommonsAttributeForm
	var row string
	err := service.FindCommonsM(&info, &list, &form, &row, commonsid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "data": info, "form": form, "row": row, "classifies": list, "msg": "成功"})
}
