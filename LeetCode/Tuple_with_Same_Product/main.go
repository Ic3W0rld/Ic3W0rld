package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/trace"
	"time"
)

func tupleSameProduct(nums []int) int {
	productCount := make(map[int]int, len(nums))
	n := len(nums)

	// Step 1: Generate all possible pairs and count their products
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			product := nums[i] * nums[j]
			productCount[product]++
		}
	}

	// Step 2: Calculate the number of valid tuples
	result := 0
	for _, count := range productCount {
		if count >= 2 {
			// For each product, the number of valid tuples is count * (count - 1) * 4
			// because for each pair of pairs (a,b) and (c,d), there are 4 permutations:
			// (a,b,c,d), (a,b,d,c), (b,a,c,d), (b,a,d,c)
			result += count * (count - 1) * 4
		}
	}

	return result
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

	arr1 := []int{2, 3, 4, 6}
	fmt.Println(tupleSameProduct(arr1))

	arr2 := []int{1, 2, 4, 5, 10}
	fmt.Println(tupleSameProduct(arr2))

	time.Sleep(5 * time.Second)
}
