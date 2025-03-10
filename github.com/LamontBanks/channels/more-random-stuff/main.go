package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := range 50 {
			ch <- i * i
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}

	fmt.Println(factorial(10))

	fmt.Println(mergeSort([]int{8, 4, 2, 7, 10}))
}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func mergeSort(nums []int) []int {
	// List of length 1 is already sorted
	if len(nums) == 1 {
		return nums
	}

	// Otherwise, split the slice in half
	mid := len(nums) / 2
	left_half := nums[:mid]
	right_half := nums[mid:]

	sorted_left_half := mergeSort(left_half)
	sorted_right_half := mergeSort(right_half)

	return merge(sorted_left_half, sorted_right_half)
}

func merge(leftHalf []int, rightHalf []int) []int {
	sortedArray := []int{}

	leftPtr := 0
	rightPtr := 0

	// Compare each element in both array, moving the pointers as needed
	for leftPtr < len(leftHalf) && rightPtr < len(rightHalf) {
		if leftHalf[leftPtr] <= rightHalf[rightPtr] {
			sortedArray = append(sortedArray, leftHalf[leftPtr])
			leftPtr++
		} else {
			sortedArray = append(sortedArray, rightHalf[rightPtr])
			rightPtr++
		}
	}

	// Copy any left over number from the left and right
	if leftPtr < len(leftHalf) {
		sortedArray = append(sortedArray, leftHalf[leftPtr:]...)
	}
	if rightPtr < len(rightHalf) {
		sortedArray = append(sortedArray, rightHalf[rightPtr:]...)
	}

	return sortedArray
}
