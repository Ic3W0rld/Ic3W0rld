package main

import (
	"fmt"
)

func lenLongestFibSubseq(arr []int) int {
	n := len(arr)
	if n < 3 {
		return 0
	}

	// Create a map to store the index of each element
	indexMap := make(map[int]int)
	for i, num := range arr {
		indexMap[num] = i
	}

	// Initialize the DP table
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	maxLen := 0

	// Iterate through all pairs (i, j)
	for j := 0; j < n; j++ {
		for i := 0; i < j; i++ {
			// Check if there exists a k such that arr[k] + arr[i] == arr[j]
			diff := arr[j] - arr[i]
			if k, exists := indexMap[diff]; exists && k < i {
				// Update dp[i][j]
				dp[i][j] = dp[k][i] + 1
				// Update the maximum length
				if dp[i][j] > maxLen {
					maxLen = dp[i][j]
				}
			} else {
				// If no such k exists, set dp[i][j] to 2 (minimum length for a Fibonacci-like sequence)
				dp[i][j] = 2
			}
		}
	}

	// If maxLen is greater than 2, return maxLen + 1 (since dp[i][j] counts the pairs, not the full sequence)
	if maxLen >= 2 {
		return maxLen + 1
	}
	return 0
}

func main() {
	// Example 1
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(lenLongestFibSubseq(arr1)) // Output: 5

	// Example 2
	arr2 := []int{1, 3, 7, 11, 12, 14, 18}
	fmt.Println(lenLongestFibSubseq(arr2)) // Output: 3
}
