package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	http := Router()
	http.Run(":8999")
}

func Router() *gin.Engine {
	//router := gin.New() 默认的没有中间件的空白 Gin
	router := gin.Default() // 默认已经连接了 Logger and Recovery 中间件
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "biny"})
	})

	router.GET("user/:name", func(ctx *gin.Context) {
		nameL := ctx.Param("name")
		ctx.String(http.StatusOK, nameL)
	})

	// 查询字符串参数使用现有的底层 request 对象解析
	// 请求响应匹配的 URL: /welcome?firstname=biny&lastname=Go
	router.GET("/welcome", func(ctx *gin.Context) {
		firstname := ctx.DefaultQuery("firstname", "Guest")
		lastname := ctx.Query("lastname")
		fmt.Println(firstname, lastname)
		ctx.String(http.StatusOK, firstname+lastname)
	})

	return router
}
