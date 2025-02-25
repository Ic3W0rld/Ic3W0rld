package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/trace"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FindElements struct {
	values map[int]bool // Use a map to store all values in the tree
}

func Constructor(root *TreeNode) FindElements {
	fe := FindElements{
		values: make(map[int]bool),
	}
	fe.recoverTree(root, 0) // Start recovering the tree from the root with value 0
	return fe
}

// Helper function to recover the tree and populate the values map
func (this *FindElements) recoverTree(node *TreeNode, value int) {
	if node == nil {
		return
	}
	node.Val = value          // Assign the correct value to the current node
	this.values[value] = true // Add the value to the map
	// Recover left and right children
	this.recoverTree(node.Left, 2*value+1)
	this.recoverTree(node.Right, 2*value+2)
}

func (this *FindElements) Find(target int) bool {
	// Check if the target exists in the values map
	return this.values[target]
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */

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

	// Example 1
	root1 := &TreeNode{Val: -1, Right: &TreeNode{Val: -1}}
	findElements1 := Constructor(root1)
	fmt.Println(findElements1.Find(1)) // Output: false
	fmt.Println(findElements1.Find(2)) // Output: true

	// Example 2
	root2 := &TreeNode{
		Val: -1,
		Left: &TreeNode{
			Val:   -1,
			Left:  &TreeNode{Val: -1},
			Right: &TreeNode{Val: -1},
		},
		Right: &TreeNode{Val: -1},
	}
	findElements2 := Constructor(root2)
	fmt.Println(findElements2.Find(1)) // Output: true
	fmt.Println(findElements2.Find(3)) // Output: true
	fmt.Println(findElements2.Find(5)) // Output: false

	// Example 3
	root3 := &TreeNode{
		Val: -1,
		Right: &TreeNode{
			Val:   -1,
			Left:  &TreeNode{Val: -1},
			Right: &TreeNode{Val: -1},
		},
	}
	findElements3 := Constructor(root3)
	fmt.Println(findElements3.Find(2)) // Output: true
	fmt.Println(findElements3.Find(3)) // Output: false
	fmt.Println(findElements3.Find(4)) // Output: false
	fmt.Println(findElements3.Find(5)) // Output: true
}
