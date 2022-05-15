package main

import "fmt"

//官方errors包 属于Wrap Error模式
// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e errorString) Error() string {
	return e.s
}

// New returns an error that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
func NewError(text string) error {
	return &errorString{text} //UnEqual Error:
	//return errorString{text} //Equal Error: 不取地址返回对比,会取里面的s字段去做一个等值的对比
}

//line:6-20以上为errors包代码

var ErrType = NewError("EOF")

func main() {
	if ErrType == NewError("EOF") {
		fmt.Println("Equal Error:", ErrType)
	} else {
		fmt.Println("UnEqual Error:", ErrType)
	}
	//UnEqual Error: EOF
}
