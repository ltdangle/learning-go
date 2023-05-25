package main

import "fmt"

func main() {
	var s []int
	s = []int{7, 2, 9, 12, 8}
	fmt.Printf("Sum of %v is %d", s, calc_max(s))
}
func calc_max(list []int) int {
	// Base case.
	if len(list) == 2 {
		if list[0] > list[1] {
			return list[0]
		} else {
			return list[1]
		}
	}
	// Recursive case.
	sub_max := calc_max(list[1:])
	if list[0] > sub_max {
		return list[0]
	} else {
		return sub_max
	}
}
