package main

import (
	"fmt"
	"sync"
)

func printEven(input int, C chan bool, wg *sync.WaitGroup) {
	for i := 2; i <= input; i += 2 {
		<-C
		fmt.Println(i)
		C <- true
	}
	wg.Done()

}

func printOdd(input int, C chan bool, wg *sync.WaitGroup) {
	for i := 1; i <= input; i += 2 {
		fmt.Println(i)
		C <- true
		<-C
	}

	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	C := make(chan bool)
	wg.Add(1)
	go printEven(100, C, &wg)
	wg.Add(1)
	go printOdd(100, C, &wg)
	wg.Wait()

}
