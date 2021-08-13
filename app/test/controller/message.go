package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/common/nsq"
	"gwsee.com.api/app/test/model"
	"gwsee.com.api/app/test/service"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func addMessage(c *gin.Context) {
	var message model.Message
	if err := c.ShouldBind(&message); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	auth := authU.GetConfig(c)
	id, err := service.AddMessage(&message, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "添加成功", "data": id})

}
func delMessage(c *gin.Context) {
	messageid := c.PostForm("messageid")
	if messageid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "参数缺失"})
		return
	}
	auth := authU.GetConfig(c)
	err := service.DelMessage(messageid, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 2, "msg": "删除成功"})
}
func readMessage(c *gin.Context) {
	messageid := c.PostForm("messageid")
	if messageid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "参数缺失"})
		return
	}
	auth := authU.GetConfig(c)
	err := service.ReadMessage(messageid, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 1, "msg": "读.."})

}
func listMessage(c *gin.Context) {
	var obj model.MessageLoad
	var data common.Data
	_ = c.BindQuery(&data)
	_ = c.BindQuery(&obj)
	data.Code = 1
	auth := authU.GetConfig(c)
	err := service.ListMessage(&obj, &data, auth)
	if err != nil {
		data.Code = 0
		data.Msg = "获取失败" + err.Error()
	}
	c.JSON(http.StatusOK, data)
}

func syncMessage(c *gin.Context) {
	//c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "获取失败" + err.Error()})
	//return
	//使用websocket 进行消息即时发送
	//也使用nsq进行消息读取
	auth := authU.GetConfig(c)
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("获取失败1" + err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "获取失败" + err.Error()})
		return
	}
	defer ws.Close()
	topic := "testMessage"
	channel := "receiverId_" + strconv.FormatUint(auth.User.UserId, 10)
	err = nsq.InitConsumer(topic, channel, ws, auth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "获取失败" + err.Error()})
		return
	}

	select {}
}

func commonAjax(c *gin.Context) {
	type ajax struct {
		Method   string ` form:"method" json:"method"`
		Url      string ` form:"url" json:"url"`
	}

	var data ajax
	data.Url = c.PostForm("url")
	dataPram:=c.PostFormMap("data")
	fmt.Println(dataPram,"dataPPPPPPPPP")
	fileValue, err := c.FormFile("file")
	if err==nil{
		file:="./runtime/"+fileValue.Filename
		c.SaveUploadedFile(fileValue,file)
		res,err:=HttpFileFunction(data.Url,file,dataPram)
		if err!=nil{
			c.JSON(http.StatusBadRequest, "文件上传失败："+err.Error())
		}else{
			var obj interface{}
			json.Unmarshal(res,&obj)
			c.JSON(http.StatusOK, obj)
		}
	}else{
		res,err:=Action("POST",data.Url,dataPram)
		if err!=nil{
			c.JSON(http.StatusBadRequest, err.Error())
		}else{
			var obj interface{}
			json.Unmarshal(res,&obj)
			c.JSON(http.StatusOK, obj)
		}
	}

	return
}


func Action(method, urls string,  postData map[string]string) (by []byte, err error) {
	fmt.Println(method,urls,postData)
	var req *http.Request
	val := url.Values{}
	for k, v := range postData {
		val.Add(k, v)
	}
	req, err = http.NewRequest(method, urls+"?"+val.Encode(),  strings.NewReader(val.Encode()))
	if err != nil {
		err = fmt.Errorf("http.NewRequest is fail: %v", err.Error())
		return
	}
	headers:=make(map[string]string)
	headers["Cookie"]="PHPSESSID=lfmjrrn5oknnfglk74igg8h8a5"
	headers["Host"]="cs1.jianceyun.net"
	headers["Origin"]="http://cs1.jianceyun.net"
	headers["Referer"]=urls
	headers["Accept"]="application/json, text/javascript, */*; q=0.01"
	headers["User-Agent"]="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36"
	headers["X-Requested-With"]="XMLHttpRequest"
	headers["Content-Type"]="application/x-www-form-urlencoded"
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("client.Do is fail: %v", err.Error())
		return
	}
	by, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("ioutil.ReadAll is fail: %v", err.Error())
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New("post amazon HTTP CODE:" + fmt.Sprint(resp.StatusCode))
		return
	}
	return
}
//在上传文件的实例中,bodyWrite不能在http.DO()之后 ,也就是这里的Client.Do())
func HttpFileFunction(Url, file_name string,mp map[string]string) (by []byte, err error) {
	bodyBuf := &bytes.Buffer{}
	bodyWrite := multipart.NewWriter(bodyBuf)

	for k,v:=range mp{
		bodyWrite.WriteField(k,v)
	}

	file, err := os.Open(file_name)
	defer file.Close()
	if err != nil {
		fmt.Println("加载文件失败", err)
		return
	}
	fileWrite, err := bodyWrite.CreateFormFile("files", file_name)
	_,err = io.Copy(fileWrite,file)
	if err != nil {
		fmt.Println("io Copy error",err)
		return
	}
	contentType := bodyWrite.FormDataContentType()
	bodyWrite.Close() //正确位置            ✓
	request, err := http.NewRequest("POST", Url, bodyBuf)

	headers:=make(map[string]string)
	headers["Cookie"]="PHPSESSID=lfmjrrn5oknnfglk74igg8h8a5"
	headers["Host"]="cs1.jianceyun.net"
	headers["Origin"]="http://cs1.jianceyun.net"
	headers["Referer"]=Url
	headers["Accept"]="application/json, text/javascript, */*; q=0.01"
	headers["User-Agent"]="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36"
	headers["X-Requested-With"]="XMLHttpRequest"
	headers["Content-Type"]="application/x-www-form-urlencoded"
	for k, v := range headers {
		request.Header.Add(k, v)
	}
	if err != nil {
		fmt.Println("http newrequest error",err)
		return
	}
	request.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		err = fmt.Errorf("client.Do is fail: %v", err.Error())
		return
	}
	by, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("ioutil.ReadAll is fail: %v", err.Error())
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New("post amazon HTTP CODE:" + fmt.Sprint(resp.StatusCode))
		return
	}
	return
}
