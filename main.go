package main

import (
	"github.com/gin-gonic/gin"
	"gwsee.com.api/app"
	"gwsee.com.api/app/common"
)

func main() {
	r := gin.Default()
	//日志中间件 -- 日志中间件中使用 analysis 进行数据处理（多线程的方式 不能影响正常数据处理响应） https://studygolang.com/articles/27098?fr=sidebar
	//待思考：使用kafka 数据先丢进去 先响应后再插入数据库（想法 ）暂未做
	//服务注册中间件 注册响应的服务
	//加载数据库并且并加载apiAPI
	r.Use(common.HandleLogger())

	app.LoadApi(r)
	r.Run(":800")
}
