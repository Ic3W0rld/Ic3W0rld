package main

import (
	"fmt"
)

func longestMonotonicSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxInc, maxDec := 1, 1
	currentInc, currentDec := 1, 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			currentInc++
			currentDec = 1
		} else if nums[i] < nums[i-1] {
			currentDec++
			currentInc = 1
		} else {
			currentInc, currentDec = 1, 1
		}

		if currentInc > maxInc {
			maxInc = currentInc
		}
		if currentDec > maxDec {
			maxDec = currentDec
		}
	}

	if maxInc > maxDec {
		return maxInc
	}
	return maxDec
}

func main() {
	arr1 := []int{1, 4, 3, 3, 2}
	fmt.Println(longestMonotonicSubarray(arr1))

	arr2 := []int{3, 3, 3, 3}
	fmt.Println(longestMonotonicSubarray(arr2))

	arr3 := []int{3, 2, 1}
	fmt.Println(longestMonotonicSubarray(arr3))
}
