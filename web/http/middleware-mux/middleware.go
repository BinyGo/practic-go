package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Logging() mux.MiddlewareFunc {
	// 创建中间件
	return func(h http.Handler) http.Handler {
		// 创建一个新的handler包装http.HandlerFunc
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 中间件的处理逻辑
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			// 调用下一个中间件或者最终的handler处理程序
			h.ServeHTTP(w, r)
		})
	}
}

func Method(m string) mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}
