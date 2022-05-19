package day070

/*
1.关于字符串连接，下面语法正确的是？
A. str := ‘abc’ + ‘123’
B. str := “abc” + “123”
C. str ：= ‘123’ + “abc”
D. fmt.Sprintf(“abc%d”, 123)
参考答案及解析：BD。知识点：单引号、双引号和字符串连接。在 Go 语言中，双引号用来表示字符串 string，其实质是一个 byte 类型的数组，单引号表示 rune 类型。
*/

// 2.下面代码能编译通过吗？可以的话，输出什么？

func Demo1() {
	println(DeferTest1(1)) //4
	println(DeferTest2(1)) //3
}

func DeferTest1(i int) (r int) {
	r = i
	defer func() {
		r += 3
	}()
	return r
}

func DeferTest2(i int) (r int) {
	defer func() {
		r += i
	}()

	return 2
}