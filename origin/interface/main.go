package main

import "fmt"

type Error interface {
	Error() string
}

type RPCError struct {
	Code    int64
	Message string
}

func (e *RPCError) Error() string {
	return fmt.Sprintf("%s,code=%d", e.Message, e.Code)
}

// 上面的代码，并没有Error接口的影子，我们只需要实现Error() string方法就实现了Error接口。在Go中，实现接口的所有方法就隐式地实现了接口。
// 我们使用上述 RPCError 结构体时并不关心它实现了哪些接口，Go 语言只会在传递参数、返回参数以及变量赋值时才会对某个类型是否实现接口进行检查。

// Go语言的这种写法很方便，不用引入包依赖。但是interface底层实现的时候会动态检测也会引入一些问题：
// 1.性能下降。使用interface作为函数参数，runtime 的时候会动态的确定行为。使用具体类型则会在编译期就确定类型。
// 2.不能清楚的看出struct实现了哪些接口，需要借助ide或其它工具。

//更多详情:https://mp.weixin.qq.com/s?__biz=MzkyNzI1NzM5NQ==&mid=2247484775&idx=1&sn=156ebaa2acfa31c9efb316569307d627&chksm=c22b833bf55c0a2dad59c6b4519a1dcf45fab5abc53f6b0c28758d06f1168b3660011c09a803&scene=178&cur_album_id=1932319304830517254#rd

// 空interface runtime.eface
func PrintStr(str []interface{}) {
	for _, val := range str {
		fmt.Println(val)
	}
}

func Demo1() {
	//names := []string{"biny", "go", "php"} //cannot use names (variable of type []string) as []interface{} value in argument to printStr
	//PrintStr(names)
	//[]interface != []string

	intSlice := []int{1, 2, 3}
	strSlice := []string{"biny", "go", "php"}
	interfaceSlice := make([]interface{}, len(strSlice))
	for i, v := range strSlice {
		interfaceSlice[i] = v
		//intSlice[i] = v //cannot use v (variable of type string) as int value in assignment
	}
	fmt.Printf("%#v \n", intSlice)       //[]int{1, 2, 3}
	fmt.Printf("%#v \n", strSlice)       //[]string{"biny", "go", "php"}
	fmt.Printf("%#v \n", interfaceSlice) //[]interface {}{"biny", "go", "php"}
}

// 非空interface runtime.iface
type Person interface {
	GetAge() int
	SetAget(int)
}

type Man struct {
	Name string
	Age  int
}

func (s Man) GetAge() int {
	return s.Age
}

func (s *Man) SetAget(age int) {
	s.Age = age
}

func f(p Person) {
	p.SetAget(10)
	fmt.Println(p.GetAge()) //10
}

func Demo2() {
	p := Man{}
	//f(p) //cannot use p (variable of type Man) as Person value in argument to f: missing method SetAget (SetAget has pointer receiver)
	f(&p)
}

type Animal interface {
	Walk()
	Eat()
}

type Dog struct {
	Name string
}

func (d *Dog) Walk() {
	fmt.Println("go")
}

func (d *Dog) Eat() {
	fmt.Println("eat shit")
}

func Demo3() {
	var d Animal = &Dog{"wangcai"}
	//var c Animal = Dog{"wangcai"} //cannot use (Dog literal) (value of type Dog) as Animal value in variable declaration: missing method Eat (Eat has pointer receiver)
	d.Eat()  //eat shit
	d.Walk() //go
}

// 类型断言
//value, ok := em.(T) 安全类型断言
// <目标类型的值>，<布尔参数> := <表达式>.( 目标类型 )

//value := em.(T)  非安全类型断言
// <目标类型的值> := <表达式>.( 目标类型 )

type Cat struct {
	Name string
}

func Demo4() {
	var c interface{} = new(Cat)
	c1, ok := c.(Cat)
	if !ok {
		fmt.Println("no cat")
	} else {
		fmt.Println(c1)
	}
	//no cat

	//var cc interface{} = new(Cat)
	//c2 := cc.(Cat)
	//fmt.Println(c2) //panic: interface conversion: interface {} is *main.Cat, not main.Cat

	var cc interface{} = Cat{Name: "cc"}
	c2 := cc.(Cat)
	fmt.Println(c2) //{cc}

}

//问题一: 下面代码，哪一行存在编译错误？（多选）

type Student struct {
}

func Set(x interface{}) {
}

func Get(x *interface{}) {
}

func Demo5() {
	s := Student{}
	p := &s
	// A B C D
	Set(s)
	//Get(s) //cannot use s (variable of type Student) as *interface{} value in argument to Get
	Set(p)
	//Get(p) //cannot use p (variable of type *Student) as *interface{} value in argument to Get
	//答案：B、D；解析：我们上文提到过，interface是所有go类型的父类，所以Get方法只能接口*interface{}类型的参数，其他任何类型都不可以。
}

//问题二: 这段代码的运行结果是什么？
func PrintInterface(val interface{}) {
	if val == nil {
		fmt.Println("this is empty interface")
		return
	}
	fmt.Println("this is non-empty interface")
}
func Demo6() {
	var pointer *string = nil
	PrintInterface(pointer)
}

//答案：this is non-empty interface。解析：这里的interface{}是空接口类型，他的结构如下:
// type eface struct { // 16 字节
// 	_type *_type
// 	data  unsafe.Pointer
// }
