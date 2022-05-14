package main

import (
	"embed"
	"fmt"
)

//go:embed hello.txt
var version string

//go:embed hello.txt
var versionByte []byte

//go:embed hello.txt
//go:embed hello2.txt
var embedFiles embed.FS

func main() {

	data, _ := embedFiles.ReadFile("hello.txt")
	println("embedFiles hello.txt:", string(data))

	data, _ = embedFiles.ReadFile("hello2.txt")
	println("embedFiles hello2.txt:", string(data))

	fmt.Printf("version: %q\n", version)
	fmt.Printf("versionByte %q\n", string(versionByte))

}
