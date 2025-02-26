package main

import (
	"fmt"
)

func reachableNodes(n int, edges [][]int, restricted []int) int {
	// Step 1: Build the adjacency list
	adj := make(map[int][]int)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// Step 2: Mark restricted nodes
	restrictedMap := make(map[int]bool)
	for _, node := range restricted {
		restrictedMap[node] = true
	}

	// Step 3: Perform BFS
	visited := make(map[int]bool)
	queue := []int{0}
	visited[0] = true
	count := 0

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		count++

		// Explore neighbors
		for _, neighbor := range adj[node] {
			if !visited[neighbor] && !restrictedMap[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return count
}

func main() {
	// Example 1
	n1 := 7
	edges1 := [][]int{{0, 1}, {1, 2}, {3, 1}, {4, 0}, {0, 5}, {5, 6}}
	restricted1 := []int{4, 5}
	fmt.Println(reachableNodes(n1, edges1, restricted1)) // Output: 4

	// Example 2
	n2 := 7
	edges2 := [][]int{{0, 1}, {0, 2}, {0, 5}, {0, 4}, {3, 2}, {6, 5}}
	restricted2 := []int{4, 2, 1}
	fmt.Println(reachableNodes(n2, edges2, restricted2)) // Output: 3
}
