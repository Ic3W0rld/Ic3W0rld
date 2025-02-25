package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/trace"
)

func maxAbsoluteSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	minSum := nums[0]
	maxAbsSum := abs(nums[0])

	currentMax := nums[0]
	currentMin := nums[0]

	for i := 1; i < len(nums); i++ {
		// Update currentMax and currentMin
		currentMax = max(nums[i], currentMax+nums[i])
		currentMin = min(nums[i], currentMin+nums[i])

		// Update maxSum and minSum
		maxSum = max(maxSum, currentMax)
		minSum = min(minSum, currentMin)

		// Update maxAbsSum
		maxAbsSum = max(maxAbsSum, abs(maxSum))
		maxAbsSum = max(maxAbsSum, abs(minSum))
	}

	return maxAbsSum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// Start tracing
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("Failed to create trace file: %v", err)
	}
	defer f.Close()

	if err := trace.Start(f); err != nil {
		log.Fatalf("Failed to start trace: %v", err)
	}
	defer trace.Stop()

	nums1 := []int{1, -3, 2, 3, -4}
	fmt.Println(maxAbsoluteSum(nums1)) // Output: 5

	nums2 := []int{2, -5, 1, -4, 3, -2}
	fmt.Println(maxAbsoluteSum(nums2)) // Output: 8
}
