package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// gorilla/mux包的使用
// 资料来源:https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247484172&idx=1&sn=6dc988c86c3572a8092bdc79feb8d4e8&chksm=fa80d29bcdf75b8d06fc56366352671131c06e1c299a4929a56d7f5ab7137d1e1aec213c5e40&cur_album_id=1323498303014780929&scene=189#wechat_redirect

//go run ./main.go handler.go router.go
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
