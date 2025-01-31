package main

import "fmt"

type myInterface interface {
	area()
}

type triangle struct {
	base   float64
	height float64
}

type rectangle struct {
	length  int
	breadth int
}

func (T triangle) area() {
	fmt.Println(0.5 * T.base * T.height)
}

func (R rectangle) area() {
	fmt.Println(R.length * R.breadth)
}

func main() {
	var I myInterface
	T := triangle{base: 10.0, height: 20.0}
	R := rectangle{length: 5, breadth: 10}
	I = R
	I.area()
	I = T
	I.area()

}
