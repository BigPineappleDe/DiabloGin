package route

import (
	"diabloGin/config"
	"diabloGin/controller"
	"diabloGin/middleware"
	"github.com/gin-gonic/gin"
	"strconv"
)

/**
路由文件
*/
func Init() {
	r := gin.Default()
	r.NoMethod(controller.HandleNotFound)
	r.NoRoute(controller.HandleNotFound)
	r.Use(middleware.BaseMiddleware())
	//分组路由
	vApi := r.Group("/api")
	{
		pApi := vApi.Group("/public")
		{
			var p *controller.PublicController
			pApi.GET("/index", p.Test)
		}

	}
	_ = r.Run(":"+strconv.Itoa(config.PORT)) // listen and serve on 0.0.0.0:8080
}
