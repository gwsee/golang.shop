package tencent

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gwsee.com.api/app/shop/model"
	"gwsee.com.api/app/shop/service"
	"io/ioutil"
	"net/http"
)

var url = "https://api.weixin.qq.com/sns/jscode2session"

type wxSessionData struct {
	Appid string `json:"appid"`
	Secret string `json:"secret"`
	JsCode string `json:"js_code"`
	GrantType string `json:"grant_type"`
}

type wxSessionRes struct {
	Openid string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid string `json:"unionid"`
	Errcode int `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func GetSessionKey(c *gin.Context)  {
	var data wxSessionData
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	var wechat model.ShopWechat
	service.FindWechat(&wechat, data.Appid)
	if wechat.WechatAppSecrect == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "小程序信息查询失败"})
		return
	}
	data.Secret =wechat.WechatAppSecrect// "435fc046c70afa90ecf459247622f0de" //有空了改成查数据库
	//435fc046c70afa90ecf459247622f0de
	resp,err:=http.Get(url+"?appid="+data.Appid+"&secret="+data.Secret+"&js_code="+data.JsCode+"&grant_type="+data.GrantType)
	if err!=nil{
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	data2, err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	fmt.Println(string(data2))
	var res wxSessionRes
	err=json.Unmarshal(data2,&res)
	if err!=nil{
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "","data": res})
	return
}