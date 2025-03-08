// Use channels to run concurrent jobs and colelct results
package main

import "fmt"

func main() {
	jobs := make(chan int)
	results := make(chan int)
	final_results := make(chan int)

	numWorkers := 100
	numJobs := 100
	numApprovers := 100

	// Starter worker pool
	for i := 0; i < numWorkers; i++ {
		go worker(jobs, results)
	}

	// Create approvers
	for i := 0; i < numApprovers; i++ {
		go approver(results, final_results)
	}

	// Create jobs
	// jobs -> results -> final_results
	go func() {
		for j := 0; j < numJobs; j++ {
			jobs <- j
		}
	}()

	// Print final results
	for r := 0; r < numJobs; r++ {
		fmt.Println(<-final_results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for num := range jobs {
		// Dummy job: Compute squares of numbers
		result := num * num
		results <- result
	}
}

func approver(results <-chan int, final_resuls chan<- int) {
	for result := range results {
		final_resuls <- result / 2
	}
}
