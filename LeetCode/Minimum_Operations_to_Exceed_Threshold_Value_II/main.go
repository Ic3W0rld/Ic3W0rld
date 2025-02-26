package main

import (
	"container/heap"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/trace"
	"time"
)

// intHeap is a min-heap implementation for integers
type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minOperations(nums []int, k int) int {
	// Define a min-heap
	h := &intHeap{}
	heap.Init(h)

	// Push all elements into the heap
	for _, num := range nums {
		heap.Push(h, num)
	}

	operations := 0
	for h.Len() >= 2 && (*h)[0] < k {
		// Pop the two smallest elements
		x := heap.Pop(h).(int)
		y := heap.Pop(h).(int)

		// Combine them using the given formula
		newNum := min(x, y)*2 + max(x, y)

		// Push the new element back into the heap
		heap.Push(h, newNum)

		// Increment the operation count
		operations++
	}

	// If the smallest element is still less than k, it's impossible to satisfy the condition
	if h.Len() > 0 && (*h)[0] < k {
		return -1 // This case should not happen as per the problem constraints
	}

	return operations
}

// Helper functions to find min and max
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
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

	arr1 := []int{2, 11, 10, 1, 3}
	k1 := 10
	fmt.Println(minOperations(arr1, k1))

	arr2 := []int{1, 1, 2, 4, 9}
	k2 := 20
	fmt.Println(minOperations(arr2, k2))

	time.Sleep(5 * time.Second)
}
