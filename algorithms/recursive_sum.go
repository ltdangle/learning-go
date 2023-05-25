package main

import (
	"fmt"
)

func main() {
	var s []int
	s = []int{1, 2, 3, 4, 5}
	fmt.Printf("Sum of %v is %d", s, calc_sum(s))

}
func calc_sum(arr []int) int {
	var sum int
	// Base case.
	if len(arr) == 1 {
		sum += arr[0]
		return sum
	}
	// Recursive case.
	sum += arr[0] + calc_sum(arr[1:])
	return sum

}
