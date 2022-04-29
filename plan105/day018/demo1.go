package day018

import "fmt"

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
		//r = t + 5  这里是r,结果为10
	}()
	return t
}

func f3() (r int) {
	defer func(r int) int {
		r = r + 5 //这里改变的传入的r
		return r
	}(r)
	return 1
}

func Demo1() {
	//返回具名参数
	fmt.Println(f1()) //1
	fmt.Println(f2()) //5
	fmt.Println(f3()) //1
}
