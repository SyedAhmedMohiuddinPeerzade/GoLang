package main

import (
	"fmt"
	"sync"
	"time"
)

var s = 0
var mutex = &sync.Mutex{}

func depositor() {
	for i := 0; i < 1000; i++ {
		mutex.Lock()
		s = s + 1
		mutex.Unlock()
	}

}

func main() {
	go depositor()
	go depositor()
	time.Sleep(5 * time.Second)
	fmt.Println("state", s)
}
