package main

import "fmt"

type rectangle struct {
	length, breadth float32
}
type circle struct {
	radius float32
}

type shape interface {
	area() float32
}

func (rec rectangle) area() float32 {
	return rec.length * rec.breadth
}

func (cir circle) area() float32 {
	return 3.14 * cir.radius * cir.radius

}


func getArea(s shape) float32 {
	return s.area()
}
func main() {
	rect := rectangle{2, 3}
	circ := circle{5}
	fmt.Println(getArea(rect))
	fmt.Println(getArea(circ))
}
