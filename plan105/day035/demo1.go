package day035

import "fmt"

// 1.关于 bool 变量 b 的赋值，下面错误的用法是？
// A. b = true
// B. b = 1
// C. b = bool(1)
// D. b = (1 == 2)
// 参考答案及解析：BC。

//2.关于GetPodAction定义，下面赋值正确的是
type TransInfo struct{}

type Fragment interface {
	Exec(transInfo *TransInfo) error
}

type GetPodAction struct{}

func (g GetPodAction) Exec(transInfo *TransInfo) error {
	return nil
}

func Demo1() {
	var fragment1 Fragment = new(GetPodAction)
	//var fragment2 Fragment = GetPodAction //GetPodAction (type) is not an expression
	var fragment3 Fragment = &GetPodAction{}
	var fragment4 Fragment = GetPodAction{}

	fmt.Println(fragment1, fragment3, fragment4)
}
