package main

import (
	"errors"
	"fmt"
)

func main() {
	var err error = &MyError{}
	println(err.Error()) //hello,it's my error

	ErrorsPkg()

	defer func() {
		if data := recover(); data != nil {
			fmt.Printf("\nhello,panic: %v \n", data)
		}
		fmt.Println("恢复之后这里继续执行")
	}()

	panic("Boom")
	//fmt.Println("这里将不会被执行到")

}

type MyError struct{}

func (m *MyError) Error() string {
	return "hello,it's my error"
}

func ErrorsPkg() {
	err := &MyError{}
	// 使用 %w 占位符，返回的是一个新错误
	// wrappedErr 是一个新类型，fmt.wrapError
	wrappedErr := fmt.Errorf("this is an wrapped error %w", err)

	fmt.Println(wrappedErr) //this is an wrapped error hello,it's my error
	fmt.Println(err)        //hello,it's my error
	// 再解出来
	if err == errors.Unwrap(wrappedErr) {
		fmt.Println("unwrapped") //unwrapped
	}

	if errors.Is(wrappedErr, err) {
		// 虽然被包了一下，但是 Is 会逐层解除包装，判断是不是该错误
		fmt.Println("wrapped is err") // wrapped is err
	}

	copyErr := &MyError{}
	// 这里尝试将 wrappedErr转换为 MyError
	// 注意我们使用了两次的取地址符号
	if errors.As(wrappedErr, &copyErr) {
		fmt.Println("convert error") // convert error
	}
}
