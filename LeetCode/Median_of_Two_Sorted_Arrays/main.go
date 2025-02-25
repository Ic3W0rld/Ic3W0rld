package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"runtime/trace"
	"time"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Ensure nums1 is the smaller array
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	m, n := len(nums1), len(nums2)
	low, high := 0, m

	for low <= high {
		i := (low + high) / 2
		j := (m+n+1)/2 - i

		if j < 0 {
			high = i - 1
			continue
		}
		if j > n {
			low = i + 1
			continue
		}

		var aLeft, aRight, bLeft, bRight int

		if i == 0 {
			aLeft = math.MinInt32
		} else {
			aLeft = nums1[i-1]
		}

		if i == m {
			aRight = math.MaxInt32
		} else {
			aRight = nums1[i]
		}

		if j == 0 {
			bLeft = math.MinInt32
		} else {
			bLeft = nums2[j-1]
		}

		if j == n {
			bRight = math.MaxInt32
		} else {
			bRight = nums2[j]
		}

		if aLeft <= bRight && bLeft <= aRight {
			if (m+n)%2 == 1 {
				return float64(max(aLeft, bLeft))
			} else {
				return float64(max(aLeft, bLeft)+min(aRight, bRight)) / 2.0
			}
		} else if aLeft > bRight {
			high = i - 1
		} else {
			low = i + 1
		}
	}
	return 0.0
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

	arr1 := []int{1, 2}
	arr11 := []int{2}
	fmt.Println(findMedianSortedArrays(arr1, arr11))

	arr2 := []int{1, 2}
	arr21 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(arr2, arr21))

	time.Sleep(5 * time.Second)
}
