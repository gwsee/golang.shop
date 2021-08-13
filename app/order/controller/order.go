package controller

import (
	"github.com/gin-gonic/gin"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/order/model"
	"gwsee.com.api/app/order/service"
	"net/http"
)

func listOrder(c *gin.Context) {
	var obj model.OrderData
	var data common.Data
	_ = c.BindQuery(&data)
	_ = c.BindQuery(&obj)
	auth := authU.GetConfig(c)
	err := service.ListOrder(&obj, &data, auth)
	data.Code = 1
	if err != nil {
		data.Code = 0
		data.Msg = "获取失败" + err.Error()
	}
	c.JSON(http.StatusOK, data)
}
func listOrderService(c *gin.Context) {
	var obj model.OrderServiceData
	var data common.Data
	_ = c.BindQuery(&data)
	_ = c.BindQuery(&obj)
	auth := authU.GetConfig(c)
	err := service.ListOrderService(&obj, &data, auth)
	data.Code = 1
	if err != nil {
		data.Code = 0
		data.Msg = "获取失败" + err.Error()
	}
	c.JSON(http.StatusOK, data)
}
