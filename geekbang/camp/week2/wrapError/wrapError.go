package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pkg/errors"
)

func WriteAll(w io.Writer, buf []byte) error {
	_, err := w.Write(buf)
	if err != nil {
		//经常有需要打印错误信息的场景
		log.Println("unable to write:", err)
		return err
	}
	return nil
}

/*
//伪代码
//不包装起来,顶层打印出来的就是一个io.EOF
func main() {
	err := WriteAll(f, []byte{1})
	fmt.Println(err) //io.EOF
}
*/

// 使用github.com/pkg/errors库
// 一般用法:
//1.在自己的应用代码中,生成error,使用errors.New()或errors.Errorf()
/*
if len(args)<3{
	return errors.New("错误信息...")
}
*/
//2.在应用内调用其他包内的函数,通常简单的直接返回error,直接往上抛
/*
if err != nil{
	return err
}
*/
//3.和其他库包括标准库进行协作,考虑用errors.Wrap() errors.Wrapf()保存堆栈信息
/*
f,err:=os.Open(path)
if err != nil {
	return errors.Wrap(err,"failed to open %q",path)
}
*/
//4.最终,直接返回错误,而不是每个错误产生的地方到处打日志
//在程序顶部或者工作的goroutine顶部(请求入口),使用fmt.Printf("stack trace:\n %+v \n", err)输出或记录堆栈详情信息

func WrapWriteAll(w io.Writer, buf []byte) error {
	_, err := w.Write(buf)
	if err != nil {
		//使用errors.Wrap,保留原始err,不会破坏原始error内容,stack,加入错误信息"unable to write"
		err = errors.Wrap(err, "unable to write")
		//err = fmt.Errorf("unable to write %w",err)
	} else {
		//这里为强加返回错误代码以查看效果
		err = errors.New("diy error")
		//err = errors.Errorf("diy error") //多参数errors.New
		err = errors.Wrap(err, "wrap error")
	}
	//errors.Wrap(err, "追加信息") 不能多次调用,会重复保存堆栈stack信息
	err = errors.WithMessage(err, "这里是追加信息")
	return err
}

func main() {
	buf := bytes.NewBufferString("sss")
	err := WrapWriteAll(buf, []byte{1})
	if err != nil {
		fmt.Printf("original error: %T,v: %v\n", errors.Cause(err), errors.Cause(err)) //根因,这里输出未被Wrap的err信息
		fmt.Printf("stack trace:\n %+v \n", err)                                       //这里输出Wrap信息
		os.Exit(1)
	}
}

/*
$ go run wrapError.go
original error: *errors.fundamental,v: diy error //根因
stack trace: //包因
 diy error
main.WrapWriteAll
        /golang/src/github.com/practic-go/geekbang/camp/week2/wrapError/wrapError.go:40
main.main
        /golang/src/github.com/practic-go/geekbang/camp/week2/wrapError/wrapError.go:50
runtime.main
        /go/go1.17.6/src/runtime/proc.go:255
runtime.goexit
        /go/go1.17.6/src/runtime/asm_amd64.s:1581
wrap error
main.WrapWriteAll
        /golang/src/github.com/practic-go/geekbang/camp/week2/wrapError/wrapError.go:41
main.main
        /golang/src/github.com/practic-go/geekbang/camp/week2/wrapError/wrapError.go:50
runtime.main
        /go/go1.17.6/src/runtime/proc.go:255
runtime.goexit
        /go/go1.17.6/src/runtime/asm_amd64.s:1581
这里是追加信息
exit status 1
*/

/*

Wrap errors
使用 errors.Cause 获取 root error，再和 sentinel error 判定。

总结:
Packages that are reusable across many projects only return root error values.
选择 wrap error 是只有 applications 可以选择应用的策略，具有最高可重用性的包只能返回根错误值。此机制与 Go 标准库中使用的相同（kit 库的 sql.ErrNoRows）。
biny:简单说就是只在你的"业务应用"中使用,而不是在地方三包中使用,不在可供他人使用的包的使用,例如官方的kit 库的 sql.ErrNoRows等,就都是直接返回根因错误值

If the error is not going to be handled, wrap and return up the call stack.
这是关于函数/方法调用返回每个错误的基本问题。如果函数/方法不打算处理错误，那么用足够的上下文 wrap errors 并将其返回到调用堆栈中。
例如，额外的上下文可以是使用的输入参数或失败的查询语句。确定你记录的上下文是足够多还是太多的一个好方法是检查日志并验证它们在开发期间是否为你工作。
biny:如果错误不打算被处理，就包起来，并向上返回调用栈。

Once an error is handled, it is not allowed to be passed up the call stack any longer.
一旦确定函数/方法将处理错误，错误就不再是错误。如果函数/方法仍然需要发出返回，则它不能返回错误值。它应该只返回零（比如降级处理中，你返回了降级数据，然后需要 return nil）。
biny:一旦一个错误被处理了，就不允许再往上传递调用栈。
*/
