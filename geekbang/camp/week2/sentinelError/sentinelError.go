package main

import (
	"errors"
	"fmt"
)

// Sentinel Error 哨兵错误处理

// 定义的特定错误，我们叫 sentinel error，这个名字来源于计算机编程中使用一个特定值来表示不
// 可能进行进一步处理的做法。所以对于 Go，我们使用特定的值来表示错误。
// if err == ErrSomething { … }
// 类似的 io.EOF，更底层的 syscall.ENOENT。

// 结论: 尽可能避免 sentinel errors。
// 建议是避免在编写的代码中使用 sentinel errors。在标准库中有一些使用它们的情况，但这不是一个你应该模仿的模式

//Create a named type for our new error type.
type errorString string //string别名

//Implement the error interface.
func (e errorString) Error() string {
	return string(e)
}

func NewError(s string) error {
	return errorString(s) //传入的string转为别名errorString类型,别名errorString带Error()方法,符合error interface
}

//需要提前简历各种错误
var (
	HttpNotFound   = NewError("EOF")
	HttpRequestBad = NewError("BAD")
	ErrNamedType   = NewError("EOF")   //自己定义的error
	ErrStructType  = errors.New("EOF") //标准库errors,errors.New()返回指针
)

func main() {
	if ErrNamedType == NewError("EOF") { //返回的都是相同的errorString类型字符串,把底层的2个字符串对比匹配,所以为true
		fmt.Println("Named Type Error") //Named Type Error
	}
	//SentinelError与Errors对比
	if ErrStructType == errors.New("EOF") { //两次都是新生成的指针errorString struct,所以为False
		fmt.Println("Struct Type Error")
	}
}
