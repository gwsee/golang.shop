package controller

import (
	"github.com/gin-gonic/gin"
	authU "gwsee.com.api/app/auth/user"
)

func LoadApi(r *gin.Engine) {
	//系统级管理员的
	system := r.Group("/system", authU.NeedAdmin)
	{
		//菜单管理  --完成
		system.GET("/listmenu", listMenu)
		system.POST("/addmenu", addMenu)
		system.POST("/delmenu", delMenu)
		system.POST("/editmenu", editMenu)
		system.POST("/setmenu", setMenu)
		system.GET("/findmenu", findMenu)

		//页面管理  --完成
		system.GET("/listurl", listUrl)
		system.POST("/addurl", addUrl)
		system.POST("/delurl", delUrl)
		system.POST("/editurl", editUrl)
		system.POST("/seturl", setUrl)
		system.GET("/findurl", findUrl)
		//品牌 --完成
		system.GET("/listbrand", listBrand)
		system.POST("/addbrand", addBrand)
		system.POST("/delbrand", delBrand)
		//品牌组 --完成
		system.GET("/listbrandgroup", listBrandGroup)
		system.POST("/addbrandgroup", addBrandGroup)
		system.POST("/delbrandgroup", delBrandGroup)
		system.GET("/listgroupbrand", listGroupBrand) //根据分组获取品牌 方便一次性进行管理
		//属性组 --完成
		system.GET("/listattribute", listAttribute)
		system.POST("/addattribute", addAttribute)
		system.POST("/delattribute", delAttribute)
		//类目  --完成

		system.GET("/treeclassify", treeClassify)
		system.POST("/addclassify", addClassify)
		system.POST("/delclassify", delClassify)
		//根类目属性管理 查询出来 然后进行编辑操作
		system.POST("/findclassifyattribute", findClassifyAttribute)
		system.POST("/updateclassifyattribute", updateClassifyAttribute)
		//给根类目关联品牌
		system.GET("/listclassifybrand", listClassifyBrand)
		system.POST("/addclassifybrand", addClassifyBrand)
		system.POST("/delclassifybrand", delClassifyBrand)
		//设置类目关联品牌的 样板
		system.GET("/listclassifytemplate", listClassifyTemplate)
		system.GET("/findclassifytemplate", findClassifyTemplate)
		system.POST("/addclassifytemplate", addClassifyTemplate)
		system.POST("/delclassifytemplate", delClassifyTemplate)
		//网站管理相关的
		system.GET("/listwebsite", listWebsite)
		system.POST("/addwebsite", addWebsite)
		system.POST("/delwebsite", delWebsite)
		//网站定时器相关
		system.GET("/listtimerclassify", listTimerClassify)
		system.POST("/deltimerclassify", delTimerClassify)
		system.POST("/addtimerclassify", addTimerClassify)
		system.GET("/listtimer", listTimer)
		system.POST("/addtimer", addTimer)
		system.POST("/deltimer", delTimer)
		system.GET("/listtimerlog", listTimerLog)
		system.POST("/deltimerlog", delTimerLog)
		system.POST("/addtimerlog", addTimerLog)

	}
	//系统级对外开放的公用接口
	common := r.Group("/system")
	{
		system.GET("/usermenu", userMenu)
		common.GET("/listclassify", listClassify)
		common.GET("/findwebsite", findWebsite)
		common.POST("/send", send)
		common.POST("/valid", valid)
	}
}
