package app

import (
	"github.com/gin-gonic/gin"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/common/nsq"
	conIndex "gwsee.com.api/app/index" //公用的部分

	conAct "gwsee.com.api/app/activity/controller" //活动
	conGoods "gwsee.com.api/app/goods/controller"  //商品
	conMem "gwsee.com.api/app/member/controller"   //会员
	conOrder "gwsee.com.api/app/order/controller"  //订单
	conShop "gwsee.com.api/app/shop/controller"    //门店
	conSys "gwsee.com.api/app/system/controller"   //系统
	conTest "gwsee.com.api/app/test/controller"    //测试
	conUser "gwsee.com.api/app/user/controller"    //用户是基础
)

func LoadApi(r *gin.Engine) {
	//加载是否登录--如果登录了 拿到用户信息
	//加载数据库
	common.InitDB()
	nsq.InitNsqProducer()
	// r.Use(common.InitDB)
	//加载缓存 判断是不是用redis缓存 如果用的话 就启动redis缓存
	common.InitCache()
	//r.Use(common.InitCache)
	//根据情况加载具体的服务
	//查看是否有用户 -- 初始化用户信息、
	// r.Use(controller.FindToken)
	//r.Use(common.InitUser)
	r.Use(authU.InitGlobalConfig) //不管用户是否存在 先解析是否有token
	//公用的部分
	conIndex.LoadApi(r)
	//各自可以独立的部分
	conAct.LoadApi(r)
	conGoods.LoadApi(r)
	conMem.LoadApi(r)
	conOrder.LoadApi(r)
	conShop.LoadApi(r)
	conSys.LoadApi(r)
	conTest.LoadApi(r)
	conUser.LoadApi(r)
}
