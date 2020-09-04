package common

import (
	"log"
	"os"
	"strings"
	"time"
)

/**
工具包
 */
//获取根目录
func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
//获取当前时间戳
func GetNowTime() int64{
	return int64(time.Now().Unix())
}
//获取当前日期 时间戳转时间  2006-01-02 15:04:05
func GetNowData(format string,timeUnix int64) string{
	if timeUnix!=0 {
		return time.Unix(timeUnix,0).Format(format)
	}
	return time.Now().Format(format)
}