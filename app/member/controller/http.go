package controller

import (
	"github.com/gin-gonic/gin"
	authU "gwsee.com.api/app/auth/user"
)

func LoadApi(r *gin.Engine) {
	member := r.Group("/member", authU.NeedShop)
	{
		//role
		member.GET("/listgrade", listGrade)
		member.POST("/addgrade", addGrade)
		member.POST("/delgrade", delGrade)

		member.GET("/listgradetemplate", listGradeTemplate)
		member.POST("/addgradetemplate", addGradeTemplate)
		member.POST("/delgradetemplate", delGradeTemplate)
	}
}
