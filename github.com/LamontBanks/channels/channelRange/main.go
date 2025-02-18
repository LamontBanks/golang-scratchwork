package main

import "fmt"

func main() {
	numSends := 21
	ch := make(chan int)

	// Function to send down channel
	// go sendInt(ch, numSends)

	// Anonymous function to send down channel
	go func() {
		for i := 0; i < numSends; i++ {
			ch <- i
		}
		close(ch) // Range over channel requires channel to be manually closed
	}()

	receivedNum := []int{}
	for receivedInt := range ch {
		receivedNum = append(receivedNum, receivedInt)
	}

	fmt.Printf("%v", receivedNum)
}

func sendInt(ch chan<- int, numSends int) {
	for i := 0; i < numSends; i++ {
		ch <- i
	}

	// close(ch)
}
