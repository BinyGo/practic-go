package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"syscall"
)

//1.1 直接指定文件名读取
//第一种：使用 os.ReadFile
func Read1() {
	content, err := os.ReadFile("a.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

//第二种：使用 ioutil.ReadFile  在 Go 1.16 开始，ioutil.ReadFile 就等价于 os.ReadFile，二者是完全一致的
func Read2() {
	content, err := ioutil.ReadFile("a.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

//1.2 先创建句柄再读取
//如果仅是读取，可以使用高级函数 os.Open

func Read3() {
	file, err := os.Open("a.txt") //os.Open 是只读模式的 os.OpenFile
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

//因此，你也可以直接使用 os.OpenFile，只是要多加两个参数
func Read4() {
	file, err := os.OpenFile("a.txt", os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

//2. 每次只读取一行
// 一次性读取所有的数据，太耗费内存，因此可以指定每次只读取一行数据。方法有三种：
// bufio.ReadLine()
// bufio.ReadBytes('\n')
// bufio.ReadString('\n')
// 在 bufio 的源码注释中，曾说道 bufio.ReadLine() 是低级库，不太适合普通用户使用，更推荐用户使用 bufio.ReadBytes 和 bufio.ReadString 去读取单行数据。
// 因此，这里不再介绍 bufio.ReadLine()
func Read5() {
	// 创建句柄
	file, err := os.Open("a.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//创建Reader
	r := bufio.NewReader(file)
	for {
		lineBytes, err := r.ReadBytes('\n')
		line := strings.TrimSpace(string(lineBytes))
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		fmt.Println(line)
	}
}

// 2.2 使用 bufio.ReadString
func Read6() {
	file, err := os.Open("a.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//创建Reader
	r := bufio.NewReader(file)
	for {
		line, err := r.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		fmt.Println(line)
	}
}

// # 3. 每次只读取固定字节数
// 每次仅读取一行数据，可以解决内存占用过大的问题，但要注意的是，并不是所有的文件都有换行符 \n。
// 因此对于一些不换行的大文件来说，还得再想想其他办法。

// 3.1 使用 os 库
// 通用的做法是：
// 先创建一个文件句柄，可以使用 os.Open 或者  os.OpenFile
// 然后 bufio.NewReader 创建一个 Reader
// 然后在 for 循环里调用  Reader 的 Read 函数，每次仅读取固定字节数量的数据。

func Read7() {
	file, err := os.Open("a.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r := bufio.NewReader(file)
	//每次读1024字节
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		fmt.Println(string(buf[:n]))
	}
}

// 3.2 使用 syscall 库
// os 库本质上也是调用 syscall 库，但由于 syscall 过于底层，如非特殊需要，一般不会使用 syscall
// 本篇为了内容的完整度，这里也使用 syscall 来举个例子。
// 本例中，会每次读取  100 字节的数据，并发送到通道中，由另外一个协程进行读取并打印出来。

func Read8() {
	fd, err := syscall.Open("a.txt", syscall.O_RDONLY, 0)
	if err != nil {
		fmt.Println("Failed on open: ", err)
	}
	defer syscall.Close(fd)

	var wg sync.WaitGroup
	wg.Add(2)
	dataChan := make(chan []byte)
	go func() {
		wg.Done()
		for {
			data := make([]byte, 100)
			n, _ := syscall.Read(fd, data)
			if n == 0 {
				break
			}
			dataChan <- data
		}
		close(dataChan)
	}()

	go func() {
		defer wg.Done()
		data := <-dataChan
		fmt.Println(string(data))
	}()
	wg.Wait()
}
