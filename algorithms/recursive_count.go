package main

import "fmt"

func main() {
	s := []int{7, 2, 9, 12, 8}
	fmt.Printf("Count of %v is %d", s, calc_count(s))
}
func calc_count(list []int) int {
	if len(list) == 0 {
		return 0
	}
	return 1 + calc_count(list[1:])
}
