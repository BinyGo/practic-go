package day079

import "fmt"

type ConfigOne struct {
	Daemon string
}

func (c *ConfigOne) String() string {
	return fmt.Sprintf("print:%v", c)
}

func Demo1() {
	c := &ConfigOne{}
	c.String() //fmt.Sprintf format %v with arg c causes recursive (*github.com/practic-go/plan105/day079.ConfigOne).String method call
	// 参考答案及解析：无限递归循环，栈溢出。知识点：类型的 String() 方法。
	// 如果类型定义了 String() 方法，使用 Printf()、Print() 、 Println() 、 Sprintf() 等格式化输出时会自动使用 String() 方法。
}
