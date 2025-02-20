package main

import "fmt"

func main() {
	strArr := []string{
		"lamont",
		"banks",
		"generics",
		"golang",
	}

	intArr := []int{
		3, 8, 7, 0,
	}

	fmt.Println(splitSlice(strArr))
	fmt.Println(splitSlice(intArr))
}

func splitSlice[T any](s []T) ([]T, []T) {
	if len(s) <= 1 {
		return s, nil
	}

	halfway := (len(s) / 2)
	return s[:halfway], s[halfway:]
}
