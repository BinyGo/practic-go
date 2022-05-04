package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// gorilla/mux包middleware的使用
// 资料来源:https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247484180&idx=1&sn=b66497c5428c25577068f18132b2d59d&chksm=fa80d283cdf75b95cc49d08c56d0fa9b00c47d0457c894be3ca3ea02bd3c404cfb5312fa1d93&cur_album_id=1323498303014780929&scene=189#wechat_redirect

//go run ./main.go handler.go router.go middleware.go
func main() {
	muxRouter := mux.NewRouter()

	RegisterRoutes(muxRouter)

	server := &http.Server{
		Addr:    ":8999",
		Handler: muxRouter,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("ListenAndServe err")
	}
}
