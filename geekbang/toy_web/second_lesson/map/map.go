package main

import "fmt"

func main() {
	// 创建了一个预估容量是2的 map
	m := make(map[string]string, 2)
	// 没有指定预估容量
	m1 := make(map[string]string)
	// 直接初始化
	m2 := map[string]string{
		"biny": "go",
	}

	//赋值
	m["hello"] = "world"
	m1["hello"] = "world"
	m2["hello"] = "world"
	// 取值
	val := m["hello"]
	println(val)

	// comm,ok写法取值，使用两个返回值，后面的ok会告诉你map有没有这个key
	val, ok := m["invalid_key"]
	if !ok {
		println("key not found")
	}
	println("invalid_key is :", val)

	for key, val := range m {
		fmt.Printf("%s =>%s \n", key, val)
	}

}
