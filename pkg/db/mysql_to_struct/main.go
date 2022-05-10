package main

import (
	"fmt"

	"github.com/gohouse/converter"
)

func main() {
	err := converter.NewTable2Struct().
		SavePath("model/blog.go").
		Dsn("root:root@tcp(localhost:30306)/blog?charset=utf8").
		//Prefix("prefix_"). // 表前缀
		EnableJsonTag(true). // 是否添加json tag
		TagKey("gorm").      // tag字段的key值,默认是gorm
		Run()
	fmt.Println(err)
}
