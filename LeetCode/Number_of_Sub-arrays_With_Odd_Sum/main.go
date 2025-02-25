package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/trace"
	"time"
)

const mod = 1e9 + 7

func numOfSubarrays(arr []int) int {
	odd, even := 0, 1 // even is initialized to 1 for prefix sum of 0
	prefixSum := 0
	result := 0

	for _, num := range arr {
		prefixSum += num
		if prefixSum&1 == 1 { // Check if prefixSum is odd using bitwise AND
			result += even
			odd++
		} else {
			result += odd
			even++
		}
	}

	return result % mod
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

	arr1 := []int{1, 3, 5}
	fmt.Println(numOfSubarrays(arr1))

	arr2 := []int{2, 4, 6}
	fmt.Println(numOfSubarrays(arr2))

	time.Sleep(5 * time.Second)
}
