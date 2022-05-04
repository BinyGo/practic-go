package main

import "fmt"

func main() {
	ChooseFruit("蓝莓")
	ChooseFruit("苹果")
	ChooseFruit("西瓜")
}

func ChooseFruit(fruit string) {
	switch fruit {
	case "苹果":
		fmt.Println("这是一个苹果")
	case "草莓", "蓝莓":
		fmt.Println("这是一个草莓或蓝莓")
	default:
		fmt.Printf("识别不出来: %s \n", fruit)
	}
}
