package main

import "fmt"

func main() {
	numSends := 21
	ch := make(chan int)

	// Function to send down channel
	go sendInt(ch, numSends)

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

	close(ch)
}
