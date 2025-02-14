package main

import (
	"fmt"
	"time"
)

type Ball struct {
	hits int
}

func player(name string, table chan *Ball) {
	ball := <-table
	ball.hits++
	fmt.Println(name, "...", ball.hits)
	time.Sleep(time.Second)
	table <- ball
}

func main() {
	table := make(chan *Ball)
	ball := new(Ball)
	go player("p1", table)
	go player("p2", table)
	time.Sleep(2 * time.Second)
	table <- ball
	time.Sleep(10 * time.Second)
	fmt.Println(ball.hits)

}
