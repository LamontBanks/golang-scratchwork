package main

import (
	"fmt"
	"time"
)

func main() {
	names := []string{"Scadriel", "Roshar", "Sel", "Nalthis", "Taldain"}

	fmt.Printf("%v", processMessages(names))
}

func processMessages(messages []string) []string {
	processedMsgCh := make(chan string)

	for _, msg := range messages {
		go func() {
			processedMsgCh <- process(msg)
		}()
	}

	processedMsgs := make([]string, len(messages))

	for i := 0; i < len(messages); i++ {
		processedMsgs[i] = <-processedMsgCh
	}

	return processedMsgs
}

// don't touch below this line

func process(message string) string {
	time.Sleep(1 * time.Second)
	return message + "-processed"
}
