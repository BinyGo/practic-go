package flag

import (
	flag2 "flag"
	"fmt"
	"time"
)

var (
	flagvar int
	name    string
	age     int
	married bool
	delay   time.Duration
)

//flag包实现了命令行参数的解析，flag包使得开发命令行工具更为简单
//go run main.go biny 18 true 20
//go run main.go -flagvar 666
//通过flag2.Args()即可取得传入参数
func init() {
	flag2.StringVar(&name, "name", "张三", "姓名")
	flag2.IntVar(&age, "age", 18, "年龄")
	flag2.BoolVar(&married, "married", false, "婚否")
	flag2.DurationVar(&delay, "d", 0, "延迟的时间间隔")
	flag2.IntVar(&flagvar, "flagname", 1234, "help message for flagname") //设置默认参数
}

func Demo1() {
	fmt.Println(name, age, married, delay, flagvar)
	//解析命令行参数
	flag2.Parse()
	fmt.Println(name, age, married, delay)
	//返回命令行参数后的其他参数
	fmt.Println("Args:", flag2.Args())
	//更多查看官方文档 https://pkg.go.dev/flag
}
