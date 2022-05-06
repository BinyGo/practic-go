package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", Login)
	mux.HandleFunc("/logout", Logout)
	mux.HandleFunc("/secret", Secret)
	http.ListenAndServe(":8998", mux)
}

//1.登录
//curl -XPOST -d 'name=biny&password=123456' -c - http://localhost:8999/login
//2.带上登录接口返回的user-session
//curl --cookie "user-session=MTU4m..." http://localhost:8998/secret
//3.退出登录
//curl --cookie "user-session=MTU4m..." http://localhost:8998/logout
