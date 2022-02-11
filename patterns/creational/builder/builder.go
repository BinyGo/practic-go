/*
 * @Auther: BinyGo
 * @Description: 建造者模式
 * @Date: 2022-02-11 10:12:08
 * @LastEditTime: 2022-02-11 19:18:08
 */
package builder

import "fmt"

//建造者模式分离的建设一个复杂对象的表示,这样相同的施工过程可以创建不同的表示。

//电脑对象的构建接口
type Computer interface {
	MakeCpu()
	MakeKeyBoard()
	MakeScreen()
}

//定义一个构建者
type Creator struct {
	Computer Computer
}

//创建构建者方法
func (builder *Creator) Construct() *Computer {
	builder.Computer.MakeCpu()
	builder.Computer.MakeKeyBoard()
	builder.Computer.MakeScreen()
	return &builder.Computer
}

//构建中国制造的电脑
type ChinaComputer struct {
	Cpu      string
	KeyBoard string
	Screen   string
}

//制作Cpu方法
func (c *ChinaComputer) MakeCpu() {
	fmt.Println("主板构建中...")
	c.Cpu = "china cpu"
}

//制作KeyBoard方法
func (c *ChinaComputer) MakeKeyBoard() {
	fmt.Println("键盘构建中...")
	c.KeyBoard = "china keyboard"
}

//制作Screen方法
func (c *ChinaComputer) MakeScreen() {
	fmt.Println("屏幕构建中...")
	c.Screen = "china screen"
}

func Builder() {
	//构建中国制造电脑结构
	c := ChinaComputer{}
	//将中国电脑结构注入构建者中
	b := Creator{Computer: &c}
	//通过构建者,生产中国电脑
	Computer := b.Construct()
	fmt.Printf("%+v", *Computer)
	fmt.Println("")

	//构建cc制造电脑结构
	u := CcComputer{}
	//将cc电脑结构注入构建者中
	b = Creator{Computer: &u}
	//通过构建者,生产cc电脑
	Uk := b.Construct()
	fmt.Printf("%+v", *Uk)
	fmt.Println("")
}

//构建中国制造的电脑
type CcComputer struct {
	Cpu      string
	KeyBoard string
	Screen   string
}

//制作Cpu方法
func (c *CcComputer) MakeCpu() {
	fmt.Println("cc主板构建中...")
	c.Cpu = "cc cpu"
}

//制作KeyBoard方法
func (c *CcComputer) MakeKeyBoard() {
	fmt.Println("cc键盘构建中...")
	c.KeyBoard = "cc keyboard"
}

//制作Screen方法
func (c *CcComputer) MakeScreen() {
	fmt.Println("cc屏幕构建中...")
	c.Screen = "cc screen"
}
