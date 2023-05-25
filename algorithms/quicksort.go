package main

import "fmt"

func main() {
	s := []int{7, 2, 9, 12, 8}
	fmt.Printf("Sorted %v is %v", s, quicksort(s))
}
func quicksort(array []int) []int {
	var less []int
	var greater []int
	var combined []int
	// Base case.
	if len(array) < 2 {
		return array
	}
	// Recursive case.
	pivot := array[0]
	for _, v := range array[1:] {
		if v <= pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}

	combined = append(combined, quicksort(less)...)
	combined = append(combined, pivot)
	combined = append(combined, quicksort(greater)...)
	return combined

}
