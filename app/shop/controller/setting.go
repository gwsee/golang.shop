package controller

import (
	"github.com/gin-gonic/gin"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/shop/model"
	"gwsee.com.api/app/shop/service"
	model2 "gwsee.com.api/app/system/model"
	"net/http"
)

//新增
func addSettingDefault(c *gin.Context) {

	type Default struct {
		model.ShopSettingDefault
		Name string ` form:"name" json:"name"`
	}
	var obj Default
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	obj.DefaultKey = obj.Name
	auth := authU.GetConfig(c)
	id, err := service.AddSettingDefault(&obj.ShopSettingDefault, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "添加成功", "data": id})
}

//列表
func listSettingDefault(c *gin.Context) {
	var obj model2.SystemSettingType
	var data common.Data
	_ = c.BindQuery(&data)
	_ = c.BindQuery(&obj)
	auth := authU.GetConfig(c)
	err := service.ListSettingDefault(&obj, &data, auth)
	data.Code = 1
	if err != nil {
		data.Code = 0
		data.Msg = "获取失败" + err.Error()
	}
	c.JSON(http.StatusOK, data)
}

//状态修改
func setSettingDefault(c *gin.Context) {
	keys := c.PostForm("keys")
	state := c.PostForm("state")
	if keys == "" || state == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数缺失"})
		return
	}
	auth := authU.GetConfig(c)
	err := service.SetSettingDefault(keys, state, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "操作成功"})
}
