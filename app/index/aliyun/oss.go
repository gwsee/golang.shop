package aliyun

import (
	"github.com/gin-gonic/gin"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/extend/aliyun/oss"
	"gwsee.com.api/utils"
	"net/http"
	"path/filepath"
	"strconv"
)

type UploadConfig struct {
	FileName string `form:"name" json:"name"`
	FileType string `form:"type" json:"type" binding:"required"`
	FileSign string `form:"sign" json:"sign" binding:"required"`
}

/**
  outer 是指的最外层的目录
*/

func buildFileName(ext string, config *UploadConfig, auth *authU.GlobalConfig) (catalogName, name string) {
	shopsn := auth.ShopSn
	if shopsn == "" {
		shopsn = "users" //指定用户上传的文件
	}
	name = strconv.FormatUint(auth.User.UserId, 10) + utils.Krand(6, 3)
	dayStr := utils.GetTimeStr("ymd", "")
	timeStr := utils.GetTimeStr("his", "")
	catalogName = shopsn + "/" + config.FileSign + "/" + config.FileType + "/" + dayStr
	// 在这之前判断临时目录是否存在 不存在就生成目录
	name = timeStr + name + ext
	return
}
func Upload(c *gin.Context) {
	//1:获取最基本的参数 与文件信息
	var config UploadConfig
	if err := c.ShouldBind(&config); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	fileValue, err := c.FormFile(config.FileName)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	auth := authU.GetConfig(c)
	// 1.2获取与文件命名
	ext := filepath.Ext(fileValue.Filename)
	path, name := buildFileName(ext, &config, auth)
	// 1.2读取文件存到字节中然后上传到服务器
	fileCountent, _ := fileValue.Open()
	var byteContent []byte
	byteContent = make([]byte, 1000000)
	fileCountent.Read(byteContent)
	// 2:将文件保存到本地
	// 2.1:生成文件名
	//fullPath := "public/files/" + path
	//utils.BuildCatalog(fullPath)
	//// 2.2:保存到本地
	//c.SaveUploadedFile(fileValue, fullPath+"/"+name)
	// 3:上传到oss
	var fileObject oss.FileObject
	fileObject.FileName = path + "/" + name
	// fileObject.FileValue = fullPath + "/" + name
	fileObject.FileByte = byteContent
	fileObject.FileType = "byte" // local 就用2*  byte  就用1.2
	name, err = oss.Upload(&fileObject)
	// 4:删除本地文件
	// 5:返回oss数据
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 1, "msg": "成功", "data": name})
}
