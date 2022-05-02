package day027

import "fmt"

//1.下面这段代码输出什么？

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) string() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

func Demo1() {
	fmt.Println(South)          //2
	fmt.Println(South.string()) //South
}

// 参考答案及解析：South。知识点：iota 的用法、类型的 String() 方法。

//2.下面代码输出什么？

type Math struct {
	x, y int
}

var m = map[string]Math{
	"foot": {2, 3},
}

func Demo2() {
	//m["foo"].x = 5 //cannot assign to struct field m["foo"].x in map
	tmp := m["foo"]
	tmp.x = 4
	m["foo"] = tmp
	fmt.Println(m["foo"].x)
}

var ma = &Math{2, 3}
var m2 = map[string]*Math{
	"foot": ma,
}

func Demo3() {
	m2["foo"].x = 4
	//error: invalid memory address or nil pointer dereference [recovered] 无效的内存地址或nil指针解除引用
	//panic: runtime error: invalid memory address or nil pointer dereference 无效的内存地址或零指针解除引用
}
