package main

import (
	"fmt"
	"net/http"
)

//http://localhost:8999/display_headers?aa=cc&bb=dd
func DisplayHeadersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Method: %s Protocol: %s \n", r.Method, r.Proto)

	// 遍历所有请求头
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header field %q,value %q \n", k, v)
	}

	fmt.Fprintf(w, "Host = %q \n", r.Host)                          //"localhost:8999"
	fmt.Fprintf(w, "RemoteAddr = %q \n", r.RemoteAddr)              //"[::1]:53718" 客户端随机端口
	fmt.Fprintf(w, "RequestURI = %q \n", r.RequestURI)              //"/display_headers?aa=cc&bb=dd"
	fmt.Fprintf(w, "URL.Path = %q \n", r.URL.Path)                  //"/display_headers"
	fmt.Fprintf(w, "URL.RawQuery = %q \n", r.URL.RawQuery)          //"aa=cc&bb=dd"
	fmt.Fprintf(w, "URL.Query = %q \n", r.URL.Query())              //map["aa":["cc"] "bb":["dd"]]
	fmt.Fprintf(w, "URL.Query aa = %q \n", r.URL.Query().Get("aa")) //"cc"
	fmt.Fprintf(w, "Form = %q \n", r.Form)                          //map[]

	// 通过 Key 获取指定请求头的值
	fmt.Fprintf(w, "\n Finding value of \"Accept\" %q\n\n", r.Header["Accept"])

	// 遍历所有cookie
	for _, cookie := range r.Cookies() {
		fmt.Fprintf(w, "Cookie field %q,value %q \n", cookie.Name, cookie.Value)
	}

}

func main() {
	http.HandleFunc("/display_headers", DisplayHeadersHandler)
	http.ListenAndServe(":8999", nil)
}
