package main

import "fmt"

// Error Types 自定义的错误类型
// Error Type 实现了 error 接口的自定义类型。
// 因为 MyError 是一个 type，调用者可以使用断言转换成这个类型，来获取更多的上下文信息。

// 与错误值相比，错误类型的一大改进是它们能够包装底层错误以提供更多上下文。
// 一个不错的例子就是 os.PathError ，它提供了底层执行了什么操作、哪个路径出了什么问题等信息

// 调用者要使用类型断言和类型 switch，就要让自定义的 error 变为 public。这种模型会
// 导致与调用者产生强耦合，从而导致 API 变得脆弱。
// 结论是尽量避免使用 error types，虽然错误类型比 sentinel errors 更好，因为它们可
// 以捕获关于出错的更多上下文，但是 error types 共享 error values 许多相同的问题。

// 建议是避免使用错误类型，或者至少避免将它们作为公共 API 的一部分。

type MyError struct {
	Msg  string
	File string
	Line int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s:%d: %s", e.File, e.Line, e.Msg)
}

func test() error {
	return &MyError{"Something happend", "server.go", 42}
}

func main() {
	err := test()
	switch err := err.(type) {
	case nil:
		//call succeeded, nothing to do
	case *MyError:
		fmt.Println("error occurred on line:", err.Line)
	default:
		//unKnow error
	}
}
