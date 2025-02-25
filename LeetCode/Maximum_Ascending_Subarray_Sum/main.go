package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/trace"
	"time"
)

func maxAscendingSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	current := nums[0]
	maxSum := current
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			current += nums[i]
		} else {
			current = nums[i]
		}
		if current > maxSum {
			maxSum = current
		}
	}
	return maxSum
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

	arr1 := []int{10, 20, 30, 5, 10, 50}
	fmt.Println(maxAscendingSum(arr1))

	arr2 := []int{10, 20, 30, 40, 50}
	fmt.Println(maxAscendingSum(arr2))

	arr3 := []int{12, 17, 15, 13, 10, 11, 12}
	fmt.Println(maxAscendingSum(arr3))

	time.Sleep(5 * time.Second)
}
