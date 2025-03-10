package main

import "fmt"

func main() {
	// deadLocks()

	noDeadlock()
}

func deadLocks() {
	ch := make(chan int)
	ch <- 1
	fmt.Println(<-ch)
}

func noDeadlock() {
	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	fmt.Println(<-ch)
}
