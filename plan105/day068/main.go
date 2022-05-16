package day068

//1.下面代码有什么问题吗？
func Demo1() {
	/* for i := 0; i < 10; i++ {
	loop:
		println(i)
	}
	goto loop //goto loop jumps into block */
	//参考答案及解析：goto 不能跳转到其他函数或者内层代码。编译报错：
}

func Demo2() {
	x := []int{0, 1, 2}
	y := [3]*int{}
	for i, v := range x {
		defer func() {
			print(v) //2 2 2
		}()
		y[i] = &v
	}
	print(*y[0], *y[1], *y[2]) //2 2 2

}
