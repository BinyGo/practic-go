package main

import (
	"fmt"
	"net/http"
)

//http://localhost:8999/display_headers?aa=cc&bb=dd
func DisplayHeadersHandler(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Domain:   "example.com", //设置的Cookie需要被其他子域名的服务访问
		HttpOnly: true,          //为避免跨域脚本 (XSS) 攻击，通过JavaScript的API无法访问带有 HttpOnly 标记的Cookie，它们只应该发送给服务端。
	}
	http.SetCookie(w, &c)
	//很多cookie设置: https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247484287&idx=1&sn=54baf9ab1739b2001a2e1aec6d6a2d66&chksm=fa80d2e8cdf75bfe8110856f8719be79db94a7de96fe771b9d38ea1d9a93957e7122c5fce780&cur_album_id=1323498303014780929&scene=189#wechat_redirect

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
	http.ListenAndServe(":8998", nil)
	//http://localhost:8999/display_headers
}
