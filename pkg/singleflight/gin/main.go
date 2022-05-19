package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
)

func main() {
	http := Router()
	http.Run(":8999")
}

func Router() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", pong)

	return router
}

var singleFly = new(singleflight.Group)

func pong(ctx *gin.Context) {
	key := "biny"

	result, _, shared := singleFly.Do(key, func() (interface{}, error) {
		err := getData(ctx)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"err": err})
		}
		msg, _ := ctx.Get("message")
		ctx.JSON(http.StatusOK, gin.H{"message": msg})
		return ctx, err
	})

	fmt.Printf("get result: %#v, shared: %v\n", result, shared)
	time.Sleep(time.Second)
	ctx.JSON(http.StatusOK, gin.H{"message": "redis获取或者降级处理"})
}

func getData(ctx *gin.Context) error {
	fmt.Printf("request time: %d\n", time.Now().UnixMicro())
	ctx.Set("message", fmt.Sprintf("source data from rpcServer,time:%d", time.Now().UnixMicro()))
	time.Sleep(time.Second)
	return nil
}

func Client() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "http://localhost:8999/ping", nil)
	if err != nil {
		log.Fatal(err)
	}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)

}
