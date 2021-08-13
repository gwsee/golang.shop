package controller

import (
	"github.com/gin-gonic/gin"
	"gwsee.com.api/app/activity/model"
	"gwsee.com.api/app/activity/service"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"net/http"
)

func listLog(c *gin.Context) {
	var obj model.Log
	var data common.Data
	_ = c.BindQuery(&data)
	_ = c.BindQuery(&obj)
	data.Code = 1
	auth := authU.GetConfig(c)
	err := service.ListLog(&obj, &data, auth)
	if err != nil {
		data.Code = 0
		data.Msg = "获取失败" + err.Error()
	}
	c.JSON(http.StatusOK, data)
}
