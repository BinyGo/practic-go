package main

import (
	"bufio"
	"fmt"
	"os"
)

//os
func Demo1() {
	filename := "./test.txt"
	//以读写模式打开文件
	fd, _ := os.OpenFile(filename, os.O_RDWR, 0666)
	b := make([]byte, 12)
	// 从文件中读取最多len(b)个字节的内容到b中
	n, err := fd.Read(b)
	fmt.Printf("n:=%d, b:=%s, err=%+v\n", n, b, err)
	//n:=10, b:=bufio包�, err=<nil>

	//上面的读取方式是通过文件系统的IO进行读取的，每次都需要一次底层的系统调用，若需要连续多次读取，那么这种方式的效率就会大大降低。
}

//bufio
func Demo2() {
	filename := "./test.txt"
	//以读写模式打开文件
	fd, _ := os.OpenFile(filename, os.O_RDWR, 0666)

	//将fd包装到buffer Reader中
	bufioReader := bufio.NewReader(fd)

	p := make([]byte, 12)
	n, _ := bufioReader.Read(p)

	fmt.Printf("n=%d,p=%s", n, p)
	// n=10,p=bufio包�PASS
}

// 更多细节:https://mp.weixin.qq.com/s/nDTzCqPwb5nriHsuhQjVpA
