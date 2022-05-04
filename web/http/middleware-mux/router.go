package main

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router) {
	r.Use(Logging()) //全局应用
	indexRouter := r.PathPrefix("/index").Subrouter()
	indexRouter.Handle("/", &helloHandler{}) //http://localhost:8999/index/

	userRouter := r.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/name/{name}/countries/{country}", ShowVisitorInfo) //http://localhost:8999/user/name/biny/countries/st
	userRouter.Use(Method("GET"))                                              //局部应用

}
