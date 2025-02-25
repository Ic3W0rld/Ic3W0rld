package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/trace"
	"time"
)

func isArraySpecial(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	for i := 0; i < len(nums)-1; i++ {
		if (nums[i] % 2) == (nums[i+1] % 2) {
			return false
		}
	}
	return true
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

	arr1 := []int{1}
	fmt.Println(isArraySpecial(arr1))

	arr2 := []int{2, 1, 4}
	fmt.Println(isArraySpecial(arr2))

	arr3 := []int{4, 3, 1, 6}
	fmt.Println(isArraySpecial(arr3))

	time.Sleep(5 * time.Second)
}
