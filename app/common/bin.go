package common

// 做这个目的主要是为了定时任务的可视化管理
// https://www.cnblogs.com/liuzhongchao/p/9521897.html  这个golang自带的定时器
import (
	"bufio"
	"fmt"
	"gopkg.in/ini.v1"
	"gwsee.com.api/config"
	"gwsee.com.api/utils"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type binFile struct {
	filePath    string
	fileName    string   //
	fileContent []string //一个代表一行
	fileCommand string   // 执行这个文件的命令
	fileExt     string   //文件后缀
}

// 这个主要用于 命令行执行并且返回；与命令文件的生成  和各种命令的执行与停止

//1：命令的执行
func ExecCommand(command string) (err error, data *[]string) {
	cmd := exec.Command("/bin/bash", "-c", command)
	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return
	}
	//执行命令
	if err = cmd.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
		return
	}
	//使用带缓冲的读取器
	outputBuf := bufio.NewReader(stdout)
	var redData []string
	for {
		//一次获取一行,_ 获取当前行是否被读完
		output, _, err := outputBuf.ReadLine()
		if err != nil {
			// 判断是否到文件的结尾了否则出错
			if err.Error() != "EOF" {
				fmt.Printf("Error :%s\n", err)
			}
			break
		}
		fmt.Printf("%s\n", string(output))
		redData = append(redData, string(output))
	}
	//wait 方法会一直阻塞到其所属的命令完全运行结束为止
	if err = cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
		return
	}
	fmt.Println("执行命令 " + command + " 运行结束")
	return err, &redData
}

//2：文件的生成

func BuildCommandFile(path, name, ext string, content []string, commandType string) (err error) {
	var crontabConfig = new(config.BinMgr)
	//0 获取基础的路径
	err = ini.MapTo(crontabConfig, "./config/bin.ini")
	if err != nil {
		return
	}
	basePath := ""
	if commandType == "" || commandType == "CrontabConfig" {
		basePath = crontabConfig.CrontabConfig.Path //这样就限制了只能是定时器了。。
	}
	fmt.Println("basePath:" + basePath)
	var bin binFile
	bin.fileName = name
	bin.filePath = path
	bin.fileContent = content
	bin.fileExt = ext
	//1 判断文件是否存在，存在就 改名字把它标识为已经删除
	nowStr := time.Now().Format("02_150405")
	//2 生成新文件
	var filePath = basePath + "/" + bin.filePath + "/"
	utils.BuildCatalog(filePath)
	var file = filePath + bin.fileName + "." + bin.fileExt // 生成的文件
	flag, _ := utils.PathExists(file)
	i := 0
	var temp string
	for flag { // 如果这个文件已经存在 就给他重命名
		i = i + 1
		temp = file + "_del_" + strconv.Itoa(i) + "_" + nowStr
		flag, _ = utils.PathExists(temp)
	}
	if i > 0 && temp != "" { //代表
		err = os.Rename(file, temp)
		if err != nil {
			fmt.Println("生成命令文件出错(重命名):Rename:" + err.Error())
			return
		}
	}
	//3 文件中写入 command
	fileObj, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("文件的打开出错:" + err.Error())
		return err
	}
	defer fileObj.Close()
	// 然后给文件里面下入信息
	bufReader := bufio.NewWriter(fileObj)
	for _, v := range bin.fileContent {
		bufReader.WriteString(v + "\n")
	}
	bufReader.Flush()                   //不执行Flush，内容不会完全保存在文件中
	ExecCommand("chmod -R 755 " + file) //将其改为可以执行的文件
	return
}

func RenameCommandFile(path, name, ext, newExt, commandType string) (err error) {
	var crontabConfig = new(config.BinMgr)
	//0 获取基础的路径
	err = ini.MapTo(crontabConfig, "./config/bin.ini")
	if err != nil {
		return
	}
	basePath := ""
	if commandType == "" || commandType == "CrontabConfig" {
		basePath = crontabConfig.CrontabConfig.Path //这样就限制了只能是定时器了。。
	}
	var filePath = basePath + "/" + path + "/"
	var file = filePath + name + "." + ext
	flag, _ := utils.PathExists(file)
	if flag {
		newFile := filePath + name + "." + newExt
		err = os.Rename(file, newFile)
		if err != nil {
			fmt.Println("状态变化时候重命名文件出错(重命名):Rename:" + err.Error())
			return
		}
	}
	return
}

//编辑crontab 文件并且重启
func RestartCrontab(list *map[string]string) (err error) {
	var crontabConfig = new(config.BinMgr)
	//0 获取基础的路径
	err = ini.MapTo(crontabConfig, "./config/bin.ini")
	if err != nil {
		return
	}
	basePath := crontabConfig.CrontabConfig.Path
	var commandArr []string
	for k, v := range *list {
		var file = basePath + "/" + v
		flag, _ := utils.PathExists(file)
		// k 是那边的 k_命令的形式组成 这边需要对他进行拆分
		kArr := strings.Split(k, "__")
		if flag && v != "" {
			commandArr = append(commandArr, kArr[1]+" "+basePath+"/"+v)
		}
	}
	//获取当前用户
	err, redData := ExecCommand("whoami")
	var whoami string
	for _, v := range *redData {
		whoami = v
	}
	//crontab是修改的/var/spool/cron/root文件  直接更改这个文件的内容 然后重启
	file := "/var/spool/cron/" + whoami
	ExecCommand("chmod -R 777 " + file)
	// chmod -R 777 /home/user
	fileObj, err := os.OpenFile(file, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("文件的打开出错:" + err.Error())
		return err
	}
	// 然后给文件里面下入信息
	bufReader := bufio.NewWriter(fileObj)
	for _, v := range commandArr {
		bufReader.WriteString(v + "\n")
	}
	bufReader.Flush() //不执行Flush，内容不会完全保存在文件中
	fileObj.Close()
	ExecCommand("chmod -R 600 " + file)
	// 最后重启
	//有人说不需要进行重启操作
	// ExecCommand("service crond restart")
	return
}
