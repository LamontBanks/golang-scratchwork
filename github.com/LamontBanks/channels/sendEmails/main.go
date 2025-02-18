package main

import (
	"fmt"
	"time"
)

type email struct {
	message string
	date    time.Time
}

func main() {
	emails := [3]email{
		{
			message: "Words are pale shadows of forgotten names. As names have power, words have power.",
			date:    time.Date(2019, 2, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			message: "It's like everyone tells a story about themselves inside their own head.",
			date:    time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			message: "Bones mend. Regret stays with you forever.",
			date:    time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
		},
	}

	fmt.Printf("%v", checkEmailAge(emails))
}

func checkEmailAge(emails [3]email) [3]bool {
	isOldChan := make(chan bool)

	go sendIsOld(isOldChan, emails)

	isOldResults := [3]bool{}
	isOldResults[0] = <-isOldChan
	isOldResults[1] = <-isOldChan
	isOldResults[2] = <-isOldChan

	return isOldResults
}

func sendIsOld(isOldChan chan<- bool, emails [3]email) {
	for _, e := range emails {
		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
			isOldChan <- true
			continue
		} else {
			isOldChan <- false
		}
	}
}
