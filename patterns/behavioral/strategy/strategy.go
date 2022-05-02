package strategy

import "fmt"

// 1. Strategy Pattern
// Strategy behavioral design pattern enables an algorithm's behavior to be selected at runtime.

// It defines algorithms, encapsulates them, and uses them interchangeably.

// 1.1. Implementation
// Implementation of an interchangeable operator object that operates on integers.
type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator Operator
}

func (o *Operation) Operate(leftValue, rightValue int) int {
	return o.Operator.Apply(leftValue, rightValue)
}

// 1.2. Usage
// 1.2.1. Addition Operator
type Addition struct{}

func (Addition) Apply(leftValue, rightValue int) int {
	return leftValue + rightValue
}

func Demo1() {
	add := Operation{Addition{}}
	fmt.Println(add.Operate(3, 5)) //8
}

//1.2.2. Multiplication Operator
type Multiplication struct{}

func (Multiplication) Apply(leftValue, rightValue int) int {
	return leftValue * rightValue
}

func Demo2() {
	mult := Operation{Multiplication{}}
	fmt.Println(mult.Operate(3, 5)) //15
}

// 1.3. Rules of Thumb
// Strategy pattern is similar to Template pattern except in its granularity.
// Strategy pattern lets you change the guts of an object. Decorator pattern lets you change the skin.
