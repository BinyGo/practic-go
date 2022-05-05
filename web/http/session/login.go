package main

import (
	"fmt"
	"net/http"
)

var sessionCookieName = "user-session"

func Login(w http.ResponseWriter, r *http.Request) {
	session, err := sessionStore.Get(r, sessionCookieName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 登录验证
	name := r.FormValue("name")
	pass := r.FormValue("password")

	fmt.Println(name, pass)

	// 在session中标记用户已经通过登录验证
	session.Values["authenticated"] = true
	err = session.Save(r, w)

	fmt.Fprintln(w, "登录成功!", err)
}
