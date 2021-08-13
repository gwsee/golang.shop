package index

import (
	"github.com/gin-gonic/gin"
	"gwsee.com.api/app/index/aliyun"
	"gwsee.com.api/app/index/tencent"
)

func LoadApi(r *gin.Engine) {
	index := r.Group("/index")
	{
		index.POST("/aliyun/upload", aliyun.Upload)
	}


	m := r.Group("/m")
	{
		m.POST("/index/tencent/get_phone", tencent.WxGetPhone)
		m.POST("/index/tencent/get_session_key", tencent.GetSessionKey)
	}
}
