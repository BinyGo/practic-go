package main

// Opaque errors 不透明错误处理

// 在我看来，这是最灵活的错误处理策略，因为它要求代码和调用者之间的耦合最少。
// 我将这种风格称为不透明错误处理，因为虽然你知道发生了错误，但你没有能力看到错误的内部。
// 作为调用者，关于操作的结果，你所知道的就是它起作用了，或者没有起作用（成功还是失败）。
// 这就是不透明错误处理的全部功能–只需返回错误而不假设其内容

//官方net包Error
type Error interface {
	error
	Timeout() bool
	Temporary() bool
}

type temporary interface {
	Temporary() bool
}

//IsTemporary returns true if err is temporary
func IsTemporary(err error) bool {
	te, ok := err.(temporary) //先断言成temporary类型,来获得上下文信息,再通过内部行为如Temporary()进行下一步处理
	return ok && te.Temporary()
	// 这里的关键是，这个逻辑可以在不导入定义错误的包或者实际上不了解 err 的底层类型的情况下实现——我们只对它的行为感兴趣
	// 这错误处理方式还是有许多麻烦,无返回错误信息
}
