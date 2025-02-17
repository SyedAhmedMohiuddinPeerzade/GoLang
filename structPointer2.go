package main

import (
	"fmt"
)

// In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both(Because of interfaces)

type myInterface interface {
	Add()
}

type A struct {
	a int
}

type B struct {
	b int
}

func (value A) Add() { // tpe A implements interface myInterface
	value.a = value.a + 10
	fmt.Println(value.a)
}

func (value *B) Add() { // tpe B implements interface myInterface
	value.b = value.b + 10
	fmt.Println(value.b)

}

func main() {
	var x myInterface

	a1 := A{a: 10}
	x = a1 // Works
	x.Add()

	b1 := B{b: 20}
	b1.Add() // Works

	//	x = b1 // ERROR if we assign value to interface instead pointer, as interface method is implemented with a pointer receiver

	b2 := B{b: 30}
	x = &b2 // Works
	x.Add()

}
