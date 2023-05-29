package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func knapsack(weights []int, values []int, capacity int) int {
	n := len(values)

	// Create a 2D array for storing intermediate results
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	// Fill the array in bottom-up fashion
	for i := 0; i <= n; i++ {
		for w := 0; w <= capacity; w++ {
			if i == 0 || w == 0 {
				dp[i][w] = 0
			} else if weights[i-1] <= w {
				dp[i][w] = max(values[i-1]+dp[i-1][w-weights[i-1]], dp[i-1][w])
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	return dp[n][capacity]
}

func main() {
	values := []int{60, 100, 120}
	weights := []int{10, 20, 30}
	capacity := 50
	fmt.Println(knapsack(weights, values, capacity)) // 220
}
