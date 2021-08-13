package controller

import (
	"github.com/gin-gonic/gin"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/system/model"
	"gwsee.com.api/app/system/service"
	"net/http"
)

//类目管理
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
func listClassify(c *gin.Context) {
	var obj model.Classify
	var data common.Data
	data.Code = 1
	_ = c.BindQuery(&data)
	_ = c.BindQuery(&obj)
	err := service.ListClassify(&obj, &data)
	if err != nil {
		data.Code = 0
		data.Msg = "获取失败" + err.Error()
	}
	c.JSON(http.StatusOK, data)

}
func treeClassify(c *gin.Context) {
	var obj model.Classify
	var data common.Data
	data.Code = 1
	_ = c.BindQuery(&data)
	_ = c.BindQuery(&obj)
	err := service.TreeClassify(&obj, &data)
	if err != nil {
		data.Code = 0
		data.Msg = "获取失败" + err.Error()
	}
	c.JSON(http.StatusOK, data)
}
func delClassify(c *gin.Context) {
	id := c.PostForm("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ID缺失"})
		return
	}
	auth := authU.GetConfig(c)
	err := service.DelClassify(id, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "删除成功"})

}

//根类目属性管理
func findClassifyAttribute(c *gin.Context) {
	classifyId := c.PostForm("classifyid")
	state := c.PostForm("state")
	if classifyId == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ID缺失"})
		return
	}
	var classify model.ClassifyAndAttribute
	err := service.FindClassifyAttribute(&classify, classifyId, state)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{"code": 1, "data": classify, "msg": "成功"})
}
func updateClassifyAttribute(c *gin.Context) {
	var classify model.ClassifyAndAttribute
	if err := c.ShouldBind(&classify); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	// cannot unmarshal string into Go struct
	// field ClassifyAttributeAndItem.items.items of type model.ClassifyAttributeItem
	auth := authU.GetConfig(c)
	err := service.UpdateClassifyAttribute(&classify, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "添加成功", "data": 0})
}

//给根类目绑定品牌
func addClassifyBrand(c *gin.Context) {
	var ids model.ClassifyBrandIds
	if err := c.ShouldBind(&ids); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	auth := authU.GetConfig(c)
	id, err := service.AddClassifyBrand(&ids, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "设置成功", "data": id})

}
func listClassifyBrand(c *gin.Context) {
	var obj model.ClassifyBrand
	var data common.Data
	data.Code = 1
	_ = c.BindQuery(&data)
	_ = c.BindQuery(&obj)
	err := service.ListClassifyBrand(&obj, &data)
	if err != nil {
		data.Code = 0
		data.Msg = "获取失败" + err.Error()
	}
	c.JSON(http.StatusOK, data)
}
func delClassifyBrand(c *gin.Context) {
	id := c.PostForm("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ID缺失"})
		return
	}
	auth := authU.GetConfig(c)
	err := service.DelClassifyBrand(id, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "删除成功"})

}

//给根类目品牌设置品牌模板
func addClassifyTemplate(c *gin.Context) {
	var classify model.ClassifyTemplateAndAttribute
	if err := c.ShouldBind(&classify); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	auth := authU.GetConfig(c)
	id, err := service.AddClassifyTemplate(&classify, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "添加成功", "data": id})

}
func listClassifyTemplate(c *gin.Context) {
	var obj model.ClassifyTemplate
	var data common.Data
	data.Code = 1
	_ = c.BindQuery(&data)
	_ = c.BindQuery(&obj)
	err := service.ListClassifyTemplate(&obj, &data)
	if err != nil {
		data.Code = 0
		data.Msg = "获取失败" + err.Error()
	}
	c.JSON(http.StatusOK, data)
}
func findClassifyTemplate(c *gin.Context) {
	templateid := c.PostForm("templateid")
	if templateid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ID缺失"})
		return
	}
	var info model.ClassifyTemplateInfo
	err := service.FindClassifyTemplate(&info, templateid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{"code": 1, "msg": "成功"})
}
func delClassifyTemplate(c *gin.Context) {
	id := c.PostForm("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ID缺失"})
		return
	}
	auth := authU.GetConfig(c)
	err := service.DelClassifyTemplate(id, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "删除成功"})
}
