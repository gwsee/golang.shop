package common

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gwsee.com.api/utils"
)

type logger struct {
	filePath string
	fileName string
	maxSize  int64
	fileObj  *os.File
}

func (f *logger) initFile() (err error) {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/runtime/log/" + now.Format("200601") + "/"
	}
	utils.BuildCatalog(logFilePath)
	logFileName := now.Format("02") + ".log"
	fileName := logFilePath + logFileName
	//如果文件不存在 应该创建文件么？
	f.filePath = logFilePath
	f.fileName = logFileName
	f.maxSize = 102400
	fileObj, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("日志文件出错（初始化）:initFile:" + err.Error())
		return
	}
	f.fileObj = fileObj
	//文件处理
	newObj, err := f.splitFile(f.fileObj)
	if err != nil {
		fmt.Println("日志文件出错（切割）:initFile:" + err.Error())
		return
	}
	f.fileObj = newObj
	return
}
func (f *logger) splitFile(file *os.File) (*os.File, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("日志文件出错(信息):splitFile:" + err.Error())
		return nil, err
	}
	if fileInfo.Size() < f.maxSize {
		return file, nil
	}
	nowStr := time.Now().Format("02_150405")
	logName := path.Join(f.filePath, fileInfo.Name())
	newLogName := path.Join(f.filePath, nowStr+".log")
	// 如果文件已经存在了  就给重复命名？
	flag, _ := utils.PathExists(newLogName)
	i := 0
	for flag {
		i = i + 1
		newLogName = f.filePath + nowStr + "_" + strconv.Itoa(i) + ".log"
		flag, _ = utils.PathExists(newLogName)
	}
	file.Close()
	err = os.Rename(logName, newLogName)
	if err != nil {
		fmt.Println("日志文件出错(重命名):Rename:" + err.Error())
		return file, err
	}
	fileObj, err := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("日志文件出错(重新设定):splitFile:" + err.Error())
		return nil, err
	}
	return fileObj, nil
}
func Logger() (*logrus.Logger, logger) {
	var log logger
	err := log.initFile()
	if err != nil {
		fmt.Println("err", err)
	}
	//实例化
	logger := logrus.New()
	//设置输出
	logger.Out = log.fileObj
	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return logger, log
}
func HandleLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqBody, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		logger, log := Logger()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		reqUse := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由 ---就已经有了参数
		reqUrl := c.Request.RequestURI
		reqUrlData := strings.Split(reqUrl, "?")
		reqUrl = reqUrlData[0]
		reqGet := ""
		if len(reqUrlData) > 1 {
			reqGet = reqUrlData[1]
		}
		// 请求头
		reqHeader := c.Request.Header
		// 请求文本
		reqPostForm := c.Request.PostForm
		if len(reqPostForm) != 0 {
			reqBody = nil
		}
		// 状态码
		reqCode := c.Writer.Status()
		reqSql := GetLog()
		// 请求IP
		reqIP := c.ClientIP()
		logger.WithFields(logrus.Fields{
			"\r\n[0]reqUrl":      reqUrl,
			"\r\n[1]reqMethod":   reqMethod,
			"\r\n[2]reqCode":     reqCode,
			"\r\n[3]reqUse":      reqUse,
			"\r\n[4]reqIP":       reqIP,
			"\r\n[5]reqHeader":   reqHeader,
			"\r\n[6]reqGet":      reqGet,
			"\r\n[7]reqPostForm": reqPostForm,
			"\r\n[8]reqBody":     string(reqBody),
			"\r\n[9]reqSql":      reqSql,
		}).Info("日志")
		log.fileObj.Close()
	}
}
