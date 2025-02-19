// Use channels to run concurrent jobs and colelct results
package main

import "fmt"

func main() {
	jobs := make(chan int)
	results := make(chan int)

	// Starter worker pool
	numWorkers := 3
	for i := 0; i < numWorkers; i++ {
		go worker(jobs, results)
	}

	// Create "jobs"
	numJobs := 20
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
		results <- result
	}
}
