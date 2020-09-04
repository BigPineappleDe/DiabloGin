package controller

import (
	"diabloGin/code"
	"fmt"
	"github.com/gin-gonic/gin"
)

type PublicController struct {

}
/**
公共类
*/
//测试方法
func (c *PublicController) Test(ctx *gin.Context) {
	da, _ := code.RedisInit().Get("abc").Result()
	fmt.Println(da)
	
	Success("数据", ctx)
}
