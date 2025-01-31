package main

import (
	"fmt"
	"sync"
)

func printEven(input int, C chan bool) {
	for i := 2; i <= input; i += 2 {
		<-C
		fmt.Println(i)
		C <- true
	}

}

func printOdd(input int, C chan bool) {
	for i := 1; i <= input; i += 2 {
		fmt.Println(i)
		C <- true
		<-C
	}

}

func main() {
	var wg sync.WaitGroup
	C := make(chan bool)
	wg.Add(1)
	go func() {
		printEven(100, C)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		printOdd(100, C)
		wg.Done()
	}()
	wg.Wait()

}
