package main

import "net/http"

func DisplayEcho(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "websockets.html")
}
