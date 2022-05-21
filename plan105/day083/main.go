package day083

import "fmt"

/*
1.同级文件的包名不允许有多个，是否正确？
A. true
B. false
参考答案及解析：A。一个文件夹下只能有一个包，可以多个.go文件，但这些文件必须属于同一个包。
*/

//2.下面的代码有什么问题，请说明
type data struct {
	name string
}

func (p *data) print() {
	fmt.Println("name:", p.name)
}

type printer interface {
	print()
}

func Demo1() {
	d1 := data{"one"}
	d1.print()

	//var i printer = data{} //cannot use (data literal) (value of type data) as printer value in variable declaration: missing method print (print has pointer receiver)
	// 结构体类型 data 没有实现接口 printer。知识点：接口。
	var i printer = &data{"two"}
	i.print()
}
