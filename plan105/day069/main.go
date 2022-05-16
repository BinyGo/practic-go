package day069

/*
1.关于 slice 或 map 操作，下面正确的是？
A.
var s []int
s = append(s,1)
B.
var m map[string]int
m["one"] = 1
C.
var s []int
s = make([]int, 0)
s = append(s,1)
D.
var m map[string]int
m = make(map[string]int)
m["one"] = 1
参考答案及解析：ACD。
*/

func Demo2() {
	a, b := test(100)
	a() //100
	b() //110
}

func test(x int) (func(), func()) {
	return func() {
			println(x)
			x += 10
		}, func() {
			println(x)
		}
}
