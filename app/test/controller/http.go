package controller

import (
	"github.com/gin-gonic/gin"
)

/**
主要用于测试
*/

func LoadApi(r *gin.Engine) {

	test := r.Group("/test")
	{
		test.GET("/ws/ping", ping)
		test.POST("/message/read", readMessage)
		test.GET("/message/list", listMessage)
		test.POST("/message/del", delMessage)
		test.POST("/message/add", addMessage)
		test.GET("/message/sync", syncMessage)
		test.POST("/common_ajax", commonAjax)
	}
}
