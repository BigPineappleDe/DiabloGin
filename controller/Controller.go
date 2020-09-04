package controller

import (
	"diabloGin/code"
	"diabloGin/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
控制器基类
*/

//成功
func Success(data interface{}, ctx *gin.Context) {
	var jsonController JsonController
	jsonController.Code = 1
	jsonController.Msg = "成功！"
	jsonController.Data = data
	ctx.JSON(http.StatusOK, jsonController)
	return
}

//失败
func Error(msg string, ctx *gin.Context) {
	var jsonController JsonController
	jsonController.Code = 1
	jsonController.Msg = msg
	jsonController.Data = nil
	ctx.JSON(http.StatusOK, jsonController)
	return
}

type JsonController struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

var (
	ServerError = OutJson(-1, nil, "系统异常，请稍后重试!")
	NotFound    = OutJson(0, nil, http.StatusText(http.StatusNotFound))
)

//其他异常
func OtherError(message string) *JsonController {
	return OutJson(0, nil, message)
}

// 404处理
func HandleNotFound(c *gin.Context) {
	err := NotFound
	if config.LOGGING {
		code.SetLogger(" IP: "+c.ClientIP()+" - "+"访问未定义路由", config.LOGGINGDIRERROR)
	}
	c.JSON(err.Code, err)
	return
}

func OutJson(statusCode int, data interface{}, msg string) *JsonController {
	return &JsonController{
		Code: statusCode,
		Data: data,
		Msg:  msg,
	}
}
