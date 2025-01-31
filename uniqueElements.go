package main

import (
	"fmt"
)

var myMap map[int]struct{}

func main() {

	myMap = make(map[int]struct{})
	mySlice := []int{1, 4, -1, -3, 2, 3, 2, -4, 4, -1, -2, -3, -4, 0, 1, -2, 3}
	for _, value := range mySlice {
		_, ok := myMap[value]
		if ok == false {
			myMap[value] = struct{}{}
			fmt.Println(value)
		}
	}
}
