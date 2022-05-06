package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	wsRouter := r.PathPrefix("/ws").Subrouter()
	wsRouter.HandleFunc("/echo", EchoMessage)
	wsRouter.HandleFunc("/echo_display", DisplayEcho)
}

func main() {
	muxRouter := mux.NewRouter()

	RegisterRoutes(muxRouter)

	server := &http.Server{
		Addr:    ":8999",
		Handler: muxRouter,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("ListenAndServe err")
	}
}
