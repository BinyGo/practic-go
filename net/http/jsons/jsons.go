package jsons

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func ParseJsonRequestHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	// 将请求体中的 JSON 数据解析到结构体中
	// 发生错误，返回400 错误码
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "User:%+v", user)
}

// router.go
//userRouter.HandleFunc("/parse_json_request", jsons.ParseJsonRequestHandler)
//curl -X POST -d '{"name": "biny", "age": 18}'  -H "Content-Type: application/json" http://localhost:8999/index/parse_json_request

func WriteJsonREsponsehandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name: "biny",
		Age:  19,
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&user)
	if err != nil {
		return
	}
}

// router.go
//userRouter.HandleFunc("/get_json_response", handler.WriteJsonResponseHandler)
//curl -X GET http://localhost:8999/index/get_json_response
//{"name":"biny","age":19}
