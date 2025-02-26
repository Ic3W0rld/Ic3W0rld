package main

import "fmt"

func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	var diffIndices []int
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diffIndices = append(diffIndices, i)
		}
	}

	if len(diffIndices) != 2 {
		return false
	}

	i, j := diffIndices[0], diffIndices[1]
	return s1[i] == s2[j] && s1[j] == s2[i]
}

func main() {
	l1 := "bank"
	l2 := "kanb"
	fmt.Println(areAlmostEqual(l1, l2))

	l3 := "attack"
	l4 := "defend"
	fmt.Println(areAlmostEqual(l3, l4))

	l5 := "kelb"
	l6 := "kelb"
	fmt.Println(areAlmostEqual(l5, l6))
}
