package main

/*
// 错误示例:
func (m map[string]string) Set(key string, value string) {
	m[key] = value
}

func Demo1() {
	m := make(map[string]string)
	fmt.Println(m)
	//m.Set("A", "One") //m.Set undefined (type map[string]string has no field or method Set)
}
*/
//参考答案及解析：Unnamed Type 不能作为方法的接收者。昨天我们讲过 Named Type 与 Unamed Type 的区别，就用 Named Type 来修复下代码：
type User map[string]string

func (m User) Set(key string, value string) {
	m[key] = value
}

func Demo2() {
	m := make(User)
	m.Set("A", "One")
}
