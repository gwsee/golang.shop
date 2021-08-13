package tencent

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"gwsee.com.api/utils"
	"net/http"
)

type wxSecretData struct {
	AppId string `json:"app_id"`
	SessionKey string `json:"session_key"`
	EncryptedData string `json:"encrypted_data"`
	Iv string `json:"iv"`
}
func WxGetPhone(c *gin.Context )  {
	//1 验证参数
	var data wxSecretData
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	userData:=make(map[string]interface{})
	err :=decrypt(data,userData)
	if err!=nil{
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	//var user wxUser
	c.JSON(200, gin.H{"code": 1, "msg": "成功", "data": userData})
}
func decrypt(data wxSecretData,userData map[string]interface{})(err error)  {
	sessionKey,err:=base64.StdEncoding.DecodeString(data.SessionKey)
	if err!=nil{
		return
	}
	aesIv,err:=base64.StdEncoding.DecodeString(data.Iv)
	if err!=nil{
		return
	}
	encryptedData,err:=base64.StdEncoding.DecodeString(data.EncryptedData)
	if err!=nil{
		return
	}
	decryptedText,err :=utils.DesDecryption(sessionKey,aesIv,encryptedData)
	if err!=nil{
		return
	}
	//3 绑定数据
	err = json.Unmarshal(decryptedText,&userData)
	if err!=nil{
		return
	}
	appid:=userData["watermark"].(map[string]interface{})["appid"]
	if appid != data.AppId{
		err = errors.New("appid不匹配:"+err.Error())
		return
	}
	return
}
//第二种


