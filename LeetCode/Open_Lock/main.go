package main

import (
	"fmt"
)

func openLock(deadends []string, target string) int {
	deadSet := make(map[string]bool)
	for _, d := range deadends {
		deadSet[d] = true
	}
	if deadSet["0000"] {
		return -1
	}
	if target == "0000" {
		return 0
	}
	visited := make(map[string]bool)
	for k := range deadSet {
		visited[k] = true
	}
	queue := []struct {
		combo string
		steps int
	}{{"0000", 0}}
	visited["0000"] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for i := 0; i < 4; i++ {
			for _, delta := range []int{1, -1} {
				next := []byte(current.combo)
				digit := int(next[i] - '0')
				newDigit := (digit + delta + 10) % 10 // Ensure non-negative
				next[i] = byte(newDigit) + '0'
				nextStr := string(next)

				if nextStr == target {
					return current.steps + 1
				}

				if !visited[nextStr] {
					visited[nextStr] = true
					queue = append(queue, struct {
						combo string
						steps int
					}{nextStr, current.steps + 1})
				}
			}
		}
	}

	return -1
}

func main() {
	arr1 := []string{"0201", "0101", "0102", "1212", "2002"}
	target1 := "0202"
	fmt.Println(openLock(arr1, target1))

	arr2 := []string{"8888"}
	target2 := "0009"
	fmt.Println(openLock(arr2, target2))

	arr3 := []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}
	target3 := "8888"
	fmt.Println(openLock(arr3, target3))
}
