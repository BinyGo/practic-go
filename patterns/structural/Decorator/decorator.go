package Decorator

import (
	"log"
)

// 1.1. Implementation
// LogDecorate decorates a function with the signature func(int) int that manipulates integers and adds input/output logging capabilities.

type Object func(int) int

func LogDecorate(fn Object) Object {
	return func(i int) int {
		log.Println("Starting the execution with the integer", i)
		result := fn(i)
		log.Println("Execution is completed with the result", result)
		return result
	}
}

//1.2 Usage
func Double(i int) int {
	return i * 2
}

func Demo1() {
	f := LogDecorate(Double)
	f(5) //10
}
