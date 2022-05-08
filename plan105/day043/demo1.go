package day043

func Demo1() {
	x := interface{}(nil)
	y := (*int)(nil)
	a := y == x //this comparison is never true; the lhs of the comparison has been assigned a concretely typed value (SA4023)
	b := y == nil
	_, c := x.(interface{}) //type assertion to the same type: x already has type interface{} (S1040)
	println(a, b, c)
	//false true false
	//知识点：类型断言。类型断言语法：i.(Type)，其中 i 是接口，Type 是类型或接口。编译时会自动检测 i 的动态类型与 Type 是否一致。但是，如果动态类型不存在，则断言总是失败
}

var m map[string]int

func Demo2() {
	//m["one"] = 1 //assignment to nil map (SA5000)
	//不能对 nil 的 map 直接赋值，需要使用 make() 初始化。
	m = make(map[string]int)
	m["one"] = 1

}
