package middleware

import (
	"diabloGin/code"
	"diabloGin/config"
	"diabloGin/controller"
	"github.com/gin-gonic/gin"
)

/**
中间件
 */
func BaseMiddleware() gin.HandlerFunc{
	return func(context *gin.Context) {
		//异常捕获
		defer func() {
			if err := recover(); err != nil {
				var Err *controller.JsonController
				if e,ok := err.(*controller.JsonController); ok {
					Err = e
				}else if e, ok := err.(error); ok {
					//其他错误
					Err = controller.OtherError(e.Error())
				}else{
					//服务器错误
					Err = controller.ServerError
				}
				// 记录一个错误的日志
				if config.LOGGING {
					code.SetLogger(" IP: "+context.ClientIP()+" - "+Err.Msg, config.LOGGINGDIRERROR)
				}
				context.JSON(Err.Code,Err)
				return
			}
		}()

		context.Next()
	}
}
