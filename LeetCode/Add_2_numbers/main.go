package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/trace"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sumVal := carry
		if l1 != nil {
			sumVal += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sumVal += l2.Val
			l2 = l2.Next
		}

		carry = sumVal / 10
		current.Next = &ListNode{Val: sumVal % 10}
		current = current.Next
	}

	return dummy.Next
}

// Helper function to convert a slice of integers to a linked list
func sliceToListNode(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return dummy.Next
}

// Helper function to print a linked list
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
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

	l1 := sliceToListNode([]int{2, 4, 3})
	l11 := sliceToListNode([]int{5, 6, 4})
	result1 := addTwoNumbers(l1, l11)
	fmt.Print("Example 1: ")
	printList(result1) // Output: 7 0 8

	l2 := sliceToListNode([]int{0})
	l21 := sliceToListNode([]int{0})
	result2 := addTwoNumbers(l2, l21)
	fmt.Print("Example 2: ")
	printList(result2) // Output: 8 9 9 9 0 0 0 1

	l3 := sliceToListNode([]int{9, 9, 9, 9, 9, 9, 9})
	l31 := sliceToListNode([]int{9, 9, 9, 9})
	result3 := addTwoNumbers(l3, l31)
	fmt.Print("Example 3: ")
	printList(result3) // Output: 8 9 9 9 0 0 0 1

	time.Sleep(5 * time.Second)
}
