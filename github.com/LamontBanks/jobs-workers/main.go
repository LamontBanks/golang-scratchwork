// Use channels to run concurrent jobs and colelct results
package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int)
	results := make(chan int)

	// Starter worker pool
	numWorkers := 10
	for i := 0; i < numWorkers; i++ {
		go worker(jobs, results)
	}

	// Create "jobs"
	numJobs := 100
	go func() {
		for numToSquare := 0; numToSquare < numJobs; numToSquare++ {
			jobs <- numToSquare
		}

		// Notify workers that there are no more jobs
		close(jobs)
	}()

	// Get results (in any order)
	jobResults := make([]int, numJobs)
	for i := 0; i < numJobs; i++ {
		jobResults[i] = <-results
	}

	// Print results
	fmt.Println("Results: ", jobResults)
}

func worker(jobs <-chan int, results chan<- int) {
	for num := range jobs {
		// Dummy job: Compute squares of numbers
		result := num * num

		// Fake job time
		time.Sleep(250 * time.Millisecond)

		results <- result
	}
}
