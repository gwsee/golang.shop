package utils

import (
	"os"
)

func BuildCatalog(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 必须分成两步
		// 先创建文件夹
		/// fmt.Println("创建目录",path) Mkdir是单级目录
		os.MkdirAll(path, 0777)
		// 再修改权限
		os.Chmod(path, 0777)
	}
}
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
