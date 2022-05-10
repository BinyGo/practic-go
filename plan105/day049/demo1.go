package day049

import (
	"encoding/json"
	"fmt"
)

//1.下面代码输出什么？
func Demo1() {
	var ch chan int
	select {
	case v, ok := <-ch:
		println(v, ok)
	default:
		println("default")
	}
	//参考答案及解析：default。ch 为 nil，读写都会阻塞。
}

//2.下面这段代码输出什么？
type People struct {
	Name string `json:"name"`
}

func Demo2() {
	js := `{
		"name":"seekload"
	}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(p) //{seekload}
}
