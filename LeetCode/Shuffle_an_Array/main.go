package main

import (
	"fmt"
)

func shuffle(nums []int, n int) []int {
	result := make([]int, 0, 2*n) // Preallocate memory for the result

	for i := 0; i < n; i++ {
		result = append(result, nums[i])   // Add x_i
		result = append(result, nums[n+i]) // Add y_i
	}

	return result
}

func main() {
	arr1 := []int{2, 5, 1, 3, 4, 7}
	n1 := 3
	fmt.Println(shuffle(arr1, n1))

	arr2 := []int{1, 2, 3, 4, 4, 3, 2, 1}
	n2 := 4
	fmt.Println(shuffle(arr2, n2))

	arr3 := []int{1, 1, 2, 2}
	n3 := 2
	fmt.Println(shuffle(arr3, n3))
}
