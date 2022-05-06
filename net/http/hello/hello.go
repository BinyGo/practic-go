package main

// http包的使用
// 资料来源:https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247484112&idx=1&sn=79d0d3167d0d962fe41ec00cdafffbb0&chksm=fa80d347cdf75a51183182f14622af766538ca0c5335012e5e1cc50b100e78f2954fa3943770&token=1184133103&lang=zh_CN&scene=21#wechat_redirect

import (
	"fmt"
	"net/http"
)

type WelcomeHandlerStruct struct{}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello biny")
}

func DemoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "demo biny")
}

func (*WelcomeHandlerStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloHandler)
	mux.HandleFunc("/demo", DemoHandler)
	mux.Handle("/welcome", &WelcomeHandlerStruct{})
	http.ListenAndServe(":8999", mux)
}

/*
func HellowHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello biny")
	//访问http://localhost:8999 返回:hello biny
}

func main() {
	http.HandleFunc("/", HellowHandler)
	http.ListenAndServe(":8999", nil)
}
*/
