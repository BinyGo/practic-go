package main

import (
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8999", nil)

	//http://localhost:8999/static/static.css
	//http://localhost:8999/static/static.js
	//http://localhost:8999/static/index.html
}
