package oss

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"strings"
)

// 这里获取 oss相关配置信息
type ossConfigs struct {
	bucket   string //bucket
	endpoint string //endpoint
	key      string //accessKeyId
	secrect  string // accessKeySecret
	site     string // url
}
type FileObject struct {
	FileName  string
	FileValue string
	FileByte  []byte
	FileType  string
}

var ossInfo ossConfigs
var bucket *oss.Bucket

func initOss() (err error) {
	ossInfo.bucket = "gwsee-go"
	ossInfo.endpoint = "oss-cn-beijing.aliyuncs.com"
	ossInfo.key = ""
	ossInfo.secrect = ""
	ossInfo.site = "https://gwsee-go.oss-cn-beijing.aliyuncs.com/"
	//ossInfo.bucket = "basetemp"
	//ossInfo.endpoint = "oss-cn-beijing.aliyuncs.com"
	//ossInfo.key = "LTAIqFIZ8Rb58Cw3"
	//ossInfo.secrect = "3hjdIBsa6LDuHChl39gLCH539WDGTE"

	client, err := oss.New(ossInfo.endpoint, ossInfo.key, ossInfo.secrect)
	if err != nil {
		return
	}
	// 获取存储空间。
	bucket, err = client.Bucket(ossInfo.bucket)
	return
}
func Upload(file *FileObject) (name string, err error) {
	err = initOss()
	if err != nil {
		return
	}
	if file.FileType == "local" {
		err = uploadLocal(file)
	} else if file.FileType == "byte" {
		err = uploadByte(file)
	} else if file.FileType == "string" {
		err = uploadString(file)
	} else {
		err = uploadFile(file)
	}
	name = ossInfo.site + file.FileName
	return
}

//上传字符串
func uploadString(file *FileObject) (err error) {
	// 创建OSSClient实例。
	// 指定存储类型为标准存储，缺省也为标准存储。
	storageType := oss.ObjectStorageClass(oss.StorageStandard)
	// 指定存储类型为归档存储。
	// storageType := oss.ObjectStorageClass(oss.StorageArchive)
	// 指定访问权限为公共读，缺省为继承bucket的权限。
	objectAcl := oss.ObjectACL(oss.ACLPublicRead)
	// 上传字符串。
	err = bucket.PutObject(file.FileName, strings.NewReader(file.FileValue), storageType, objectAcl)
	if err != nil {
		return
	}
	return
}

//上传Byte数组
func uploadByte(file *FileObject) (err error) {
	// 创建OSSClient实例。
	// 上传Byte数组。
	err = bucket.PutObject(file.FileName, bytes.NewReader(file.FileByte))
	//if err != nil {
	//	fmt.Println("Error:", err)
	//}
	return
}

//上传文件流
func uploadFile(file *FileObject) (err error) {
	// 创建OSSClient实例。
	// 读取本地文件。
	fd, err := os.Open(file.FileValue)
	if err != nil {
		//fmt.Println("Error:", err)
		//os.Exit(-1)
		return
	}
	defer fd.Close()

	// 上传文件流。
	err = bucket.PutObject(file.FileName, fd)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	os.Exit(-1)
	//}
	return
}

//上传本地文件
func uploadLocal(file *FileObject) (err error) {
	err = bucket.PutObjectFromFile(file.FileName, file.FileValue)
	return
}
