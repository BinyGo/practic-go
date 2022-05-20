package main

// Go下的linux进程,NameSpace,Cgroups剖析,该示例需要搭配图文查看为佳
// 更多信息:https://shimo.im/docs/m4kMLR1rZBCXgEqD

import (
	"os"
	"os/exec"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("help")
	}
	// 这个程序接收用户命令行传递的参数，并使用 exec.Command 运行，
	// 例如当我们执行 go run main.go run echo hello 时，
	// 会创建出 main 进程， main 进程内执行 echo hello 命令创建出一个新的 echo 进程，
	// 最后随着 echo 进程的执行完毕，main 进程也随之结束并退出。
}

func run() {
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
