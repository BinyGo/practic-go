package main

import (
	"fmt"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, sessionCookieName)

	fmt.Println("authenticated:", session.Values["authenticated"])

	session.Values["authenticated"] = false //保存不生效?bug?
	session.Options.MaxAge = 0
	session.ID = "22"

	session.Save(r, w)

}
