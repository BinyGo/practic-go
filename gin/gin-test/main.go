package main

import "github.com/gin-gonic/gin"

func main() {
	http := Router()
	http.Run()
}

func Router() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "ping")
	})
	return router
}
