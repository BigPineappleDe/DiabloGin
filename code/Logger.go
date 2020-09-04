package code

import (
	"fmt"
	"diabloGin/common"
	"diabloGin/config"
	"os"
)

/**
日志模块  消息  文件名  0保存到根目录
*/
func SetLogger(msg string, fileName string) {
	FileUrlDir := common.GetCurrentPath() + "/"
	if config.ISLOGGINGDIROUT {
		if config.ISLOGGINGDIR==""{
			FileUrlDir += "../"
		}else{
			FileUrlDir = config.ISLOGGINGDIR
		}
	}
	//路径
	FileUrl := FileUrlDir + config.LOGGINGFILEURL + "/" + fileName + "/"
	//有无都创建
	_ = os.MkdirAll(FileUrl, os.ModeDir|os.ModePerm)
	//写入文件末尾 无则创建
	logFile, err := os.OpenFile(FileUrl+common.GetNowData("2006_01_02",0)+".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer logFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	_, _ = logFile.Write([]byte("[ " + common.GetNowData("2006-01-02 15:04:05",0) + " ] " + msg + "\r\n"))

}
