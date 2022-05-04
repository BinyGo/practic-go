package main

import (
	"github.com/gorilla/mux"
)

func ResisterRoutes(r *mux.Router) {
	indexRouter := r.PathPrefix("/index").Subrouter()
	indexRouter.Handle("/", &helloHandler{}) //http://localhost:8999/index/

	userRouter := r.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/name/{name}/countries/{country}", ShowVisitorInfo) //http://localhost:8999/user/name/biny/countries/st
}
