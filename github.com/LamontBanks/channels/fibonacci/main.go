package main

import "fmt"

func main() {
	fmt.Printf("%v", concurrentFib(24))
}

func concurrentFib(n int) []int {
	fibCh := make(chan int)

	go fibonacci(n, fibCh)

	fibSeq := []int{}
	for v := range fibCh {
		fibSeq = append(fibSeq, v)
	}

	return fibSeq
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
